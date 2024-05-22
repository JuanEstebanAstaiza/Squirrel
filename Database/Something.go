package Database

import (
	"database/sql"
	"errors"
	"github.com/JuanEstebanAstaiza/Squirrel/Models"
	"github.com/JuanEstebanAstaiza/Squirrel/Utils"
)

func FindUserByEmail(email string) (*Models.Profile, error) {
	// Consulta a la base de datos para buscar al usuario por su email.
	query := "SELECT id, nickname, email FROM users WHERE email = ?"
	row := Utils.DB.QueryRow(query, email)

	// Escanear el resultado de la consulta en una estructura de usuario.
	var user Models.Profile
	err := row.Scan(&user.ID, &user.Nickname, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Si no se encontró ningún usuario con el email dado, devolvemos nil.
			return nil, nil
		}
		// Si hubo otro error durante la consulta, lo devolvemos.
		return nil, err
	}

	// Si se encontró un usuario, lo devolvemos.
	return &user, nil
}

func FindSquiresByUserId(user_id string) (*[]Models.Squire, error) {
	// Consulta a la base de datos para buscar al usuario por su email.
	query := "SELECT * FROM mainpage WHERE user_id = ?"
	rows, err := Utils.DB.Query(query, user_id)

	if err != nil {
		return nil, err
	}

	// Escanear el resultado de la consulta en una estructura de usuario.
	var squires []Models.Squire
	for rows.Next() {
		var squire Models.Squire
		if err := rows.Scan(&squire.ID, &squire.UserId, &squire.Url, &squire.Username, &squire.Password); err != nil {
			return nil, err
		}
		squires = append(squires, squire)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &squires, nil
}
