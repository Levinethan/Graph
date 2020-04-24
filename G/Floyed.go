package Graph

import "math"

func FloydWarShall(graph WeightGraph,
	init func(WeightGraph)([] [] []int,[]interface{}),
		handler func(*[][][]int,int,int,int),
		rebuild func([]interface{},[][]int))  {

			array , vertices := init(graph)
			for k := range  array[:len(array)-1]{
				for i:=range array[k]{
					for j := range array[k][i]{
						handler(&array,k,i,j)
					}
				}
			}
	rebuild(vertices,array[len(array)-1])
}

func DisFloyedInit(graph WeightGraph) ([][][]int,[]interface{})  {
	//初始化权重图距离
	vertices := graph.AllVertices()
	array := make([][][]int,len(vertices)-1,len(vertices)-1) //分配内存
	for k:= range array {
		array[k] = make([][]int,len(vertices),len(vertices))
		for i := range array[k]{
			array[k][i] = make([]int,len(vertices),len(vertices))
			if k==0{
				for j := range array[k][i]{
					//迭代式循环
					cureEdge := Edge{vertices[i],vertices[j]}
					if i==j{
						array[k][i][j]=0
					}else if !graph.CheckEdge(cureEdge){ //判断如果边长不存在
						array[k][i][j]= math.MaxInt32 //不联通
					}else {
						array[k][i][j] = graph.Weight(cureEdge) //获取当前边长
					}
				}
			}
		}
	}
	return array,vertices  //返回数组  图上的节点

}
//判断 取得最短距离

func DisFloyedHander(array *[][][]int,k,i,j int)  {
	(*array)[k+1][i][j]=(*array)[k][i][j] //取出数据

	if (*array)[k][i][k] + (*array)[k][k][j] < (*array)[k+1][i][j]{
		(*array)[k+1][i][j] = (*array)[k][i][k]+(*array)[k][k][j]
	}//保存最短节点

}

func DistFloyedWarShall(graph WeightGraph) WeightGraph  {
	newGraph := NewAdjacencyMatrixWithWeight()
	//构建一个权重图
	rebuild := func(vertices []interface{},array [][]int) {

		for i:= range vertices{
			for j := range vertices{
				if array[i][j] < math.MaxInt32{
					newGraph.AddWeightEdge(Edge{vertices[i],vertices[j]},array[i][j])
				}
			}
		}
	}
	FloydWarShall(graph,DisFloyedInit,DisFloyedHander,rebuild)
	return newGraph
}

func PathFloyedWarShall(graph WeightGraph)map[interface{}]WeightGraph  {
	var pathArray [][]int

	var pathForest map[interface{}]WeightGraph

	init := func(graph WeightGraph)([][][]int,[]interface{}) {
		disArray,vertices := DisFloyedInit(graph)//初始化
		pathArray = make([][]int,len(vertices),len(vertices)) //计算数组
		for i := range pathArray{
			pathArray[i] = make([]int,len(vertices),len(vertices))
			for j:= range pathArray[i]{
				curEdge := Edge{vertices[i],vertices[j]}
				if i==j || !graph.CheckEdge(curEdge){
					pathArray[i][j] = math.MaxInt32
				}else {
					pathArray[i][j] = i
				}

			}
		}
		return disArray,vertices


	}
	handler := func(array*[][][]int,k,i,j int) {
		if (*array)[k][i][j]>(*array)[k][k][j] + (*array)[k][i][k]{
			pathArray[i][j] = pathArray[k][j]  //保存最短
		}
		DisFloyedHander(array,k,i,j)
	}
	rebuild := func(vertices []interface{},array [][]int) {
		pathForest = make(map[interface{}]WeightGraph)

		for i := range  vertices{
			pathForest[vertices[i]] = NewAdjacencyMatrixWithWeight()
			for j := range  vertices{
				if pathArray[i][j] < math.MaxInt32{
					curEdge := Edge{vertices[pathArray[i][j]],vertices[j]}
					pathForest[vertices[i]].AddWeightEdge(curEdge,graph.Weight(curEdge))
				}
			}
		}
	}
	FloydWarShall(graph,init,handler,rebuild)
	return pathForest
}

func Johnson(graph WeightGraph)WeightGraph  {
	TempG := NewAdjacencyMatrixWithWeight()

	s := struct {}{}
	for _,e := range graph.AllEdges(){
		TempG.AddWeightEdge(e,graph.Weight(e))
		TempG.AddWeightEdge(Edge{s,e.Start},0)
	}
	//relax := new(DefRelax) //新建类型 协助判断最短路径



	return TempG
}
