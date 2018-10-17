package diaryWorkerPool


import (
	"fmt"
	"merlot_project/merlot_write/model"
)

var (
	//MaxWorker = os.Getenv("MAX_WORKERS")
	//MaxQueue = os.Getenv("MAX_QUEUE")
	MaxWorker = 5
	MaxQueue = 1
)

//For example usage


func Do(diary model.Diary) (err error){

	//fmt.Println("title=====",diary.Title)
	//fmt.Println("aid====",diary.DiaryID)


	err = nil
	return
}

// Job represents the job to be run
//type Job struct {
//	Payload Payload
//}




type Job struct {
	diary model.Diary
}
// A buffered channel that we can send work requests on.

var JobQueue chan Job
func init(){
	JobQueue = make(chan Job, 100000)
}
// Worker represents the worker that executes the job

type Worker struct{
	WorkerPool chan chan Job
	JobChannel chan Job
	quit	   chan bool
}

func NewWorker(workerPool chan chan Job) Worker{
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit: make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel
// in case we need to stop it
func (w Worker) Start(){

	go func(){

		for{
			//register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job:= <- w.JobChannel:
				// we have received a work request.
				if err := Do(job.diary); err != nil{
					fmt.Printf("Error do payload function ：%s", err.Error())
				}
			case <- w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop()  {
	go func() {
		w.quit <-true
	}()
}