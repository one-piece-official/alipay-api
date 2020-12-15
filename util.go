package alipay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// RsaSignWithSha256Hex 签名：采用sha256算法进行签名并输出为hex格式（私钥PKCS1格式）.
// NOTE: https://blog.csdn.net/whatday/article/details/97623948
func RsaSignWithSha256Hex(data string, secret string) (signature string, err error) {
	pk, err := parseRawPrivateKey([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("invalid private key: %w", err)
	}

	bs, err := privateSign(pk, []byte(data))
	if err != nil {
		return "", fmt.Errorf("private sign failed: %w", err)
	}

	return base64.StdEncoding.EncodeToString(bs), nil
}

// ParseRawPrivateKey returns a private key from a PEM encoded private key.
// It supports RSA (PKCS#1) private keys.
// If the private key is encrypted, it will return a PassphraseMissingError.
// NOTE: https://sourcegraph.com/github.com/golang/crypto/-/blob/ssh/keys.go#L1120:1
func parseRawPrivateKey(pemBytes []byte) (pk *rsa.PrivateKey, err error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("no key found")
	}

	if encryptedBlock(block) {
		return nil, errors.New("this private key is passphrase protected")
	}

	// RFC5208 - https://tools.ietf.org/html/rfc5208
	if block.Type != "RSA PRIVATE KEY" {
		err = errors.New("block type is wrong")

		return nil, fmt.Errorf("%w, unsupported key type %q", err, block.Type)
	}

	blockBytes := block.Bytes
	ok := x509.IsEncryptedPEMBlock(block)

	if ok {
		blockBytes, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, fmt.Errorf("decrypt pem block failed: %w", err)
		}
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(blockBytes)
	if err != nil {
		return nil, fmt.Errorf("parse pcks1 private key failed: %w", err)
	}

	return privateKey, nil
}

// encryptedBlock tells whether a private key is
// encrypted by examining its Proc-Type header
// for a mention of ENCRYPTED
// according to RFC 1421 Section 4.6.1.1.
func encryptedBlock(block *pem.Block) bool {
	return strings.Contains(block.Headers["Proc-Type"], "ENCRYPTED")
}

// NOTE: https://gist.github.com/wongoo/2b974a9594627114bea3e53c794980cd
func privateSign(key *rsa.PrivateKey, data []byte) ([]byte, error) {
	hashedData, err := hash(data)
	if err != nil {
		return data, err
	}

	return rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashedData)
}

// NOTE: https://blog.csdn.net/whatday/article/details/97623948
func hash(data []byte) ([]byte, error) {
	hash := sha256.New()
	_, err := hash.Write(data)
	if err != nil {
		return data, errors.New("hash write failed")
	}

	s := hash.Sum(nil)

	return s[:], nil
}
