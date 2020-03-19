// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
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
	return instance.SetProperty("eventSource", value)
}

// GeteventSource gets the value of eventSource for the instance
func (instance *RSOP_ExtensionEventSourceLink) GetPropertyeventSource() (value RSOP_ExtensionEventSource, err error) {
	retValue, err := instance.GetProperty("eventSource")
	if err != nil {
		return
	}
	value, ok := retValue.(RSOP_ExtensionEventSource)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetextensionStatus sets the value of extensionStatus for the instance
func (instance *RSOP_ExtensionEventSourceLink) SetPropertyextensionStatus(value RSOP_ExtensionStatus) (err error) {
	return instance.SetProperty("extensionStatus", value)
}

// GetextensionStatus gets the value of extensionStatus for the instance
func (instance *RSOP_ExtensionEventSourceLink) GetPropertyextensionStatus() (value RSOP_ExtensionStatus, err error) {
	retValue, err := instance.GetProperty("extensionStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(RSOP_ExtensionStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}
