package graphs

/**
 * Calculates the adjacency-list for a given adjacency-matrix.
 * 
 * 
 * @author Frank Meyer, web@normalisiert.de
 * @version 1.0, 26.08.2006
 *
 */
	/**
	 * Calculates a adjacency-list for a given array of an adjacency-matrix.
	 *
	 * @param adjacencyMatrix array with the adjacency-matrix that represents
	 * the graph
	 * @return int[][]-array of the adjacency-list of given nodes. The first
	 * dimension in the array represents the same node as in the given
	 * adjacency, the second dimension represents the indicies of those nodes,
	 * that are direct successornodes of the node.
	 */
	func GetAdjacencyList(adjacencyMatrix [][]bool) [][]int {
		var list [][]int
			list = make([]int,len(adjacencyMatrix) )

		for i := 0; i < len(adjacencyMatrix) ; i++ {
			var v []int
			for j := 0; j < len(adjacencyMatrix[i]); j++ {
				if adjacencyMatrix[i][j] {
					v = append(v , j)
				}
			}

			list[i] = make( []int , len(v) )
			for j := 0; j < len(v); j++ {
				in := v[j]
				list[i][j] = in
			}
		}
		
		return list;
	}
