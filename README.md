a goroutine pool

老龚 said goroutine pool is meaningless, but I still wrote one.

# usage

```
package main

import (
    "github.com/cst05001/goroutinepool"
    "log"
    "fmt"
    "time"
)

type MyJob struct {
    I   int
}

func (this *MyJob) Do(executer *goroutinepool.Executer) {
    log.Print(fmt.Sprintf("#%d: MyJob.Do() start\n", executer.Id))
    fmt.Printf("Hello, %d\n", this.I)
    time.Sleep(time.Second * 5)
    log.Print(fmt.Sprintf("#%d: MyJob.Do() end\n", executer.Id))
}

func processor() {
    pool := *goroutinepool.NewPool(4)
    for i := 0; i < 10; i++ {
        executer := pool.GetExecuter()
        go executer.Execute(&MyJob{I: i})
    }
}

func main() {
    go processor()
    time.Sleep(time.Second * 30)
}

```

output

![screenshot1][1]

[1]: screenshot1.png "screenshot1.png"
