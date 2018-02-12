package activedirectory

// Group represents an active directory group
type Group struct {
	DistinguishedName string
	GroupCategory     string
	GroupScope        string
	Name              string
	ObjectClass       string
	ObjectGUID        string
	SAMAccountName    string
	SID               string
}

const (
	// GroupScopeDomainLocal ...
	GroupScopeDomainLocal = 0
	// GroupScopeGlobal ...
	GroupScopeGlobal = 1
	// GroupScopeUniversal ...
	GroupScopeUniversal = 2
)

// ParseFromText takes a block of text and parses the resulting active directory group
func (group *Group) ParseFromText(text string) error {
	return fillStructFromPowershellOutput(group, text)
}
