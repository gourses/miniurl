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
	const payload = `{'"url": " https://github.com/gourses/miniurl/blob/main/LICENSE"}`
	const expectedBody = `{'"url": " https://github.com/gourses/miniurl/blob/main/LICENSE", "hash":"testvalue"}`
	const expectedStatuscode = http.StatusOK

	req := httptest.NewRequest(http.MethodPost, "/api/v1/url", strings.NewReader(payload))
	rr := httptest.NewRecorder()
	r := httprouter.New()
	api.Bind(r, nil)
	r.ServeHTTP(rr, req)

	assert.Equal(t, expectedStatuscode, rr.Result().Body)
	// go palauttaa virheen funkction
	// jos haluaa ignorata _ scorella, body,_
	body, err := io.ReadAll(rr.Result().Body)
	//require stoppaa suorituksen ja toista assertia ei suoriteta
	require.NoError(t, err)
	assert.Equal(t, expectedBody, string(body))
}
