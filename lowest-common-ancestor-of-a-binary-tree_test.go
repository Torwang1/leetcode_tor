package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lowestCommonAncestor(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}

	require.Equal(t, 3, lowestCommonAncestor(root, root.Left, root.Right))
}
