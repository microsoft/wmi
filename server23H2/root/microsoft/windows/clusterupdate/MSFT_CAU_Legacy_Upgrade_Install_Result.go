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

// MSFT_CAU_Legacy_Upgrade_Install_Result struct
type MSFT_CAU_Legacy_Upgrade_Install_Result struct {
	*cim.WmiInstance

	//
	LaunchInstallerHResult uint32

	//
	ReturnCode int32
}

func NewMSFT_CAU_Legacy_Upgrade_Install_ResultEx1(instance *cim.WmiInstance) (newInstance *MSFT_CAU_Legacy_Upgrade_Install_Result, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_CAU_Legacy_Upgrade_Install_Result{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_CAU_Legacy_Upgrade_Install_ResultEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_CAU_Legacy_Upgrade_Install_Result, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_CAU_Legacy_Upgrade_Install_Result{
		WmiInstance: tmp,
	}
	return
}

// SetLaunchInstallerHResult sets the value of LaunchInstallerHResult for the instance
func (instance *MSFT_CAU_Legacy_Upgrade_Install_Result) SetPropertyLaunchInstallerHResult(value uint32) (err error) {
	return instance.SetProperty("LaunchInstallerHResult", (value))
}

// GetLaunchInstallerHResult gets the value of LaunchInstallerHResult for the instance
func (instance *MSFT_CAU_Legacy_Upgrade_Install_Result) GetPropertyLaunchInstallerHResult() (value uint32, err error) {
	retValue, err := instance.GetProperty("LaunchInstallerHResult")
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

// SetReturnCode sets the value of ReturnCode for the instance
func (instance *MSFT_CAU_Legacy_Upgrade_Install_Result) SetPropertyReturnCode(value int32) (err error) {
	return instance.SetProperty("ReturnCode", (value))
}

// GetReturnCode gets the value of ReturnCode for the instance
func (instance *MSFT_CAU_Legacy_Upgrade_Install_Result) GetPropertyReturnCode() (value int32, err error) {
	retValue, err := instance.GetProperty("ReturnCode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}
