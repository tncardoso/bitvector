Simple BitVector for Go
=======================

This package contains a simple BitVector implementation. 

Installing
----------

To install `bitvector` run:

    go get github.com/tncardoso/bitvector

Usage Example
--------------

    package main
    
    import (
        "fmt"
        "github.com/tncardoso/bitvector"
    )
    
    func main () {
        bv1 := bitvector.New(10)
        bv1.Set(0, true)
        bv1.Set(1, true)
    
        bv2 := bitvector.New(10)
        bv2.Set(1, true)
        bv2.Set(2, true)
    
        and, err := bv1.And(bv2)
        val, err := and.Get(1)
        if err != nil || val == false {
            fmt.Printf("error\n")
        } else {
            fmt.Printf("ok\n")
        }   
    
        val, err = and.Get(2)
        if err != nil || val == true {
            fmt.Printf("error\n")
        } else {
            fmt.Printf("ok\n")  
        }   
    }

