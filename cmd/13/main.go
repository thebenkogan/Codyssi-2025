package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/samber/lo"
	lib "github.com/thebenkogan/Codyssi-2025"
	"github.com/zyedidia/generic/heap"
)

func main() {
	input := lib.GetInput()

	adj1 := make(map[string][]edge)
	adj2 := make(map[string][]edge)
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, " ")
		adj1[split[0]] = append(adj1[split[0]], edge{to: split[2], cost: 1})
		cost := lib.ParseNums(line)[0]
		adj2[split[0]] = append(adj2[split[0]], edge{to: split[2], cost: cost})
	}

	fmt.Println(minPathsProduct(adj1))
	fmt.Println(minPathsProduct(adj2))
	fmt.Println(longestCycle(adj2))
}

type edge struct {
	to   string
	cost int
}

type node struct {
	curr  string
	steps int
}

func minPathsProduct(adj map[string][]edge) int {
	q := heap.New(func(a, b node) bool { return a.steps < b.steps })
	q.Push(node{curr: "STT", steps: 0})
	seen := make(map[string]struct{})
	costs := make(map[string]int)

	for q.Size() > 0 {
		n, _ := q.Pop()
		if _, ok := seen[n.curr]; ok {
			continue
		}
		seen[n.curr] = struct{}{}
		costs[n.curr] = n.steps

		for _, e := range adj[n.curr] {
			q.Push(node{curr: e.to, steps: n.steps + e.cost})
		}
	}

	vals := lo.Values(costs)
	slices.Sort(vals)
	return lo.Product(vals[len(costs)-3:])
}

type cycleNode struct {
	curr   string
	length int
	seen   map[string]struct{}
}

func longestCycle(adj map[string][]edge) int {
	best := 0
	for start := range adj {
		q := []cycleNode{{curr: start, length: 0, seen: make(map[string]struct{})}}
		for len(q) > 0 {
			n := q[0]
			q = q[1:]
			if n.length > 0 && n.curr == start {
				best = max(best, n.length)
				continue
			}
			for _, e := range adj[n.curr] {
				if _, ok := n.seen[e.to]; ok {
					continue
				}
				newSeen := maps.Clone(n.seen)
				newSeen[e.to] = struct{}{}
				q = append(q, cycleNode{curr: e.to, length: n.length + e.cost, seen: newSeen})
			}
		}
	}
	return best
}
