package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// default JWT token for testing
var headers = map[string]string{"Authorization": "BEARER eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMTF9._cLJn0xFS0Mdr_4L_8XF8-8tv7bHyOQJXyWaNsSqlEs"}

func TestAuthorization(t *testing.T) {
        testCases := []struct {
                name string
                path string
                method string
                headers map[string]string
                data interface{}
                expectedStatus int
                expectedBody string
        }{
                {
                        name: "most_runs without Token",
                        path: "http://localhost:8080/api/v1/players/most_runs",
                        method: "GET",
                        headers: nil,
                        data: nil,
                        expectedStatus: http.StatusUnauthorized,
                        expectedBody: "no token found\n",

                },
                {
                        name: "most_runs with Token",
                        path: "http://localhost:8080/api/v1/players/most_runs",
                        method: "GET",
                        headers: headers,
                        data: nil,
                        expectedStatus: http.StatusOK,
                        expectedBody: "{\"name\":\"SR Tendulkar (INDIA)\",\"runs\":18426}\n",
                },
                {
                        name: "active without Token",
                        path: "http://localhost:8080/api/v1/players/active?careerEndYear=2018",
                        method: "GET",
                        headers: nil,
                        data: nil,
                        expectedStatus: http.StatusUnauthorized,
                        expectedBody: "no token found\n",
                },
                {
                        name: "active with Token",
                        path: "http://localhost:8080/api/v1/players/active?careerYear=1970",
                        method: "GET",
                        headers: headers,
                        data: nil,
                        expectedStatus: http.StatusOK,
                        expectedBody: "{\"names\":[]}\n",
                },
        }

        for _, tc := range testCases {
                t.Run(tc.name, func(t *testing.T) {
                        resp, body := makeRequest(t, tc.method, tc.path, tc.data, tc.headers)
                        require.Equal(t, tc.expectedStatus, resp.StatusCode)
                        require.Equal(t, tc.expectedBody, string(body))
                })
        }
}

func makeRequest(t *testing.T, method string, path string, data interface{}, headers map[string]string) (*http.Response, []byte) {
	var body []byte

	if data != nil {
		var err error
		body, err = json.Marshal(data)
		assert.NoError(t, err)
	}

	req, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
        for k, v := range headers {
	        req.Header.Set(k, v)
        }

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	// resp.Body is io.Reader and is treated as stream, so you can read from it once.
	// We must set body again to allow read it again.
	bodyCopy, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	resp.Body.Close()

	resp.Body = io.NopCloser(bytes.NewBuffer(bodyCopy))

	fmt.Printf("Request %s %s done, response status: %d: body: %s\n", method, path, resp.StatusCode, bodyCopy)

	return resp, bodyCopy
}
