package omdb

import (
	"github.com/golang/mock/gomock"
	"github.com/henprasetya/omdbapi/pkg/lib/http"
	mockHttp "github.com/henprasetya/omdbapi/pkg/mock/lib/http"
	"github.com/henprasetya/omdbapi/pkg/model"
	"github.com/pkg/errors"
	"reflect"
	"testing"
)

func Test_movieRepo_SearchMovie(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockCtrl.Finish()

	mockClient := mockHttp.NewMockHttpClient(mockCtrl)

	type fields struct {
		client  http.HttpClient
		baseUrl string
		apikey  string
	}
	type args struct {
		s string
		page string
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
			name:    "error get",
			fields:  fields{
				client:  mockClient,
				baseUrl: "",
				apikey:  "",
			},
			args:    args{
				s: "Batman",
				page:"2",
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				mockClient.EXPECT().GET(gomock.Any()).Return(nil, errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &movieRepo{
				client:  tt.fields.client,
				baseUrl: tt.fields.baseUrl,
				apikey:  tt.fields.apikey,
			}
			tt.mock()
			got, err := m.SearchMovie(tt.args.s, tt.args.page)
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

func Test_movieRepo_MovieDetail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockCtrl.Finish()

	mockClient := mockHttp.NewMockHttpClient(mockCtrl)

	type fields struct {
		client  http.HttpClient
		baseUrl string
		apikey  string
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
			name:    "error get",
			fields:  fields{
				client:  mockClient,
				baseUrl: "",
				apikey:  "",
			},
			args:    args{
				omdbID: "1",
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				mockClient.EXPECT().GET(gomock.Any()).Return(nil, errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &movieRepo{
				client:  tt.fields.client,
				baseUrl: tt.fields.baseUrl,
				apikey:  tt.fields.apikey,
			}
			tt.mock()
			got, err := m.MovieDetail(tt.args.omdbID)
			if (err != nil) != tt.wantErr {
				t.Errorf("MovieDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovieDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
