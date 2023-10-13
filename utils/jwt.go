/**
 * Created by GOLAND.
 * User: pengyu
 * Time: 2023/10/12 17:52
 */

package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func MakeJwtToken(secretKey string, seconds int64, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Duration(seconds) * time.Second).Unix()
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
