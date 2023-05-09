package controllers_test

import (
	"backend-challenge-api/internal/application/services/mocks"
	"backend-challenge-api/internal/domain/entities"
	"backend-challenge-api/internal/infraestructure/http/auth"
	"backend-challenge-api/internal/infraestructure/http/fiber/controllers"
	"backend-challenge-api/internal/infraestructure/http/fiber/routes"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterExpression(t *testing.T) {

	expressionJsonSerialized, err := os.ReadFile("../../../../../test/resources/expression.json")
	if err != nil {
		t.Errorf("Error on read expression json from resources, %s", err)
	}

	createdExpressionJsonSerialized, err := os.ReadFile("../../../../../test/resources/created_expression.json")
	if err != nil {
		t.Errorf("Error on read created expression json from resources, %s", err)
	}

	expressionOutput := entities.Expression{}
	if err := json.Unmarshal(createdExpressionJsonSerialized, &expressionOutput); err != nil {
		t.Errorf("Error on parse serialized expressions output")
	}

	tests := []struct {
		name                             string
		input                            *bytes.Buffer
		expectedStatusCode               int
		expectedRegisterExpressionResult entities.Expression
		expectedRegisterExpressionError  error
	}{
		{
			name:                             "Success on register an expression",
			input:                            bytes.NewBuffer(expressionJsonSerialized),
			expectedStatusCode:               200,
			expectedRegisterExpressionResult: expressionOutput,
			expectedRegisterExpressionError:  nil,
		},
		{
			name:                             "Error on register an expression",
			input:                            bytes.NewBuffer(expressionJsonSerialized),
			expectedStatusCode:               500,
			expectedRegisterExpressionResult: entities.Expression{},
			expectedRegisterExpressionError:  assert.AnError,
		},
		{
			name:                             "Error on register an expression, broken input",
			input:                            bytes.NewBuffer([]byte("{ex:'1'}")),
			expectedStatusCode:               500,
			expectedRegisterExpressionResult: entities.Expression{},
			expectedRegisterExpressionError:  assert.AnError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			servicesMock := mocks.NewExpressionServiceMock()
			servicesMock.On("RegisterExpression", mock.AnythingOfType("*entities.Expression")).
				Return(test.expectedRegisterExpressionResult, test.expectedRegisterExpressionError)

			controllers := controllers.NewAPIControllers(servicesMock)

			app := fiber.New()
			routes.SetupRoutes(app, controllers)

			validToken, _ := auth.GenerateToken()

			req := httptest.NewRequest(fiber.MethodPost, "/expressions", test.input)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Add("Authorization", "Bearer "+validToken)

			response, err := app.Test(req)
			if err != nil {
				t.Errorf("Error on test app with created req")
			}

			responseReadCloser, _ := ioutil.ReadAll(response.Body)
			var actualResponse entities.Expression
			if err := json.Unmarshal(responseReadCloser, &actualResponse); err != nil {
				t.Errorf("Error on parse bodyResp")
			}

			assert.Equal(t, test.expectedStatusCode, response.StatusCode)
			assert.Equal(t, test.expectedRegisterExpressionResult, actualResponse)
		})
	}
}
