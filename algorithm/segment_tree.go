package algorithm

type SegTreeNode struct {
	lazy int
	l, r int
	sum  int
}

type SegTree []*SegTreeNode

func BuildSegTree(n int) SegTree {
	s := make(SegTree, n*4)
	s.built(0, n-1, 0)

	return s
}

func (s SegTree) built(l, r, i int) {
	s[i] = &SegTreeNode{
		l: l,
		r: r,
	}
	if r == l {
		return
	}

	mid := (l + r) / 2
	s.built(l, mid, i*2+1)
	s.built(mid+1, r, i*2+2)
}

func (s SegTree) passdown(i int) {
	n := s[i]

	if n.lazy == 0 {
		return
	}

	ll, rr := i*2+1, i*2+2
	s[ll].lazy += n.lazy
	s[rr].lazy += n.lazy

	s[ll].sum += n.lazy * (s[ll].r - s[ll].l + 1)
	s[rr].sum += n.lazy * (s[rr].r - s[rr].l + 1)

	n.lazy = 0
}

func (s SegTree) RangeAdd(l, r, v int) {
	s.rangeAdd(l, r, v, 0)
}

func (s SegTree) rangeAdd(l, r, v, i int) {
	n := s[i]

	if outRange(l, r, n) {
		return
	}

	ll, rr := maxInt(l, n.l), minInt(r, n.r)
	n.sum += (rr - ll + 1) * v

	if inRange(l, r, n) {
		n.lazy += v
		return
	}

	s.passdown(i)
	s.rangeAdd(l, r, v, i*2+1)
	s.rangeAdd(l, r, v, i*2+2)
}

func (s SegTree) GetRange(l, r int) int {
	return s.getRange(l, r, 0)
}

func (s SegTree) getRange(l, r, i int) int {
	n := s[i]

	if inRange(l, r, n) {
		return n.sum
	}

	if outRange(l, r, n) {
		return 0
	}

	s.passdown(i)
	return s.getRange(l, r, i*2+1) + s.getRange(l, r, i*2+2)
}

func inRange(l, r int, n *SegTreeNode) bool {
	return l <= n.l && n.r <= r
}

func outRange(l, r int, n *SegTreeNode) bool {
	return n.l > r || n.r < l
}

func maxInt(l, r int) int {
	if l > r {
		return l
	}
	return r
}

func minInt(l, r int) int {
	if l < r {
		return l
	}
	return r
}