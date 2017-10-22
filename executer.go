package goroutinepool

import (
)

type Executer struct {
    Job    Job
    Id  int
    ChanExecuter    chan *Executer
}

func (this *Executer) Execute(job Job) {
    job.Do(this)
    this.ChanExecuter <- this
}


