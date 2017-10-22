package goroutinepool

type Job interface {
    Do(*Executer, ...chan interface{})
}

