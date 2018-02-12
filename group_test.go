package activedirectory

import (
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGroup(t *testing.T) {
	Convey("Given active directory group", t, func() {
		group := Group{}

		Convey("When parsing from text", func() {
			bytes, _ := ioutil.ReadFile("./samples/get-adgroup.txt")
			adGroupSample := string(bytes)
			err := group.ParseFromText(adGroupSample)

			So(err, ShouldBeNil)

			Convey("Should parse distinguished name correctly", func() {
				So(group.DistinguishedName, ShouldEqual, "CN=SomeGroup,OU=SomeOU Groups,DC=somedc,DC=local")
			})

			Convey("Should parse group category correctly", func() {
				So(group.GroupCategory, ShouldEqual, "Security")
			})

			Convey("Should parse group scope correctly", func() {
				So(group.GroupScope, ShouldEqual, "Universal")
			})

			Convey("Should parse name correctly", func() {
				So(group.Name, ShouldEqual, "SomeGroup")
			})

			Convey("Should parse object class correctly", func() {
				So(group.ObjectClass, ShouldEqual, "group")
			})

			Convey("Should parse object guid correctly", func() {
				So(group.ObjectGUID, ShouldEqual, "8562e62c-ebd9-474f-a1d6-bdd79e4b678c")
			})

			Convey("Should parse sam accout name correctly", func() {
				So(group.SAMAccountName, ShouldEqual, "SomeAccount")
			})

			Convey("Should parse sid correctly", func() {
				So(group.SID, ShouldEqual, "S-0-0-00-0000000000-000000000-000000000-00000")
			})
		})
	})
}
