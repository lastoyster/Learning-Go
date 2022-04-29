package main

import . "fmt"


func main() {
   Supernatural := Supernatural { name:"Goblin", age:4000};
    Printf("%#v\n", Supernatural)
    name, age := Supernatural.destruct(); // 👈
    Printf("name = %q\nage = %d", name, age)
}


type Supernatural struct { name string; age int }


// 👇 destruct() method for Person struct
func (this Supernatural) destruct() (string, int) {
    return this.name, this.age;
    
}
