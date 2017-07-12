package graphs

import (
	"fmt"
	"math"
)

/**
 * This is a helpclass for the search of all elementary cycles in a graph
 * with the algorithm of Johnson. For this it searches for strong connected
 * components, using the algorithm of Tarjan. The constructor gets an
 * adjacency-list of a graph. Based on this graph, it gets a nodenumber s,
 * for which it calculates the subgraph, containing all nodes
 * {s, s + 1, ..., n}, where n is the highest nodenumber in the original
 * graph (e.g. it builds a subgraph with all nodes with higher or same
 * nodenumbers like the given node s). It returns the strong connected
 * component of this subgraph which contains the lowest nodenumber of all
 * nodes in the subgraph.
 *
 * For a description of the algorithm for calculating the strong connected
 * components see:
 * Robert Tarjan: Depth-first search and linear graph algorithms. In: SIAM
 * Journal on Computing. Volume 1, Nr. 2 (1972), pp. 146-160.
 * For a description of the algorithm for searching all elementary cycles in
 * a directed graph see:
 * Donald B. Johnson: Finding All the Elementary Circuits of a Directed Graph.
 * SIAM Journal on Computing. Volumne 4, Nr. 1 (1975), pp. 77-84.
 *
 * This is based on the Java implementation of :
 * @author Frank Meyer, web_at_normalisiert_dot_de
 * @version 1.1, 22.03.2009
 *
 */

// StrongConnectedComponents : represent a set of scc's
type StrongConnectedComponents struct {
	/** Adjacency-list of original graph */
	adjListOriginal [][]int

	/** Adjacency-list of currently viewed subgraph */
	adjList [][]int

	/** Helpattribute for finding scc's */
	visited []bool

	/** Helpattribute for finding scc's */
	stack []int

	/** Helpattribute for finding scc's */
	lowlink []int

	/** Helpattribute for finding scc's */
	number []int

	/** Helpattribute for finding scc's */
	sccCounter int

	/** Helpattribute for finding scc's */
	currentSCCs [][]int
}

// NewStrongConnectedComponents :
/**
 * Constructor.
 *
 * @param adjList adjacency-list of the graph
 */
func NewStrongConnectedComponents(adjList [][]int) *StrongConnectedComponents {
	this := new(StrongConnectedComponents)
	this.adjListOriginal = adjList
	return this
}

// getAdjacencyList :
/**
 * This method returns the adjacency-structure of the strong connected
 * component with the least vertex in a subgraph of the original graph
 * induced by the nodes {s, s + 1, ..., n}, where s is a given node. Note
 * that trivial strong connected components with just one node will not
 * be returned.
 *
 * @param node node s
 * @return SCCResult with adjacency-structure of the strong
 * connected component; null, if no such component exists
 */
func (this *StrongConnectedComponents) getAdjacencyList(node int) *SCCResult {

	this.visited = make([]bool, len(this.adjListOriginal))
	this.lowlink = make([]int, len(this.adjListOriginal))
	this.number = make([]int, len(this.adjListOriginal))
	this.visited = make([]bool, len(this.adjListOriginal))
	this.stack = make([]int, 0)
	this.currentSCCs = make([][]int, 0)

	this.makeAdjListSubgraph(node)

	for i := node; i < len(this.adjListOriginal); i++ {
		if !this.visited[i] {
			this.getStrongConnectedComponents(i)
			var nodes []int
			nodes = this.getLowestIdComponent()
			if nodes != nil && !contains(nodes, node) && !contains(nodes, node+1) {
				return this.getAdjacencyList(node + 1)
			} else {
				var adjacencyList [][]int
				adjacencyList = this.getAdjList(nodes)
				if adjacencyList != nil {
					for j := 0; j < len(this.adjListOriginal); j++ {
						if len(adjacencyList[j]) > 0 {
							result := NewSCCResult(adjacencyList, j)
							return result
						}
					}
				}
			}
		}
	}

	return nil
}

// makeAdjListSubgraph :
/**
 * Builds the adjacency-list for a subgraph containing just nodes
 * >= a given index.
 *
 * @param node Node with lowest index in the subgraph
 */
func (this *StrongConnectedComponents) makeAdjListSubgraph(node int) {
	this.adjList = make([][]int, len(this.adjListOriginal)) // = new int[this.adjListOriginal.length][0];
	for i := range this.adjList {
		this.adjList[i] = make([]int, 0)
	}

	for i := node; i < len(this.adjList); i++ {
		var successors []int
		successors = make([]int, 0)
		for j := 0; j < len(this.adjListOriginal[i]); j++ {
			if this.adjListOriginal[i][j] >= node {
				successors = append(successors, this.adjListOriginal[i][j])
			}
		}
		if len(successors) > 0 {
			this.adjList[i] = make([]int, len(successors))
			for j := 0; j < len(successors); j++ {
				var succ int
				succ = successors[j]
				this.adjList[i][j] = succ
			}
		}
	}
}

// getLowestIdComponent :
/**
 * Calculates the strong connected component out of a set of scc's, that
 * contains the node with the lowest index.
 *
 * @return Vector::Integer of the scc containing the lowest nodenumber
 */
func (this *StrongConnectedComponents) getLowestIdComponent() []int {
	min := len(this.adjList)
	var currScc []int
	currScc = nil

	for i := 0; i < len(this.currentSCCs); i++ {
		var scc []int
		scc = this.currentSCCs[i]
		for j := 0; j < len(scc); j++ {
			var node int
			node = scc[j]
			if node < min {
				currScc = scc
				min = node
			}
		}
	}

	return currScc
}

// getAdjList
/**
 * @return Vector[]::Integer representing the adjacency-structure of the
 * strong connected component with least vertex in the currently viewed
 * subgraph
 */
func (this *StrongConnectedComponents) getAdjList(nodes []int) [][]int {
	//Vector[] lowestIdAdjacencyList = null;
	var lowestIdAdjacencyList [][]int
	lowestIdAdjacencyList = nil

	if nodes != nil {
		lowestIdAdjacencyList = make([][]int, len(this.adjList))
		for i := 0; i < len(lowestIdAdjacencyList); i++ {
			lowestIdAdjacencyList[i] = make([]int, 0)
		}
		for i := 0; i < len(nodes); i++ {
			node := nodes[i]
			for j := 0; j < len(this.adjList[node]); j++ {
				succ := this.adjList[node][j]
				if contains(nodes, succ) {
					lowestIdAdjacencyList[node] = append(lowestIdAdjacencyList[node], succ)
				}
			}
		}
	}

	return lowestIdAdjacencyList
}

// getStrongConnectedComponents :
/**
 * Searches for strong connected components reachable from a given node.
 *
 * @param root node to start from.
 */
func (this *StrongConnectedComponents) getStrongConnectedComponents(root int) {
	this.sccCounter++
	this.lowlink[root] = this.sccCounter
	this.number[root] = this.sccCounter
	this.visited[root] = true
	this.stack = append(this.stack, root)

	for i := 0; i < len(this.adjList[root]); i++ {
		w := this.adjList[root][i]
		if !this.visited[w] {
			this.getStrongConnectedComponents(w)
			this.lowlink[root] = int(math.Min(float64(this.lowlink[root]), float64(this.lowlink[w])))
		} else if this.number[w] < this.number[root] {
			if contains(this.stack, w) {
				this.lowlink[root] = int(math.Min(float64(this.lowlink[root]), float64(this.number[w])))
			}
		}
	}

	// found scc
	if (this.lowlink[root] == this.number[root]) && (len(this.stack) > 0) {
		next := -1
		var scc []int
		scc = make([]int, 0)

		// do while equivalent
		for ok := true; ok; ok = (this.number[next] > this.number[root]) {
			next = this.stack[len(this.stack)-1]
			this.stack = this.stack[:len(this.stack)-1]
			scc = append(scc, next)
		}

		// simple scc's with just one node will not be added
		if len(scc) > 1 {
			this.currentSCCs = append(this.currentSCCs, scc)
		}
	}
}

func main() {
	adjMatrix := make([][]bool, 10)

	for i := 0; i < 10; i++ {
		adjMatrix[i] = make([]bool, 10)
	}

	/*adjMatrix[0][1] = true;
	adjMatrix[1][2] = true;
	adjMatrix[2][0] = true;
	adjMatrix[2][4] = true;
	adjMatrix[1][3] = true;
	adjMatrix[3][6] = true;
	adjMatrix[6][5] = true;
	adjMatrix[5][3] = true;
	adjMatrix[6][7] = true;
	adjMatrix[7][8] = true;
	adjMatrix[7][9] = true;
	adjMatrix[9][6] = true;*/

	adjMatrix[0][1] = true
	adjMatrix[1][2] = true
	adjMatrix[2][0] = true
	adjMatrix[2][6] = true
	adjMatrix[3][4] = true
	adjMatrix[4][5] = true
	adjMatrix[4][6] = true
	adjMatrix[5][3] = true
	adjMatrix[6][7] = true
	adjMatrix[7][8] = true
	adjMatrix[8][6] = true

	adjMatrix[6][1] = true

	var adjList [][]int
	adjList = GetAdjacencyList(adjMatrix)

	var scc *StrongConnectedComponents
	scc = NewStrongConnectedComponents(adjList)

	for i := 0; i < len(adjList); i++ {
		fmt.Print("i: ", i, "\n")
		var r *SCCResult
		r = scc.getAdjacencyList(i)
		if r != nil {
			var al [][]int
			al = scc.getAdjacencyList(i).getAdjList()
			for j := i; j < len(al); j++ {
				if len(al[j]) > 0 {
					fmt.Print("j: ", j)
					for k := 0; k < len(al[j]); k++ {
						fmt.Print(" _", al[j][k])
					}
					fmt.Print("\n")
				}
			}
			fmt.Print("\n")
		}
	}
}
