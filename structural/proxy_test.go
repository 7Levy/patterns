package structural

import "testing"

func TestProxyObject_ObjDo(t *testing.T) {
	pb := ProxyObject{}
	pb.ObjDo("Run")
}
