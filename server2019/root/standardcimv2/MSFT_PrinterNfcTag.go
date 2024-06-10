// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_PrinterNfcTag struct
type MSFT_PrinterNfcTag struct {
	*cim.WmiInstance

	//
	Locked bool

	//
	SharePath []string

	//
	WsdAddress []string
}

func NewMSFT_PrinterNfcTagEx1(instance *cim.WmiInstance) (newInstance *MSFT_PrinterNfcTag, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_PrinterNfcTag{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_PrinterNfcTagEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_PrinterNfcTag, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_PrinterNfcTag{
		WmiInstance: tmp,
	}
	return
}

// SetLocked sets the value of Locked for the instance
func (instance *MSFT_PrinterNfcTag) SetPropertyLocked(value bool) (err error) {
	return instance.SetProperty("Locked", (value))
}

// GetLocked gets the value of Locked for the instance
func (instance *MSFT_PrinterNfcTag) GetPropertyLocked() (value bool, err error) {
	retValue, err := instance.GetProperty("Locked")
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

// SetSharePath sets the value of SharePath for the instance
func (instance *MSFT_PrinterNfcTag) SetPropertySharePath(value []string) (err error) {
	return instance.SetProperty("SharePath", (value))
}

// GetSharePath gets the value of SharePath for the instance
func (instance *MSFT_PrinterNfcTag) GetPropertySharePath() (value []string, err error) {
	retValue, err := instance.GetProperty("SharePath")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetWsdAddress sets the value of WsdAddress for the instance
func (instance *MSFT_PrinterNfcTag) SetPropertyWsdAddress(value []string) (err error) {
	return instance.SetProperty("WsdAddress", (value))
}

// GetWsdAddress gets the value of WsdAddress for the instance
func (instance *MSFT_PrinterNfcTag) GetPropertyWsdAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("WsdAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
