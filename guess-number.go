package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	target := rand.Intn(100) + 1
	fmt.Println("Выбрано число от 1 до 100")
	fmt.Println("Это число: ", target)

	count, n := 10, 10

	for i := 0; i < n; i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Введите число: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		guess, _ := strconv.Atoi(input)

		if guess == target {
			fmt.Println("Поздравляем, вы угадали число!")
			break
		} else {
			count--
			fmt.Println("К сожалению, вы не угадали число, попробуйте еще раз!", "У вас осталось", count, "попыток")
			if guess < target {
				fmt.Println("Введенное Вами число является МЕНЬШЕ, чем загаданное")
			} else {
				fmt.Println("Введенное вами число является БОЛЬШЕ, чем загаданное")
			}
		}
	}

	if count == 0 {
		fmt.Println("К сожалению, Вы проиграли!")
	} else {
		fmt.Println("Поздравляем, Вы выиграли!")
	}
}
