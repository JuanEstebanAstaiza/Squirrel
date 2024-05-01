package Services

import (
	"github.com/JuanEstebanAstaiza/Squirrel/Models"
	"github.com/JuanEstebanAstaiza/Squirrel/Utils"
)
func AuthenticateUser (credentials Models.Credentials) (bool, error){

	// Buscar al usuario por email en la base de datos
	user, err := datab
}




func RegisterUser(user Models.Credentials) error {
	// Generar un ID único para el usuario
	userID, err := Utils.GenerateUserID()
	if err != nil {
		return err
	}

	// Encriptar la contraseña utilizando MD5
	encryptedPassword := Utils.EncryptPassword(user.Password)

	// Insertar el usuario en la base de datos con el ID único generado
	_, err = Utils.DB.Exec("INSERT INTO users (id, nickname, email, password) VALUES (?, ?, ?, ?)", userID, user.Nickname, user.Email, encryptedPassword)
	if err != nil {
		return err
	}

	return nil
}


func LoginUser(credentials Models.Credentials)(*Models.Credentials, error){
	// Verificar las credenciales del usuario

	authenticaded, err :=
}