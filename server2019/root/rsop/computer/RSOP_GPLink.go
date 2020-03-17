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

// RSOP_GPLink struct
type RSOP_GPLink struct {
	cim.WmiInstance

	//
	appliedOrder uint32

	//
	enabled bool

	//
	GPO RSOP_GPO

	//
	linkOrder uint32

	//
	noOverride bool

	//
	SOM RSOP_SOM

	//
	somOrder uint32
}

// SetappliedOrder sets the value of appliedOrder for the instance
func (instance *RSOP_GPLink) SetPropertyappliedOrder(value uint32) (err error) {
	return instance.SetProperty("appliedOrder", value)
}

// GetappliedOrder gets the value of appliedOrder for the instance
func (instance *RSOP_GPLink) GetPropertyappliedOrder() (value uint32, err error) {
	retValue, err := instance.GetProperty("appliedOrder")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setenabled sets the value of enabled for the instance
func (instance *RSOP_GPLink) SetPropertyenabled(value bool) (err error) {
	return instance.SetProperty("enabled", value)
}

// Getenabled gets the value of enabled for the instance
func (instance *RSOP_GPLink) GetPropertyenabled() (value bool, err error) {
	retValue, err := instance.GetProperty("enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGPO sets the value of GPO for the instance
func (instance *RSOP_GPLink) SetPropertyGPO(value RSOP_GPO) (err error) {
	return instance.SetProperty("GPO", value)
}

// GetGPO gets the value of GPO for the instance
func (instance *RSOP_GPLink) GetPropertyGPO() (value RSOP_GPO, err error) {
	retValue, err := instance.GetProperty("GPO")
	if err != nil {
		return
	}
	value, ok := retValue.(RSOP_GPO)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetlinkOrder sets the value of linkOrder for the instance
func (instance *RSOP_GPLink) SetPropertylinkOrder(value uint32) (err error) {
	return instance.SetProperty("linkOrder", value)
}

// GetlinkOrder gets the value of linkOrder for the instance
func (instance *RSOP_GPLink) GetPropertylinkOrder() (value uint32, err error) {
	retValue, err := instance.GetProperty("linkOrder")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetnoOverride sets the value of noOverride for the instance
func (instance *RSOP_GPLink) SetPropertynoOverride(value bool) (err error) {
	return instance.SetProperty("noOverride", value)
}

// GetnoOverride gets the value of noOverride for the instance
func (instance *RSOP_GPLink) GetPropertynoOverride() (value bool, err error) {
	retValue, err := instance.GetProperty("noOverride")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSOM sets the value of SOM for the instance
func (instance *RSOP_GPLink) SetPropertySOM(value RSOP_SOM) (err error) {
	return instance.SetProperty("SOM", value)
}

// GetSOM gets the value of SOM for the instance
func (instance *RSOP_GPLink) GetPropertySOM() (value RSOP_SOM, err error) {
	retValue, err := instance.GetProperty("SOM")
	if err != nil {
		return
	}
	value, ok := retValue.(RSOP_SOM)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetsomOrder sets the value of somOrder for the instance
func (instance *RSOP_GPLink) SetPropertysomOrder(value uint32) (err error) {
	return instance.SetProperty("somOrder", value)
}

// GetsomOrder gets the value of somOrder for the instance
func (instance *RSOP_GPLink) GetPropertysomOrder() (value uint32, err error) {
	retValue, err := instance.GetProperty("somOrder")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
