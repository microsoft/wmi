// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_UserPrivilegeRight struct
type RSOP_UserPrivilegeRight struct {
	RSOP_SecuritySettings

	//
	AccountList []string

	//
	UserRight string
}

// SetAccountList sets the value of AccountList for the instance
func (instance *RSOP_UserPrivilegeRight) SetPropertyAccountList(value []string) (err error) {
	return instance.SetProperty("AccountList", value)
}

// GetAccountList gets the value of AccountList for the instance
func (instance *RSOP_UserPrivilegeRight) GetPropertyAccountList() (value []string, err error) {
	retValue, err := instance.GetProperty("AccountList")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserRight sets the value of UserRight for the instance
func (instance *RSOP_UserPrivilegeRight) SetPropertyUserRight(value string) (err error) {
	return instance.SetProperty("UserRight", value)
}

// GetUserRight gets the value of UserRight for the instance
func (instance *RSOP_UserPrivilegeRight) GetPropertyUserRight() (value string, err error) {
	retValue, err := instance.GetProperty("UserRight")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
