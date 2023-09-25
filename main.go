package main

//TODO Добавить в англ версию фичу использования taskName
import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Print(chooselangNotification)
	fmt.Scanln(&lang)
	switch strings.ToLower(lang) {
	default:
		//Rus lang scenario started...
		fmt.Print(disclaimerRus)
		fmt.Println(TasksNumberNotificationRus)
		fmt.Scan(&countTasks)
		countTasks = int(countTasks)
		checkInput(countTasks, lang)
		tasks := make([]int, countTasks)          //массив задач
		fmt.Print(TaskNamesChoiceNotificationRus) //Используем наззвания задач?
		fmt.Scan(&decision)
		if !detectTrueFalseString(decision) {
			//Не используем названия задач
			fmt.Println(CountWorkersNotificationRus)
			fmt.Scan(&workers)
			checkInput(workers, lang)
			fmt.Print(NameDecisionNotifierRus)
			fmt.Scan(&decision)
			//заполняем массив с задачами
			for i := 0; i <= countTasks-1; i++ {
				tasks[i] = i
			}
			switch {
			case detectTrueFalseString(decision):
				//массив имен работников
				workerNames := make([]string, workers)
				//Записываем имена работников в массив
				for i := 0; i < workers; i++ {
					var name string
					fmt.Printf("Введите имя %d работника\n", i+1)
					fmt.Scan(&name)
					workerNames[i] = name
				}
				matchTaskIdToWorkerName(workers, tasks, workerNames, lang)
			case !detectTrueFalseString(decision):
				workerIds := make([]int, workers)
				//Записываем id работников в массив
				for i := 0; i < workers; i++ {
					workerIds[i] = i + 1
				}
				matchTaskIdToWorkerId(workers, tasks, workerIds, lang)
			}
		} else {
			//Используем названия задач
			taskNameArr := make([]string, countTasks+1)
			//заполняем массив названиями задач
			for i := 0; i < countTasks; i++ {
				var taskName string
				fmt.Printf("Введите имя %d задачи\n", i+1)
				fmt.Scan(&taskName)
				taskNameArr[i] = taskName
			}
			fmt.Println(CountWorkersNotificationRus)
			fmt.Scan(&workers)
			checkInput(workers, lang)
			fmt.Print(NameDecisionNotifierRus)
			fmt.Scan(&decision)
			switch {
			case detectTrueFalseString(decision):
				//массив имен работников
				workerNames := make([]string, workers)
				//Записываем имена работников в массив
				for i := 0; i < workers; i++ {
					var name string
					fmt.Printf("Введите имя %d работника\n", i+1)
					fmt.Scan(&name)
					workerNames[i] = name
				}
				matchTaskNameToWorkerName(workers, workerNames, taskNameArr, lang)
			case !detectTrueFalseString(decision):
				workerIds := make([]int, workers)
				//Записываем id работников в массив
				for i := 0; i < workers; i++ {
					workerIds[i] = i + 1
				}
				matchTaskIdToWorkerId(workers, tasks, workerIds, lang)
			}

		}
	case "eng":
		//Eng lang scenario started...
		fmt.Print(disclaimerEng)
		fmt.Println(TasksNumberNotificationEng)
		fmt.Scan(&countTasks)
		countTasks = int(countTasks)
		checkInput(countTasks, lang)
		tasks := make([]int, countTasks) //array of tasks
		fmt.Println(CountWorkersNotificationEng)
		fmt.Scan(&workers)
		checkInput(workers, lang)
		fmt.Println(NameDecisionNotifierEng)
		fmt.Scan(&decision)
		//write task ids to array
		for i := 0; i <= countTasks-1; i++ {
			tasks[i] = i
		}
		switch {
		case detectTrueFalseString(decision):
			//array of workers name
			workerNames := make([]string, workers)
			//write worker's names to array
			for i := 0; i < workers; i++ {
				var name string
				fmt.Printf("Type name of worker №%d\n", i+1)
				fmt.Scan(&name)
				workerNames[i] = name
			}
			matchTaskIdToWorkerName(workers, tasks, workerNames, lang)
		case !detectTrueFalseString(decision):
			workerIds := make([]int, workers)
			//write worker's ids to array
			for i := 0; i < workers; i++ {
				workerIds[i] = i + 1
			}
			matchTaskIdToWorkerId(workers, tasks, workerIds, lang)
		}
	}

}
