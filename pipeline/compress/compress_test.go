package compress

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtWindows(t *testing.T) {
	assert.Equal(t, ext("windows"), ".exe")
}

func TestExtOthers(t *testing.T) {
	assert.Empty(t, ext("linux"))
}
