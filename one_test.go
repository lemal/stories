package main

import "testing"

func TestOne(t *testing.T){
        t.Run("init num pines", func(t *testing.T){
	num := 5
	want := wood_log{num, "pine"}
        got := initialize(4)
        if want != got {
            t.Errorf("got %q, want %q", got, want)
        }
     })
     
     t.Run("actions done test func", func(t *testing.T){
        want := 1
        got := collect()
        if want != got {
            t.Errorf("got %d, want %d", got, want)
        }
     })
}//collect inits the 

/*
func TestTwo(t *testing.T){
     t.Run("first", func(t *testing.T){
     want := 1
     got := collect()
     if want != got {
     	t.Errorf("got %d, want %d", got, want)
     }
     })
}*/