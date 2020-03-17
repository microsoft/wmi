// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_RestrictedGroupEx struct
type RSOP_RestrictedGroupEx struct {
	RSOP_SecuritySettings

	//
	GroupName string

	//
	Members []string

	//
	MembersOf []string
}

// SetGroupName sets the value of GroupName for the instance
func (instance *RSOP_RestrictedGroupEx) SetPropertyGroupName(value string) (err error) {
	return instance.SetProperty("GroupName", value)
}

// GetGroupName gets the value of GroupName for the instance
func (instance *RSOP_RestrictedGroupEx) GetPropertyGroupName() (value string, err error) {
	retValue, err := instance.GetProperty("GroupName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMembers sets the value of Members for the instance
func (instance *RSOP_RestrictedGroupEx) SetPropertyMembers(value []string) (err error) {
	return instance.SetProperty("Members", value)
}

// GetMembers gets the value of Members for the instance
func (instance *RSOP_RestrictedGroupEx) GetPropertyMembers() (value []string, err error) {
	retValue, err := instance.GetProperty("Members")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMembersOf sets the value of MembersOf for the instance
func (instance *RSOP_RestrictedGroupEx) SetPropertyMembersOf(value []string) (err error) {
	return instance.SetProperty("MembersOf", value)
}

// GetMembersOf gets the value of MembersOf for the instance
func (instance *RSOP_RestrictedGroupEx) GetPropertyMembersOf() (value []string, err error) {
	retValue, err := instance.GetProperty("MembersOf")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
