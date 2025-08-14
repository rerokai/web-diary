package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type Task struct {
	Title string
	Done  bool
}

type Habit struct {
	Title string
	Done  bool
}

type TodayData struct {
	Date         string
	Tasks        []Task
	Notes        string
	Habits       []Habit
	Productivity int
}

func TodayHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("internal/templates/today.html")
	if err != nil {
		log.Fatal(err) // если шаблон не найден, приложение упадет
	}

	tasks := []Task{
		{Title: "Buy groceries", Done: true},
		{Title: "Read book", Done: false},
	}

	habits := []Habit{
		{Title: "Drink water", Done: true},
		{Title: "Meditate", Done: false},
	}

	data := TodayData{
		Date:         "2025-08-14",
		Tasks:        tasks,
		Notes:        "Here my notes for today, Im so tired...",
		Habits:       habits,
		Productivity: 60,
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}
