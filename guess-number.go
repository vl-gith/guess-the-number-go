package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	target := rand.Intn(100) + 1
	fmt.Println("Выбрано число от 1 до 100")

	count := 10
	playerWon := false

	reader := bufio.NewReader(os.Stdin)

	for count > 0 {

		fmt.Print("Введите число от 1 до 100: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal("Ошибка ввода: ", err)
		}

		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Ошибка! Введенное значение не является целым числом! Введите целое число!")
			continue
		}

		if guess < 1 || guess > 100 {
			continue
		}

		if guess == target {
			playerWon = true
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

	if !playerWon {
		fmt.Println("К сожалению, Вы проиграли!")
	} else {
		fmt.Println("Поздравляем, Вы выиграли!")
	}
}
