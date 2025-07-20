package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/samber/lo"
	lib "github.com/thebenkogan/Codyssi-2025"
	"github.com/zyedidia/generic/queue"
)

type debt struct {
	to   *account
	owed int
}

type account struct {
	total int
	debts *queue.Queue[*debt]
}

func NewAccount(amnt int) *account {
	return &account{total: amnt, debts: queue.New[*debt]()}
}

func (a *account) Pay(o *account, amnt int) {
	owed := amnt - a.total
	o.Receive(min(amnt, a.total))
	a.total = max(0, a.total-amnt)
	if owed > 0 {
		a.debts.Enqueue(&debt{o, owed})
	}
}

func (a *account) Receive(amnt int) {
	a.total += amnt
	for !a.debts.Empty() && a.total > 0 {
		d := a.debts.Peek()
		if a.total >= d.owed {
			a.debts.Dequeue()
			d.to.Receive(d.owed)
			a.total -= d.owed
		} else {
			d.owed -= a.total
			d.to.Receive(a.total)
			a.total = 0
		}
	}
}

func main() {
	input := lib.GetInput()
	sections := strings.Split(input, "\n\n")

	p1Accounts := make(map[string]*account)
	p2Accounts := make(map[string]*account)
	p3Accounts := make(map[string]*account)
	for _, line := range strings.Split(sections[0], "\n") {
		p := strings.Split(line, " ")[0]
		n := lib.ParseNums(line)[0]
		p1Accounts[p] = NewAccount(n)
		p2Accounts[p] = NewAccount(n)
		p3Accounts[p] = NewAccount(n)
	}

	for _, line := range strings.Split(sections[1], "\n") {
		n := lib.ParseNums(line)[0]
		from := strings.Split(line, " ")[1]
		to := strings.Split(line, " ")[3]

		p1Accounts[from].total -= n
		p1Accounts[to].total += n

		p2Accounts[to].total += min(n, p2Accounts[from].total)
		p2Accounts[from].total = max(0, p2Accounts[from].total-n)

		p3Accounts[from].Pay(p3Accounts[to], n)
	}

	fmt.Println(top3(p1Accounts))
	fmt.Println(top3(p2Accounts))
	fmt.Println(top3(p3Accounts))
}

func top3(accountMap map[string]*account) int {
	accounts := lo.Values(accountMap)
	slices.SortFunc(accounts, func(a, b *account) int {
		return b.total - a.total
	})
	return lo.SumBy(accounts[:3], func(a *account) int {
		return a.total
	})
}
