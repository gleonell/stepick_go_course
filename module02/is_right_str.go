// На вход подается строка. Нужно определить, является ли она правильной или нет. Правильная строка начинается с заглавной буквы и заканчивается точкой. Если строка правильная - вывести Right иначе - вывести Wrong

// Sample Input:

// Быть или не быть.
// Sample Output:

// Right

package main

import (
    "fmt"
    "strings"
    "os"
    "bufio"
    "unicode"
)

func main() {
    text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    str := []rune(text)
    if unicode.IsUpper(str[0]) == true &&  strings.HasSuffix(text, ".") {
        fmt.Print("Right")
    } else {
        fmt.Print("Wrong")
    }
}
