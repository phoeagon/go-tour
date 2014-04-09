package main

import (
    "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    // a call to fmt.Print(e) inside the Error method will send 
    //the program into an infinite loop. You can avoid this by 
    //converting e first: fmt.Print(float64(e)). Why?
    return fmt.Sprintf("cannot Sqrt negative number: %v",float64(e))
}

func Sqrt(f float64) (float64, error) {
    if  f<0 {
    	return 0,ErrNegativeSqrt(f)
    }
    x:=f;
    z:=1.0;
    for i:=0; i<10; i+=1 {
        z=z-(z*z-x)/(2*z);
    }
    return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
