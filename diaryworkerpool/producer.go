package diaryWorkerPool

import "merlot_project/merlot_write/model"

type Producer struct {
	job Job
}

func NewProducer(diary model.Diary) (*Producer) {
	job := Job{diary:diary}
	return &Producer{job: job}
}

func (p Producer) Output(diary model.Diary) {

	work:= Job{diary:diary}

	JobQueue <- work
	//time.Sleep(time.Second*1)

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