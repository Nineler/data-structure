package MaxHeap

import "data_structure/Array"

type MaxHeap struct {
	data Array.IArray
}

func NewMaxHeap(cap int) *MaxHeap {
	return &MaxHeap{data: Array.NewArray(cap)}
}

func (m *MaxHeap) Size() int {
	return m.data.GetSize()
}

func (m *MaxHeap) IsEmpty() bool {
	return m.data.GetSize() == 0
}

//返回完全二叉树的数组表示中，一个索引所表示的元素的父亲节点的索引
func (m *MaxHeap) Parent(index int) int {
	if index == 0 {
		panic("index-0 does't have parent.")
	}
	return (index - 1) / 2
}

//返回完全二叉树的数组表示中，一个索引表示的元素的左孩子的索引
func (m *MaxHeap) LeftChild(index int) int {
	return index*2 + 1
}

//返回完全二叉树的数组表示中，一个索引表示的元素的左孩子的索引
func (m *MaxHeap) RightChild(index int) int {
	return index*2 + 2
}

//向堆中添加元素
func (m *MaxHeap) Add(e int) {
	m.data.AddLast(e)
	m.siftUp(m.data.GetSize() - 1)
}

func (m *MaxHeap) siftUp(k int) {
	for k > 0 && m.data.GetIndex(m.Parent(k)).(int) < m.data.GetIndex(k).(int) {
		m.data.Swap(m.Parent(k), k)
		k = m.Parent(k)
	}
}
func (m *MaxHeap) FindMax() int {
	if m.data.GetSize() == 0 {
		panic("Can not findMax when heap is empty")
	} else {
		return m.data.GetIndex(0).(int)
	}
}

//取出堆中最大的元素
func (m *MaxHeap) ExtractMax() int {
	ret := m.FindMax()
	m.data.Swap(0, m.data.GetSize()-1)
	m.data.DelectLast()
	m.siftDown(0)
	return ret
}

func (m *MaxHeap) siftDown(k int) {
	for m.LeftChild(k) < m.data.GetSize() {
		j := m.LeftChild(k)
		if j+1 < m.data.GetSize() && m.data.GetIndex(j+1).(int) > m.data.GetIndex(j).(int) {
			j = m.RightChild(k)
			//data[j]是leftChild和rightChild中的最大值
		}
		if m.data.GetIndex(k).(int) >= m.data.GetIndex(j).(int) {
			break
		}
		m.data.Swap(k, j)
		k = j
	}
}

//取出堆中的最大元素，并且替换成元素e
func (m *MaxHeap) Replace(e int) int {
	ret := m.FindMax()
	m.data.SetIndex(0, e)
	m.siftDown(0)
	return ret
}
