// Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.

 

// Формат входных данных

// На вход подается номер билета - одно шестизначное  число.

// Формат выходных данных
// Выведите "YES", если билет счастливый, в противном случае - "NO".

// Sample Input:

// 613244
// Sample Output:

// YES

package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    
    sum1 := ((n / 1000) % 10) + ((n / 10000) % 10) + (n / 100000)
    sum2 := ((n % 1000) % 10) + (((n % 1000) / 10) % 10) + (n % 1000) / 100
    
    if sum1 == sum2 {
        fmt.Print("YES")
    } else {
        fmt.Print("NO")
    }
}