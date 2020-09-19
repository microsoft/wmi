// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_RestrictedGroupExBlocked struct
type RSOP_RestrictedGroupExBlocked struct {
	*RSOP_SecuritySettingsBlocked

	//
	GroupName string

	//
	Members []string

	//
	MembersOf []string
}

func NewRSOP_RestrictedGroupExBlockedEx1(instance *cim.WmiInstance) (newInstance *RSOP_RestrictedGroupExBlocked, err error) {
	tmp, err := NewRSOP_SecuritySettingsBlockedEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_RestrictedGroupExBlocked{
		RSOP_SecuritySettingsBlocked: tmp,
	}
	return
}

func NewRSOP_RestrictedGroupExBlockedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_RestrictedGroupExBlocked, err error) {
	tmp, err := NewRSOP_SecuritySettingsBlockedEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_RestrictedGroupExBlocked{
		RSOP_SecuritySettingsBlocked: tmp,
	}
	return
}

// SetGroupName sets the value of GroupName for the instance
func (instance *RSOP_RestrictedGroupExBlocked) SetPropertyGroupName(value string) (err error) {
	return instance.SetProperty("GroupName", (value))
}

// GetGroupName gets the value of GroupName for the instance
func (instance *RSOP_RestrictedGroupExBlocked) GetPropertyGroupName() (value string, err error) {
	retValue, err := instance.GetProperty("GroupName")
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

// SetMembers sets the value of Members for the instance
func (instance *RSOP_RestrictedGroupExBlocked) SetPropertyMembers(value []string) (err error) {
	return instance.SetProperty("Members", (value))
}

// GetMembers gets the value of Members for the instance
func (instance *RSOP_RestrictedGroupExBlocked) GetPropertyMembers() (value []string, err error) {
	retValue, err := instance.GetProperty("Members")
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

// SetMembersOf sets the value of MembersOf for the instance
func (instance *RSOP_RestrictedGroupExBlocked) SetPropertyMembersOf(value []string) (err error) {
	return instance.SetProperty("MembersOf", (value))
}

// GetMembersOf gets the value of MembersOf for the instance
func (instance *RSOP_RestrictedGroupExBlocked) GetPropertyMembersOf() (value []string, err error) {
	retValue, err := instance.GetProperty("MembersOf")
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
