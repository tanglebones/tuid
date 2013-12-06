package tuid_test

import (
  `github.com/tanglebones/tuid`
  `fmt`
)

func ExampleNew() {
  ta := tuid.New()
  tb := tuid.New() // will be different than ta but have a similar prefix when converted to string
  fmt.Printf("%v %v\n", ta, tb)
}

func ExampleZero() {
  ta := tuid.Zero // placeholder tuid
  fmt.Printf("%v\n", ta)
}

func ExampleParse() {
  tas := tuid.New().String()
  ta, err := tuid.Parse(tas)
  if err != nil {
    fmt.Printf("%v parsed to %v", tas, ta)
  } else {
    fmt.Printf("%v did not parse", tas)
  }
}
