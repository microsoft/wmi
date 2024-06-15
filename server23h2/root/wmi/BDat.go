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

// BDat struct
type BDat struct {
	*cim.WmiInstance

	// data
	Bytes []uint8
}

func NewBDatEx1(instance *cim.WmiInstance) (newInstance *BDat, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &BDat{
		WmiInstance: tmp,
	}
	return
}

func NewBDatEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BDat, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BDat{
		WmiInstance: tmp,
	}
	return
}

// SetBytes sets the value of Bytes for the instance
func (instance *BDat) SetPropertyBytes(value []uint8) (err error) {
	return instance.SetProperty("Bytes", (value))
}

// GetBytes gets the value of Bytes for the instance
func (instance *BDat) GetPropertyBytes() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Bytes")
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
