// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_StorageObject struct
type MSFT_StorageObject struct {
	cim.WmiInstance

	//
	ObjectId string

	//
	PassThroughClass string

	//
	PassThroughIds string

	//
	PassThroughNamespace string

	//
	PassThroughServer string

	//
	UniqueId string
}

// SetObjectId sets the value of ObjectId for the instance
func (instance *MSFT_StorageObject) SetPropertyObjectId(value string) (err error) {
	return instance.SetProperty("ObjectId", value)
}

// GetObjectId gets the value of ObjectId for the instance
func (instance *MSFT_StorageObject) GetPropertyObjectId() (value string, err error) {
	retValue, err := instance.GetProperty("ObjectId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPassThroughClass sets the value of PassThroughClass for the instance
func (instance *MSFT_StorageObject) SetPropertyPassThroughClass(value string) (err error) {
	return instance.SetProperty("PassThroughClass", value)
}

// GetPassThroughClass gets the value of PassThroughClass for the instance
func (instance *MSFT_StorageObject) GetPropertyPassThroughClass() (value string, err error) {
	retValue, err := instance.GetProperty("PassThroughClass")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPassThroughIds sets the value of PassThroughIds for the instance
func (instance *MSFT_StorageObject) SetPropertyPassThroughIds(value string) (err error) {
	return instance.SetProperty("PassThroughIds", value)
}

// GetPassThroughIds gets the value of PassThroughIds for the instance
func (instance *MSFT_StorageObject) GetPropertyPassThroughIds() (value string, err error) {
	retValue, err := instance.GetProperty("PassThroughIds")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPassThroughNamespace sets the value of PassThroughNamespace for the instance
func (instance *MSFT_StorageObject) SetPropertyPassThroughNamespace(value string) (err error) {
	return instance.SetProperty("PassThroughNamespace", value)
}

// GetPassThroughNamespace gets the value of PassThroughNamespace for the instance
func (instance *MSFT_StorageObject) GetPropertyPassThroughNamespace() (value string, err error) {
	retValue, err := instance.GetProperty("PassThroughNamespace")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPassThroughServer sets the value of PassThroughServer for the instance
func (instance *MSFT_StorageObject) SetPropertyPassThroughServer(value string) (err error) {
	return instance.SetProperty("PassThroughServer", value)
}

// GetPassThroughServer gets the value of PassThroughServer for the instance
func (instance *MSFT_StorageObject) GetPropertyPassThroughServer() (value string, err error) {
	retValue, err := instance.GetProperty("PassThroughServer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUniqueId sets the value of UniqueId for the instance
func (instance *MSFT_StorageObject) SetPropertyUniqueId(value string) (err error) {
	return instance.SetProperty("UniqueId", value)
}

// GetUniqueId gets the value of UniqueId for the instance
func (instance *MSFT_StorageObject) GetPropertyUniqueId() (value string, err error) {
	retValue, err := instance.GetProperty("UniqueId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
