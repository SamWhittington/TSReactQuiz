package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true
  eludedGuards := rand.Intn(100)
  fmt.Println("eludedGuards has a value of", eludedGuards) 
  if eludedGuards > 50 {
  fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
} else {
  isHeistOn = false 
  fmt.Println("Plan a better disguise next time?")
}
openedVault := rand.Intn(100)
fmt.Println("Is the heist on????!!!!!", isHeistOn)
if openedVault > 70 && isHeistOn {
  fmt.Println("Grab and Go!")
} else if isHeistOn && openedVault < 70 {
  isHeistOn = false 
  fmt.Println("The vault cannot be opened :((((((")
}

leftSafely := rand.Intn(5)
if isHeistOn { 
  switch leftSafely {
  case 0: 
    isHeistOn = false 
    fmt.Print("Looks like you tripped an alarm... run?")
  case 1: 
    isHeistOn = false 
    fmt.Print("Turns out vault doors don't open from the inside...")   
  case 2: 
    isHeistOn = false 
    fmt.Print("When did they start raising dogs in vaults??") 
  case 3: 
    isHeistOn = false 
    fmt.Print("Looks like this fingerprint scanner won’t accept any fingerprint :O") 
  case 4: 
    isHeistOn = false 
    fmt.Print("Did I even pack the burlap bags?") 
  default:
    fmt.Println("Start the getaway car!")
  }

  if isHeistOn {
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Println("$", amtStolen, "not bad!")
  }


}


}
