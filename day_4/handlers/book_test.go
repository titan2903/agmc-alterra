package handlers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBooksSuccess(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	fmt.Println("============== rec", rec)
	fmt.Println("============== rec", rec.Code)
	c.SetPath("/v1/books")

	asserts := assert.New(t)

	// testing
	if asserts.NoError(h.GetAllBooks(c)) {
		fmt.Println("====================masuk")
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "code")
		asserts.Contains(body, "status")
		asserts.Contains(body, "message")
		// asserts.Contains(body, "data")
	}
}
