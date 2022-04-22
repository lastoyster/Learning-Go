/*
 A bunch of ways to take user input
   1. fmt.Scanln(&text)
   2. fmt.Scanf("%s %d", &text, &num, ... )
   3. reader.ReadString()
   4. scanner.Scan()

 Here we use the Reader option.
    https://golang.org/pkg/bufio/#Reader****************************************************************************/

package main
import (
    "fmt"
    "bufio"
    "os"
    strs "strings"
    )

func main() {
    Println,Print := fmt.Println, fmt.Print
    Println("Hello Gophers!\n")
    
    reader := bufio.NewReader(os.Stdin)
    
    Print("Enter text:")
    line1,_ := reader.ReadString('\n')
    Print("\nEnter more text: ")
    line2, _:= reader.ReadString('\n')
    
    Trim := strs.TrimSpace
    line1 = Trim(line1)
    line2 = Trim(line2)
    fmt.Printf("\n\nText1:\"%v\"\nText2: \"%v\"\n",
    line1,
    line2)
}
