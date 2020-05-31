package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAllEndpointAttributesHaveName(t *testing.T) {
	for i := endpointAttribute(0); i < attrInvalid; i++ {
		require.NotEmpty(t, i.String())
	}
}
