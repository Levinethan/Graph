package Graph

import (
	"container/list"

)

//求关键点 ，切割成几个子图
func vertexBCC(graph Graph) (cuts Graph,comps []Graph)  {
	cuts = NewGraph()  //cut切割  新的图去掉这个点与相关边长


	comps = make([]Graph,0,0) //图的数组，子图

	lows := make(map[interface{}]int)  //存储map


	child := make(map[interface{}]int) //存储子节点

	edgeStack := list.New()  //模拟构造一个栈 不断循环出栈进栈

	handler := NewDESVisitHandler()  //DFS访问器 ，内部防止指针


	handler.BeforeDFSHander = func(v *DFS_Element) {
		lows[v.V] = v.Dist //根据点存储距离
	}
	handler.TreeEdgeHander = func(start *DFS_Element, end *DFS_Element) {
		child[start.V]++
		edgeStack.PushBack(Edge{start.V,end.V})

	}

	handler.BackEdgeHander = func(start *DFS_Element, end *DFS_Element) {
		if start.Dist < lows[start.V] && start.P!=end{
			//满足上诉条件 压入栈
			edgeStack.PushBack(Edge{start.V,end.V})   //Push data in stack
			lows[start.V] = end.Dist  //save the distence
		}
	}

	handler.AfterDFSHander = func(v *DFS_Element) {
		//重点
		p := v.P  //mark down the point
		if p == nil{
			return
		}
		if lows[v.V] < lows[p.V]{

			//save the shortest point
			lows[p.V] = lows[v.V]
		}
		if lows[v.V] > lows[p.V]{
			if !(p.Dist==1 && child[p.V]<2){
				cuts.AddVertex(p.V) //find the important point
			}
			comps = append(comps,NewGraph())  //追加的集合   extra adding set
			curEdge := Edge{p.V,v.V}
			comps[len(comps)-1].AddEdgeBi(curEdge) //add double connected
			for e:=edgeStack.Back().Value;e!=curEdge;e=edgeStack.Back(){
				edgeStack.Remove(edgeStack.Back())
				comps[len(comps)-1].AddEdgeBi(curEdge)
			}
			edgeStack.Remove(edgeStack.Back())
			  //pop data
		}

	}

	for _ , v := range graph.AllVertices(){
		if !handler.Element.exist(v){
			dfsVisit(graph,v,handler)  //for range graph
		}
	}
	return

}


func edgeBCC(graph Graph) (bridges Graph,comps []Graph)  {
	bridges = NewGraph()
	comps = make([]Graph,0,0)
	lows := make(map[interface{}]int)
	edgeStack := list.New()
	handler := NewDESVisitHandler()  //DFS访问器 ，内部防止指针
	handler.BeforeDFSHander = func(v *DFS_Element) {
		lows[v.V] = v.Dist //根据点存储距离
	}
	handler.TreeEdgeHander = func(start *DFS_Element, end *DFS_Element) {
		edgeStack.PushBack(Edge{start.V,end.V})

	}
	handler.BackEdgeHander = func(start *DFS_Element, end *DFS_Element) {
		if start.Dist < lows[start.V] && start.P!=end{
			//满足上诉条件 压入栈
			edgeStack.PushBack(Edge{start.V,end.V})   //Push data in stack
			lows[start.V] = end.Dist  //save the distence
		}
	}
	handler.AfterDFSHander = func(v *DFS_Element) {
		//重点
		p := v.P  //mark down the point
		if p == nil{
			return
		}
		if lows[v.V] < lows[p.V]{

			//save the shortest point
			lows[p.V] = lows[v.V]
		}
		if lows[v.V] >= lows[p.V]{
			//find the right path
			comp := NewGraph() //new graph
			curEdge := Edge{p.V,v.V}  //curent edge
			if lows[v.V] > p.Dist{
				bridges.AddEdgeBi(curEdge) //找到关键边长

			}else {
				comp.AddEdgeBi(curEdge)
			}
			for e:=edgeStack.Back().Value;e!=curEdge;e=edgeStack.Back(){
				edgeStack.Remove(edgeStack.Back())
				comp.AddEdgeBi(curEdge)
			}
			edgeStack.Remove(edgeStack.Back())

			if len(comp.AllVertices()) >0{
				comps=append(comps,comp) //追加节点
			}
		}
	}


	for _ , v := range graph.AllVertices(){
		if !handler.Element.exist(v){
			dfsVisit(graph,v,handler)  //for range graph
		}
	}
	return

}