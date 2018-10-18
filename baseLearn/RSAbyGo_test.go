package baseLearn

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"net/http"
	"testing"
)

type TestWriter struct {
	header  http.Header
	status  int
	message string
}

func (w *TestWriter) Header() http.Header {
	return w.header
}

func (w *TestWriter) Write(b []byte) (int, error) {
	w.message = string(b)
	return len(b), nil
}

func (w *TestWriter) WriteHeader(s int) {
	w.status = s
}

func TestWrapAllValid(t *testing.T) {
	pk, _ := rsa.GenerateKey(rand.Reader, 1024)
	pubDer, _ := x509.MarshalPKIXPublicKey(&pk.PublicKey)
	pubPem := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Headers: nil, Bytes: pubDer})

	content := "body"
	newHash := crypto.SHA256.New()
	newHash.Write([]byte(content))
	opts := rsa.PSSOptions{SaltLength: 20}
	sign, _ := rsa.SignPSS(rand.Reader, pk, crypto.SHA256, newHash.Sum(nil), &opts)

	body := bytes.NewBufferString(content)
	req, _ := http.NewRequest("GET", "http://valami", body)
	req.Header.Add("signature", base64.StdEncoding.EncodeToString(sign))
	writer := new(TestWriter)
	writer.header = req.Header
	handler := Wrap(func(w http.ResponseWriter, req *http.Request) {}, pubPem)
	handler.ServeHTTP(writer, req)
	if writer.status != 0 {
		t.Errorf("writer.status 0 == %d", writer.status)
	}
}
