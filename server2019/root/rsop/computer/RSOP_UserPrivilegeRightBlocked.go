// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.RSOP.Computer
//
// ////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_UserPrivilegeRightBlocked struct
type RSOP_UserPrivilegeRightBlocked struct {
	*RSOP_SecuritySettingsBlocked

	//
	AccountList []string

	//
	UserRight string
}

func NewRSOP_UserPrivilegeRightBlockedEx1(instance *cim.WmiInstance) (newInstance *RSOP_UserPrivilegeRightBlocked, err error) {
	tmp, err := NewRSOP_SecuritySettingsBlockedEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_UserPrivilegeRightBlocked{
		RSOP_SecuritySettingsBlocked: tmp,
	}
	return
}

func NewRSOP_UserPrivilegeRightBlockedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_UserPrivilegeRightBlocked, err error) {
	tmp, err := NewRSOP_SecuritySettingsBlockedEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_UserPrivilegeRightBlocked{
		RSOP_SecuritySettingsBlocked: tmp,
	}
	return
}

// SetAccountList sets the value of AccountList for the instance
func (instance *RSOP_UserPrivilegeRightBlocked) SetPropertyAccountList(value []string) (err error) {
	return instance.SetProperty("AccountList", (value))
}

// GetAccountList gets the value of AccountList for the instance
func (instance *RSOP_UserPrivilegeRightBlocked) GetPropertyAccountList() (value []string, err error) {
	retValue, err := instance.GetProperty("AccountList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetUserRight sets the value of UserRight for the instance
func (instance *RSOP_UserPrivilegeRightBlocked) SetPropertyUserRight(value string) (err error) {
	return instance.SetProperty("UserRight", (value))
}

// GetUserRight gets the value of UserRight for the instance
func (instance *RSOP_UserPrivilegeRightBlocked) GetPropertyUserRight() (value string, err error) {
	retValue, err := instance.GetProperty("UserRight")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
