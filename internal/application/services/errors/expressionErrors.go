package expressionErros

import "fmt"

type ExpressionErrors struct {
	Message string `json:"message"`
}

func NewExpressionErrors(message string) *ExpressionErrors {
	return &ExpressionErrors{
		Message: message,
	}
}

func (e *ExpressionErrors) Error() string {
	return fmt.Sprintf("Expression Error: %s", e.Message)
}
