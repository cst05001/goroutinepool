/*
      
          ┌─┐       ┌─┐
       ┌──┘ ┴───────┘ ┴──┐
       │                 │
       │       ───       │
       │  ─┬┘       └┬─  │
       │                 │
       │       ─┴─       │
       │                 │
       └───┐         ┌───┘
           │         │
           │         │
           │         │
           │         └──────────────┐
           │                        │
           │                        ├─┐
           │                        ┌─┘    
           │                        │
           └─┐  ┐  ┌───────┬──┐  ┌──┘         
             │ ─┤ ─┤       │ ─┤ ─┤         
             └──┴──┘       └──┴──┘ 
       向 Python 艺术家徐凡、路斌致敬。
*/
package goroutinepool

import (
    "errors"
)

var ErrorExecuterIsBusy error
var ErrorAllExecutersAreBusy error

func init() {
    ErrorExecuterIsBusy = errors.New("executer is busy")
    ErrorAllExecutersAreBusy = errors.New("all executers are busy")
}

