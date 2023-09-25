package main

const (
	disclaimerRus                  = "ДИСКЛЕЙМЕР!\nДанная программа производит автоматическое распределение задач между исполнителями.\nИспользование алгоритма данной программы не гарантирует честное и справедливое распределение задач между исполнителями.\nБудьте добры к коллегам и не злоупотребляйте везением.\nУдачи!\n\n\n"
	disclaimerEng                  = "DISCLAIMER!\nThis program automatically assign tasks between performers.\nUsing the algorithm of this program does not guarantee a fair and equitable distribution of tasks between performers.\nBe kind to your colleagues and do not abuse your luck.\nGood luck!\n\n\n"
	chooselangNotification         = "Choose language | Выбор языка\nНапишите \"рус\" для использования программы на русском языке.\n Type \"eng\" for using english language.\n"
	TasksNumberNotificationRus     = "Введите количество задач, которое необходимо распланировать: "
	TasksNumberNotificationEng     = "Type the number of tasks to be scheduled: "
	TaskWorkerTableRus             = "|Задача| Исполнитель"
	TaskWorkerTableEng             = "| Task | Worker"
	AppEndingRus                   = "Программа завершена..."
	AppEndingEng                   = "The application ending..."
	CountWorkersNotificationRus    = "Введите количество работников, между которыми нужно распланировать задачи: "
	CountWorkersNotificationEng    = "Type the number of workers between whom you want to schedule tasks: "
	NameDecisionNotifierRus        = "Если желаете использовать имена работников при распределении, отправьте \"Да\"\nЕсли желаете использовать только id работников пришлите \"Нет\":"
	NameDecisionNotifierEng        = "If you want to use employee names for distribution, send \"Yes\"\nIf you want to use only employee ids, send \"No\":"
	AgreeOfDistributionRus         = "\nВас устраивает распределение задач?\n(Введите \"Да\" если устраивает или \"Нет\" если нужно повторить процедуру распределения задач)\n"
	AgreeOfDistributionEng         = "\nAre you satisfied with the task distribution?\n(Enter \"Yes\" if you are satisfied or \"No\" if you need to repeat the task distribution procedure)\n"
	RepeatingAppRus                = "Инициировано повторное распределение задач..."
	RepeatingAppEng                = "Redistribution of tasks has been initiated..."
	InputZeroTaskRus               = "Невозможно продолжение со значением меньше единицы. Завершение программы."
	InputZeroTaskEng               = "Cannot continue with less one value. Terminating programm."
	TaskNamesChoiceNotificationRus = "Желаете ли использовать названия задач вместо порядковых номеров?\nВведите \"Да\", чтобы использовать названия задач и \"Нет\" чтобы использовать порядковые номера задач (этот метод быстрее).\n"
	TaskNamesChoiceNotificationEng = "Would you like to use task names instead of serial numbers?\nEnter \"Yes\" to use task names and \"No\" to use task serial numbers (this method is faster).\n"
)

var (
	lang       string // user's language
	countTasks int    //number of tasks
	workers    int    //number of workers
	decision   string //user's decisions
)
