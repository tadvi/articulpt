package articulpt

import "testing"

func TestAP(t *testing.T) {

	g1 := NewGraph(4)
	g1.AddEdge(0, 1)
	g1.AddEdge(1, 2)
	g1.AddEdge(2, 3)
	g1.AddEdge(0, 2)

	ret := g1.FindAP()
	want := []int{2}

	if !sameValues(ret, want) {
		// expect [2]
		t.Errorf("got: %v, want: %v", ret, want)
	}

	g2 := NewGraph(8)
	g2.AddEdge(0, 1)
	g2.AddEdge(1, 2)
	g2.AddEdge(3, 2)
	g2.AddEdge(0, 2)
	g2.AddEdge(4, 3)
	g2.AddEdge(5, 4)
	g2.AddEdge(5, 6)
	g2.AddEdge(4, 6)
	g2.AddEdge(7, 6)

	ret = g2.FindAP()
	want = []int{2, 3, 4, 6}

	if !sameValues(ret, want) {
		// expect [2, 3, 4, 6]
		t.Errorf("got: %v, want: %v", ret, want)
	}
}

func sameValues(ret, want []int) bool {
	if len(ret) != len(want) {
		return false
	}

Found:
	for _, v1 := range ret {
		for _, v2 := range want {
			if v1 == v2 {
				continue Found
			}
		}
		return false
	}
	return true
}
