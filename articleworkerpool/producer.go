package articleWorkerPool

import (
	"merlot_project/merlot_write/model"
	"time"
	"fmt"
)

type Producer struct {
	job Job
}

func NewProducer(art model.Article) (*Producer) {
	job := Job{article:art}
	return &Producer{job: job}
}

func (p Producer) Output(art model.Article) {

	work:= Job{article:art}
    fmt.Println("work=======",work)
	JobQueue <- work
	time.Sleep(time.Second*1)

}


//func payloadHandler(){
//	//Go through each payload and queue items individually to be posted to S3
//	//for _, payload := range
//	payload := Payload(4)
//	work := Job{Payload: payload}
//
//	//Push the work onto the queue.
//	JobQueue <- work
//}