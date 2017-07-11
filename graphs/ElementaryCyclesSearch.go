package graphs

/**
 * Searchs all elementary cycles in a given directed graph. The implementation
 * is independent from the concrete objects that represent the graphnodes, it
 * just needs an array of the objects representing the nodes the graph
 * and an adjacency-matrix of type boolean, representing the edges of the
 * graph. It then calculates based on the adjacency-matrix the elementary
 * cycles and returns a list, which contains lists itself with the objects of the
 * concrete graphnodes-implementation. Each of these lists represents an
 * elementary cycle.<br><br>
 *
 * The implementation uses the algorithm of Donald B. Johnson for the search of
 * the elementary cycles. For a description of the algorithm see:<br>
 * Donald B. Johnson: Finding All the Elementary Circuits of a Directed Graph.
 * SIAM Journal on Computing. Volumne 4, Nr. 1 (1975), pp. 77-84.<br><br>
 *
 * The algorithm of Johnson is based on the search for strong connected
 * components in a graph. For a description of this part see:<br>
 * Robert Tarjan: Depth-first search and linear graph algorithms. In: SIAM
 * Journal on Computing. Volume 1, Nr. 2 (1972), pp. 146-160.<br>
 *
 * @author Frank Meyer, web_at_normalisiert_dot_de
 * @version 1.2, 22.03.2009
 *
 */
type ElementaryCyclesSearch struct {
	/** List of cycles */
	// private List<List<Integer>> cycles = null // TODO : look into how to reproduce same behavior
	cycles [][]int

	/** Adjacency-list of graph */
	adjList [][]int

	/** Graphnodes */
	//private Integer[] graphNodes = null;
	graphNodes []int

	/** Blocked nodes, used by the algorithm of Johnson */
	blocked []bool

	/** B-Lists, used by the algorithm of Johnson */
	//private Vector[] B = null;
	B [][]int

	/** Stack for nodes, used by the algorithm of Johnson */
	//private Vector stack = null;
	stack []int
}

/**
 * Constructor.
 *
 * @param matrix adjacency-matrix of the graph
 * @param graphNodes array of the graphnodes of the graph; this is used to
 * build sets of the elementary cycles containing the objects of the original
 * graph-representation
 */
func NewElementaryCyclesSearch(matrix [][]bool, graphNodes []int) *ElementaryCyclesSearch {
	ecs := new(ElementaryCyclesSearch)
	ecs.graphNodes = graphNodes
	ecs.adjList = GetAdjacencyList(matrix)
	return ecs
}

/**
 * Returns List::List::Object with the Lists of nodes of all elementary
 * cycles in the graph.
 *
 * @return List::List::Object with the Lists of the elementary cycles.
 */
func (this *ElementaryCyclesSearch) GetElementaryCycles() [][]int {
	this.cycles = make([][]int, 0)
	this.blocked = make([]bool, len(this.adjList))
	this.B = make([][]int, len(this.adjList))
	this.stack = make([]int, 0)
	var sccs *StrongConnectedComponents
	sccs = NewStrongConnectedComponents(this.adjList)
	s := 0

	for true {
		var sccResult *SCCResult
		sccResult = sccs.getAdjacencyList(s)
		if sccResult != nil && sccResult.getAdjList() != nil {
			var scc [][]int
			scc = sccResult.getAdjList()
			s = sccResult.getLowestNodeId()
			for j := 0; j < len(scc); j++ {
				if (scc[j] != nil) && (len(scc[j]) > 0) {
					this.blocked[j] = false
					this.B[j] = make([]int, 0)
				}
			}

			this.findCycles(s, s, scc)
			s++
		} else {
			break
		}
	}

	return this.cycles
}

/**
 * Calculates the cycles containing a given node in a strongly connected
 * component. The method calls itself recursivly.
 *
 * @param v
 * @param s
 * @param adjList adjacency-list with the subgraph of the strongly
 * connected component s is part of.
 * @return true, if cycle found; false otherwise
 */
func (this *ElementaryCyclesSearch) findCycles(v int, s int, adjList [][]int) bool {
	f := false
	//this.stack[len(this.stack)-1] = v
	this.stack = append(this.stack, v)
	this.blocked[v] = true

	for i := 0; i < len(adjList[v]); i++ {
		w := adjList[v][i]
		// found cycle
		if w == s {
			var cycle []int
			cycle = make([]int, 0)
			for j := 0; j < len(this.stack); j++ {
				index := this.stack[j]
				cycle = append(cycle, this.graphNodes[index])
			}
			this.cycles = append(this.cycles, cycle)
			f = true
		} else if !this.blocked[w] {
			if this.findCycles(w, s, adjList) {
				f = true
			}
		}
	}

	if f {
		this.unblock(v)
	} else {
		for i := 0; i < len(adjList[v]); i++ {
			w := adjList[v][i]
			if !contains(this.B[w], v) {
				this.B[w] = append(this.B[w], v)
			}
		}
	}

	//fmt.Println(len(this.stack)) // = 6
	//fmt.Println(v) // =8
	this.stack = append(this.stack[:v], this.stack[v+1:]...) //this.stack.remove(v) : v is the index to remove
	return f
}

//func remove(s []int, r int) []int {
//	for i, v := range s {
//		if v == r {
//			return append(s[:i], s[i+1:]...)
//		}
//	}
//	return s
//}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

/**
 * Unblocks recursivly all blocked nodes, starting with a given node.
 *
 * @param node node to unblock
 */
func (this *ElementaryCyclesSearch) unblock(node int) {
	this.blocked[node] = false
	Bnode := this.B[node]
	for len(Bnode) > 0 {
		w := Bnode[0]
		Bnode = append(Bnode[:0], Bnode[0+1:]...) // remove first element - TODO : check if element are shifted correctly (i.e index 1 becomes 0 etc.)
		if this.blocked[w] {
			this.unblock(w)
		}
	}
}
