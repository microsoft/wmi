// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Cluster.Scaleout
//////////////////////////////////////////////
package scaleout

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ClusterSetMember struct
type MSFT_ClusterSetMember struct {
	*cim.WmiInstance

	//
	ClusterName string

	//
	Id uint64

	//
	NetworkPrefixes string

	//
	State uint32

	//
	Tags []string

	//
	TopologyLabel string

	//
	WorkloadCount uint32
}

func NewMSFT_ClusterSetMemberEx1(instance *cim.WmiInstance) (newInstance *MSFT_ClusterSetMember, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ClusterSetMember{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ClusterSetMemberEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ClusterSetMember, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ClusterSetMember{
		WmiInstance: tmp,
	}
	return
}

// SetClusterName sets the value of ClusterName for the instance
func (instance *MSFT_ClusterSetMember) SetPropertyClusterName(value string) (err error) {
	return instance.SetProperty("ClusterName", value)
}

// GetClusterName gets the value of ClusterName for the instance
func (instance *MSFT_ClusterSetMember) GetPropertyClusterName() (value string, err error) {
	retValue, err := instance.GetProperty("ClusterName")
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
func (instance *MSFT_ClusterSetMember) SetPropertyId(value uint64) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *MSFT_ClusterSetMember) GetPropertyId() (value uint64, err error) {
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

// SetNetworkPrefixes sets the value of NetworkPrefixes for the instance
func (instance *MSFT_ClusterSetMember) SetPropertyNetworkPrefixes(value string) (err error) {
	return instance.SetProperty("NetworkPrefixes", value)
}

// GetNetworkPrefixes gets the value of NetworkPrefixes for the instance
func (instance *MSFT_ClusterSetMember) GetPropertyNetworkPrefixes() (value string, err error) {
	retValue, err := instance.GetProperty("NetworkPrefixes")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSFT_ClusterSetMember) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSFT_ClusterSetMember) GetPropertyState() (value uint32, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTags sets the value of Tags for the instance
func (instance *MSFT_ClusterSetMember) SetPropertyTags(value []string) (err error) {
	return instance.SetProperty("Tags", value)
}

// GetTags gets the value of Tags for the instance
func (instance *MSFT_ClusterSetMember) GetPropertyTags() (value []string, err error) {
	retValue, err := instance.GetProperty("Tags")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTopologyLabel sets the value of TopologyLabel for the instance
func (instance *MSFT_ClusterSetMember) SetPropertyTopologyLabel(value string) (err error) {
	return instance.SetProperty("TopologyLabel", value)
}

// GetTopologyLabel gets the value of TopologyLabel for the instance
func (instance *MSFT_ClusterSetMember) GetPropertyTopologyLabel() (value string, err error) {
	retValue, err := instance.GetProperty("TopologyLabel")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkloadCount sets the value of WorkloadCount for the instance
func (instance *MSFT_ClusterSetMember) SetPropertyWorkloadCount(value uint32) (err error) {
	return instance.SetProperty("WorkloadCount", value)
}

// GetWorkloadCount gets the value of WorkloadCount for the instance
func (instance *MSFT_ClusterSetMember) GetPropertyWorkloadCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkloadCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="Flags" type="uint32 "></param>
// <param name="Force" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetMember) Evict( /* IN */ Force bool,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Evict", Force, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="KeyType" type="uint32 "></param>

// <param name="key" type="MSFT_ClusterSetKey "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetMember) GetKey( /* IN */ KeyType uint32,
	/* IN */ Flags uint32,
	/* OUT */ key MSFT_ClusterSetKey) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetKey", KeyType, Flags)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AvailabilitySetName" type="string "></param>
// <param name="Id" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="vmConfig" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="VM" type="MSFT_ClusterSetVM "></param>
func (instance *MSFT_ClusterSetMember) ClusterLocalVm( /* IN */ vmConfig string,
	/* IN */ Id string,
	/* IN */ Name string,
	/* IN */ AvailabilitySetName string,
	/* OUT */ VM MSFT_ClusterSetVM) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ClusterLocalVm", vmConfig, Id, Name, AvailabilitySetName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Tag" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetMember) AddTags( /* IN */ Tag []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddTags", Tag)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Tag" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetMember) RemoveTags( /* IN */ Tag []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveTags", Tag)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="AvailabilitySetName" type="string "></param>
// <param name="RegisterAll" type="bool "></param>
// <param name="VMId" type="string "></param>
// <param name="VMName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="VM" type="MSFT_ClusterSetVM []"></param>
func (instance *MSFT_ClusterSetMember) AddVM( /* IN */ VMId string,
	/* IN */ VMName string,
	/* IN */ RegisterAll bool,
	/* IN */ AvailabilitySetName string,
	/* OUT */ VM []MSFT_ClusterSetVM) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddVM", VMId, VMName, RegisterAll, AvailabilitySetName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
