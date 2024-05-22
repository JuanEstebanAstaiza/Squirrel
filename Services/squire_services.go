package Services

import (
	"github.com/JuanEstebanAstaiza/Squirrel/Models"
	"github.com/JuanEstebanAstaiza/Squirrel/Utils"
)

//Insertar datos en la base de datos

func AddSquireToDB(squire Models.Squire) (string, error) {
	query := "INSERT INTO mainpage (id, user_id, url,User, password) VALUES (?, ?, ?, ?, ?)"

	squireID, err := Utils.GenerateSquireID()
	if err != nil {
		return "", err
	}

	// Ejecutar la consulta SQL con los valores proporcionados
	_, err = Utils.DB.Exec(query, squireID, squire.UserId, squire.Url, squire.Username, squire.Password)
	if err != nil {
		return "", err
	}

	return squireID, nil
}

func GetSquire() ([]Models.Squire, error) {
	query := "SELECT url, user, password FROM mainpage"

	// Ejecutar la consulta SQL
	rows, err := Utils.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterar sobre las filas y guardar los datos en una lista
	var passwords []Models.Squire
	for rows.Next() {
		var password Models.Squire
		err := rows.Scan(&password.Url, &password.Username, &password.Password)
		if err != nil {
			return nil, err
		}
		passwords = append(passwords, password)
	}

	return passwords, nil
}
