package version

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestGetBuildInfo(t *testing.T) {
    assert := assert.New(t)

    expected := ""
    actual := ""

    assert.Equal(actual, expected)
}

// https://github.com/google/ko/blob/main/pkg/commands/version.go
