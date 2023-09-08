package many_one

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

const taskNum int64 = 100
const nums int64 = 100 //nums是taskNum的因数

func producer(wo chan<- Task, startNum int64, nums int64) {
	for i := startNum; i < startNum+nums; i++ {
		t := Task{
			ID: i,
		}
		wo <- t
	}
}

func consumer(ro <-chan Task) {
	for t := range ro {
		if t.ID != 0 {
			t.run()
		}
	}
}

func Exec() {
	wg := sync.WaitGroup{}
	pwg := sync.WaitGroup{}

	for i := int64(0); i < taskNum; i += nums {
		if i >= taskNum {
			break
		}
		wg.Add(1)
		pwg.Add(1)
		go func(i int64) {
			defer wg.Done()
			defer pwg.Done()
			producer(taskCh, i, nums)
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		consumer(taskCh)
	}()

	pwg.Wait()
	close(taskCh)

	wg.Wait()

	out.Println("执行成功")
}
