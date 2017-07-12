package cycles

// SCCResult : represents the adjacency structure of a set of Strongly Connected Components
type SCCResult struct {
	nodeIDsOfSCC []int //private Set nodeIDsOfSCC = null; // Set is a list without duplicates || apparently works with ordinary list, but we should probably replicate the behaviour of Set
	adjList      [][]int
	lowestNodeId int //private int lowestNodeId = -1;
}

// NewSCCResult : Constructor
func NewSCCResult(adjList [][]int, lowestNodeId int) *SCCResult {
	this := new(SCCResult)
	this.adjList = adjList
	this.lowestNodeId = lowestNodeId
	this.nodeIDsOfSCC = make([]int, 0) //this.nodeIDsOfSCC = new HashSet(); // TODO : replicate behaviour of Set
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
