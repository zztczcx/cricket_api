package batting

import (
	"testing"

	mockdb "cricket/db/mock"
	"go.uber.org/mock/gomock"
)

func Test_Load(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := mockdb.NewMockStore(ctrl)

        testFile := "../../../db/seeds/ODI_data.csv"
	l := loader{
                store: store,
                dataFile: &testFile, 
        }

	store.EXPECT().
		CreatePlayer(gomock.Any(), gomock.Any()).
		Times(2500)

        l.Load()
}
