package Utils

import (
	"crypto/rand"
	"fmt"
)

// GenerateUserID genera un ID único para el usuario.
func GenerateUserID() (string, error) {
	// Generar un UUID (identificador único universal) v4
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return "", fmt.Errorf("error al generar UUID: %v", err)
	}
	// La versión 4 de UUID tiene los bits de la posición 6 y 7 establecidos en 01
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	// Formatear el UUID como una cadena de texto hexadecimal separada por guiones
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func GenerateSquireID() (string, error) {
	// Generar un UUID (identificador único universal) v4
	suid := make([]byte, 16)
	n, err := rand.Read(suid)
	if n != len(suid) || err != nil {
		return "", fmt.Errorf("error al generar UUID: %v", err)
	}
	// La versión 4 de UUID tiene los bits de la posición 6 y 7 establecidos en 01
	suid[6] = (suid[6] & 0x0f) | 0x40
	suid[8] = (suid[8] & 0x3f) | 0x80
	// Formatear el UUID como una cadena de texto hexadecimal separada por guiones
	return fmt.Sprintf("%x-%x-%x-%x-%x", suid[0:4], suid[4:6], suid[6:8], suid[8:10], suid[10:]), nil
}
