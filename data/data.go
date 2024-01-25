package data

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/KindOf/itsuxel/data/evaluator"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

func ConnectStorage() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
}

type (
	ValueResolver struct {
		sheet, cell string
	}

	CellValue struct {
		Value  string
		Result any
	}
)

func (vr *ValueResolver) Resolve(ident string) (any, error) {
	ctx := context.Background()
	cmd := RedisClient.HGet(ctx, vr.sheet, ident)
	if err := cmd.Err(); err != nil {
		return nil, err
	}
	intValue, err := strconv.Atoi(cmd.Val())
	if err != nil {
		return cmd.Val(), nil
	}
	return intValue, nil
}

func SetCell(sheet, cell, value string) error {
	ctx := context.Background()
	s := RedisClient.HSet(ctx, sheet, strings.ToUpper(cell), value)

	if err := s.Err(); err != nil {
		return err
	}

	return nil
}

func GetCell(sheet, cell string) (CellValue, error) {
	ctx := context.Background()

	cmd := RedisClient.HGet(ctx, sheet, strings.ToUpper(cell))

	val, err := cmd.Result()
	if err != nil {
		return CellValue{}, err
	}

	if isExp(val) {
		env, err := evaluator.GetEnv(val, &ValueResolver{sheet, strings.ToUpper(cell)}, 0)
		if err != nil {
			return CellValue{}, err
		}
		res, err := evaluator.Evaluate(val[1:], env)
		if err != nil {
			return CellValue{}, err
		}

		return CellValue{val, res}, err
	}

	return CellValue{val, val}, nil
}

func GetSheet(sheet string) (map[string]CellValue, error) {
	ctx := context.Background()
	cmd := RedisClient.HGetAll(ctx, sheet)

	r := map[string]CellValue{}

	m, err := cmd.Result()
	if err != nil {
		return map[string]CellValue{}, err
	}

	for k, v := range m {
		if isExp(v) {
			env, err := evaluator.GetEnv(v, &ValueResolver{sheet, k}, 0)
			if err != nil {
                fmt.Println("ERROR", err)
				r[k] = CellValue{v, "Can't evaluate"}
				continue
			}
            res, err := evaluator.Evaluate(v[1:], env)
			if err != nil {
                fmt.Println("ERROR", err)
				r[k] = CellValue{v, "Can't evaluate"}
				continue
			}

			r[k] = CellValue{v, res}
		} else {
			r[k] = CellValue{v, v}
		}
	}

	return r, nil
}

func isExp(val string) bool {
	return string(val[0]) == "="
}
