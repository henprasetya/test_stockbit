package omdb

import (
	"github.com/golang/mock/gomock"
	mockOmdb "github.com/henprasetya/omdbapi/pkg/mock/repo/omdb"
	"github.com/henprasetya/omdbapi/pkg/model"
	"github.com/henprasetya/omdbapi/pkg/repo/omdb"
	"github.com/pkg/errors"
	"reflect"
	"testing"
)

func Test_service_SearchMovie(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockCtrl.Finish()

	mockApi := mockOmdb.NewMockMovieRepo(mockCtrl)

	type fields struct {
		api omdb.MovieRepo
	}
	type args struct {
		search string
		page   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Response
		wantErr bool
		mock func()
	}{
		{
			name:    "error search movie",
			fields:  fields{
				api: mockApi,
			},
			args:    args{
				search: "a",
				page:   "1",
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				mockApi.EXPECT().SearchMovie(gomock.Any(), gomock.Any()).Return(nil, errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				api: tt.fields.api,
			}
			tt.mock()
			got, err := s.SearchMovie(tt.args.search, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchMovie() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_MovieDetail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockCtrl.Finish()

	mockApi := mockOmdb.NewMockMovieRepo(mockCtrl)

	type fields struct {
		api omdb.MovieRepo
	}
	type args struct {
		omdbID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Movie
		wantErr bool
		mock func()
	}{
		{
			name:    "error movie detail",
			fields:  fields{
				api: mockApi,
			},
			args:    args{
				omdbID: "1",
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				mockApi.EXPECT().MovieDetail(gomock.Any(), gomock.Any()).Return(nil, errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				api: tt.fields.api,
			}
			tt.mock()
			got, err := s.SearchMovie(tt.args.search, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchMovie() got = %v, want %v", got, tt.want)
			}
		})
	}
}