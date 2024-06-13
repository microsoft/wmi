// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_HCSVM struct
type MSCluster_HCSVM struct {
	*cim.WmiInstance

	//
	CpuCount uint32

	//
	ExtendedVmConfiguration string

	//
	MacAddress string

	//
	MemorySizeInMb uint32

	//
	Name string

	//
	NetworkEndpointId string

	//
	OfflineAction uint32

	//
	ResState string

	//
	SwitchName string

	//
	VhdPath string

	//
	VmName string
}

func NewMSCluster_HCSVMEx1(instance *cim.WmiInstance) (newInstance *MSCluster_HCSVM, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSCluster_HCSVM{
		WmiInstance: tmp,
	}
	return
}

func NewMSCluster_HCSVMEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_HCSVM, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_HCSVM{
		WmiInstance: tmp,
	}
	return
}

// SetCpuCount sets the value of CpuCount for the instance
func (instance *MSCluster_HCSVM) SetPropertyCpuCount(value uint32) (err error) {
	return instance.SetProperty("CpuCount", (value))
}

// GetCpuCount gets the value of CpuCount for the instance
func (instance *MSCluster_HCSVM) GetPropertyCpuCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("CpuCount")
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

// SetExtendedVmConfiguration sets the value of ExtendedVmConfiguration for the instance
func (instance *MSCluster_HCSVM) SetPropertyExtendedVmConfiguration(value string) (err error) {
	return instance.SetProperty("ExtendedVmConfiguration", (value))
}

// GetExtendedVmConfiguration gets the value of ExtendedVmConfiguration for the instance
func (instance *MSCluster_HCSVM) GetPropertyExtendedVmConfiguration() (value string, err error) {
	retValue, err := instance.GetProperty("ExtendedVmConfiguration")
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

// SetMacAddress sets the value of MacAddress for the instance
func (instance *MSCluster_HCSVM) SetPropertyMacAddress(value string) (err error) {
	return instance.SetProperty("MacAddress", (value))
}

// GetMacAddress gets the value of MacAddress for the instance
func (instance *MSCluster_HCSVM) GetPropertyMacAddress() (value string, err error) {
	retValue, err := instance.GetProperty("MacAddress")
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

// SetMemorySizeInMb sets the value of MemorySizeInMb for the instance
func (instance *MSCluster_HCSVM) SetPropertyMemorySizeInMb(value uint32) (err error) {
	return instance.SetProperty("MemorySizeInMb", (value))
}

// GetMemorySizeInMb gets the value of MemorySizeInMb for the instance
func (instance *MSCluster_HCSVM) GetPropertyMemorySizeInMb() (value uint32, err error) {
	retValue, err := instance.GetProperty("MemorySizeInMb")
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

// SetName sets the value of Name for the instance
func (instance *MSCluster_HCSVM) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSCluster_HCSVM) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetNetworkEndpointId sets the value of NetworkEndpointId for the instance
func (instance *MSCluster_HCSVM) SetPropertyNetworkEndpointId(value string) (err error) {
	return instance.SetProperty("NetworkEndpointId", (value))
}

// GetNetworkEndpointId gets the value of NetworkEndpointId for the instance
func (instance *MSCluster_HCSVM) GetPropertyNetworkEndpointId() (value string, err error) {
	retValue, err := instance.GetProperty("NetworkEndpointId")
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

// SetOfflineAction sets the value of OfflineAction for the instance
func (instance *MSCluster_HCSVM) SetPropertyOfflineAction(value uint32) (err error) {
	return instance.SetProperty("OfflineAction", (value))
}

// GetOfflineAction gets the value of OfflineAction for the instance
func (instance *MSCluster_HCSVM) GetPropertyOfflineAction() (value uint32, err error) {
	retValue, err := instance.GetProperty("OfflineAction")
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

// SetResState sets the value of ResState for the instance
func (instance *MSCluster_HCSVM) SetPropertyResState(value string) (err error) {
	return instance.SetProperty("ResState", (value))
}

// GetResState gets the value of ResState for the instance
func (instance *MSCluster_HCSVM) GetPropertyResState() (value string, err error) {
	retValue, err := instance.GetProperty("ResState")
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

// SetSwitchName sets the value of SwitchName for the instance
func (instance *MSCluster_HCSVM) SetPropertySwitchName(value string) (err error) {
	return instance.SetProperty("SwitchName", (value))
}

// GetSwitchName gets the value of SwitchName for the instance
func (instance *MSCluster_HCSVM) GetPropertySwitchName() (value string, err error) {
	retValue, err := instance.GetProperty("SwitchName")
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

// SetVhdPath sets the value of VhdPath for the instance
func (instance *MSCluster_HCSVM) SetPropertyVhdPath(value string) (err error) {
	return instance.SetProperty("VhdPath", (value))
}

// GetVhdPath gets the value of VhdPath for the instance
func (instance *MSCluster_HCSVM) GetPropertyVhdPath() (value string, err error) {
	retValue, err := instance.GetProperty("VhdPath")
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

// SetVmName sets the value of VmName for the instance
func (instance *MSCluster_HCSVM) SetPropertyVmName(value string) (err error) {
	return instance.SetProperty("VmName", (value))
}

// GetVmName gets the value of VmName for the instance
func (instance *MSCluster_HCSVM) GetPropertyVmName() (value string, err error) {
	retValue, err := instance.GetProperty("VmName")
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

//

// <param name="CpuCount" type="uint32 "></param>
// <param name="ExtendedVmConfiguration" type="string "></param>
// <param name="MemorySizeInMb" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="OfflineAction" type="uint32 "></param>
// <param name="SwitchName" type="string "></param>
// <param name="VhdPath" type="string "></param>
// <param name="VmName" type="string "></param>

// <param name="CreatedHCSVMCluster" type="MSCluster_HCSVM "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_HCSVM) NewClusterHCSVM( /* IN */ Name string,
	/* IN */ SwitchName string,
	/* IN */ ExtendedVmConfiguration string,
	/* IN */ MemorySizeInMb uint32,
	/* IN */ CpuCount uint32,
	/* IN */ VhdPath string,
	/* IN */ VmName string,
	/* IN */ OfflineAction uint32,
	/* OUT */ CreatedHCSVMCluster MSCluster_HCSVM) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewClusterHCSVM", Name, SwitchName, ExtendedVmConfiguration, MemorySizeInMb, CpuCount, VhdPath, VmName, OfflineAction)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="Name" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_HCSVM) RemoveClusterHCSVM( /* IN */ Name string,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveClusterHCSVM", Name, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ExtendedVmConfiguration" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="NewName" type="string "></param>
// <param name="OfflineAction" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_HCSVM) SetClusterHCSVM( /* IN */ Name string,
	/* IN */ NewName string,
	/* IN */ ExtendedVmConfiguration string,
	/* IN */ OfflineAction uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetClusterHCSVM", Name, NewName, ExtendedVmConfiguration, OfflineAction)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Name" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_HCSVM) StartClusterHCSVM( /* IN */ Name string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("StartClusterHCSVM", Name)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Name" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_HCSVM) StopClusterHCSVM( /* IN */ Name string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("StopClusterHCSVM", Name)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
