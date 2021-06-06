package asymmetric_jwt

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
	"math/big"
)

type Encoder struct {
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
	// Use this string for a public endpoint to allow JWT verification
	JwkString string
}

func NewEncoder(rsaPrivateKey []byte, rsaPublicKey []byte) (*Encoder, error) {
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(rsaPrivateKey)
	if err != nil {
		return nil, err
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(rsaPublicKey)
	if err != nil {
		return nil, err
	}

	jwkString, err := JwkEncode(verifyKey)
	if err != nil {
		return nil, err
	}

	return &Encoder{
		signKey:   signKey,
		verifyKey: verifyKey,
		JwkString: jwkString,
	}, nil
}

func (service *Encoder) CreateAndSignToken(claims map[string]interface{}) (string, error) {
	jwtClaims := jwt.MapClaims{}
	for key, value := range claims {
		jwtClaims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwtClaims)
	tokenString, err := token.SignedString(service.signKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (service *Encoder) Verify(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return service.verifyKey, nil
	})
}

func (service *Encoder) VerifyFromJwk(tokenString string, jwkString string) (*jwt.Token, error) {
	parsedKey, err := jwk.ParseKey([]byte(jwkString))
	if err != nil {
		return nil, err
	}

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		var raw interface{}

		return raw, parsedKey.Raw(&raw)
	})
}

// JwkEncode encodes public part of an RSA or ECDSA key into a JWK.
// The result is also suitable for creating a JWK thumbprint.
// https://tools.ietf.org/html/rfc7517
// Example from https://github.com/golang/crypto/blob/master/acme/jws.go
func JwkEncode(pub crypto.PublicKey) (string, error) {
	switch pub := pub.(type) {
	case *rsa.PublicKey:
		// https://tools.ietf.org/html/rfc7518#section-6.3.1
		n := pub.N
		e := big.NewInt(int64(pub.E))
		// Field order is important.
		// See https://tools.ietf.org/html/rfc7638#section-3.3 for details.
		return fmt.Sprintf(`{"e":"%s","kty":"RSA","n":"%s"}`,
			base64.RawURLEncoding.EncodeToString(e.Bytes()),
			base64.RawURLEncoding.EncodeToString(n.Bytes()),
		), nil
	case *ecdsa.PublicKey:
		// https://tools.ietf.org/html/rfc7518#section-6.2.1
		p := pub.Curve.Params()
		n := p.BitSize / 8
		if p.BitSize%8 != 0 {
			n++
		}
		x := pub.X.Bytes()
		if n > len(x) {
			x = append(make([]byte, n-len(x)), x...)
		}
		y := pub.Y.Bytes()
		if n > len(y) {
			y = append(make([]byte, n-len(y)), y...)
		}
		// Field order is important.
		// See https://tools.ietf.org/html/rfc7638#section-3.3 for details.
		return fmt.Sprintf(`{"crv":"%s","kty":"EC","x":"%s","y":"%s"}`,
			p.Name,
			base64.RawURLEncoding.EncodeToString(x),
			base64.RawURLEncoding.EncodeToString(y),
		), nil
	}
	return "", errors.New("unsupported key")
}

func GenerateRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, _ := rsa.GenerateKey(rand.Reader, 4096)
	return privkey, &privkey.PublicKey
}

func ExportRsaPrivateKeyAsPem(key *rsa.PrivateKey) []byte {
	privateKey := x509.MarshalPKCS1PrivateKey(key)
	privateKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privateKey,
		},
	)
	return privateKeyPem
}

func ParseRsaPrivateKeyFromPem(privatePem []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privatePem)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return private, nil
}

func ExportRsaPublicKeyAsPem(publicKey *rsa.PublicKey) ([]byte, error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	publicKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: publicKeyBytes,
		},
	)

	return publicKeyPem, nil
}

func ParseRsaPublicKeyFromPem(pubPEM []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pubPEM)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break
	}
	return nil, errors.New("key type is not RSA")
}
