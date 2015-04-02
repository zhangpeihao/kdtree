package kdtree

import (
	"io"
	"os"
	"sort"
)

type Tree struct {
	root       *Node
	dimensions int
}

func NewTree(nodes []*Node, dimensions int) (tree *Tree, err error) {
	var root *Node
	root, err = createTree(nodes, dimensions, 0, nil)
	if err != nil {
		return
	}
	tree = &Tree{
		root:       root,
		dimensions: dimensions,
	}
	return
}

type WalkFunc func(*Node) bool

func (tree *Tree) Search(center *Coordinate, radius float64, walker WalkFunc) (err error) {
	if tree.dimensions != center.Dimensions() {
		return ErrDimensionUnmatch
	}
	if tree.root != nil {
		return tree.root.search(center, radius, walker)
	} else {
		return nil
	}
}

func createTree(nodes []*Node, dimensions, depth int, parent *Node) (root *Node, err error) {
	nodes_len := len(nodes)
	switch nodes_len {
	case 0:
		root = nil
	case 1:
		if dimensions != nodes[0].Coordinate.Dimensions() {
			return nil, ErrDimensionUnmatch
		}
		root = nodes[0]
		root.axis = depth % dimensions
		root.parent = parent
		root.leftLeaf = nil
		root.rightLeaf = nil
	default:
		if dimensions != nodes[0].Coordinate.Dimensions() {
			return nil, ErrDimensionUnmatch
		}
		median := (nodes_len / 2)

		nodeList := new(NodeList)
		nodeList.Axis = depth % dimensions
		nodeList.Nodes = make([]*Node, nodes_len)
		copy(nodeList.Nodes, nodes)
		sort.Sort(nodeList)

		root = nodeList.Nodes[median]

		root.axis = nodeList.Axis
		root.parent = parent
		root.leftLeaf, err = createTree(nodeList.Nodes[0:median], dimensions, depth+1, root)
		if err != nil {
			return nil, err
		}
		root.rightLeaf, err = createTree(nodeList.Nodes[median+1:], dimensions, depth+1, root)
		if err != nil {
			return nil, err
		}
	}

	return
}

func (tree *Tree) DumpStdout() {
	tree.Dump(os.Stdout)
}
func (tree *Tree) Dump(w io.Writer) {
	if tree.root != nil {
		tree.root.Dump(w, " ")
	} else {
		io.WriteString(w, "+nil\n")
	}
}
