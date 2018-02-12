package activedirectory

// Container ...
type Container struct {
	DistinguishedName string
	Name              string
	ObjectClass       string
	ObjectGUID        string
}

// ParseFromText takes a block of text and parses the resulting active directory container
func (container *Container) ParseFromText(text string) error {
	return fillStructFromPowershellOutput(container, text)
}
