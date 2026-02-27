package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-go/internal/task"
)

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	var t task.Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
	}

	newTask := task.AddTask(t.Title, t.Content)
	json.NewEncoder(w).Encode(newTask)
}

func ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := task.ListTasks()
	json.NewEncoder(w).Encode(tasks)
}


func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	if task.MarkDone(id) {
		w.WriteHeader(http.StatusOK)
	}else {
		http.Error(w, "not found", http.StatusNotFound)
	}
}

func MarkDoneHandler(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, _ := strconv.Atoi(idStr)
    if task.MarkDone(id) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("done"))
    } else {
        http.Error(w, "not found", http.StatusNotFound)
    }
}