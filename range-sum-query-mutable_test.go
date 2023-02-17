package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConstructorI(t *testing.T) {

	// ["NumArray","sumRange","update","sumRange"]
	// [[[1,3,5]],[0,2],[1,2],[0,2]]

	segTree := ConstructorI([]int{1, 3, 5})

	require.Equal(t, 9, segTree.SumRange(0, 2))

	segTree.Update(1, 2)

	require.Equal(t, 8, segTree.SumRange(0, 2))

}
