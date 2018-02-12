package activedirectory

import (
	"fmt"
	"strings"

	ps "github.com/gorillalabs/go-powershell"
	"github.com/gorillalabs/go-powershell/backend"
)

// Client for active directory
type Client struct {
	username           string
	password           string
	credentialsCommand string
}

// NewClient creates a new active directory client
func NewClient(username string, password string) Client {
	securePasswordCommand := fmt.Sprintf("$securePass = ConvertTo-SecureString \"%s\" -AsPlainText -Force", password)
	credentialsCommand := fmt.Sprintf("$credentials = New-Object System.Management.Automation.PSCredential (\"%s\", $securePass)", username)

	return Client{
		username:           username,
		password:           password,
		credentialsCommand: fmt.Sprintf("%s\n%s\n", securePasswordCommand, credentialsCommand),
	}
}

// AddADGroupMember adds an active directory group member to a group
func (client Client) AddADGroupMember(name string, member string) error {
	back := &backend.Local{}
	shell, err := ps.New(back)
	if err != nil {
		return err
	}

	defer shell.Exit()
	cmd := fmt.Sprintf("%s\n", client.credentialsCommand)
	cmd += fmt.Sprintf("Add-ADGroupMember -Identity \"%s\" -Members \"%s\" -Credential $credentials", name, member)
	_, _, err = shell.Execute(cmd)
	if err != nil {
		return err
	}

	return nil
}

// GetADGroup return an active directory group
func (client Client) GetADGroup(name string) (*Group, error) {
	back := &backend.Local{}
	shell, err := ps.New(back)
	if err != nil {
		return nil, err
	}

	defer shell.Exit()
	cmd := fmt.Sprintf("%s\n", client.credentialsCommand)
	cmd += fmt.Sprintf("Get-ADGroup \"%s\" -Credential $credentials", name)
	stout, _, err := shell.Execute(cmd)
	if err != nil {
		if strings.Contains(err.Error(), "Cannot find an object with identity") {
			return nil, nil
		}

		return nil, err
	}

	stout = strings.TrimLeft(stout, "\r\n")
	stout = strings.TrimRight(stout, "\r\n")
	group := Group{}
	group.ParseFromText(stout)
	return &group, nil
}

// GetADGroupMembers ...
func (client Client) GetADGroupMembers(name string) (GroupMembers, error) {
	back := &backend.Local{}
	shell, err := ps.New(back)
	if err != nil {
		return nil, err
	}

	defer shell.Exit()
	cmd := fmt.Sprintf("%s\n", client.credentialsCommand)
	cmd += fmt.Sprintf("Get-ADGroupMember -Identity \"%s\" -Credential $credentials", name)
	stout, _, err := shell.Execute(cmd)
	if err != nil {
		return nil, err
	}

	stout = strings.TrimLeft(stout, "\r\n")
	stout = strings.TrimRight(stout, "\r\n")
	groupmembers := GroupMembers{}
	groupmembers.ParseFromText(stout)
	return groupmembers, nil
}

// NewADGroup creates an active directory group
func (client Client) NewADGroup(name string, groupScope int, path string) error {
	back := &backend.Local{}
	shell, err := ps.New(back)
	if err != nil {
		return err
	}

	defer shell.Exit()
	cmd := fmt.Sprintf("%s\n", client.credentialsCommand)
	cmd += fmt.Sprintf("New-ADGroup -Name \"%s\" -GroupScope %d -Credential $credentials", name, groupScope)
	if path != "" {
		cmd += fmt.Sprintf(" -Path \"%s\"", path)
	}

	_, _, err = shell.Execute(cmd)
	if err != nil {
		return err
	}

	return nil
}

// RemoveADGroup removes an active directory group
func (client Client) RemoveADGroup(name string) error {
	back := &backend.Local{}
	shell, err := ps.New(back)
	if err != nil {
		return err
	}

	defer shell.Exit()
	cmd := fmt.Sprintf("%s\n", client.credentialsCommand)
	cmd += fmt.Sprintf("Remove-ADGroup -Identity \"%s\" -Confirm:$false -Credential $credentials", name)
	_, _, err = shell.Execute(cmd)
	if err != nil {
		return err
	}

	return nil
}
