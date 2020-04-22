package units

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleParser(t *testing.T) {
	assert := assert.New(t)

	parser := &parseSimpleLog{
		Key:     "foo",
		Content: []string{"foo", "bar"},
	}
	assert.NoError(parser.Validate())
	parser.Run(context.TODO())
	// assert.True(parser.Completed())

	// TODO: setupChangePointsTest local queue or db service as a prereq

	// assert.NoError(parser.Error())

	// assert.Equal(parser.numLines, 2)
	// assert.Equal(parser.freq["o"], 2)
}
