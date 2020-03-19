// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Cluster.Scaleout
//////////////////////////////////////////////
package scaleout

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
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
	return instance.SetProperty("original_error", value)
}

// Getoriginal_error gets the value of original_error for the instance
func (instance *MSFT_ExtendedStatus) GetPropertyoriginal_error() (value interface{}, err error) {
	retValue, err := instance.GetProperty("original_error")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}
