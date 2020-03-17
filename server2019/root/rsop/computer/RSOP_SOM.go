// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_SOM struct
type RSOP_SOM struct {
	cim.WmiInstance

	//
	blocked bool

	//
	blocking bool

	//
	id string

	//
	reason uint32

	//
	SOMOrder uint32

	//
	_type uint32
}

// Setblocked sets the value of blocked for the instance
func (instance *RSOP_SOM) SetPropertyblocked(value bool) (err error) {
	return instance.SetProperty("blocked", value)
}

// Getblocked gets the value of blocked for the instance
func (instance *RSOP_SOM) GetPropertyblocked() (value bool, err error) {
	retValue, err := instance.GetProperty("blocked")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setblocking sets the value of blocking for the instance
func (instance *RSOP_SOM) SetPropertyblocking(value bool) (err error) {
	return instance.SetProperty("blocking", value)
}

// Getblocking gets the value of blocking for the instance
func (instance *RSOP_SOM) GetPropertyblocking() (value bool, err error) {
	retValue, err := instance.GetProperty("blocking")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setid sets the value of id for the instance
func (instance *RSOP_SOM) SetPropertyid(value string) (err error) {
	return instance.SetProperty("id", value)
}

// Getid gets the value of id for the instance
func (instance *RSOP_SOM) GetPropertyid() (value string, err error) {
	retValue, err := instance.GetProperty("id")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setreason sets the value of reason for the instance
func (instance *RSOP_SOM) SetPropertyreason(value uint32) (err error) {
	return instance.SetProperty("reason", value)
}

// Getreason gets the value of reason for the instance
func (instance *RSOP_SOM) GetPropertyreason() (value uint32, err error) {
	retValue, err := instance.GetProperty("reason")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSOMOrder sets the value of SOMOrder for the instance
func (instance *RSOP_SOM) SetPropertySOMOrder(value uint32) (err error) {
	return instance.SetProperty("SOMOrder", value)
}

// GetSOMOrder gets the value of SOMOrder for the instance
func (instance *RSOP_SOM) GetPropertySOMOrder() (value uint32, err error) {
	retValue, err := instance.GetProperty("SOMOrder")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Settype sets the value of type for the instance
func (instance *RSOP_SOM) SetPropertytype(value uint32) (err error) {
	return instance.SetProperty("type", value)
}

// Gettype gets the value of type for the instance
func (instance *RSOP_SOM) GetPropertytype() (value uint32, err error) {
	retValue, err := instance.GetProperty("type")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
