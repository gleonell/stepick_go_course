// Напишите программу, принимающая на вход число 
// N(N≥4), а затем 
// N целых чисел-элементов среза. На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.

// Sample Input:

// 5
// 41 -231 24 49 6
// Sample Output:

// 49

package main
import "fmt"

func main() {
    var len int
    
    fmt.Scan(&len)
    
    var array = make([]int, len)
    
    for i := 0; i < len; i++ {
        var k int
        fmt.Scan(&k)
        array[i] = k
    }

    fmt.Print(array[3])
}