package main

import (
	"fmt"
	"strings"

	lib "github.com/thebenkogan/Codyssi-2025"
	"github.com/zyedidia/generic/queue"
	"github.com/zyedidia/generic/set"
)

const (
	maxX = 10
	maxY = 15
	maxZ = 60
)

type debris struct {
	x, y, z, a     int
	vx, vy, vz, va int
}

func (d debris) posAfterDuration(t int) (int, int, int, int) {
	x := d.x + d.vx*t
	y := d.y + d.vy*t
	z := d.z + d.vz*t
	a := d.a + d.va*t
	return mod(x, maxX), mod(y, maxY), mod(z, maxZ), correctA(a)
}

type node struct {
	x, y, z, a, t, collisions int
}

func (n node) hash() string {
	return fmt.Sprintf("%d,%d,%d,%d,%d,%d", n.x, n.y, n.z, n.a, n.t, n.collisions)
}

func main() {
	input := lib.GetInput()

	ds := make([]debris, 0)
	for _, line := range strings.Split(input, "\n") {
		ns := lib.ParseNums(line)
		c1, c2, c3, c4, d, r, vx, vy, vz, va := ns[1], ns[2], ns[3], ns[4], ns[5], ns[6], ns[7], ns[8], ns[9], ns[10]
		for x := range maxX {
			for y := range maxY {
				for z := range maxZ {
					for a := -1; a < 2; a++ {
						sum := x*c1 + y*c2 + z*c3 + a*c4
						if ((sum%d)+d)%d == r {
							ds = append(ds, debris{x, y, z, a, vx, vy, vz, va})
						}
					}
				}
			}
		}
	}

	fmt.Println(len(ds))
	fmt.Println(lowestDuration(ds, 1))
	fmt.Println(lowestDuration(ds, 4))
}

func lowestDuration(ds []debris, maxCollisions int) int {
	q := queue.New[node]()
	q.Enqueue(node{})
	seen := set.NewMapset(node{}.hash())
	for !q.Empty() {
		n := q.Dequeue()
		if n.x == maxX-1 && n.y == maxY-1 && n.z == maxZ-1 && n.a == 0 {
			return n.t
		}

		for _, nd := range neighbors {
			dx, dy, dz := nd[0], nd[1], nd[2]
			px, py, pz := n.x+dx, n.y+dy, n.z+dz
			if !inBounds(px, py, pz) {
				continue
			}
			nxt := node{px, py, pz, n.a, n.t + 1, n.collisions}
			if checkCollision(ds, nxt) {
				if n.collisions == maxCollisions-1 {
					continue
				}
				nxt.collisions += 1
			}
			if seen.Has(nxt.hash()) {
				continue
			}
			seen.Put(nxt.hash())
			q.Enqueue(nxt)
		}
	}

	panic("no path")
}

var debrisCache = make(map[int]set.Set[string])

func checkCollision(ds []debris, n node) bool {
	if n.x == 0 && n.y == 0 && n.z == 0 {
		return false
	}
	cached, ok := debrisCache[n.t]
	if !ok {
		cached = set.NewMapset[string]()
		for _, d := range ds {
			px, py, pz, pa := d.posAfterDuration(n.t)
			p := node{px, py, pz, pa, n.t, 0}
			cached.Put(p.hash())
		}
		debrisCache[n.t] = cached
	}
	n.collisions = 0
	return cached.Has(n.hash())
}

var neighbors = [][]int{
	{1, 0, 0},
	{0, 1, 0},
	{0, 0, 1},
	{0, 0, 0},
	{0, -1, 0},
	{0, 0, -1},
	{-1, 0, 0},
}

func inBounds(x, y, z int) bool {
	return x >= 0 && x < maxX && y >= 0 && y < maxY && z >= 0 && z < maxZ
}

func mod(a, b int) int {
	return ((a % b) + b) % b
}

func correctA(a int) int {
	return mod((a+1), 3) - 1
}
