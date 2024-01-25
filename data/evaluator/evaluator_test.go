package evaluator

import (
	"fmt"
	"reflect"
	"testing"
)

type ResolverStub struct {
	storageStub map[string]any
}

func (r *ResolverStub) Resolve(ident string) (any, error) {
	return r.storageStub[ident], nil
}

func TestEvaluate(t *testing.T) {
	tests := []struct {
		expr string
		want any
		err  bool
	}{
		{"2+2", 4, false},
		{"3*(1+2)", 9, false},
		{"A1+1", 2, false},
		{"A1+A2", 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.expr, func(t1 *testing.T) {
			got, err := Evaluate(tt.expr, map[string]any{"A1": 1, "A2": 2})

			if tt.err && err == nil {
				t1.Fatalf("Error expected got nil")
			}

			if tt.want != got {
				t1.Fatalf("Got = %v, Want = %v", got, tt.want)
			}

		})
	}
}

func TestGetEnv(t *testing.T) {
	valueMap := map[string]any{"A1": 1, "A2": 2}
	tests := []struct {
		expr    string
		storage map[string]any
		want    map[string]any
		err     bool
	}{
		{"2+2", valueMap, map[string]any{}, false},
		{"A1+A2", valueMap, valueMap, false},
		{"A1+A2+A3", valueMap, map[string]any{"A1": 1, "A2": 2, "A3": nil}, false},
		{
			"A1+A2+A3",
			map[string]any{"A1": 1, "A2": 2, "A3": "=A2-A4", "A4": 4},
			map[string]any{"A1": 1, "A2": 2, "A3": -2, "A4": 4},
			false,
		},
        {
            "A1+A2",
            map[string]any{"A1": "=A2", "A2": "=A1"},
            map[string]any{},
            false,
        },
	}

	for _, tt := range tests {
		t.Run(tt.expr, func(t1 *testing.T) {
			resolver := &ResolverStub{tt.storage}
			got, err := GetEnv(tt.expr, resolver, 0)
			if err != nil && !tt.err {
				t1.Fatalf("Got error %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				fmt.Println(got, tt.want)
				t1.Fatalf("Got = %v; Want = %v", got, tt.want)
			}
		})
	}
}

func TestIsExpr(t *testing.T) {
	tests := []struct {
		name string
		val  any
		want string
		ok   bool
	}{
		{"A1", "A1", "A1", false},
		{"1", 1, "", false},
		{"=A1", "=A1", "=A1", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			got, ok := isExpr(tt.val)

			if tt.ok && !ok {
				t1.Fatalf("Expected %v to be an expression", tt.val)
			}

			if tt.want != got {
				t1.Fatalf("Got = %v; Want = %v", got, tt.want)
			}
		})
	}
}
