// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_RedundancySet struct
type CIM_RedundancySet struct {
	*CIM_SystemSpecificCollection

	// The current load balance algorithm.
	///Least Blocks, Least IO, and Address Region are used in storage device path redundancy drivers to optimize load balancing by routing requests to a path with the least queued blocks or IO requests, or based on locality of reference.
	///'Product Specific' indicates that the algorithm is optimized for a particular type of product. Information about that product SHOULD be provided in an associated CIM_Product instance.
	LoadBalanceAlgorithm RedundancySet_LoadBalanceAlgorithm

	// MaxNumberSupported indicates the largest number of elements that can participate in the RedundancySet. A value of 0 indicates there is no limit on the number of elements.
	MaxNumberSupported uint32

	// MinNumberNeeded indicates the smallest number of elements that MUST be operational in order to function. For example, in an N+1 redundancy relationship, the MinNumberNeeded property is set equal to N. In a 'LimitedSparing' environment, this property is meaningless and SHOULD be set to zero.
	MinNumberNeeded uint32

	// When LoadBalanceAlgorithm is Other, this property describes the algorithm.
	OtherLoadBalanceAlgorithm string

	// When the corresponding array entry in TypeOfSet[] is 'Other', this entry provides a string describing the type of set.
	OtherTypeOfSet []string

	// RedundancyStatus provides information on the state of the RedundancyGroup. 'Fully Redundant' (value=2) means that all of the configured redundancy is still available; 'Degraded Redundancy' (3) means that some configured elements are degraded, missing or failed but that the number of elements in the set is still greater than the minimum required ('MinNumberNeeded'); 'Redundancy Lost' (4) means that sufficient configured elements are missing or failed that no redundancy is available and the next failure experienced will cause overall failure. 'Overall Failure' (5) means that there has been an overall failure of the RedundancySet.
	RedundancyStatus RedundancySet_RedundancyStatus

	// TypeOfSet provides information on the type of redundancy. N+1 (=2) indicates all members are active, are unaware and function independent of one another. However, there exist at least one extra member to achieve functionality. 'Sparing' is implied (i.e. each member can be a spare for the other(s). An example of N+1 is a system that has 2 power supplies, but needs only 1 power supply to functioning properly. N+1 is a special case of N+M redundancy where M=1. A value of N+1 (=2) shall be used for N+M redundancy. - Load Balanced (=3) indicates all members are active. However, there functionality is not independent of each other. Their functioning is determined by some sort of load balancing algrothim (implemented in hardware and/or software). 'Sparing' is implied (i.e. each member can be a spare for the other(s).
	///- Sparing (=4) indicates that all members are active and are aware of each others. However, their functionality is independent until failover. Each member can be a spare for the other(s).
	///- Limited Sparing (=5) indicates that all members are active, and they may or may not be aware of each and they are not spares for each other. Instead, their redundancy is indicated by the IsSpare relationship.
	TypeOfSet []RedundancySet_TypeOfSet

	// VendorIdentifyingInfo captures the vendor identifying data for the RedundancySet. One example is the product name for a cluster.
	VendorIdentifyingInfo string
}

func NewCIM_RedundancySetEx1(instance *cim.WmiInstance) (newInstance *CIM_RedundancySet, err error) {
	tmp, err := NewCIM_SystemSpecificCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_RedundancySet{
		CIM_SystemSpecificCollection: tmp,
	}
	return
}

func NewCIM_RedundancySetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_RedundancySet, err error) {
	tmp, err := NewCIM_SystemSpecificCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_RedundancySet{
		CIM_SystemSpecificCollection: tmp,
	}
	return
}

// SetLoadBalanceAlgorithm sets the value of LoadBalanceAlgorithm for the instance
func (instance *CIM_RedundancySet) SetPropertyLoadBalanceAlgorithm(value RedundancySet_LoadBalanceAlgorithm) (err error) {
	return instance.SetProperty("LoadBalanceAlgorithm", value)
}

// GetLoadBalanceAlgorithm gets the value of LoadBalanceAlgorithm for the instance
func (instance *CIM_RedundancySet) GetPropertyLoadBalanceAlgorithm() (value RedundancySet_LoadBalanceAlgorithm, err error) {
	retValue, err := instance.GetProperty("LoadBalanceAlgorithm")
	if err != nil {
		return
	}
	value, ok := retValue.(RedundancySet_LoadBalanceAlgorithm)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxNumberSupported sets the value of MaxNumberSupported for the instance
func (instance *CIM_RedundancySet) SetPropertyMaxNumberSupported(value uint32) (err error) {
	return instance.SetProperty("MaxNumberSupported", value)
}

// GetMaxNumberSupported gets the value of MaxNumberSupported for the instance
func (instance *CIM_RedundancySet) GetPropertyMaxNumberSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxNumberSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinNumberNeeded sets the value of MinNumberNeeded for the instance
func (instance *CIM_RedundancySet) SetPropertyMinNumberNeeded(value uint32) (err error) {
	return instance.SetProperty("MinNumberNeeded", value)
}

// GetMinNumberNeeded gets the value of MinNumberNeeded for the instance
func (instance *CIM_RedundancySet) GetPropertyMinNumberNeeded() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinNumberNeeded")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherLoadBalanceAlgorithm sets the value of OtherLoadBalanceAlgorithm for the instance
func (instance *CIM_RedundancySet) SetPropertyOtherLoadBalanceAlgorithm(value string) (err error) {
	return instance.SetProperty("OtherLoadBalanceAlgorithm", value)
}

// GetOtherLoadBalanceAlgorithm gets the value of OtherLoadBalanceAlgorithm for the instance
func (instance *CIM_RedundancySet) GetPropertyOtherLoadBalanceAlgorithm() (value string, err error) {
	retValue, err := instance.GetProperty("OtherLoadBalanceAlgorithm")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherTypeOfSet sets the value of OtherTypeOfSet for the instance
func (instance *CIM_RedundancySet) SetPropertyOtherTypeOfSet(value []string) (err error) {
	return instance.SetProperty("OtherTypeOfSet", value)
}

// GetOtherTypeOfSet gets the value of OtherTypeOfSet for the instance
func (instance *CIM_RedundancySet) GetPropertyOtherTypeOfSet() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherTypeOfSet")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRedundancyStatus sets the value of RedundancyStatus for the instance
func (instance *CIM_RedundancySet) SetPropertyRedundancyStatus(value RedundancySet_RedundancyStatus) (err error) {
	return instance.SetProperty("RedundancyStatus", value)
}

// GetRedundancyStatus gets the value of RedundancyStatus for the instance
func (instance *CIM_RedundancySet) GetPropertyRedundancyStatus() (value RedundancySet_RedundancyStatus, err error) {
	retValue, err := instance.GetProperty("RedundancyStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(RedundancySet_RedundancyStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTypeOfSet sets the value of TypeOfSet for the instance
func (instance *CIM_RedundancySet) SetPropertyTypeOfSet(value []RedundancySet_TypeOfSet) (err error) {
	return instance.SetProperty("TypeOfSet", value)
}

// GetTypeOfSet gets the value of TypeOfSet for the instance
func (instance *CIM_RedundancySet) GetPropertyTypeOfSet() (value []RedundancySet_TypeOfSet, err error) {
	retValue, err := instance.GetProperty("TypeOfSet")
	if err != nil {
		return
	}
	value, ok := retValue.([]RedundancySet_TypeOfSet)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVendorIdentifyingInfo sets the value of VendorIdentifyingInfo for the instance
func (instance *CIM_RedundancySet) SetPropertyVendorIdentifyingInfo(value string) (err error) {
	return instance.SetProperty("VendorIdentifyingInfo", value)
}

// GetVendorIdentifyingInfo gets the value of VendorIdentifyingInfo for the instance
func (instance *CIM_RedundancySet) GetPropertyVendorIdentifyingInfo() (value string, err error) {
	retValue, err := instance.GetProperty("VendorIdentifyingInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// This method forces a failover from one ManagedElement to another. There are two parameters to the Failover method.
///- FailoverFrom is a reference to an 'active' ManagedElement that will become inactive after the method. This element SHOULD be part of the RedundancySet via a MemberOfCollection relationship.
///- FailoverTo is a reference to the ManagedElement that will take over for the FailoverFrom element. This element SHOULD either be a member of the RedundancySet or be associated with the RedundancySet via an IsSpare relationship.
///
///Upon sucessful completion:
///- the FailoverTo element SHOULD be associated to the RedundancySet via MemberOfCollection.
///- the FailFrom element SHOULD either still be associated to the RedundandySet via MemberOfCollection with a OperationalStatus or EnableState that indicates it not active, or it SHOULD be associated to the 'Spared' collection via the MemberOfCollection association.

// <param name="FailoverFrom" type="CIM_ManagedElement ">The primary ManagedSystemElement that will become inactive after the method.</param>
// <param name="FailoverTo" type="CIM_ManagedElement ">The ManagedSystemElement that will take over from the primary MSE.</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_RedundancySet) Failover( /* IN */ FailoverFrom CIM_ManagedElement,
	/* IN */ FailoverTo CIM_ManagedElement) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Failover", FailoverFrom, FailoverTo)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
