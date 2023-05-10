package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/RichieRock/gourses-miniurl/api"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPI_AddUrl(t *testing.T) {
	tests := []struct{
		name string
		handler api.Handler
		payload string
		expectedStatusCode int
		expectedBody string
	}{
		{
			name: "OK",
			handler: &strHandler{str: "testvalue"},
			payload: `{"url": "https://github.com/gourses/miniurl/blob/main/LICENSE"}`,
			expectedStatusCode: http.StatusOK,
			expectedBody: `{"url": "https://github.com/gourses/miniurl/blob/main/LICENSE", "hash": "testvalue"}`,
		},
		{
			name: "bad request",
			handler: nil,
			payload: `{"url": "https://github.com/gourses/miniurl/blob/main/LICENSE"}`,
			expectedStatusCode: http.StatusBadRequest,
			expectedBody: `{"msg": "Bad request"}`,
		},
	}
	for _, tc:= range tests {
		t.Run(tc.name, func (t *testing.T) {

			req := httptest.NewRequest(http.MethodPost, "/api/v1/url", strings.NewReader(tc.payload))
			rr := httptest.NewRecorder()
		
			r := httprouter.New()
			api.Bind(r, tc.handler)
			r.ServeHTTP(rr, req)
		
			// check status code
			assert.Equal(t, tc.expectedStatusCode, rr.Result().StatusCode)
			// read the body
			body, err := io.ReadAll(rr.Result().Body)
			require.NoError(t, err)
			assert.JSONEq(t, tc.expectedBody, string (body))
		})
	}
}

type strHandler struct {
	str string
}

func (h *strHandler) AddUrl(url string) (hash string, err error) {
	return h.str, nil
}