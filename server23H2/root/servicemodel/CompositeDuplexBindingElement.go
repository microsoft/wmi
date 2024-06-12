// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CompositeDuplexBindingElement struct
type CompositeDuplexBindingElement struct {
	*BindingElement

	// The base address of the client.
	ClientBaseAddress string
}

func NewCompositeDuplexBindingElementEx1(instance *cim.WmiInstance) (newInstance *CompositeDuplexBindingElement, err error) {
	tmp, err := NewBindingElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CompositeDuplexBindingElement{
		BindingElement: tmp,
	}
	return
}

func NewCompositeDuplexBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CompositeDuplexBindingElement, err error) {
	tmp, err := NewBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CompositeDuplexBindingElement{
		BindingElement: tmp,
	}
	return
}

// SetClientBaseAddress sets the value of ClientBaseAddress for the instance
func (instance *CompositeDuplexBindingElement) SetPropertyClientBaseAddress(value string) (err error) {
	return instance.SetProperty("ClientBaseAddress", (value))
}

// GetClientBaseAddress gets the value of ClientBaseAddress for the instance
func (instance *CompositeDuplexBindingElement) GetPropertyClientBaseAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ClientBaseAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
