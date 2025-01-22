package graphs

import (
	"math"
)

// CREATING A PAIR STRUCTURE TO USE WHEN NEEDING PAIRS
type pair struct {
	first  int
	second int
}

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

// TOPOLOGICAL SORT
func topoSort(node int, visited *[]bool, adj [][]int, st *[]int) {
	(*visited)[node] = true
	n := len(adj[node])
	for i := 0; i < n; i++ {
		neighbor := adj[node][i]
		if (*visited)[neighbor] == false {
			topoSort(neighbor, visited, adj, st)
		}
	}
	(*st) = append((*st), node)
}
func topologicalSort(V int, adj [][]int) []int {
	var ans []int
	var st []int
	visited := make([]bool, V)
	for i := 0; i < V; i++ {
		if visited[i] == false {
			topoSort(i, &visited, adj, &st)
		}
	}
	for len(st) > 0 {
		topElement := st[len(st)-1]
		st = st[:len(st)-1]
		ans = append(ans, topElement)
	}
	return ans
}

// KAHNS ALGORITHM FOR TOPOLOGICAL SORT
func kahnsAlgorithms(V int, adj [][]int) []int {
	inDegree := make([]int, V)
	for i := 0; i < V; i++ {
		n := len(adj[i])
		for j := 0; j < n; j++ {
			node := adj[i][j]
			inDegree[node]++
		}
	}
	var q []int
	for i := 0; i < V; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}
	var ans []int
	for len(q) > 0 {
		node := q[0]
		ans = append(ans, node)
		q = q[1:]
		n := len(adj[node])
		for i := 0; i < n; i++ {
			neighbor := adj[node][i]
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				q = append(q, neighbor)
			}
		}
	}
	return ans
}

// CHECK FOR CYCLE IN DIRECTED GRAPH
func checkForCycleDirectedFunction(node int, visited *[]bool, path *[]bool, adj [][]int) bool {
	(*visited)[node] = true
	(*path)[node] = true
	n := len(adj[node])

	for i := 0; i < n; i++ {
		neighbor := adj[node][i]
		if (*visited)[neighbor] == true {
			continue
		}
		if (*path)[node] == true {
			return true
		}
		if (*visited)[node] == false {
			if checkForCycleDirectedFunction(neighbor, visited, path, adj) == true {
				return true
			}
		}
	}
	(*path)[node] = false
	return false
}
func checkForCycleDirected(V int, adj [][]int) bool {
	visited := make([]bool, V)
	path := make([]bool, V)
	for i := 0; i < V; i++ {
		if visited[i] == false {
			if checkForCycleDirectedFunction(i, &visited, &path, adj) == true {
				return true
			}
		}
	}
	return false
}

// CHECK FOR CYCLE IN DIRECTED GRAPH USING KAHNS ALGORITHMS
func checkForCycleDirected2(V int, adj [][]int) bool {
	var indegree []int
	for i := 0; i < V; i++ {
		n := len(adj[i])
		for j := 0; j < n; j++ {
			node := adj[i][j]
			indegree[node]++
		}
	}
	var q []int
	for i := 0; i < V; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}
	var cnt int
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		n := len(adj[node])
		for i := 0; i < n; i++ {
			neighbor := adj[node][i]
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				q = append(q, neighbor)
			}
		}
	}
	if cnt != V {
		return true
	}
	return false
}

// FIND THE SHORTEST DISTANCE FROM SOURCE NODE TO ALL NODES IN UNDIRECTED RAPH
func findShortestDistance(N int, M int, edges [][]int, SRC int) []int {
	distance := make([]int, N)
	for i := 0; i < N; i++ {
		distance[i] = math.MaxInt32
	}
	adj := make([][]int, N)
	for i := 0; i < M; i++ {
		u := edges[i][0]
		v := edges[i][1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	visited := make([]bool, N)
	var q []int
	q = append(q, SRC)
	visited[SRC] = true
	distance[SRC] = 0

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		n := len(adj[node])
		for i := 0; i < n; i++ {
			neighbor := adj[node][i]
			distanceToNeighor := 1 + distance[node]
			if distance[neighbor] > distanceToNeighor {
				distance[neighbor] = distanceToNeighor
			}
			q = append(q, neighbor)
		}
	}
	for i := 0; i < N; i++ {
		if distance[i] == math.MaxInt32 {
			distance[i] = -1
		}
	}
	return distance
}

// FIND THE SHORTEST DISTANCE FROM 0 TO ALL NODES IN DIRECTED ACYCLIC GRAPH
func topoSortFunctionForDAG(node int, visited *[]bool, st *[]int, adj [][]pair) {
	(*visited)[node] = true
	n := len(adj[node])
	for i := 0; i < n; i++ {
		neighbor := adj[node][i].first
		if (*visited)[neighbor] == false {
			topoSortFunctionForDAG(neighbor, visited, st, adj)
		}
	}
	(*st) = append((*st), node)
}
func findShortestDAG(V int, edges [][]int, M int) []int {
	adj := make([][]pair, V)

	for i := 0; i < M; i++ {
		u := edges[i][0]
		v := edges[i][1]
		weight := edges[i][2]

		adj[u] = append(adj[u], pair{v, weight})
	}
	var st []int
	visited := make([]bool, V)
	distance := make([]int, V)
	for i := 0; i < V; i++ {
		distance[i] = math.MaxInt32
	}
	topoSortFunctionForDAG(0, &visited, &st, adj)
	distance[0] = 0
	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]
		n := len(adj[node])
		for i := 0; i < n; i++ {
			neighbor := adj[node][i].first
			weight := adj[node][i].second

			distanceToNeighor := weight + distance[node]
			if distance[neighbor] > distanceToNeighor {
				distance[neighbor] = distanceToNeighor
			}
		}
	}
	for i := 0; i < V; i++ {
		if distance[i] == math.MaxInt32 {
			distance[i] = -1
		}
	}
	return distance
}

// CHECK WHEHTER GRAPH IS BIPARTITE USING DFS
func checkfunctionBipartite(node int, color *[]int, adj [][]int) bool {
	n := len(adj[node])
	for i := 0; i < n; i++ {
		neighbor := adj[node][i]
		if (*color)[neighbor] == -1 {
			(*color)[neighbor] = 1 - (*color)[node]
			if checkfunctionBipartite(neighbor, color, adj) == false {
				return false
			}
		} else {
			if (*color)[neighbor] == (*color)[node] {
				return false
			}
		}
	}
	return true
}
func isBipartiteDFS(V int, adj [][]int) bool {
	color := make([]int, V)
	for i := 0; i < V; i++ {
		color[i] = -1
	}
	for i := 0; i < V; i++ {
		if color[i] == -1 {
			color[i] = 0
			if checkfunctionBipartite(i, &color, adj) == false {
				return false
			}
		}
	}
	return true
}

// CHECK FOR BIPARTITE USING BFS
func isBipartiteBFS(V int, adj [][]int) bool {
	color := make([]int, V)
	for i := 0; i < V; i++ {
		color[i] = -1
	}
	var q []int
	for i := 0; i < V; i++ {
		if color[i] == -1 {
			color[i] = 0
			q = append(q, i)
			for len(q) > 0 {
				node := q[0]
				q = q[1:]
				n := len(adj[node])
				for j := 0; j < n; j++ {
					neighbor := adj[node][j]
					if color[neighbor] == -1 {
						color[neighbor] = 1 - color[node]
						q = append(q, neighbor)
					} else {
						if color[neighbor] == color[node] {
							return false
						}
					}
				}

			}
		}
	}
	return true
}

// DIJKSTRA ALGORITHM
func dijkstra(V int, adj [][][]int, SRC int) []int {
	distance := make([]int, V)
	visited := make([]bool, V)
	for i := 0; i < V; i++ {
		distance[i] = math.MaxInt32
		visited[i] = false
	}
	distance[SRC] = 0
	for i := 0; i < V-1; i++ {
		node := -1
		value := math.MaxInt
		for j := 0; j < V; j++ {
			if visited[j] == false && distance[j] < value {
				node = j
				value = distance[j]
			}
		}
		if node == -1 {
			break
		}
		visited[node] = true
		n := len(adj[node])
		for j := 0; j < n; j++ {
			neighbor := adj[node][j][0]
			weight := adj[node][j][1]

			if distance[neighbor] > distance[node]+weight {
				distance[neighbor] = distance[node] + weight
			}
		}

	}
	for i := 0; i < len(distance); i++ {
		if distance[i] == math.MaxInt32 {
			distance[i] = -1
		}
	}
	return distance
}

// BELLMAN FORD ALGORITHM
func bellmanford(V int, edges [][]int, SRC int) []int {
	distance := make([]int, V)
	for i := 0; i < V; i++ {
		distance[i] = math.MaxInt32
	}
	distance[SRC] = 0
	for i := 0; i < V-1; i++ {
		for j := 0; j < len(edges); j++ {
			u := edges[j][0]
			v := edges[j][1]
			weight := edges[j][2]
			if distance[u] != math.MaxInt32 && distance[v] > weight+distance[u] {
				distance[v] = distance[u] + weight
			}
		}
	}
	for i := 0; i < len(edges); i++ {
		u := edges[i][0]
		v := edges[i][1]
		weight := edges[i][2]

		if distance[u] != math.MaxInt32 && distance[v] > distance[u]+weight {
			var new []int
			new = append(new, -1)
			return new
		}
	}
	for i := 0; i < len(distance); i++ {
		if distance[i] == math.MaxInt32 {
			distance[i] = -1
		}
	}
	return distance
}

// FLOYD WARSHALL ALGORHTM
func floydwarshall(matrix *[][]int) {
	inf := math.MaxInt32 / 2
	for i := 0; i < len(*matrix); i++ {
		for j := 0; j < len(*matrix); j++ {
			if (*matrix)[i][j] == -1 && i != j {
				(*matrix)[i][j] = inf
			}
		}
	}
	for k := 0; k < len(*matrix); k++ {
		for i := 0; i < len(*matrix); i++ {
			for j := 0; j < len(*matrix); j++ {
				if (*matrix)[i][k] != inf && (*matrix)[k][j] != inf {
					if (*matrix)[i][j] > (*matrix)[i][k]+(*matrix)[k][j] {
						(*matrix)[i][j] = (*matrix)[i][k] + (*matrix)[k][j]
					}
				}
			}
		}
	}
	for i := 0; i < len(*matrix); i++ {
		for j := 0; j < len(*matrix); j++ {
			if (*matrix)[i][j] == inf {
				(*matrix)[i][j] = -1
			}
		}
	}
}

// EULER CIRCUIT IN GRAPH
func dfsfunctionEuler(node int, visited *[]bool, adj [][]int) {
	(*visited)[node] = true
	n := len(adj[node])

	for i := 0; i < n; i++ {
		neighbor := adj[node][i]
		if (*visited)[neighbor] == false {
			dfsfunctionEuler(neighbor, visited, adj)
		}
	}
}
func isEuler(V int, adj [][]int) int {
	inDegree := make([]int, V)
	for i := 0; i < V; i++ {
		n := len(adj[i])
		for j := 0; j < n; j++ {
			node := adj[i][j]
			inDegree[node]++
		}
	}
	odd := 0
	for i := 0; i < len(inDegree); i++ {
		if inDegree[i]%2 == 1 {
			odd++
		}
	}
	if odd != 0 && odd != 2 {
		return 0
	}
	visited := make([]bool, V)
	for i := 0; i < V; i++ {
		if visited[i] == false {
			dfsfunctionEuler(i, &visited, adj)
			break
		}
	}
	for i := 0; i < V; i++ {
		if visited[i] == false && inDegree[i] != 0 {
			return 0
		}
	}
	if odd == 0 {
		return 2
	}
	if odd == 2 {
		return 1
	}
	return 0
}

// COVID SPREAD
func isValid(x int, y int, matrix [][]int) bool {
	if x >= 0 && y >= 0 && x < len(matrix) && y < len(matrix[0]) && matrix[x][y] == 1 {
		return true
	}
	return false
}
func covidSpread(matrix [][]int) int {
	rows := [4]int{0, 0, -1, 1}
	cols := [4]int{1, -1, 0, 0}
	var q []pair
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 2 {
				q = append(q, pair{i, j})
			}
		}
	}
	cnt := 0
	for len(q) > 0 {
		spread := false
		n := len(q)
		for i := 0; i < n; i++ {
			x := q[0].first
			y := q[0].second
			q = q[1:]
			for j := 0; j < 4; j++ {
				newX := x + rows[j]
				newY := y + cols[j]

				if isValid(newX, newY, matrix) == true {
					matrix[newX][newY] = 2
					q = append(q, pair{newX, newY})
					spread = true
				}

			}
		}
		if spread == true {
			cnt++
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 2 {
				return -1
			}
		}
	}
	return cnt - 1
}
