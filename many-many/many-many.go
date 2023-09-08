package many_many

import (
	"fmt"
	"go-producer-consumer/out"
	"time"
)

type Task struct {
	ID int64
}

func (t *Task) run() {
	out.Println(t.ID)
}

var taskCh = make(chan Task, 10)
var done = make(chan struct{})

const taskNum int64 = 10000

func producer(wo chan<- Task, done chan struct{}) {
	var i int64
	for {
		if i >= taskNum {
			i = 0
		}
		i++
		t := Task{
			ID: i,
		}

		select {
		case wo <- t:
		case <-done:
			out.Println("生产者退出")
			return
		}
	}
}

func consumer(ro <-chan Task, done chan struct{}) {
	for {
		select {
		case t := <-ro:
			if t.ID != 0 {
				t.run()
			}
		case <-done:
			for t := range ro { //信号结束时带缓冲通道里可能还有内容
				if t.ID != 0 {
					t.run()
				}
			}
			out.Println("消费者退出")
			return
		}
	}
}

func Exec() {
	for i := 1; i <= 10; i++ {
		go producer(taskCh, done)
	}
	for i := 1; i <= 10; i++ {
		go consumer(taskCh, done)
	}

	time.Sleep(time.Second * 3)
	close(done)
	time.Sleep(time.Second * 1)
	close(taskCh)
	fmt.Println("执行完毕")
}
