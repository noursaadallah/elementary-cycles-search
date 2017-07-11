package graphs

type SCCResult struct {
	//private Set nodeIDsOfSCC = null; // Set is a list without duplicates || apparently works with ordinary list, but we should probably reproduce the behaviour of Set
	nodeIDsOfSCC []int
	//private Vector[] adjList = null;
	adjList [][]int
	//private int lowestNodeId = -1;
	lowestNodeId int
}

// Constructor
func NewSCCResult(adjList [][]int, lowestNodeId int) *SCCResult {
	this := new(SCCResult)
	this.adjList = adjList
	this.lowestNodeId = lowestNodeId
	//this.nodeIDsOfSCC = new HashSet();
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
