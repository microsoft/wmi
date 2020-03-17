// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetSwitchTeam struct
type MSFT_NetSwitchTeam struct {
	MSFT_NetImPlatTeam
}

//

// <param name="Name" type="string "></param>
// <param name="TeamMembers" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetSwitchTeam) Create( /* IN */ Name string,
	/* IN */ TeamMembers []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Create", Name, TeamMembers)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Name" type="string "></param>
// <param name="NewName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetSwitchTeam) Rename( /* IN */ Name string,
	/* IN */ NewName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Rename", Name, NewName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Name" type="string "></param>
// <param name="Team" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetSwitchTeam) AddMember( /* IN */ Name string,
	/* IN */ Team string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddMember", Name, Team)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Name" type="string "></param>
// <param name="Team" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetSwitchTeam) RemoveMember( /* IN */ Name string,
	/* IN */ Team string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveMember", Name, Team)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
