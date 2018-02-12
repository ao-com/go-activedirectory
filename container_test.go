package activedirectory

import (
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestContainer(t *testing.T) {
	Convey("Given active directory container", t, func() {
		container := Container{}

		Convey("When parsing from text", func() {
			bytes, _ := ioutil.ReadFile("./samples/get-adgroup.txt")
			adGroupSample := string(bytes)
			err := container.ParseFromText(adGroupSample)

			So(err, ShouldBeNil)

			Convey("Should parse distinguished name correctly", func() {
				So(container.DistinguishedName, ShouldEqual, "CN=SomeGroup,OU=SomeOU Groups,DC=somedc,DC=local")
			})

			Convey("Should parse name correctly", func() {
				So(container.Name, ShouldEqual, "SomeGroup")
			})

			Convey("Should parse object class correctly", func() {
				So(container.ObjectClass, ShouldEqual, "group")
			})

			Convey("Should parse object guid correctly", func() {
				So(container.ObjectGUID, ShouldEqual, "8562e62c-ebd9-474f-a1d6-bdd79e4b678c")
			})
		})
	})
}
