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
			expectedStatusCode:               201,
			expectedRegisterExpressionResult: Expression,
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
			expectedStatusCode:               400,
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
				t.Errorf("Error on test app with created req: %s", err)
			}

			var actualResponse entities.Expression
			responseReadCloser, _ := ioutil.ReadAll(response.Body)
			if err := json.Unmarshal(responseReadCloser, &actualResponse); err != nil {
				t.Errorf("Error on parse bodyResp: %s", err)
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
			expectedGetExpressionsResult: Expressions,
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
				t.Errorf("Error on test app with created req: %s", err)
			}

			var expressions []entities.Expression
			responseReadCloser, _ := ioutil.ReadAll(response.Body)
			if response.StatusCode == 200 {
				if err := json.Unmarshal(responseReadCloser, &expressions); err != nil {
					t.Errorf(fmt.Sprintf("Error on parse: %s", err))
				}
			}

			assert.Equal(t, test.expectedStatusCode, response.StatusCode)
			assert.Equal(t, test.expectedGetExpressionsResult, expressions)
		})
	}
}

func TestUpdateExpressions(t *testing.T) {
	tests := []struct {
		name                           string
		input                          *bytes.Buffer
		expressionID                   string
		expectedUpdateExpressionResult entities.Expression
		expectedUpdateExpressionError  error
		expectedStatusCode             int
		authToken                      string
	}{
		{
			name:                           "Success on update an expression",
			input:                          bytes.NewBuffer(ExpressionJsonSerialized),
			expressionID:                   "1",
			expectedUpdateExpressionResult: Expression,
			expectedUpdateExpressionError:  nil,
			expectedStatusCode:             200,
			authToken:                      ValidToken,
		},
		{
			name:                           "Error on update an expression",
			input:                          bytes.NewBuffer(ExpressionJsonSerialized),
			expressionID:                   "1",
			expectedUpdateExpressionResult: entities.Expression{},
			expectedUpdateExpressionError:  assert.AnError,
			expectedStatusCode:             500,
			authToken:                      ValidToken,
		},
		{
			name:                           "Error on update an expression, invalid expressionID",
			input:                          bytes.NewBuffer(ExpressionJsonSerialized),
			expressionID:                   "xyz",
			expectedUpdateExpressionResult: entities.Expression{},
			expectedUpdateExpressionError:  nil,
			expectedStatusCode:             400,
			authToken:                      ValidToken,
		},
		{
			name:                           "Error on update an expression, broken json",
			input:                          bytes.NewBuffer([]byte("{ex:'1'}")),
			expressionID:                   "1",
			expectedUpdateExpressionResult: entities.Expression{},
			expectedUpdateExpressionError:  assert.AnError,
			expectedStatusCode:             400,
			authToken:                      ValidToken,
		},
		{
			name:                           "Error on update an expression, invalid auth token",
			input:                          bytes.NewBuffer(ExpressionJsonSerialized),
			expressionID:                   "1",
			expectedUpdateExpressionResult: entities.Expression{},
			expectedUpdateExpressionError:  nil,
			expectedStatusCode:             401,
			authToken:                      ValidToken + "invalidating auth token",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			servicesMock := mocks.NewExpressionServiceMock()
			servicesMock.On("UpdateExpression", mock.AnythingOfType("uint64"), mock.AnythingOfType("*string")).
				Return(test.expectedUpdateExpressionResult, test.expectedUpdateExpressionError)

			controllers := controllers.NewAPIControllers(servicesMock)

			app := fiber.New()
			routes.SetupRoutes(app, controllers)

			req := httptest.NewRequest("PATCH", fmt.Sprintf("/expressions/%s", test.expressionID), test.input)
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Authorization", "Bearer "+test.authToken)

			response, err := app.Test(req)
			if err != nil {
				t.Errorf("Error on test app with an created req: %s", err)
			}

			var updatedExpression entities.Expression
			responseReadCloser, _ := ioutil.ReadAll(response.Body)
			if err := json.Unmarshal(responseReadCloser, &updatedExpression); err != nil {
				t.Errorf("Error on parse bodyResp: %s", err)
			}

			assert.Equal(t, test.expectedStatusCode, response.StatusCode)
			assert.Equal(t, test.expectedUpdateExpressionResult, updatedExpression)

		})
	}
}

func TestDeleteExpression(t *testing.T) {
	tests := []struct {
		name                           string
		expressionID                   string
		expectedDeleteExpressionResult error
		expectedStatusCode             int
		authToken                      string
	}{
		{
			name:                           "Success on delete an expression",
			expressionID:                   "1",
			expectedDeleteExpressionResult: nil,
			expectedStatusCode:             204,
			authToken:                      ValidToken,
		},
		{
			name:                           "Error on delete an expression",
			expressionID:                   "1",
			expectedDeleteExpressionResult: assert.AnError,
			expectedStatusCode:             500,
			authToken:                      ValidToken,
		},
		{
			name:                           "Error on delete an expression, invalid expressionID",
			expressionID:                   "abcde",
			expectedDeleteExpressionResult: nil,
			expectedStatusCode:             400,
			authToken:                      ValidToken,
		},
		{
			name:                           "Error on delete an expression, invalid auth token",
			expressionID:                   "1",
			expectedDeleteExpressionResult: assert.AnError,
			expectedStatusCode:             401,
			authToken:                      ValidToken + "invalidating auth token",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			servicesMock := mocks.NewExpressionServiceMock()
			servicesMock.On("DeleteExpressionById", mock.AnythingOfType("uint64")).Return(test.expectedDeleteExpressionResult)

			controllers := controllers.NewAPIControllers(servicesMock)

			app := fiber.New()
			routes.SetupRoutes(app, controllers)

			req := httptest.NewRequest("DELETE", fmt.Sprintf("/expressions/%s", test.expressionID), nil)
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Authorization", "Bearer "+test.authToken)

			response, err := app.Test(req)
			if err != nil {
				t.Errorf("Error on test app with an created req: %s", err)
			}

			assert.Equal(t, test.expectedStatusCode, response.StatusCode)
		})
	}
}

func TestCalculateExpression(t *testing.T) {
	tests := []struct {
		name                             string
		expressionID                     string
		expectedStatusCode               int
		expectedCalculateExpressionError error
		authToken                        string
	}{
		{
			name:                             "Success on calculate expressions",
			expressionID:                     "1",
			expectedStatusCode:               200,
			expectedCalculateExpressionError: nil,
			authToken:                        ValidToken,
		},
		{
			name:                             "Error on calculate expressions",
			expressionID:                     "1",
			expectedStatusCode:               500,
			expectedCalculateExpressionError: assert.AnError,
			authToken:                        ValidToken,
		},
		{
			name:                             "Error on calculate expressions, invalid token",
			expressionID:                     "1",
			expectedStatusCode:               401,
			expectedCalculateExpressionError: nil,
			authToken:                        ValidToken + "invalidatingTokenString",
		},
		{
			name:                             "Error on calculate expressions, invalid expressionID",
			expressionID:                     "abcde",
			expectedStatusCode:               500,
			expectedCalculateExpressionError: assert.AnError,
			authToken:                        ValidToken,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			servicesMock := mocks.NewExpressionServiceMock()
			servicesMock.On("CalculateExpression", mock.AnythingOfType("uint64"), mock.AnythingOfType("*fiber.Ctx")).
				Return(mock.AnythingOfType("interface{}"), test.expectedCalculateExpressionError)
			controllers := controllers.NewAPIControllers(servicesMock)

			app := fiber.New()
			routes.SetupRoutes(app, controllers)

			req := httptest.NewRequest(fiber.MethodGet, fmt.Sprintf("/expressions/%s", test.expressionID), nil)
			req.Header.Add("Authorization", "Bearer "+test.authToken)
			req.Header.Add("Content-Type", "application/json")

			response, err := app.Test(req)
			if err != nil {
				t.Errorf("Error on test app with created req: %s", err)
			}

			assert.Equal(t, test.expectedStatusCode, response.StatusCode)
		})
	}
}
