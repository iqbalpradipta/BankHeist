package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  var isHeistOn bool = true
  var eludedGuards = rand.Intn(100)
  if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    fmt.Println("Plan a better disguise next time?")
  }
  fmt.Println("Status isHeistOn = ", isHeistOn)
  var openedVault = rand.Intn(100)
  if isHeistOn == true && openedVault >= 70 {
    fmt.Println("Grab and GO!")
  } else if isHeistOn == true && openedVault < 70 {
    fmt.Println("Vault Can't be opened!")
    fmt.Println("Status isHeistOn = ", !isHeistOn)
  }
  var leftSafely = rand.Intn(5)
  if isHeistOn == true {
    switch leftSafely{
      case 0:
        isHeistOn = false
        fmt.Println("Turns out vault doors don't open from the inside...")
      case 1:
        isHeistOn = false
        fmt.Println("Police will be coming, Hide from the police")
      case 2:
        isHeistOn = false
        fmt.Println("Heist already failed, because Exit door is broke")
      case 3:
        isHeistOn = false
        fmt.Println("Go away from there, heist already failed")
      default:
        isHeistOn = false
        fmt.Println("Start the gateway car")
    }
  }
  if isHeistOn {
    var amtStolen = 10000 + rand.Intn(1000000)
    fmt.Println(amtStolen)
  }
}
