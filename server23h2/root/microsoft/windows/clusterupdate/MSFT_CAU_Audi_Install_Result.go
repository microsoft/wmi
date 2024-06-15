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

// MSFT_CAU_Audi_Install_Result struct
type MSFT_CAU_Audi_Install_Result struct {
	*cim.WmiInstance

	//
	AudiInstallHResult uint32

	//
	RebootRequired bool

	//
	UpdateId string
}

func NewMSFT_CAU_Audi_Install_ResultEx1(instance *cim.WmiInstance) (newInstance *MSFT_CAU_Audi_Install_Result, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_CAU_Audi_Install_Result{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_CAU_Audi_Install_ResultEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_CAU_Audi_Install_Result, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_CAU_Audi_Install_Result{
		WmiInstance: tmp,
	}
	return
}

// SetAudiInstallHResult sets the value of AudiInstallHResult for the instance
func (instance *MSFT_CAU_Audi_Install_Result) SetPropertyAudiInstallHResult(value uint32) (err error) {
	return instance.SetProperty("AudiInstallHResult", (value))
}

// GetAudiInstallHResult gets the value of AudiInstallHResult for the instance
func (instance *MSFT_CAU_Audi_Install_Result) GetPropertyAudiInstallHResult() (value uint32, err error) {
	retValue, err := instance.GetProperty("AudiInstallHResult")
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

// SetRebootRequired sets the value of RebootRequired for the instance
func (instance *MSFT_CAU_Audi_Install_Result) SetPropertyRebootRequired(value bool) (err error) {
	return instance.SetProperty("RebootRequired", (value))
}

// GetRebootRequired gets the value of RebootRequired for the instance
func (instance *MSFT_CAU_Audi_Install_Result) GetPropertyRebootRequired() (value bool, err error) {
	retValue, err := instance.GetProperty("RebootRequired")
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

// SetUpdateId sets the value of UpdateId for the instance
func (instance *MSFT_CAU_Audi_Install_Result) SetPropertyUpdateId(value string) (err error) {
	return instance.SetProperty("UpdateId", (value))
}

// GetUpdateId gets the value of UpdateId for the instance
func (instance *MSFT_CAU_Audi_Install_Result) GetPropertyUpdateId() (value string, err error) {
	retValue, err := instance.GetProperty("UpdateId")
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
