// Входные данные

// Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только арабские цифры.

// Выходные данные

// Выведите максимальную цифру, которая встречается во введенной строке.

// Sample Input:

// 1112221112
// Sample Output:

// 2

package main

import "fmt"

func main() {
    var num string
    fmt.Scan(&num)
    
    var r rune = 0
    for _, n := range num {
        if n > r {
            r = n
        }
    }
    fmt.Printf("%c",r)
}
