package main

import "fmt"

type person struct {
    name string
    age  int
}

type student struct {
    p person
    grade int
    sex  int
}

func newPerson(name string) *person {

    p := person{name: name}
    p.age = 42
    return &p
}


func main() {

    fmt.Println(person{"Bob", 20})

    fmt.Println(person{name: "Alice", age: 30})

    fmt.Println(person{name: "Fred"})

    fmt.Println(&person{name: "Ann", age: 40})

    fmt.Println(newPerson("Jon"))

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)

    stu1 := student { person{"bob", 20}, 1, 0 }

    fmt.Println(stu1)

}