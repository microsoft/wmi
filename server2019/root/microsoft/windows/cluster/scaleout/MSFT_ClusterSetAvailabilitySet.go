// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.Cluster.Scaleout
//////////////////////////////////////////////
package scaleout

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_ClusterSetAvailabilitySet struct
type MSFT_ClusterSetAvailabilitySet struct {
	*cim.WmiInstance

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

func NewMSFT_ClusterSetAvailabilitySetEx1(instance *cim.WmiInstance) (newInstance *MSFT_ClusterSetAvailabilitySet, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ClusterSetAvailabilitySet{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ClusterSetAvailabilitySetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ClusterSetAvailabilitySet, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ClusterSetAvailabilitySet{
		WmiInstance: tmp,
	}
	return
}

// SetAvailabilitySetName sets the value of AvailabilitySetName for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) SetPropertyAvailabilitySetName(value string) (err error) {
	return instance.SetProperty("AvailabilitySetName", (value))
}

// GetAvailabilitySetName gets the value of AvailabilitySetName for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) GetPropertyAvailabilitySetName() (value string, err error) {
	retValue, err := instance.GetProperty("AvailabilitySetName")
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

// SetDomainAssignments sets the value of DomainAssignments for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) SetPropertyDomainAssignments(value []string) (err error) {
	return instance.SetProperty("DomainAssignments", (value))
}

// GetDomainAssignments gets the value of DomainAssignments for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) GetPropertyDomainAssignments() (value []string, err error) {
	retValue, err := instance.GetProperty("DomainAssignments")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetFaultDomains sets the value of FaultDomains for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) SetPropertyFaultDomains(value uint32) (err error) {
	return instance.SetProperty("FaultDomains", (value))
}

// GetFaultDomains gets the value of FaultDomains for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) GetPropertyFaultDomains() (value uint32, err error) {
	retValue, err := instance.GetProperty("FaultDomains")
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

// SetFaultDomainType sets the value of FaultDomainType for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) SetPropertyFaultDomainType(value uint32) (err error) {
	return instance.SetProperty("FaultDomainType", (value))
}

// GetFaultDomainType gets the value of FaultDomainType for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) GetPropertyFaultDomainType() (value uint32, err error) {
	retValue, err := instance.GetProperty("FaultDomainType")
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

// SetId sets the value of Id for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) SetPropertyId(value uint64) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) GetPropertyId() (value uint64, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetParticipantIds sets the value of ParticipantIds for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) SetPropertyParticipantIds(value []uint64) (err error) {
	return instance.SetProperty("ParticipantIds", (value))
}

// GetParticipantIds gets the value of ParticipantIds for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) GetPropertyParticipantIds() (value []uint64, err error) {
	retValue, err := instance.GetProperty("ParticipantIds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint64)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint64(valuetmp))
	}

	return
}

// SetSoftEnforcement sets the value of SoftEnforcement for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) SetPropertySoftEnforcement(value bool) (err error) {
	return instance.SetProperty("SoftEnforcement", (value))
}

// GetSoftEnforcement gets the value of SoftEnforcement for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) GetPropertySoftEnforcement() (value bool, err error) {
	retValue, err := instance.GetProperty("SoftEnforcement")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetUpdateDomains sets the value of UpdateDomains for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) SetPropertyUpdateDomains(value uint32) (err error) {
	return instance.SetProperty("UpdateDomains", (value))
}

// GetUpdateDomains gets the value of UpdateDomains for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) GetPropertyUpdateDomains() (value uint32, err error) {
	retValue, err := instance.GetProperty("UpdateDomains")
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

// SetWorkloads sets the value of Workloads for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) SetPropertyWorkloads(value []uint64) (err error) {
	return instance.SetProperty("Workloads", (value))
}

// GetWorkloads gets the value of Workloads for the instance
func (instance *MSFT_ClusterSetAvailabilitySet) GetPropertyWorkloads() (value []uint64, err error) {
	retValue, err := instance.GetProperty("Workloads")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint64)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint64(valuetmp))
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
