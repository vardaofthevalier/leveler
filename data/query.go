package leveler

import (
	"os"
	"fmt"
)

const (
	conditional := []string{"==", "!=", ">=", "<=", ">", "<", "IN"}
	boolean := []string{"AND", "OR", "NOT"}
)

type QueryError struct {
	Message string
}

type TokenSplitError struct {}

func (e InvalidQueryError) Error() string {
	return fmt.Sprintf(e.Message)
}

func in(elem Type, slice []Type) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}

	return false
}

func executeQuery(q string, dataset d) {
	tokenized := tokenize(q)

	head, tail, err := splitTokens(tokenized)
	if err != nil {
		fmt.Println("Invalid query")
		os.Exit(1)
	}

	return evaluateExpression(head, tail)
}

func tokenize(s string) {

}

func splitTokens(tokens []string) (string, []string, error) {
	if len(tokens) > 2 {
		return tokens[0], tokens[1:], nil
	} else if len(tokens) == 1 {
		return tokens[0], [], nil
	} else {
		return &TokenSplitError{}
	}
}

func evaluateExpression(head string, tail []string) {
	if head == "(" {
		stuff
	} else {
		op, tail, err := splitTokens(tail)
		if err != nil {
			
		}
	}
}