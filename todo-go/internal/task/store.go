package task 

var tasks  []Task
var nexID int = 1

func AddTask (title, content string) Task {
	t := Task{
		ID: nexID,
		Title: title,
		Content: content,
		Done: false,
	}
	tasks = append(tasks, t)
	nexID++
	return t
}

func ListTasks() []Task {
	return tasks
}

func DeleteTask(id int) bool {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}

func MarkDone(id int) bool {
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Done = true
			return true
		}
	}
	return false
}