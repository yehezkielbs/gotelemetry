package gotelemetry

import (
	"crypto/rand"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestChannels(t *testing.T) {
	credentials, _ := NewCredentials(getTestKey())

	p := make([]byte, 10)
	rand.Read(p)

	name := fmt.Sprintf("board-%x", p)

	Convey("Board export", t, func() {
		Convey("Should properly export a board", func() {
			b, err := NewBoard(credentials, name, "dark", true, "HDTV")

			So(err, ShouldBeNil)
			So(b, ShouldNotBeNil)
			So(b.Id, ShouldNotBeNil)
			So(b.ChannelIds, ShouldNotBeEmpty)

			_, err = NewWidget(credentials, b, "value", 1, 2, 3, 4, 0, "normal")

			So(err, ShouldBeNil)

			_, err = NewWidget(credentials, b, "value", 1, 2, 3, 4, 0, "normal")

			So(err, ShouldBeNil)

			b1, err := GetBoard(credentials, b.Id)

			So(err, ShouldBeNil)
			So(b1, ShouldNotBeNil)

			exported, err := b1.Export()

			So(err, ShouldBeNil)
			So(exported, ShouldNotBeNil)

			b1.Delete()

			b3, err := ImportBoard(credentials, "test", exported)

			So(err, ShouldBeNil)
			So(b3, ShouldNotBeNil)

			b4, err := GetBoard(credentials, b3.Id)

			So(err, ShouldBeNil)

			So(b3.Name, ShouldEqual, "test"+b.Name)
			So(len(b3.Widgets), ShouldEqual, len(b4.Widgets))

			b3.Delete()
		})

		Convey("Should not import a board with a name that already exists", func() {

			newBoard, _ := NewBoard(credentials, name, "dark", true, "HDTV")

			_, _ = NewWidget(credentials, newBoard, "value", 1, 2, 3, 4, 0, "normal")
			_, _ = NewWidget(credentials, newBoard, "value", 1, 2, 3, 4, 0, "normal")

			getBoard, _ := GetBoard(credentials, newBoard.Id)

			exportedBoard, _ := getBoard.Export()

			importedBoard, err := ImportBoard(credentials, "fail_import", exportedBoard)
			importedBoardWithSameName, err := ImportBoard(credentials, "fail_import", exportedBoard)

			So(importedBoardWithSameName, ShouldBeNil)
			So(err, ShouldNotBeNil)

			// Clean up
			newBoard.Delete()
			importedBoard.Delete()

		})
	})

}
