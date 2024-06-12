// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetSwitchTeam_TeamMember struct
type MSFT_NetSwitchTeam_TeamMember struct {
	*cim.WmiInstance

	//
	MemberOfTheTeam MSFT_NetSwitchTeamMember

	//
	TeamOfTheMember MSFT_NetSwitchTeam
}

func NewMSFT_NetSwitchTeam_TeamMemberEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetSwitchTeam_TeamMember, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetSwitchTeam_TeamMember{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetSwitchTeam_TeamMemberEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetSwitchTeam_TeamMember, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetSwitchTeam_TeamMember{
		WmiInstance: tmp,
	}
	return
}

// SetMemberOfTheTeam sets the value of MemberOfTheTeam for the instance
func (instance *MSFT_NetSwitchTeam_TeamMember) SetPropertyMemberOfTheTeam(value MSFT_NetSwitchTeamMember) (err error) {
	return instance.SetProperty("MemberOfTheTeam", (value))
}

// GetMemberOfTheTeam gets the value of MemberOfTheTeam for the instance
func (instance *MSFT_NetSwitchTeam_TeamMember) GetPropertyMemberOfTheTeam() (value MSFT_NetSwitchTeamMember, err error) {
	retValue, err := instance.GetProperty("MemberOfTheTeam")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetSwitchTeamMember)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetSwitchTeamMember is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetSwitchTeamMember(valuetmp)

	return
}

// SetTeamOfTheMember sets the value of TeamOfTheMember for the instance
func (instance *MSFT_NetSwitchTeam_TeamMember) SetPropertyTeamOfTheMember(value MSFT_NetSwitchTeam) (err error) {
	return instance.SetProperty("TeamOfTheMember", (value))
}

// GetTeamOfTheMember gets the value of TeamOfTheMember for the instance
func (instance *MSFT_NetSwitchTeam_TeamMember) GetPropertyTeamOfTheMember() (value MSFT_NetSwitchTeam, err error) {
	retValue, err := instance.GetProperty("TeamOfTheMember")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetSwitchTeam)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetSwitchTeam is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetSwitchTeam(valuetmp)

	return
}
