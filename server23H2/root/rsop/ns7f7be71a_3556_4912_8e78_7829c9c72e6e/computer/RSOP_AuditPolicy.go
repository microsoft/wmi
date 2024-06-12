// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS7F7BE71A_3556_4912_8E78_7829C9C72E6E.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_AuditPolicy struct
type RSOP_AuditPolicy struct {
	*RSOP_SecuritySettings

	//
	Category string

	//
	Failure bool

	//
	Success bool
}

func NewRSOP_AuditPolicyEx1(instance *cim.WmiInstance) (newInstance *RSOP_AuditPolicy, err error) {
	tmp, err := NewRSOP_SecuritySettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_AuditPolicy{
		RSOP_SecuritySettings: tmp,
	}
	return
}

func NewRSOP_AuditPolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_AuditPolicy, err error) {
	tmp, err := NewRSOP_SecuritySettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_AuditPolicy{
		RSOP_SecuritySettings: tmp,
	}
	return
}

// SetCategory sets the value of Category for the instance
func (instance *RSOP_AuditPolicy) SetPropertyCategory(value string) (err error) {
	return instance.SetProperty("Category", (value))
}

// GetCategory gets the value of Category for the instance
func (instance *RSOP_AuditPolicy) GetPropertyCategory() (value string, err error) {
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
func (instance *RSOP_AuditPolicy) SetPropertyFailure(value bool) (err error) {
	return instance.SetProperty("Failure", (value))
}

// GetFailure gets the value of Failure for the instance
func (instance *RSOP_AuditPolicy) GetPropertyFailure() (value bool, err error) {
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
func (instance *RSOP_AuditPolicy) SetPropertySuccess(value bool) (err error) {
	return instance.SetProperty("Success", (value))
}

// GetSuccess gets the value of Success for the instance
func (instance *RSOP_AuditPolicy) GetPropertySuccess() (value bool, err error) {
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
