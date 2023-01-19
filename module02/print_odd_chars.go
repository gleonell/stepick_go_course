// На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

// Sample Input:

// ihgewlqlkot
// Sample Output:

// hello

package main 

import "fmt"

func main() {
    var str string
    
    fmt.Scan(&str)
    
    for i, key := range str {
        if i % 2 != 0 {
            fmt.Printf("%c", key)
        }
    }
}