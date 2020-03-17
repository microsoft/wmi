// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// CIM_AllocationCapabilities struct
type CIM_AllocationCapabilities struct {
	CIM_Capabilities

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

// SetOtherResourceType sets the value of OtherResourceType for the instance
func (instance *CIM_AllocationCapabilities) SetPropertyOtherResourceType(value string) (err error) {
	return instance.SetProperty("OtherResourceType", value)
}

// GetOtherResourceType gets the value of OtherResourceType for the instance
func (instance *CIM_AllocationCapabilities) GetPropertyOtherResourceType() (value string, err error) {
	retValue, err := instance.GetProperty("OtherResourceType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequestTypesSupported sets the value of RequestTypesSupported for the instance
func (instance *CIM_AllocationCapabilities) SetPropertyRequestTypesSupported(value AllocationCapabilities_RequestTypesSupported) (err error) {
	return instance.SetProperty("RequestTypesSupported", value)
}

// GetRequestTypesSupported gets the value of RequestTypesSupported for the instance
func (instance *CIM_AllocationCapabilities) GetPropertyRequestTypesSupported() (value AllocationCapabilities_RequestTypesSupported, err error) {
	retValue, err := instance.GetProperty("RequestTypesSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(AllocationCapabilities_RequestTypesSupported)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceSubType sets the value of ResourceSubType for the instance
func (instance *CIM_AllocationCapabilities) SetPropertyResourceSubType(value string) (err error) {
	return instance.SetProperty("ResourceSubType", value)
}

// GetResourceSubType gets the value of ResourceSubType for the instance
func (instance *CIM_AllocationCapabilities) GetPropertyResourceSubType() (value string, err error) {
	retValue, err := instance.GetProperty("ResourceSubType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceType sets the value of ResourceType for the instance
func (instance *CIM_AllocationCapabilities) SetPropertyResourceType(value AllocationCapabilities_ResourceType) (err error) {
	return instance.SetProperty("ResourceType", value)
}

// GetResourceType gets the value of ResourceType for the instance
func (instance *CIM_AllocationCapabilities) GetPropertyResourceType() (value AllocationCapabilities_ResourceType, err error) {
	retValue, err := instance.GetProperty("ResourceType")
	if err != nil {
		return
	}
	value, ok := retValue.(AllocationCapabilities_ResourceType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSharingMode sets the value of SharingMode for the instance
func (instance *CIM_AllocationCapabilities) SetPropertySharingMode(value AllocationCapabilities_SharingMode) (err error) {
	return instance.SetProperty("SharingMode", value)
}

// GetSharingMode gets the value of SharingMode for the instance
func (instance *CIM_AllocationCapabilities) GetPropertySharingMode() (value AllocationCapabilities_SharingMode, err error) {
	retValue, err := instance.GetProperty("SharingMode")
	if err != nil {
		return
	}
	value, ok := retValue.(AllocationCapabilities_SharingMode)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedAddStates sets the value of SupportedAddStates for the instance
func (instance *CIM_AllocationCapabilities) SetPropertySupportedAddStates(value []AllocationCapabilities_SupportedAddStates) (err error) {
	return instance.SetProperty("SupportedAddStates", value)
}

// GetSupportedAddStates gets the value of SupportedAddStates for the instance
func (instance *CIM_AllocationCapabilities) GetPropertySupportedAddStates() (value []AllocationCapabilities_SupportedAddStates, err error) {
	retValue, err := instance.GetProperty("SupportedAddStates")
	if err != nil {
		return
	}
	value, ok := retValue.([]AllocationCapabilities_SupportedAddStates)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedRemoveStates sets the value of SupportedRemoveStates for the instance
func (instance *CIM_AllocationCapabilities) SetPropertySupportedRemoveStates(value []AllocationCapabilities_SupportedRemoveStates) (err error) {
	return instance.SetProperty("SupportedRemoveStates", value)
}

// GetSupportedRemoveStates gets the value of SupportedRemoveStates for the instance
func (instance *CIM_AllocationCapabilities) GetPropertySupportedRemoveStates() (value []AllocationCapabilities_SupportedRemoveStates, err error) {
	retValue, err := instance.GetProperty("SupportedRemoveStates")
	if err != nil {
		return
	}
	value, ok := retValue.([]AllocationCapabilities_SupportedRemoveStates)
	if !ok {
		// TODO: Set an error
	}
	return
}
