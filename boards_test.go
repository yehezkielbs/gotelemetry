package gotelemetry

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBoards(t *testing.T) {

	board := Board{Name: "junitTestBoard", Theme: "dark", Display_board_name: true, Aspect_ratio: "HDTV"}
	credentials, _ := NewCredentials(api_key)

	Convey("Boards", t, func() {
		Convey("Should return status 201 when a Board is created", func() {
			results := board.CreateBoard(credentials)
			So(results.Error(), ShouldContainSubstring, "201")
		})

		Convey("Should return status 204 when a Board is deleted", func() {
			results := board.DeleteBoard(credentials)
			So(results.Error(), ShouldContainSubstring, "204")
		})
	})

}
