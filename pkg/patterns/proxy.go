package patterns

import "fmt"

// IRealObject interface
type IRealObject interface {
	performAction()
}

// RealObject struct
type RealObject struct{}

func (realObject *RealObject) performAction(){
	fmt.Println("RealObject performAcction()")
}

//VirtualProxy struct 
type VirtualProxy struct{
	realObject *RealObject 
} 

// VirtualProxy class method performAction
func (virtualProxy *VirtualProxy) performAction() {
	if virtualProxy.realObject == nil {
		virtualProxy.realObject = new(RealObject)
	}
	fmt.Println("Virtual Proxy performAction()")
	virtualProxy.realObject.performAction()
}

func ProxyFunc() {
	object := VirtualProxy{}
	object.performAction()
}
