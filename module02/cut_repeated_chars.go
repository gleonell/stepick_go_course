// Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку

// Sample Input:

// zaabcbd
// Sample Output:

// zcd


package main

import ("fmt"
        "strings")

func main() {
    var str string
    
    fmt.Scan(&str)
    
    for i := 0; i < len(str); i++ {
        if strings.Count(str, string(str[i])) == 1 {
            fmt.Printf("%c", str[i])
        }
    }
}