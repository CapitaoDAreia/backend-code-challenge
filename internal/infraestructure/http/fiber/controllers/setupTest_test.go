package controllers_test

import (
	"backend-challenge-api/internal/domain/entities"
	"backend-challenge-api/internal/infraestructure/http/auth"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

var (
	Err                              error
	ValidToken                       string
	ExpressionOutput                 entities.Expression
	ExpressionsOutput                []entities.Expression
	ExpressionJsonSerialized         []byte
	CreatedExpressionJsonSerialized  []byte
	CreatedExpressionsJsonSerialized []byte
)

func TestMain(m *testing.M) {
	ValidToken, _ = auth.GenerateToken()

	ExpressionJsonSerialized, Err = os.ReadFile("../../../../../test/resources/expression.json")
	if Err != nil {
		panic(fmt.Sprintf("Error on parsing setup test variables, %s", Err))
	}

	CreatedExpressionJsonSerialized, Err = os.ReadFile("../../../../../test/resources/created_expression.json")
	if Err != nil {
		panic(fmt.Sprintf("Error on parsing setup test variables, %s", Err))
	}

	if err := json.Unmarshal(CreatedExpressionJsonSerialized, &ExpressionOutput); err != nil {
		panic(fmt.Sprintf("Error on read created expression json from resources, %s", err))
	}

	CreatedExpressionsJsonSerialized, Err = os.ReadFile("../../../../../test/resources/expressions.json")
	if Err != nil {
		panic(fmt.Sprintf("Error on read created expression json from resources, %s", Err))
	}

	if Err = json.Unmarshal(CreatedExpressionsJsonSerialized, &ExpressionsOutput); Err != nil {
		panic(fmt.Sprintf("Error on read created expression json from resources, %s", Err))
	}

	os.Exit(m.Run())
}
