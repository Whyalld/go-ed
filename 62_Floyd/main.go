package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	inf := math.MaxInt / 2

	dist := make([][]int, n)

	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)

		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = inf
			}
		}
	}

	for edgeNumber := 0; edgeNumber < m; edgeNumber++ {
		var a, b, weight int
		fmt.Scan(&a, &b, &weight)

		a--
		b--

		if weight < dist[a][b] {
			dist[a][b] = weight
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] == inf || dist[k][j] == inf {
					continue
				}

				newDistance := dist[i][k] + dist[k][j]

				if newDistance < dist[i][j] {
					dist[i][j] = newDistance
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dist[i][j] == inf {
				fmt.Print("INF ")
			} else {
				fmt.Printf("%d ", dist[i][j])
			}
		}

		fmt.Println()
	}
}
