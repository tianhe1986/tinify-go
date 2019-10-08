package tinifytest

import (
	"testing"

	//"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/tianhe1986/tinify-go/tinify"
)

func TestKey(t *testing.T) {
	var key string = "ok your key"
	tinify.SetKey(key)
	assert := assert.New(t)
	// assert equality
	assert.Equal(key, tinify.GetKey(), "they should be equal")
}
