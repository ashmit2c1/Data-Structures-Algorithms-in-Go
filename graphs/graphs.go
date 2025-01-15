package graphs

// BREADTH FIRST SEARCH IN GRAPH
func breadthFirstSearch(V int, adj [][]int) []int {
	var ans []int
	visited := make([]bool, V)
	var q []int
	q = append(q, 0)
	visited[0] = true
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		n := len(adj[node])
		ans = append(ans, node)
		for i := 0; i < n; i++ {
			neighbor := adj[node][i]
			if visited[neighbor] == false {
				visited[neighbor] = true
				q = append(q, neighbor)
			}
		}
	}
	return ans
}

// DEPTH FIRST SEARCH IN GRAPH
func dfsFunction(node int, visited *[]bool, ans *[]int, adj [][]int) {
	(*visited)[node] = true
	(*ans) = append((*ans), node)
	n := len(adj[node])
	for i := 0; i < n; i++ {
		neighbor := adj[node][i]
		if (*visited)[neighbor] == false {
			dfsFunction(neighbor, visited, ans, adj)
		}
	}
}

func depthFirstSearch(V int, adj [][]int) []int {
	var ans []int
	visited := make([]bool, V)
	start := 0
	dfsFunction(start, &visited, &ans, adj)
	return ans
}

// CHECK FOR CYCLE IN UNDIRECTED GRAPH USING DFS
func checkForCycle(node int, parent int, visited *[]bool, adj [][]int) bool {
	(*visited)[node] = true
	n := len(adj[node])
	for i := 0; i < n; i++ {
		neighbor := adj[node][i]
		if neighbor == parent {
			continue
		}
		if (*visited)[neighbor] == true {
			return true
		}
		if (*visited)[neighbor] == false {
			if checkForCycle(neighbor, node, visited, adj) == true {
				return true
			}
		}
	}
	return false
}
func isCycle(V int, adj [][]int) bool {
	visited := make([]bool, V)
	for i := 0; i < V; i++ {
		if visited[i] == false {
			check := checkForCycle(i, -1, &visited, adj)
			if check == true {
				return true
			}
		}
	}
	return false
}

type pair struct {
	first  int
	second int
}

// CHECK FOR CYCLE IN GRAPH USING BFS
func checkForCycleBFS(node int, visited *[]bool, adj [][]int) bool {
	var q []pair
	q = append(q, pair{node, -1})
	(*visited)[node] = true
	for len(q) > 0 {
		node := q[0].first
		parent := q[0].second
		q = q[1:]
		n := len(adj[node])
		for i := 0; i < n; i++ {
			neighbor := adj[node][i]
			if neighbor == parent {
				continue
			}
			if (*visited)[neighbor] == true {
				return true
			}
			if (*visited)[neighbor] == false {
				q = append(q, pair{neighbor, node})
				(*visited)[neighbor] = true
			}
		}
	}
	return false
}
func isCycleBFS(V int, adj [][]int) bool {
	visited := make([]bool, V)
	for i := 0; i < V; i++ {
		if visited[i] == false {
			check := checkForCycleBFS(i, &visited, adj)
			if check == true {
				return true
			}
		}
	}
	return false
}
