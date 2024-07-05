package main

import "fmt"

// init ititializes stacks of wood each 90 in size

type wood_log struct {
     number int
     wood_type string
}

const pine = "pine"

func initialize() wood_log{
     return wood_log{1, pine}
}

func (w_l *wood_log) insert_num(num int) wood_log {
     w_l.number = num
     return *w_l
}//refactor to make it a receiver method

func collect(w_l wood_log){
     fmt.Println("collected wood", w_l.number, w_l.wood_type)
}