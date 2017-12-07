package service_test

import (
    "fmt"
    "os"
    "testing"
)

func TestFailGo(t *testing.T) {
    fmt.Printf("fail fail")
    os.Exit(3)
}