// На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является палиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.

// Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").

// Sample Input:

// топот
// Sample Output:

// Палиндром


package main 
import (
    "fmt"
)

func main() {
    var str string 
    
    fmt.Scan(&str)
    
    var palindrom int 
    
    strrune := []rune(str)
    
    strlen := len(strrune)
   
    for i, j := 0, strlen - 1; j >= 0; i, j = i + 1, j - 1 {
        if strrune[i] == strrune[j] {
                palindrom++
            }
    }
    
    if palindrom == strlen {
        fmt.Print("Палиндром")
    } else {
        fmt.Print("Не палиндром")
    }
}
