package Controllers

import (
	"encoding/json"
	"github.com/JuanEstebanAstaiza/Squirrel/Models"
	"github.com/JuanEstebanAstaiza/Squirrel/Services"
	"github.com/gorilla/mux"
	"net/http"
)

func AddSquire(w http.ResponseWriter, r *http.Request) {
	var squire Models.Squire
	err := json.NewDecoder(r.Body).Decode(&squire)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := Services.AddSquireToDB(squire)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func GetSquiresByUser(w http.ResponseWriter, r *http.Request) {
	// Obtener el userID de los parámetros de la URL
	vars := mux.Vars(r)
	userID := vars["userID"]

	// Obtener las contraseñas de la base de datos
	squires, err := Services.GetSquiresByUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Establecer el encabezado de respuesta
	w.Header().Set("Content-Type", "application/json")

	// Escribir la respuesta JSON
	if err = json.NewEncoder(w).Encode(squires); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func EditSquire(w http.ResponseWriter, r *http.Request) {
	var squire Models.Squire
	err := json.NewDecoder(r.Body).Decode(&squire)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := Services.EditSquire(squire)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !updated {
		http.Error(w, "Squire not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteSquire(w http.ResponseWriter, r *http.Request) {
	// Obtener el squireID de los parámetros de la URL
	vars := mux.Vars(r)
	squireID := vars["squireID"]

	deleted, err := Services.DeleteSquire(squireID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !deleted {
		http.Error(w, "Squire not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(deleted); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
