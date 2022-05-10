package lettucejwt

import (
	"crypto/ed25519"
	"encoding/base64"
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Issuer    string `json:"iss"`
	ExpiresAt int64  `json:"exp"`
	IssuedAt  int64  `json:"iat"`
	Subject   string `json:"sub"`
}

type Algorithm = jwt.SigningMethodEd25519

const Issuer = "Laravel"
const Type = "JWT"

// Read will parse and validate the JWT token from the Lettuce login service, and verify that, in
// addition to usual JWT validation, some specific properties will match what we expect of Lettuce,
// and all expected fields are present. It will then return all claims as a Claims struct and
// potentially a corresponding error
func Read(tokenString string, pubKey string) (claims Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validate algorithm
		if _, ok := token.Method.(*Algorithm); !ok {
			return nil, errors.New("invalid algorithm")
		}

		// validate issuer
		if !token.Claims.(jwt.MapClaims).VerifyIssuer(Issuer, true) {
			return nil, errors.New("issuer doesn't match")
		}

		// validate type
		if typ, ok := token.Header["typ"].(string); !ok {
			return nil, errors.New("couldn't read header")
		} else if typ != Type {
			return nil, errors.New("type does not match")
		}

		// verify signature
		decodedKey, err := base64.StdEncoding.DecodeString(pubKey)
		if err != nil {
			return nil, errors.New("invalid key")
		}
		return ed25519.PublicKey(decodedKey), nil
	})

	if err != nil {
		return claims, err
	}

	if mapClaims, ok := token.Claims.(jwt.MapClaims); ok {
		if iss, ok := mapClaims["iss"].(string); ok {
			claims.Issuer = iss
		}

		if exp, ok := mapClaims["exp"].(float64); ok {
			claims.ExpiresAt = int64(exp)
		}

		if iat, ok := mapClaims["iat"].(float64); ok {
			claims.IssuedAt = int64(iat)
		}

		if sub, ok := mapClaims["sub"].(string); ok {
			claims.Subject = sub
		}
	}

	return claims, err
}
