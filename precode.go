package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Task ...
type Task struct {
	ID           string   `json:"id"`
	Description  string   `json:"description"`
	Note         string   `json:"note"`
	Applications []string `json:"applications"`
}

var tasks = map[string]Task{
	"1": {
		ID:          "1",
		Description: "Сделать финальное задание темы REST API",
		Note:        "Если сегодня сделаю, то завтра будет свободный день. Ура!",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
		},
	},
	"2": {
		ID:          "2",
		Description: "Протестировать финальное задание с помощью Postmen",
		Note:        "Лучше это делать в процессе разработки, каждый раз, когда запускаешь сервер и проверяешь хендлер",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
			"Postman",
		},
	},
}

// Ниже напишите обработчики для каждого эндпоинта
// ...
/*
Обработчик должен вернуть все задачи, которые хранятся в мапе.
*/
func getAllTasks(w http.ResponseWriter, r *http.Request) {

	resp, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// в заголовок записываем тип контента, у нас это данные в формате JSON
	w.Header().Set("Content-Type", "application/json")
	// так как все успешно, то статус OK
	w.WriteHeader(http.StatusOK)
	// записываем сериализованные в JSON данные в тело ответа
	w.Write(resp)

}

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

/*
Обработчик для получения задачи по ID
Обработчик должен вернуть задачу с указанным в запросе пути ID, если такая есть в мапе.
В мапе ключами являются ID задач. Вспомните, как проверить, есть ли ключ в мапе. Если такого ID нет, верните соответствующий статус.
Конечная точка /tasks/{id}.
Метод GET.
При успешном выполнении запроса сервер должен вернуть статус 200 OK.
В случае ошибки или отсутствия задачи в мапе сервер должен вернуть статус 400 Bad Request.
*/
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

/*
Обработчик удаления задачи по ID
Обработчик должен удалить задачу из мапы по её ID. Здесь так же нужно сначала проверить, есть ли задача с таким ID в мапе, если нет вернуть соответствующий статус.
Конечная точка /tasks/{id}.
Метод DELETE.
При успешном выполнении запроса сервер должен вернуть статус 200 OK.
В случае ошибки или отсутствия задачи в мапе сервер должен вернуть статус 400 Bad Request.
Во всех обработчиках тип контента Content-Type — application/json.
*/
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

func main() {
	mux := chi.NewRouter()

	// здесь регистрируйте ваши обработчики
	// ...
	mux.Get("/tasks", getAllTasks)
	mux.Post("/tasks", createTask)
	mux.Get("/tasks/{id}", getTaskByID)
	mux.Delete("/tasks/{id}", deleteTaskByID)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
