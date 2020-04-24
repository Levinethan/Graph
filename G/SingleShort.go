package Graph

import (
	"container/list"
	"math"
)

type relax interface {
	InitValue()int   //初始化
	Compare(*ssspElement,*ssspElement,int)bool //对比
	Relax(*ssspElement,*ssspElement,int)bool //放松
}
type ssspElement struct {
	D int   //长度
	P*ssspElement    //指向自己，链表结构记录走过的路
	V interface{} //数据类型
}
func (e*ssspElement)init(v interface{},d int )*ssspElement{
	e.V=v
	e.D=d
	e.P=nil
	return e
}
func NewssspElement(v interface{},d int)*ssspElement{
	return new(ssspElement).init(v,d)
}


type  DefaultRelax struct {
	relax
}
func (r* DefaultRelax)InitValue()int{
	return math.MaxInt32
	//return 10000000
}
func (r* DefaultRelax)Compare(start *ssspElement,end *ssspElement,weight int)bool{
	return end.D>start.D+weight
}
func (r* DefaultRelax)Relax(start *ssspElement,end *ssspElement,weight int)bool{
	if r.Compare(start,end,weight){
		end.D=start.D+weight
		end.P=start
		return true //保存最短路径
	}
	return false
}

func  initSingleSource(g Graph,d int)map[interface{}]* ssspElement{
	ssspE:=make(map[interface{}]* ssspElement)
	for _,v :=range  g.AllVertices(){ //循环所有节点
		ssspE[v]=NewssspElement(v,d)
	}
	return ssspE
}
//增加一个边长
func addSsspGEdge(g,ssspG WeightGraph,ssspE *ssspElement){
	if ssspE.P!=nil{
		//fmt.Println("nimei",ssspE.P.V,ssspE)
		//fmt.Println("xxx",g.Weight(Edge{ssspE.P.V,ssspE.V}))
		ssspG.AddWeightEdge(Edge{ssspE.P.V,ssspE.V},g.Weight(Edge{ssspE.P.V,ssspE.V}))
	}
}
//提取图
func GetSsspGEdge(g WeightGraph,ssspE map[interface{}] *ssspElement) WeightGraph{
	sssPG:=NewAdjacencyMatrixWithWeight()
	for _,e:=range g.AllEdges(){
		addSsspGEdge(g,sssPG,ssspE[e.End]) //追加边长
	}
	return sssPG//返回图
}
//检查存在与否，提取
func CheackorGetSsspGEdge(g WeightGraph,ssspE map[interface{}] *ssspElement,r relax) WeightGraph{
	ssspG:=NewAdjacencyMatrixWithWeight()//新建图
	//fmt.Println("1",ssspG.TotalWeight(),ssspG.AllEdges())
	for _,e:=range g.AllEdges(){
		if ssspE==nil || r.Compare(ssspE[e.Start],ssspE[e.End],g.Weight(e)){
			return nil
		}
		addSsspGEdge(g,ssspG,ssspE[e.End])
	}
	return ssspG
}

//包装器,所有的最短路径算法都用这个接口
func sssssssssssssssspWrapper(core func( WeightGraph,interface{},relax) map[interface{}] *ssspElement )func(graph WeightGraph,s interface{},relax relax)WeightGraph  {
	return func(graph WeightGraph,s interface{},r relax)WeightGraph{
		return CheackorGetSsspGEdge(graph,core(graph,s,r),r)
	}
}
func ssspposWeightWrapper(core func( WeightGraph,interface{},relax) map[interface{}] *ssspElement)func(graph WeightGraph,s interface{},relax relax)WeightGraph{
	return func(graph WeightGraph,s interface{},r relax)WeightGraph{
		return  GetSsspGEdge(graph,core(graph,s,r))
	}
}

func bellmanFordCore(graph WeightGraph,s interface{},r relax)map[interface{}]*ssspElement{
	sssPE:=initSingleSource(graph,r.InitValue()) //初始化

	sssPE[s].D=0

	for i:=0;i<len(sssPE)-1;i++{
		for _,e:=range graph.AllEdges(){
			r.Relax(sssPE[e.Start],sssPE[e.End],graph.Weight(e))//统计距离
		}
	}
	return sssPE


}
func BellmanFord(graph WeightGraph,s interface{},r relax)WeightGraph{
	return sssssssssssssssspWrapper(bellmanFordCore)(graph,s,r) //包装调用
}

func BellmanFordQueueCore(graph WeightGraph,s interface{},r relax)map[interface{}]*ssspElement  {
	sssPE:=initSingleSource(graph,r.InitValue()) //初始化

	sssPE[s].D=0
	visit := make(map[interface{}]int)  //访问器
	//进队 出队
	queue := list.New() //新建一个队列
	queue.PushBack(sssPE[s])

	for queue.Len() !=0{
		v := queue.Front().Value.(*ssspElement)
		iter := graph.IterConnectedVertices(v.V)
		for e:= iter.Value();e!=nil;e = iter.Next(){
			//对每个节点 判断
			if r.Relax(v,sssPE[e],graph.Weight(Edge{v.V,e})){
				//取最短  内部还需要进行操作
				if cnt,ok := visit[e];ok && cnt > len(visit){
					return nil
				}
			}
			queue.PushBack(sssPE[e])
			visit[e]++


		}
		queue.Remove(queue.Front())
	}
	return sssPE
}

func BellmanFordQueue(graph WeightGraph,s interface{},r relax) WeightGraph {
	return sssssssssssssssspWrapper(BellmanFordQueueCore)(graph ,s,r)  //包装调用
}

func DijstraCore(graph WeightGraph,s interface{},r relax)map[interface{}]*ssspElement  {
	sssPE:=initSingleSource(graph,r.InitValue()) //初始化

	sssPE[s].D=0
	visit := make(map[interface{}]int)  //访问器
	//进队 出队
	queue := list.New() //新建一个队列
	queue.PushBack(sssPE[s])

	for queue.Len() !=0{
		v := queue.Front().Value.(*ssspElement)
		iter := graph.IterConnectedVertices(v.V)
		for e:= iter.Value();e!=nil;e = iter.Next(){
			//对每个节点 判断
			if r.Relax(v,sssPE[e],graph.Weight(Edge{v.V,e})){
				//取最短  内部还需要进行操作
				if cnt,ok := visit[e];ok && cnt > len(visit){
					return nil
				}
			}
			queue.PushBack(sssPE[e])
			visit[e]++


		}
		queue.Remove(queue.Front())
	}
	return sssPE
}

func Dijstra(graph WeightGraph,s interface{},r relax) WeightGraph {
	return sssssssssssssssspWrapper(BellmanFordQueueCore)(graph ,s,r)  //包装调用
}

//强连通分量 ， 也可以用于寻路，强连通切割 调用上述寻路算法
func Gabow(graph WeightGraph,s interface{},r relax , k uint32) WeightGraph  {
	degree := k
	if degree ==0{
		degree = 32  //标记边长的出度与入度

	}
	SSSSSSPE := initSingleSource(graph,r.InitValue())
	//update 函数指针
	updataSSSSSSSSSSSSPE := func(currentSSSSSSPE map[interface{}]*ssspElement) {
		//内部调用
		for v := range currentSSSSSSPE{
			if SSSSSSPE[v].D != r.InitValue(){
				currentSSSSSSPE[v].D = currentSSSSSSPE[v].D+(SSSSSSPE[v].D<<1)  //判断联通

			}

			SSSSSSPE[v] = currentSSSSSSPE[v] //指针用于强联通判断  保存
		}
	}
	gi := NewAdjacencyMatrixWithWeight()  //新建一个权重图
	updateGi := func(j uint32) {
		for _ , e := range graph.AllEdges(){
			gi.AddWeightEdge(e,
				(graph.Weight(e)>>j)+((SSSSSSPE[e.Start].D-SSSSSSPE[e.End].D)<<1))

		}//处理距离


	}
	for i := uint32(0);i<degree;i++{
		updateGi(degree-1-i)
		CurE :=BellmanFordQueueCore(gi,s,r)
		updataSSSSSSSSSSSSPE(CurE)
	}
	return GetSsspGEdge(graph,SSSSSSPE)
}


//定义一个结构体  解决  近临值问题
type NestedBoxesRelax struct {
	maxLen int
	lastE *ssspElement
	DefaultRelax
}

func (n *NestedBoxesRelax)Init () *NestedBoxesRelax  {
	n.maxLen = 0
	n.lastE = nil
	return n
}

func (n *NestedBoxesRelax)Relax (start , end  *ssspElement, weight int) bool  {
	update := n.DefaultRelax.Relax(start,end,weight)  //更新
	if update {
		if end.D < n.maxLen{
			n.maxLen, n.lastE = end.D,end
		}
	}
	return update
}

func NestedBoxes(boxes [][]int) [][]int  {
	//KNN预测  传入二维数组参数  传出的也是 二维数组参数
	//新建一个权重图
	g := NewAdjacencyMatrixWithWeight()


	//判断两个数 是否接近
	nested := func(box1 , box2 []int) bool {
		if len(box1) != len(box2){
			return false
		}

		for i:= range box1{
			//长度不相等  没必要对比
			if box1[i] >= box2[i]{
				return false
			}
		}

		return true
	}
	roooooooot := struct {}{}
	for i := range boxes{
		g.AddWeightEdge(Edge{roooooooot,&boxes[i]},0)

		for j := i+1;i<len(boxes);j++{
			if nested(boxes[i],boxes[j]){
				//判断是否近似值
				g.AddWeightEdge(Edge{&boxes[i],&boxes[j]},-1)
			}else if nested(boxes[j],boxes[i]){
				g.AddWeightEdge(Edge{&boxes[j],&boxes[i]},-1)
			}
		}
	}
	NestedBoxesR := new(NestedBoxesRelax)   //初始化
	DijstraCore(g,roooooooot,NestedBoxesR)  //寻路算法

	seq := make([][]int,0,0)
	for e := NestedBoxesR.lastE;e.V !=roooooooot;e = e.P{
		seq = append(seq,*e.V.(*[]int))
	}
	return seq
}

type KarpElement struct {
	//数据结构  最大流karp算法   核心：使用到队列
	k int

	u float64

	sssssssssPE []*ssspElement
}

func (e *KarpElement)Init(n int, v interface{},init int) *KarpElement  {
	e.sssssssssPE = make([]*ssspElement,n+1,n+1) //开辟内存
	for i := range e.sssssssssPE{
		e.sssssssssPE[i] = NewssspElement(v,init)
	}
	e.k = 0
	e.u = math.MaxInt32
	return e
}


func (e *KarpElement)GetssssssssssssssssssPE() *ssspElement  {
	return e.sssssssssPE[e.k]  //数据结构中的 k 为索引
}

func (e *KarpElement)sum ()  {
	//max  (d_n - d_k)/ n-k

	for i:= range e.sssssssssPE[:len(e.sssssssssPE)-1]{
		if max := float64(e.sssssssssPE[len(e.sssssssssPE)-1].D - e.sssssssssPE[i].D)  / float64(len(e.sssssssssPE)-1-i);max > e.u{
			e.u = max  //用 e.u 存储最大值
		}
	}
}

func Karp(graph WeightGraph, s interface{}) float64  {
	//深度遍历用 栈  。广度遍历用队列
	karPE := make(map[interface{}]*KarpElement)

	r := new(DefaultRelax)
	vertices := graph.AllVertices()  //所有顶点
	for _ , v := range vertices {
		karPE[v] = new(KarpElement).Init(len(vertices),v,r.InitValue())  //初始化数据
	}

	karPE[s].sssssssssPE[0].D = 0  //把第一个位置  设置为0
	queue := list.New()
	queue.PushBack(karPE[s]) //压入数据

	//递归计算
	for queue.Len()!=0{
		ke := queue.Front().Value.(*KarpElement) //初始胡  取出队列第一个位置
		v := ke.GetssssssssssssssssssPE()
		iter := graph.IterConnectedVertices(v.V)
		//构造 迭代器
		//循环 求出最大流量
		for e:= iter.Value();e!=nil;e=iter.Next(){
			if ke.k < len(vertices){
				//如果 k 小于当前节点数量
				karPE[e].k = ke.k + 1
				if r.Relax(v,karPE[e].GetssssssssssssssssssPE(),graph.Weight(Edge{v.V,e})){
					queue.PushBack(karPE[e])
				}
			}
		}
		queue.Remove(queue.Front()) //弹出队列
	}
	//计算最短的  最大流
	u := float64(math.MaxInt64)
	for v:= range karPE{
		if karPE[v].sum();karPE[v].u <u{
			u = karPE[v].u
		}
	}
	return u
}