package baseLearn

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func Wrap(handler func(w http.ResponseWriter, req *http.Request), signatureKey []byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := new(bytes.Buffer)
		defer r.Body.Close()
		ioutil.ReadAll(io.TeeReader(r.Body, body))
		r.Body = ioutil.NopCloser(body) // 我们读取主体两次, 我们必须包装原始的 ReadCloser
		signature := strings.TrimSpace(r.Header.Get("signature"))
		if err := CheckSignature(signature, signatureKey, body.Bytes()); err != nil {
			// Error handling
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write([]byte("406 Not Acceptable"))
			return
		}
		http.HandlerFunc(handler).ServeHTTP(w, r)
	})
}

func CheckSignature(rawSign string, pubPem []byte, data []byte) error {
	var err error
	var sign []byte
	var pub interface{}
	sign, err = base64.StdEncoding.DecodeString(rawSign)
	if err != nil {
		return err
	}
	block, _ := pem.Decode(pubPem)
	if block == nil {
		return errors.New("Failed to decode public PEM")
	}
	pub, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	newHash := crypto.SHA256.New()
	newHash.Write(data)
	opts := rsa.PSSOptions{SaltLength: 20} // Java default salt length
	err = rsa.VerifyPSS(pub.(*rsa.PublicKey), crypto.SHA256, newHash.Sum(nil), sign, &opts)
	return err
}

func PostItHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("ok"))
}

func RegisterHandler() {
	signature, _ := ioutil.ReadFile("/path/of/public/key")

	r := mux.NewRouter()
	r.Handle("/postit", Wrap(PostItHandler, signature)).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe("8080", nil)
}

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
