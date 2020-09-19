// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_GuestClusterInformation struct
type Msvm_GuestClusterInformation struct {
	*cim.WmiInstance

	//
	ClusterId string

	//
	ClusterSize uint16

	//
	IsActiveActive []bool

	//
	IsClustered []bool

	//
	IsOnline []bool

	//
	IsOwned []bool

	//
	LastResourceMoveTime uint64

	//
	SharedVirtualHardDiskPaths []string

	//
	SharedVirtualHardDisks []string
}

func NewMsvm_GuestClusterInformationEx1(instance *cim.WmiInstance) (newInstance *Msvm_GuestClusterInformation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Msvm_GuestClusterInformation{
		WmiInstance: tmp,
	}
	return
}

func NewMsvm_GuestClusterInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_GuestClusterInformation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_GuestClusterInformation{
		WmiInstance: tmp,
	}
	return
}

// SetClusterId sets the value of ClusterId for the instance
func (instance *Msvm_GuestClusterInformation) SetPropertyClusterId(value string) (err error) {
	return instance.SetProperty("ClusterId", (value))
}

// GetClusterId gets the value of ClusterId for the instance
func (instance *Msvm_GuestClusterInformation) GetPropertyClusterId() (value string, err error) {
	retValue, err := instance.GetProperty("ClusterId")
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

// SetClusterSize sets the value of ClusterSize for the instance
func (instance *Msvm_GuestClusterInformation) SetPropertyClusterSize(value uint16) (err error) {
	return instance.SetProperty("ClusterSize", (value))
}

// GetClusterSize gets the value of ClusterSize for the instance
func (instance *Msvm_GuestClusterInformation) GetPropertyClusterSize() (value uint16, err error) {
	retValue, err := instance.GetProperty("ClusterSize")
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

// SetIsActiveActive sets the value of IsActiveActive for the instance
func (instance *Msvm_GuestClusterInformation) SetPropertyIsActiveActive(value []bool) (err error) {
	return instance.SetProperty("IsActiveActive", (value))
}

// GetIsActiveActive gets the value of IsActiveActive for the instance
func (instance *Msvm_GuestClusterInformation) GetPropertyIsActiveActive() (value []bool, err error) {
	retValue, err := instance.GetProperty("IsActiveActive")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(bool)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, bool(valuetmp))
	}

	return
}

// SetIsClustered sets the value of IsClustered for the instance
func (instance *Msvm_GuestClusterInformation) SetPropertyIsClustered(value []bool) (err error) {
	return instance.SetProperty("IsClustered", (value))
}

// GetIsClustered gets the value of IsClustered for the instance
func (instance *Msvm_GuestClusterInformation) GetPropertyIsClustered() (value []bool, err error) {
	retValue, err := instance.GetProperty("IsClustered")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(bool)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, bool(valuetmp))
	}

	return
}

// SetIsOnline sets the value of IsOnline for the instance
func (instance *Msvm_GuestClusterInformation) SetPropertyIsOnline(value []bool) (err error) {
	return instance.SetProperty("IsOnline", (value))
}

// GetIsOnline gets the value of IsOnline for the instance
func (instance *Msvm_GuestClusterInformation) GetPropertyIsOnline() (value []bool, err error) {
	retValue, err := instance.GetProperty("IsOnline")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(bool)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, bool(valuetmp))
	}

	return
}

// SetIsOwned sets the value of IsOwned for the instance
func (instance *Msvm_GuestClusterInformation) SetPropertyIsOwned(value []bool) (err error) {
	return instance.SetProperty("IsOwned", (value))
}

// GetIsOwned gets the value of IsOwned for the instance
func (instance *Msvm_GuestClusterInformation) GetPropertyIsOwned() (value []bool, err error) {
	retValue, err := instance.GetProperty("IsOwned")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(bool)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, bool(valuetmp))
	}

	return
}

// SetLastResourceMoveTime sets the value of LastResourceMoveTime for the instance
func (instance *Msvm_GuestClusterInformation) SetPropertyLastResourceMoveTime(value uint64) (err error) {
	return instance.SetProperty("LastResourceMoveTime", (value))
}

// GetLastResourceMoveTime gets the value of LastResourceMoveTime for the instance
func (instance *Msvm_GuestClusterInformation) GetPropertyLastResourceMoveTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("LastResourceMoveTime")
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

// SetSharedVirtualHardDiskPaths sets the value of SharedVirtualHardDiskPaths for the instance
func (instance *Msvm_GuestClusterInformation) SetPropertySharedVirtualHardDiskPaths(value []string) (err error) {
	return instance.SetProperty("SharedVirtualHardDiskPaths", (value))
}

// GetSharedVirtualHardDiskPaths gets the value of SharedVirtualHardDiskPaths for the instance
func (instance *Msvm_GuestClusterInformation) GetPropertySharedVirtualHardDiskPaths() (value []string, err error) {
	retValue, err := instance.GetProperty("SharedVirtualHardDiskPaths")
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

// SetSharedVirtualHardDisks sets the value of SharedVirtualHardDisks for the instance
func (instance *Msvm_GuestClusterInformation) SetPropertySharedVirtualHardDisks(value []string) (err error) {
	return instance.SetProperty("SharedVirtualHardDisks", (value))
}

// GetSharedVirtualHardDisks gets the value of SharedVirtualHardDisks for the instance
func (instance *Msvm_GuestClusterInformation) GetPropertySharedVirtualHardDisks() (value []string, err error) {
	retValue, err := instance.GetProperty("SharedVirtualHardDisks")
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
