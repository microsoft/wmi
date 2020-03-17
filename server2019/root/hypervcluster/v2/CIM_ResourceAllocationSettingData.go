// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// CIM_ResourceAllocationSettingData struct
type CIM_ResourceAllocationSettingData struct {
	CIM_SettingData

	// The address of the resource. For example, the MAC address of a Ethernet port.
	Address string

	// Describes the address of this resource in the context of the Parent. The Parent/AddressOnParent properties are used to describe the controller relationship as well the ordering of devices on a controller.For example, if the parent is a PCI Controller, this property would specify the PCI slot of this child device.
	AddressOnParent string

	// This property specifies the units of allocation used by the Reservation and Limit properties. For example, when ResourceType=Processor, AllocationUnits may be set to hertz*10^6 or percent. When ResourceType=Memory, AllocationUnits may be set to bytes*10^3.
	///It is expected that profiles constrain the units that apply in context of particular resource types.
	///The value of this property shall be a legal value of the Programmatic Units qualifier as defined in Annex C.1 of DSP0004 V2.5 or later.
	AllocationUnits string

	// This property specifies if the resource will be automatically allocated. For example when set to true, when the consuming virtual computer system is powered on, this resource would be allocated. A value of false indicates the resource must be explicitly allocated. For example, the setting may represent removable media (cdrom, floppy, etc.) where at power on time, the media is not present. An explicit operation is required to allocate the resource.
	AutomaticAllocation bool

	// This property specifies if the resource will be automatically de-allocated. For example, when set to true, when the consuming virtual computer system is powered off, this resource would be de-allocated. When set to false, the resource will remain allocated and must be explicitly de-allocated.
	AutomaticDeallocation bool

	// The thing to which this resource is connected. For example, a named network or switch port.
	Connection []string

	// Describes the consumers visibility to the allocated resource.
	///A value of "Passed-Through" indicates the underlying or host resource is utilized and passed through to the consumer, possibly using partitioning. At least one item shall be present in the HostResource property.
	///A value of "Virtualized" indicates the resource is virtualized and may not map directly to an underlying/host resource. Some implementations may support specific assignment for virtualized resources, in which case the host resource(s) are exposed using the HostResource property.
	///A value of "Not represented" indicates a representation of the resource does not exist within the context of the resource consumer.
	ConsumerVisibility ResourceAllocationSettingData_ConsumerVisibility

	// This property exposes specific assignment of resources. Each non-null value of the HostResource property shall be formated as a URI per RFC3986.
	///If this resource is modeled then a value should be a WBEM URI (DSP0207). If the resource is not modeled then see the appropriate profile.
	///Profiles may further constrain the type of URI. A NULL value or empty array requests the implementation decide the kind of host resource.
	///If the virtual resource is mapped to more than oneunderlying resource, this property may be left NULL.
	///If NULL, the DeviceAllocatedFromPool or ResourceAllocationFromPool associations may be used to determine the pool of host resources this virtual resource may use. If specific assignment is utilized, all underlying resources used by this virtual resource should be listed.The kind of dependency is specified by the ConsumerVisibility and the MappingBehavior properties. Typically the array contains one item, however multiple host resources may be specified.
	///A client may set the value(s) to indicate that the requested virtual resource allocation be based on host resources that are identified by element values.
	HostResource []string

	// This property specifies the upper bound, or maximum amount of resource that will be granted for this allocation. For example, a system which supports memory paging may support setting the Limit of a Memory allocation below that of the VirtualQuantity, thus forcing paging to occur for this allocation.
	///The value of the Limit property is expressed in the unit specified by the value of the AllocationUnits property.
	Limit uint64

	// Specifies how this resource maps to underlying resourcesIf the HostResource array contains any entries, this property reflects how the resource maps to those specific resources.
	MappingBehavior ResourceAllocationSettingData_MappingBehavior

	// A string that describes the resource type when a well defined value is not available and ResourceType has the value "Other".
	OtherResourceType string

	// The Parent of the resource. For example, a controller for the current allocation
	Parent string

	// This property specifies which ResourcePool the resource is currently allocated from, or which ResourcePool the resource will be allocated from when the allocation occurs.
	PoolID string

	// This property specifies the amount of resource guaranteed to be available for this allocation. On system which support over-commitment of resources, this value is typically used for admission control to prevent an an allocation from being accepted thus preventing starvation.
	///The value of the Reservation property is expressed in the unit specified by the value of the AllocationUnits property.
	Reservation uint64

	// A string describing an implementation specific sub-type for this resource. For example, this may be used to distinguish different models of the same resource type.The property value shall conform to this format (in ABNF): vs-type = dmtf-value / other-org-value / legacy-value dmtf-value = "DMTF:" defining-org ":" org-vs-type org-value = defining-org ":" org-vs-type
	///Where: dmtf-value: is a property value defined by DMTF and is defined in the description of this property. other-org-value: is a property value defined by a business entity other than DMTF and is not defined in the description of this property. legacy-value: is a property value defined by a business entity other than DMTF and is not defined in the description of this property. These values are permitted but recommended to be deprecated over time. defining-org:
	///is an identifier for the business entity that defines the virtual system type. It shall include a copyrighted, trademarked, or otherwise unique name that is owned by that business entity. It shall not be "DMTF" and shall not contain a colon (:). org-vs-type:
	///is an identifier for the virtual system type within the defining business entity. It shall be unique within the defining-org. It may use any character allowed for CIM strings, except for the following: U0000-U001F (Unicode C0 controls) U0020 (space), note that the reason is that OVF allows for multiple space-separated vs-type values in this property. U007F (Unicode C0 controls) U0080-U009F (Unicode C1 controls)
	///If there is a need to structure the value into segments, the segments should be separated with a single colon (:).
	///The values of this property shall be processed case sensitively. They are intended to be processed programmatically (instead of being a display name) and should be short.
	///The following DMTF values are defined: DMTF:unknown - the resource sub-type is unknown or cannot be determined
	///Developers should consult the relevant profile for defined values.
	ResourceSubType string

	// The type of resource this allocation setting represents.
	ResourceType ResourceAllocationSettingData_ResourceType

	// This property specifies the quantity of resources presented to the consumer. For example, when ResourceType=Processor, this property would reflect the number of discrete Processors presented to the virtual computer system. When ResourceType=Memory, this property could reflect the number of MB reported to the virtual computer system.
	///The value of the VirtualQuantity property should be expressed in units as defined by the value of the VirtualQuantityUnits property.
	VirtualQuantity uint64

	// This property specifies the units used by the VirtualQuantity property. For example
	///- if ResourceType=Processor, the value of the VirtualQuantityUnits property may be set to "count", indicating that the value of the VirtualQuantity property is expressed as a count.
	///- if ResourceType=Memory, the value of the VirtualQuantityUnits property may be set to "bytes*10^3", indicating that the value of the VirtualQuantity property is expressed in kilobyte.
	///It is expected that profiles constrain the units that apply in context of particular resource types.
	///The value of this property shall be a legal value of the Programmatic Units qualifier as defined in Annex C.1 of DSP0004 V2.5 or later.
	VirtualQuantityUnits string

	// This property specifies a relative priority for this allocation in relation to other allocations from the same ResourcePool. This property has no unit of measure, and is only relevant when compared to other allocations vying for the same host resources.
	Weight uint32
}

// SetAddress sets the value of Address for the instance
func (instance *CIM_ResourceAllocationSettingData) SetPropertyAddress(value string) (err error) {
	return instance.SetProperty("Address", value)
}

// GetAddress gets the value of Address for the instance
func (instance *CIM_ResourceAllocationSettingData) GetPropertyAddress() (value string, err error) {
	retValue, err := instance.GetProperty("Address")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAddressOnParent sets the value of AddressOnParent for the instance
func (instance *CIM_ResourceAllocationSettingData) SetPropertyAddressOnParent(value string) (err error) {
	return instance.SetProperty("AddressOnParent", value)
}

// GetAddressOnParent gets the value of AddressOnParent for the instance
func (instance *CIM_ResourceAllocationSettingData) GetPropertyAddressOnParent() (value string, err error) {
	retValue, err := instance.GetProperty("AddressOnParent")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllocationUnits sets the value of AllocationUnits for the instance
func (instance *CIM_ResourceAllocationSettingData) SetPropertyAllocationUnits(value string) (err error) {
	return instance.SetProperty("AllocationUnits", value)
}

// GetAllocationUnits gets the value of AllocationUnits for the instance
func (instance *CIM_ResourceAllocationSettingData) GetPropertyAllocationUnits() (value string, err error) {
	retValue, err := instance.GetProperty("AllocationUnits")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutomaticAllocation sets the value of AutomaticAllocation for the instance
func (instance *CIM_ResourceAllocationSettingData) SetPropertyAutomaticAllocation(value bool) (err error) {
	return instance.SetProperty("AutomaticAllocation", value)
}

// GetAutomaticAllocation gets the value of AutomaticAllocation for the instance
func (instance *CIM_ResourceAllocationSettingData) GetPropertyAutomaticAllocation() (value bool, err error) {
	retValue, err := instance.GetProperty("AutomaticAllocation")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutomaticDeallocation sets the value of AutomaticDeallocation for the instance
func (instance *CIM_ResourceAllocationSettingData) SetPropertyAutomaticDeallocation(value bool) (err error) {
	return instance.SetProperty("AutomaticDeallocation", value)
}

// GetAutomaticDeallocation gets the value of AutomaticDeallocation for the instance
func (instance *CIM_ResourceAllocationSettingData) GetPropertyAutomaticDeallocation() (value bool, err error) {
	retValue, err := instance.GetProperty("AutomaticDeallocation")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnection sets the value of Connection for the instance
func (instance *CIM_ResourceAllocationSettingData) SetPropertyConnection(value []string) (err error) {
	return instance.SetProperty("Connection", value)
}

// GetConnection gets the value of Connection for the instance
func (instance *CIM_ResourceAllocationSettingData) GetPropertyConnection() (value []string, err error) {
	retValue, err := instance.GetProperty("Connection")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConsumerVisibility sets the value of ConsumerVisibility for the instance
func (instance *CIM_ResourceAllocationSettingData) SetPropertyConsumerVisibility(value ResourceAllocationSettingData_ConsumerVisibility) (err error) {
	return instance.SetProperty("ConsumerVisibility", value)
}

// GetConsumerVisibility gets the value of ConsumerVisibility for the instance
func (instance *CIM_ResourceAllocationSettingData) GetPropertyConsumerVisibility() (value ResourceAllocationSettingData_ConsumerVisibility, err error) {
	retValue, err := instance.GetProperty("ConsumerVisibility")
	if err != nil {
		return
	}
	value, ok := retValue.(ResourceAllocationSettingData_ConsumerVisibility)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostResource sets the value of HostResource for the instance
func (instance *CIM_ResourceAllocationSettingData) SetPropertyHostResource(value []string) (err error) {
	return instance.SetProperty("HostResource", value)
}

// GetHostResource gets the value of HostResource for the instance
func (instance *CIM_ResourceAllocationSettingData) GetPropertyHostResource() (value []string, err error) {
	retValue, err := instance.GetProperty("HostResource")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLimit sets the value of Limit for the instance
func (instance *CIM_ResourceAllocationSettingData) SetPropertyLimit(value uint64) (err error) {
	return instance.SetProperty("Limit", value)
}

// GetLimit gets the value of Limit for the instance
func (instance *CIM_ResourceAllocationSettingData) GetPropertyLimit() (value uint64, err error) {
	retValue, err := instance.GetProperty("Limit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMappingBehavior sets the value of MappingBehavior for the instance
func (instance *CIM_ResourceAllocationSettingData) SetPropertyMappingBehavior(value ResourceAllocationSettingData_MappingBehavior) (err error) {
	return instance.SetProperty("MappingBehavior", value)
}

// GetMappingBehavior gets the value of MappingBehavior for the instance
func (instance *CIM_ResourceAllocationSettingData) GetPropertyMappingBehavior() (value ResourceAllocationSettingData_MappingBehavior, err error) {
	retValue, err := instance.GetProperty("MappingBehavior")
	if err != nil {
		return
	}
	value, ok := retValue.(ResourceAllocationSettingData_MappingBehavior)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherResourceType sets the value of OtherResourceType for the instance
func (instance *CIM_ResourceAllocationSettingData) SetPropertyOtherResourceType(value string) (err error) {
	return instance.SetProperty("OtherResourceType", value)
}

// GetOtherResourceType gets the value of OtherResourceType for the instance
func (instance *CIM_ResourceAllocationSettingData) GetPropertyOtherResourceType() (value string, err error) {
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

// SetParent sets the value of Parent for the instance
func (instance *CIM_ResourceAllocationSettingData) SetPropertyParent(value string) (err error) {
	return instance.SetProperty("Parent", value)
}

// GetParent gets the value of Parent for the instance
func (instance *CIM_ResourceAllocationSettingData) GetPropertyParent() (value string, err error) {
	retValue, err := instance.GetProperty("Parent")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPoolID sets the value of PoolID for the instance
func (instance *CIM_ResourceAllocationSettingData) SetPropertyPoolID(value string) (err error) {
	return instance.SetProperty("PoolID", value)
}

// GetPoolID gets the value of PoolID for the instance
func (instance *CIM_ResourceAllocationSettingData) GetPropertyPoolID() (value string, err error) {
	retValue, err := instance.GetProperty("PoolID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReservation sets the value of Reservation for the instance
func (instance *CIM_ResourceAllocationSettingData) SetPropertyReservation(value uint64) (err error) {
	return instance.SetProperty("Reservation", value)
}

// GetReservation gets the value of Reservation for the instance
func (instance *CIM_ResourceAllocationSettingData) GetPropertyReservation() (value uint64, err error) {
	retValue, err := instance.GetProperty("Reservation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceSubType sets the value of ResourceSubType for the instance
func (instance *CIM_ResourceAllocationSettingData) SetPropertyResourceSubType(value string) (err error) {
	return instance.SetProperty("ResourceSubType", value)
}

// GetResourceSubType gets the value of ResourceSubType for the instance
func (instance *CIM_ResourceAllocationSettingData) GetPropertyResourceSubType() (value string, err error) {
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
func (instance *CIM_ResourceAllocationSettingData) SetPropertyResourceType(value ResourceAllocationSettingData_ResourceType) (err error) {
	return instance.SetProperty("ResourceType", value)
}

// GetResourceType gets the value of ResourceType for the instance
func (instance *CIM_ResourceAllocationSettingData) GetPropertyResourceType() (value ResourceAllocationSettingData_ResourceType, err error) {
	retValue, err := instance.GetProperty("ResourceType")
	if err != nil {
		return
	}
	value, ok := retValue.(ResourceAllocationSettingData_ResourceType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualQuantity sets the value of VirtualQuantity for the instance
func (instance *CIM_ResourceAllocationSettingData) SetPropertyVirtualQuantity(value uint64) (err error) {
	return instance.SetProperty("VirtualQuantity", value)
}

// GetVirtualQuantity gets the value of VirtualQuantity for the instance
func (instance *CIM_ResourceAllocationSettingData) GetPropertyVirtualQuantity() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualQuantity")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualQuantityUnits sets the value of VirtualQuantityUnits for the instance
func (instance *CIM_ResourceAllocationSettingData) SetPropertyVirtualQuantityUnits(value string) (err error) {
	return instance.SetProperty("VirtualQuantityUnits", value)
}

// GetVirtualQuantityUnits gets the value of VirtualQuantityUnits for the instance
func (instance *CIM_ResourceAllocationSettingData) GetPropertyVirtualQuantityUnits() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualQuantityUnits")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWeight sets the value of Weight for the instance
func (instance *CIM_ResourceAllocationSettingData) SetPropertyWeight(value uint32) (err error) {
	return instance.SetProperty("Weight", value)
}

// GetWeight gets the value of Weight for the instance
func (instance *CIM_ResourceAllocationSettingData) GetPropertyWeight() (value uint32, err error) {
	retValue, err := instance.GetProperty("Weight")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
