package main

import (
	"fmt"
	"log"
)

type Box []string

var boxes []*Box

func (b Box) Size() int {
	return len(b)
}

func (b Box) Top() string {
	if b.Size() <= 0 {
		return ""
	}

	return b[len(b)-1]
}

func (b *Box) Push(w string) {
	*b = append(*b, w)
}

func (b *Box) Pop() {
	if len(*b) <= 0 {
		return
	}
	*b = (*b)[:len(*b)-1]
}

func rearrange(op Operand) {
	from := boxes[op.From-1]
	to := boxes[op.To-1]
	log.Printf("%+v: %+v", op, boxes)
	for i := 0; i < op.Count; i++ {
		to.Push(from.Top())
		from.Pop()
	}
}

func solve(strs []string) string {
	var res string
	for _, str := range strs {
		operand := extractOperand(str)
		rearrange(operand)
	}

	for _, box := range boxes {
		res += box.Top()
	}

	// box length의 마지막
	return res
}

type Operand struct {
	Count int
	From  int
	To    int
}

func extractOperand(str string) Operand {
	operand := Operand{}
	fmt.Sscanf(str, "move %d from %d to %d", &operand.Count, &operand.From, &operand.To)
	return operand
}
