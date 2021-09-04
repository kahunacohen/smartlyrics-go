package album_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	albumHandlers "github.com/kahunacohen/smartlyrics/handlers/album"
	albumModel "github.com/kahunacohen/smartlyrics/models/album"
	"github.com/kahunacohen/smartlyrics/routers"
)

func TestGetAll(t *testing.T) {
	r := routers.Get()
	r.GET("/api/v1/albums", albumHandlers.GetAll)
	req, _ := http.NewRequest("GET", "/api/v1/albums", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	var albums []albumModel.Album
	json.Unmarshal([]byte(resp.Body.Bytes()), &albums)
	wantedLen := 3
	gotLen := len(albums)
	if gotLen != wantedLen {
		t.Errorf("wanted album length: %v, got: %v", wantedLen, gotLen)
	}
	blueTrain := albums[0]
	wantedTitle := "Blue Train"
	gotTitle := blueTrain.Title
	if gotTitle != wantedTitle {
		t.Errorf("wanted first title to be: %v, got: %v", wantedTitle, gotTitle)
	}
}

func TestGet(t *testing.T) {
	r := routers.Get()
	r.GET("/api/v1/albums/1", albumHandlers.Get)
	req, _ := http.NewRequest("GET", "/api/v1/albums/1", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	var album albumModel.Album
	json.Unmarshal([]byte(resp.Body.Bytes()), &album)
	wantedTitle := "Blue Train"
	gotTitle := album.Title
	if gotTitle != wantedTitle {
		t.Errorf("wanted first title to be: %v, got: %v", wantedTitle, gotTitle)
	}
}
