// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS989BFFCE_BB44_4232_AD47_587E230268AB.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_ExtensionEventSource struct
type RSOP_ExtensionEventSource struct {
	*cim.WmiInstance

	//
	eventLogName string

	//
	eventLogSource string

	//
	id string
}

func NewRSOP_ExtensionEventSourceEx1(instance *cim.WmiInstance) (newInstance *RSOP_ExtensionEventSource, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_ExtensionEventSource{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_ExtensionEventSourceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_ExtensionEventSource, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_ExtensionEventSource{
		WmiInstance: tmp,
	}
	return
}

// SeteventLogName sets the value of eventLogName for the instance
func (instance *RSOP_ExtensionEventSource) SetPropertyeventLogName(value string) (err error) {
	return instance.SetProperty("eventLogName", (value))
}

// GeteventLogName gets the value of eventLogName for the instance
func (instance *RSOP_ExtensionEventSource) GetPropertyeventLogName() (value string, err error) {
	retValue, err := instance.GetProperty("eventLogName")
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

// SeteventLogSource sets the value of eventLogSource for the instance
func (instance *RSOP_ExtensionEventSource) SetPropertyeventLogSource(value string) (err error) {
	return instance.SetProperty("eventLogSource", (value))
}

// GeteventLogSource gets the value of eventLogSource for the instance
func (instance *RSOP_ExtensionEventSource) GetPropertyeventLogSource() (value string, err error) {
	retValue, err := instance.GetProperty("eventLogSource")
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

// Setid sets the value of id for the instance
func (instance *RSOP_ExtensionEventSource) SetPropertyid(value string) (err error) {
	return instance.SetProperty("id", (value))
}

// Getid gets the value of id for the instance
func (instance *RSOP_ExtensionEventSource) GetPropertyid() (value string, err error) {
	retValue, err := instance.GetProperty("id")
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
