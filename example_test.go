package tuid_test

import (
  `fmt`
  `github.com/tanglebones/tuid`
)

func ExampleTuidProvider() {
  tp := tuid.NewTuidProvider(tuid.DefaultResolver)
  ta := tp.New()
  tb := tp.New() // will be different than ta but have a similar prefix when converted to string
  fmt.Printf("%v %v\n", ta, tb)
}

func Example_zero() {
  ta := tuid.Zero // placeholder tuid
  fmt.Printf("%v", ta)
  // Output:
  // AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
}

func ExampleParse() {
  tp := tuid.NewTuidProvider(tuid.DefaultResolver)
  tas := tp.New().String()
  ta, err := tuid.Parse(tas)
  if err != nil {
    fmt.Printf("%v parsed to %v", tas, ta)
  } else {
    fmt.Printf("%v did not parse", tas)
  }
}
