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

		credentials, _ := NewCredentials(getTestKey())

		Convey("Should successfully create, retrieve, and delete a Widget", func() {
			w, err := NewWidget(credentials, board, "value", 1, 1, 1, 1, 0, "default")

			So(err, ShouldBeNil)
			So(w, ShouldNotBeNil)
			So(w.Id, ShouldNotBeEmpty)
			So(len(w.FlowIds), ShouldBeGreaterThan, 0)
			So(w.BoardId, ShouldNotBeEmpty)

			w2, err := GetWidget(credentials, w.Id)

			So(err, ShouldBeNil)
			So(w2, ShouldNotBeNil)
			So(w2.Id, ShouldEqual, w.Id)

			err = w2.Delete()

			So(err, ShouldBeNil)

			w3, err := GetWidget(credentials, w.Id)

			So(w3, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(err.(*Error).StatusCode, ShouldEqual, 404)
		})
	})

}
