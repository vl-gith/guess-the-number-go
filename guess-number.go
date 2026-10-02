package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	target := rand.Intn(100) + 1
	fmt.Println("Выбрано число от 1 до 100")

	playerWon := false
	levelChosen := false
	userNameChosen := false

	var count int
	var level int
	var strLvl string
	var userName string

	reader := bufio.NewReader(os.Stdin)

	for !userNameChosen {
		fmt.Print("Введите свое имя: ")
		result, err := getUserName(reader)

		if err != nil {
			fmt.Println(err)
			continue
		}
		userName = result
		userNameChosen = true
	}

	fmt.Println("Здравствуйте,", userName)

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

	countInit := count

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

		count--

		if guess == target {
			playerWon = true
			break
		} else {
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

	playerScore := calcResult(count, countInit, level, playerWon)

	fmt.Printf("Ваш результат: %0.3f", playerScore)
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

/*Функция расчета результата игрока, в зависимости от уровня сложности и использованных попыток, возвращает число с плавающей точкой*/
/*При заданном уровне сложности задается максимальный результат*/
func calcResult(count int, countInit int, level int, playerWon bool) float64 {

	if !playerWon {
		return 0.0
	}

	attemptsUsed := countInit - count
	var maxScore float64

	if level == 1 {
		maxScore = 1800
	} else if level == 2 {
		maxScore = 3000
	} else if level == 3 {
		maxScore = 7000
	}

	return maxScore / float64(attemptsUsed)
}

/*Функция запрашивает у пользователя имя и возвращает значение и статус ошибки*/
func getUserName(reader *bufio.Reader) (string, error) {

	userName, err := reader.ReadString('\n')

	if err != nil {
		return "", fmt.Errorf("Ошибка ввода: %w", err)
	}

	name := strings.TrimSpace(userName)

	if name == "" {
		return "", fmt.Errorf("Ошибка, введите имя!")
	}

	for _, ch := range name {
		if !unicode.IsLetter(ch) {
			return "", fmt.Errorf("Ошибка, в имени могут быть только буквы!")
		}
	}

	return name, nil
}
