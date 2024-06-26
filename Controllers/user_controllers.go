package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/JuanEstebanAstaiza/Squirrel/Models"
	"github.com/JuanEstebanAstaiza/Squirrel/Services"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user Models.Credentials
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	if err != nil {
		http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	// Registrar al usuario
	id, err := Services.RegisterUser(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al registrar el usuario: %e", err), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(id); err != nil {
		http.Error(w, fmt.Sprintf("Error al registrar el usuario: %e", err), http.StatusInternalServerError)
	}

	// Responder con éxito
	w.WriteHeader(http.StatusCreated)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Parsear los datos del cuerpo de la solicitud
	var user Models.Credentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtener el usuario con las credenciales proporcionadas
	userInfo, err := Services.LoginUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if userInfo == nil {
		// Las credenciales son inválidas
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	// Si las credenciales son válidas, enviar la información del usuario
	if userInfo != nil {
		var userProfile Models.Profile
		userProfile.ID = userInfo.ID
		userProfile.Nickname = userInfo.Nickname
		userProfile.Email = userInfo.Email
		// Escribir encabezado de respuesta exitosa
		w.WriteHeader(http.StatusOK)
		// Codificar el mensaje de "Acceso concedido" junto con los detalles del usuario
		response := map[string]interface{}{
			"message":  "Acceso concedido",
			"userInfo": userProfile,
		}
		// Codificar la respuesta en formato JSON y enviarla
		json.NewEncoder(w).Encode(response)
	}

}
