package main

import (
	"fmt"

	"strconv"

	"github.com/noursaadallah/elementary-cycles-golang/graphs"
)

/**
 * Testfile for elementary cycle search.
 *
 * @author Frank Meyer
 *
 */

/**
 * @param args
 */
func main() {
	//String nodes[] = new String[10];
	//nodes := make([]string, 10)
	nodes := make([]int, 10)
	//boolean adjMatrix[][] = new boolean[10][10];
	adjMatrix := make([][]bool, 10)

	for i := 0; i < 10; i++ {
		//nodes[i] = "Node " + strconv.Itoa(i)
		adjMatrix[i] = make([]bool, 10)
	}

	for i := 0; i < 10; i++ {
		//nodes[i] = "Node " + strconv.Itoa(i)
		nodes[i] = i
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

	var ecs *graphs.ElementaryCyclesSearch
	ecs = graphs.NewElementaryCyclesSearch(adjMatrix, nodes)
	cycles := ecs.GetElementaryCycles()
	for i := 0; i < len(cycles); i++ {
		cycle := cycles[i]
		for j := 0; j < len(cycle); j++ {
			var node string
			node = strconv.Itoa(cycle[j])
			if j < len(cycle)-1 {
				fmt.Print(node + " -> ")
			} else {
				fmt.Print(node)
			}
		}
		fmt.Print("\n")
	}
}
