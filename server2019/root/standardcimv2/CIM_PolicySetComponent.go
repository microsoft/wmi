// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_PolicySetComponent struct
type CIM_PolicySetComponent struct {
	*CIM_PolicyComponent

	//
	Priority uint16
}

func NewCIM_PolicySetComponentEx1(instance *cim.WmiInstance) (newInstance *CIM_PolicySetComponent, err error) {
	tmp, err := NewCIM_PolicyComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_PolicySetComponent{
		CIM_PolicyComponent: tmp,
	}
	return
}

func NewCIM_PolicySetComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_PolicySetComponent, err error) {
	tmp, err := NewCIM_PolicyComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_PolicySetComponent{
		CIM_PolicyComponent: tmp,
	}
	return
}

// SetPriority sets the value of Priority for the instance
func (instance *CIM_PolicySetComponent) SetPropertyPriority(value uint16) (err error) {
	return instance.SetProperty("Priority", value)
}

// GetPriority gets the value of Priority for the instance
func (instance *CIM_PolicySetComponent) GetPropertyPriority() (value uint16, err error) {
	retValue, err := instance.GetProperty("Priority")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
