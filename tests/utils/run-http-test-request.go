package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
)

type HTTPTestOptions struct {
	Headers  *map[string]string
	Cookies  *[]*http.Cookie
	Body     *map[string]interface{}
	FormData *HTTPFormData
}

type HTTPFormData struct {
	Buffer *bytes.Buffer
	Writer *multipart.Writer
}

type HTTPTestResponse struct {
	StatusCode int
	Cookies    []*http.Cookie
	Obj        map[string]interface{}
	Arr        []map[string]interface{}
}

/**
 * Executes an HTTP request within a specific application module and returns the parsed response.
 *
 * @param t - The *testing.T instance used for asserting and reporting test failures.
 * @param module - The name of the application module to initialize for handling the request (e.g., "users", "chats").
 * @param method - The HTTP method to use for the request (e.g., "GET", "POST", "PUT", "DELETE").
 * @param url - The URL path of the request (relative to the test server's base URL).
 * @param opt - Additional request options such as headers, body (payload), and cookies.
 *
 * @returns HTTPTestResult - The result of the request, including the HTTP status code, cookies,
 * and the decoded response body either as an object (`Obj`) or an array of objects (`Arr`),
 * depending on the response format.
 */

func RunHTTPTestRequest(t *testing.T, module, method, url string, opt HTTPTestOptions) HTTPTestResponse {
	app := SetupTestModule(Module(module))
	server := httptest.NewServer(fiberToHttpHandler(app.Handler()))
	defer server.Close()

	var requestBody io.Reader
	var contentType string

	// Define o corpo como JSON
	if opt.Body != nil {
		bodyBytes, err := json.Marshal(opt.Body)
		require.NoError(t, err, "falha ao codificar payload")
		contentType = "application/json"
		requestBody = bytes.NewBuffer(bodyBytes)
	}

	// Define o corpo como multipart/form-data
	if opt.FormData != nil {
		requestBody = opt.FormData.Buffer
		contentType = opt.FormData.Writer.FormDataContentType()
	}

	req := httptest.NewRequest(method, server.URL+url, requestBody)
	if opt.Headers != nil {
		for key, value := range *opt.Headers {
			req.Header.Set(key, value)
		}
	}

	// Seta Content-Type se foi definido
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	if opt.Cookies != nil {
		for _, cookie := range *opt.Cookies {
			req.AddCookie(cookie)
		}
	}

	resp, err := app.Test(req)
	require.NoError(t, err, "erro ao executar a requisição")
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "erro ao ler corpo da resposta")

	var obj map[string]interface{}
	var arr []map[string]interface{}

	result := HTTPTestResponse{
		StatusCode: resp.StatusCode,
		Cookies:    resp.Cookies(),
	}

	// Tenta decodificar como objeto
	if err := json.Unmarshal(respBody, &obj); err == nil {
		result.Obj = obj
		return result
	}

	// Tenta decodificar como array de objetos
	if err := json.Unmarshal(respBody, &arr); err == nil {
		result.Arr = arr
		return result
	}

	// Se nenhum dos dois, retorna com ambos nil
	return result

}

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
		for key, value := range ctx.Response.Header.All() {
			w.Header().Set(string(key), string(value))
		}

		w.Write(ctx.Response.Body())
	})
}
