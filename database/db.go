package db

import (
	"encoding/json"
	"strconv"
	"team1_qgame/conf"
	"github.com/go-redis/redis/v8"
)

var storage = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

var ctx = storage.Context()

func SaveUser(user *conf.User) {
	j, _ := json.Marshal(user)
	err := storage.Set(ctx, strconv.Itoa(int(user.Id)), string(j), 0).Err()
	if err != nil {
		panic(err)
	}
}

func GetUser(id string) conf.User {
	u, err := storage.Get(ctx, id).Result()
	if err != nil {
		panic(err)
	}
	var user conf.User

	json.Unmarshal([]byte(u), &user)

	return user
}