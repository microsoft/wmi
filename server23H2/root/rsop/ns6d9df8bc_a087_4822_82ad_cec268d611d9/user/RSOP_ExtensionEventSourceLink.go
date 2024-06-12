// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS6D9DF8BC_A087_4822_82AD_CEC268D611D9.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_ExtensionEventSourceLink struct
type RSOP_ExtensionEventSourceLink struct {
	*cim.WmiInstance

	//
	eventSource RSOP_ExtensionEventSource

	//
	extensionStatus RSOP_ExtensionStatus
}

func NewRSOP_ExtensionEventSourceLinkEx1(instance *cim.WmiInstance) (newInstance *RSOP_ExtensionEventSourceLink, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_ExtensionEventSourceLink{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_ExtensionEventSourceLinkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_ExtensionEventSourceLink, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_ExtensionEventSourceLink{
		WmiInstance: tmp,
	}
	return
}

// SeteventSource sets the value of eventSource for the instance
func (instance *RSOP_ExtensionEventSourceLink) SetPropertyeventSource(value RSOP_ExtensionEventSource) (err error) {
	return instance.SetProperty("eventSource", (value))
}

// GeteventSource gets the value of eventSource for the instance
func (instance *RSOP_ExtensionEventSourceLink) GetPropertyeventSource() (value RSOP_ExtensionEventSource, err error) {
	retValue, err := instance.GetProperty("eventSource")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RSOP_ExtensionEventSource)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RSOP_ExtensionEventSource is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RSOP_ExtensionEventSource(valuetmp)

	return
}

// SetextensionStatus sets the value of extensionStatus for the instance
func (instance *RSOP_ExtensionEventSourceLink) SetPropertyextensionStatus(value RSOP_ExtensionStatus) (err error) {
	return instance.SetProperty("extensionStatus", (value))
}

// GetextensionStatus gets the value of extensionStatus for the instance
func (instance *RSOP_ExtensionEventSourceLink) GetPropertyextensionStatus() (value RSOP_ExtensionStatus, err error) {
	retValue, err := instance.GetProperty("extensionStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RSOP_ExtensionStatus)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RSOP_ExtensionStatus is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RSOP_ExtensionStatus(valuetmp)

	return
}
