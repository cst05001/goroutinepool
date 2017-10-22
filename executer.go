package goroutinepool

import (
)

type Executer struct {
    Job    Job
    Id  int
    ChanExecuter    chan *Executer
}

func (this *Executer) Execute(job Job, params ...interface{}) {
    job.Do(this, params...)
    this.ChanExecuter <- this
}


