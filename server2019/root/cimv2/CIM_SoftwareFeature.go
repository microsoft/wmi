// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_SoftwareFeature struct
type CIM_SoftwareFeature struct {
	CIM_LogicalElement

	//
	IdentifyingNumber string

	//
	ProductName string

	//
	Vendor string

	//
	Version string
}

// SetIdentifyingNumber sets the value of IdentifyingNumber for the instance
func (instance *CIM_SoftwareFeature) SetPropertyIdentifyingNumber(value string) (err error) {
	return instance.SetProperty("IdentifyingNumber", value)
}

// GetIdentifyingNumber gets the value of IdentifyingNumber for the instance
func (instance *CIM_SoftwareFeature) GetPropertyIdentifyingNumber() (value string, err error) {
	retValue, err := instance.GetProperty("IdentifyingNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProductName sets the value of ProductName for the instance
func (instance *CIM_SoftwareFeature) SetPropertyProductName(value string) (err error) {
	return instance.SetProperty("ProductName", value)
}

// GetProductName gets the value of ProductName for the instance
func (instance *CIM_SoftwareFeature) GetPropertyProductName() (value string, err error) {
	retValue, err := instance.GetProperty("ProductName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVendor sets the value of Vendor for the instance
func (instance *CIM_SoftwareFeature) SetPropertyVendor(value string) (err error) {
	return instance.SetProperty("Vendor", value)
}

// GetVendor gets the value of Vendor for the instance
func (instance *CIM_SoftwareFeature) GetPropertyVendor() (value string, err error) {
	retValue, err := instance.GetProperty("Vendor")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVersion sets the value of Version for the instance
func (instance *CIM_SoftwareFeature) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *CIM_SoftwareFeature) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
