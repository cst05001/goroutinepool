package goroutinepool

import (
)

type Pool struct {
    executers  []*Executer
    poolSize    int
    chanExecuter    chan *Executer
}

func NewPool(poolSize int) *Pool {
    pool := &Pool{}
    pool.Init(poolSize)
    return pool
}

func (this *Pool) Init(poolSize int) {
    this.poolSize = poolSize
    this.executers = make([]*Executer, this.poolSize)
    this.chanExecuter = make(chan *Executer, this.poolSize)
    for i := 0; i < this.poolSize; i++ {
        this.executers[i] = &Executer{}
        this.executers[i].Id = i
        this.executers[i].ChanExecuter = this.chanExecuter
        this.executers[i].ChanExecuter <- this.executers[i]
    }
}

func (this *Pool) GetExecuter() *Executer {
    return <- this.chanExecuter
}

