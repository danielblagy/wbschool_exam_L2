package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

type Event struct {
	EventId int       `json:"event_id"`
	Content string    `json: "content"`
	Date    time.Time `json:"date"`
}

var (
	events      []Event
	eventsMutex *sync.Mutex
)

type Logger struct {
	handler http.Handler
}

// adding logging
func (l *Logger) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	initTime := time.Now()
	l.handler.ServeHTTP(w, r)
	endTime := time.Now()
	difference := endTime.Sub(initTime)
	log.Printf("%s %s %v", r.Method, r.URL.Path, difference)
}

func WrapHandler(h http.Handler) *Logger {
	var wrapLoger Logger
	wrapLoger.handler = h
	return &wrapLoger
}

func parseJSON(r *http.Request) (Event, error) {
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		return event, errors.New("invalid json file")
	}
	return event, nil
}

func errorResponse(w http.ResponseWriter, e string, status int) {
	errorResponse := struct {
		Error string `json:"error"`
	}{Error: e}
	json, err := json.Marshal(errorResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func resultResponse(w http.ResponseWriter, r string, e []Event, status int) {
	resultResponse := struct {
		Result string  `json:"result"`
		Events []Event `json:"events"`
	}{Result: r, Events: e}

	json, err := json.Marshal(resultResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func isValidEvent(event Event) bool {
	return !(event.EventId <= 0 || event.Content == "")
}

func CreateNewEvent(event Event) error {
	eventsMutex.Lock()
	defer eventsMutex.Unlock()
	for _, x := range events {
		if x.EventId == event.EventId {
			return errors.New("event already exists")
		}
	}
	events = append(events, event)
	return nil
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, "must be POST", http.StatusBadRequest)
		return
	}
	newEvent, err := parseJSON(r)
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if isValidEvent(newEvent) {
		errorResponse(w, err.Error(), http.StatusBadRequest)
	}

	if err := CreateNewEvent(newEvent); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
	}
	resultResponse(w, "Success", []Event{newEvent}, http.StatusCreated)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, "must be POST", http.StatusBadRequest)
		return
	}
	newEvent, err := parseJSON(r)
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if isValidEvent(newEvent) {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	eventsMutex.Lock()
	defer eventsMutex.Unlock()
	for i, x := range events {
		if x.EventId == newEvent.EventId {
			events[i] = newEvent
			return
		}
	}
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, "must be POST", http.StatusBadRequest)
		return
	}
	newEvent, err := parseJSON(r)
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if isValidEvent(newEvent) {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	eventsMutex.Lock()
	defer eventsMutex.Unlock()
	for i, x := range events {
		if x.EventId == newEvent.EventId {
			events = append(events[:i], events[i+1:]...)
		}
	}
}

func EventsByDay(date time.Time) []Event {
	var result []Event
	eventsMutex.Lock()
	defer eventsMutex.Unlock()
	for _, x := range events {
		if x.Date.Year() == date.Year() && x.Date.Month() == date.Month() && x.Date.Day() == date.Day() {
			result = append(result, x)
		}
	}
	return result
}

func EventsByWeek(date time.Time) []Event {
	var result []Event
	eventsMutex.Lock()
	defer eventsMutex.Unlock()
	for _, x := range events {
		difference := date.Sub(x.Date)
		if difference < 0 {
			difference = -difference
		}
		if difference <= time.Duration(7*24)*time.Hour {
			result = append(result, x)
		}
	}
	return result

}

func EventsByMonth(date time.Time) []Event {
	var result []Event
	eventsMutex.Lock()
	defer eventsMutex.Unlock()
	for _, x := range events {
		if x.Date.Year() == date.Year() && x.Date.Month() == date.Month() {
			result = append(result, x)
		}
	}
	return result

}

func EventsDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorResponse(w, "must be GET", http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2019-09-09", r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := EventsByDay(date)
	resultResponse(w, "Success", result, http.StatusOK)
}

func EventsWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorResponse(w, "must be GET", http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2019-09-09", r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := EventsByWeek(date)
	resultResponse(w, "Success", result, http.StatusOK)
}
func EventsMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorResponse(w, "must be GET", http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2019-09-09", r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := EventsByMonth(date)
	resultResponse(w, "Success", result, http.StatusOK)
}

func main() {
	port := ":8080"

	httpMultyPlex := http.NewServeMux()
	httpMultyPlex.HandleFunc("/create_event", CreateEvent)
	httpMultyPlex.HandleFunc("/update_event", UpdateEvent)
	httpMultyPlex.HandleFunc("/delete_event", DeleteEvent)
	httpMultyPlex.HandleFunc("/events_for_day", EventsDay)
	httpMultyPlex.HandleFunc("/events_for_week", EventsWeek)
	httpMultyPlex.HandleFunc("/events_for_month", EventsMonth)

	midleLogger := WrapHandler(httpMultyPlex)
	log.Fatalln(http.ListenAndServe(port, midleLogger.handler))
}
