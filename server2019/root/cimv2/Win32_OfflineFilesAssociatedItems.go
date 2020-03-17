// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_OfflineFilesAssociatedItems struct
type Win32_OfflineFilesAssociatedItems struct {
	cim.WmiInstance

	//
	Antecedent Win32_OfflineFilesCache

	//
	Dependent Win32_OfflineFilesItem
}

// SetAntecedent sets the value of Antecedent for the instance
func (instance *Win32_OfflineFilesAssociatedItems) SetPropertyAntecedent(value Win32_OfflineFilesCache) (err error) {
	return instance.SetProperty("Antecedent", value)
}

// GetAntecedent gets the value of Antecedent for the instance
func (instance *Win32_OfflineFilesAssociatedItems) GetPropertyAntecedent() (value Win32_OfflineFilesCache, err error) {
	retValue, err := instance.GetProperty("Antecedent")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_OfflineFilesCache)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDependent sets the value of Dependent for the instance
func (instance *Win32_OfflineFilesAssociatedItems) SetPropertyDependent(value Win32_OfflineFilesItem) (err error) {
	return instance.SetProperty("Dependent", value)
}

// GetDependent gets the value of Dependent for the instance
func (instance *Win32_OfflineFilesAssociatedItems) GetPropertyDependent() (value Win32_OfflineFilesItem, err error) {
	retValue, err := instance.GetProperty("Dependent")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_OfflineFilesItem)
	if !ok {
		// TODO: Set an error
	}
	return
}
