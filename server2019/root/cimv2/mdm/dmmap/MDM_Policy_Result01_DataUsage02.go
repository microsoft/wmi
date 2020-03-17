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

// MDM_Policy_Result01_DataUsage02 struct
type MDM_Policy_Result01_DataUsage02 struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	ParentID string

	//
	SetCost3G string

	//
	SetCost4G string
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_DataUsage02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_DataUsage02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Result01_DataUsage02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_DataUsage02) GetPropertyParentID() (value string, err error) {
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

// SetSetCost3G sets the value of SetCost3G for the instance
func (instance *MDM_Policy_Result01_DataUsage02) SetPropertySetCost3G(value string) (err error) {
	return instance.SetProperty("SetCost3G", value)
}

// GetSetCost3G gets the value of SetCost3G for the instance
func (instance *MDM_Policy_Result01_DataUsage02) GetPropertySetCost3G() (value string, err error) {
	retValue, err := instance.GetProperty("SetCost3G")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSetCost4G sets the value of SetCost4G for the instance
func (instance *MDM_Policy_Result01_DataUsage02) SetPropertySetCost4G(value string) (err error) {
	return instance.SetProperty("SetCost4G", value)
}

// GetSetCost4G gets the value of SetCost4G for the instance
func (instance *MDM_Policy_Result01_DataUsage02) GetPropertySetCost4G() (value string, err error) {
	retValue, err := instance.GetProperty("SetCost4G")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
