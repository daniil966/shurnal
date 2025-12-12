package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	FIO    string
	Grades []int
	Avg    float64
}

func (s *Student) calculateAverage() {
	if len(s.Grades) == 0 {
		s.Avg = 0.0
		return
	}
	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	s.Avg = float64(sum) / float64(len(s.Grades))
}

func addStudent(students map[string]Student, reader *bufio.Reader) {
	fmt.Print("Введите фамилию и имя студента: ")
	fio, _ := reader.ReadString('\n')
	fio = strings.TrimSpace(fio)

	if _, exists := students[fio]; exists {
		fmt.Println("Студент с таким ФИО уже существует.")
		return
	}

	var grades []int
	fmt.Println("Введите оценки студента через пробел. Нажмите Enter еще раз для завершения.")
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			if len(grades) == 0 {
				fmt.Println("Оценки не были введены. Введите оценки или нажмите Enter еще раз для пропуска.")
				continue
			}
			break
		}

		gradesStrSlice := strings.Fields(input)
		for _, gradeStr := range gradesStrSlice {
			grade, err := strconv.Atoi(gradeStr)
			if err != nil {
				fmt.Printf("Некорректный ввод оценки '%s'. Пожалуйста, вводите только числа.\n", gradeStr)
				continue
			}
			if grade < 1 || grade > 5 {
				fmt.Printf("Оценка %d вне допустимого диапазона (1-5)\n", grade)
				continue
			}
			grades = append(grades, grade)
		}
	}

	student := Student{
		FIO:    fio,
		Grades: grades,
	}
	student.calculateAverage()

	students[fio] = student
	fmt.Println("Студент", fio, "успешно добавлен!")
}

func filterStudentsByAvg(students map[string]Student, threshold float64) []Student {
	filteredStudents := []Student{}
	for _, student := range students {
		if student.Avg < threshold {
			filteredStudents = append(filteredStudents, student)
		}
	}
	return filteredStudents
}

func printStudentInfo(student Student) {
	fmt.Printf("  ФИ: %s, Оценки: %v, Средний балл: %.2f\n", student.FIO, student.Grades, student.Avg)
}

func printAllStudents(students map[string]Student) {
	fmt.Println("Список всех студентов:")
	for _, student := range students {
		printStudentInfo(student)
	}
}

func main() {
	students := make(map[string]Student)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Добро пожаловать в журнал!")

	for {
		fmt.Print("\nВведите команду (help - для вывода всех команд): ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "add":
			addStudent(students, reader)
		case "list":
			if len(students) > 0 {
				printAllStudents(students)
			} else {
				fmt.Println("В базе данных пока нет студентов.")
			}
		case "filter":
			fmt.Print("Введите максимальный средний балл: ")
			thresholdStr, _ := reader.ReadString('\n')
			thresholdStr = strings.TrimSpace(thresholdStr)
			threshold, err := strconv.ParseFloat(thresholdStr, 64)
			if err != nil {
				fmt.Println("Некорректный ввод среднего балла. Пожалуйста, введите число.")
				continue
			}

			filteredStudents := filterStudentsByAvg(students, threshold)

			if len(filteredStudents) > 0 {
				fmt.Printf("Студенты со средним баллом ниже %.2f:\n", threshold)
				for _, student := range filteredStudents {
					printStudentInfo(student)
				}
			} else {
				fmt.Printf("Нет студентов со средним баллом ниже %.2f.\n", threshold)
			}
		case "help":
			fmt.Println("Доступные команды:")
			fmt.Println("  add    - Добавить нового студента.")
			fmt.Println("  list   - Вывести информацию о всех студентах.")
			fmt.Println("  filter - Отфильтровать студентов по среднему баллу (ниже заданного порога).")
			fmt.Println("  help   - Показать список доступных команд.")
			fmt.Println("  exit   - Выйти из программы.")
		case "exit":
			fmt.Println("Выход из программы. До свидания!")
			return
		default:
			fmt.Println("Неизвестная команда. Введите 'help' для просмотра списка команд.")
		}
	}
}
