//+build disable

package main

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}

	item, items := q.items[0], q.items[1:]
	q.items = items
	return item
}

func main() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)

	println(q.Dequeue())
	println(q.Dequeue())
	println(q.Dequeue())
	println(q.Dequeue())
}
