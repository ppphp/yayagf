package swag

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOperation(t *testing.T) {
	operation := NewOperation()
	assert.NotNil(t, operation)
}

func TestOperation_ParseDescriptionComment(t *testing.T) {
	comment := `line one`
	operation := NewOperation()
	operation.parser = New()
	operation.ParseDescriptionComment(comment)
	comment = `line two x`
	operation.ParseDescriptionComment(comment)

	b, err := json.MarshalIndent(operation, "", "    ")
	assert.NoError(t, err)

	expected := `"description": "line one\nline two x"`
	assert.Contains(t, string(b), expected)
}

func TestOperation_ParseTagsComment(t *testing.T) {
	expected := `{
    "tags": [
        "pet",
        "store",
        "user"
    ]
}`
	comment := `pet, store,user`
	operation := NewOperation()
	operation.ParseTagsComment(comment)
	b, err := json.MarshalIndent(operation, "", "    ")
	assert.NoError(t, err)
	assert.Equal(t, expected, string(b))
}

func TestOperation_ParseAcceptComment(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		expected := `{
    "consumes": [
        "application/json",
        "text/xml",
        "text/plain",
        "text/html",
        "multipart/form-data",
        "application/x-www-form-urlencoded",
        "application/vnd.api+json",
        "application/x-json-stream",
		"application/octet-stream",
		"image/png",
		"image/jpeg",
		"image/gif",
		"application/xhtml+xml",
		"application/health+json"
    ]
}`
		comment := `json,xml,plain,html,mpfd,x-www-form-urlencoded,json-api,json-stream,octet-stream,png,jpeg,gif,application/xhtml+xml,application/health+json`
		operation := NewOperation()
		err := operation.ParseAcceptComment(comment)
		assert.NoError(t, err)
		b, _ := json.MarshalIndent(operation, "", "    ")
		assert.JSONEq(t, expected, string(b))
	})

	t.Run("error", func(t *testing.T) {
		comment := `@Accept unknown`
		operation := NewOperation()
		err := operation.ParseAcceptComment(comment)
		assert.Error(t, err)
	})
}

func TestParseProduceComment(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		expected := `{
    "produces": [
        "application/json",
        "text/xml",
        "text/plain",
        "text/html",
        "multipart/form-data",
        "application/x-www-form-urlencoded",
        "application/vnd.api+json",
        "application/x-json-stream",
		"application/octet-stream",
		"image/png",
		"image/jpeg",
		"image/gif",
		"application/health+json"
    ]
}`
		comment := `json,xml,plain,html,mpfd,x-www-form-urlencoded,json-api,json-stream,octet-stream,png,jpeg,gif,application/health+json`
		operation := new(Operation)
		err := operation.ParseProduceComment(comment)
		assert.NoError(t, err, "ParseComment should not fail")
		b, _ := json.MarshalIndent(operation, "", "    ")
		assert.JSONEq(t, expected, string(b))
	})
	t.Run("error", func(t *testing.T) {
		comment := `@Produce foo`
		operation := new(Operation)
		err := operation.ParseProduceComment(comment)
		assert.Error(t, err)
	})
}

func TestOperation_ParseComment(t *testing.T) {
	t.Run("test empty",
		func(t *testing.T) {
			operation := NewOperation()
			err := operation.ParseComment("//", nil)
			assert.NoError(t, err)
		})
}
