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

// Msvm_WiFiPort struct
type Msvm_WiFiPort struct {
	*CIM_WiFiPort

	//
	IsBound bool
}

func NewMsvm_WiFiPortEx1(instance *cim.WmiInstance) (newInstance *Msvm_WiFiPort, err error) {
	tmp, err := NewCIM_WiFiPortEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_WiFiPort{
		CIM_WiFiPort: tmp,
	}
	return
}

func NewMsvm_WiFiPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_WiFiPort, err error) {
	tmp, err := NewCIM_WiFiPortEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_WiFiPort{
		CIM_WiFiPort: tmp,
	}
	return
}

// SetIsBound sets the value of IsBound for the instance
func (instance *Msvm_WiFiPort) SetPropertyIsBound(value bool) (err error) {
	return instance.SetProperty("IsBound", value)
}

// GetIsBound gets the value of IsBound for the instance
func (instance *Msvm_WiFiPort) GetPropertyIsBound() (value bool, err error) {
	retValue, err := instance.GetProperty("IsBound")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
