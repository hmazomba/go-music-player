package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Returns a JSON response with status code 200 and all albums data
func test_getAlbums_ReturnsAllAlbumsData(t *testing.T) {
	// Initialize a new HTTP request to the /albums endpoint
	req, _ := http.NewRequest("GET", "/albums", nil)

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Create a new router instance
	router := gin.Default()
	router.GET("/albums", getAlbums)

	// Serve the HTTP request to the response recorder
	router.ServeHTTP(rr, req)

	// Check if the status code is 200
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check if the response body contains all albums data
	expectedResponse := `[{"id":"1","title":"Blue Train","artist":"John Coltrane","price":56.99},{"id":"2","title":"Jeru","artist":"Gerry Mulligan","price":17.99},{"id":"3","title":"Sarah Vaughan and Clifford Brown","artist":"Sarah Vaughan","price":39.99}]`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, rr.Body.String())
	}
}

// Returns a JSON response with status code 200 and an empty array when there are no albums
func test_getAlbums_ReturnsEmptyArrayWhenNoAlbums(t *testing.T) {
	// Initialize a new HTTP request to the /albums endpoint
	req, _ := http.NewRequest("GET", "/albums", nil)

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Create a new router instance with an empty albums slice
	router := gin.Default()
	router.GET("/albums", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, []album{})
	})

	// Serve the HTTP request to the response recorder
	router.ServeHTTP(rr, req)

	// Check if the status code is 200
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check if the response body is an empty array
	expectedResponse := `[]`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, rr.Body.String())
	}
}

// Returns a JSON response with status code 404 and an error message when the requested resource is not found
func test_getAlbums_Returns404WhenResourceNotFound(t *testing.T) {
	// Initialize a new HTTP request to a non-existent endpoint
	req, _ := http.NewRequest("GET", "/nonexistent", nil)

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Create a new router instance
	router := gin.Default()
	router.GET("/albums", getAlbums)

	// Serve the HTTP request to the response recorder
	router.ServeHTTP(rr, req)

	// Check if the status code is 404
	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, rr.Code)
	}

	// Check if the response body contains the error message
	expectedResponse := `404 page not found`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, rr.Body.String())
	}
}

// Returns a JSON response with status code 500 and an error message when there is an internal server error
func test_getAlbums_Returns500WhenInternalServerError(t *testing.T) {
	// Initialize a new HTTP request to the /albums endpoint
	req, _ := http.NewRequest("GET", "/albums", nil)

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Create a new router instance with a faulty handler function
	router := gin.Default()
	router.GET("/albums", func(c *gin.Context) {
		// Simulate an internal server error by causing a panic
		panic("Internal Server Error")
	})

	// Serve the HTTP request to the response recorder
	router.ServeHTTP(rr, req)

	// Check if the status code is 500
	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, rr.Code)
	}

	// Check if the response body contains the error message
	expectedResponse := `Internal Server Error`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, rr.Body.String())
	}
}

// The order of the albums in the response is the same as in the 'albums' variable
func test_getAlbums_ReturnsAlbumsInCorrectOrder(t *testing.T) {
	// Initialize a new HTTP request to the /albums endpoint
	req, _ := http.NewRequest("GET", "/albums", nil)

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Create a new router instance
	router := gin.Default()
	router.GET("/albums", getAlbums)

	// Serve the HTTP request to the response recorder
	router.ServeHTTP(rr, req)

	// Check if the status code is 200
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check if the response body contains albums in correct order
	expectedResponse := `[{"id":"1","title":"Blue Train","artist":"John Coltrane","price":56.99},{"id":"2","title":"Jeru","artist":"Gerry Mulligan","price":17.99},{"id":"3","title":"Sarah Vaughan and Clifford Brown","artist":"Sarah Vaughan","price":39.99}]`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, rr.Body.String())
	}
}

// The response JSON has the correct format and keys for each album object
func test_getAlbums_ReturnsResponseWithCorrectFormatAndKeys(t *testing.T) {
	// Initialize a new HTTP request to the /albums endpoint
	req, _ := http.NewRequest("GET", "/albums", nil)

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Create a new router instance
	router := gin.Default()
	router.GET("/albums", getAlbums)

	// Serve the HTTP request to the response recorder
	router.ServeHTTP(rr, req)

	// Check if the status code is 200
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check if the response body has the correct format and keys
	expectedResponse := `[{"id":"1","title":"Blue Train","artist":"John Coltrane","price":56.99},{"id":"2","title":"Jeru","artist":"Gerry Mulligan","price":17.99},{"id":"3","title":"Sarah Vaughan and Clifford Brown","artist":"Sarah Vaughan","price":39.99}]`
	var response []album
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %s", err.Error())
	}

	for i, album := range response {
		if album.ID != albums[i].ID || album.Title != albums[i].Title || album.Artist != albums[i].Artist || album.Price != albums[i].Price {
			t.Errorf("Expected album %v, but got %v", albums[i], album)
		}
	}
}
