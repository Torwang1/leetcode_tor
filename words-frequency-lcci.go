package leetcode

// trie 字典树, 特征:
// - 组成每个 word 的字符有限, 26 个英文小写字母;

type WordsFrequency struct {
	root trieNode
}

// 构造过程: 把每个 string 按字节拆解, 构造 trieNode 节点;
func Constructor(book []string) WordsFrequency {
	result := WordsFrequency{
		root: trieNode{
			Children: map[byte]*trieNode{},
		},
	}

	for _, str := range book {
		node := &result.root
		for i := 0; i < len(str); i++ {
			char := str[i]

			charNode := node.Children[char]
			if charNode == nil {
				charNode = &trieNode{
					Children: map[byte]*trieNode{},
				}
			}
			node.Children[char] = charNode

			node = charNode
		}
		node.Count++
	}

	return result
}

func (this *WordsFrequency) Get(word string) int {
	node := &this.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		node = node.Children[char]
		if node == nil {
			break
		}
	}

	if node == nil {
		return 0
	}
	return node.Count
}

type trieNode struct {
	Count    int
	Children map[byte]*trieNode
}
