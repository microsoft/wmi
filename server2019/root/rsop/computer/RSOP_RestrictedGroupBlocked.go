// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_RestrictedGroupBlocked struct
type RSOP_RestrictedGroupBlocked struct {
	RSOP_SecuritySettingsBlocked

	//
	GroupName string

	//
	Members []string
}

// SetGroupName sets the value of GroupName for the instance
func (instance *RSOP_RestrictedGroupBlocked) SetPropertyGroupName(value string) (err error) {
	return instance.SetProperty("GroupName", value)
}

// GetGroupName gets the value of GroupName for the instance
func (instance *RSOP_RestrictedGroupBlocked) GetPropertyGroupName() (value string, err error) {
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
func (instance *RSOP_RestrictedGroupBlocked) SetPropertyMembers(value []string) (err error) {
	return instance.SetProperty("Members", value)
}

// GetMembers gets the value of Members for the instance
func (instance *RSOP_RestrictedGroupBlocked) GetPropertyMembers() (value []string, err error) {
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
