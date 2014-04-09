package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    result:=make([][]uint8,dy);
    for ind := range result {
        result[ind] = make([]uint8,dx);
        for jnd := range result[ind]{
            result[ind][jnd] = uint8(jnd^ind);
        }
    }
    return result
}

func main() {
    pic.Show(Pic)
}
