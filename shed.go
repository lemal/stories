package main

import (
       "fmt"
       )

func main(){
     count := 0
     count += collect()
     count += load()
     fmt.Println("unloaded")
}
