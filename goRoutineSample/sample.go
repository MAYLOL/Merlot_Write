package goRoutineSample

import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
)

type Message struct {
	Id   int
	Name string
}

func main() {
	messages := make(chan Message, 100000)
	result := make(chan error, 100000)

	// 创建任务处理Worker
	for i := 0; i < 100; i ++ {
		go worker(i, messages, result)
	}





	total := 0
	// 发布任务
	for k := 1; k <= 10000; k ++ {
		messages <- Message{Id: k, Name: "job" + strconv.Itoa(k)}
		total += 1
	}

	close(messages)

	// 接收任务处理结果
	for j := 1; j <= total; j ++ {
		res := <-result
		if res != nil {
			fmt.Println(res.Error())
		}
	}

	close(result)
}

func worker(worker int, msg <-chan Message, result chan<- error) {
	// 从通道 chan Message 中监听&接收新的任务
	for job := range msg {
		fmt.Println("worker:", worker, "msg: ", job.Id, ":", job.Name)

		// 模拟任务执行时间
		time.Sleep(time.Second * time.Duration(RandInt(1, 3)))

		// 通过通道返回执行结果
		result <- nil
	}
}

func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}