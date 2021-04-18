package main

import (
    "fmt"
    // "time"
)

func main() {

    a := 2

    whoami := func(i interface{}) {
        switch t:= i.(type) {
        case int:
            println(" int ")
        case bool:
            println(" bool ")
        case string:
            println(" string ")
        default:
            fmt.Printf(" unknow type %T\n", t)
        }
    }

    whoami(a)
    whoami(true)
    whoami("stringa")
    whoami(1.2233)
    whoami([]int{1,2,3})

}