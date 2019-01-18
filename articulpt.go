// Package for articulation point discovery in the graphs using Tarjan's algorithm.
// First edge has to have index=0.
package articulpt

type Graph struct {
	vertexCount int
	time        int
	edges       map[int][]int

	visited []bool // visited vertices
	disc    []int  // discovery times of visited vertices
	low     []int  // low time
	parent  []int  // parent vertices of DFS tree
	ap      []bool // articulation points
}

func NewGraph(vertexCount int) *Graph {
	return &Graph{
		vertexCount: vertexCount,
		edges:       make(map[int][]int),

		visited: make([]bool, vertexCount),
		disc:    make([]int, vertexCount),
		low:     make([]int, vertexCount),
		parent:  make([]int, vertexCount),
		ap:      make([]bool, vertexCount),
	}
}

func (gp *Graph) AddEdge(v1, v2 int) {
	gp.edges[v1] = append(gp.edges[v1], v2)
	gp.edges[v2] = append(gp.edges[v2], v1)
}

func (gp *Graph) dfs(u int) {
	// Count of children in DFS tree.
	children := 0
	// Mark the current node as visited.
	gp.visited[u] = true
	// Init discovery time and low value.
	gp.time++
	gp.disc[u] = gp.time
	gp.low[u] = gp.time

	for _, v := range gp.edges[u] {

		// if v is not visited yet, then make it a child of u and recur for it
		if !gp.visited[v] {
			children++
			gp.parent[v] = u
			gp.dfs(v) // recur

			// Check if the subtree rooted with v has a connection to one of the ancestors of u
			gp.low[u] = min(gp.low[u], gp.low[v])

			// u is an articulation point in following cases
			// (1) u is root of DFS tree and has two or more children
			if gp.parent[u] == -1 && children > 1 {
				gp.ap[u] = true
			}

			// (2) if u is not root and low value of one of its child is more
			// than discovery value of u.
			if gp.parent[u] != -1 && gp.low[v] >= gp.disc[u] {
				gp.ap[u] = true
			}
		} else if v != gp.parent[u] {
			// Update low value of u for parent function calls.
			gp.low[u] = min(gp.low[u], gp.disc[v])
		}
	}
}

// FindAP finds articulation points.
func (gp *Graph) FindAP() (ret []int) {
	for i := 0; i < gp.vertexCount; i++ {
		gp.parent[i] = -1
	}

	// Run DFS on all yet unvisited vertices.
	for i := 0; i < gp.vertexCount; i++ {
		if !gp.visited[i] {
			gp.dfs(i)
		}
	}

	// Gather results.
	for i, v := range gp.ap {
		if v {
			ret = append(ret, i)
		}
	}
	return
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}
