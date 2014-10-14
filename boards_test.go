package gotelemetry

import (
	"crypto/rand"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBoards(t *testing.T) {
	credentials, _ := NewCredentials(getTestKey())

	p := make([]byte, 10)
	rand.Read(p)

	name := fmt.Sprintf("board-%x", p)

	Convey("Boards", t, func() {
		Convey("Should properly create and delete boards", func() {
			b, err := NewBoard(credentials, name, "dark", true, "HDTV")

			So(err, ShouldBeNil)
			So(b, ShouldNotBeNil)

			err = b.DeleteBoard(credentials)
			So(err, ShouldBeNil)
		})
	})

}
