package main

import (
	"flag"
	"fmt"
	. "github.com/zhangpeihao/kdtree"
	"math/rand"
	"time"
)

var (
	number *int     = flag.Int("Number", 1000000, "The number of nodes.")
	search *int     = flag.Int("Search", 10000, "The search times.")
	radius *float64 = flag.Float64("Redius", 2.0, "The search radius.")
)

type user struct {
	id   int
	x, y float64
}

func main() {
	flag.Parse()
	users := make(map[int]*user, *number)
	var nodes []*Node
	var err error
	var tree *Tree
	rand.Seed(time.Now().Unix() % 1234567)
	for i := 0; i < *number; i++ {
		users[i] = &user{
			id: i,
			x:  float64(rand.Int()%3600000)/10000.0 - 180.0,
			y:  float64(rand.Int()%1800000)/10000.0 - 90.0,
		}
	}
	startAt := time.Now()
	for _, user := range users {
		nodes = append(nodes, &Node{
			Coordinate: &Coordinate{
				Values: []float64{user.x, user.y},
			},
		})
	}
	cost := float64(time.Now().Sub(startAt))
	fmt.Printf("Load %d nodes cost %.3fs\n", *number, cost/1000000000.0)
	startAt = time.Now()
	if err, tree = NewTree(nodes, 2); err != nil {
		fmt.Printf("NewTree err: %s\n", err.Error())
		return
	}
	cost = float64(time.Now().Sub(startAt))
	fmt.Printf("%d - %.3fs - %.3fns/op\n", *number, cost/1000000000.0, cost/float64(*number))

	var retNodes []*Node
	walker := func(node *Node) bool {
		retNodes = append(retNodes, node)
		if len(retNodes) >= 25 {
			return true
		}
		return false
	}
	var center *Coordinate
	var err_counter, unfinished_counter, unfound_counter int
	startAt = time.Now()
	for i := 0; i < *search; i++ {
		center = &Coordinate{
			Values: []float64{
				float64(rand.Int()%3600000)/10000.0 - 180.0,
				float64(rand.Int()%1800000)/10000.0 - 90.0,
			},
		}
		err = tree.Search(center, *radius, walker)
		switch err {
		case nil:
			if len(retNodes) == 0 {
				unfound_counter++
			} else {
				unfinished_counter++
			}
		case ErrSearchStopped:
		default:
			err_counter++
		}
	}
	cost = float64(time.Now().Sub(startAt))
	fmt.Printf("%d - search cost %fs, %.3fns/op\n", *search, cost/1000000000.0, cost/float64(*search))
	fmt.Printf("err: (%d)%.3f%%, unfinished: (%d)%.3f%%, unfound: (%d)%.3f%%\n",
		err_counter, 100.0*float64(err_counter)/float64(*search),
		unfinished_counter, 100.0*float64(unfinished_counter)/float64(*search),
		unfound_counter, 100.0*float64(unfound_counter)/float64(*search))
}
