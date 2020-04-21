package main

import (
	"./G"
	"fmt"
)
func Tee()  {
	g := Graph.NewGraph()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddEdgeBi(Graph.Edge{3,4})
	g.AddEdgeBi(Graph.Edge{1,4})
	g.AddEdgeBi(Graph.Edge{2,4})

	fmt.Println(g.AllEdges())
	fmt.Println(g.AllVertices())
	connect := make([][]interface{},len(g.AllVertices()),len(g.AllVertices()))
	for i := range g.AllVertices(){
		connect[i] = g.AllConnectedVertices(g.AllVertices()[i])
	}
	fmt.Println(connect)

}

func T_BFS()  {
	g := Graph.NewGraph()
	g.AddVertex("r")
	g.AddVertex("x")
	g.AddVertex("t")
	g.AddVertex("u")
	g.AddVertex("v")
	g.AddVertex("w")
	g.AddVertex("x")
	g.AddVertex("y")
	g.AddVertex("z")
	g.AddEdgeBi(Graph.Edge{"r","s"})
	g.AddEdgeBi(Graph.Edge{"s","w"})
	g.AddEdgeBi(Graph.Edge{"r","v"})
	g.AddEdgeBi(Graph.Edge{"w","t"})
	g.AddEdgeBi(Graph.Edge{"w","x"})
	g.AddEdgeBi(Graph.Edge{"t","x"})
	g.AddEdgeBi(Graph.Edge{"t","u"})
	g.AddEdgeBi(Graph.Edge{"x","u"})

	g.AddEdgeBi(Graph.Edge{"x","y"})
	g.AddEdgeBi(Graph.Edge{"u","y"})


	mynew_Graph := Graph.DFS(g,nil)
	fmt.Println(mynew_Graph.AllVertices())
	fmt.Println(mynew_Graph.AllEdges())
	for _ , v := range mynew_Graph.AllVertices(){
		fmt.Println(v)
	}
}

func TestStronglyConnectedComp()  {
	g :=Graph.NewGraph()
	g.AddVertex("a")
	g.AddVertex("b")
	g.AddVertex("c")
	g.AddVertex("d")
	g.AddVertex("e")
	g.AddVertex("f")
	g.AddVertex("g")
	g.AddVertex("h")


	g.AddEdge(Graph.Edge{"a","b"})
	g.AddEdge(Graph.Edge{"b","e"})
	g.AddEdge(Graph.Edge{"e","a"})
	g.AddEdge(Graph.Edge{"e","f"})
	g.AddEdge(Graph.Edge{"b","f"})
	g.AddEdge(Graph.Edge{"b","c"})
	g.AddEdge(Graph.Edge{"c","d"})
	g.AddEdge(Graph.Edge{"d","c"})
	g.AddEdge(Graph.Edge{"c","g"})
	g.AddEdge(Graph.Edge{"f","g"})
	g.AddEdge(Graph.Edge{"g","f"})
	g.AddEdge(Graph.Edge{"g","h"})
	g.AddEdge(Graph.Edge{"d","h"})
	g.AddEdge(Graph.Edge{"h","h"})
	sc := Graph.StronglyConnect(g)  //计算的图
	fmt.Println(sc.AllVertices())
	fmt.Println(sc.AllEdges())

}

func Euler()  {
	g := Graph.NewGraph()
	g.AddEdge(Graph.Edge{2,3})
	g.AddEdge(Graph.Edge{2,5})
	g.AddEdge(Graph.Edge{3,4})
	g.AddEdge(Graph.Edge{1,2})
	g.AddEdge(Graph.Edge{4,2})
	g.AddEdge(Graph.Edge{5,1})
	//g.AddEdge(Graph.Edge{6,1}) 增加6 1 节点 欧拉回路不成立 返回空集
	//g.AddVertex(6)
	//g.AddEdgeBi(Graph.Edge{6,7})

	path := Graph.EluerCircuit(g)
	fmt.Println(path)
	fmt.Println(g.AllVertices())
	fmt.Println(g.AllEdges())

}
func main()  {
	Euler()
}