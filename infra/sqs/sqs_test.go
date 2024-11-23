package sqs

import (
	"fiap-fast-food-ms-producao/infra/ctx"
	"testing"
)

func TestNewSQSClient(t *testing.T) {
	ctx := ctx.NewContextManager()
	sqsUrl := ctx.Get("aws_production_update_sqs_url")
	NewSQSClient("us-east-1", sqsUrl.(string))
}
