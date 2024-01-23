package evaluator

import "testing"

func TestEvaluate(t * testing.T) {
    got, err := Evaluate("2 + 2")
    want := "4"

    if err != nil {
        t.Fatalf("Got an error: %s", err)
    }

    if got != want {
        t.Fatalf("Got = %v, Want = %v", got, want)
    }
}
