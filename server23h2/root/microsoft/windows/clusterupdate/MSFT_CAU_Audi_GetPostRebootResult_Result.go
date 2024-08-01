// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.ClusterUpdate
//////////////////////////////////////////////
package clusterupdate

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_CAU_Audi_GetPostRebootResult_Result struct
type MSFT_CAU_Audi_GetPostRebootResult_Result struct {
	*cim.WmiInstance

	//
	HResult uint32

	//
	PostRebootHResult uint32
}

func NewMSFT_CAU_Audi_GetPostRebootResult_ResultEx1(instance *cim.WmiInstance) (newInstance *MSFT_CAU_Audi_GetPostRebootResult_Result, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_CAU_Audi_GetPostRebootResult_Result{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_CAU_Audi_GetPostRebootResult_ResultEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_CAU_Audi_GetPostRebootResult_Result, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_CAU_Audi_GetPostRebootResult_Result{
		WmiInstance: tmp,
	}
	return
}

// SetHResult sets the value of HResult for the instance
func (instance *MSFT_CAU_Audi_GetPostRebootResult_Result) SetPropertyHResult(value uint32) (err error) {
	return instance.SetProperty("HResult", (value))
}

// GetHResult gets the value of HResult for the instance
func (instance *MSFT_CAU_Audi_GetPostRebootResult_Result) GetPropertyHResult() (value uint32, err error) {
	retValue, err := instance.GetProperty("HResult")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPostRebootHResult sets the value of PostRebootHResult for the instance
func (instance *MSFT_CAU_Audi_GetPostRebootResult_Result) SetPropertyPostRebootHResult(value uint32) (err error) {
	return instance.SetProperty("PostRebootHResult", (value))
}

// GetPostRebootHResult gets the value of PostRebootHResult for the instance
func (instance *MSFT_CAU_Audi_GetPostRebootResult_Result) GetPropertyPostRebootHResult() (value uint32, err error) {
	retValue, err := instance.GetProperty("PostRebootHResult")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
