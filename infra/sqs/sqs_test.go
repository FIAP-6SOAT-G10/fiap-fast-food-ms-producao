package sqs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSQSClient(t *testing.T) {
	client, err := NewSQSClient("us-east-1")

	assert.NotNil(t, client)
	assert.Nil(t, err)
}
