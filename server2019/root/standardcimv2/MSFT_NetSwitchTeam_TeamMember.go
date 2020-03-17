// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetSwitchTeam_TeamMember struct
type MSFT_NetSwitchTeam_TeamMember struct {
	cim.WmiInstance

	//
	MemberOfTheTeam MSFT_NetSwitchTeamMember

	//
	TeamOfTheMember MSFT_NetSwitchTeam
}

// SetMemberOfTheTeam sets the value of MemberOfTheTeam for the instance
func (instance *MSFT_NetSwitchTeam_TeamMember) SetPropertyMemberOfTheTeam(value MSFT_NetSwitchTeamMember) (err error) {
	return instance.SetProperty("MemberOfTheTeam", value)
}

// GetMemberOfTheTeam gets the value of MemberOfTheTeam for the instance
func (instance *MSFT_NetSwitchTeam_TeamMember) GetPropertyMemberOfTheTeam() (value MSFT_NetSwitchTeamMember, err error) {
	retValue, err := instance.GetProperty("MemberOfTheTeam")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetSwitchTeamMember)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTeamOfTheMember sets the value of TeamOfTheMember for the instance
func (instance *MSFT_NetSwitchTeam_TeamMember) SetPropertyTeamOfTheMember(value MSFT_NetSwitchTeam) (err error) {
	return instance.SetProperty("TeamOfTheMember", value)
}

// GetTeamOfTheMember gets the value of TeamOfTheMember for the instance
func (instance *MSFT_NetSwitchTeam_TeamMember) GetPropertyTeamOfTheMember() (value MSFT_NetSwitchTeam, err error) {
	retValue, err := instance.GetProperty("TeamOfTheMember")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetSwitchTeam)
	if !ok {
		// TODO: Set an error
	}
	return
}
