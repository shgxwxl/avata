package main
import (
    "flag"
    "frame"
)
var s frame.Server
func main() {
    flag.Parse()
    err := s.Init()
    if err != nil {
        panic(err)
    }
    s.Run()
}