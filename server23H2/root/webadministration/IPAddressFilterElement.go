// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// IPAddressFilterElement struct
type IPAddressFilterElement struct {
	*CollectionElement

	//
	Allowed bool

	//
	DomainName string

	//
	IpAddress string

	//
	SubnetMask string
}

func NewIPAddressFilterElementEx1(instance *cim.WmiInstance) (newInstance *IPAddressFilterElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &IPAddressFilterElement{
		CollectionElement: tmp,
	}
	return
}

func NewIPAddressFilterElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *IPAddressFilterElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &IPAddressFilterElement{
		CollectionElement: tmp,
	}
	return
}

// SetAllowed sets the value of Allowed for the instance
func (instance *IPAddressFilterElement) SetPropertyAllowed(value bool) (err error) {
	return instance.SetProperty("Allowed", (value))
}

// GetAllowed gets the value of Allowed for the instance
func (instance *IPAddressFilterElement) GetPropertyAllowed() (value bool, err error) {
	retValue, err := instance.GetProperty("Allowed")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetDomainName sets the value of DomainName for the instance
func (instance *IPAddressFilterElement) SetPropertyDomainName(value string) (err error) {
	return instance.SetProperty("DomainName", (value))
}

// GetDomainName gets the value of DomainName for the instance
func (instance *IPAddressFilterElement) GetPropertyDomainName() (value string, err error) {
	retValue, err := instance.GetProperty("DomainName")
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

// SetIpAddress sets the value of IpAddress for the instance
func (instance *IPAddressFilterElement) SetPropertyIpAddress(value string) (err error) {
	return instance.SetProperty("IpAddress", (value))
}

// GetIpAddress gets the value of IpAddress for the instance
func (instance *IPAddressFilterElement) GetPropertyIpAddress() (value string, err error) {
	retValue, err := instance.GetProperty("IpAddress")
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

// SetSubnetMask sets the value of SubnetMask for the instance
func (instance *IPAddressFilterElement) SetPropertySubnetMask(value string) (err error) {
	return instance.SetProperty("SubnetMask", (value))
}

// GetSubnetMask gets the value of SubnetMask for the instance
func (instance *IPAddressFilterElement) GetPropertySubnetMask() (value string, err error) {
	retValue, err := instance.GetProperty("SubnetMask")
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
