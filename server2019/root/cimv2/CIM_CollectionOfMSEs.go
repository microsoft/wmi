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

// CIM_CollectionOfMSEs struct
type CIM_CollectionOfMSEs struct {
	cim.WmiInstance

	//
	Caption string

	//
	CollectionID string

	//
	Description string
}

// SetCaption sets the value of Caption for the instance
func (instance *CIM_CollectionOfMSEs) SetPropertyCaption(value string) (err error) {
	return instance.SetProperty("Caption", value)
}

// GetCaption gets the value of Caption for the instance
func (instance *CIM_CollectionOfMSEs) GetPropertyCaption() (value string, err error) {
	retValue, err := instance.GetProperty("Caption")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCollectionID sets the value of CollectionID for the instance
func (instance *CIM_CollectionOfMSEs) SetPropertyCollectionID(value string) (err error) {
	return instance.SetProperty("CollectionID", value)
}

// GetCollectionID gets the value of CollectionID for the instance
func (instance *CIM_CollectionOfMSEs) GetPropertyCollectionID() (value string, err error) {
	retValue, err := instance.GetProperty("CollectionID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *CIM_CollectionOfMSEs) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *CIM_CollectionOfMSEs) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
