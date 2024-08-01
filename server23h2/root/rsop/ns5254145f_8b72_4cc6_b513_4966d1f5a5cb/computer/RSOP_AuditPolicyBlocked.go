// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.NS5254145F_8B72_4CC6_B513_4966D1F5A5CB.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_AuditPolicyBlocked struct
type RSOP_AuditPolicyBlocked struct {
	*RSOP_SecuritySettingsBlocked

	//
	Category string

	//
	Failure bool

	//
	Success bool
}

func NewRSOP_AuditPolicyBlockedEx1(instance *cim.WmiInstance) (newInstance *RSOP_AuditPolicyBlocked, err error) {
	tmp, err := NewRSOP_SecuritySettingsBlockedEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_AuditPolicyBlocked{
		RSOP_SecuritySettingsBlocked: tmp,
	}
	return
}

func NewRSOP_AuditPolicyBlockedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_AuditPolicyBlocked, err error) {
	tmp, err := NewRSOP_SecuritySettingsBlockedEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_AuditPolicyBlocked{
		RSOP_SecuritySettingsBlocked: tmp,
	}
	return
}

// SetCategory sets the value of Category for the instance
func (instance *RSOP_AuditPolicyBlocked) SetPropertyCategory(value string) (err error) {
	return instance.SetProperty("Category", (value))
}

// GetCategory gets the value of Category for the instance
func (instance *RSOP_AuditPolicyBlocked) GetPropertyCategory() (value string, err error) {
	retValue, err := instance.GetProperty("Category")
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

// SetFailure sets the value of Failure for the instance
func (instance *RSOP_AuditPolicyBlocked) SetPropertyFailure(value bool) (err error) {
	return instance.SetProperty("Failure", (value))
}

// GetFailure gets the value of Failure for the instance
func (instance *RSOP_AuditPolicyBlocked) GetPropertyFailure() (value bool, err error) {
	retValue, err := instance.GetProperty("Failure")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetSuccess sets the value of Success for the instance
func (instance *RSOP_AuditPolicyBlocked) SetPropertySuccess(value bool) (err error) {
	return instance.SetProperty("Success", (value))
}

// GetSuccess gets the value of Success for the instance
func (instance *RSOP_AuditPolicyBlocked) GetPropertySuccess() (value bool, err error) {
	retValue, err := instance.GetProperty("Success")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}
