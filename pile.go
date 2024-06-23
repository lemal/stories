package main

import "fmt"

// init ititializesstacks of wood each 90 in size

type wood_log struct {
     number int
     wood_type string
}

const pine = "pine"

func initialize(i int) wood_log{
     return wood_log{i, pine}
}
//init initializes the logs based on the starting conditions - birch, pine, palm - selects randomly the type


func collect() int {
     fmt.Println("found wood")
     kek := wood_log{number:1, wood_type:"pine"}
     var lol wood_log
     lol = wood_log{number:2, wood_type:"birch"}
     fmt.Println("collected wood", kek, lol)
     return 1
}