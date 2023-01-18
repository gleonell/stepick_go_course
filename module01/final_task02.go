// Дано трехзначное число. Переверните его, а затем выведите. 

// Формат входных данных
// На вход дается трехзначное число, не оканчивающееся на ноль.

// Формат выходных данных
// Выведите перевернутое число.

// Sample Input:

// 843
// Sample Output:

// 348

package main 

import "fmt"

func main() {
    var num int
    
    fmt.Scan(&num)
    rev_num := ((num % 10) * 100) + (((num % 100) / 10) * 10) + num / 100
    fmt.Print(rev_num)
}