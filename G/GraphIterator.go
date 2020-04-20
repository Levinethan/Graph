package Graph

//迭代器接口
type  iterator interface {
	Len()int
	Next()interface{}
	Value()interface{}
}

type LinkedMapItertorator struct {
	m* LinkedMap//访问的结构
	key interface{}
	iterator
}
func NewLinkedMapItertorator(m* LinkedMap)*LinkedMapItertorator{
	return new(LinkedMapItertorator).init(m)
}
func(i*LinkedMapItertorator )init(m* LinkedMap)*LinkedMapItertorator{
	i.m=m
	i.key=i.m.frontkey()//第一个
	return i
}
func(i*LinkedMapItertorator )Len()int{
	return i.m.keyL.Len()//长度
}
func(i*LinkedMapItertorator )Next()interface{}{
	if i.key==nil{
		return nil
	}
	if i.key=i.m.nextkey(i.key);i.key==nil{
		return nil
	}
	return i.key
}
func(i*LinkedMapItertorator )Value()interface{}{
	if i.key==nil{
		return nil
	}
	return i.key
}