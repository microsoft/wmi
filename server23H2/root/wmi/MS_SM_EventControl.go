// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MS_SM_EventControl struct
type MS_SM_EventControl struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string
}

func NewMS_SM_EventControlEx1(instance *cim.WmiInstance) (newInstance *MS_SM_EventControl, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MS_SM_EventControl{
		WmiInstance: tmp,
	}
	return
}

func NewMS_SM_EventControlEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MS_SM_EventControl, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MS_SM_EventControl{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MS_SM_EventControl) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MS_SM_EventControl) GetPropertyActive() (value bool, err error) {
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
func (instance *MS_SM_EventControl) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MS_SM_EventControl) GetPropertyInstanceName() (value string, err error) {
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

//

// <param name="AllTargets" type="uint32 "></param>
// <param name="DiscoveredPortWWN" type="uint8 []"></param>
// <param name="DomainPortWWN" type="uint8 []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MS_SM_EventControl) SM_AddTarget( /* IN */ HbaPortWWN []uint8,
	/* IN */ DiscoveredPortWWN []uint8,
	/* IN */ DomainPortWWN []uint8,
	/* IN */ AllTargets uint32,
	/* OUT */ HBAStatus uint32) (err error) {
	_, err = instance.InvokeMethod("SM_AddTarget", HbaPortWWN, DiscoveredPortWWN, DomainPortWWN, AllTargets)
	if err != nil {
		return
	}
	return

}

//

// <param name="AllTargets" type="uint32 "></param>
// <param name="DiscoveredPortWWN" type="uint8 []"></param>
// <param name="DomainPortWWN" type="uint8 []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MS_SM_EventControl) SM_RemoveTarget( /* IN */ HbaPortWWN []uint8,
	/* IN */ DiscoveredPortWWN []uint8,
	/* IN */ DomainPortWWN []uint8,
	/* IN */ AllTargets uint32,
	/* OUT */ HBAStatus uint32) (err error) {
	_, err = instance.InvokeMethod("SM_RemoveTarget", HbaPortWWN, DiscoveredPortWWN, DomainPortWWN, AllTargets)
	if err != nil {
		return
	}
	return

}

//

// <param name="EventType" type="uint32 "></param>
// <param name="PortWWN" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MS_SM_EventControl) SM_AddPort( /* IN */ PortWWN []uint8,
	/* IN */ EventType uint32,
	/* OUT */ HBAStatus uint32) (err error) {
	_, err = instance.InvokeMethod("SM_AddPort", PortWWN, EventType)
	if err != nil {
		return
	}
	return

}

//

// <param name="EventType" type="uint32 "></param>
// <param name="PortWWN" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MS_SM_EventControl) SM_RemovePort( /* IN */ PortWWN []uint8,
	/* IN */ EventType uint32,
	/* OUT */ HBAStatus uint32) (err error) {
	_, err = instance.InvokeMethod("SM_RemovePort", PortWWN, EventType)
	if err != nil {
		return
	}
	return

}

//

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MS_SM_EventControl) SM_AddLink( /* OUT */ HBAStatus uint32) (err error) {
	_, err = instance.InvokeMethod("SM_AddLink")
	if err != nil {
		return
	}
	return

}

//

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MS_SM_EventControl) SM_RemoveLink( /* OUT */ HBAStatus uint32) (err error) {
	_, err = instance.InvokeMethod("SM_RemoveLink")
	if err != nil {
		return
	}
	return

}
