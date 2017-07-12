package main

import (
	"fmt"

	"github.com/noursaadallah/elementary-cycles-golang/cycles"

	"strconv"
)

/**
 * Testfile for elementary cycle search.
 *
 * Based on work of :
 * @author Frank Meyer
 *
 */

/**
 * @param args
 */
func main() {

	// init 2D array
	adjMatrix := make([][]bool, 10)
	for i := 0; i < 10; i++ {
		adjMatrix[i] = make([]bool, 10)
	}

	// init nodes
	nodes := make([]int, 10) //nodes := make([]string, 10)
	for i := 0; i < 10; i++ {
		//nodes[i] = "Node " + strconv.Itoa(i)
		nodes[i] = i
	}

	/*adjMatrix[0][1] = true
	adjMatrix[1][2] = true
	adjMatrix[2][0] = true
	adjMatrix[2][4] = true
	adjMatrix[1][3] = true
	adjMatrix[3][6] = true
	adjMatrix[6][5] = true
	adjMatrix[5][3] = true
	adjMatrix[6][7] = true
	adjMatrix[7][8] = true
	adjMatrix[7][9] = true
	adjMatrix[9][6] = true*/

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

	// search for cycles using the adjMatrix and the list of nodes
	var ecs *cycles.ElementaryCyclesSearch
	ecs = cycles.NewElementaryCyclesSearch(adjMatrix, nodes)
	cycles := ecs.GetElementaryCycles()
	// range through cycles and print them
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
