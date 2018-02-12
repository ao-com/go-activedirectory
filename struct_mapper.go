package activedirectory

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func setField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByNameFunc(func(key string) bool {
		return strings.ToLower(key) == strings.ToLower(name)
	})

	if !structFieldValue.IsValid() {
		return nil
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

func fillStruct(s interface{}, m map[string]interface{}) error {
	for k, v := range m {
		err := setField(s, k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func fillStructFromPowershellOutput(s interface{}, text string) error {
	structMap := make(map[string]interface{})
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if line == "\r" || line == "\n" || line == "\r\n" || line == "" {
			continue
		}

		lineSplit := strings.Split(line, ": ")
		lineEnd := lineSplit[1]
		lineEndTrimmed := strings.TrimRight(lineEnd, "\r\n")
		words := strings.Fields(line)
		structMap[words[0]] = lineEndTrimmed
	}

	return fillStruct(s, structMap)
}

func getTextBlocksFromPowershellOutput(text string) []string {
	textLines := strings.Split(text, "\n")
	firstLine := textLines[0]
	if isTextBlank(firstLine) {
		textLines = textLines[1 : len(textLines)-1]
	}

	lastLine := textLines[len(textLines)-1]
	if isTextBlank(lastLine) {
		textLines = textLines[:len(textLines)-2]
	}

	textBlocks := []string{}
	textBlock := ""
	for index, line := range textLines {
		isLastLine := index == len(textLines)-1
		if isTextBlank(line) {
			textBlocks = append(textBlocks, textBlock)
			textBlock = ""
			continue
		} else if isLastLine {
			textBlock += fmt.Sprintf("\n%s", line)
			textBlocks = append(textBlocks, textBlock)
			textBlock = ""
			continue
		}

		nextLine := ""
		if index+1 < len(textLines)-1 {
			nextLine = textLines[index+1]
		}

		if isTextBlank(nextLine) {
			textBlock += fmt.Sprintf("%s", line)
		} else {
			textBlock += fmt.Sprintf("%s\n", line)
		}
	}

	return textBlocks
}

// func fillStructsFromPowershellOutput(structs interface{}, text string) error {
// 	textLines := strings.Split(text, "\n")
// 	firstLine := textLines[0]
// 	if isTextBlank(firstLine) {
// 		textLines = textLines[1 : len(textLines)-1]
// 	}

// 	lastLine := textLines[len(textLines)-1]
// 	if isTextBlank(lastLine) {
// 		textLines = textLines[:len(textLines)-2]
// 	}

// 	textLines = textLines[1 : len(textLines)-1]
// 	textBlocks := []string{}
// 	textBlock := ""
// 	for index, line := range textLines {
// 		// textBlock += fmt.Sprintf("%s\n", line)
// 		// if we're the last line and the line isn't blank
// 		if isTextBlank(line) {
// 			textBlocks = append(textBlocks, textBlock)
// 			textBlock = ""
// 			continue
// 		}

// 		nextLine := ""
// 		if index+1 < len(textLines)-1 {
// 			nextLine = textLines[index+1]
// 		}

// 		if isTextBlank(nextLine) {
// 			textBlock += fmt.Sprintf("%s", line)
// 		} else {
// 			textBlock += fmt.Sprintf("%s\n", line)
// 		}
// 	}

// 	for _, block := range textBlocks {
// 		for _, l := range strings.Split(block, "\n") {
// 			fmt.Printf("l: %s\n", l)
// 		}

// 		var s interface{}

// 		t := reflect.TypeOf(GroupMember{}).Elem()
// 		ms := reflect.New(t).Elem().Interface()
// 		// works = &GroupMember{}

// 		err := fillStructFromPowershellOutput(&ms, block)
// 		if err != nil {
// 			return err
// 		}

// 		structs = append(structs.([]interface{}), s)
// 	}

// 	return nil
// }

func isTextBlank(t string) bool {
	if t == "\r" || t == "\n" || t == "\r\n" || t == "" {
		return true
	}

	return false
}
