// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_SystemSpecificCollection struct
type CIM_SystemSpecificCollection struct {
	*CIM_Collection

	//
	InstanceID string
}

func NewCIM_SystemSpecificCollectionEx1(instance *cim.WmiInstance) (newInstance *CIM_SystemSpecificCollection, err error) {
	tmp, err := NewCIM_CollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_SystemSpecificCollection{
		CIM_Collection: tmp,
	}
	return
}

func NewCIM_SystemSpecificCollectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_SystemSpecificCollection, err error) {
	tmp, err := NewCIM_CollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_SystemSpecificCollection{
		CIM_Collection: tmp,
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *CIM_SystemSpecificCollection) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *CIM_SystemSpecificCollection) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
