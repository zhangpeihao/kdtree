package kdtree

import (
	"fmt"
	"testing"
)

func TestCreateTree(t *testing.T) {
	// Empty tree
	err, tree := NewTree(nil, 2)
	if err != nil {
		t.Errorf("NewTree(nil), err: %s", err.Error())
	}
	fmt.Println()
	tree.DumpStdout()

	var nodes []*Node
	nodes = append(nodes, &Node{
		Coordinate: &Coordinate{
			Values: []float64{2.0, 3.0},
		},
	})
	tree, err = NewTree(nodes, 2)
	if err != nil {
		t.Errorf("NewTree(nil), err: %s", err.Error())
	}
	fmt.Println()
	tree.DumpStdout()

	nodes = append(nodes, &Node{
		Coordinate: &Coordinate{
			Values: []float64{5.0, 4.0},
		},
	})
	nodes = append(nodes, &Node{
		Coordinate: &Coordinate{
			Values: []float64{9.0, 6.0},
		},
	})
	nodes = append(nodes, &Node{
		Coordinate: &Coordinate{
			Values: []float64{4.0, 7.0},
		},
	})
	nodes = append(nodes, &Node{
		Coordinate: &Coordinate{
			Values: []float64{8.0, 1.0},
		},
	})
	nodes = append(nodes, &Node{
		Coordinate: &Coordinate{
			Values: []float64{7.0, 2.0},
		},
	})
	err, tree = NewTree(nodes, 2)
	if err != nil {
		t.Errorf("NewTree(nil), err: %s", err.Error())
	}
	fmt.Println()
	tree.DumpStdout()

	var retNodes []*Node
	walker := func(node *Node) bool {
		retNodes = append(retNodes, node)
		return false
	}

	err = tree.Search(&Coordinate{Values: []float64{6.0, 3.0}}, 2.0, walker)
	switch err {
	case nil:
		if len(retNodes) == 0 {
			t.Errorf("No nodes found\n")
		} else {
			fmt.Print("Found nodes: ")
			for _, retNode := range retNodes {
				fmt.Print(retNode.Coordinate.String())
			}
			fmt.Println()
		}
	case ErrSearchStopped:
		fmt.Print("Found finished: ")
		for _, retNode := range retNodes {
			fmt.Print(retNode.Coordinate.String())
		}
		fmt.Println()
	default:
		t.Errorf("tree.Search error: %s\n", err.Error())
	}

	retNodes = nil
	err = tree.Search(&Coordinate{Values: []float64{6.0, 3.0}}, 4.000001, walker)
	switch err {
	case nil:
		if len(retNodes) == 0 {
			t.Errorf("No nodes found\n")
		} else {
			fmt.Print("Found nodes: ")
			for _, retNode := range retNodes {
				fmt.Print(retNode.Coordinate.String())
			}
			fmt.Println()
		}
	case ErrSearchStopped:
		fmt.Print("Found finished: ")
		for _, retNode := range retNodes {
			fmt.Print(retNode.Coordinate.String())
		}
		fmt.Println()
	default:
		t.Errorf("tree.Search error: %s\n", err.Error())
	}
}
