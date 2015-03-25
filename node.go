package kdtree

import (
	"io"
	"os"
)

type Node struct {
	Coordinate *Coordinate
	Context    NodeContext
	axis       int
	parent     *Node
	leftLeaf   *Node
	rightLeaf  *Node
}

func (node *Node) DumpStdout(prefix string) (err error) {
	return node.Dump(os.Stdout, prefix)
}

func (node *Node) Dump(w io.Writer, prefix string) (err error) {
	str := `+` + node.Coordinate.String()
	strlen := len(str)
	//	str = prefix + str
	if _, err = io.WriteString(w, str); err != nil {
		return
	}
	var space string
	var minus string
	for i := 0; i < strlen-1; i++ {
		space += ` `
		minus += `-`
	}

	if node.leftLeaf != nil {
		if err = node.leftLeaf.Dump(w, prefix+space+`|`); err != nil {
			return
		}
	} else {
		if _, err = io.WriteString(w, "+nil\n"); err != nil {
			return
		}
	}
	if _, err = io.WriteString(w, prefix+space); err != nil {
		return
	}
	if node.rightLeaf != nil {
		if err = node.rightLeaf.Dump(w, prefix+` `+space); err != nil {
			return
		}
	} else {
		if _, err = io.WriteString(w, "+nil\n"); err != nil {
			return
		}
	}
	return
}

func (node *Node) search(center *Coordinate, radius float64, walker WalkFunc) (err error) {
	if center.Values[node.axis] <= node.Coordinate.Values[node.axis] {
		if node.leftLeaf != nil {
			err = node.leftLeaf.search(center, radius, walker)
			if err != nil {
				return err
			}
		}
		if center.DistanceTo(node.Coordinate) <= radius && walker(node) {
			// Finished
			return ErrSearchStopped
		}
		if node.rightLeaf != nil && center.Values[node.axis]+radius >= node.Coordinate.Values[node.axis] {
			err = node.rightLeaf.search(center, radius, walker)
			if err != nil {
				return err
			}
		}
	} else {
		if node.rightLeaf != nil {
			err = node.rightLeaf.search(center, radius, walker)
			if err != nil {
				return err
			}
		}
		if center.DistanceTo(node.Coordinate) < radius && walker(node) {
			// Finished
			return ErrSearchStopped
		}
		if node.leftLeaf != nil && center.Values[node.axis]-radius <= node.Coordinate.Values[node.axis] {
			err = node.leftLeaf.search(center, radius, walker)
			if err != nil {
				return err
			}
		}

	}
	return
}
