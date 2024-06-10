// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.ServiceModel
//
// ////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// PeerCustomResolverBindingElement struct
type PeerCustomResolverBindingElement struct {
	*PeerResolverBindingElement

	// The address of the peer custom resolver.
	Address string

	// The configuration name of the binding.
	Binding string
}

func NewPeerCustomResolverBindingElementEx1(instance *cim.WmiInstance) (newInstance *PeerCustomResolverBindingElement, err error) {
	tmp, err := NewPeerResolverBindingElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &PeerCustomResolverBindingElement{
		PeerResolverBindingElement: tmp,
	}
	return
}

func NewPeerCustomResolverBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PeerCustomResolverBindingElement, err error) {
	tmp, err := NewPeerResolverBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PeerCustomResolverBindingElement{
		PeerResolverBindingElement: tmp,
	}
	return
}

// SetAddress sets the value of Address for the instance
func (instance *PeerCustomResolverBindingElement) SetPropertyAddress(value string) (err error) {
	return instance.SetProperty("Address", (value))
}

// GetAddress gets the value of Address for the instance
func (instance *PeerCustomResolverBindingElement) GetPropertyAddress() (value string, err error) {
	retValue, err := instance.GetProperty("Address")
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

// SetBinding sets the value of Binding for the instance
func (instance *PeerCustomResolverBindingElement) SetPropertyBinding(value string) (err error) {
	return instance.SetProperty("Binding", (value))
}

// GetBinding gets the value of Binding for the instance
func (instance *PeerCustomResolverBindingElement) GetPropertyBinding() (value string, err error) {
	retValue, err := instance.GetProperty("Binding")
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
