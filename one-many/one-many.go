package one_many

import (
	"go-producer-consumer/out"
	"sync"
)

type Task struct {
	ID int64
}

func (t *Task) run() {
	out.Println(t.ID)
}

var taskCh = make(chan Task, 10)

const taskNum int64 = 10000

func producer(wo chan<- Task) {
	var i int64
	for i = 1; i <= taskNum; i++ {
		t := Task{
			ID: i,
		}
		wo <- t
	}

	//生产者关闭通道，仅限于生产者任务是有限的
	close(wo)
}

func consumer(ro <-chan Task) {
	for t := range ro {
		if t.ID != 0 {
			t.run()
		}
	}
}

func Exec() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		producer(taskCh)
	}()

	for i := int64(0); i < taskNum; i++ {
		if i%100 == 0 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				consumer(taskCh)
			}()
		}
	}

	wg.Wait()
	out.Println("执行成功")
}
