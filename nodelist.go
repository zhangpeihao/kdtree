package kdtree

type NodeList struct {
	Axis  int
	Nodes []*Node
}

func (nodeList *NodeList) Len() int {
	return len(nodeList.Nodes)
}

func (nodeList *NodeList) Less(i, j int) bool {
	nodeLen := len(nodeList.Nodes)
	if nodeLen <= i {
		return false
	}
	if nodeLen <= j {
		return true
	}
	co1 := nodeList.Nodes[i].Coordinate
	co2 := nodeList.Nodes[j].Coordinate
	return co1.Values[nodeList.Axis] < co2.Values[nodeList.Axis]
}

func (nodeList *NodeList) Swap(i, j int) {
	nodeLen := len(nodeList.Nodes)
	if nodeLen <= i || nodeLen <= j {
		return
	}
	nodeList.Nodes[i], nodeList.Nodes[j] = nodeList.Nodes[j], nodeList.Nodes[i]
}
