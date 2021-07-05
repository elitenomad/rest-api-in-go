// +build e2e

package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetComments(t *testing.T) {
	fmt.Println("Test GetAllComments API endpoint...")

	client := resty.New()
	response, err := client.R().Get("http://localhost:8080/api/comments")

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, response.StatusCode())
	// assert.Equal(t, [], response.Body())
}

func TestPostComments(t *testing.T) {
	fmt.Println("Test PostComment API endpoint...")

	client := resty.New()
	response, err := client.
		R().
		SetBody(`{"slug": "pranava", "Author": "pranava", "Body": "Tested it"	}`).
		Get("http://localhost:8080/api/comments")

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, response.StatusCode())
	assert.NotEmpty(t, response.Body())
	assert.NoError(t, err)
}
