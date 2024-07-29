package api

import (
	"encoding/json"
	"messaggioService/db"
	"messaggioService/kafka"
	"messaggioService/models"
	"net/http"
)

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.DB.Create(&message).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := kafka.SendMessage(message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Message created successfully"}`))
}

func GetStats(w http.ResponseWriter, r *http.Request) {
	var count int64
	if err := db.DB.Model(&models.Message{}).Where("processed = ?", true).Count(&count).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]int64{"processed_messages": count})
}
