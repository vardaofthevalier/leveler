package leveler

import (
	"regexp"
	redis "github.com/mediocregopher/radix.v2/redis"
	redis_pool "github.com/mediocregopher/radix.v2/pool"
)

type Database interface {
	Create()
	Get()
	List()
	Update()
	Delete()
	Flush()
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

func NewRedisDatabaseRangeQueryConfig (t interface{}) *RedisDatabaseRangeQueryConfig, error {
	rangeQueryNumerical := RedisDatabaseRangeQueryConfig{
		Command: "ZRANGEBYSCORE",
		RangeqBound: func(boundType string, x string, suffix string) string { return x },
		RangeqInclusive: "",
		RangeqExclusive: "",
		RangeqPosInf: "+inf",
		RangeqNegInf: "-inf",
		RangeqSpecial: ""
	}

	rangeQueryLexical := RedisDatabaseRangeQueryConfig{
		Command: "ZRANGEBYLEX",
		RangeqBound: func(bType string, x string, suffix string) string { return fmt.Sprintf("%s%s%s", bType, x, suffix) },
		RangeqInclusive: "[",
		RangeqExclusive: "]",
		RangeqPosInf: "+",
		RangeqNegInf: "-",
		RangeqSpecial: "\\xff\\xff\\xff\\xff"
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
		return &RedisDatabaseRangeQueryConfig{}, QueryError(fmt.Sprintf("Can't perform range query on type '%s'", typ))
	}
}

type RedisDatabase struct {
	DatabaseConnectionPool redis_pool.Pool
}

const (
	action = iota
	requirement
	role
	host
)

const (
	conditional := []string{"==", "!=", ">=", "<=", ">", "<", "IN"}
	boolean := []string{"AND", "OR", "NOT"}
	tokenRegex := regexp.MustCompile(fmt.Sprintf("([a-zA-Z0-9_\-,.@]+|[\(\)]|%s)", strings.Join(conditional, "|")))
)

type QueryError struct {
	Message string
}

type TokenSplitError struct {}

func (e QueryError) Error() string {
	return fmt.Sprintf(e.Message)
}

func (db *RedisDatabase) selectDatabase(kind string) *redis.Client {
	conn, err := db.Get()
	if err != nil {
		// TODO
	}

	switch kind {
	case "action":
		conn.Cmd("select", action)
	case "requirement": 
		conn.Cmd("select", requirement)
	case "role":
		conn.Cmd("select", requirement)
	case "host":
		comm.Cmd("select", host)
	}

	return conn
}

func (db *RedisDatabase) executeQuery(q string, kind string, conn *redis.Client) {
	func tokenize(s string) {
		return tokenRegex.Split(q, -1)
	}

	func splitTokens(tokens []string) (string, []string, error) {
		if len(tokens) > 2 {
			return tokens[0], tokens[1:], nil
		} else if len(tokens) == 1 {
			return tokens[0], [], nil
		} else {
			return &TokenSplitError{}
		}
	}

	func evaluateConditionalExpression(head string, tail[]string) (*redis.Resp, []string, error) {
		op, tail, err := splitTokens(tail)
		if err != nil {
			// TODO
		}

		test, tail, err := splitTokens(tail)
		if err != nil {

		}

		rangeqName := fmt.Sprintf("%s.%s.index", kind, head)
		rangeqConfig := GetRedisDatabaseRangeQueryConfig(test)

		switch op {
		case "==":
			return conn.Cmd(rangeqConfig.Command, rangeqName, rangeqConfig.RangeqBound(rangeqConfig.RangeqInclusive, test, ""), rangeqConfig.RangeqBound(rangeqConfig.RangeqInclusive, test, rangeqConfig.RangeqSpecial)), tail
		case ">=":
			return conn.Cmd(rangeqConfig.Command, rangeqName, rangeqConfig.RangeqBound(rangeqConfig.RangeqInclusive, test, ""), rangeqConfig.RangeqPosInf)), tail
		case ">": 
			return conn.Cmd(rangeqConfig.Command, rangeqName, rangeqConfig.RangeqBound(rangeqConfig.RangeqExclusive, test, ""), rangeqConfig.RangeqPosInf)), tail
		case "<=":
			return conn.Cmd(rangeqConfig.Command, rangeqName, rangeqConfig.RangeqNegInf, rangeqConfig.RangeqBound(rangeqConfig.RangeqInclusive, test, "")), tail
		case "<":
			return conn.Cmd(rangeqConfig.Command, rangeqName, rangeqConfig.RangeqNegInf, rangeqConfig.RangeqBound(rangeqConfig.RangeqExclusive, test, "")), tail
		default:
			return *redis.Resp, [], QueryError{fmt.Sprintf("Invalid conditional operator '%s'", op)}
		}

	}

	func evaluateSetExpression(tokens []string, previous []string) (*redis.Resp, error) {
		head, tail, err := splitTokens(tokens)
		if err != nil {
			return previous_result
		}

		if head == "(" {
			return evaluateSetExpression(tail, previous)
		} else if head == ")" {
			if len(tail) > 0 {
				return evaluateSetExpression(tail, previous)
			} else {
				return previous
			}
		} else if in(head, boolean) {
			if head == "AND" {
				head, tail, err := splitTokens(tail)
				if err != nil {
					// TODO: malformed query
				}
				if head == "NOT" {
					rh, rem, err := evaluateConditionalExpression(head, tail)
					if err != nil {
						return rh, err
					}
					return evaluateSetExpression(rem, intersection(previous, complement(rh)))
				} else if head == "(" {
					rh = evaluateSetExpression(tail, [])
					return intersection(rh, previous)
				} else {
					rh, rem, err := evaluateConditionalExpression(tail)
					if err != nil {
						return rh, err
					}
					return evaluateSetExpression(rem, intersection(previous, rh))
				}
			} else {
				head, tail, err := splitTokens(tail)
				if err != nil {
					// TODO: malformed query
				}
				if head == "NOT" {
					rh, rem, err := evaluateConditionalExpression(head, tail)
					if err != nil {
						return rh, err
					}
					return evaluateSetExpression(rem, union(previous, complement(rh)))
				} else if head == "(" {
					rh = evaluateSetExpression(tail, [])
					return intersection(rh, previous)
				} else {
					rh, rem, err := evaluateConditionalExpression(tail)
					if err != nil {
						return rh, err
					}
					return evaluateSetExpression(rem, union(previous, rh))
			}
		} else {
			head, tail, err := splitTokens(tail)
			if err != nil {
				// TODO
			}
			lh, rem, err := evaluateConditionalExpression(head, tail)
			if err != nil {
				return lh, err
			}
			return evaluateSetExpression(rem, lh)
		}
	}

	func complement() {

	}

	func intersection() {
		
	}

	func union() {
		
	}

	func in(elem string, slice []string) bool {
		for _, v := range slice {
			if v == elem {
				return true
			}
		}

		return false
	}

	tokenized := tokenize(q)

	head, tail, err := splitTokens(tokenized)
	if err != nil {
		fmt.Println("Invalid query")
		os.Exit(1)
	}

	return evaluateSetExpression(head, tail, [])
}

func (db *RedisDatabase) Create(kind string, obj map[string]interface{}) error {
	conn := db.selectDatabase(kind)
	defer db.Put(conn)

	// create a new hash in the database
	// create secondary keys, if applicable
}

func (db *RedisDatabase) Get(kind string, id string) error {
	conn := db.selectDatabase(kind)
	defer db.Put(conn)

	rawResult, err := conn.Cmd("HGETALL", fmt.Sprintf("%s:%s", kind, id))
	if err != nil {
		return map[string][string]{}, err 
	}

	result, err2 := result.Map()
	if err2 != nil {
		return map[string][string]{}, err2
	}

	return result, nil
}

func (db *RedisDatabase) List(kind string, query string) error {
	conn := db.selectDatabase(kind)
	defer db.Put(conn)

	if len(query) > 0 {
		rawResult, err := db.executeQuery(query, kind, conn)
	} else {
		rawResult, err := conn.Cmd("")
	}

	if err != nil {
		return rawResult, err
	}

	result, err2 := result.Map()
	if err2 != nil {
		return map[string][string]{}, err2
	}

	return result, nil
}

func (db *RedisDatabase) Update(kind string, id string, obj map[string]interface{}) error {
	conn := db.selectDatabase(kind)
	defer db.Put(conn)

	// full replace update on obj
}

func (db *RedisDatabase) Delete(kind string, id string) error {
	conn := db.selectDatabase(kind)
	defer db.Put(conn)

	// delete the hash stored at key = id
}

func (db *RedisDatabase) Flush() error {
	conn := db.DatabaseConnectionPool.Get()
	defer db.Put(conn)

	// Flush all contents of the database, i.e. FLUSHALL
	// Question: how should this be exposed?  Probably should require authentication!
}

func NewRedisDatabase(protocol string, host string, port int, size int) RedisDatabase {
	pool, err := redis_pool.New(protocol, fmt.Sprintf("%s:%d", host, port), size)  
	if err != nil {
		log.Fatalf("Couldn't connect to Redis server: %s", err)
	}

	return RedisDatabase{DatabaseConnectionPool: *pool}
}