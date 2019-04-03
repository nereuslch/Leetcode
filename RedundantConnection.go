package Leetcode

func findRedundantConnection(edges [][]int) []int {
	union := NewUnionSet(len(edges) + 1)
	last := make([]int, 2)
	for _, edge := range edges {
		if !union.connected(edge[0], edge[1]) {
			union.join(edge[0], edge[1])
			continue
		}
		copy(last, edge)
	}
	return last
}

type UnionSet []int

func NewUnionSet(n int) UnionSet {
	var u UnionSet
	u = make([]int, n)
	for i := 0; i < len(u); i ++ {
		u[i] = i
	}
	return u

}

func (u UnionSet) find(i int) int {
	tmp := i
	for u[tmp] != tmp {
		tmp = u[tmp]
	}

	j := i
	for j != tmp {
		tt := u[j]
		u[j] = tmp
		j = tt
	}

	return tmp
}

func (u UnionSet) connected(i, j int) bool {
	return u.find(i) == u.find(j)
}

func (u UnionSet) counts() int {
	var count int
	for idx, rec := range u {
		if idx == rec {
			count++
		}
	}
	return count
}

func (u UnionSet) join(i, j int) {
	x, y := u.find(i), u.find(j)
	if x != y {
		if y > x {
			u[x] = y
		} else {
			u[y] = x
		}
	}
}
