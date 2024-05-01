package Services

func RegisterUser(user) error {
	// Generar un ID único para el usuario
	userID, err := utils.GenerateUserID()
	if err != nil {
		return err
	}

	// Encriptar la contraseña utilizando MD5
	encryptedPassword := utils.EncryptPassword(user.Password)

	// Insertar el usuario en la base de datos con el ID único generado
	_, err = utils.DB.Exec("INSERT INTO users (id, nickname, email, password) VALUES (?, ?, ?, ?)", userID, user.Nickname, user.Email, encryptedPassword)
	if err != nil {
		return err
	}

	return nil
}
