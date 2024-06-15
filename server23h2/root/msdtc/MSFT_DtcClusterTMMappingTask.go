// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.msdtc
//////////////////////////////////////////////
package msdtc

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DtcClusterTMMappingTask struct
type MSFT_DtcClusterTMMappingTask struct {
	*cim.WmiInstance
}

func NewMSFT_DtcClusterTMMappingTaskEx1(instance *cim.WmiInstance) (newInstance *MSFT_DtcClusterTMMappingTask, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_DtcClusterTMMappingTask{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_DtcClusterTMMappingTaskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DtcClusterTMMappingTask, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DtcClusterTMMappingTask{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="Name" type="string "></param>

// <param name="cmdletOutput" type="DtcClusterTMMapping []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcClusterTMMappingTask) Get( /* IN */ Name string,
	/* OUT */ cmdletOutput []DtcClusterTMMapping) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", Name)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ClusterResourceName" type="string "></param>
// <param name="ComPlusAppId" type="string "></param>
// <param name="Local" type="bool "></param>
// <param name="Name" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcClusterTMMappingTask) AddByComPlusSet( /* IN */ Name string,
	/* IN */ ClusterResourceName string,
	/* IN */ Local bool,
	/* IN */ ComPlusAppId string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddByComPlusSet", Name, ClusterResourceName, Local, ComPlusAppId)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ClusterResourceName" type="string "></param>
// <param name="ExecutablePath" type="string "></param>
// <param name="Local" type="bool "></param>
// <param name="Name" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcClusterTMMappingTask) AddByExeSet( /* IN */ Name string,
	/* IN */ ClusterResourceName string,
	/* IN */ Local bool,
	/* IN */ ExecutablePath string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddByExeSet", Name, ClusterResourceName, Local, ExecutablePath)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ClusterResourceName" type="string "></param>
// <param name="Local" type="bool "></param>
// <param name="Name" type="string "></param>
// <param name="Service" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcClusterTMMappingTask) AddByServiceSet( /* IN */ Name string,
	/* IN */ ClusterResourceName string,
	/* IN */ Local bool,
	/* IN */ Service string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddByServiceSet", Name, ClusterResourceName, Local, Service)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="All" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcClusterTMMappingTask) RemoveByAllParameterSet( /* IN */ All bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveByAllParameterSet", All)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Name" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcClusterTMMappingTask) RemoveByMappingNameParameterSet( /* IN */ Name string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveByMappingNameParameterSet", Name)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ClusterResourceName" type="string "></param>
// <param name="ComPlusAppId" type="string "></param>
// <param name="Local" type="bool "></param>
// <param name="Name" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcClusterTMMappingTask) SetByComPlusSet( /* IN */ Name string,
	/* IN */ ComPlusAppId string,
	/* IN */ ClusterResourceName string,
	/* IN */ Local bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetByComPlusSet", Name, ComPlusAppId, ClusterResourceName, Local)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ClusterResourceName" type="string "></param>
// <param name="ExecutablePath" type="string "></param>
// <param name="Local" type="bool "></param>
// <param name="Name" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcClusterTMMappingTask) SetByExeSet( /* IN */ Name string,
	/* IN */ ExecutablePath string,
	/* IN */ ClusterResourceName string,
	/* IN */ Local bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetByExeSet", Name, ExecutablePath, ClusterResourceName, Local)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ClusterResourceName" type="string "></param>
// <param name="Local" type="bool "></param>
// <param name="Name" type="string "></param>
// <param name="Service" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcClusterTMMappingTask) SetByServiceSet( /* IN */ Name string,
	/* IN */ Service string,
	/* IN */ ClusterResourceName string,
	/* IN */ Local bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetByServiceSet", Name, Service, ClusterResourceName, Local)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
