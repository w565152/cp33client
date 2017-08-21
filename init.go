package main

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-redis/redis"
)

func redisInit() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := redisClient.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

var (
	Db *pg.DB
)

func Database() *pg.DB {
	return Db
}

func createSchema(db *pg.DB) error {
	return nil
}

func dbInit() {
	Db = pg.Connect(&pg.Options{
		Network:            "tcp",
		Addr:               fmt.Sprintf("%s:%s", "127.0.0.1", "5432"),
		User:               "root",
		Password:           "root",
		Database:           "cp33",
		DialTimeout:        3 * time.Second,
		ReadTimeout:        3 * time.Second,
		WriteTimeout:       3 * time.Second,
		PoolSize:           99,
		PoolTimeout:        time.Second,
		IdleTimeout:        time.Second,
		IdleCheckFrequency: time.Second,
	})

	err := createSchema(Db)
	if err != nil {
		fmt.Println(err.Error())
		t, _ := time.ParseDuration("5s")
		time.Sleep(t)
		dbInit()
	}
}

func init() {
	redisInit()
	dbInit()
}