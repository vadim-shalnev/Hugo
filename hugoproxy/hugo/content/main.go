package main

import (
	"fmt"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"math/rand"
	"time"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			nodesCount := rand.Intn(26) + 5 // От 5 до 30 элементов

			g := simple.NewUndirectedGraph()
			nodes := make([]graph.Node, nodesCount)

			// Создание узлов графа
			for i := 0; i < nodesCount; i++ {
				nodes[i] = g.NewNode()
				g.AddNode(nodes[i])
			}

			// Создание ребер между узлами
			for i := 0; i < nodesCount-1; i++ {
				g.SetEdge(g.NewEdge(nodes[i], nodes[i+1]))
			}

			// Вывод графа
			fmt.Println("Граф:")
			for _, n := range nodes {
				edges := g.From(n.ID())
				for edges.Next() {
					fmt.Printf("%d -> %d\n", n.ID(), edges.Node().ID())
				}
			}
			fmt.Println()
		}
	}
}
