// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_CollectionOfMSEs struct
type CIM_CollectionOfMSEs struct {
	*CIM_Collection

	// The identification of the Collection object. When subclassed, the CollectionID property can be overridden to be a Key property.
	CollectionID string
}

func NewCIM_CollectionOfMSEsEx1(instance *cim.WmiInstance) (newInstance *CIM_CollectionOfMSEs, err error) {
	tmp, err := NewCIM_CollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_CollectionOfMSEs{
		CIM_Collection: tmp,
	}
	return
}

func NewCIM_CollectionOfMSEsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_CollectionOfMSEs, err error) {
	tmp, err := NewCIM_CollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_CollectionOfMSEs{
		CIM_Collection: tmp,
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
