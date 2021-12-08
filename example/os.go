package main

import (
    "os"
    "fmt"
    "time"
    "sync"
)

var globalMap sync.Map


func routine(){
    globalMap.Store(os.Getpid(),1)
    globalMap.Store(os.Getppid(),-1)
//    fmt.Println(os.Getpid())
//    fmt.Println(os.Getppid())
    time.Sleep(30 * time.Second)
}

func main() {
    fmt.Println(os.Getpid())
    fmt.Println(os.Getppid())
    for i := 0; i<10000; i++ {
         go routine()
    }
    time.Sleep(10 * time.Second)
    fmt.Println(globalMap)
//    time.Sleep(100 * time.Second)
}
