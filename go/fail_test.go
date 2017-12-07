package service_test

import (
    "fmt"
    "os"
    "testing"
)

func FailTestGo(t *testing.T) {
    fmt.Printf("fail fail")
    os.Exit(3)
}