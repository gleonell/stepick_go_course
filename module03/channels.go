// Напишите функцию которая принимает канал и число N типа int. Необходимо вернуть значение N+1 в канал. 
// Функция должна называться task().

// Внимание! Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!

// Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!

func task(c chan int, N int) {
    c <- N + 1
}


// Напишите функцию которая принимает канал и строку. Необходимо отправить эту же строку в заданный канал 5 раз, добавив к ней пробел.

// Функция должна называться task2().

// Внимание! Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!

// Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!

func task2(c chan string, s string) {
    for i := 0; i < 5; i++ {
        c <- s + " " 
    }
}
