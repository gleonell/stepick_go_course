// Из натурального числа удалить заданную цифру.

// Входные данные

// Вводятся натуральное число и цифра, которую нужно удалить.

// Выходные данные

// Вывести число без заданных цифр.

// Sample Input:

// 38012732
// 3
// Sample Output:

// 801272

package main

import "fmt"

func main() {
    var num, del int
    
    fmt.Scan(&num, &del)
    
    str := fmt.Sprintf("%v", num)
    
    for i := range str {
        if int(str[i]) - '0' != del {
            fmt.Print(int(str[i]) - '0' )
        }
    }
}
