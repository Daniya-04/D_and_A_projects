package main

import (
	"fmt"

	"D_and_A_projects/internal/stack/graph"
)

func main() {
	// ---------- Example 1: string graph ----------
	//
	//	A ──► B ──► D
	//	│     │
	//	│     └───► E ──► F
	//	└───► C ──────────▲
	//	            F ──► A  (cycle back to A)
	//	G  (isolated vertex)
	g := &graph.Graph[string]{}

	a := g.AddVertex("A")
	b := g.AddVertex("B")
	c := g.AddVertex("C")
	d := g.AddVertex("D")
	e := g.AddVertex("E")
	f := g.AddVertex("F")
	_ = g.AddVertex("G")

	a.AddEdge(b, c)
	b.AddEdge(d, e)
	e.AddEdge(f)
	 // cycle: shows that visited-tracking prevents infinite loops

	fmt.Println("=== Adjacency list ===")
	fmt.Print(g.ToString())

	fmt.Println("\n=== DFS from A to F ===")
	g.DFS(a, f)

	fmt.Println("\n=== BFS from A to F ===")
	g.BFS(a, f)

	fmt.Println("\n=== DFS from A to D ===")
	g.DFS(a, d)

	fmt.Println("\n=== BFS from A to D ===")
	g.BFS(a, d)

	// ---------- Example 2: int graph (generics) ----------
	ig := &graph.Graph[int]{}

	n1 := ig.AddVertex(1)
	n2 := ig.AddVertex(2)
	n3 := ig.AddVertex(3)
	n4 := ig.AddVertex(4)
	n5 := ig.AddVertex(5)

	n1.AddEdge(n2, n3)
	n2.AddEdge(n4)
	n3.AddEdge(n4, n5)
	n4.AddEdge(n5)

	fmt.Println("\n=== Int graph: adjacency list ===")
	fmt.Print(ig.ToString())

	fmt.Println("\n=== Int graph: DFS from 1 to 5 ===")
	ig.DFS(n1, n5)

	fmt.Println("\n=== Int graph: BFS from 1 to 5 ===")
	ig.BFS(n1, n5)

	// ---------- Example 3: unreachable target ----------
	// G is isolated, so the search visits everything reachable from A
	// and then simply ends without printing a newline-terminated "found" path.
	fmt.Println("\n=== BFS from A to unreachable vertex ===")
	g.BFS(a, g.Vertices[len(g.Vertices)-1])
	fmt.Println()

	// ---------- Example 4: nil safety ----------
	fmt.Println("=== nil start/end (no output expected) ===")
	g.DFS(nil, f)
	g.BFS(a, nil)
	fmt.Println("done")
}