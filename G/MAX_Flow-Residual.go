package Graph

import "math"

type ResidualGraph interface {

	FlowGraph
	RCap (edge Edge)int
}

type AdjacencyMatrixResidual struct {
	AdjacencyMatrixWithFlow
}

func (g *AdjacencyMatrixResidual) Init() *AdjacencyMatrixResidual  {
	g.AdjacencyMatrixWithFlow.init()
	return g
}

func (g *AdjacencyMatrixResidual)RCap(edge Edge) int {
	return g.Cap(edge) - g.Flow(edge)
}

func NewResidualGraph(g FlowGraph)ResidualGraph  {
	residualG := new(AdjacencyMatrixResidual).Init()
	for _ , e := range g.AllEdges(){
		residualG.AddEdgeWithCap(e,g.Cap(e)) //追加边长
	}
	return residualG
}

//根据流量增加边长
func (g *AdjacencyMatrixResidual)AddEdgeWithFlow(edge Edge,flow int)  {
	g.AdjacencyMatrixWithFlow.AddEdgeWithFlow(edge,flow)

	if g.RCap(edge) == 0{
		g.AdjacencyMatrixWithFlow.DeleteEdge(edge) //删除=0的边长
	}

}

//建议路径
func AugementingPath(graph ResidualGraph,s interface{}, t interface{})(int ,[]Edge)  {
	augmentEdge := make([]Edge,0,0)
	minRC := math.MaxInt32
	handler := NewBFSVisitHandler()
	handler.EdgeHander = func(start *BFSElement, end *BFSElement) {
		if end.V == t{
			for v := end;v.P !=nil;v = v.P{
				curentEdge := Edge{v.P.V,v.V}  //获取当前边长
				augmentEdge = append(augmentEdge,curentEdge) //叠加
				if rc := graph.RCap(curentEdge);rc < minRC {
					minRC = rc
				}
			}
		}
	}
	bfsVisit(graph,s,handler)  //广度遍历  要设置参数
	return minRC,augmentEdge
}

//
func updateFlow(rg ResidualGraph,g FlowGraph,rc int , edges []Edge)  {
	for _, e := range edges{
		flow := g.Flow(e) + rc
		g.AddEdgeWithFlow(e,flow)
		rg.AddEdgeWithFlow(e,flow)
		re := Edge{e.End,e.Start}
		rg.AddEdgeWithFlow(re, 0-flow)
	}
}

func edmondesKarp(g FlowGraph , s interface{} , t interface{})  {
	rG := NewResidualGraph(g)  //新建一个剩余图
	for rc , edges := AugementingPath(rG,s,t);len(edges)>0;rc,edges=AugementingPath(rG,s,t){
		updateFlow(rG,g,rc,edges) //刷新图
	}
}


