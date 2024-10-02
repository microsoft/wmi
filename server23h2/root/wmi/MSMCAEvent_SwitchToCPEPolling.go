// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSMCAEvent_SwitchToCPEPolling struct
type MSMCAEvent_SwitchToCPEPolling struct {
	*WMIEvent

	//
	Active bool

	//
	InstanceName string
}

func NewMSMCAEvent_SwitchToCPEPollingEx1(instance *cim.WmiInstance) (newInstance *MSMCAEvent_SwitchToCPEPolling, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSMCAEvent_SwitchToCPEPolling{
		WMIEvent: tmp,
	}
	return
}

func NewMSMCAEvent_SwitchToCPEPollingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSMCAEvent_SwitchToCPEPolling, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSMCAEvent_SwitchToCPEPolling{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSMCAEvent_SwitchToCPEPolling) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSMCAEvent_SwitchToCPEPolling) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSMCAEvent_SwitchToCPEPolling) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSMCAEvent_SwitchToCPEPolling) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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
