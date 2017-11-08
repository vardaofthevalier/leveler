package leveler

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	config "leveler/config"
	uuid "github.com/satori/go.uuid"
	redis "github.com/mediocregopher/radix.v2/redis"
	redis_pool "github.com/mediocregopher/radix.v2/pool"
)

type RedisDatabase struct {
	DatabaseConnectionPool redis_pool.Pool
	DatabaseLabels map[string]int
}

type RedisDatabaseRangeQueryConfig struct {
	Command string
	RangeqBound func(bType string, x string, suffix string) string
	RangeqInclusive string
	RangeqExclusive string
	RangeqPosInf string
	RangeqNegInf string
	RangeqSpecial string
}

func GetRedisDatabaseRangeQueryConfig (t interface{}) (*RedisDatabaseRangeQueryConfig, error) {
	rangeQueryNumerical := RedisDatabaseRangeQueryConfig{
		Command: "ZRANGEBYSCORE",
		RangeqBound: func(boundType string, x string, suffix string) string { return x },
		RangeqInclusive: "",
		RangeqExclusive: "",
		RangeqPosInf: "+inf",
		RangeqNegInf: "-inf",
		RangeqSpecial: "",
	}

	rangeQueryLexical := RedisDatabaseRangeQueryConfig{
		Command: "ZRANGEBYLEX",
		RangeqBound: func(bType string, x string, suffix string) string { return fmt.Sprintf("%s%s%s", bType, x, suffix) },
		RangeqInclusive: "[",
		RangeqExclusive: "]",
		RangeqPosInf: "+",
		RangeqNegInf: "-",
		RangeqSpecial: "\\xff\\xff\\xff\\xff",
	}

	switch typ := t.(type) {
	case string:
		return &rangeQueryLexical, nil
	case int:
		return &rangeQueryNumerical, nil
	case int8:
		return &rangeQueryNumerical, nil
	case int16: 
		return &rangeQueryNumerical, nil
	case int32:
		return &rangeQueryNumerical, nil
	case int64:
		return &rangeQueryNumerical, nil
	case uint:
		return &rangeQueryNumerical, nil
	case uint8:
		return &rangeQueryNumerical, nil
	case uint16:
		return &rangeQueryNumerical, nil
	case uint32:
		return &rangeQueryNumerical, nil
	case uint64:
		return &rangeQueryNumerical, nil
	case float32:
		return &rangeQueryNumerical, nil
	case float64:
		return &rangeQueryNumerical, nil
	default:
		return &RedisDatabaseRangeQueryConfig{}, QueryError{fmt.Sprintf("Can't perform range query on type '%s'", typ)}
	}
}

var conditional = []string{"==", "!=", ">=", "<=", ">", "<", "IN"}
var boolean = []string{"AND", "OR", "NOT"}
var tokenRegex = regexp.MustCompile(fmt.Sprintf("([a-zA-Z0-9_\\-,.@]+|[\\(\\)]|%s)", strings.Join(conditional, "|")))


type QueryError struct {
	Message string
}

type TokenSplitError struct {}

func (e QueryError) Error() string {
	return fmt.Sprintf(e.Message)
}

func (e TokenSplitError) Error() string {
	return "No tokens left to split"
}

func tokenize(s string) []string {
	return tokenRegex.Split(s, -1)
}

func splitTokens(tokens []string) (string, []string, error) {
	if len(tokens) > 2 {
		return tokens[0], tokens[1:], nil

	} else if len(tokens) == 1 {
		return tokens[0], []string{}, nil

	} else {
		return "", []string{}, &TokenSplitError{}
	}
}

func evaluateConditionalExpression(kind string, conn *redis.Client, head string, tail[]string) (map[string]string, []string, error) {
	op, tail, err := splitTokens(tail)
	if err != nil {
		return map[string]string{}, []string{}, QueryError{fmt.Sprintf("Malformed query")}
	}

	test, tail, err := splitTokens(tail)
	if err != nil {
		return map[string]string{}, []string{}, QueryError{fmt.Sprintf("Malformed query")}
	}

	rangeqName := fmt.Sprintf("%s.%s.index", kind, head)
	rangeqConfig, ok := GetRedisDatabaseRangeQueryConfig(test)
	if ok != nil {
		return map[string]string{}, []string{}, QueryError{fmt.Sprintf("Unknown type for range query")}
	}

	var result map[string]string

	switch op {
	case "==":
		result, err = conn.Cmd(rangeqConfig.Command, rangeqName, rangeqConfig.RangeqBound(rangeqConfig.RangeqInclusive, test, ""), rangeqConfig.RangeqBound(rangeqConfig.RangeqInclusive, test, rangeqConfig.RangeqSpecial)).Map()
	
	case ">=":
		result, err = conn.Cmd(rangeqConfig.Command, rangeqName, rangeqConfig.RangeqBound(rangeqConfig.RangeqInclusive, test, ""), rangeqConfig.RangeqPosInf).Map()
	
	case ">": 
		result, err = conn.Cmd(rangeqConfig.Command, rangeqName, rangeqConfig.RangeqBound(rangeqConfig.RangeqExclusive, test, ""), rangeqConfig.RangeqPosInf).Map()
	
	case "<=":
		result, err = conn.Cmd(rangeqConfig.Command, rangeqName, rangeqConfig.RangeqNegInf, rangeqConfig.RangeqBound(rangeqConfig.RangeqInclusive, test, "")).Map()
	
	case "<":
		result, err = conn.Cmd(rangeqConfig.Command, rangeqName, rangeqConfig.RangeqNegInf, rangeqConfig.RangeqBound(rangeqConfig.RangeqExclusive, test, "")).Map()
	
	default:
		return map[string]string{}, []string{}, QueryError{fmt.Sprintf("Invalid conditional operator '%s'", op)}
	}

	if err != nil {
		return map[string]string{}, []string{}, err
	}

	return result, tail, nil

}

func evaluateSetExpression(kind string, conn *redis.Client, tokens []string, previous map[string]string) (map[string]string, error) {
	head, tail, err := splitTokens(tokens)
	if err != nil {
		return previous, nil
	}

	if head == "(" {
		return evaluateSetExpression(kind, conn, tail, previous)

	} else if head == ")" {
		if len(tail) > 0 {
			return evaluateSetExpression(kind, conn, tail, previous)

		} else {
			return previous, nil

		}
	} else if in(head, boolean) {
		if head == "AND" {
			head, tail, err := splitTokens(tail)
			if err != nil {
				return map[string]string{}, QueryError{fmt.Sprintf("Malformed query")}
			}

			if head == "NOT" {
				rh, rem, err := evaluateConditionalExpression(kind, conn, head, tail)
				if err != nil {
					return rh, err
				}

				c, err := complement(rh, conn)
				if err != nil {
					return c, err
				}
				return evaluateSetExpression(kind, conn, rem, intersection(c, previous))

			} else if head == "(" {
				rh, err := evaluateSetExpression(kind, conn, tail, previous)
				if err != nil {
					return rh, err
				}

				return intersection(rh, previous), nil

			} else {
				rh, rem, err := evaluateConditionalExpression(kind, conn, head, tail)
				if err != nil {
					return rh, err
				}

				return evaluateSetExpression(kind, conn, rem, intersection(rh, previous))
			}
		} else {
			head, tail, err := splitTokens(tail)
			if err != nil {
				return map[string]string{}, QueryError{fmt.Sprintf("Malformed query")}
			}

			if head == "NOT" {
				rh, rem, err := evaluateConditionalExpression(kind, conn, head, tail)
				if err != nil {
					return rh, err
				}

				c, err := complement(rh, conn)
				if err != nil {
					return c, err
				}

				return evaluateSetExpression(kind, conn, rem, union(c, previous))

			} else if head == "(" {
				rh, err := evaluateSetExpression(kind, conn, tail, previous)
				if err != nil {
					return rh, err
				}

				return intersection(rh, previous), nil

			} else {
				rh, rem, err := evaluateConditionalExpression(kind, conn, head, tail)
				if err != nil {
					return rh, err
				}

				return evaluateSetExpression(kind, conn, rem, union(rh, previous))
			}	
		}
	} else {
		head, tail, err := splitTokens(tail)
		if err != nil {
			return map[string]string{}, QueryError{fmt.Sprintf("Malformed query")}
		}

		lh, rem, err := evaluateConditionalExpression(kind, conn, head, tail)
		if err != nil {
			return lh, err
		}

		return evaluateSetExpression(kind, conn, rem, lh)
	}
}

func complement(a map[string]string, conn *redis.Client) (map[string]string, error) {
	var results map[string]string
	u, err := conn.Cmd("KEYS", "*").Array()
	if err != nil {
		return results, err
	}

	for _, k := range u {
		s := k.String()
		if _, ok := a[s]; !ok {
			r, err := conn.Cmd("HGETALL", s).Map()
			if err != nil {
				return map[string]string{}, err
			}
			for w, v := range r {
				results[w] = v
			}
		}
	}

	return results, nil
}

func intersection(a map[string] string, b map[string]string) map[string] string {
	var results map[string]string
	var smaller map[string]string
	var larger map[string]string

	if len(a) > len(b) {
		smaller = b
		larger = a
	} else {
		smaller = a
		larger = b
	}

	for k, v := range smaller {
		if _, ok := larger[k]; ok {
			results[k] = v
		} 
	}

	return results
}

func union(a map[string] string, b map[string]string) map[string] string {
	var results map[string]string
	var smaller map[string]string
	var larger map[string]string

	if len(a) > len(b) {
		smaller = b
		larger = a
	} else {
		smaller = a
		larger = b
	}

	for k, _ := range smaller {
		if exists, ok := larger[k]; !ok {
			results[k] = exists
		}
	}

	return results
}

func in(elem string, slice []string) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}

	return false
}

func (db *RedisDatabase) selectDatabase(kind string) (*redis.Client, error) {
	conn, err := db.DatabaseConnectionPool.Get()
	if err != nil {
		return conn, err
	}

	conn.Cmd("select", db.DatabaseLabels[kind])

	return conn, nil
}

func (db *RedisDatabase) executeQuery(q string, kind string, conn *redis.Client) (map[string]string, error) {
	return evaluateSetExpression(kind, conn, tokenize(q), map[string]string{})
}

func (db RedisDatabase) Create(kind string, obj map[string]interface{}) (string, error) {
	conn, err := db.selectDatabase(kind)
	if err != nil {
		return "", err
	}

	defer db.DatabaseConnectionPool.Put(conn)

	// generate a uuid for the hash key
	id := uuid.NewV4().String()

	// create a new hash in the database
	_ = conn.Cmd("HMSET", fmt.Sprintf("%s:%s", kind, id), obj).String()

	// create secondary keys, if applicable

	return id, nil
}

func (db RedisDatabase) Get(kind string, id string) (map[string]string, error) {
	conn, err := db.selectDatabase(kind)
	if err != nil {
		return map[string]string{}, err
	}

	defer db.DatabaseConnectionPool.Put(conn)

	result, err := conn.Cmd("HGETALL", fmt.Sprintf("%s:%s", kind, id)).Map()
	log.Println(result)
	if err != nil {
		return map[string]string{}, err 
	}

	return result, nil
}

func (db RedisDatabase) List(kind string, query string) (map[string]string, error) {
	conn, err := db.selectDatabase(kind)
	if err != nil {
		return map[string]string{}, err
	}

	defer db.DatabaseConnectionPool.Put(conn)

	var result map[string]string

	if len(query) > 0 {
		result, err = db.executeQuery(query, kind, conn)
	} else {
		result, err = conn.Cmd("GET", kind).Map()
	}

	if err != nil {
		return result, err
	}

	return result, nil
}

func (db RedisDatabase) Update(kind string, id string, obj map[string]interface{}) error {
	conn, err := db.selectDatabase(kind)
	if err != nil {
		return err
	}

	defer db.DatabaseConnectionPool.Put(conn)

	// full replace update on obj
	_ = conn.Cmd("HMSET", fmt.Sprintf("%s:%s", kind, id), obj).String()

	return err
}

func (db RedisDatabase) Delete(kind string, id string) error {
	conn, err := db.selectDatabase(kind)
	if err != nil {
		return err
	}

	defer db.DatabaseConnectionPool.Put(conn)

	// delete the hash stored at key = id
	_ = conn.Cmd("DEL", fmt.Sprintf("%s:%s", kind, id)).String()

	return err
}

func (db RedisDatabase) Flush(kind string) error {
	conn, err := db.selectDatabase(kind)
	if err != nil {
		return err
	}
	defer db.DatabaseConnectionPool.Put(conn)

	// if kind == "" flush all contents of the database, i.e. FLUSHALL
	if len(kind) == 0 {
		err := conn.Cmd("FLUSHALL").Err
		if err != nil {
			return err
		}
	} else {
		err := conn.Cmd("FLUSHDB", db.DatabaseLabels[kind]).Err
		if err != nil {
			return err
		}
	}

	// Question: how should this be exposed?  Probably should require authentication!
	return nil
}

func NewRedisDatabase(protocol string, host string, port int32, size int32) RedisDatabase {
	pool, err := redis_pool.New(protocol, fmt.Sprintf("%s:%d", host, port), int(size))  
	if err != nil {
		log.Fatalf("Couldn't connect to Redis server: %s", err)
	}

	return RedisDatabase{DatabaseConnectionPool: *pool, DatabaseLabels: config.DatabaseLabelsMap}
}