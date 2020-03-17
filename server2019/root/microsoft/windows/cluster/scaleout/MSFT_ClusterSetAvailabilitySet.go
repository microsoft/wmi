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

// MSFT_ClusterSetAvailabilitySet struct
type MSFT_ClusterSetAvailabilitySet struct {
	cim.WmiInstance

	//
	AvailabilitySetName string

	//
	DomainAssignments []string

	//
	FaultDomains uint32

	//
	FaultDomainType uint32

	//
	Id uint64

	//
	ParticipantIds []uint64

	//
	SoftEnforcement bool

	//
	UpdateDomains uint32

	//
	Workloads []uint64
}

// SetAvailabilitySetName sets the value of AvailabilitySetName for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) SetPropertyAvailabilitySetName(value string) (err error) {
	return instance.SetProperty("AvailabilitySetName", value)
}

// GetAvailabilitySetName gets the value of AvailabilitySetName for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) GetPropertyAvailabilitySetName() (value string, err error) {
	retValue, err := instance.GetProperty("AvailabilitySetName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDomainAssignments sets the value of DomainAssignments for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) SetPropertyDomainAssignments(value []string) (err error) {
	return instance.SetProperty("DomainAssignments", value)
}

// GetDomainAssignments gets the value of DomainAssignments for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) GetPropertyDomainAssignments() (value []string, err error) {
	retValue, err := instance.GetProperty("DomainAssignments")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFaultDomains sets the value of FaultDomains for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) SetPropertyFaultDomains(value uint32) (err error) {
	return instance.SetProperty("FaultDomains", value)
}

// GetFaultDomains gets the value of FaultDomains for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) GetPropertyFaultDomains() (value uint32, err error) {
	retValue, err := instance.GetProperty("FaultDomains")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFaultDomainType sets the value of FaultDomainType for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) SetPropertyFaultDomainType(value uint32) (err error) {
	return instance.SetProperty("FaultDomainType", value)
}

// GetFaultDomainType gets the value of FaultDomainType for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) GetPropertyFaultDomainType() (value uint32, err error) {
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

// SetId sets the value of Id for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) SetPropertyId(value uint64) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) GetPropertyId() (value uint64, err error) {
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

// SetParticipantIds sets the value of ParticipantIds for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) SetPropertyParticipantIds(value []uint64) (err error) {
	return instance.SetProperty("ParticipantIds", value)
}

// GetParticipantIds gets the value of ParticipantIds for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) GetPropertyParticipantIds() (value []uint64, err error) {
	retValue, err := instance.GetProperty("ParticipantIds")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSoftEnforcement sets the value of SoftEnforcement for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) SetPropertySoftEnforcement(value bool) (err error) {
	return instance.SetProperty("SoftEnforcement", value)
}

// GetSoftEnforcement gets the value of SoftEnforcement for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) GetPropertySoftEnforcement() (value bool, err error) {
	retValue, err := instance.GetProperty("SoftEnforcement")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUpdateDomains sets the value of UpdateDomains for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) SetPropertyUpdateDomains(value uint32) (err error) {
	return instance.SetProperty("UpdateDomains", value)
}

// GetUpdateDomains gets the value of UpdateDomains for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) GetPropertyUpdateDomains() (value uint32, err error) {
	retValue, err := instance.GetProperty("UpdateDomains")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkloads sets the value of Workloads for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) SetPropertyWorkloads(value []uint64) (err error) {
	return instance.SetProperty("Workloads", value)
}

// GetWorkloads gets the value of Workloads for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) GetPropertyWorkloads() (value []uint64, err error) {
	retValue, err := instance.GetProperty("Workloads")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="AvailabilitySetName" type="string "></param>
// <param name="FaultDomains" type="uint32 "></param>
// <param name="FDType" type="uint32 "></param>
// <param name="Flags" type="uint32 "></param>
// <param name="ParticipantName" type="string []"></param>
// <param name="SoftEnforcement" type="bool "></param>
// <param name="UpdateDomains" type="uint32 "></param>

// <param name="CreatedAvailabilitySet" type="MSFT_ClusterSetAvailabilitySet "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetAvailabilitySet) CreateAvailabilitySet( /* IN */ AvailabilitySetName string,
	/* IN */ FaultDomains uint32,
	/* IN */ UpdateDomains uint32,
	/* IN */ FDType uint32,
	/* IN */ ParticipantName []string,
	/* IN */ SoftEnforcement bool,
	/* IN */ Flags uint32,
	/* OUT */ CreatedAvailabilitySet MSFT_ClusterSetAvailabilitySet) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateAvailabilitySet", AvailabilitySetName, FaultDomains, UpdateDomains, FDType, ParticipantName, SoftEnforcement, Flags)
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
func (instance *MSFT_ClusterSetAvailabilitySet) RemoveAvailabilitySet( /* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveAvailabilitySet", Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="ParticipantName" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetAvailabilitySet) AddParticipant( /* IN */ ParticipantName []string,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddParticipant", ParticipantName, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="ParticipantName" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetAvailabilitySet) RemoveParticipant( /* IN */ ParticipantName []string,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveParticipant", ParticipantName, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
