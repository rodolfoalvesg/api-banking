package accounts

import (
	"context"
	"errors"
	"testing"
)

// TestGeneratePasswdHash, gera um hash da senha
func TestGeneratePasswdHash(t *testing.T) {
	t.Parallel()

	t.Run("Generate Successfully Passwd Hash ", func(t *testing.T) {
		TestCaseA := Account{
			Secret: "",
		}

		_, err := GeneratePasswdHash(context.Background(), TestCaseA)
		if err != nil {
			t.Errorf("got %v, want %v", err, nil)
		}

	})

	t.Run("Error creating hash", func(t *testing.T) {
		TestCaseB := Account{
			Secret: "7&%",
		}

		_, err := GeneratePasswdHash(context.Background(), TestCaseB)
		if err != nil {
			t.Errorf("got %v, want %v", err, nil)
		}

	})
}

func TestValidateCreateAccountData(t *testing.T) {
	t.Parallel()

	t.Run("Invalid Secret", func(t *testing.T) {
		TestCaseA := Account{
			Secret: "",
		}

		err := ValidateCreateAccountData(TestCaseA)
		if !errors.Is(err, ErrInvalidPasswd) {
			t.Errorf("got %v, want %v", err, ErrInvalidPasswd)
		}
	})

	t.Run("Invalid CPF", func(t *testing.T) {
		TestCaseA := Account{
			CPF: "",
		}

		err := ValidateCreateAccountData(TestCaseA)
		if !errors.Is(err, ErrInvalidPasswd) {
			t.Errorf("got %v, want %v", err, ErrInvalidLengthCPF)
		}
	})

}
