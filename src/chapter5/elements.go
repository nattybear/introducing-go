package main

import "fmt"

func main() {
  elements := map[string]string{
    "H"  : "Hydrogen",
    "He" : "Helium",
    "Li" : "Lithium",
    "Be" : "Beryllium",
    "B"  : "Boron",
    "C"  : "Carbon",
    "N"  : "Nitogen",
    "O"  : "Oxygen",
    "F"  : "Fluorine",
    "Ne" : "Neon",
  }

  if name, ok := elements["Un"]; ok {
    fmt.Println(name, ok)
  }
}
