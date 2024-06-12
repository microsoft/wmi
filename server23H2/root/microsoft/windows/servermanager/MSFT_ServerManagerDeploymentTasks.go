// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ServerManagerDeploymentTasks struct
type MSFT_ServerManagerDeploymentTasks struct {
	*cim.WmiInstance
}

func NewMSFT_ServerManagerDeploymentTasksEx1(instance *cim.WmiInstance) (newInstance *MSFT_ServerManagerDeploymentTasks, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerManagerDeploymentTasks{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ServerManagerDeploymentTasksEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ServerManagerDeploymentTasks, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerManagerDeploymentTasks{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="RequestGuid" type="MSFT_ServerManagerRequestGuid "></param>

// <param name="EnumerationState" type="MSFT_ServerManagerRequestState "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="ServerComponents" type="MSFT_ServerManagerServerComponent []"></param>
func (instance *MSFT_ServerManagerDeploymentTasks) GetServerComponentsAsync( /* IN */ RequestGuid MSFT_ServerManagerRequestGuid,
	/* OUT */ EnumerationState MSFT_ServerManagerRequestState,
	/* OUT */ ServerComponents []MSFT_ServerManagerServerComponent) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetServerComponentsAsync", RequestGuid)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RequestGuid" type="MSFT_ServerManagerRequestGuid "></param>
// <param name="VhdPath" type="string "></param>

// <param name="EnumerationState" type="MSFT_ServerManagerRequestState "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="ServerComponents" type="MSFT_ServerManagerServerComponent []"></param>
func (instance *MSFT_ServerManagerDeploymentTasks) GetServerComponentsVhdAsync( /* IN */ RequestGuid MSFT_ServerManagerRequestGuid,
	/* IN */ VhdPath string,
	/* OUT */ EnumerationState MSFT_ServerManagerRequestState,
	/* OUT */ ServerComponents []MSFT_ServerManagerServerComponent) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetServerComponentsVhdAsync", RequestGuid, VhdPath)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RequestGuid" type="MSFT_ServerManagerRequestGuid "></param>

// <param name="EnumerationState" type="MSFT_ServerManagerRequestState "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="ServerComponents" type="MSFT_ServerManagerServerComponent []"></param>
func (instance *MSFT_ServerManagerDeploymentTasks) GetEnumerationRequestState( /* IN */ RequestGuid MSFT_ServerManagerRequestGuid,
	/* OUT */ EnumerationState MSFT_ServerManagerRequestState,
	/* OUT */ ServerComponents []MSFT_ServerManagerServerComponent) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetEnumerationRequestState", RequestGuid)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RequestGuid" type="MSFT_ServerManagerRequestGuid "></param>
// <param name="ScanForUpdates" type="bool "></param>
// <param name="ServerComponentDescriptors" type="MSFT_ServerManagerServerComponentDescriptor []"></param>
// <param name="Source" type="string []"></param>

// <param name="AlterationState" type="MSFT_ServerManagerRequestState "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ServerManagerDeploymentTasks) AddServerComponentAsync( /* IN */ RequestGuid MSFT_ServerManagerRequestGuid,
	/* IN */ Source []string,
	/* IN */ ScanForUpdates bool,
	/* IN */ ServerComponentDescriptors []MSFT_ServerManagerServerComponentDescriptor,
	/* OUT */ AlterationState MSFT_ServerManagerRequestState) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddServerComponentAsync", RequestGuid, Source, ScanForUpdates, ServerComponentDescriptors)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RequestGuid" type="MSFT_ServerManagerRequestGuid "></param>
// <param name="ScanForUpdates" type="bool "></param>
// <param name="ServerComponentDescriptors" type="MSFT_ServerManagerServerComponentDescriptor []"></param>
// <param name="Source" type="string []"></param>
// <param name="VhdPath" type="string "></param>

// <param name="AlterationState" type="MSFT_ServerManagerRequestState "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ServerManagerDeploymentTasks) AddServerComponentVhdAsync( /* IN */ RequestGuid MSFT_ServerManagerRequestGuid,
	/* IN */ Source []string,
	/* IN */ ScanForUpdates bool,
	/* IN */ ServerComponentDescriptors []MSFT_ServerManagerServerComponentDescriptor,
	/* IN */ VhdPath string,
	/* OUT */ AlterationState MSFT_ServerManagerRequestState) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddServerComponentVhdAsync", RequestGuid, Source, ScanForUpdates, ServerComponentDescriptors, VhdPath)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DeleteComponents" type="bool "></param>
// <param name="RequestGuid" type="MSFT_ServerManagerRequestGuid "></param>
// <param name="ServerComponentDescriptors" type="MSFT_ServerManagerServerComponentDescriptor []"></param>

// <param name="AlterationState" type="MSFT_ServerManagerRequestState "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ServerManagerDeploymentTasks) RemoveServerComponentAsync( /* IN */ RequestGuid MSFT_ServerManagerRequestGuid,
	/* IN */ DeleteComponents bool,
	/* IN */ ServerComponentDescriptors []MSFT_ServerManagerServerComponentDescriptor,
	/* OUT */ AlterationState MSFT_ServerManagerRequestState) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RemoveServerComponentAsync", RequestGuid, DeleteComponents, ServerComponentDescriptors)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DeleteComponents" type="bool "></param>
// <param name="RequestGuid" type="MSFT_ServerManagerRequestGuid "></param>
// <param name="ServerComponentDescriptors" type="MSFT_ServerManagerServerComponentDescriptor []"></param>
// <param name="VhdPath" type="string "></param>

// <param name="AlterationState" type="MSFT_ServerManagerRequestState "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ServerManagerDeploymentTasks) RemoveServerComponentVhdAsync( /* IN */ RequestGuid MSFT_ServerManagerRequestGuid,
	/* IN */ DeleteComponents bool,
	/* IN */ ServerComponentDescriptors []MSFT_ServerManagerServerComponentDescriptor,
	/* IN */ VhdPath string,
	/* OUT */ AlterationState MSFT_ServerManagerRequestState) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RemoveServerComponentVhdAsync", RequestGuid, DeleteComponents, ServerComponentDescriptors, VhdPath)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="KeepAlterationStateOnRestartRequired" type="bool "></param>
// <param name="RequestGuid" type="MSFT_ServerManagerRequestGuid "></param>

// <param name="AlterationState" type="MSFT_ServerManagerRequestState "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="ServerComponents" type="MSFT_ServerManagerServerComponent []"></param>
func (instance *MSFT_ServerManagerDeploymentTasks) GetAlterationRequestState( /* IN */ RequestGuid MSFT_ServerManagerRequestGuid,
	/* IN */ KeepAlterationStateOnRestartRequired bool,
	/* OUT */ AlterationState MSFT_ServerManagerRequestState,
	/* OUT */ ServerComponents []MSFT_ServerManagerServerComponent) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetAlterationRequestState", RequestGuid, KeepAlterationStateOnRestartRequired)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
