package verify

import "testing"

func TestVerifyShouldBeSuccess(t *testing.T) {
	t.Run("verify should be success", func(t *testing.T) {
		_ = LineJWT()
	})
}

func TestGetTokenShouldBeSuccess(t *testing.T) {
	t.Run("get token should be success", func(t *testing.T) {
		token := `Bearer eyJhbGciOiJIUzI1NiJ9.gYp-4ieS6GdXCU3QOduonlmLAU7E1Tz-DdksGX0a6SinuXLOSqos2Q6zBJiSze7leabHwQQT6DE4qGWtX4dHyzr5CfZZ_HYvU9d7bYRAAIVlCZttEcR4B5q_2SN1KOjD7Ly_I5j6btG0VoXwcNqtZqfxXNUeL9kc4SVVx9g6tFI.CKFYwLEzx63vtmFcWJGIZa1nLllmlpZquIUfhZI7h4E`
		expected := "eyJhbGciOiJIUzI1NiJ9.gYp-4ieS6GdXCU3QOduonlmLAU7E1Tz-DdksGX0a6SinuXLOSqos2Q6zBJiSze7leabHwQQT6DE4qGWtX4dHyzr5CfZZ_HYvU9d7bYRAAIVlCZttEcR4B5q_2SN1KOjD7Ly_I5j6btG0VoXwcNqtZqfxXNUeL9kc4SVVx9g6tFI.CKFYwLEzx63vtmFcWJGIZa1nLllmlpZquIUfhZI7h4E"
		actual := getToken(token)
		if actual != expected {
			t.Errorf("get token should be %v but get %v", expected, actual)
		}
	})

	t.Run("get token should be empty", func(t *testing.T) {
		token := ``
		expected := ""
		actual := getToken(token)
		if actual != expected {
			t.Errorf("get token should be %v but get %v", expected, actual)
		}
	})
}
