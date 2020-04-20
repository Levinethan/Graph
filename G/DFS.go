package Graph

import (
	"container/list"
)

type DFS_Element struct {
	Color int //走过的路标记颜色脚印
	Dist,F  int //距离 F标记深度
	P,Root *DFS_Element  //指向元素
	V interface{}
	Iter iterator  //迭代器

}

func NewDFS_Element(v interface{}) *DFS_Element  {
	return new(DFS_Element).Init(v) //初始化
}

func (e *DFS_Element)Init(v interface{}) *DFS_Element  {
	e.V=v
	e.Color = while
	e.P = nil
	e.F =0
	e.Dist = 0
	e.Root = e //根节点代表自己
	e.Iter = nil
	return e
}

func (e *DFS_Element)FindRoot() *DFS_Element  {
	_e := e
	for _e.Root!=_e{
		if _e.Root.Root!=_e.Root{
			_e.Root = _e.Root.Root
		}
		_e = _e.Root
	}
	return _e
}

type DFSforest struct {
	Trees,BackEdges,ForwardEdges,CrossEdges Graph
	Comps map[*DFS_Element]*DFSforest  //开发标准结构  用于实现深度遍历
}


//保存深度遍历的记录
func NewDFS_forest() *DFSforest  {
	return new(DFSforest).Init()
}

func (dfs *DFSforest)Init() *DFSforest  {
	dfs.Trees = NewGraph()
	dfs.BackEdges = NewGraph()
	dfs.ForwardEdges = NewGraph()
	dfs.CrossEdges = NewGraph()
	dfs.Comps = make(map[*DFS_Element]*DFSforest)
	return dfs
}

type DES_VisitHandler struct {
	TreeEdgeHander,BackEdgeHander,ForwardEdgeHander,CrossEdgeHander func(*DFS_Element,*DFS_Element)
	BeforeDFSHander,AfterDFSHander func(element *DFS_Element)
	Element *LinkedMap
	timer int  //深度遍历有时候会栈溢出  走不出来
}

func NewDESVisitHandler() *DES_VisitHandler  {
	return new(DES_VisitHandler).Init()
}

func (d *DES_VisitHandler)Init() *DES_VisitHandler  {
	d.TreeEdgeHander= func(element *DFS_Element, element2 *DFS_Element) {}
	d.BackEdgeHander= func(element *DFS_Element, element2 *DFS_Element) {}
	d.ForwardEdgeHander= func(element *DFS_Element, element2 *DFS_Element) {}
	d.CrossEdgeHander= func(element *DFS_Element, element2 *DFS_Element) {}

	d.AfterDFSHander = func(element *DFS_Element) {}
	d.BeforeDFSHander = func(element *DFS_Element) {}
	d.Element  = new(LinkedMap).init()
	d.timer = 0
	return d
}

func (dfs *DFSforest)addVertex(v *DFS_Element)  {
	dfs.Trees.AddVertex(v)
	dfs.BackEdges.AddVertex(v)
	dfs.ForwardEdges .AddVertex(v)
	dfs.CrossEdges.AddVertex(v)
}

func (dfs *DFSforest)addEdgesTrees(edge Edge)  {
	dfs.Trees.AddEdge(edge)
}
func (dfs *DFSforest)addEdgesBackEdges(edge Edge)  {
	dfs.BackEdges.AddEdge(edge)
}
func (dfs *DFSforest)addEdgesForwardEdges(edge Edge)  {
	dfs.ForwardEdges.AddEdge(edge)
}
func (dfs *DFSforest)addEdgesCrossEdges(edge Edge)  {
	dfs.CrossEdges.AddEdge(edge)
}
func (dfs *DFSforest)getRoot(e *DFS_Element) *DFS_Element  {
	root := e.FindRoot()
	if _, ok := dfs.Comps[root] ; !ok{
		dfs.Comps[root] = NewDFS_forest()  //抓取根节点
	}
	return root
}

func (dfs *DFSforest)AddVertex(v *DFS_Element)  {
	dfs.addVertex(v)
	dfs.Comps[dfs.getRoot(v)].addVertex(v)
}
func (dfs *DFSforest)AddEdgesTrees(edge Edge)  {
	dfs.addEdgesTrees(edge)
	root := dfs.getRoot(edge.Start.(*DFS_Element))
	if root == dfs.getRoot(edge.End.(*DFS_Element)){
		dfs.Comps[root].addEdgesTrees(edge)
	}
}
func (dfs *DFSforest)AddEdgesBackEdges(edge Edge)  {
	dfs.addEdgesBackEdges(edge)
	root := dfs.getRoot(edge.Start.(*DFS_Element))
	if root == dfs.getRoot(edge.End.(*DFS_Element)){
		dfs.Comps[root].addEdgesBackEdges(edge)
	}
}
func (dfs *DFSforest)AddEdgesForwardEdges(edge Edge)  {
	dfs.addEdgesForwardEdges(edge)
	root := dfs.getRoot(edge.Start.(*DFS_Element))
	if root == dfs.getRoot(edge.End.(*DFS_Element)){
		dfs.Comps[root].addEdgesForwardEdges(edge)
	}
}
func (dfs *DFSforest)AddEdgesCrossEdges(edge Edge)  {
	dfs.addEdgesCrossEdges(edge)
	root := dfs.getRoot(edge.Start.(*DFS_Element))
	if root == dfs.getRoot(edge.End.(*DFS_Element)){
		dfs.Comps[root].addEdgesCrossEdges(edge)
	}
}

//结构体的更高级封装
func (dfs *DFSforest)AllVertices() []interface{}   {
	return dfs.Trees.AllVertices()
}

func (dfs *DFSforest)AllTreeEdges() []Edge   {
	return dfs.Trees.AllEdges()
}

func (dfs *DFSforest)AllTreeBackEdges() []Edge   {
	return dfs.BackEdges.AllEdges()
}
func (dfs *DFSforest)AllTreeForwardEdges() []Edge   {
	return dfs.ForwardEdges.AllEdges()
}
func (dfs *DFSforest)AllTreeCrossEdges() []Edge   {
	return dfs.CrossEdges.AllEdges()
}

//所有边长
func (dfs *DFSforest)AllEdges() []Edge   {
	edges := dfs.AllTreeEdges()
	edges = append(edges,dfs.AllTreeForwardEdges()...)
	edges = append(edges,dfs.AllTreeBackEdges()...)
	edges = append(edges,dfs.AllTreeCrossEdges()...)
	return edges
}

func (d *DES_VisitHandler)Counting() int  {
	d.timer ++
	return d.timer
}

//深度遍历有一个问题  死循环  因为长度有可能会无限叠加
func dfsVisit(graph Graph,v interface{},hander *DES_VisitHandler)  {
	if hander == nil{
		panic("hander is nill")
	}

	//压栈函数
	stack := list.New()
	pushStack := func(v interface{}) *DFS_Element {
		newE := NewDFS_Element(v)
		newE.Color = gray
		newE.Dist = hander.Counting()
		newE.Iter = graph.IterConnectedEdges(v)
		//计算迭代器
		hander.Element.add(v,newE)

		stack.PushBack(newE)

		hander.BeforeDFSHander(newE)

		return newE
	}

	//出栈函数

	popStack := func(e *DFS_Element) {
		e.Color = black
		e.F = hander.Counting()
		e.Iter = graph.IterConnectedEdges(e.V)
		stack.Remove(stack.Back())  //删除最后一个 出栈
		hander.AfterDFSHander(e)
	}

	pushStack(v)
	for stack.Len()!=0{

		e:= stack.Back().Value.(*DFS_Element)  //取出数据
		for c := hander.Element.get(e.V).(*DFS_Element).Iter.Value();c!=nil;{

			if !hander.Element.exist(c){
				newE := pushStack(c)
				newE.P = e
				newE.Root = e
				if hander!=nil{
					hander.TreeEdgeHander(e,newE)
				}
				break
			}else if hander.Element.get(c).(*DFS_Element).Color == gray{

				hander.BackEdgeHander(e,hander.Element.get(c).(*DFS_Element))
			}else if e.Dist > hander.Element.get(c).(*DFS_Element).Dist{

				hander.CrossEdgeHander(e,hander.Element.get(c).(*DFS_Element))
			}else if hander.Element.get(c).(*DFS_Element).Dist-e.Dist >1{
				hander.ForwardEdgeHander(e,hander.Element.get(c).(*DFS_Element))
			}
			c = hander.Element.get(e.V).(*DFS_Element).Iter.Next()

		}
		if e == stack.Back().Value.(*DFS_Element){
			popStack(e)
		}

		//提取数据 然后实例化
	}
}

func DFS(graph Graph,sorted func([]interface{})) (dfsforest *DFSforest)  {
	dfsforest = NewDFS_forest() //新建深度遍历保存图
	hander := NewDESVisitHandler()
	hander.BeforeDFSHander= func(element *DFS_Element) {
		dfsforest.AddVertex(element)
	}
	hander.TreeEdgeHander= func(start,end *DFS_Element) {
		dfsforest.AddEdgesTrees(Edge{start,end})
	}
	hander.BackEdgeHander= func(start,end *DFS_Element) {
		dfsforest.AddEdgesBackEdges(Edge{start,end})
	}
	hander.ForwardEdgeHander= func(start,end *DFS_Element) {
		dfsforest.AddEdgesForwardEdges(Edge{start,end})
	}
	hander.CrossEdgeHander= func(start,end *DFS_Element) {
		dfsforest.AddEdgesCrossEdges(Edge{start,end})
	}
	vs := graph.AllVertices()
	//所有顶点
	if sorted!=nil{
		sorted(vs)
	}
	for _ , v := range vs{
		if !hander.Element.exist(v){
			dfsVisit(graph,v,hander)
		}
	}

	return
}


func CheckConnectivity(graph Graph)bool  {
	dfsforest :=  DFS(graph,nil)
	return len(dfsforest.Comps)==1 //判断是否存在回路 是否所有节点联通
}
