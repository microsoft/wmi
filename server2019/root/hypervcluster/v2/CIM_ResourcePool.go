// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// CIM_ResourcePool struct
type CIM_ResourcePool struct {
	CIM_LogicalElement

	// This property specifies the units of allocation used by the Reservation and Limit properties. For example, when ResourceType=Processor, AllocationUnits may be set to hertz*10^6 or percent. When ResourceType=Memory, AllocationUnits may be set to bytes*10^3. The value of this property shall be a legal value of the Programmatic Units qualifier as defined in Appendix C.1 of DSP0004 V2.4 or later.
	AllocationUnits string

	// This property represents the maximum amount (in units of AllocationUnits) of reservations that the ResourcePool can support.
	Capacity uint64

	// This property specifies the units for the MaxConsumable and the Consumed properties.
	ConsumedResourceUnits string

	// This property specifies the amount of resource that the resource pool currently presents to consumers.
	///This property is different from the Reserved property in that it describes the consumers view of the resource while the Reserved property describes the producers view of the resource.
	CurrentlyConsumedResource uint64

	// This property specifies the maximum of amount of consumable resource that the resource pool can present to consumers.
	///This property is different from the Capacity property in that it describes the consumers view of the resource while the Capacity property describes the producers view of the resource.
	MaxConsumableResource uint64

	// A string that describes the resource type when a well defined value is not available and ResourceType is set to 0 for Other.
	OtherResourceType string

	// An opaque identifier for the pool. This property is used to provide correlation across save and restore of configuration data to underlying persistent storage.
	PoolID string

	// If true, "Primordial" indicates that this ResourcePool is a base from which resources are drawn and returned in the activity of resource management. Being primordial means that this ResourcePool shall not be created or deleted by consumers of this model. However, other actions, modeled or not, may affect the characteristics or size of primordial ResourcePools. If false, "Primordial" indicates that the ResourcePool, a concrete Resource Pool, is subject to resource allocation services functions. This distinction is important because higher-level ResourcePools may be assembled using the Component or ElementAllocatedFromPool associations. Although the higher-level abstractions can be created and deleted, the most basic, (i.e. primordial), hardware-based ResourcePools cannot. They are physically realized as part of the System, or are actually managed by some other System and imported as if they were physically realized.
	Primordial bool

	// This property represents the current reservations (in units of AllocationUnits) spread across all active allocations from this pool. In a hierarchical configuration, this represents the sum of all descendant ResourcePool current reservations.
	Reserved uint64

	// A string describing an implementation specific sub-type for this pool. For example, this may be used to distinguish different models of the same resource type.
	ResourceSubType string

	// The type of resource this ResourcePool may allocate.
	ResourceType ResourcePool_ResourceType
}

// SetAllocationUnits sets the value of AllocationUnits for the instance
func (instance *CIM_ResourcePool) SetPropertyAllocationUnits(value string) (err error) {
	return instance.SetProperty("AllocationUnits", value)
}

// GetAllocationUnits gets the value of AllocationUnits for the instance
func (instance *CIM_ResourcePool) GetPropertyAllocationUnits() (value string, err error) {
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

// SetCapacity sets the value of Capacity for the instance
func (instance *CIM_ResourcePool) SetPropertyCapacity(value uint64) (err error) {
	return instance.SetProperty("Capacity", value)
}

// GetCapacity gets the value of Capacity for the instance
func (instance *CIM_ResourcePool) GetPropertyCapacity() (value uint64, err error) {
	retValue, err := instance.GetProperty("Capacity")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConsumedResourceUnits sets the value of ConsumedResourceUnits for the instance
func (instance *CIM_ResourcePool) SetPropertyConsumedResourceUnits(value string) (err error) {
	return instance.SetProperty("ConsumedResourceUnits", value)
}

// GetConsumedResourceUnits gets the value of ConsumedResourceUnits for the instance
func (instance *CIM_ResourcePool) GetPropertyConsumedResourceUnits() (value string, err error) {
	retValue, err := instance.GetProperty("ConsumedResourceUnits")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentlyConsumedResource sets the value of CurrentlyConsumedResource for the instance
func (instance *CIM_ResourcePool) SetPropertyCurrentlyConsumedResource(value uint64) (err error) {
	return instance.SetProperty("CurrentlyConsumedResource", value)
}

// GetCurrentlyConsumedResource gets the value of CurrentlyConsumedResource for the instance
func (instance *CIM_ResourcePool) GetPropertyCurrentlyConsumedResource() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentlyConsumedResource")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxConsumableResource sets the value of MaxConsumableResource for the instance
func (instance *CIM_ResourcePool) SetPropertyMaxConsumableResource(value uint64) (err error) {
	return instance.SetProperty("MaxConsumableResource", value)
}

// GetMaxConsumableResource gets the value of MaxConsumableResource for the instance
func (instance *CIM_ResourcePool) GetPropertyMaxConsumableResource() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxConsumableResource")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherResourceType sets the value of OtherResourceType for the instance
func (instance *CIM_ResourcePool) SetPropertyOtherResourceType(value string) (err error) {
	return instance.SetProperty("OtherResourceType", value)
}

// GetOtherResourceType gets the value of OtherResourceType for the instance
func (instance *CIM_ResourcePool) GetPropertyOtherResourceType() (value string, err error) {
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

// SetPoolID sets the value of PoolID for the instance
func (instance *CIM_ResourcePool) SetPropertyPoolID(value string) (err error) {
	return instance.SetProperty("PoolID", value)
}

// GetPoolID gets the value of PoolID for the instance
func (instance *CIM_ResourcePool) GetPropertyPoolID() (value string, err error) {
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

// SetPrimordial sets the value of Primordial for the instance
func (instance *CIM_ResourcePool) SetPropertyPrimordial(value bool) (err error) {
	return instance.SetProperty("Primordial", value)
}

// GetPrimordial gets the value of Primordial for the instance
func (instance *CIM_ResourcePool) GetPropertyPrimordial() (value bool, err error) {
	retValue, err := instance.GetProperty("Primordial")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReserved sets the value of Reserved for the instance
func (instance *CIM_ResourcePool) SetPropertyReserved(value uint64) (err error) {
	return instance.SetProperty("Reserved", value)
}

// GetReserved gets the value of Reserved for the instance
func (instance *CIM_ResourcePool) GetPropertyReserved() (value uint64, err error) {
	retValue, err := instance.GetProperty("Reserved")
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
func (instance *CIM_ResourcePool) SetPropertyResourceSubType(value string) (err error) {
	return instance.SetProperty("ResourceSubType", value)
}

// GetResourceSubType gets the value of ResourceSubType for the instance
func (instance *CIM_ResourcePool) GetPropertyResourceSubType() (value string, err error) {
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
func (instance *CIM_ResourcePool) SetPropertyResourceType(value ResourcePool_ResourceType) (err error) {
	return instance.SetProperty("ResourceType", value)
}

// GetResourceType gets the value of ResourceType for the instance
func (instance *CIM_ResourcePool) GetPropertyResourceType() (value ResourcePool_ResourceType, err error) {
	retValue, err := instance.GetProperty("ResourceType")
	if err != nil {
		return
	}
	value, ok := retValue.(ResourcePool_ResourceType)
	if !ok {
		// TODO: Set an error
	}
	return
}
