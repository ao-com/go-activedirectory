package activedirectory

import (
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGroupMember(t *testing.T) {
	Convey("Given active directory group members", t, func() {
		members := GroupMembers{}

		Convey("When parsing from text", func() {
			bytes, _ := ioutil.ReadFile("./samples/get-adgroupmember.txt")
			adGroupMemberSample := string(bytes)
			members.ParseFromText(adGroupMemberSample)

			Convey("Should parse entires correctly", func() {
				So(len(members), ShouldEqual, 4)
				So(members[3].DistinguishedName, ShouldEqual, "CN=Some Person 4,OU=Some OU 1,OU=Some OU 2,OU=Some OU 3,OU=Some OU 4,OU=Some OU 5,DC=Some DC 1,DC=Some DC 2")
				So(members[3].Name, ShouldEqual, "Some Person 4")
				So(members[3].ObjectClass, ShouldEqual, "user")
				So(members[3].ObjectGUID, ShouldEqual, "0c110996-40de-4060-9927-ec3517d350b2")
				So(members[3].SAMAccountName, ShouldEqual, "SPERSON4")
				So(members[3].SID, ShouldEqual, "S-0-0-00-0000000000-000000000-000000000-0000")
			})
		})
	})
}
