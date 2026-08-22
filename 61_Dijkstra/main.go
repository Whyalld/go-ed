package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
)

type Edge struct {
	to     int
	weight int
}

type Item struct {
	dist   int
	vertex int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func(pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Item))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)

	item := old[n-1]
	*pq = old[:n-1]

	return item
}

func main() {
	inf := math.MaxInt
	var v, e, a, b, c int

	_, err := fmt.Scan(&v, &e)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ошибка: %v\n", err)
		os.Exit(1)
	}
	graph := make([][]Edge, v)
	for ; e > 0; e-- {
		fmt.Scan(&a, &b, &c)
		a--
		b--
		graph[a] = append(graph[a], Edge{
			to:     b,
			weight: c,
		})
	}

	dist := make([]int, v)
	for i := range dist {
		dist[i] = inf
	}
	dist[0] = 0

	q := &PriorityQueue{}
	heap.Init(q)

	heap.Push(q, Item{
		dist:   0,
		vertex: 0,
	})

	for q.Len() > 0 {
		item := heap.Pop(q).(Item)
		length := item.dist
		vertex := item.vertex

		// В очереди может находится уже старое, неактуальное расстояние.
		if length > dist[vertex] {
			continue
		}

		for _, edge := range graph[vertex] {
			newDist := dist[vertex] + edge.weight

			if newDist < dist[edge.to] {
				dist[edge.to] = newDist

				heap.Push(q, Item{
					dist: newDist,
					vertex: edge.to,
				})
			}
		}
	}

	for _, distance := range dist {
		fmt.Printf("%d ", distance)
	}

	fmt.Println()

}
