package Graph

import "container/list"

type LinkedMap struct{
	keyL  *list.List //链表-所有边长，
	mymap   map[interface{}]interface{} //映射关系
}
//数据结构初始化
func  (lm*LinkedMap )init()*LinkedMap {
	lm.keyL=list.New()
	lm.mymap=make( map[interface{}]interface{})
	return lm
}
//判断数据存在或者不存在
func  (lm*LinkedMap )exist(key  interface{})bool {
	_,ok:=lm.mymap[key]
	return ok
}
//插入数据，key,value,存在更新，不存在插入
func  (lm*LinkedMap )add(key  interface{},value interface{}  ) {
	if !lm.exist(key){
		e:=lm.keyL.PushBack(key)//压入数据
		lm.mymap[key]=[]interface{}{e,value}
	}else{
		lm.mymap[key].([]interface{})[1]=value
	}
}
//抓取数据
func  (lm*LinkedMap )get(key  interface{}) interface{}{
	if lm.exist(key){
		return lm.mymap[key].([]interface{})[1]
	}else{
		return nil
	}
}
//删除数据
func  (lm*LinkedMap )delete(key  interface{}) {
	if lm.exist(key){
		i:=lm.mymap[key].([]interface{})
		lm.keyL.Remove(i[0].(*list.Element))//列表删除
		delete (lm.mymap,key)//删除
	}
}
//返回第一个数据
func (lm*LinkedMap )frontkey()interface{}{
	if lm.keyL.Len()==0{
		return nil
	}
	return lm.keyL.Front().Value
}
func (lm*LinkedMap )backkey()interface{}{
	if lm.keyL.Len()==0{
		return nil
	}
	return lm.keyL.Back().Value
}
//取得上一个数据
func (lm*LinkedMap )nextkey(key interface{})interface{}{
	if e:=lm.mymap[key].([]interface{})[0].(*list.Element).Next();e!=nil{
		return e.Value
	}
	return nil
}
//取得下一个数据
func (lm*LinkedMap )prevkey(key interface{})interface{}{
	if e:=lm.mymap[key].([]interface{})[0].(*list.Element).Prev();e!=nil{
		return e.Value
	}
	return nil
}
















