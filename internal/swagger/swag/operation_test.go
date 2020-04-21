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

func TestOperation_ParseComment(t *testing.T) {
	t.Run("test empty",
		func(t *testing.T) {
			operation := NewOperation()
			err := operation.ParseComment("//", nil)
			assert.NoError(t, err)
		})
}
