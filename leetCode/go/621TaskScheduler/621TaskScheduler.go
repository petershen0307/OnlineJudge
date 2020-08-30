package taskscheduler

func leastInterval(tasks []byte, n int) int {
	taskIntervalCount := make(map[byte]int, 0)
	toDoTaskCount := make(map[byte]int, 0)
	for _, task := range tasks {
		taskIntervalCount[task] = 0
		toDoTaskCount[task]++
	}
	result := 0
	for len(toDoTaskCount) != 0 {
		doTask := byte(0)
		pickCount := 0
		for task, taskCount := range toDoTaskCount {
			// task 0 is idle
			if taskIntervalCount[task] == 0 {
				if pickCount < taskCount {
					doTask = task
					pickCount = taskCount
				}
			} else if taskIntervalCount[task] > 0 {
				taskIntervalCount[task]--
			}
		}
		if doTask != 0 {
			toDoTaskCount[doTask]--
			taskIntervalCount[doTask] = n
			if toDoTaskCount[doTask] == 0 {
				delete(toDoTaskCount, doTask)
			}
		}
		result++
	}
	return result
}
