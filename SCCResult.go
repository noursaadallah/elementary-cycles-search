package ElementaryCyclesSearch

// SCCResult : represents the adjacency structure of a set of Strongly Connected Components
type SCCResult struct {
	nodeIDsOfSCC []int // TODO : replicate behaviour of Set
	adjList      [][]int
	lowestNodeId int // TODO : init lowestNodeId to -1
}

func NewSCCResult(adjList [][]int, lowestNodeId int) *SCCResult {
	this := new(SCCResult)
	this.adjList = adjList
	this.lowestNodeId = lowestNodeId
	this.nodeIDsOfSCC = make([]int, 0)
	if this.adjList != nil {
		for i := this.lowestNodeId; i < len(this.adjList); i++ {
			if len(this.adjList[i]) > 0 {
				this.nodeIDsOfSCC = append(this.nodeIDsOfSCC, i)
			}
		}
	}
	return this
}

func (this *SCCResult) getAdjList() [][]int {
	return this.adjList
}

func (this *SCCResult) getLowestNodeId() int {
	return this.lowestNodeId
}
