package Graph

import (
	"container/list"
	"math"
)

type BFSElement struct {
	Color int //走过的标记颜色
	Dist  int //距离
	P*BFSElement//指向元素
	V interface{}
	Iter iterator //迭代器
}
//新建广度遍历的节点
func  NewBFSElement(v interface{})*BFSElement{
	return new(BFSElement).Init(v)
}
//初始化
func ( e*BFSElement)Init(v interface{})*BFSElement{
	e.V=v
	e.Color=while
	e.Dist=math.MaxInt64 //标记距离
	e.P=nil //为空
	e.Iter=nil
	return e
}
//深度遍历的访问者
type  BFSVisitHandler struct {
	BeforeBFSHander,AfterBFSHander func(*BFSElement)
	EdgeHander func(*BFSElement,*BFSElement)
	Elements map[interface{}]*BFSElement
}
//初始化
func NewBFSVisitHandler()*BFSVisitHandler{
	return new(BFSVisitHandler).Init()

}
func (bfs *BFSVisitHandler)Init()*BFSVisitHandler{
	//初始化
	bfs.AfterBFSHander= func(element *BFSElement) {}
	bfs.BeforeBFSHander= func(element *BFSElement) {}
	bfs.EdgeHander= func(element *BFSElement, element2 *BFSElement) {}
	bfs.Elements=make(map[interface{}]*BFSElement)

	return bfs
}
func bfsVisit(g Graph, s interface{},hander *BFSVisitHandler)  {
	if hander==nil{
		panic("hander is null")
	}
	queue:=list.New()//开辟list当作队列
	pushQueue:= func(v interface{}) *BFSElement {
		hander.Elements[v]=NewBFSElement(v)//初始化
		hander.Elements[v].Color=gray
		hander.Elements[v].Iter=g.IterConnectedEdges(v)//处理迭起器
		queue.PushBack(hander.Elements[v]) //压入队列
		hander.BeforeBFSHander(hander.Elements[v])
		return hander.Elements[v]
	}
	pushQueue(s).Dist=0//设定距离从0开始
	for  queue.Len()!=0{
		v:=queue.Front().Value.(*BFSElement)//提取队列第一个数据
		//循环遍历所有节点
		for  c:=v.Iter.Value();c!=nil;c=v.Iter.Next(){
			//节点压入队列
			if _,ok :=hander.Elements[c];!ok{
				newE:=pushQueue(c)
				newE.Dist=v.Dist+1
				newE.P=v
				hander.EdgeHander(v,newE) //广度遍历
			}
		}
		v.Color=black//走过
		v.Iter=g.IterConnectedEdges(v.V)//继续循环
		queue.Remove(queue.Front())//删除数据
		hander.AfterBFSHander(v)//继续循环

	}



}

func BFS(graph Graph, s interface{})(bfsgraph Graph){
	bfsgraph=NewGraph()//新建一个图
	handler:=NewBFSVisitHandler()//新建一个访问器
	//函数设置为增加边长
	handler.EdgeHander= func(start *BFSElement, end *BFSElement) {
		bfsgraph.AddEdge(Edge{start.V,end.V})
	}
	bfsVisit(graph,s,handler)
	return


}