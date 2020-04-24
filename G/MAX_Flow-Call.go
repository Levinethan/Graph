package Graph

import "container/list"

func relableToFront(g FlowGraph,s,t interface{})  {
	allowedGraph := NewAllowGraph(g,s,t)  //创建一个允许运行的图
	l := list.New()  //创建列表
	for _ , v := range g.AllVertices() {
		if v != s && v !=t {
			l.PushBack(v)  //保存点
		}
	}


	for e := l.Front();e!=nil;{
		oldH := allowedGraph.Height(e.Value)  //获取旧的高度
		allowedGraph.Discharge(e.Value)  //计算一些允许运行
		if allowedGraph.Height(e.Value) > oldH{
			l.MoveToFront(e)
		}
		e = e.Next()
	}

	for _ , e := range g.AllEdges(){
		g.AddEdgeWithFlow(e,allowedGraph.Flow(e))
	}
}

func bipGraphMaxMatch(g Graph,l []interface{},flowAlg func(g FlowGraph,s,t interface{}))Graph  {
	fG := NewAdjacencyMatrixWithFlow()  //初始化一个流量图

	s := struct {
		start string
	}{"s"}

	t := struct {
		end string
	}{"t"}

	for _ , vl := range l{
		fG.AddEdgeWithCap(Edge{s,vl},1)
		iter := g.IterConnectedVertices(vl)
		for rv := iter.Value();rv !=nil;rv= iter.Next(){
			fG.AddEdgeWithCap(Edge{vl,rv},1)
			fG.AddEdgeWithCap(Edge{rv,t},1)
		}
	}
	
	flowAlg(fG,s,t)  //传递函数指针 调用
	matchG := NewGraph()  //构造一个图
	for _ , e:= range fG.AllEdges(){
		if fG.Flow(e) > 0 && e.Start != s && e.End !=t{
			matchG.AddEdgeBi(e)
		}
	}
	
	return matchG
}
