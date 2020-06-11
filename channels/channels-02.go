package main

import (
  "fmt"
  "time"
)

func write(flag chan bool) {
  time.Sleep(2000 * time.Millisecond)
  fmt.Println("Write")
  flag <- true
}

func main() {
  flag := make(chan bool)  
  
  go write(flag)
  fmt.Println("Waiting...")
  
  // Thread locks 
  // until `flag` has a value
  fmt.Println("Unlocked", <- flag)
}