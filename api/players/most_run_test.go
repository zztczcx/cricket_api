package players

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "cricket/db/mock"
	db "cricket/db/sqlc"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_HandleMostRuns(t *testing.T) {
	testCases := []struct {
		name          string
		url           string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "missing query parameter",
			url:  "/players/most_runs",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetPlayerOfMostRuns(gomock.Any()).
					Times(1).
					Return(db.Player{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				data, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				require.Equal(t,
					"{\"name\":\"\",\"runs\":null}\n",
					string(data),
				)
			},
		},
		{
			name: "invalid query parameter",
			url:  "/players/most_runs?careerYear=202a",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetPlayerOfMostRuns(gomock.Any()).
					Times(1).
					Return(db.Player{}, nil)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				data, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				require.Equal(t,
					"{\"name\":\"\",\"runs\":null}\n",
					string(data),
				)
			},
		},
		{
			name: "Empty data",
			url: "/players/most_runs?careerEndYear=2020",
			buildStubs: func(store *mockdb.MockStore) {
				p := db.ToNullInt64("2020")
				store.EXPECT().
					GetPlayerOfMostRunsByCareerEndYear(gomock.Any(), gomock.Eq(p)).
					Times(1).
					Return(db.Player{}, nil)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				data, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				require.Equal(t,
					"{\"name\":\"\",\"runs\":null}\n",
					string(data),
                                )
			},
		},
		{
			name: "valid data",
			url: "/players/most_runs?careerEndYear=2020",
			buildStubs: func(store *mockdb.MockStore) {
				p := db.ToNullInt64("2020")
		                              store.EXPECT().
					GetPlayerOfMostRunsByCareerEndYear(gomock.Any(), gomock.Eq(p)).
					Times(1).
					Return(db.Player{
                                                Name: "a", Runs: db.ToNullInt64("1000"),
					}, nil)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				data, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				require.Equal(t, "{\"name\":\"a\",\"runs\":1000}\n", string(data))
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

			router := chi.NewRouter()
			NewPlayersHandler(router, store)
			recorder := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodGet, tc.url, nil)

			require.NoError(t, err)

			router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}
