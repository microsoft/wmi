// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_EventClusterCallback struct
type MSCluster_EventClusterCallback struct {
	*MSCluster_Event

	//
	ObjectName string

	//
	PercentComplete int32

	//
	PhaseSeverity int32

	//
	PhaseType int32

	//
	SetupPhase int32

	//
	Status int32
}

func NewMSCluster_EventClusterCallbackEx1(instance *cim.WmiInstance) (newInstance *MSCluster_EventClusterCallback, err error) {
	tmp, err := NewMSCluster_EventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_EventClusterCallback{
		MSCluster_Event: tmp,
	}
	return
}

func NewMSCluster_EventClusterCallbackEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_EventClusterCallback, err error) {
	tmp, err := NewMSCluster_EventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_EventClusterCallback{
		MSCluster_Event: tmp,
	}
	return
}

// SetObjectName sets the value of ObjectName for the instance
func (instance *MSCluster_EventClusterCallback) SetPropertyObjectName(value string) (err error) {
	return instance.SetProperty("ObjectName", (value))
}

// GetObjectName gets the value of ObjectName for the instance
func (instance *MSCluster_EventClusterCallback) GetPropertyObjectName() (value string, err error) {
	retValue, err := instance.GetProperty("ObjectName")
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

// SetPercentComplete sets the value of PercentComplete for the instance
func (instance *MSCluster_EventClusterCallback) SetPropertyPercentComplete(value int32) (err error) {
	return instance.SetProperty("PercentComplete", (value))
}

// GetPercentComplete gets the value of PercentComplete for the instance
func (instance *MSCluster_EventClusterCallback) GetPropertyPercentComplete() (value int32, err error) {
	retValue, err := instance.GetProperty("PercentComplete")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetPhaseSeverity sets the value of PhaseSeverity for the instance
func (instance *MSCluster_EventClusterCallback) SetPropertyPhaseSeverity(value int32) (err error) {
	return instance.SetProperty("PhaseSeverity", (value))
}

// GetPhaseSeverity gets the value of PhaseSeverity for the instance
func (instance *MSCluster_EventClusterCallback) GetPropertyPhaseSeverity() (value int32, err error) {
	retValue, err := instance.GetProperty("PhaseSeverity")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetPhaseType sets the value of PhaseType for the instance
func (instance *MSCluster_EventClusterCallback) SetPropertyPhaseType(value int32) (err error) {
	return instance.SetProperty("PhaseType", (value))
}

// GetPhaseType gets the value of PhaseType for the instance
func (instance *MSCluster_EventClusterCallback) GetPropertyPhaseType() (value int32, err error) {
	retValue, err := instance.GetProperty("PhaseType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetSetupPhase sets the value of SetupPhase for the instance
func (instance *MSCluster_EventClusterCallback) SetPropertySetupPhase(value int32) (err error) {
	return instance.SetProperty("SetupPhase", (value))
}

// GetSetupPhase gets the value of SetupPhase for the instance
func (instance *MSCluster_EventClusterCallback) GetPropertySetupPhase() (value int32, err error) {
	retValue, err := instance.GetProperty("SetupPhase")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetStatus sets the value of Status for the instance
func (instance *MSCluster_EventClusterCallback) SetPropertyStatus(value int32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MSCluster_EventClusterCallback) GetPropertyStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}
