package auth

import (
	"bytes"
	"testing"
)

func TestJWT(t *testing.T) {
	want := []byte("-----BEGIN PUBLIC KEY-----")
	if !bytes.Contains(rawPublicKey, want) {
		t.Errorf("rawPublicKey = %v, want %v", rawPublicKey, want)
	}

	want = []byte("-----BEGIN RSA PRIVATE KEY-----")
	if !bytes.Contains(rawPrivateKey, want) {
		t.Errorf("rawPrivateKey = %v, want %v", rawPrivateKey, want)
	}
}
