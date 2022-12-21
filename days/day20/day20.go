package day20

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	val  int
	next *Node
	prev *Node
}

func (node *Node) String() string {
	return fmt.Sprint(node.val)
}

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func (l LinkedList) String() string {
	n := l.head
	str := fmt.Sprint(n)
	i := 0

	for n.next != nil && i < l.length-1 {
		n = n.next
		str += fmt.Sprintf(" -> %v", n)
		i++
	}

	return str
}

func (l LinkedList) Len() int {
	return l.length
}

func (l *LinkedList) Push(v int) {
	n := Node{v, nil, nil}
	if l.tail != nil {
		n.prev = l.tail
		l.tail.next = &n
	}

	l.tail = &n

	if l.head == nil {
		l.head = &n
	}

	l.length++
}

func (node *Node) Right() int {
	if node.next == nil {
		return -1
	}

	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev

	node.prev = next
	node.next = next.next

	next.next.prev = node
	next.next = node

	return 1
}

func (node *Node) Left() int {
	if node.prev == nil {
		return -1
	}

	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev

	node.next = prev
	node.prev = prev.prev

	prev.prev.next = node
	prev.prev = node

	return 1
}

func (l *LinkedList) Shift(node *Node, n int) {
	if n > 0 {
		for i := 0; i < n; i++ {
			node.Right()

			if l.head == node {
				l.head = node.prev
			}

			if l.tail == node {
				l.tail = node.prev
			}
		}
	} else {
		for i := 0; i > n; i-- {
			if l.head == node {
				l.head = node.next
			}

			node.Left()

			if l.tail == node {
				l.tail = node.next
			}
		}
	}
}

func parseInput(input string) LinkedList {
	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)

	nums := LinkedList{}

	for scanner.Scan() {
		line := scanner.Text()
		val, _ := strconv.Atoi(line)
		nums.Push(val)
	}

	nums.head.prev = nums.tail
	nums.tail.next = nums.head

	return nums
}

func Part1(input string) int {
	nums := parseInput(input)
	vals := make([]*Node, nums.length)

	// fmt.Println(nums.head, nums.tail)

	n := nums.head
	zero := n

	for i := 0; i < len(vals); i++ {
		vals[i] = n

		if n.val == 0 {
			zero = n
		}

		n = n.next
	}

	// fmt.Println(nums)
	for _, v := range vals {
		nums.Shift(v, v.val)
	}
	// fmt.Println(nums)

	n1 := 1000 % len(vals)
	n2 := 2000 % len(vals)
	n3 := 3000 % len(vals)
	// fmt.Println(n1, n2, n3)

	t1, t2, t3 := 0, 0, 0

	z := zero
	for i := 1; i <= len(vals); i++ {
		z = z.next

		if i == n1 {
			t1 = z.val
		}
		if i == n2 {
			t2 = z.val
		}
		if i == n3 {
			t3 = z.val
		}
	}
	// fmt.Println(t1, t2, t3)

	return t1 + t2 + t3
}

func Part2(input string) int {
	nums := parseInput(input)
	vals := make([]*Node, nums.length)

	// fmt.Println(nums.head, nums.tail)

	n := nums.head
	zero := n

	multiplier := 811589153
	// shifter := multiplier % nums.length

	for i := 0; i < len(vals); i++ {
		n.val = n.val * multiplier
		vals[i] = n

		if n.val == 0 {
			zero = n
		}

		n = n.next
	}

	for i := 0; i < 10; i++ {
		for _, v := range vals {
			nums.Shift(v, v.val%(len(vals)-1))
		}
		// fmt.Println("after", i, nums)
	}

	n1 := 1000 % len(vals)
	n2 := 2000 % len(vals)
	n3 := 3000 % len(vals)
	// fmt.Println(n1, n2, n3)

	t1, t2, t3 := 0, 0, 0
	// fmt.Println(nums)

	z := zero
	for i := 1; i <= len(vals); i++ {
		z = z.next

		if i == n1 {
			t1 = z.val
		}
		if i == n2 {
			t2 = z.val
		}
		if i == n3 {
			t3 = z.val
		}
	}
	// fmt.Println(t1, t2, t3)

	return t1 + t2 + t3
}

func Run(input string) {
	fmt.Println("Day 20 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
