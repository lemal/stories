package main

import "testing"

func TestOne(t *testing.T){
     t.Run("filler function", func(t *testing.T){
     want := 1
     got := collect()
     if want != got {
     	t.Errorf("got %d, want %d", got, want)
     }
     })
}

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