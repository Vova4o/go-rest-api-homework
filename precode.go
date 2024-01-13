package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Регистрируем роутер который будет обрабатывать запросы
	mux := chi.NewRouter()

	// Регистрируем обработчики
	// Обратите внимание на то, что мы передаем функции, а не вызываем их
	mux.Get("/tasks", getAllTasks)
	mux.Post("/tasks", createTask)
	mux.Get("/tasks/{id}", getTaskByID)
	mux.Delete("/tasks/{id}", deleteTaskByID)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
