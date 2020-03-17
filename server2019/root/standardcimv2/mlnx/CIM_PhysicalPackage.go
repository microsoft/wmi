// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_PhysicalPackage struct
type CIM_PhysicalPackage struct {
	CIM_PhysicalElement

	//
	Depth float32

	//
	Height float32

	//
	HotSwappable bool

	//
	OtherPackageType string

	//
	PackageType PhysicalPackage_PackageType

	//
	Removable bool

	//
	RemovalConditions PhysicalPackage_RemovalConditions

	//
	Replaceable bool

	//
	VendorCompatibilityStrings []string

	//
	Weight float32

	//
	Width float32
}

// SetDepth sets the value of Depth for the instance
func (instance *CIM_PhysicalPackage) SetPropertyDepth(value float32) (err error) {
	return instance.SetProperty("Depth", value)
}

// GetDepth gets the value of Depth for the instance
func (instance *CIM_PhysicalPackage) GetPropertyDepth() (value float32, err error) {
	retValue, err := instance.GetProperty("Depth")
	if err != nil {
		return
	}
	value, ok := retValue.(float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHeight sets the value of Height for the instance
func (instance *CIM_PhysicalPackage) SetPropertyHeight(value float32) (err error) {
	return instance.SetProperty("Height", value)
}

// GetHeight gets the value of Height for the instance
func (instance *CIM_PhysicalPackage) GetPropertyHeight() (value float32, err error) {
	retValue, err := instance.GetProperty("Height")
	if err != nil {
		return
	}
	value, ok := retValue.(float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHotSwappable sets the value of HotSwappable for the instance
func (instance *CIM_PhysicalPackage) SetPropertyHotSwappable(value bool) (err error) {
	return instance.SetProperty("HotSwappable", value)
}

// GetHotSwappable gets the value of HotSwappable for the instance
func (instance *CIM_PhysicalPackage) GetPropertyHotSwappable() (value bool, err error) {
	retValue, err := instance.GetProperty("HotSwappable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherPackageType sets the value of OtherPackageType for the instance
func (instance *CIM_PhysicalPackage) SetPropertyOtherPackageType(value string) (err error) {
	return instance.SetProperty("OtherPackageType", value)
}

// GetOtherPackageType gets the value of OtherPackageType for the instance
func (instance *CIM_PhysicalPackage) GetPropertyOtherPackageType() (value string, err error) {
	retValue, err := instance.GetProperty("OtherPackageType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPackageType sets the value of PackageType for the instance
func (instance *CIM_PhysicalPackage) SetPropertyPackageType(value PhysicalPackage_PackageType) (err error) {
	return instance.SetProperty("PackageType", value)
}

// GetPackageType gets the value of PackageType for the instance
func (instance *CIM_PhysicalPackage) GetPropertyPackageType() (value PhysicalPackage_PackageType, err error) {
	retValue, err := instance.GetProperty("PackageType")
	if err != nil {
		return
	}
	value, ok := retValue.(PhysicalPackage_PackageType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemovable sets the value of Removable for the instance
func (instance *CIM_PhysicalPackage) SetPropertyRemovable(value bool) (err error) {
	return instance.SetProperty("Removable", value)
}

// GetRemovable gets the value of Removable for the instance
func (instance *CIM_PhysicalPackage) GetPropertyRemovable() (value bool, err error) {
	retValue, err := instance.GetProperty("Removable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemovalConditions sets the value of RemovalConditions for the instance
func (instance *CIM_PhysicalPackage) SetPropertyRemovalConditions(value PhysicalPackage_RemovalConditions) (err error) {
	return instance.SetProperty("RemovalConditions", value)
}

// GetRemovalConditions gets the value of RemovalConditions for the instance
func (instance *CIM_PhysicalPackage) GetPropertyRemovalConditions() (value PhysicalPackage_RemovalConditions, err error) {
	retValue, err := instance.GetProperty("RemovalConditions")
	if err != nil {
		return
	}
	value, ok := retValue.(PhysicalPackage_RemovalConditions)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplaceable sets the value of Replaceable for the instance
func (instance *CIM_PhysicalPackage) SetPropertyReplaceable(value bool) (err error) {
	return instance.SetProperty("Replaceable", value)
}

// GetReplaceable gets the value of Replaceable for the instance
func (instance *CIM_PhysicalPackage) GetPropertyReplaceable() (value bool, err error) {
	retValue, err := instance.GetProperty("Replaceable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVendorCompatibilityStrings sets the value of VendorCompatibilityStrings for the instance
func (instance *CIM_PhysicalPackage) SetPropertyVendorCompatibilityStrings(value []string) (err error) {
	return instance.SetProperty("VendorCompatibilityStrings", value)
}

// GetVendorCompatibilityStrings gets the value of VendorCompatibilityStrings for the instance
func (instance *CIM_PhysicalPackage) GetPropertyVendorCompatibilityStrings() (value []string, err error) {
	retValue, err := instance.GetProperty("VendorCompatibilityStrings")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWeight sets the value of Weight for the instance
func (instance *CIM_PhysicalPackage) SetPropertyWeight(value float32) (err error) {
	return instance.SetProperty("Weight", value)
}

// GetWeight gets the value of Weight for the instance
func (instance *CIM_PhysicalPackage) GetPropertyWeight() (value float32, err error) {
	retValue, err := instance.GetProperty("Weight")
	if err != nil {
		return
	}
	value, ok := retValue.(float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWidth sets the value of Width for the instance
func (instance *CIM_PhysicalPackage) SetPropertyWidth(value float32) (err error) {
	return instance.SetProperty("Width", value)
}

// GetWidth gets the value of Width for the instance
func (instance *CIM_PhysicalPackage) GetPropertyWidth() (value float32, err error) {
	retValue, err := instance.GetProperty("Width")
	if err != nil {
		return
	}
	value, ok := retValue.(float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ElementToCheck" type="CIM_PhysicalElement "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_PhysicalPackage) IsCompatible( /* IN */ ElementToCheck CIM_PhysicalElement) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("IsCompatible", ElementToCheck)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
