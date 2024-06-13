// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_DynamicForwardingEntry struct
type CIM_DynamicForwardingEntry struct {
	*CIM_LogicalElement

	// CreationClassName indicates the name of the class or the subclass used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
	CreationClassName string

	// The status of the entry.
	DynamicStatus DynamicForwardingEntry_DynamicStatus

	// Unicast MAC address for which the TransparentBridging Service has forwarding and/or filtering information. Note that the MAC address is formatted as twelve hexadecimal digits (e.g., "010203040506"), with each pair representing one of the six octets of the MAC address in "canonical" bit order according to RFC 2469.
	MACAddress string

	// The scoping Service's CreationClassName.
	ServiceCreationClassName string

	// The scoping Service's Name.
	ServiceName string

	// The scoping System's CreationClassName.
	SystemCreationClassName string

	// The scoping System's Name.
	SystemName string
}

func NewCIM_DynamicForwardingEntryEx1(instance *cim.WmiInstance) (newInstance *CIM_DynamicForwardingEntry, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_DynamicForwardingEntry{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewCIM_DynamicForwardingEntryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_DynamicForwardingEntry, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_DynamicForwardingEntry{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *CIM_DynamicForwardingEntry) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", (value))
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *CIM_DynamicForwardingEntry) GetPropertyCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CreationClassName")
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

// SetDynamicStatus sets the value of DynamicStatus for the instance
func (instance *CIM_DynamicForwardingEntry) SetPropertyDynamicStatus(value DynamicForwardingEntry_DynamicStatus) (err error) {
	return instance.SetProperty("DynamicStatus", (value))
}

// GetDynamicStatus gets the value of DynamicStatus for the instance
func (instance *CIM_DynamicForwardingEntry) GetPropertyDynamicStatus() (value DynamicForwardingEntry_DynamicStatus, err error) {
	retValue, err := instance.GetProperty("DynamicStatus")
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

	value = DynamicForwardingEntry_DynamicStatus(valuetmp)

	return
}

// SetMACAddress sets the value of MACAddress for the instance
func (instance *CIM_DynamicForwardingEntry) SetPropertyMACAddress(value string) (err error) {
	return instance.SetProperty("MACAddress", (value))
}

// GetMACAddress gets the value of MACAddress for the instance
func (instance *CIM_DynamicForwardingEntry) GetPropertyMACAddress() (value string, err error) {
	retValue, err := instance.GetProperty("MACAddress")
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

// SetServiceCreationClassName sets the value of ServiceCreationClassName for the instance
func (instance *CIM_DynamicForwardingEntry) SetPropertyServiceCreationClassName(value string) (err error) {
	return instance.SetProperty("ServiceCreationClassName", (value))
}

// GetServiceCreationClassName gets the value of ServiceCreationClassName for the instance
func (instance *CIM_DynamicForwardingEntry) GetPropertyServiceCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("ServiceCreationClassName")
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

// SetServiceName sets the value of ServiceName for the instance
func (instance *CIM_DynamicForwardingEntry) SetPropertyServiceName(value string) (err error) {
	return instance.SetProperty("ServiceName", (value))
}

// GetServiceName gets the value of ServiceName for the instance
func (instance *CIM_DynamicForwardingEntry) GetPropertyServiceName() (value string, err error) {
	retValue, err := instance.GetProperty("ServiceName")
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

// SetSystemCreationClassName sets the value of SystemCreationClassName for the instance
func (instance *CIM_DynamicForwardingEntry) SetPropertySystemCreationClassName(value string) (err error) {
	return instance.SetProperty("SystemCreationClassName", (value))
}

// GetSystemCreationClassName gets the value of SystemCreationClassName for the instance
func (instance *CIM_DynamicForwardingEntry) GetPropertySystemCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemCreationClassName")
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

// SetSystemName sets the value of SystemName for the instance
func (instance *CIM_DynamicForwardingEntry) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", (value))
}

// GetSystemName gets the value of SystemName for the instance
func (instance *CIM_DynamicForwardingEntry) GetPropertySystemName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemName")
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
