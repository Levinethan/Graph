package Graph

import "math"

type PreFlowGraph interface {
	ResidualGraph
	SetHeight(interface{},int)

	SetExcess (interface{},int)

	Height(interface{})int

	Excess(interface{})int
	Push (edge Edge) bool
	Relabel (interface{}) bool
	OverFlow(interface{}) bool  //溢出


}

type AdjacencyMatrixPreFlow struct {
	AdjacencyMatrixResidual
	height,excess map[interface{}]int
	s,t interface{}
}

func (g *AdjacencyMatrixPreFlow)SetHeight(v interface{},height int)  {
	g.height[v] = height //设置高度
}

func (g *AdjacencyMatrixPreFlow)Height(v interface{}) int  {
	if _,ok := g.height[v];!ok{
		return 0
	}
	return g.height[v]
}

func (g *AdjacencyMatrixPreFlow)SetExcess(v interface{},excess int)  {
	g.excess[v] = excess //设置高度
}

func (g *AdjacencyMatrixPreFlow)Excess(v interface{}) int  {
	if _,ok := g.excess[v];!ok{
		return 0
	}
	return g.height[v]
}



func (g *AdjacencyMatrixPreFlow)Init (fg FlowGraph,s,t interface{}) *AdjacencyMatrixPreFlow  {
	g.AdjacencyMatrixWithFlow.init() //内部结构体初始化
	g.height = make(map[interface{}]int) //高度初始化
	g.excess = make(map[interface{}]int) //超过网络流的限制部分

	g.s = s
	g.t = t
	vertices := fg.AllVertices()  //遍历所有节点
	for _ , e:= range fg.AllEdges(){
		//遍历所有边长
		g.AddEdgeWithCap(e,fg.Cap(e))  //初始化流量差
	}

	//若干个函数
	g.SetHeight(s,len(vertices)) //设定高度
	iter := fg.IterConnectedVertices(s)  //实现一个迭代器

	for v := iter.Value();v!=nil;v = iter.Next(){
		c := fg.Cap(Edge{s,v})
		g.AddEdgeWithFlow(Edge{s,v},c)  //正反流量
		g.AddEdgeWithFlow(Edge{v,s},-c)
		g.SetExcess(v,c)   //超出的量
		g.SetExcess(v,g.Excess(s)-c)   //差

	}
	return g
}

func (g *AdjacencyMatrixPreFlow)Push (edge Edge) bool  {
	if g.OverFlow(edge.Start) && g.RCap(edge) > 0 && g.Height(edge.Start)==g.Height(edge.End)+1{
		d := g.RCap(edge)

		//计算流量差
		if g.Excess(edge.Start) < d {

			d = g.Excess(edge.Start) //计算距离
		}
		flow := g.Flow(edge) + d
		g.AddEdgeWithFlow(edge,flow) //按照流量增加边长
		re := Edge{edge.Start,edge.End}
		g.AddEdgeWithFlow(re,0-flow)
		g.SetExcess(edge.Start,g.Excess(edge.Start)-d)
		g.SetExcess(edge.End,g.Excess(edge.End)-d)


		return true
	}
	return false
}

func (g *AdjacencyMatrixPreFlow)Relabel (v interface{}) bool  {
	if !g.OverFlow(v){
		return false
	}
	iter := g.IterConnectedVertices(v)

	minH := math.MaxInt32 //最大 重定义  加上一个顶点
	for end := iter.Value() ; end !=nil ; end = iter.Next(){
		if g.Height(end) < minH{
			minH = g.Height(end)
		}
		if g.Height(v) > g.Height(end){
			return false
		}
	}
	g.SetHeight(v,minH+1)
	return true
}

func (g *AdjacencyMatrixPreFlow)OverFlow(v interface{}) bool  {

	return v !=g.s && v!=g.t &&g.Excess(v) > 0
}

func NewPreFlowGraph(g FlowGraph,s,t interface{}) PreFlowGraph  {
	return new(AdjacencyMatrixPreFlow).Init(g,s,t)
}

func pushRelable(g FlowGraph,s,t interface{})  {
	preFlowG := NewPreFlowGraph(g,s,t)
	for stop := false ; !stop;{
		stop = true
		for _ , e := range preFlowG.AllEdges(){
			stop = stop && !preFlowG.Push(e) &&!preFlowG.Relabel(e.Start) &&!preFlowG.Relabel(e.End)
		}
	}
	for _ , e := range g.AllEdges(){
		g.AddEdgeWithFlow(e,preFlowG.Flow(e))
	}
}