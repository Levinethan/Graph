package Graph



const (
	while=0  //图的类型，无向图，有向图，混合图
	gray=1
	black=2
)
type  Edge struct{
	Start ,End interface{} //开始点，结束点
}

type Graph interface {
	AddVertex(interface{})//增加顶点
	CheckVertex(interface{})bool //判断顶点是否存在
	DeleteVertex(interface{})//删除顶点

	AddEdge(edge Edge)//增加边长
	CheckEdge(edge Edge)bool  //判断边长是否联通
	DeleteEdge(edge Edge)//删除边长

	AddEdgeBi(edge Edge)//增加边长,处理无向
	DeleteEdgeBi(edge Edge)//删除边长


	AllVertices()[]interface{}//返回所有的顶点
	AllEdges()[] Edge //返回所有的边长

	AllConnectedVertices(v interface{})[]interface{}//返回所有的联通的顶点
	IterConnectedEdges(v interface{})iterator //返回与一个点所有的联通的边长


	Transpose()Graph //返回旋转图




}
//t图的数据结构
type  AdjacencyMatrix struct {
	matrix LinkedMap
	Graph
}
//构建图的数据结构初始化
func (g*AdjacencyMatrix)init()*AdjacencyMatrix{
	g.matrix=*new(LinkedMap).init() //初始化
	return g
}
//增加定点
func (g*AdjacencyMatrix)AddVertex(vertex  interface{}){
	if !g.matrix.exist(vertex){
		g.matrix.add(vertex,new(LinkedMap).init())
	}
}
//判断顶点是否存在
func (g*AdjacencyMatrix)CheckVertex(vertex  interface{})bool {
	return  g.matrix.exist(vertex)
}
//删除顶点
func (g*AdjacencyMatrix)DeleteVertex(vertex   interface{}){
	g.matrix.delete(vertex)//删除定点
	//循环删除每个边长
	for v:=g.matrix.frontkey();v!=nil;v=g.matrix.nextkey(v){
		g.DeleteEdge(Edge{v,vertex}) //删除相关的边长
	}
}
//增加边长
func (g*AdjacencyMatrix)AddEdge(edge Edge){
	g.AddVertex(edge.Start)
	g.AddVertex(edge.End)
	g.matrix.get(edge.Start).(*LinkedMap).add(edge.End,true)//增加边长
}
//判断边长是否联通
func (g*AdjacencyMatrix)CheckEdge(edge Edge)bool{
	if  !g.CheckVertex(edge.Start){  //判断点是否存在
		return false
	}
	return g.matrix.get(edge).(*LinkedMap).exist(edge.End) //返回边长是否存在
}
//删除边长
func (g*AdjacencyMatrix)DeleteEdge(edge Edge){
	if  g.matrix.exist(edge.Start){
		g.matrix.get(edge.Start).(*LinkedMap).delete(edge.End)
	}
}



//增加边长,处理无向
func (g*AdjacencyMatrix)AddEdgeBi(edge Edge){
	g.AddVertex(edge.Start)
	g.AddVertex(edge.End)
	g.matrix.get(edge.Start).(*LinkedMap).add(edge.End,true)//增加边长
	g.matrix.get(edge.End).(*LinkedMap).add(edge.Start,true)//增加边长
}
//删除边长  处理无向
func (g*AdjacencyMatrix)DeleteEdgeBi(edge Edge){
	g.DeleteEdge(edge)
	g.DeleteEdge(Edge{edge.End,edge.Start})
}


//返回所有的顶点
func (g*AdjacencyMatrix) AllVertices()[]interface{}{
	keys:=make([]interface{},0,0)//开辟内存
	for v:=g.matrix.frontkey();v!=nil;v=g.matrix.nextkey(v){
		keys=append(keys,v)
	}
	return keys
}
//返回所有的边长,没有联通的边长算。
func (g*AdjacencyMatrix) AllEdges()[] Edge{
	edges:=make([]Edge,0,0)//开辟内存
	for start:=g.matrix.frontkey();start!=nil;start=g.matrix.nextkey(start){
		for end:=g.matrix.get(start).(*LinkedMap).frontkey();end!=nil;end=g.matrix.get(start).(*LinkedMap).nextkey(end){
			edges=append(edges,Edge{start,end})
		}
	}
	return edges
}
//返回所有的联通的顶点与某一个定点
func (g*AdjacencyMatrix)AllConnectedVertices(v interface{})[]interface{}{
	keys:=make([]interface{},0,0)//开辟内存
	if g.matrix.exist(v){
		for key:=g.matrix.get(v).(*LinkedMap).frontkey();key!=nil;key=g.matrix.get(v).(*LinkedMap).nextkey(key){
			keys=append(keys,key)
		}
	}

	return keys
}
//返回所有的联通的边长
func (g*AdjacencyMatrix) 	IterConnectedEdges(v interface{})iterator{

	if g.matrix.exist(v) {
		return  NewLinkedMapItertorator(g.matrix.get(v).(*LinkedMap))
	}else{
		return nil
	}
}

//返回旋转图
func  NewGraph()Graph{
	return  new(AdjacencyMatrix).init()
}
//返回景象图
func (g*AdjacencyMatrix)Transpose()Graph {
	gt:= NewGraph()
	for _,e:=range g.AllEdges(){
		gt.AddEdge(Edge{e.End,e.Start})
	}
	return gt
}









