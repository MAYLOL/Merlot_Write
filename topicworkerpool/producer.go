package topicWorkerPool

import "merlot_project/merlot_write/model"

type Producer struct {
	job Job
}

func NewProducer(topic model.Topic) (*Producer) {
	job := Job{topic:topic}
	return &Producer{job: job}
}

func (p Producer) Output(top model.Topic) {

	work:= Job{topic:top}

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