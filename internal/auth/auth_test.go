package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	t.Run("fails with empty Authorization header", func(t *testing.T) {

		headers := http.Header{
			"Authorization": {""},
		}
		actual, actualerr := GetAPIKey(headers)

		if actual != "" || actualerr == nil {
			t.Errorf("expected error, got %s", actual)
		}
	})

	t.Run("succeeds with valid api key", func(t *testing.T) {
		headers := http.Header{
			"Authorization": {"ApiKey 12345"},
		}

		expected := "12345"
		actual, actualerr := GetAPIKey(headers)

		if actualerr != nil {
			t.Fatalf("unexpected error %s", actualerr.Error())
		}

		if actual != expected {
			t.Errorf("expected %s got %s", expected, actual)
		}
	})

}
