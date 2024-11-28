package sqs

import (
	"fiap-fast-food-ms-producao/adapter/context_manager"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

// NewSQSClient initializes the SQS client.
func NewSQSClient(ctx context_manager.ContextManager, region string) (sqsiface.SQSAPI, error) {
	aws_access_key_id := ctx.Get("AWS_ACCESS_KEY_ID")
	aws_secret_access_key := ctx.Get("AWS_SECRET_ACCESS_KEY")
	aws_session_token := ctx.Get("AWS_SESSION_TOKEN")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(aws_access_key_id.(string), aws_secret_access_key.(string), aws_session_token.(string)),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	return sqs.New(sess), nil
}
