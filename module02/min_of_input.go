// Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

// Входные данные

// Вводится четыре числа.

// Выходные данные

// Необходимо вернуть из функции наименьшее из 4-х данных чисел

// Sample Input:

// 4 5 6 7
// Sample Output:

// 4

func minimumFromFour() int {
    var a, tmp int
    var flag bool = false
   
    for i := 0; i < 4; i++ {
        fmt.Scan(&a)
        if a < tmp || !flag {
            tmp = a
            flag = true
        }
    }
    
    return tmp
}
