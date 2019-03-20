package Leetcode

func dailyTemperatures(T []int) []int {
	ms := &mystack{}
	res := make([]int,len(T))

	for idx, rec := range T {
		for !ms.Empty() && rec > T[ms.Top()] {
			res[ms.Top()] = idx - ms.Top()
			ms.Pop()
		}
		ms.Push(idx)
	}

	for !ms.Empty() {
		res[ms.Pop()] = 0
	}

	return res
}

type mystack struct {
	data []int
}

func (m *mystack) Empty() bool {
	return len(m.data) == 0
}

func (m *mystack) Push(d int) {
	m.data = append(m.data, d)
}

func (m *mystack) Top() int {
	return m.data[len(m.data)-1]
}

func (m *mystack) Pop() int {
	if m.Empty() {
		return -1
	}

	tmp := m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	return tmp
}
