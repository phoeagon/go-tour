package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    f1 := 1
    f2 := 2
    cnt := 0
    return func()int{
        cnt += 1;
        if cnt >= 3 {
            tmp := f1+f2;
            f1 = f2;
            f2 = tmp;
            return tmp
        }else if cnt == 1{
        	return f1
        }else{ return f2;
             }
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
