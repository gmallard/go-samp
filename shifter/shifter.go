/*
Effects of shifts on signed and unsigned numbers.
*/
package main

import (
  "fmt"
)

func main() {
  d := 0x7fffffff
  fmt.Printf("positive %d, hex is %08x\n", d, d)
  m := -(d)
  fmt.Printf("negative %d, hex is %08x\n", m, m)
  //
  var i int = m
  fmt.Println("i1", i)
  fmt.Printf("i1 is: %d, and hex is: %08x\n", i, i)
  //

  var u uint = uint(m)
  fmt.Println("u1", u)
  fmt.Printf("u1 is: %d, and hex is: %08x\n", u, u)
  //
  var is int = i >> 1
  fmt.Println("is1", is)
  fmt.Printf("is1 is: %d, and hex is: %08x\n", is, is)
  //
  var is31 int = i >> 31
  fmt.Println("is31", is)
  fmt.Printf("is31 is: %d, and hex is: %08x\n", is31, is31)
  //
  var us uint = u >> 1
  fmt.Println("us1", us)
  fmt.Printf("us1 is: %d, and hex is: %08x\n", us, us)
  //
  var us31 uint = u >> 31
  fmt.Println("us31", us31)
  fmt.Printf("us31 is: %d, and hex is: %08x\n", us31, us31)
}

