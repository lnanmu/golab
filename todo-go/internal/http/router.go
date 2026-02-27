package http 

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/tasks/add", AddTaskHandler)
	http.HandleFunc("/tasks/list", ListTasksHandler)
	http.HandleFunc("/tasks/delete",DeleteTaskHandler)
	http.HandleFunc("/tasks/done",MarkDoneHandler)
}