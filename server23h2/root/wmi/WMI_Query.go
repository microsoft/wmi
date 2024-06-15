// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
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

// WMI_Query struct
type WMI_Query struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string

	// BIOS WMI info
	QDATA QDat
}

func NewWMI_QueryEx1(instance *cim.WmiInstance) (newInstance *WMI_Query, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &WMI_Query{
		WmiInstance: tmp,
	}
	return
}

func NewWMI_QueryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WMI_Query, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WMI_Query{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *WMI_Query) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *WMI_Query) GetPropertyActive() (value bool, err error) {
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
func (instance *WMI_Query) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *WMI_Query) GetPropertyInstanceName() (value string, err error) {
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

// SetQDATA sets the value of QDATA for the instance
func (instance *WMI_Query) SetPropertyQDATA(value QDat) (err error) {
	return instance.SetProperty("QDATA", (value))
}

// GetQDATA gets the value of QDATA for the instance
func (instance *WMI_Query) GetPropertyQDATA() (value QDat, err error) {
	retValue, err := instance.GetProperty("QDATA")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(QDat)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " QDat is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = QDat(valuetmp)

	return
}
