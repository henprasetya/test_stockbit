package mysql

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	mockDB "github.com/henprasetya/omdbapi/pkg/mock/repo/mysql"
)

func Test_movieDB_SelectFromDB(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockCtrl.Finish()

	mockDB := mockDB.NewMockOmdbMysql(mockCtrl)
	db, _, _ := sqlmock.New()
	type fields struct {
		db *sql.DB
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		mock    func()
	}{
		{
			name: "error get",
			fields: fields{
				db: db,
			},

			wantErr: true,
			mock: func() {
				mockDB.EXPECT().SelectFromDb().Return().Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &omdbmysql{
				db: tt.fields.db,
			}
			tt.mock()
			m.SelectFromDb()

		})
	}
}
