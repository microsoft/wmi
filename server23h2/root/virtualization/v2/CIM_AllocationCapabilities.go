// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_AllocationCapabilities struct
type CIM_AllocationCapabilities struct {
	*CIM_Capabilities

	// A string that describes the resource type when a well defined value is not available and ResourceType has the value "Other".
	OtherResourceType string

	// Indicates whether requesting a specific resource is supported
	///"Specific" -- request can include a request for specific resource
	///"General" -- request does not include specific resource
	///"Both" -- both specific and general requests are supported.
	RequestTypesSupported AllocationCapabilities_RequestTypesSupported

	// A string describing an implementation specific sub-type for this resource. For example, this may be used to distinguish different models of the same resource type.
	ResourceSubType string

	// The type of resource this allocation setting represents.
	ResourceType AllocationCapabilities_ResourceType

	// Indicates how access to underlying resource is granted:
	///"Dedicated" -- exclusive access to underlying resource
	///"Shared" -- shared use of underlying resource.
	///Actual quantity is controlled by min, max size, weights, etc.
	SharingMode AllocationCapabilities_SharingMode

	// Indicates the states that the System, to which the resource will be associated via SystemDevice, may be in when a new resource is created.
	SupportedAddStates []AllocationCapabilities_SupportedAddStates

	// Indicates the states that the System, to which the resource is associated via SystemDevice, may be in when a the resource is removed .
	SupportedRemoveStates []AllocationCapabilities_SupportedRemoveStates
}

func NewCIM_AllocationCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *CIM_AllocationCapabilities, err error) {
	tmp, err := NewCIM_CapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_AllocationCapabilities{
		CIM_Capabilities: tmp,
	}
	return
}

func NewCIM_AllocationCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_AllocationCapabilities, err error) {
	tmp, err := NewCIM_CapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_AllocationCapabilities{
		CIM_Capabilities: tmp,
	}
	return
}

// SetOtherResourceType sets the value of OtherResourceType for the instance
func (instance *CIM_AllocationCapabilities) SetPropertyOtherResourceType(value string) (err error) {
	return instance.SetProperty("OtherResourceType", (value))
}

// GetOtherResourceType gets the value of OtherResourceType for the instance
func (instance *CIM_AllocationCapabilities) GetPropertyOtherResourceType() (value string, err error) {
	retValue, err := instance.GetProperty("OtherResourceType")
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

// SetRequestTypesSupported sets the value of RequestTypesSupported for the instance
func (instance *CIM_AllocationCapabilities) SetPropertyRequestTypesSupported(value AllocationCapabilities_RequestTypesSupported) (err error) {
	return instance.SetProperty("RequestTypesSupported", (value))
}

// GetRequestTypesSupported gets the value of RequestTypesSupported for the instance
func (instance *CIM_AllocationCapabilities) GetPropertyRequestTypesSupported() (value AllocationCapabilities_RequestTypesSupported, err error) {
	retValue, err := instance.GetProperty("RequestTypesSupported")
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

	value = AllocationCapabilities_RequestTypesSupported(valuetmp)

	return
}

// SetResourceSubType sets the value of ResourceSubType for the instance
func (instance *CIM_AllocationCapabilities) SetPropertyResourceSubType(value string) (err error) {
	return instance.SetProperty("ResourceSubType", (value))
}

// GetResourceSubType gets the value of ResourceSubType for the instance
func (instance *CIM_AllocationCapabilities) GetPropertyResourceSubType() (value string, err error) {
	retValue, err := instance.GetProperty("ResourceSubType")
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

// SetResourceType sets the value of ResourceType for the instance
func (instance *CIM_AllocationCapabilities) SetPropertyResourceType(value AllocationCapabilities_ResourceType) (err error) {
	return instance.SetProperty("ResourceType", (value))
}

// GetResourceType gets the value of ResourceType for the instance
func (instance *CIM_AllocationCapabilities) GetPropertyResourceType() (value AllocationCapabilities_ResourceType, err error) {
	retValue, err := instance.GetProperty("ResourceType")
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

	value = AllocationCapabilities_ResourceType(valuetmp)

	return
}

// SetSharingMode sets the value of SharingMode for the instance
func (instance *CIM_AllocationCapabilities) SetPropertySharingMode(value AllocationCapabilities_SharingMode) (err error) {
	return instance.SetProperty("SharingMode", (value))
}

// GetSharingMode gets the value of SharingMode for the instance
func (instance *CIM_AllocationCapabilities) GetPropertySharingMode() (value AllocationCapabilities_SharingMode, err error) {
	retValue, err := instance.GetProperty("SharingMode")
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

	value = AllocationCapabilities_SharingMode(valuetmp)

	return
}

// SetSupportedAddStates sets the value of SupportedAddStates for the instance
func (instance *CIM_AllocationCapabilities) SetPropertySupportedAddStates(value []AllocationCapabilities_SupportedAddStates) (err error) {
	return instance.SetProperty("SupportedAddStates", (value))
}

// GetSupportedAddStates gets the value of SupportedAddStates for the instance
func (instance *CIM_AllocationCapabilities) GetPropertySupportedAddStates() (value []AllocationCapabilities_SupportedAddStates, err error) {
	retValue, err := instance.GetProperty("SupportedAddStates")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, AllocationCapabilities_SupportedAddStates(valuetmp))
	}

	return
}

// SetSupportedRemoveStates sets the value of SupportedRemoveStates for the instance
func (instance *CIM_AllocationCapabilities) SetPropertySupportedRemoveStates(value []AllocationCapabilities_SupportedRemoveStates) (err error) {
	return instance.SetProperty("SupportedRemoveStates", (value))
}

// GetSupportedRemoveStates gets the value of SupportedRemoveStates for the instance
func (instance *CIM_AllocationCapabilities) GetPropertySupportedRemoveStates() (value []AllocationCapabilities_SupportedRemoveStates, err error) {
	retValue, err := instance.GetProperty("SupportedRemoveStates")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, AllocationCapabilities_SupportedRemoveStates(valuetmp))
	}

	return
}
