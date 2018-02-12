package activedirectory

// GroupMember ...
type GroupMember struct {
	DistinguishedName string
	Name              string
	ObjectClass       string
	ObjectGUID        string
	SAMAccountName    string
	SID               string
}

// GroupMembers ...
type GroupMembers []GroupMember

// ParseFromText takes a block of text and parses the resulting active directory group member
func (member *GroupMember) ParseFromText(text string) error {
	return fillStructFromPowershellOutput(member, text)
}

// ParseFromText takes a block of text and parses the resulting active directory group members
func (members *GroupMembers) ParseFromText(text string) error {
	textBlocks := getTextBlocksFromPowershellOutput(text)

	for _, textBlock := range textBlocks {
		member := GroupMember{}
		err := member.ParseFromText(textBlock)
		if err != nil {
			return err
		}

		*members = append(*members, member)
	}

	return nil
}
