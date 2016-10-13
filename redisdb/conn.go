package redisdb

import (
	"os"

	"github.com/mediocregopher/radix.v2/redis"
)

var rdb *redis.Client

// Connect is a helper function to stablish a new Redis connection
func Connect() (*redis.Client, error) {
	if rdb != nil {
		return rdb, nil
	}

	var err error

	rdb, err = redis.Dial("tcp", os.Getenv("Redis_Addr"))

	if err != nil {
		panic(err)
	}

	return rdb, nil
}
