package goroutinepool

type Job interface {
    Do(*Executer, ...interface{})
}

