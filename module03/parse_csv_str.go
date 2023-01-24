// // В привычных нам редакторах электронных таблиц присутствует удобное представление числа с разделителем разрядов в виде пробела, кроме того в России целая часть от дробной отделяется запятой. Набор таких чисел был экспортирован в формат CSV, где в качестве разделителя используется символ ";".

// // На стандартный ввод вы получаете 2 таких вещественных числа, в качестве результата требуется вывести частное от деления первого числа на второе с точностью до четырех знаков после "запятой" (на самом деле после точки, результат не требуется приводить к исходному формату).

// Sample Input:

// 1 149,6088607594936;1 179,0666666666666
// Sample Output:

// 0.9750


package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	err := scanner.Err()
	if err != nil {
		panic(err)
	}
    str = strings.ReplaceAll(str, " ", "")
	//fmt.Println(strings.Trim(str, " "))
	str = strings.ReplaceAll(str, ",", ".")

    strSplited := strings.Split(str, ";")
    
    num1, err := strconv.ParseFloat(strSplited[0], 64)
	if err != nil {
		panic(err)
	}
    num2, err := strconv.ParseFloat(strSplited[1], 64)
	if err != nil {
		panic(err)
	}
    fmt.Printf("%.4f", num1/num2)
}
