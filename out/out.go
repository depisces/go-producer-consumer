package out

import "fmt"

type Out struct {
	data chan interface{}
	//data <-chan interface{} //只读
	//data chan<- interface{}	//只写

}

var out *Out

func NewOut() *Out {
	if out == nil {
		out = &Out{
			data: make(chan interface{}, 65535),
		}
	}
	return out
}

func Println(i interface{}) {
	out.data <- i
}

func (o *Out) OutPut() {
	for i := range o.data {
		fmt.Println(i)
	}
}
