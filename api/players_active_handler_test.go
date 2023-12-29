package api

import (
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "cricket/db/mock"
	db "cricket/db/sqlc"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_handlePlayersActive(t *testing.T) {
	testCases := []struct {
		name          string
		url           string
		setupAuth     func(request *http.Request)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "Invalid Auth Token",
			url:  "/api/v1/players/active",
			setupAuth: func(request *http.Request) {
				request.Header.Set("Authorization", "invalid token")
			},
			buildStubs: func(store *mockdb.MockStore) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "valid Token, missing query parameter",
			url:  "/api/v1/players/active",
			setupAuth: func(request *http.Request) {
				request.Header.Set("Authorization", jwtToken)
			},
			buildStubs: func(store *mockdb.MockStore) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				data, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				require.Equal(t,
					"{\"status\":\"Bad request\",\"error\":\"careerYear is required as query parameter\"}\n",
					string(data),
				)
			},
		},
		{
			name: "valid Token, invalid query parameter",
			url:  "/api/v1/players/active?careerYear=202a",
			setupAuth: func(request *http.Request) {
				request.Header.Set("Authorization", jwtToken)
			},
			buildStubs: func(store *mockdb.MockStore) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				data, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				require.Equal(t,
					"{\"status\":\"Bad request\",\"error\":\"careerYear should be a number\"}\n",
					string(data),
				)
			},
		},
		{
			name: "valid Token, Empty data",
			url:  "/api/v1/players/active?careerYear=2020",
			setupAuth: func(request *http.Request) {
				request.Header.Set("Authorization", jwtToken)
			},
			buildStubs: func(store *mockdb.MockStore) {
				p := db.GetPlayersByCareerYearParams{
					CareerYear: sql.NullInt64{Int64: int64(2020), Valid: true},
				}
				store.EXPECT().
					GetPlayersByCareerYear(gomock.Any(), gomock.Eq(p)).
					Times(1).
					Return([]db.Player{}, nil)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				data, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				require.Equal(t, "{\"names\":[]}\n", string(data))
			},
		},
		{
			name: "valid Token, valid data",
			url:  "/api/v1/players/active?careerYear=2020",
			setupAuth: func(request *http.Request) {
				request.Header.Set("Authorization", jwtToken)
			},
			buildStubs: func(store *mockdb.MockStore) {
				p := db.GetPlayersByCareerYearParams{
					CareerYear: sql.NullInt64{Int64: int64(2020), Valid: true},
				}
				store.EXPECT().
					GetPlayersByCareerYear(gomock.Any(), gomock.Eq(p)).
					Times(1).
					Return([]db.Player{
						{Name: "a"},
						{Name: "b"},
						{Name: "c"},
					}, nil)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				data, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				require.Equal(t, "{\"names\":[\"a\",\"b\",\"c\"]}\n", string(data))
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockStore(ctrl)

			tc.buildStubs(store)

			server := newTestServer(store)
			recorder := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodGet, tc.url, nil)

			tc.setupAuth(request)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}
