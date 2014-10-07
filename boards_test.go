package gotelemetry

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBoards(t *testing.T) {
	Convey("Boards", t, func() {

		board := Board{"junitTestBoard", "dark", true, "HDTV"}
		credentials, _ := NewCredentials(api_key)

		Convey("Should return status 201 when a Board is created", func() {
			results := board.CreateBoard(credentials)
			So(results.Error(), ShouldContainSubstring, "201")
		})

		Convey("Should return status 200 when a Board is deleted", func() {
			// TODO so previous tests can run indefintely
		})
	})

}
