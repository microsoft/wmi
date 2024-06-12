// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfigurationProxy
//////////////////////////////////////////////
package desiredstateconfigurationproxy

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DscProxy struct
type MSFT_DscProxy struct {
	*cim.WmiInstance
}

func NewMSFT_DscProxyEx1(instance *cim.WmiInstance) (newInstance *MSFT_DscProxy, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_DscProxy{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_DscProxyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DscProxy, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DscProxy{
		WmiInstance: tmp,
	}
	return
}

// Get resource state.

// <param name="ConfigurationData" type="uint8 []"></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="state" type="bool "></param>
func (instance *MSFT_DscProxy) GetResourceState( /* IN */ ConfigurationData []uint8,
	/* OUT */ state bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetResourceState", ConfigurationData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
