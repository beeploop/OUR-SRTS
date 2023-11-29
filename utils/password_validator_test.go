package utils

import "testing"

type MockPassword struct {
	Password string
	Result   bool
}

func TestValidPassword(t *testing.T) {
	mockPasswords := []MockPassword{
		{
			Password: "password",
			Result:   false,
		},
		{
			Password: "Password",
			Result:   false,
		},
		{
			Password: "Password1",
			Result:   false,
		},
		{
			Password: "password 1",
			Result:   false,
		},
        {
            Password: "pass_1",
            Result:   false,
        },
		{
			Password: "Password1!",
			Result:   true,
		},
	}

	for i, mockPassword := range mockPasswords {
		result := ValidPassword(mockPassword.Password)
		if result != mockPassword.Result {
			t.Errorf("Test %d: FAILED!: Expected %v, got %v", i, mockPassword.Result, result)
		} else {
			t.Logf("Test %d: PASSED!: Expected %v, got %v", i, mockPassword.Result, result)
		}
	}
}
