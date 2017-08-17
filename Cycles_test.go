package ElementaryCyclesSearch

import (
	"fmt"
	"testing"
)

func TestCycles1(t *testing.T) {

	// init 2D array
	adjMatrix := make([][]bool, 10)
	for i := 0; i < 10; i++ {
		adjMatrix[i] = make([]bool, 10)
	}

	// init nodes
	nodes := make([]int, 10)
	for i := 0; i < 10; i++ {
		nodes[i] = i
	}

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
	var ecs *ElementaryCyclesSearch
	ecs = NewElementaryCyclesSearch(adjMatrix, nodes)
	cycles := ecs.GetElementaryCycles()

	if cycles == nil {
		t.Log("result shouldn't be nil")
		t.Fail()
	}

	if len(cycles) != 4 {
		t.Log("result should contain 4 cycles")
		t.Fail()
	}

	result := fmt.Sprint(cycles)
	if result != "[[0 1 2] [1 2 6] [3 4 5] [6 7 8]]" {
		t.Log("incorrect result")
		t.Fail()
	}

}

func TestCycles2(t *testing.T) {

	// init 2D array
	adjMatrix := make([][]bool, 10)
	for i := 0; i < 10; i++ {
		adjMatrix[i] = make([]bool, 10)
	}

	// init nodes
	nodes := make([]int, 10)
	for i := 0; i < 10; i++ {
		nodes[i] = i
	}

	adjMatrix[0][1] = true
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
	adjMatrix[9][6] = true

	// search for cycles using the adjMatrix and the list of nodes
	var ecs *ElementaryCyclesSearch
	ecs = NewElementaryCyclesSearch(adjMatrix, nodes)
	cycles := ecs.GetElementaryCycles()

	if cycles == nil {
		t.Log("result shouldn't be nil")
		t.Fail()
	}

	if len(cycles) != 3 {
		t.Log("result should contain 3 cycles")
		t.Fail()
	}

	result := fmt.Sprint(cycles)
	if result != "[[0 1 2] [3 6 5] [6 7 9]]" {
		t.Log("incorrect result")
		t.Fail()
	}

}

func TestCycles3(t *testing.T) {

	// init 2D array
	adjMatrix := make([][]bool, 10)
	for i := 0; i < 10; i++ {
		adjMatrix[i] = make([]bool, 10)
	}

	// init nodes
	nodes := make([]int, 10)
	for i := 0; i < 10; i++ {
		nodes[i] = i
	}

	// search for cycles using the adjMatrix and the list of nodes
	var ecs *ElementaryCyclesSearch
	ecs = NewElementaryCyclesSearch(adjMatrix, nodes)
	cycles := ecs.GetElementaryCycles()

	if cycles == nil {
		t.Log("result shouldn't be nil")
		t.Fail()
	}

	if len(cycles) != 0 {
		t.Log("result should contain 0 cycles")
		t.Fail()
	}

}

func TestCycles4(t *testing.T) {

	// init 2D array
	adjMatrix := make([][]bool, 10)
	for i := 0; i < 10; i++ {
		adjMatrix[i] = make([]bool, 10)
	}

	// init nodes
	nodes := make([]int, 10)
	for i := 0; i < 10; i++ {
		nodes[i] = i
	}

	adjMatrix[0][1] = true
	adjMatrix[1][0] = true

	// search for cycles using the adjMatrix and the list of nodes
	var ecs *ElementaryCyclesSearch
	ecs = NewElementaryCyclesSearch(adjMatrix, nodes)
	cycles := ecs.GetElementaryCycles()

	if cycles == nil {
		t.Log("result shouldn't be nil")
		t.Fail()
	}

	if len(cycles) != 1 {
		t.Log("result should contain 4 cycles")
		t.Fail()
	}

	result := fmt.Sprint(cycles)
	if result != "[[0 1]]" {
		t.Log("incorrect result")
		t.Fail()
	}

}
