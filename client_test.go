package activedirectory

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestClient(t *testing.T) {
	Convey("Given new client", t, func() {
		username := "workingusername"
		password := "workingpassword"
		client := NewClient(username, password)

		Convey("username should be set", func() {
			So(client.username, ShouldEqual, username)
		})

		Convey("password should be set", func() {
			So(client.password, ShouldEqual, password)
		})

		Convey("credentials command should be set correctly", func() {
			securePasswordCommand := fmt.Sprintf("$securePass = ConvertTo-SecureString \"%s\" -AsPlainText -Force", password)
			credentialsCommand := fmt.Sprintf("$credentials = New-Object System.Management.Automation.PSCredential (\"%s\", $securePass)", username)

			So(client.credentialsCommand, ShouldEqual, fmt.Sprintf("%s\n%s\n", securePasswordCommand, credentialsCommand))
		})

		Convey("NewADGroup should create a new active directory group", func() {
			err := client.NewADGroup("go-activedirectory-test", GroupScopeUniversal, "")

			So(err, ShouldBeNil)
		})

		Convey("AddADGroupMember should add a group to another group", func() {
			client.NewADGroup("go-activedirectory-test-2", GroupScopeUniversal, "")
			err := client.AddADGroupMember("go-activedirectory-test", "go-activedirectory-test-2")

			So(err, ShouldBeNil)
		})

		Convey("GetADGroupMembers should return the correct active directory group members", func() {
			members, err := client.GetADGroupMembers("go-activedirectory-test")

			So(err, ShouldBeNil)
			So(len(members), ShouldEqual, 1)
			So(members[0].Name, ShouldEqual, "go-activedirectory-test-2")
		})

		Convey("GetADGroup should return an active directory group", func() {
			group, err := client.GetADGroup("go-activedirectory-test")

			So(err, ShouldBeNil)
			So(group, ShouldNotBeNil)
		})

		Convey("GetADGroup should return nil when active directory group doesn't exist", func() {
			group, err := client.GetADGroup("thisdoesntexist")

			So(err, ShouldBeNil)
			So(group, ShouldBeNil)
		})

		Convey("RemoveADGroup should remove an active directory group", func() {
			err := client.RemoveADGroup("go-activedirectory-test")
			err = client.RemoveADGroup("go-activedirectory-test-2")

			So(err, ShouldBeNil)
		})

		Convey("GetADGroups should return some active directory groups", func() {
			groups, err := client.GetADGroups("", "")

			So(err, ShouldBeNil)
			So(groups, ShouldNotBeNil)
		})
	})

	Convey("Given new client with bad credentials", t, func() {
		username := "badusername"
		password := "badpassword"
		client := NewClient(username, password)

		Convey("NewADGroup should return rejected credentials error", func() {
			err := client.NewADGroup("go-activedirectory-test", GroupScopeUniversal, "")

			So(err, ShouldNotBeNil)
		})
	})
}
