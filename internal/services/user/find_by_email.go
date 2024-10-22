package user

import (
	"back-end-inventory/internal/domains"
)

//	type TokenG struct{
//		TokenGenerator ports.TokenGenerator
//	}
// var jwtKey = []byte("mi_secret_clave")

// func generateJWT(email string) (string, error) {
// 	expirationTime := time.Now().Add(24 * time.Hour)

// 	claims := &jwt.StandardClaims{
// 		ExpiresAt: expirationTime.Unix(),
// 		Subject:   email,
// 	}

// 	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

// 	tokenString, err := tokenJWT.SignedString(jwtKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }

func (s Services) FindUserByEmail(email string) (*domains.Users, error) {
	return s.Repo.FindUserByEmail(email)
	// user, err := s.Repo.FindUserByEmail(email)
	// if err != nil || user == nil {
	// 	return nil, errors.New("correo o contraseña incorrectos")
	// }

	// // token, err := s.TokenGenerator.GenerateToken(user.Email)
	// // if err != nil {
	// // 	return nil, err
	// // }

	// token, err := generateJWT(user.Email)
	// if err != nil {
	// 	return nil, err
	// }

	// userResponse := &domains.UsersResponse{
	// 	Token: token,
	// 	User:  user,
	// }

	// return userResponse, nil

}
