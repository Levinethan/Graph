package Graph

type AllowGraph interface {
	PreFlowGraph
	Discharge(interface{})
	
}

type AdjacencyMatrixAllowed struct {

	AdjacencyMatrixPreFlow
	edges Graph
}

func (g *AdjacencyMatrixAllowed) Init(fg FlowGraph,s,t interface{}) *AdjacencyMatrixAllowed {
	g.AdjacencyMatrixPreFlow.Init(fg,s,t)
	g.edges = NewGraph()
	for _ , e := range fg.AllEdges(){
		g.edges.AddEdgeBi(e)  //增加无向图
	}
	return g
}

func NewAllowGraph(fg FlowGraph,s,t interface{}) *AdjacencyMatrixAllowed  {
	return new(AdjacencyMatrixAllowed).Init(fg,s,t) //初始化
}

func (g *AdjacencyMatrixAllowed)Discharge(v interface{})  {
	iter := g.edges.IterConnectedVertices(v)  //链接呆呆其
	for g.OverFlow(v){
		if iter.Value() == nil{
			g.Relabel(v)
			iter = g.edges.IterConnectedVertices(v)  //没溢出的 加入
		}else if !g.Push(Edge{v,iter.Value()}){
			//压入可以走的边

			iter.Next()  //下一个
		}


	}

	//
}
