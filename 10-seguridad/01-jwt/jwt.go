// Generación y validación de JWT. El token viaja en el header Authorization
// como "Bearer <token>" y el middleware lo verifica en cada request protegido.

package main

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var claveSecreta = []byte("go-api-secret-2026-no-hardcodear-en-prod")

type Claims struct {
	UsuarioID string `json:"usuario_id"`
	Email     string `json:"email"`
	Rol       string `json:"rol"`
	// jwt.RegisteredClaims incluye los campos estándar: ExpiresAt, IssuedAt, etc.
	jwt.RegisteredClaims
}

func GenerarToken(usuarioID, email, rol string) (string, error) {
	claims := Claims{
		UsuarioID: usuarioID,
		Email:     email,
		Rol:       rol,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-learning-api",
		},
	}
	// HS256 firma con una clave simétrica. En producción se usa RS256 con clave privada/pública.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(claveSecreta)
}

func ValidarToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
		// Verificamos que el algoritmo sea el esperado. Evita ataques de downgrade a "none".
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("algoritmo de firma inesperado")
		}
		return claveSecreta, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("token inválido")
	}
	return claims, nil
}
