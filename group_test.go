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

	Convey("Given active directory groups", t, func() {
		groups := Groups{}

		Convey("When parsing from text", func() {
			bytes, _ := ioutil.ReadFile("./samples/get-adgroups.txt")
			adGroupsSample := string(bytes)
			err := groups.ParseFromText(adGroupsSample)

			So(err, ShouldBeNil)

			Convey("Should parse entries correctly", func() {
				So(len(groups), ShouldEqual, 3)
				So(groups[2].DistinguishedName, ShouldEqual, "CN=SomeGroup3,OU=SomeOU Groups,DC=somedc,DC=local")
				So(groups[2].Name, ShouldEqual, "SomeGroup3")
				So(groups[2].ObjectClass, ShouldEqual, "group")
				So(groups[2].ObjectGUID, ShouldEqual, "214d570c-e4c8-4d13-b7e1-040c9a738750")
				So(groups[2].SAMAccountName, ShouldEqual, "SomeAccount3")
				So(groups[2].SID, ShouldEqual, "S-0-0-00-0000000000-000000000-000000000-00000")
				So(groups[2].GroupCategory, ShouldEqual, "Security")
				So(groups[2].GroupScope, ShouldEqual, "Universal")
			})
		})
	})

	Convey("Given active directory group with multiline DistinguishedName", t, func() {
		group := Group{}

		Convey("When parsing from text", func() {
			bytes, _ := ioutil.ReadFile("./samples/get-adgroups_multiline_dn.txt")
			adGroupsSample := string(bytes)
			err := group.ParseFromText(adGroupsSample)

			So(err, ShouldBeNil)

			Convey("Should parse entries correctly", func() {
				So(group.DistinguishedName, ShouldEqual, "CN=Some OU Administrators,OU=Some OU Location,OU=Some OU Groups,OU=Some OU Team Department,DC=domain,DC=local")
				So(group.Name, ShouldEqual, "Some OU Administrators")
				So(group.ObjectClass, ShouldEqual, "group")
				So(group.ObjectGUID, ShouldEqual, "efaf7005-ed81-40e2-b6bb-083e346333ef")
				So(group.SAMAccountName, ShouldEqual, "Some OU Administrators")
				So(group.SID, ShouldEqual, "S-1-5-21-0000000000-000000000-000000000-00000")
				So(group.GroupCategory, ShouldEqual, "Security")
				So(group.GroupScope, ShouldEqual, "Global")
			})
		})
	})
}
