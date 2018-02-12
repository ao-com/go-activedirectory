package activedirectory

import (
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStructMapper(t *testing.T) {
	Convey("Given map", t, func() {
		testMap := map[string]interface{}{
			"SomeKeyOne": "somevalueone",
			"somekeytwo": "somevaluetwo",
		}

		Convey("Should be mapped to struct correctly", func() {
			testStruct := &struct {
				SomeKeyOne string
				SomeKeyTwo string
			}{}
			err := fillStruct(testStruct, testMap)

			So(err, ShouldBeNil)
			So(testStruct.SomeKeyOne, ShouldEqual, testMap["SomeKeyOne"])
			So(testStruct.SomeKeyTwo, ShouldEqual, testMap["somekeytwo"])
		})
	})

	Convey("Given powershell output for single object", t, func() {
		bytes, _ := ioutil.ReadFile("./samples/get-adgroup.txt")
		sample := string(bytes)

		Convey("Should be mapped to struct correctly", func() {
			testStuct := &struct {
				DistinguishedName string
				SID               string
			}{}
			err := fillStructFromPowershellOutput(testStuct, sample)

			So(err, ShouldBeNil)
			So(testStuct.DistinguishedName, ShouldEqual, "CN=SomeGroup,OU=SomeOU Groups,DC=somedc,DC=local")
			So(testStuct.SID, ShouldEqual, "S-0-0-00-0000000000-000000000-000000000-00000")
		})
	})

	Convey("Given powershell output for multiple objects", t, func() {
		bytes, _ := ioutil.ReadFile("./samples/get-adgroupmember.txt")
		sample := string(bytes)

		Convey("Should return correct text blocks", func() {
			textBlocks := getTextBlocksFromPowershellOutput(sample)

			So(len(textBlocks), ShouldEqual, 4)
		})
	})
}
