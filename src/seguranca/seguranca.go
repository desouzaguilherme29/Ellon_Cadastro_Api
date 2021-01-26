package seguranca

import (
	"golang.org/x/crypto/bcrypt"
)

//Hash gerar senha com hash
func Hash(senha string) ([]byte, error){
	return bcrypt.GenerateFromPassword([]byte(senha),bcrypt.DefaultCost)
}

//VerificarSenha compara as duas senhas
func VerificarSenha(senhaComHash string, senhaString string) error{
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}