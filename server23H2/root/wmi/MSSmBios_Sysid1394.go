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

// MSSmBios_Sysid1394 struct
type MSSmBios_Sysid1394 struct {
	*MS_SmBios

	//
	x1394 []uint8
}

func NewMSSmBios_Sysid1394Ex1(instance *cim.WmiInstance) (newInstance *MSSmBios_Sysid1394, err error) {
	tmp, err := NewMS_SmBiosEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSSmBios_Sysid1394{
		MS_SmBios: tmp,
	}
	return
}

func NewMSSmBios_Sysid1394Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSSmBios_Sysid1394, err error) {
	tmp, err := NewMS_SmBiosEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSSmBios_Sysid1394{
		MS_SmBios: tmp,
	}
	return
}

// Setx1394 sets the value of x1394 for the instance
func (instance *MSSmBios_Sysid1394) SetPropertyx1394(value []uint8) (err error) {
	return instance.SetProperty("x1394", (value))
}

// Getx1394 gets the value of x1394 for the instance
func (instance *MSSmBios_Sysid1394) GetPropertyx1394() (value []uint8, err error) {
	retValue, err := instance.GetProperty("x1394")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}
