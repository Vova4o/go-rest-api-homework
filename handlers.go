package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// getAllTasks возвращает список всех задач в формате JSON
// Пример запроса: GET /tasks
// Возвращает код 200 и список задач в формате JSON
// В случае ошибки возвращает код 500
func getAllTasks(w http.ResponseWriter, r *http.Request) {

	resp, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}

// createTask создает новую задачу
// Пример запроса: POST /tasks
// В теле запроса передается JSON с описанием задачи
// Возвращает код 201 в случае успеха
// В случае ошибки возвращает код 400
func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tasks[task.ID] = task

	w.WriteHeader(http.StatusCreated)
}

// getTaskByID возвращает задачу по ее ID
// Пример запроса: GET /tasks/1
// Возвращает код 200 и задачу в формате JSON
// В случае ошибки возвращает код 400
func getTaskByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	task, ok := tasks[id]
	if !ok {
		http.Error(w, "task not found", http.StatusBadRequest)
		return
	}

	resp, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// deleteTaskByID удаляет задачу по ее ID
// Пример запроса: DELETE /tasks/1
// Возвращает код 200 в случае успеха
// В случае ошибки возвращает код 400
func deleteTaskByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, ok := tasks[id]
	if !ok {
		http.Error(w, "task not found", http.StatusBadRequest)
		return
	}

	delete(tasks, id)

	w.WriteHeader(http.StatusOK)
}
