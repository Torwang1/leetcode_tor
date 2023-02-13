package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// ["LRUCache","put","put","get","put","get","put","get","get","get"]
// [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
func TestLRUCache(t *testing.T) {
	lru := ConstructorLRU(2)

	lru.Put(1, 1)
	require.Equal(t, 1, lru.list.Len())
	require.Equal(t, 1, len(lru.hash))

	lru.Put(2, 2)
	require.Equal(t, 2, lru.list.Len())
	require.Equal(t, 2, len(lru.hash))

	lru.Get(1)
	require.Equal(t, 2, lru.list.Len())
	require.Equal(t, 2, len(lru.hash))

	lru.Put(3, 3)
	require.Equal(t, 2, lru.list.Len())
	require.Equal(t, 2, len(lru.hash))

	require.Equal(t, 2, lru.Get(2))

}
