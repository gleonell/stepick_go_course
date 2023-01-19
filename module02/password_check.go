// Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. 
// Длина пароля должна быть не менее 5 символов, он должен содержать только арабские цифры и буквы латинского алфавита. 
// На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

// Sample Input:

// fdsghdfgjsdDD1
// Sample Output:

// Ok

package main

import (
    "fmt"
    "unicode"
)

func main() {
    var pswd string
    
    fmt.Scan(&pswd)
    
    flag := true
    
    for _, k := range pswd {
        if unicode.IsDigit(k) || (unicode.IsLetter(k) == true && unicode.Is(unicode.Latin, k) == true) {
            continue
        } else {
            flag = false
        }
    }
    
    if flag == true && len(pswd) >= 5 {
        fmt.Print("Ok")
    } else {
        fmt.Print("Wrong password")
    }
}
