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

	playerWon := false
	levelChosen := false

	var count int
	var level int
	var strLvl string

	reader := bufio.NewReader(os.Stdin)

	for !levelChosen {
		fmt.Print("Введите желаемый уровень сложности, где: 1-легкий, 2-средний, 3-сложный: ")
		chooseLevel, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода!")
			continue
		}

		chooseLevel = strings.TrimSpace(chooseLevel)
		level, err = strconv.Atoi(chooseLevel)

		if err != nil {
			fmt.Println("Ошибка ввода, число не является целым!")
			continue
		}

		count, strLvl, err = chooseDifficultyLevel(level)

		if err != nil {
			fmt.Println("Ошибка ввода, выбранного Вами уровня не существует!")
			continue
		}

		levelChosen = true
		fmt.Println("Вы выбрали:", strLvl)
		fmt.Println("Вам доступно", count, "попыток")
	}

	for count > 0 {

		fmt.Print("Введите число от 1 до 100: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Ошибка ввода!")
			continue
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

	fmt.Println("Загаданное число:", target)
}

/*Функция выбора сложности игры, возвращает количество попыток, выбранный уровень и статус ошибки*/
/*Игрок будет выбирать сложность игры от 1 до 3, где 1 - легкий, 2 - средний, 3 - сложный*/
/*При этом, легкий уровень включает в себя 10 попыток; средний уровень 5 попыток, сложный - 3 попытки*/
func chooseDifficultyLevel(chooseLvl int) (int, string, error) {
	if chooseLvl == 1 {
		return 10, "легкий уровень", nil
	} else if chooseLvl == 2 {
		return 5, "средний уровень", nil
	} else if chooseLvl == 3 {
		return 3, "сложный уровень", nil
	} else {
		return 0, "", fmt.Errorf("Ошибка выбора уровня, уровня %d не существует!", chooseLvl)
	}
}
