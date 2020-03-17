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

// MDM_VPNv2_Certificate04 struct
type MDM_VPNv2_Certificate04 struct {
	cim.WmiInstance

	//
	Eku string

	//
	InstanceID string

	//
	Issuer string

	//
	ParentID string
}

// SetEku sets the value of Eku for the instance
func (instance *MDM_VPNv2_Certificate04) SetPropertyEku(value string) (err error) {
	return instance.SetProperty("Eku", value)
}

// GetEku gets the value of Eku for the instance
func (instance *MDM_VPNv2_Certificate04) GetPropertyEku() (value string, err error) {
	retValue, err := instance.GetProperty("Eku")
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
func (instance *MDM_VPNv2_Certificate04) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_VPNv2_Certificate04) GetPropertyInstanceID() (value string, err error) {
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

// SetIssuer sets the value of Issuer for the instance
func (instance *MDM_VPNv2_Certificate04) SetPropertyIssuer(value string) (err error) {
	return instance.SetProperty("Issuer", value)
}

// GetIssuer gets the value of Issuer for the instance
func (instance *MDM_VPNv2_Certificate04) GetPropertyIssuer() (value string, err error) {
	retValue, err := instance.GetProperty("Issuer")
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
func (instance *MDM_VPNv2_Certificate04) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_VPNv2_Certificate04) GetPropertyParentID() (value string, err error) {
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
