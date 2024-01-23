package data

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
    RedisClient *redis.Client
)

func ConnectStorage() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr: ":6379",
    })
}

func SetCell(sheet, cell, value string) error {
    ctx := context.Background()
    s := RedisClient.HSet(ctx, sheet, cell, value)

    if isExp(value) {
        fmt.Println("Value is expression")
    }

    if err := s.Err(); err != nil {
        return err
    }

    fmt.Println(s.Val())
    return nil
}

func GetCell(sheet, cell string) {
//    ctx := context.Background()
}

func GetSheet(sheet string) {
    //res := RedisClient.HGetAll(ctx, sheet)
}

func isExp(val string) bool {
    return string(val[0]) == "="
}
