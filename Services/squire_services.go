package Services

import (
	"fmt"
	"github.com/JuanEstebanAstaiza/Squirrel/Models"
	"github.com/JuanEstebanAstaiza/Squirrel/Utils"
)

// Function to add a squire to the database
func AddSquireToDB(squire Models.Squire) (string, error) {
	query := "INSERT INTO mainpage (id, user_id, url, user_name, password) VALUES (?, ?, ?, ?, ?)"

	squireID, err := Utils.GenerateSquireID()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Execute the SQL query with the provided values
	_, err = Utils.DB.Exec(query, squireID, squire.UserId, squire.Url, squire.Username, squire.Password)
	if err != nil {
		return "", err
	}

	return squireID, nil
}

// Function to get squires by user
func GetSquiresByUser(userID string) ([]Models.Squire, error) {
	query := "SELECT id, url, user_name, password FROM mainpage WHERE user_id=?"

	// Execute the SQL query
	rows, err := Utils.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and save the data in a list
	var squires []Models.Squire
	for rows.Next() {
		var squire Models.Squire
		err := rows.Scan(&squire.ID, &squire.Url, &squire.Username, &squire.Password)
		if err != nil {
			return nil, err
		}
		squires = append(squires, squire)
	}
	fmt.Println(squires)
	return squires, nil
}

// Function to edit a squire
func EditSquire(squire Models.Squire) (bool, error) {
	exists, err := squireExists(squire.ID)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, nil
	}

	query := "UPDATE mainpage SET url=?, user_name=?, password=? WHERE id=?"

	// Execute the SQL query
	_, err = Utils.DB.Exec(query, squire.Url, squire.Username, squire.Password, squire.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Function to check if a squire exists
func squireExists(ID string) (bool, error) {
	query := "SELECT COUNT(*) FROM mainpage WHERE id=?"

	rows, err := Utils.DB.Query(query, ID)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return false, err
		}
	}
	return count != 0, nil
}

// Function to delete a squire
func DeleteSquire(squireID string) (bool, error) {
	exists, err := squireExists(squireID)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, nil
	}

	query := "DELETE FROM mainpage WHERE id=?"

	// Execute the SQL query
	_, err = Utils.DB.Exec(query, squireID)
	if err != nil {
		return false, err
	}
	return true, nil
}
