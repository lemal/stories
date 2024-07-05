package main

import "testing"

func TestOne(t *testing.T){
     t.Run("init pine", func(t *testing.T){
	want := wood_log{1, "pine"}
        got := initialize()
        if want != got {
            t.Errorf("got %q, want %q", got, want)
        }
     })
     
     t.Run("init number of log type by its receiver method", func(t *testing.T){
        want := 2
	got := wood_log{0,"pine"}.insert_num(2).number
	if got != want {
	   t.Errorf("got %q, want %q", got, want)
	}
     })
     
}