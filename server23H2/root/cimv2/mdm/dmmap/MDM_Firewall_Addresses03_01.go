// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Firewall_Addresses03_01 struct
type MDM_Firewall_Addresses03_01 struct {
	*cim.WmiInstance

	//
	Addresses string

	//
	AutoResolve bool

	//
	InstanceID string

	//
	Keyword string

	//
	ParentID string
}

func NewMDM_Firewall_Addresses03_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_Firewall_Addresses03_01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Firewall_Addresses03_01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Firewall_Addresses03_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Firewall_Addresses03_01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Firewall_Addresses03_01{
		WmiInstance: tmp,
	}
	return
}

// SetAddresses sets the value of Addresses for the instance
func (instance *MDM_Firewall_Addresses03_01) SetPropertyAddresses(value string) (err error) {
	return instance.SetProperty("Addresses", (value))
}

// GetAddresses gets the value of Addresses for the instance
func (instance *MDM_Firewall_Addresses03_01) GetPropertyAddresses() (value string, err error) {
	retValue, err := instance.GetProperty("Addresses")
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

// SetAutoResolve sets the value of AutoResolve for the instance
func (instance *MDM_Firewall_Addresses03_01) SetPropertyAutoResolve(value bool) (err error) {
	return instance.SetProperty("AutoResolve", (value))
}

// GetAutoResolve gets the value of AutoResolve for the instance
func (instance *MDM_Firewall_Addresses03_01) GetPropertyAutoResolve() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoResolve")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Firewall_Addresses03_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Firewall_Addresses03_01) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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

// SetKeyword sets the value of Keyword for the instance
func (instance *MDM_Firewall_Addresses03_01) SetPropertyKeyword(value string) (err error) {
	return instance.SetProperty("Keyword", (value))
}

// GetKeyword gets the value of Keyword for the instance
func (instance *MDM_Firewall_Addresses03_01) GetPropertyKeyword() (value string, err error) {
	retValue, err := instance.GetProperty("Keyword")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Firewall_Addresses03_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Firewall_Addresses03_01) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
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
