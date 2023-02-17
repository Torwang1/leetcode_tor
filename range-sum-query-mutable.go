package leetcode

// 线段树:
// 1) Build: O(n)
// 2) Update: O(logn) 对范围内的节点, 做相同的操作;
// 3) Query: O(logn) 查询范围的特征;
//
// 典型特征:
// 1) 父节点的特征, 可以通过“相邻的”“两个字节点”推理出来;
//

type NumArray struct {
	size    int
	segTree []int // 空间复杂度: O(4n)
}

// Build 静态构建过程;
func ConstructorI(nums []int) NumArray {
	result := NumArray{
		size:    len(nums),
		segTree: make([]int, 4*len(nums)),
	}
	result.BuildImpl(nums, 0, 0, result.size-1)

	return result
}

// 递归构造过程:
// 1) l, r: 分别代表 range 的左右边界, [l, r];
// 2) pos : 线段树节点下标;
func (this *NumArray) BuildImpl(nums []int, pos int, l, r int) {
	if l == r {
		this.segTree[pos] = nums[l]
		return
	}

	left, right := 2*pos+1, 2*pos+2

	mid := l + (r-l)/2
	this.BuildImpl(nums, left, l, mid)
	this.BuildImpl(nums, right, mid+1, r)

	// 线段和
	this.segTree[pos] = this.segTree[left] + this.segTree[right]
}

func (this *NumArray) Update(index int, val int) {
	this.UpdateImpl(index, 0, 0, this.size-1, val)
}

// 对区间内的所有元素, 做相同的更新操作. 其中:
// [s,e] [l,r] 表示原始数组上的区间;
// segTree[pos] 表示 [l,r] 区间和;
func (this *NumArray) UpdateImpl(index int, pos int, l, r int, val int) {

	if l == r {
		this.segTree[pos] = val
		return
	}

	left, right := 2*pos+1, 2*pos+2

	mid := l + (r-l)/2
	if mid >= index {
		this.UpdateImpl(index, left, l, mid, val)
	} else {
		this.UpdateImpl(index, right, mid+1, r, val)
	}
	this.segTree[pos] = this.segTree[left] + this.segTree[right]
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.Query(left, right, 0, 0, this.size-1)
}

// 查询区间和. 其中:
// [s,e] [l,r] 表示原始数组上的区间;
// segTree[pos] 表示 [l,r] 区间和;
func (this *NumArray) Query(s, e int, pos int, l, r int) int {
	// [l,r] 是 [s,e] 的子区间. segTree[pos] 表示 [l,r] 的区间和.
	if l >= s && r <= e {
		return this.segTree[pos]
	}

	var ans int

	left, right := 2*pos+1, 2*pos+2

	mid := l + (r-l)/2
	if mid >= s {
		ans += this.Query(s, e, left, l, mid)
	}
	if mid+1 <= e {
		ans += this.Query(s, e, right, mid+1, r)
	}

	return ans
}
