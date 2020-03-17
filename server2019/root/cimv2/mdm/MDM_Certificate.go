// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Certificate struct
type MDM_Certificate struct {
	cim.WmiInstance

	//
	Blob string

	//
	IsInstalled bool

	//
	StoreLocation uint8

	//
	StoreName string

	//
	Thumbprint string
}

// SetBlob sets the value of Blob for the instance
func (instance *MDM_Certificate) SetPropertyBlob(value string) (err error) {
	return instance.SetProperty("Blob", value)
}

// GetBlob gets the value of Blob for the instance
func (instance *MDM_Certificate) GetPropertyBlob() (value string, err error) {
	retValue, err := instance.GetProperty("Blob")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsInstalled sets the value of IsInstalled for the instance
func (instance *MDM_Certificate) SetPropertyIsInstalled(value bool) (err error) {
	return instance.SetProperty("IsInstalled", value)
}

// GetIsInstalled gets the value of IsInstalled for the instance
func (instance *MDM_Certificate) GetPropertyIsInstalled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsInstalled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStoreLocation sets the value of StoreLocation for the instance
func (instance *MDM_Certificate) SetPropertyStoreLocation(value uint8) (err error) {
	return instance.SetProperty("StoreLocation", value)
}

// GetStoreLocation gets the value of StoreLocation for the instance
func (instance *MDM_Certificate) GetPropertyStoreLocation() (value uint8, err error) {
	retValue, err := instance.GetProperty("StoreLocation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStoreName sets the value of StoreName for the instance
func (instance *MDM_Certificate) SetPropertyStoreName(value string) (err error) {
	return instance.SetProperty("StoreName", value)
}

// GetStoreName gets the value of StoreName for the instance
func (instance *MDM_Certificate) GetPropertyStoreName() (value string, err error) {
	retValue, err := instance.GetProperty("StoreName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThumbprint sets the value of Thumbprint for the instance
func (instance *MDM_Certificate) SetPropertyThumbprint(value string) (err error) {
	return instance.SetProperty("Thumbprint", value)
}

// GetThumbprint gets the value of Thumbprint for the instance
func (instance *MDM_Certificate) GetPropertyThumbprint() (value string, err error) {
	retValue, err := instance.GetProperty("Thumbprint")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
