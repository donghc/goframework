package password

import "testing"

func TestGenerateLoginToken(t *testing.T) {

	tests := []struct {
		name      string
		id        int32
		wantToken string
	}{
		{name: "steven", id: 123456, wantToken: "d0caa5ee0477830efe166a524ec7824d"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotToken := GenerateLoginToken(tt.id); gotToken != tt.wantToken {
				t.Errorf("GenerateLoginToken() = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}

func TestGeneratePassword(t *testing.T) {
	tests := []struct {
		name         string
		str          string
		wantPassword string
	}{
		{name: "steven", str: "123456", wantPassword: "bff46e451a3fb194db9800562a48bd993b40f2d79fb72628c1ee0373e013fa6f"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPassword := GeneratePassword(tt.str); gotPassword != tt.wantPassword {
				t.Errorf("GeneratePassword() = %v, want %v", gotPassword, tt.wantPassword)
			}
		})
	}
}

func TestResetPassword(t *testing.T) {
	tests := []struct {
		name         string
		wantPassword string
	}{
		{name: "123456", wantPassword: "5b37b751fd3ccee62d95da9b9ed9e08a850e9833d6c2b4110f99daa9219dcff0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPassword := ResetPassword(); gotPassword != tt.wantPassword {
				t.Errorf("ResetPassword() = %v, want %v", gotPassword, tt.wantPassword)
			}
		})
	}
}
