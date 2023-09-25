package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func checkInput(num int, lang string) {
	switch lang {
	default:
		if num < 1 {
			fmt.Println(InputZeroTaskRus)
			time.Sleep(2 * time.Second)
			os.Exit(0)
		}
	case "eng":
		if num < 1 {
			fmt.Println(InputZeroTaskEng)
			time.Sleep(2 * time.Second)
			os.Exit(0)
		}
	}

}

// detection user's decisions from string and convert to bool, p.s. if unrecognized then false
func detectTrueFalseString(arg string) bool {
	arg = strings.ToLower(arg)
	switch arg {
	case "yes":
		return true
	case "y":
		return true
	case "ye":
		return true
	case "да":
		return true
	case "д":
		return true
	case "1":
		return true
	default:
		return false
	}
}

func matchTaskNameToWorkerName(countWorkets int, names, tasksNames []string, lang string) {
	tasksWithWorkers := make(map[string]string)
	var decision string
	switch lang {
	default:
		fmt.Println(TaskWorkerTableRus)
		for i := 0; i < len(tasksNames)-1; i++ {
			//fmt.Printf("len(tasksNames) = %v\n", len(tasksNames))
			tasksWithWorkers[tasksNames[i]] = names[rand.Intn(countWorkets)]
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("|  %v   |  %v\n", tasksNames[i], tasksWithWorkers[tasksNames[i]])
		}
		fmt.Print(AgreeOfDistributionRus)
		fmt.Scan(&decision)
		switch {
		case detectTrueFalseString(decision):
			fmt.Println(AppEndingRus)
			time.Sleep(2 * time.Second)
			os.Exit(0)
		case !detectTrueFalseString(decision):
			fmt.Println(RepeatingAppRus)
			time.Sleep(1 * time.Second)
			matchTaskNameToWorkerName(countWorkets, names, tasksNames, lang)
		}
	}
}

// Recursionly matching task ids with workers names
func matchTaskIdToWorkerName(countWorkets int, tasks []int, names []string, lang string) {
	tasksWithWorkers := make(map[int]string)
	var decision string
	switch lang {
	default:
		fmt.Println(TaskWorkerTableRus)
		for i := 0; i < len(tasks); i++ {
			secIndex := rand.Intn(countWorkets)
			tasksWithWorkers[i] = names[secIndex]
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("|  %v   |  %v\n", i+1, tasksWithWorkers[i])
		}
		fmt.Print(AgreeOfDistributionRus)
		switch {
		case detectTrueFalseString(decision):
			fmt.Println(AppEndingRus)
			time.Sleep(2 * time.Second)
			os.Exit(0)
		case !detectTrueFalseString(decision):
			fmt.Println(RepeatingAppRus)
			time.Sleep(1 * time.Second)
			matchTaskIdToWorkerName(countWorkets, tasks, names, lang)
		}

	case "eng":
		fmt.Println(TaskWorkerTableEng)
		for i := 0; i < len(tasks); i++ {
			secIndex := rand.Intn(countWorkets)
			tasksWithWorkers[i] = names[secIndex]
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("|  %v   |  %v\n", i+1, tasksWithWorkers[i])
		}
		fmt.Print(AgreeOfDistributionEng)
		fmt.Scan(&decision)
		switch {
		case detectTrueFalseString(decision):
			fmt.Println(AppEndingEng)
			time.Sleep(2 * time.Second)
			os.Exit(0)
		case !detectTrueFalseString(decision):
			fmt.Println(RepeatingAppEng)
			time.Sleep(1 * time.Second)
			matchTaskIdToWorkerName(countWorkets, tasks, names, lang)
		}
	}

}

// Recursionly matching task ids with worker ids
func matchTaskIdToWorkerId(countWorkets int, tasks, workerIds []int, lang string) {
	tasksWithWorkers := make(map[int]int)
	var decision string
	switch lang {
	default:
		fmt.Println(TaskWorkerTableRus)
		for i := 0; i < len(tasks); i++ {
			secIndex := rand.Intn(countWorkets)
			tasksWithWorkers[i] = workerIds[secIndex]
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("|  %v   |  %v\n", i+1, tasksWithWorkers[i])
		}
		fmt.Print(AgreeOfDistributionRus)

		fmt.Scan(&decision)
		switch {
		case detectTrueFalseString(decision):
			fmt.Println(AppEndingRus)
			time.Sleep(1 * time.Second)
			os.Exit(0)
		case !detectTrueFalseString(decision):
			fmt.Println(RepeatingAppRus)
			time.Sleep(2 * time.Second)
			matchTaskIdToWorkerId(countWorkets, tasks, workerIds, lang)
		}
	case "eng":
		fmt.Println(TaskWorkerTableEng)
		for i := 0; i < len(tasks); i++ {
			secIndex := rand.Intn(countWorkets)
			tasksWithWorkers[i] = workerIds[secIndex]
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("|  %v   |  %v\n", i+1, tasksWithWorkers[i])
		}
		fmt.Print(AgreeOfDistributionEng)

		fmt.Scan(&decision)
		switch {
		case detectTrueFalseString(decision):
			fmt.Println(AppEndingEng)
			time.Sleep(1 * time.Second)
			os.Exit(0)
		case !detectTrueFalseString(decision):
			fmt.Println(RepeatingAppEng)
			time.Sleep(2 * time.Second)
			matchTaskIdToWorkerId(countWorkets, tasks, workerIds, lang)
		}
	}

}
