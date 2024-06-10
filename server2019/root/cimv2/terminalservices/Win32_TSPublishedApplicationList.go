// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_TSPublishedApplicationList struct
type Win32_TSPublishedApplicationList struct {
	*CIM_LogicalElement

	// Whether the Terminal Server restricts remote applications to those on the list.
	Disabled bool

	// Indicates whether the property Disabled is configured by Server (0),Group Policy (1), Default (2)
	PolicySourceDisabled uint32
}

func NewWin32_TSPublishedApplicationListEx1(instance *cim.WmiInstance) (newInstance *Win32_TSPublishedApplicationList, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TSPublishedApplicationList{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewWin32_TSPublishedApplicationListEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TSPublishedApplicationList, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TSPublishedApplicationList{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetDisabled sets the value of Disabled for the instance
func (instance *Win32_TSPublishedApplicationList) SetPropertyDisabled(value bool) (err error) {
	return instance.SetProperty("Disabled", (value))
}

// GetDisabled gets the value of Disabled for the instance
func (instance *Win32_TSPublishedApplicationList) GetPropertyDisabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Disabled")
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

// SetPolicySourceDisabled sets the value of PolicySourceDisabled for the instance
func (instance *Win32_TSPublishedApplicationList) SetPropertyPolicySourceDisabled(value uint32) (err error) {
	return instance.SetProperty("PolicySourceDisabled", (value))
}

// GetPolicySourceDisabled gets the value of PolicySourceDisabled for the instance
func (instance *Win32_TSPublishedApplicationList) GetPropertyPolicySourceDisabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceDisabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
