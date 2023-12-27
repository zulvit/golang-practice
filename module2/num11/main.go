package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Event представляет событие в календаре
type Event struct {
	ID    int       `json:"id"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}

// Response представляет стандартный ответ сервера
type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

// encodeJSON преобразует объект в JSON и записывает его в ResponseWriter
func encodeJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// Обработчик для создания события
func createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Здесь должна быть логика для обработки данных события
	encodeJSON(w, Response{Result: "Event created"}, http.StatusOK)
}

// Обработчик для обновления события
func updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Здесь должна быть логика для обновления данных события
	encodeJSON(w, Response{Result: "Event updated"}, http.StatusOK)
}

// Обработчик для удаления события
func deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Здесь должна быть логика для удаления события
	encodeJSON(w, Response{Result: "Event deleted"}, http.StatusOK)
}

// Обработчик для получения событий за день
func eventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Здесь должна быть логика для получения событий за день
	encodeJSON(w, Response{Result: "Events for day"}, http.StatusOK)
}

// Обработчик для получения событий за неделю
func eventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Здесь должна быть логика для получения событий за неделю
	encodeJSON(w, Response{Result: "Events for week"}, http.StatusOK)
}

// Обработчик для получения событий за месяц
func eventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Здесь должна быть логика для получения событий за месяц
	encodeJSON(w, Response{Result: "Events for month"}, http.StatusOK)
}

// Middleware для логирования
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/create_event", createEvent)
	http.HandleFunc("/update_event", updateEvent)
	http.HandleFunc("/delete_event", deleteEvent)
	http.HandleFunc("/events_for_day", eventsForDay)
	http.HandleFunc("/events_for_week", eventsForWeek)
	http.HandleFunc("/events_for_month", eventsForMonth)

	// Обертывание всех обработчиков middleware для логирования
	http.Handle("/", loggingMiddleware(http.DefaultServeMux))

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
