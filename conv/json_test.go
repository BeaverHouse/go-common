package conv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalJSONOrDefault(t *testing.T) {
	b, err := MarshalJSONOrDefault(nil, "[]")
	assert.NoError(t, err)
	assert.Equal(t, "[]", string(b))

	b, err = MarshalJSONOrDefault([]int{1, 2}, "[]")
	assert.NoError(t, err)
	assert.Equal(t, "[1,2]", string(b))
}

func TestMarshalJSONOrNil(t *testing.T) {
	b, err := MarshalJSONOrNil(nil)
	assert.NoError(t, err)
	assert.Nil(t, b)

	b, err = MarshalJSONOrNil(map[string]int{"a": 1})
	assert.NoError(t, err)
	assert.Equal(t, `{"a":1}`, string(b))
}
