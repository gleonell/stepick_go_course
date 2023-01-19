// Дана строка, содержащая только английские буквы (большие и маленькие). Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).

// Входные данные

// Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков.

// Выходные данные

// Вывести строку, которая получится после добавления символов '*'.

// Sample Input:

// LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO
// Sample Output:

// L*I*t*B*e*o*F*L*c*S*G*B*O*F*Q*x*M*H*o*I*u*D*D*W*c*q*c*V*g*k*c*R*o*A*e*o*c*X*O
// Sample Input:

// LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO
// Sample Output:

// L*I*t*B*e*o*F*L*c*S*G*B*O*F*Q*x*M*H*o*I*u*D*D*W*c*q*c*V*g*k*c*R*o*A*e*o*c*X*O


package main 

import "fmt"

func main(){
    var str string
    fmt.Scan(&str)
    len := len(str)
    for i := 0; i < len; i++ {
        if i == len - 1{
            fmt.Printf("%c", str[i])
        } else {
            fmt.Printf("%c*", str[i])
        }
    }
}
