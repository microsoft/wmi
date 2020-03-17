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

// RSOP_IEESC struct
type RSOP_IEESC struct {
	cim.WmiInstance

	//
	EscEnabled bool

	//
	rsopID string

	//
	rsopPrecedence int32
}

// SetEscEnabled sets the value of EscEnabled for the instance
func (instance *RSOP_IEESC) SetPropertyEscEnabled(value bool) (err error) {
	return instance.SetProperty("EscEnabled", value)
}

// GetEscEnabled gets the value of EscEnabled for the instance
func (instance *RSOP_IEESC) GetPropertyEscEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("EscEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetrsopID sets the value of rsopID for the instance
func (instance *RSOP_IEESC) SetPropertyrsopID(value string) (err error) {
	return instance.SetProperty("rsopID", value)
}

// GetrsopID gets the value of rsopID for the instance
func (instance *RSOP_IEESC) GetPropertyrsopID() (value string, err error) {
	retValue, err := instance.GetProperty("rsopID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetrsopPrecedence sets the value of rsopPrecedence for the instance
func (instance *RSOP_IEESC) SetPropertyrsopPrecedence(value int32) (err error) {
	return instance.SetProperty("rsopPrecedence", value)
}

// GetrsopPrecedence gets the value of rsopPrecedence for the instance
func (instance *RSOP_IEESC) GetPropertyrsopPrecedence() (value int32, err error) {
	retValue, err := instance.GetProperty("rsopPrecedence")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
