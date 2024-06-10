// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
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

// MSCluster_HealthFault struct
type MSCluster_HealthFault struct {
	*cim.WmiInstance

	//
	FaultId string

	//
	FaultingObjectDescription string

	//
	FaultingObjectLocation string

	//
	FaultingObjectType string

	//
	FaultingObjectUniqueId string

	//
	FaultTime string

	//
	FaultType string

	//
	PerceivedSeverity uint16

	//
	Reason string

	//
	RecommendedActions []string
}

func NewMSCluster_HealthFaultEx1(instance *cim.WmiInstance) (newInstance *MSCluster_HealthFault, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSCluster_HealthFault{
		WmiInstance: tmp,
	}
	return
}

func NewMSCluster_HealthFaultEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_HealthFault, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_HealthFault{
		WmiInstance: tmp,
	}
	return
}

// SetFaultId sets the value of FaultId for the instance
func (instance *MSCluster_HealthFault) SetPropertyFaultId(value string) (err error) {
	return instance.SetProperty("FaultId", (value))
}

// GetFaultId gets the value of FaultId for the instance
func (instance *MSCluster_HealthFault) GetPropertyFaultId() (value string, err error) {
	retValue, err := instance.GetProperty("FaultId")
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

// SetFaultingObjectDescription sets the value of FaultingObjectDescription for the instance
func (instance *MSCluster_HealthFault) SetPropertyFaultingObjectDescription(value string) (err error) {
	return instance.SetProperty("FaultingObjectDescription", (value))
}

// GetFaultingObjectDescription gets the value of FaultingObjectDescription for the instance
func (instance *MSCluster_HealthFault) GetPropertyFaultingObjectDescription() (value string, err error) {
	retValue, err := instance.GetProperty("FaultingObjectDescription")
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

// SetFaultingObjectLocation sets the value of FaultingObjectLocation for the instance
func (instance *MSCluster_HealthFault) SetPropertyFaultingObjectLocation(value string) (err error) {
	return instance.SetProperty("FaultingObjectLocation", (value))
}

// GetFaultingObjectLocation gets the value of FaultingObjectLocation for the instance
func (instance *MSCluster_HealthFault) GetPropertyFaultingObjectLocation() (value string, err error) {
	retValue, err := instance.GetProperty("FaultingObjectLocation")
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

// SetFaultingObjectType sets the value of FaultingObjectType for the instance
func (instance *MSCluster_HealthFault) SetPropertyFaultingObjectType(value string) (err error) {
	return instance.SetProperty("FaultingObjectType", (value))
}

// GetFaultingObjectType gets the value of FaultingObjectType for the instance
func (instance *MSCluster_HealthFault) GetPropertyFaultingObjectType() (value string, err error) {
	retValue, err := instance.GetProperty("FaultingObjectType")
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

// SetFaultingObjectUniqueId sets the value of FaultingObjectUniqueId for the instance
func (instance *MSCluster_HealthFault) SetPropertyFaultingObjectUniqueId(value string) (err error) {
	return instance.SetProperty("FaultingObjectUniqueId", (value))
}

// GetFaultingObjectUniqueId gets the value of FaultingObjectUniqueId for the instance
func (instance *MSCluster_HealthFault) GetPropertyFaultingObjectUniqueId() (value string, err error) {
	retValue, err := instance.GetProperty("FaultingObjectUniqueId")
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

// SetFaultTime sets the value of FaultTime for the instance
func (instance *MSCluster_HealthFault) SetPropertyFaultTime(value string) (err error) {
	return instance.SetProperty("FaultTime", (value))
}

// GetFaultTime gets the value of FaultTime for the instance
func (instance *MSCluster_HealthFault) GetPropertyFaultTime() (value string, err error) {
	retValue, err := instance.GetProperty("FaultTime")
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

// SetFaultType sets the value of FaultType for the instance
func (instance *MSCluster_HealthFault) SetPropertyFaultType(value string) (err error) {
	return instance.SetProperty("FaultType", (value))
}

// GetFaultType gets the value of FaultType for the instance
func (instance *MSCluster_HealthFault) GetPropertyFaultType() (value string, err error) {
	retValue, err := instance.GetProperty("FaultType")
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

// SetPerceivedSeverity sets the value of PerceivedSeverity for the instance
func (instance *MSCluster_HealthFault) SetPropertyPerceivedSeverity(value uint16) (err error) {
	return instance.SetProperty("PerceivedSeverity", (value))
}

// GetPerceivedSeverity gets the value of PerceivedSeverity for the instance
func (instance *MSCluster_HealthFault) GetPropertyPerceivedSeverity() (value uint16, err error) {
	retValue, err := instance.GetProperty("PerceivedSeverity")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetReason sets the value of Reason for the instance
func (instance *MSCluster_HealthFault) SetPropertyReason(value string) (err error) {
	return instance.SetProperty("Reason", (value))
}

// GetReason gets the value of Reason for the instance
func (instance *MSCluster_HealthFault) GetPropertyReason() (value string, err error) {
	retValue, err := instance.GetProperty("Reason")
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

// SetRecommendedActions sets the value of RecommendedActions for the instance
func (instance *MSCluster_HealthFault) SetPropertyRecommendedActions(value []string) (err error) {
	return instance.SetProperty("RecommendedActions", (value))
}

// GetRecommendedActions gets the value of RecommendedActions for the instance
func (instance *MSCluster_HealthFault) GetPropertyRecommendedActions() (value []string, err error) {
	retValue, err := instance.GetProperty("RecommendedActions")
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
