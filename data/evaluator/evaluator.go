package evaluator

import (
	"fmt"

	"github.com/expr-lang/expr"
)

func Evaluate(e string) (string, error) {
	program, err := expr.Compile(e)
	if err != nil {
        return "", err
	}

	output, err := expr.Run(program, nil)
	if err != nil {
        return "", err
	}

	return fmt.Sprintf("%v", output), nil
}
