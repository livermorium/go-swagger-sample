package services

import (
	"fmt"

	swag "github.com/go-openapi/swag"
	middleware "github.com/go-openapi/runtime/middleware"

	"o/samples/swagger-sample/gen/restapi/operations"
	"o/samples/swagger-sample/gen/restapi/operations/greetings"
	"o/samples/swagger-sample/gen/models"
)

func Handle(api *operations.GreetingsAPI)  {
	
	api.GreetingsGetHelloHandler = greetings.GetHelloHandlerFunc(func(params greetings.GetHelloParams) middleware.Responder {
		name := swag.StringValue(params.Name)
		if name == "" {
			name = "World"
		}

		greeting := models.Greetings(fmt.Sprintf("Hello, %s!", name))
		return greetings.NewGetHelloOK().WithPayload(greeting)})
}