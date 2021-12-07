package dcache

//ByteView 只读元素 缓存值的抽象与封装
type ByteView struct {
	b []byte //byte类型支持任意的数据类型的存储，例如字符串、图片等
}

func (v ByteView) Len() int {
	return len(v.b)
}

//ByteSlice 返回克隆的值，防止被修改
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
