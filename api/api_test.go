package api_test

import (
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"miniurl/api"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestApi_AddUrl(t *testing.T) {
	test := []struct {
		name               string
		handler            api.Handler
		payload            string
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name:               "OK",
			handler:            &stringHandler{str: "testvalue"},
			payload:            `{"url": "https://github.com/gourses/miniurl/blob/main/LICENSE"}`,
			expectedStatusCode: http.StatusOK,
			expectedBody:       `{"url": "https://github.com/gourses/miniurl/blob/main/LICENSE", "hash":"testvalue"}`,
		},
		{
			name:               "bad request",
			handler:            &stringHandler{str: "testvalue"},
			payload:            `invalid data json`,
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"msg": "bad request"}`,
		},
	}

	const payload = `{"url": "https://github.com/gourses/miniurl/blob/main/LICENSE"}`
	const expectedBody = `{"url": "https://github.com/gourses/miniurl/blob/main/LICENSE", "hash":"testvalue"}`
	const expectedStatuscode = http.StatusOK
	for _, tc := range test {
		req := httptest.NewRequest(http.MethodPost, "/api/v1/url", strings.NewReader(payload))
		rr := httptest.NewRecorder()
		r := httprouter.New()

		api.Bind(r, tc.handler)
		r.ServeHTTP(rr, req)

		//assert.Equal(t, expectedStatuscode, rr.Result().Body)
		// go palauttaa virheen funkction
		// jos haluaa ignorata _ scorella, body,_
		body, err := io.ReadAll(rr.Result().Body)
		//require stoppaa suorituksen ja toista assertia ei suoriteta
		require.NoError(t, err)
		assert.JSONEq(t, expectedBody, string(body))
	}
}

type stringHandler struct {
	str string
}

func (h *stringHandler) AddUrl(Url string) (hash string, err error) {
	return h.str, nil
}
