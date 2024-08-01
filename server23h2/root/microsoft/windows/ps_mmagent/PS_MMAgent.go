// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.PS_MMAgent
//////////////////////////////////////////////
package ps_mmagent

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_MMAgent struct
type PS_MMAgent struct {
	*cim.WmiInstance
}

func NewPS_MMAgentEx1(instance *cim.WmiInstance) (newInstance *PS_MMAgent, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_MMAgent{
		WmiInstance: tmp,
	}
	return
}

func NewPS_MMAgentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_MMAgent, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_MMAgent{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ApplicationLaunchPrefetching" type="bool "></param>
// <param name="ApplicationPreLaunch" type="bool "></param>
// <param name="MemoryCompression" type="bool "></param>
// <param name="OperationAPI" type="bool "></param>
// <param name="PageCombining" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_MMAgent) Enable( /* IN */ ApplicationLaunchPrefetching bool,
	/* IN */ OperationAPI bool,
	/* IN */ PageCombining bool,
	/* IN */ ApplicationPreLaunch bool,
	/* IN */ MemoryCompression bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable", ApplicationLaunchPrefetching, OperationAPI, PageCombining, ApplicationPreLaunch, MemoryCompression)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ApplicationLaunchPrefetching" type="bool "></param>
// <param name="ApplicationPreLaunch" type="bool "></param>
// <param name="MemoryCompression" type="bool "></param>
// <param name="OperationAPI" type="bool "></param>
// <param name="PageCombining" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_MMAgent) Disable( /* IN */ ApplicationLaunchPrefetching bool,
	/* IN */ OperationAPI bool,
	/* IN */ PageCombining bool,
	/* IN */ ApplicationPreLaunch bool,
	/* IN */ MemoryCompression bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable", ApplicationLaunchPrefetching, OperationAPI, PageCombining, ApplicationPreLaunch, MemoryCompression)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="MaxOperationAPIFiles" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_MMAgent) Set( /* IN */ MaxOperationAPIFiles uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Set", MaxOperationAPIFiles)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="cmdletOutput" type="MMAgentComponents "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_MMAgent) Get( /* OUT */ cmdletOutput MMAgentComponents) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DisableDebugMode" type="bool "></param>
// <param name="PackageFullName" type="string "></param>
// <param name="PackageRelativeAppId" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_MMAgent) Debug( /* IN */ PackageFullName string,
	/* IN */ DisableDebugMode bool,
	/* IN */ PackageRelativeAppId string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Debug", PackageFullName, DisableDebugMode, PackageRelativeAppId)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
