package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
	lib "github.com/thebenkogan/Codyssi-2025"
	"github.com/zyedidia/generic/queue"
)

type node struct {
	left  *node
	right *node
	id    int
	code  string
}

type bst struct {
	root *node
}

// Insert adds a new node to the tree with the given id and code, returning the sequence of codes
// traversed during the insertion excluding the new code.
func (b *bst) Insert(id int, code string) []string {
	newNode := &node{id: id, code: code}
	if b.root == nil {
		b.root = newNode
		return nil
	}
	parent, sequence := b.Find(id)
	if parent.id == id {
		panic("node already exists with id: " + strconv.Itoa(id))
	}
	if id > parent.id {
		parent.right = newNode
	} else {
		parent.left = newNode
	}
	return sequence
}

type levelOrderTraversalNode struct {
	n     *node
	layer int
}

// LevelOrder returns a level order traversal of the tree, where arr[i] is the IDs
// from left to right on layer i + 1.
func (b *bst) LevelOrder() [][]int {
	if b.root == nil {
		return nil
	}
	order := make([][]int, 0)
	q := queue.New[*levelOrderTraversalNode]()
	q.Enqueue(&levelOrderTraversalNode{n: b.root})
	for !q.Empty() {
		ln := q.Dequeue()
		if ln.layer >= len(order) {
			order = append(order, make([]int, 0))
		}
		order[ln.layer] = append(order[ln.layer], ln.n.id)
		if ln.n.left != nil {
			q.Enqueue(&levelOrderTraversalNode{n: ln.n.left, layer: ln.layer + 1})
		}
		if ln.n.right != nil {
			q.Enqueue(&levelOrderTraversalNode{n: ln.n.right, layer: ln.layer + 1})
		}
	}
	return order
}

// Find returns the node with the given id, along with the sequence of codes traversed excluding
// the found node's code. If no node in the tree has the given ID, it returns the parent of where
// the node would be.
func (b *bst) Find(id int) (*node, []string) {
	sequence := make([]string, 0)
	curr := b.root
	for curr != nil {
		if curr.id == id {
			return curr, sequence
		}
		sequence = append(sequence, curr.code)
		if id > curr.id && curr.right != nil {
			curr = curr.right
		} else if id < curr.id && curr.left != nil {
			curr = curr.left
		} else {
			return curr, sequence
		}
	}
	return nil, sequence
}

func main() {
	input := lib.GetInput()
	sections := strings.Split(input, "\n\n")

	tree := new(bst)
	for _, line := range strings.Split(sections[0], "\n") {
		id, code := parseLine(line)
		tree.Insert(id, code)
	}

	order := tree.LevelOrder()
	maxLayer := lo.Max(lo.Map(order, func(l []int, _ int) int { return lo.Sum(l) }))
	fmt.Println(maxLayer * len(order))

	sequence := tree.Insert(500000, "code")
	fmt.Println(strings.Join(sequence, "-"))

	lines := strings.Split(sections[1], "\n")
	id1, _ := parseLine(lines[0])
	id2, _ := parseLine(lines[1])
	_, path1 := tree.Find(id1)
	_, path2 := tree.Find(id2)

	var lca string
	for _, n := range lo.Zip2(path1, path2) {
		if n.A != n.B {
			break
		}
		lca = n.A
	}
	fmt.Println(lca)
}

func parseLine(line string) (int, string) {
	return lib.ParseNums(line)[0], strings.Split(line, " | ")[0]
}
