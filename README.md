# MyFirstGoProject

In this project, I wrote basic main file in golang for my internship at the blockchain center.


## How it works 

open a go file, mark it as package main and then import the "fmt" package which formats go programs.

```package main

import (
	"fmt"
	
)
```
Next you will need to implement a function to the addition called "Plus" whose parameters are a and b, both being integers, returning a + b.

```
func Plus(a int, b int) int {

    return a + b
}
```

Next we write the main function, which inputs the values for a and b and exxecutes the code.

```
func main() {

    res := Plus(1, 2)
    fmt.Println("1+2 =", res)  
 
}
```

This code should complete the main function.

### How to test it

At the top mark your package as package_main and import the test package and the go module for the main file.

```
package main_test

import (
	
	"testing"
	main "github.com/kaleb/firstproject"
)
```

Next you must write the function to test the main function, hence you will name the function 'TestMain' whose parameters are t *testing.T .

Inside the function, you will set total equal to the two integers of main.Plus which a re 1 and 2.

```
unc TestMain( t *testing.T) {
 
total := main.Plus(1, 2)
```

The function total must return 3 otherwise, the error message will read "Incorrect total, got: %d, want: %d".

```
if total != 3 {
	t.Errorf("Incorrect total, got: %d, want: %d", total, 3)
}
 
}
```

To run the test in the terminal, simply run go test .

```
kalebamarante@KALEBS-MBP ~ % go test .
ok  	github.com/kaleb/firstproject	0.492s
```
