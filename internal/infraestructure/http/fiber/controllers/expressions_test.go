package controllers_test

import (
	"backend-challenge-api/internal/application/services/mocks"
	"backend-challenge-api/internal/domain/entities"
	"backend-challenge-api/internal/infraestructure/http/fiber/controllers"
	"backend-challenge-api/internal/infraestructure/http/fiber/routes"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterExpression(t *testing.T) {

	tests := []struct {
		name                             string
		input                            *bytes.Buffer
		expectedStatusCode               int
		expectedRegisterExpressionResult entities.Expression
		expectedRegisterExpressionError  error
		authToken                        string
	}{
		{
			name:                             "Success on register an expression",
			input:                            bytes.NewBuffer(ExpressionJsonSerialized),
			expectedStatusCode:               200,
			expectedRegisterExpressionResult: ExpressionOutput,
			expectedRegisterExpressionError:  nil,
			authToken:                        ValidToken,
		},
		{
			name:                             "Error on register an expression",
			input:                            bytes.NewBuffer(ExpressionJsonSerialized),
			expectedStatusCode:               500,
			expectedRegisterExpressionResult: entities.Expression{},
			expectedRegisterExpressionError:  assert.AnError,
			authToken:                        ValidToken,
		},
		{
			name:                             "Error on register an expression, broken input",
			input:                            bytes.NewBuffer([]byte("{ex:'1'}")),
			expectedStatusCode:               500,
			expectedRegisterExpressionResult: entities.Expression{},
			expectedRegisterExpressionError:  assert.AnError,
			authToken:                        ValidToken,
		},
		{
			name:                             "Invalid token",
			input:                            bytes.NewBuffer(ExpressionJsonSerialized),
			expectedStatusCode:               401,
			expectedRegisterExpressionResult: entities.Expression{},
			expectedRegisterExpressionError:  nil,
			authToken:                        ValidToken + "invalidatingTokenString",
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

			req := httptest.NewRequest(fiber.MethodPost, "/expressions", test.input)
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Authorization", "Bearer "+test.authToken)

			response, err := app.Test(req)
			if err != nil {
				t.Errorf("Error on test app with created req")
			}

			var actualResponse entities.Expression
			responseReadCloser, _ := ioutil.ReadAll(response.Body)
			if err := json.Unmarshal(responseReadCloser, &actualResponse); err != nil {
				t.Errorf("Error on parse bodyResp")
			}

			assert.Equal(t, test.expectedStatusCode, response.StatusCode)
			assert.Equal(t, test.expectedRegisterExpressionResult, actualResponse)
		})
	}
}

func TestGetExpressions(t *testing.T) {
	tests := []struct {
		name                         string
		expectedStatusCode           int
		expectedGetExpressionsResult []entities.Expression
		expectedGetExpressionsError  error
		authToken                    string
	}{
		{
			name:                         "Success on get expressions",
			expectedStatusCode:           200,
			expectedGetExpressionsResult: ExpressionsOutput,
			expectedGetExpressionsError:  nil,
			authToken:                    ValidToken,
		},
		{
			name:                         "Error on get expressions",
			expectedStatusCode:           500,
			expectedGetExpressionsResult: []entities.Expression(nil),
			expectedGetExpressionsError:  assert.AnError,
			authToken:                    ValidToken,
		},
		{
			name:                         "Error on get expressions, invalid token",
			expectedStatusCode:           401,
			expectedGetExpressionsResult: []entities.Expression(nil),
			expectedGetExpressionsError:  nil,
			authToken:                    ValidToken + "invalidatingTokenString",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			servicesMock := mocks.NewExpressionServiceMock()
			servicesMock.On("GetExpressions").Return(test.expectedGetExpressionsResult, test.expectedGetExpressionsError)
			controllers := controllers.NewAPIControllers(servicesMock)

			app := fiber.New()
			routes.SetupRoutes(app, controllers)

			req := httptest.NewRequest(fiber.MethodGet, "/expressions", nil)
			req.Header.Add("Authorization", "Bearer "+test.authToken)
			req.Header.Add("Content-Type", "application/json")

			response, err := app.Test(req)
			if err != nil {
				t.Errorf("Error on test app with created req")
			}

			var actualResponse []entities.Expression
			responseReadCloser, _ := ioutil.ReadAll(response.Body)
			if response.StatusCode == 200 {
				if err := json.Unmarshal(responseReadCloser, &actualResponse); err != nil {
					t.Errorf(fmt.Sprintf("Error on parse: %s", err))
				}
			}

			assert.Equal(t, test.expectedStatusCode, response.StatusCode)
			assert.Equal(t, test.expectedGetExpressionsResult, actualResponse)
		})
	}
}
