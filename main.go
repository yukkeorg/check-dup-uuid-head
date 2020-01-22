package main

import (
    "fmt"
    "github.com/google/uuid"
)


func main() {
    d := make(map[string]bool)
    c := 0
    q := 10_000_000
    for i := 0; i < q; i++ {
        guid, err := uuid.NewRandom()
        if err != nil {
            continue
        }

        key := guid.String()[0:7]

        if _, exists := d[key]; exists {
            c++
        } else {
            d[key] = true
        }
    }

    fmt.Printf("%d / %d = %f\n", c, q, float64(c)/float64(q))
}
