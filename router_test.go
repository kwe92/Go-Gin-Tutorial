package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAlbums(t *testing.T) {

	// create a request to desired endpoint
	req, err := http.NewRequest("GET", "/albums", nil)

	// check if there was an error generating the request
	checkErr(t, err)

	// create a ResponseRecorder to act as a ResponseWriter
	res := httptest.NewRecorder()

	// setup router
	router := setupGinRouter()

	// match request URL to a pattern of a registered handler
	// excute the handler, writing to the response body
	router.ServeHTTP(res, req)

	var actual []Album

	// write the response body bytes to a GO data structure
	json.NewDecoder(res.Body).Decode(&actual)

	assert.Equal(t, 200, res.Code)

	assert.Equal(t, albumsList, actual)
}

func TestPostAlbums(t *testing.T) {

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
				expected := Album{
					ID:     "1004",
					Title:  "The Modern Sound of Betty Carter",
					Artist: "Betty Carter",
					Price:  49.99,
				}

				// create buffer with io.Read and io.Write methods representing the reqest body
				var reqBody bytes.Buffer

				// write expected data to request body buffer
				json.NewEncoder(&reqBody).Encode(expected)

				// create a request to desired endpoint with request body
				req, err := http.NewRequest("POST", "/albums", &reqBody)

				fmt.Println("Test Post Albums request body:", req.Body)

				checkErr(t, err)

				return args{
					req: req,
				}

			},

			wantCode: http.StatusOK,
			wantBody: Album{
				ID:     "1004",
				Title:  "The Modern Sound of Betty Carter",
				Artist: "Betty Carter",
				Price:  49.99,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			response := httptest.NewRecorder()

			router := setupGinRouter()

			router.ServeHTTP(response, tArgs.req)

			var responseData []Album

			json.NewDecoder(response.Body).Decode(&responseData)

			actual := responseData[len(responseData)-1]

			fmt.Println("Test Post Albums expected:", actual)

			assert.Equal(t, http.StatusOK, response.Code)

			assert.Equal(t, tt.wantBody, actual)

		})
	}

}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}

}
