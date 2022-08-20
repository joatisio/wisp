package encryption

import "testing"

func BenchmarkGetPasswordHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetPasswordHash("randpass")
	}
}

func BenchmarkCheckPassword(b *testing.B) {
	pass := "randpass"

	hash := GetPasswordHash(pass)

	for i := 0; i < b.N; i++ {
		_ = CheckPassword(hash, pass)
	}
}

func TestCheckPassword(t *testing.T) {
	const pass = "randompass"

	hash := GetPasswordHash(pass)
	if err := CheckPassword(hash, pass); err != nil {
		t.Error(err)
	}
}
