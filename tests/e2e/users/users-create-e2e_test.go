package users_create_e2e__test

import (
	"api-gym-on-go/tests/utils"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func fiberToHttpHandler(handler fasthttp.RequestHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req := &fasthttp.Request{}
		req.SetRequestURI(r.URL.String())
		req.Header.SetMethod(r.Method)

		for k, vv := range r.Header {
			for _, v := range vv {
				req.Header.Add(k, v)
			}
		}

		if r.Body != nil {
			bodyBytes, _ := io.ReadAll(r.Body)
			req.SetBody(bodyBytes)
		}

		ctx := &fasthttp.RequestCtx{}

		handler(ctx)

		w.WriteHeader(ctx.Response.StatusCode())
		ctx.Response.Header.VisitAll(func(key, value []byte) {
			w.Header().Set(string(key), string(value))
		})

		w.Write(ctx.Response.Body())
	})
}

func TestUserRegisterE2E(t *testing.T) {
	utils.ResetDb()
	app := utils.SetupTestApp()
	server := httptest.NewServer(fiberToHttpHandler(app.Handler()))
	defer server.Close()

	t.Run("should be able to register", func(t *testing.T) {
		payload := map[string]interface{}{
			"user_name":     "Jhon Doe",
			"email":         "user@email.com",
			"password_hash": "123456",
		}

		body, err := json.Marshal(payload)
		if err != nil {
			t.Fatalf("falha ao codificar payload: %v", err)
		}

		req := httptest.NewRequest("POST", "/users/create", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req, -1)

		assert.Equalf(t, 200, resp.StatusCode, "get HTTP status 200")
	})
}
