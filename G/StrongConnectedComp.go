package Graph

import (
	"fmt"
	"sort"
)

func StronglyConnect(g Graph) (scc Graph)  {
	scc = NewGraph()

	dfsGraph , gt := DFS(g,nil),g.Transpose()  //将有向图 反转
	//获取所有顶点

	dfsVertices := dfsGraph.AllVertices()
	sort.Slice(dfsVertices, func(i, j int) bool {
		return dfsVertices[i].(*DFS_Element).F > dfsVertices[j].(*DFS_Element).F
	})

	//新建子图  深度遍历镜像图
	dfsGraphfT := DFS(gt, func(vertices []interface{}) {
		for i, v := range dfsVertices{
			vertices[i] = v.(*DFS_Element).V  //取出联通分量
		}
	})

	//fmt.Println(dfsGraphfT.AllVertices())
	//fmt.Println(dfsGraphfT.AllEdges())
	//存储所有子图
	commpent := make(map[*DFS_Element]Graph)
	for i:= range dfsGraphfT.Comps{
		commpent[i] =  NewGraph() //新建图
		//返回所有的节点   也就是节点集合
		for _ , v := range dfsGraphfT.Comps[i].AllVertices(){
			//所有顶点
			commpent[i].AddVertex(v.(*DFS_Element).V)  //追加所有顶点
		}

		//追加所有的边长集合
		for _ , e := range dfsGraphfT.Comps[i].AllEdges(){
			//所有顶点
			commpent[i].AddEdge(Edge{e.End.(*DFS_Element).V,e.Start.(*DFS_Element).V})  //追加所有顶点
		}
	}
	fmt.Println(commpent)
	//判断是否联通
	for _ , e := range dfsGraphfT.AllTreeCrossEdges(){
		//循环走过的边长
		if e.End.(*DFS_Element).FindRoot()!=e.Start.(*DFS_Element).FindRoot(){
			//如果初始点start 找不到根节点  增加边长
			scc.AddEdge(Edge{commpent[e.End.(*DFS_Element).FindRoot()],commpent[e.Start.(*DFS_Element).FindRoot()]})
		}
	}


	return scc
}
