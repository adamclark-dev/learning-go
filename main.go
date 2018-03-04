package main

import (
    "test/model"
    "test/postcode"
    "fmt"
)

func main() {
    var test []Model.Vertex

    for i := 0; i < 10; i++ {
        test = append(test, Model.Vertex{i,i})
    }

    for _, value := range test {
        fmt.Println(value, value.Abs())
    }

    result,_ := postcode.Get("sk92bu")

    fmt.Println(result.Result.Country)
}
