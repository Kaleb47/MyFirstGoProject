package main_test

import (
	
	"testing"
	main "github.com/kaleb/firstproject"
)

func TestMain( t *testing.T) {
 
total := main.Plus(1, 2)

if total != 3 {
	t.Errorf("Incorrect total, got: %d, want: %d", total, 3)
}



  
}