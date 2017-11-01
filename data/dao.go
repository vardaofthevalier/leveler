package leveler

import (
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

func (db *RedisDatabase) Create() {

}

func (db *RedisDatabase) Get() {
	
}

func (db *RedisDatabase) List() {
	
}

func (db *RedisDatabase) Update() {
	
}

func (db *RedisDatabase) Delete() {
	
}