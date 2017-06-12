package graphs

import "math"

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
 * nodes in the subgraph.<br><br>
 *
 * For a description of the algorithm for calculating the strong connected
 * components see:<br>
 * Robert Tarjan: Depth-first search and linear graph algorithms. In: SIAM
 * Journal on Computing. Volume 1, Nr. 2 (1972), pp. 146-160.<br>
 * For a description of the algorithm for searching all elementary cycles in
 * a directed graph see:<br>
 * Donald B. Johnson: Finding All the Elementary Circuits of a Directed Graph.
 * SIAM Journal on Computing. Volumne 4, Nr. 1 (1975), pp. 77-84.<br><br>
 *
 * @author Frank Meyer, web_at_normalisiert_dot_de
 * @version 1.1, 22.03.2009
 *
 */
type StrongConnectedComponents struct {
	/** Adjacency-list of original graph */
	//private int[][] adjListOriginal = null;
	adjListOriginal [][]int

	/** Adjacency-list of currently viewed subgraph */
	//private int[][] adjList = null;
	adjList [][]int
	
	/** Helpattribute for finding scc's */
	//private boolean[] visited = null;
	visited []bool

	/** Helpattribute for finding scc's */
	//private Vector stack = null;
	stack []int

	/** Helpattribute for finding scc's */
	//private int[] lowlink = null;
	lowlink []int

	/** Helpattribute for finding scc's */
	//private int[] number = null;
	number []int

	/** Helpattribute for finding scc's */
	//private int sccCounter = 0;
	sccCounter int

	/** Helpattribute for finding scc's */
	//private Vector currentSCCs = null;
	currentSCCs [][]int
}

	/**
	 * Constructor.
	 *
	 * @param adjList adjacency-list of the graph
	 */
	func NewStrongConnectedComponents(adjList [][]int ) *StrongConnectedComponents {
		this := new(StrongConnectedComponents)
		this.adjListOriginal = adjList
		return this
	}

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
func (this *StrongConnectedComponents) getAdjacencyList(node int ) SCCResult{
		//this.visited = new boolean[this.adjListOriginal.length];
	this.visited = make([]bool , len(this.adjListOriginal) )
		//this.lowlink = new int[this.adjListOriginal.length];
	this.lowlink = make([]int , len(this.adjListOriginal) )
		//this.number = new int[this.adjListOriginal.length];
	this.number = make([]int , len(this.adjListOriginal) )
		//this.visited = new boolean[this.adjListOriginal.length];
	this.visited = make([]bool , len(this.adjListOriginal) )
		//this.stack = new Vector();
		//this.currentSCCs = new Vector();

		this.makeAdjListSubgraph(node)

		for i := node; i < len(this.adjListOriginal) ; i++ {
			if (!this.visited[i]) {
				this.getStrongConnectedComponents(i);
				var nodes []int
				nodes = this.getLowestIdComponent();
				if (nodes != nil && !contains(nodes , node) && !contains(nodes , node+1) ) {
					return this.getAdjacencyList(node + 1);
				} else {
					var adjacencyList []int
					adjacencyList = this.getAdjList(nodes);
					if (adjacencyList != nil) {
						for j := 0; j < len(this.adjListOriginal); j++ {
							if (len(adjacencyList[j]) > 0) {
								result := NewSCCResult(adjacencyList , j)
								return  result
							}
						}
					}
				}
			}
		}

		return nil
	}

	/**
	 * Builds the adjacency-list for a subgraph containing just nodes
	 * >= a given index.
	 *
	 * @param node Node with lowest index in the subgraph
	 */
func (this *StrongConnectedComponents) makeAdjListSubgraph(node int ) {
		this.adjList = make([][]int , len(this.adjListOriginal) )	// = new int[this.adjListOriginal.length][0];

		for  i := node; i < len(this.adjList); i++ {
			var successors []int
			for j := 0; j < len(this.adjListOriginal[i]); j++ {
				if (this.adjListOriginal[i][j] >= node) {
					successors = append (successors , this.adjListOriginal[i][j])
				}
			}
			if (len( successors ) > 0) {
				this.adjList[i] = make ([]int, len(successors))
				for j := 0; j < len(successors); j++ {
					var succ int
					succ = successors[j]
					this.adjList[i][j] = succ
				}
			}
		}
	}

	/**
	 * Calculates the strong connected component out of a set of scc's, that
	 * contains the node with the lowest index.
	 *
	 * @return Vector::Integer of the scc containing the lowest nodenumber
	 */
func (this *StrongConnectedComponents) getLowestIdComponent() []int {
		min := len(this.adjList)
		var currScc []int

		for i := 0; i < len(this.currentSCCs) ; i++ {
			var scc []int
			scc = this.currentSCCs[i]
			for j := 0; j < len(scc); j++ {
				var node int
				node = scc[j]
				if (node < min) {
					currScc = scc
					min = node
				}
			}
		}

		return currScc;
	}

	/**
	 * @return Vector[]::Integer representing the adjacency-structure of the
	 * strong connected component with least vertex in the currently viewed
	 * subgraph
	 */
func (this *StrongConnectedComponents) getAdjList(nodes []int) [][]int {
		//Vector[] lowestIdAdjacencyList = null;
	var lowestIdAdjacencyList [][]int

	if (nodes != nil) {
			//lowestIdAdjacencyList = new Vector[this.adjList.length];
		lowestIdAdjacencyList = make([][]int , len(this.adjList) )
			//for i := 0; i < len(lowestIdAdjacencyList); i++ {
				//lowestIdAdjacencyList[i] = new Vector();
			//}
			for i := 0; i < len(nodes); i++ {
				//int node = ((Integer) nodes.get(i)).intValue();
				node := nodes[i]
				for j := 0; j < len(this.adjList[node]); j++ {
					succ := this.adjList[node][j];
					if (contains(nodes, succ)) {
						lowestIdAdjacencyList[node] = append(lowestIdAdjacencyList[node] , succ)
					}
				}
			}
		}

		return lowestIdAdjacencyList;
	}

	/**
	 * Searchs for strong connected components reachable from a given node.
	 *
	 * @param root node to start from.
	 */
func (this *StrongConnectedComponents) getStrongConnectedComponents(root int) {
		this.sccCounter++
		this.lowlink[root] = this.sccCounter
		this.number[root] = this.sccCounter
		this.visited[root] = true
		this.stack = append(this.stack , root)

		for i := 0; i < len(this.adjList[root]); i++ {
			w := this.adjList[root][i];
			if (!this.visited[w]) {
				this.getStrongConnectedComponents(w);
				//this.lowlink[root] = Math.min(lowlink[root], lowlink[w]);
				this.lowlink[root] = math.Min(this.lowlink[root], this.lowlink[w])
			} else if (this.number[w] < this.number[root]) {
				if (contains(this.stack , w )){
					//lowlink[root] = Math.min(this.lowlink[root], this.number[w]);
					this.lowlink[root] = math.Min(this.lowlink[root], this.number[w])
				}
			}
		}

		// found scc
		if ((this.lowlink[root] == this.number[root]) && (len(this.stack) > 0)) {
			next := -1
			//Vector scc = new Vector();
			var scc []int

			//do {
			//	next = ((Integer) this.stack.get(stack.size() - 1)).intValue();
			//	this.stack.remove(stack.size() - 1);
			//	scc.add(new Integer(next));
			//} while (this.number[next] > this.number[root]);

			for ok := true ; ok ; ok = (this.number[next] > this.number[root]){
				next = this.stack[len(this.stack) - 1]
			}


			// simple scc's with just one node will not be added
			if (len(scc) > 1) {
				this.currentSCCs = append(this.currentSCCs, scc)
			}
		}
	}

//	public static void main(String[] args) {
//		boolean[][] adjMatrix = new boolean[10][];
//
//		for (int i = 0; i < 10; i++) {
//			adjMatrix[i] = new boolean[10];
//		}
//
//		/*adjMatrix[0][1] = true;
//		adjMatrix[1][2] = true;
//		adjMatrix[2][0] = true;
//		adjMatrix[2][4] = true;
//		adjMatrix[1][3] = true;
//		adjMatrix[3][6] = true;
//		adjMatrix[6][5] = true;
//		adjMatrix[5][3] = true;
//		adjMatrix[6][7] = true;
//		adjMatrix[7][8] = true;
//		adjMatrix[7][9] = true;
//		adjMatrix[9][6] = true;*/
//
//        adjMatrix[0][1] = true;
//        adjMatrix[1][2] = true;
//        adjMatrix[2][0] = true; adjMatrix[2][6] = true;
//        adjMatrix[3][4] = true;
//        adjMatrix[4][5] = true; adjMatrix[4][6] = true;
//        adjMatrix[5][3] = true;
//        adjMatrix[6][7] = true;
//        adjMatrix[7][8] = true;
//        adjMatrix[8][6] = true;
//
//        adjMatrix[6][1] = true;
//
//		int[][] adjList = AdjacencyList.getAdjacencyList(adjMatrix);
//		StrongConnectedComponents scc = new StrongConnectedComponents(adjList);
//		for (int i = 0; i < adjList.length; i++) {
//			System.out.print("i: " + i + "\n");
//			SCCResult r = scc.getAdjacencyList(i);
//			if (r != null) {
//				Vector[] al = scc.getAdjacencyList(i).getAdjList();
//				for (int j = i; j < al.length; j++) {
//					if (al[j].size() > 0) {
//						System.out.print("j: " + j);
//						for (int k = 0; k < al[j].size(); k++) {
//							System.out.print(" _" + al[j].get(k).toString());
//						}
//						System.out.print("\n");
//					}
//				}
//				System.out.print("\n");
//			}
//		}
//	}
//}
