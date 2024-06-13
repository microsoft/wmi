// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_AvailableStoragePool struct
type MSCluster_AvailableStoragePool struct {
	*MSCluster_LogicalElement

	//
	Attributes uint64

	//
	ConnectedNodes []string

	//
	HealthStatus uint32

	//
	Id string

	//
	QuorumStatus uint32

	//
	TotalSize uint64

	//
	Usage uint64
}

func NewMSCluster_AvailableStoragePoolEx1(instance *cim.WmiInstance) (newInstance *MSCluster_AvailableStoragePool, err error) {
	tmp, err := NewMSCluster_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_AvailableStoragePool{
		MSCluster_LogicalElement: tmp,
	}
	return
}

func NewMSCluster_AvailableStoragePoolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_AvailableStoragePool, err error) {
	tmp, err := NewMSCluster_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_AvailableStoragePool{
		MSCluster_LogicalElement: tmp,
	}
	return
}

// SetAttributes sets the value of Attributes for the instance
func (instance *MSCluster_AvailableStoragePool) SetPropertyAttributes(value uint64) (err error) {
	return instance.SetProperty("Attributes", (value))
}

// GetAttributes gets the value of Attributes for the instance
func (instance *MSCluster_AvailableStoragePool) GetPropertyAttributes() (value uint64, err error) {
	retValue, err := instance.GetProperty("Attributes")
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

// SetConnectedNodes sets the value of ConnectedNodes for the instance
func (instance *MSCluster_AvailableStoragePool) SetPropertyConnectedNodes(value []string) (err error) {
	return instance.SetProperty("ConnectedNodes", (value))
}

// GetConnectedNodes gets the value of ConnectedNodes for the instance
func (instance *MSCluster_AvailableStoragePool) GetPropertyConnectedNodes() (value []string, err error) {
	retValue, err := instance.GetProperty("ConnectedNodes")
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

// SetHealthStatus sets the value of HealthStatus for the instance
func (instance *MSCluster_AvailableStoragePool) SetPropertyHealthStatus(value uint32) (err error) {
	return instance.SetProperty("HealthStatus", (value))
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *MSCluster_AvailableStoragePool) GetPropertyHealthStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("HealthStatus")
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
func (instance *MSCluster_AvailableStoragePool) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *MSCluster_AvailableStoragePool) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetQuorumStatus sets the value of QuorumStatus for the instance
func (instance *MSCluster_AvailableStoragePool) SetPropertyQuorumStatus(value uint32) (err error) {
	return instance.SetProperty("QuorumStatus", (value))
}

// GetQuorumStatus gets the value of QuorumStatus for the instance
func (instance *MSCluster_AvailableStoragePool) GetPropertyQuorumStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuorumStatus")
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

// SetTotalSize sets the value of TotalSize for the instance
func (instance *MSCluster_AvailableStoragePool) SetPropertyTotalSize(value uint64) (err error) {
	return instance.SetProperty("TotalSize", (value))
}

// GetTotalSize gets the value of TotalSize for the instance
func (instance *MSCluster_AvailableStoragePool) GetPropertyTotalSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalSize")
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

// SetUsage sets the value of Usage for the instance
func (instance *MSCluster_AvailableStoragePool) SetPropertyUsage(value uint64) (err error) {
	return instance.SetProperty("Usage", (value))
}

// GetUsage gets the value of Usage for the instance
func (instance *MSCluster_AvailableStoragePool) GetPropertyUsage() (value uint64, err error) {
	retValue, err := instance.GetProperty("Usage")
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

//
func (instance *MSCluster_AvailableStoragePool) AddToCluster() (err error) {
	_, err = instance.InvokeMethodWithReturn("AddToCluster")
	if err != nil {
		return
	}
	return

}
