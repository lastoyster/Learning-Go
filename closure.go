/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/


// You can edit this code!
// Click here and start typing.
package main

import (
"fmt"
)

func makeWorld() func() int {
        count := 0
        return func() int{
            count += 3
            return count
        }
    }
    
    func main(){
        world1 := makeWorld()
        
        fmt.Println(world1())
        fmt.Println(world1())
        fmt.Println(world1())
    
        fmt.Println()
    
        fmt.Println(world1())
        fmt.Println(world1())
    }
    