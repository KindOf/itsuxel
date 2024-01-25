package evaluator

import (
	"fmt"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/ast"
	"github.com/expr-lang/expr/parser"
	"github.com/labstack/gommon/log"
)

type (
	resolver interface {
		Resolve(ident string) (any, error)
	}

	envCollector struct {
		identifiers []string
	}
)

func (v *envCollector) Visit(n *ast.Node) {
	if n, ok := (*n).(*ast.IdentifierNode); ok {
		v.identifiers = append(v.identifiers, n.String())
	}
}

func (v *envCollector) GetIdentifiers() []string {
	return v.identifiers
}

func Evaluate(e string, env map[string]any) (any, error) {
    fmt.Println("EVAL", e, env)
	output, err := expr.Eval(e, env)
	if err != nil {
		log.Warn(err)
		return "", err
	}

	return output, nil
}

func GetEnv(expr string, r resolver, deps int) (map[string]any, error) {
	env := make(map[string]any)
    env["SUM"] = func(a, b int) int {
        return a + b
    }
    if deps > 10 {
        return env, fmt.Errorf("Circular dependencies")
    }

    tree, err := parser.Parse(expr[1:])
	if err != nil {
		return env, err
	}

	visitor := &envCollector{}
	ast.Walk(&tree.Node, visitor)

	for _, ident := range visitor.GetIdentifiers() {
		v, err := r.Resolve(ident)
        if err != nil && env[ident] == nil {
            env[ident] = nil
            continue
        }
        fmt.Println("RESOLVE", ident, v)

        if sv, ok := isExpr(v); ok {
            innerEnv, err := GetEnv(sv, r, deps + 1)
            if err != nil {
                return env, err
            }

            for k, v := range innerEnv {
                env[k] = v
            }
        }
		if err == nil {
			env[ident] = v
		} else {
            fmt.Println(err)
        }
	}

    for k, v := range env {
        if exp, ok := isExpr(v); ok {
            val, err := Evaluate(exp[1:], env);
            if err != nil {
                return env, err
            }
            env[k] = val
        }
    }

	return env, nil
}

func isExpr(v any) (string, bool) {
    val, ok := v.(string)

    if ok {
        if string(val[0]) == "=" {
            return val, true
        }
        return val, false
    }
    return "", false
}
