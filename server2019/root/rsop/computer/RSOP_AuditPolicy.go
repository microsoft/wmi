// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
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
	return instance.SetProperty("Category", value)
}

// GetCategory gets the value of Category for the instance
func (instance *RSOP_AuditPolicy) GetPropertyCategory() (value string, err error) {
	retValue, err := instance.GetProperty("Category")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFailure sets the value of Failure for the instance
func (instance *RSOP_AuditPolicy) SetPropertyFailure(value bool) (err error) {
	return instance.SetProperty("Failure", value)
}

// GetFailure gets the value of Failure for the instance
func (instance *RSOP_AuditPolicy) GetPropertyFailure() (value bool, err error) {
	retValue, err := instance.GetProperty("Failure")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSuccess sets the value of Success for the instance
func (instance *RSOP_AuditPolicy) SetPropertySuccess(value bool) (err error) {
	return instance.SetProperty("Success", value)
}

// GetSuccess gets the value of Success for the instance
func (instance *RSOP_AuditPolicy) GetPropertySuccess() (value bool, err error) {
	retValue, err := instance.GetProperty("Success")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
