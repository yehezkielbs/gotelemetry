package gotelemetry

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// Constants
const board_id string = "542ef7b37765627da4090000"

func TestWidgets(t *testing.T) {

	Convey("Widgets", t, func() {
		board := &Board{Id: board_id}

		widget := Widget{"log", board.Id, 1, 1, 10, 10, 0, "default"}
		credentials, _ := NewCredentials(getTestKey())

		Convey("Should successfully create a Widget", func() {
			widget, err := NewWidget(credentials, board, "value", 1, 1, 1, 1, 0, "default")
			So(widget, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})

		Convey("Should return an invalid board_id error when the board does not exist", func() {
			testWidget := widget
			testWidget.BoardId = "I am not a valid board"
			err := testWidget.Save(credentials)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "400")
		})

		Convey("Should return a 400 when a invalid variant is passed", func() {
			testWidget := widget
			testWidget.Variant = "I am not a valid variant"
			err := testWidget.Save(credentials)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "400")
		})
	})

}
