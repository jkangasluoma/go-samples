package main

import (
	"context"
	"fmt"

	lambdaGo "github.com/aws/aws-lambda-go/lambda"
)

// Response lambda response type.
type Response struct {
	Message string `json:"message"`
}

// Handler lambda main handler.
func Handler(ctx context.Context) (Response, error) {
	err := do()
	if err != nil {
		fmt.Println("Error", err)
	}
	return Response{
		Message: "Go Serverless v1.0! Your function executed successfully!",
	}, nil
}

func do() error {
	fmt.Println("Lambda was here.")
	return nil
}

// https://www.enotes.com/shakespeare-quotes/once-more-unto-breach
func main() {
	lambdaGo.Start(Handler)
}
