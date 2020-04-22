package swag

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	goparser "go/parser"
	"go/token"
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

func TestOperation_ParseProduceComment(t *testing.T) {
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
		b, err := json.MarshalIndent(operation, "", "    ")
		assert.NoError(t, err)
		assert.JSONEq(t, expected, string(b))
	})
	t.Run("error", func(t *testing.T) {
		comment := `@Produce foo`
		operation := new(Operation)
		err := operation.ParseProduceComment(comment)
		assert.Error(t, err)
	})
}

func TestOperation_ParseRouterComment(t *testing.T) {
	t.Run("comment", func(t *testing.T) {
		comment := `/customer/get-wishlist/{wishlist_id} [get]`
		operation := NewOperation()
		err := operation.ParseRouterComment(comment)
		assert.NoError(t, err)
		assert.Equal(t, "/customer/get-wishlist/{wishlist_id}", operation.Path)
		assert.Equal(t, "GET", operation.HTTPMethod)
	})

	t.Run("slash", func(t *testing.T) {
		comment := `/ [get]`
		operation := NewOperation()
		err := operation.ParseRouterComment(comment)
		assert.NoError(t, err)
		assert.Equal(t, "/", operation.Path)
		assert.Equal(t, "GET", operation.HTTPMethod)
	})

	t.Run("plus", func(t *testing.T) {
		comment := `/customer/get-wishlist/{proxy+} [post]`
		operation := NewOperation()
		err := operation.ParseRouterComment(comment)
		assert.NoError(t, err)
		assert.Equal(t, "/customer/get-wishlist/{proxy+}", operation.Path)
		assert.Equal(t, "POST", operation.HTTPMethod)
	})

	t.Run("colon", func(t *testing.T) {
		comment := `/customer/get-wishlist/{wishlist_id}:move [post]`
		operation := NewOperation()
		err := operation.ParseRouterComment(comment)
		assert.NoError(t, err)
		assert.Equal(t, "/customer/get-wishlist/{wishlist_id}:move", operation.Path)
		assert.Equal(t, "POST", operation.HTTPMethod)
	})

	t.Run("start error", func(t *testing.T) {
		comment := `:customer/get-wishlist/{wishlist_id}:move [post]`
		operation := NewOperation()
		err := operation.ParseRouterComment(comment)
		assert.Error(t, err)
	})

	t.Run("method error", func(t *testing.T) {
		comment := `/api/{id}|,*[get`
		operation := NewOperation()
		err := operation.ParseRouterComment(comment)
		assert.Error(t, err)
	})

	t.Run("method missing", func(t *testing.T) {
		comment := `/customer/get-wishlist/{wishlist_id}`
		operation := NewOperation()
		err := operation.ParseRouterComment(comment)
		assert.Error(t, err)
	})
}

func TestOperation_ParseParamComment(t *testing.T) {

}

func TestOperation_ParseResponseComment(t *testing.T) {

}

func TestOperation_ParseResponseHeaderComment(t *testing.T) {

}

func TestOperation_ParseEmptyResponseComment(t *testing.T) {
}

func TestOperation_ParseEmptyResponseOnly(t *testing.T) {
	comment := `200`
	operation := NewOperation()
	err := operation.ParseEmptyResponseOnly(comment)
	assert.NoError(t, err, "ParseComment should not fail")

	b, _ := json.MarshalIndent(operation, "", "    ")

	expected := `{
    "responses": {
        "200": {}
    }
}`
	assert.Equal(t, expected, string(b))
}

func TestOperation_ParseSecurityComment(t *testing.T) {

}

func TestOperation_ParseMetadata(t *testing.T) {
	t.Run("need a value", func(t *testing.T) {
		operation := NewOperation()
		operation.parser = New()
		err := operation.ParseMetadata("@x-amazon-apigateway-integration", "@x-amazon-apigateway-integration", "")
		assert.EqualError(t, err, "annotation @x-amazon-apigateway-integration need a value")
	})

	// Fail if args of attributes are broken.
	t.Run("attributes are broken", func(t *testing.T) {
		operation := NewOperation()
		operation.parser = New()

		err := operation.ParseMetadata("@x-amazon-apigateway-integration", "@x-amazon-apigateway-integration", "[\"broken\"}]")
		assert.EqualError(t, err, "annotation @x-amazon-apigateway-integration need a valid json value")
	})

	t.Run("pass", func(t *testing.T) {
		operation := NewOperation()
		operation.parser = New()

		err := operation.ParseMetadata("@x-amazon-apigateway-integration","@x-amazon-apigateway-integration", "{\"uri\": \"${some_arn}\", \"passthroughBehavior\": \"when_no_match\", \"httpMethod\": \"POST\", \"type\": \"aws_proxy\"}")
		assert.NoError(t, err)

		expected := `{
    "x-amazon-apigateway-integration": {
        "httpMethod": "POST",
        "passthroughBehavior": "when_no_match",
        "type": "aws_proxy",
        "uri": "${some_arn}"
    }
}`
		b, _ := json.MarshalIndent(operation, "", "    ")
		assert.Equal(t, expected, string(b))
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

func TestOperation_registerSchemaType(t*testing.T){

	operation := NewOperation()

	fset := token.NewFileSet()
	astFile, err := goparser.ParseFile(fset, "main.go", `package main
	import "timer"
`, goparser.ParseComments)

	assert.NoError(t, err)

	operation.parser = New()
	_, _, err = operation.registerSchemaType("timer.Location", astFile)
	assert.Error(t, err)
}
