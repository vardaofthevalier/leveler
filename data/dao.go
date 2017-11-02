package leveler

import (
	redis "github.com/mediocregopher/radix.v2/redis"
	redis_pool "github.com/mediocregopher/radix.v2/pool"
)

type Database interface {
	Create()
	Get()
	List()
	Update()
	Delete()
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

func (db *RedisDatabase) Create(kind string, obj map[string]interface{}) error {
	conn := db.selectDatabase(kind)
	defer db.Put(conn)

	// create a new hash in the database
	// create secondary keys, if applicable
}

func (db *RedisDatabase) Get(kind string, id string) error {
	conn := db.selectDatabase(kind)
	defer db.Put(conn)

	// return the hash stored at key = id
}

func (db *RedisDatabase) List(kind string, query string) error {
	conn := db.selectDatabase(kind)
	defer db.Put(conn)

	// if query is the empty string, return a full listing of all keys
	// otherwise, execute the query to filter results
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

func NewRedisDatabase(protocol string, host string, port int, size int) RedisDatabase {
	pool, err := redis_pool.New(protocol, fmt.Sprintf("%s:%d", host, port), size)  
	if err != nil {
		log.Fatalf("Couldn't connect to Redis server: %s", err)
	}

	return RedisDatabase{DatabaseConnectionPool: *pool}
}