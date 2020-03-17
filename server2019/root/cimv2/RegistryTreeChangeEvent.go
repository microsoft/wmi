// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// RegistryTreeChangeEvent struct
type RegistryTreeChangeEvent struct {
	RegistryEvent

	//
	Hive string

	//
	RootPath string
}

// SetHive sets the value of Hive for the instance
func (instance *RegistryTreeChangeEvent) SetPropertyHive(value string) (err error) {
	return instance.SetProperty("Hive", value)
}

// GetHive gets the value of Hive for the instance
func (instance *RegistryTreeChangeEvent) GetPropertyHive() (value string, err error) {
	retValue, err := instance.GetProperty("Hive")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRootPath sets the value of RootPath for the instance
func (instance *RegistryTreeChangeEvent) SetPropertyRootPath(value string) (err error) {
	return instance.SetProperty("RootPath", value)
}

// GetRootPath gets the value of RootPath for the instance
func (instance *RegistryTreeChangeEvent) GetPropertyRootPath() (value string, err error) {
	retValue, err := instance.GetProperty("RootPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
