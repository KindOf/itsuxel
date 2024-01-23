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

func GetCell(sheet, cell string) (string, error) {
	ctx := context.Background()

	cmd := RedisClient.HGet(ctx, sheet, cell)

	val, err := cmd.Result()
	if err != nil {
		return "", err
	}

	if isExp(val) {
		fmt.Println("Value is expression")
	}

	return val, nil
}

func GetSheet(sheet string) (map[string]string, error) {
	ctx := context.Background()

	cmd := RedisClient.HGetAll(ctx, sheet)

    m, err := cmd.Result()
    if err != nil {
        return map[string]string{}, err
    }

    return m, nil
}

func isExp(val string) bool {
	return string(val[0]) == "="
}
