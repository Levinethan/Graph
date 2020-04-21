package Graph

type VertexDegree struct {
	InputDegree int
	OutputDegree int
}

func NewEulerVertex(vertex interface{}) *VertexDegree  {
	return new(VertexDegree).init(vertex)
}

func (e *VertexDegree)init (vertex interface{})*VertexDegree  {
	e.InputDegree= 0
	e.OutputDegree = 0 //出度 和入度  都为零
	return e
}

//判断 出度和入度  是否相等
func checkDegree(degree *VertexDegree,e *DFS_Element)bool  {
	if degree.InputDegree == 0 || degree.OutputDegree ==0{
		return false
	}

	if e.Iter.Len() ==1 && e.Iter.Value()==e.V{
		return false
	}

	return degree.InputDegree == degree.OutputDegree
}

func checkVertexAndEdgeCo(vetexCount , edgeCount int)bool  {
	//假如有5个节点  判断是否有6个边长

	return edgeCount == vetexCount +1
}


//传入  图 参数  返回  欧拉回路
func EluerCircuit(graph Graph)[]Edge  {
	degrees := make(map[interface{}]*VertexDegree)  //degrees 返回值是图VertexDegree

	path := make([]Edge,0,0)

	vetexCount , edgeCount :=0,0 //初始化节点和边长数量

	nonTreeEdgeHander := func(start, end *DFS_Element) {
		//函数指针
		if start.P!=nil&&start.P!=end{
			degrees[start.V].OutputDegree++

			//统计每个节点

			degrees[end.V].InputDegree++

			//出度++  入读++
			edgeCount ++

			//边长数量++
			path = append(path,Edge{start.V,end.V})
			//记录路径 追加进path

		}

	}

	handler := NewDESVisitHandler() //深度遍历访问器

	handler.BeforeDFSHander = func(v *DFS_Element) {
		degrees[v.V]= NewEulerVertex(v.V) //记录出度入度
	}

	handler.TreeEdgeHander  = func(start *DFS_Element, end *DFS_Element) {
		edgeCount ++
		degrees[start.V].OutputDegree++
		degrees[end.V].InputDegree = 1
		path = append(path,Edge{start.V,end.V})


	}
	handler.BackEdgeHander = nonTreeEdgeHander

	for _ , v := range  graph.AllVertices(){
		//遍历所有节点
		vetexCount ++
		if !handler.Element.exist(v){
			dfsVisit(graph,v,handler)  //深度遍历
		}
	}

	if !checkVertexAndEdgeCo(vetexCount,edgeCount){
		return nil
	}

	for _ , e:= range  path{
		if !checkDegree(degrees[e.Start],handler.Element.get(e.Start).(*DFS_Element)) ||!checkDegree(degrees[e.End],handler.Element.get(e.End).(*DFS_Element)) {
			//判断开头结尾的出度 入度
			return nil
		}
	}
	return path  //返回路径
}