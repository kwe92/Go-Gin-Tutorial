package router

import (
	"bytes"
	"encoding/json"
	"example/web-service-gin/model"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAlbums(t *testing.T) {

	// create request to endpoint
	req, err := http.NewRequest("GET", "/albums", nil)

	// check error generating request
	checkErr(t, err)

	// create ResponseRecorder acting as ResponseWriter
	res := httptest.NewRecorder()

	// setup router
	router := SetupGinRouter()

	// match request URL to pattern of registered handler
	// execute handler, writing to response body
	router.ServeHTTP(res, req)

	var actual []model.Album

	// write response body bytes to GO data structure
	json.NewDecoder(res.Body).Decode(&actual)

	assert.Equal(t, 200, res.Code)

	assert.Equal(t, albumsList, actual)
}

func TestPostAlbums(t *testing.T) {

	router := SetupGinRouter()

	type args struct {
		req *http.Request
	}

	// Table Driven Test collection
	tests := []struct {
		name     string
		args     func(t *testing.T) args
		wantCode int
		wantBody any
	}{
		{
			name: "must return expected result for valid request body.",
			args: func(t *testing.T) args {
				expected := model.Album{
					ID:     "1004",
					Title:  "The Modern Sound of Betty Carter",
					Artist: "Betty Carter",
					Price:  49.99,
				}

				// create buffer with io.Read and io.Write methods representing reqest body
				var reqBody bytes.Buffer

				// write expected data to request body buffer
				json.NewEncoder(&reqBody).Encode(expected)

				// create a request to endpoint with request body
				req, err := http.NewRequest("POST", "/albums", &reqBody)

				checkErr(t, err)

				return args{
					req: req,
				}

			},

			wantCode: http.StatusOK,
			wantBody: []interface{}{map[string]interface{}{"artist": "John Coltrane", "id": "1001", "price": 56.99, "title": "Blue Train"}, map[string]interface{}{"artist": "Gerry Mulligan", "id": "1002", "price": 17.99, "title": "Jeru"}, map[string]interface{}{"artist": "Sarah Vaughan", "id": "1003", "price": 39.99, "title": "Sarah Vaughan and Clifford Brown"}, map[string]interface{}{"artist": "Betty Carter", "id": "1004", "price": 49.99, "title": "The Modern Sound of Betty Carter"}},
		},

		{
			name: "must return error message for invalid request body.",
			args: func(t *testing.T) args {

				// create buffer with io.Read and io.Write methods representing reqest body
				var reqBody bytes.Buffer

				// write expected data to request body buffer
				json.NewEncoder(&reqBody).Encode("unexpected value")

				// create a request to desired endpoint with request body
				req, err := http.NewRequest("POST", "/albums", &reqBody)

				checkErr(t, err)

				return args{
					req: req,
				}

			},

			wantCode: http.StatusBadRequest,
			wantBody: map[string]interface{}(map[string]interface{}{"error": "json: cannot unmarshal string into Go value of type model.Album"}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tArgs := tt.args(t)

			response := httptest.NewRecorder()

			router.ServeHTTP(response, tArgs.req)

			var responseData any

			json.NewDecoder(response.Body).Decode(&responseData)

			assert.Equal(t, tt.wantCode, response.Code)

			assert.Equal(t, tt.wantBody, responseData)

		})
	}

}

func TestGetAlbumById(t *testing.T) {

	router := SetupGinRouter()

	type args struct {
		req *http.Request
	}

	tests := []struct {
		name     string
		args     func(*testing.T) args
		wantCode int
		wantBody any
	}{
		{
			name: "given a valid album id, return the associated album data.",
			args: func(t *testing.T) args {

				req, err := http.NewRequest("GET", "/albums/1001", nil)

				checkErr(t, err)

				return args{
					req: req,
				}
			},
			wantCode: http.StatusOK,
			wantBody: map[string]interface{}(map[string]interface{}{"artist": "John Coltrane", "id": "1001", "price": 56.99, "title": "Blue Train"}),
		},
		{
			name: "given a invalid album id, return an error message.",
			args: func(t *testing.T) args {

				req, err := http.NewRequest("GET", "/albums/9999", nil)

				checkErr(t, err)

				return args{
					req: req,
				}
			},
			wantCode: http.StatusBadRequest,
			wantBody: map[string]interface{}(map[string]interface{}{"error": "could not locate an album with the id: 9999"}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tArgs := tt.args(t)

			response := httptest.NewRecorder()

			router.ServeHTTP(response, tArgs.req)

			var responseData any

			json.NewDecoder(response.Body).Decode(&responseData)

			fmt.Println("responseData from TestGetAlbumById:", responseData)

			assert.Equal(t, tt.wantCode, response.Code)

			assert.Equal(t, tt.wantBody, responseData)

		})
	}

}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err.Error())
	}

}
