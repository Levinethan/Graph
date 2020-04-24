package Graph

import (
	"container/list"
	"math"
)

type HopcraftKarp struct {
	g Graph
	dis int
	xMatch , yMatch map[interface{}]interface{}
	xlevel , ylevel map[interface{}]int
	matches int

}

func (h *HopcraftKarp)Init (g Graph)*HopcraftKarp  {
	h.g = g
	h.xMatch,h.yMatch = make(map[interface{}]interface{}),make(map[interface{}]interface{})
	h.matches = 0
	return  h
}

func (h *HopcraftKarp)bfs() bool  {
	h.dis = math.MaxInt32  //首先距离设置为 最大 然和广度遍历
	h.xlevel,h.ylevel = make(map[interface{}]int), make(map[interface{}]int) //初始化
	 queue := list.New()
	 for _ , x := range h.g.AllVertices(){
	 	if _ , ok := h.xMatch[x];!ok{
	 		queue.PushBack(x)
	 		h.xlevel[x] = 0
		}
	 }
	 for queue.Len() != 0{
	 	s := queue.Front().Value //取出数据
	 	queue.Remove(queue.Front()) //删除数据
	 	if v,ok := h.xlevel[s];ok && v >h.dis{
	 		break
		}
		iter := h.g.IterConnectedVertices(s)
		for y := iter.Value();y !=nil;y=iter.Next(){
			if _ ,ok := h.ylevel[y];!ok{
				h.ylevel[y] = h.xlevel[s] + 1
				if _,ok := h.yMatch[y];!ok{
					h.dis = h.ylevel[y]
				}else {
					h.xlevel[h.yMatch[y]] = h.ylevel[y] + 1
					queue.PushBack(h.yMatch[y])
				}
			}
		}
	 }
	 return h.dis!=math.MaxInt32
}

func (h *HopcraftKarp)dfs(x interface{}, yVsit map[interface{}]bool) bool  {
	iter := h.g.IterConnectedVertices(x)  //获取迭代器
	for y := iter.Value();y !=nil;y=iter.Next(){
		if _,ok := yVsit[y];!ok && h.ylevel[y]==h.xlevel[x]+1{
			yVsit[y] = true
			if _ , ok := h.yMatch[y];ok && h.ylevel[y] == h.dis{
				continue
			}
			if _,ok := h.yMatch[y];!ok{
				h.xMatch[x] = y
				h.xMatch[y] = x
				return true
			}else if h.dfs(h.yMatch[y],yVsit){
				h.xMatch[x] = y
				h.xMatch[y] = x
				return true
			}
		}
	}
	return false
}

func (h *HopcraftKarp)maxMatch() int  {
	for h.bfs(){
		yVisti := make(map[interface{}]bool)
		for _ , x := range h.g.AllVertices(){
			if _ , ok := h.xMatch[x];!ok{
				if h.dfs(x,yVisti){
					h.matches --
				}
			}
		}
	}

	return h.matches
}