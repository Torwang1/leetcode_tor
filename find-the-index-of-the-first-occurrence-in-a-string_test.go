package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_strStr(t *testing.T) {
	index := strStr("ababcaababcaabc", "ababcaabc")

	require.Equal(t, 6, index)
}
