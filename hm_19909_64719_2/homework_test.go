package homework

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendRequest(t *testing.T) {
	t.Run("200 OK response", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusOK)
		}))
		defer svr.Close()

		statusCode, err := SendRequest(svr.URL)

		// Используем require для проверок с автоматическим фейлом теста
		require.NoError(t, err, "Request should not return error")
		require.Equal(t, http.StatusOK, statusCode, "Response status code should be 200")
	})

	t.Run("404 Error client response", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusNotFound)
		}))
		defer svr.Close()

		statusCode, err := SendRequest(svr.URL + "/404")

		// Используем require для проверок с автоматическим фейлом теста
		require.NoError(t, err, "Request should not return error")
		require.Equal(t, http.StatusNotFound, statusCode, "Response status code should be 404")
	})

	t.Run("Handle request error", func(t *testing.T) {
		// Специально используем невалидный URL
		_, err := SendRequest("http://invalid-url-that-does-not-exist")

		// Проверяем что получили ожидаемую ошибку
		require.Error(t, err, "Should return error for invalid request")
		require.Contains(t, err.Error(), "request failed", "Error message should contain context")
	})
}
