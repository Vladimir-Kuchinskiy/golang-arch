package jwt

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/gofrs/uuid"
)

type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("token has expired")
	}

	if u.SessionID == 0 {
		return fmt.Errorf("invalid session ID")
	}

	return nil
}

type key struct {
	key     []byte
	created time.Time
}

var currentKID = ""
var keys = map[string]key{}

func GenerateNewKey() error {
	newKey := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, newKey)

	if err != nil {
		return fmt.Errorf("error in generateNewKey while generating a key: %w", err)
	}

	uid, err := uuid.NewV4()

	if err != nil {
		return fmt.Errorf("error in generateNewKey while generating UUID: %w", err)
	}

	keys[uid.String()] = key{
		key:     newKey,
		created: time.Now(),
	}
	currentKID = uid.String()

	return nil
}

func CreateToken(u *UserClaims, key []byte) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, u)
	signedToken, err := t.SignedString(currentKID)
	if err != nil {
		return "", fmt.Errorf("error in creatToken on token sign: %w", err)
	}

	return signedToken, nil
}

func ParseToken(token string, key []byte) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(token, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("invalid signin algorithm")
		}

		kid, ok := t.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid key ID")
		}

		k, ok := keys[kid]
		if !ok {
			return nil, fmt.Errorf("invalid key ID")
		}

		return k.key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error while verifying a token: %w", err)
	}

	if !t.Valid {
		return nil, fmt.Errorf("jwt token in not valid")
	}

	return t.Claims.(*UserClaims), nil
}
