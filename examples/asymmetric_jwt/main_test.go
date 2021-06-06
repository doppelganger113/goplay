package asymmetric_jwt

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

const invalidJwt = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

func TestGenerateRsaKeyPair(t *testing.T) {
	priv, pub := GenerateRsaKeyPair()

	// Export the keys to pem string
	privatePem := ExportRsaPrivateKeyAsPem(priv)
	pubPem, err := ExportRsaPublicKeyAsPem(pub)
	assert.Nil(t, err)

	// Import the keys from pem string
	privateParsed, err := ParseRsaPrivateKeyFromPem(privatePem)
	assert.Nil(t, err)

	pubParsed, err := ParseRsaPublicKeyFromPem(pubPem)
	assert.Nil(t, err)

	// Export the newly imported keys
	privateParsedPem := ExportRsaPrivateKeyAsPem(privateParsed)
	pubParsedPem, err := ExportRsaPublicKeyAsPem(pubParsed)
	assert.Nil(t, err)

	//  Check that the exported/imported keys match the original keys
	assert.Equal(t, privatePem, privateParsedPem)
	assert.Equal(t, pubPem, pubParsedPem)
}

func TestNewEncoder(t *testing.T) {
	priv, pub := GenerateRsaKeyPair()

	// Export the keys to pem string
	privatePem := ExportRsaPrivateKeyAsPem(priv)
	pubPem, err := ExportRsaPublicKeyAsPem(pub)
	assert.Nil(t, err)

	jwtService, err := NewEncoder(pubPem, privatePem)
	assert.Nil(t, err)

	assert.NotNil(t, jwtService)
}

func TestJwtService_CreateAndSignToken(t *testing.T) {
	priv, pub := GenerateRsaKeyPair()

	// Export the keys to pem string
	privatePem := ExportRsaPrivateKeyAsPem(priv)
	pubPem, err := ExportRsaPublicKeyAsPem(pub)
	assert.Nil(t, err)

	jwtService, err := NewEncoder(pubPem, privatePem)
	assert.Nil(t, err)

	myClaims := make(map[string]interface{})
	myClaims["foo"] = "my foo"
	myClaims["bar"] = "my bar"

	token, err := jwtService.CreateAndSignToken(myClaims)
	assert.Nil(t, err)

	parsedToken, err := jwtService.Verify(token)
	assert.Nil(t, err)

	assert.True(t, parsedToken.Valid)
	assert.Equal(t, parsedToken.Claims.(jwt.MapClaims)["foo"], "my foo")
	assert.Equal(t, parsedToken.Claims.(jwt.MapClaims)["bar"], "my bar")
}

func TestJwtService_Verify(t *testing.T) {
	priv, pub := GenerateRsaKeyPair()

	// Export the keys to pem string
	privatePem := ExportRsaPrivateKeyAsPem(priv)
	pubPem, err := ExportRsaPublicKeyAsPem(pub)
	assert.Nil(t, err)

	jwtService, err := NewEncoder(pubPem, privatePem)
	assert.Nil(t, err)

	claims := make(map[string]interface{})
	claims["foo"] = "my foo"

	signedToken, err := jwtService.CreateAndSignToken(claims)
	assert.Nil(t, err)
	assert.NotEmpty(t, signedToken)

	parsedToken, err := jwtService.Verify(signedToken)
	assert.Nil(t, err)
	assert.NotNil(t, parsedToken)
	assert.True(t, parsedToken.Valid)
	assert.Equal(t, parsedToken.Claims.(jwt.MapClaims)["foo"], "my foo")

	_, err = jwtService.Verify(invalidJwt)
	assert.NotNil(t, err)
}

func TestJwtService_VerifyFromJwk(t *testing.T) {
	priv, pub := GenerateRsaKeyPair()

	// Export the keys to pem string
	privatePem := ExportRsaPrivateKeyAsPem(priv)
	pubPem, err := ExportRsaPublicKeyAsPem(pub)
	assert.Nil(t, err)

	jwtService, err := NewEncoder(pubPem, privatePem)
	assert.Nil(t, err)

	claims := make(map[string]interface{})
	claims["foo"] = "my foo"
	token, err := jwtService.CreateAndSignToken(claims)
	assert.Nil(t, err)

	verified, err := jwtService.VerifyFromJwk(token, jwtService.JwkString)
	assert.Nil(t, err)
	assert.True(t, verified.Valid)
	assert.Equal(t, verified.Claims.(jwt.MapClaims)["foo"], "my foo")
}

func TestJwkEncode(t *testing.T) {
	_, pub := GenerateRsaKeyPair()
	jwkString, err := JwkEncode(pub)
	assert.Nil(t, err)
	assert.NotEmpty(t, jwkString)

	var result map[string]interface{}
	err = json.Unmarshal([]byte(jwkString), &result)
	assert.Nil(t, err)

	assert.NotEmpty(t, result["e"])
	assert.NotEmpty(t, result["n"])
	assert.Equal(t, result["kty"], "RSA")
}
