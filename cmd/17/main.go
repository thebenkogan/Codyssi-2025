package main

import (
	"fmt"
	"maps"
	"math/big"
	"slices"
	"strings"

	lib "github.com/thebenkogan/Codyssi-2025"
	"github.com/zyedidia/generic/queue"
	"github.com/zyedidia/generic/set"
)

type staircase struct {
	id, start, end int
}

type node struct {
	id, step, length int
}

func fromHash(s string) node {
	ns := lib.ParseNums(s)
	return node{id: ns[0], step: ns[1]}
}

func (n node) hash() string {
	return fmt.Sprintf("S%d_%d", n.id, n.step)
}

func main() {
	input := lib.GetInput()
	sections := strings.Split(input, "\n\n")

	lines := strings.Split(sections[0], "\n")
	moves := lib.ParseNums(sections[1])
	fmt.Println(countPathsInStaircase(lib.ParseNums(lines[0])[2], moves))

	staircases := make(map[int]staircase)
	adj := make(map[string][]node)
	for _, line := range lines {
		ns := lib.ParseNums(line)
		s, startStep, endStep := ns[0], ns[1], ns[2]
		staircases[s] = staircase{id: s, start: startStep, end: endStep}
		for i := startStep; i < endStep; i++ {
			fromHash := node{id: s, step: i}.hash()
			to := node{id: s, step: i + 1}
			adj[fromHash] = append(adj[fromHash], to)
		}
		if len(ns) > 3 {
			startStair, endStair := staircases[ns[3]], staircases[ns[4]]
			fromStart := node{id: startStair.id, step: startStep}
			fromEnd := node{id: s, step: startStep}
			toStart := node{id: s, step: endStep}
			toEnd := node{id: endStair.id, step: endStep}
			adj[fromStart.hash()] = append(adj[fromStart.hash()], fromEnd)
			adj[toStart.hash()] = append(adj[toStart.hash()], toEnd)
		}
	}

	moveSet := set.NewMapset(moves...)
	moveSet.Remove(1)
	biggestMove := moves[len(moves)-1]

	adjWithSkips := maps.Clone(adj)
	for hash, neighbors := range adj {
		n := fromHash(hash)
		adjWithSkips[hash] = append(slices.Clone(neighbors), getNextPositions(n, adj, moveSet, biggestMove)...)
	}

	startStaircase := staircases[1]
	start := node{id: 1, step: startStaircase.start}
	end := node{id: 1, step: startStaircase.end}
	numPaths := countPathsInDAG(adjWithSkips, start)
	fmt.Println(numPaths[end.hash()])
}

func countPathsInDAG(adj map[string][]node, start node) map[string]*big.Int {
	numPaths := make(map[string]*big.Int, len(adj))
	numPaths[start.hash()] = big.NewInt(1)

	for _, n := range topologicalSort(adj) {
		total := numPaths[n.hash()]
		for _, neighbor := range adj[n.hash()] {
			neighborTotal := numPaths[neighbor.hash()]
			if neighborTotal == nil {
				neighborTotal = big.NewInt(0)
			}
			numPaths[neighbor.hash()] = neighborTotal.Add(neighborTotal, total)
		}
	}

	return numPaths
}

func topologicalSort(adj map[string][]node) []node {
	indegree := make(map[string]int)
	for hash, neighbors := range adj {
		if _, ok := indegree[hash]; !ok {
			indegree[hash] = 0
		}
		for _, n := range neighbors {
			indegree[n.hash()] += 1
		}
	}

	q := queue.New[node]()
	for hash, count := range indegree {
		if count == 0 {
			q.Enqueue(fromHash(hash))
		}
	}

	order := make([]node, 0, len(adj))
	for !q.Empty() {
		n := q.Dequeue()
		order = append(order, n)
		for _, neighbor := range adj[n.hash()] {
			neighborHash := neighbor.hash()
			indegree[neighborHash] -= 1
			if indegree[neighborHash] == 0 {
				q.Enqueue(neighbor)
			}
		}
	}

	return order
}

func getNextPositions(start node, adj map[string][]node, moveSet set.Set[int], biggestMove int) []node {
	finals := make([]node, 0)
	q := queue.New[node]()
	q.Enqueue(start)
	seen := set.NewMapset(start.hash())
	for !q.Empty() {
		n := q.Dequeue()
		if moveSet.Has(n.length) && !seen.Has(n.hash()) {
			finals = append(finals, n)
			seen.Put(n.hash())
		}
		if n.length == biggestMove {
			continue
		}
		for _, neighbor := range adj[n.hash()] {
			q.Enqueue(node{id: neighbor.id, step: neighbor.step, length: n.length + 1})
		}
	}
	return finals
}

func countPathsInStaircase(end int, moves []int) *big.Int {
	dp := make([]*big.Int, end+1)
	dp[0] = big.NewInt(1)
	for i := 1; i <= end; i++ {
		// dp[i] = sum(dp[i-m]) for all moves m where i - m >= 0
		total := big.NewInt(0)
		for _, m := range moves {
			if i-m >= 0 {
				total.Add(total, dp[i-m])
			}
		}
		dp[i] = total
	}

	return dp[end]
}
