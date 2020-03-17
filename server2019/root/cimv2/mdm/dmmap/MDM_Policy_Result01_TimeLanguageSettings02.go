// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Policy_Result01_TimeLanguageSettings02 struct
type MDM_Policy_Result01_TimeLanguageSettings02 struct {
	cim.WmiInstance

	//
	ConfigureTimeZone string

	//
	InstanceID string

	//
	ParentID string
}

// SetConfigureTimeZone sets the value of ConfigureTimeZone for the instance
func (instance *MDM_Policy_Result01_TimeLanguageSettings02) SetPropertyConfigureTimeZone(value string) (err error) {
	return instance.SetProperty("ConfigureTimeZone", value)
}

// GetConfigureTimeZone gets the value of ConfigureTimeZone for the instance
func (instance *MDM_Policy_Result01_TimeLanguageSettings02) GetPropertyConfigureTimeZone() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigureTimeZone")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_TimeLanguageSettings02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_TimeLanguageSettings02) GetPropertyInstanceID() (value string, err error) {
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_TimeLanguageSettings02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_TimeLanguageSettings02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
