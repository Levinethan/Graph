package Graph

type WeightGraph interface {
	Graph
	Weight(edge Edge)int //返回边长的权重
	AddWeightEdge(edge Edge,weight int)//增加有向图的权重
	AddWeightEdgeBi(edge Edge,weight int)//增加有无向图的权重
	TotalWeight()int //统计权重的总量
}
//t图的数据结构
type  AdjacencyMatrixWithWeight struct {
	AdjacencyMatrix
	weights map[Edge]int
	totalweight  int
}
func NewAdjacencyMatrixWithWeight()WeightGraph{
	return new(AdjacencyMatrixWithWeight).Init()
}

func ( g*AdjacencyMatrixWithWeight)Init()WeightGraph{
	g.AdjacencyMatrix.init()//初始化图
	g.weights=make(map[Edge]int) //权重
	g.totalweight=0 //总重量
	return g
}
//返回边长的权重
func ( g*AdjacencyMatrixWithWeight) Weight(edge Edge)int{
	if  value,ok:=g.weights[edge];ok{
		return value
	}
	return -1
}
//增加有向图的权重
func ( g*AdjacencyMatrixWithWeight)AddWeightEdge(edge Edge,weight int){
	g.AdjacencyMatrix.AddEdge(edge)//增加边长
	if _,ok:=g.weights[edge];ok{
		g.totalweight=g.totalweight-g.weights[edge]+weight
	}else{
		g.totalweight=	g.totalweight+weight
	}
	g.weights[edge]=weight
}
//增加有无向图的权重
func ( g*AdjacencyMatrixWithWeight)AddWeightEdgeBi(edge Edge,weight int){
	g.AddWeightEdge(edge,weight)
	g.AddWeightEdge(Edge{edge.End,edge.Start},weight)
}
//统计权重的总量
func ( g*AdjacencyMatrixWithWeight)TotalWeight()int {
	return  g.totalweight
}

//继承的覆盖
func ( g*AdjacencyMatrixWithWeight)DeleteEdge(edge Edge){
	g.AdjacencyMatrix.DeleteEdge(edge)
	if w,ok:=g.weights[edge];ok{
		g.totalweight-=w
		delete(g.weights,edge) //删除权重
	}
}
//删除
func ( g*AdjacencyMatrixWithWeight)DeleteEdgeBi(edge Edge){
	g.DeleteEdge(edge)
	g.DeleteEdge(Edge{edge.End,edge.Start})
}

































