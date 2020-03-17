// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Cluster.Scaleout
//////////////////////////////////////////////
package scaleout

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ClusterSetFaultDomain struct
type MSFT_ClusterSetFaultDomain struct {
	cim.WmiInstance

	//
	ClusterName []string

	//
	Description string

	//
	FaultDomainType uint32

	//
	FDName string

	//
	Id uint64
}

// SetClusterName sets the value of ClusterName for the instance
func (instance *MSFT_ClusterSetFaultDomain) SetPropertyClusterName(value []string) (err error) {
	return instance.SetProperty("ClusterName", value)
}

// GetClusterName gets the value of ClusterName for the instance
func (instance *MSFT_ClusterSetFaultDomain) GetPropertyClusterName() (value []string, err error) {
	retValue, err := instance.GetProperty("ClusterName")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_ClusterSetFaultDomain) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_ClusterSetFaultDomain) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFaultDomainType sets the value of FaultDomainType for the instance
func (instance *MSFT_ClusterSetFaultDomain) SetPropertyFaultDomainType(value uint32) (err error) {
	return instance.SetProperty("FaultDomainType", value)
}

// GetFaultDomainType gets the value of FaultDomainType for the instance
func (instance *MSFT_ClusterSetFaultDomain) GetPropertyFaultDomainType() (value uint32, err error) {
	retValue, err := instance.GetProperty("FaultDomainType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFDName sets the value of FDName for the instance
func (instance *MSFT_ClusterSetFaultDomain) SetPropertyFDName(value string) (err error) {
	return instance.SetProperty("FDName", value)
}

// GetFDName gets the value of FDName for the instance
func (instance *MSFT_ClusterSetFaultDomain) GetPropertyFDName() (value string, err error) {
	retValue, err := instance.GetProperty("FDName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *MSFT_ClusterSetFaultDomain) SetPropertyId(value uint64) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *MSFT_ClusterSetFaultDomain) GetPropertyId() (value uint64, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ClusterName" type="string []"></param>
// <param name="Description" type="string "></param>
// <param name="FDName" type="string "></param>
// <param name="FDType" type="uint32 "></param>
// <param name="Flags" type="uint32 "></param>

// <param name="CreatedFaultDomain" type="MSFT_ClusterSetFaultDomain "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetFaultDomain) CreateFaultDomain( /* IN */ FDName string,
	/* IN */ ClusterName []string,
	/* IN */ FDType uint32,
	/* IN */ Description string,
	/* IN */ Flags uint32,
	/* OUT */ CreatedFaultDomain MSFT_ClusterSetFaultDomain) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateFaultDomain", FDName, ClusterName, FDType, Description, Flags)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Flags" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetFaultDomain) RemoveFaultDomain( /* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveFaultDomain", Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ClusterName" type="string []"></param>
// <param name="Flags" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetFaultDomain) AddMembers( /* IN */ ClusterName []string,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddMembers", ClusterName, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ClusterName" type="string []"></param>
// <param name="Flags" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetFaultDomain) RemoveMembers( /* IN */ ClusterName []string,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveMembers", ClusterName, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
