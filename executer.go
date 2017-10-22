package goroutinepool

import (
)

type Executer struct {
    Job    Job
    Id  int
    ChanExecuter    chan *Executer
}

func (this *Executer) Execute(job Job, chans ...chan interface{}) {
    job.Do(this, chans...)
    this.ChanExecuter <- this
}


