// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS0177EE87_7BD5_4C68_920A_6D1F4A019AF1
//////////////////////////////////////////////
package ns0177ee87_7bd5_4c68_920a_6d1f4a019af1

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_ExtendedStatus struct
type MSFT_ExtendedStatus struct {
	*MSFT_WmiError

	//
	original_error interface{}
}

func NewMSFT_ExtendedStatusEx1(instance *cim.WmiInstance) (newInstance *MSFT_ExtendedStatus, err error) {
	tmp, err := NewMSFT_WmiErrorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExtendedStatus{
		MSFT_WmiError: tmp,
	}
	return
}

func NewMSFT_ExtendedStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ExtendedStatus, err error) {
	tmp, err := NewMSFT_WmiErrorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ExtendedStatus{
		MSFT_WmiError: tmp,
	}
	return
}

// Setoriginal_error sets the value of original_error for the instance
func (instance *MSFT_ExtendedStatus) SetPropertyoriginal_error(value interface{}) (err error) {
	return instance.SetProperty("original_error", (value))
}

// Getoriginal_error gets the value of original_error for the instance
func (instance *MSFT_ExtendedStatus) GetPropertyoriginal_error() (value interface{}, err error) {
	retValue, err := instance.GetProperty("original_error")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}
