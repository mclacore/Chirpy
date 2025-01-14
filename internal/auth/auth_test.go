package auth_test

import (
	// "fmt"
	"testing"

	"github.com/mclacore/Chirpy/internal/auth"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	testPass := "TestPassword123"
	gotHash, gotErr := auth.HashPassword(testPass)
	if gotErr != nil {
		t.Errorf("Could not hash password value: %q\n", gotErr)
	}

	wantHash, wantErr := bcrypt.GenerateFromPassword([]byte(testPass), auth.Cost)
	if wantErr != nil {
		t.Errorf("Could not generate hash pass from bcrypt library: %q\n", wantErr)
	}

	got := auth.CheckPasswordHash(testPass, gotHash)
	want := bcrypt.CompareHashAndPassword(wantHash, []byte(testPass))

	if got != want {
		t.Errorf("\ngot: %v\nwant: %v\n", gotHash, string(wantHash))
	}
}
