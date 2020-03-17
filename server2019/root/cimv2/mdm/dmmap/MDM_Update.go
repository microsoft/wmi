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

// MDM_Update struct
type MDM_Update struct {
	cim.WmiInstance

	//
	DeferUpgrade int32

	//
	InstanceID string

	//
	LastSuccessfulScanTime string

	//
	ParentID string
}

// SetDeferUpgrade sets the value of DeferUpgrade for the instance
func (instance *MDM_Update) SetPropertyDeferUpgrade(value int32) (err error) {
	return instance.SetProperty("DeferUpgrade", value)
}

// GetDeferUpgrade gets the value of DeferUpgrade for the instance
func (instance *MDM_Update) GetPropertyDeferUpgrade() (value int32, err error) {
	retValue, err := instance.GetProperty("DeferUpgrade")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Update) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Update) GetPropertyInstanceID() (value string, err error) {
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

// SetLastSuccessfulScanTime sets the value of LastSuccessfulScanTime for the instance
func (instance *MDM_Update) SetPropertyLastSuccessfulScanTime(value string) (err error) {
	return instance.SetProperty("LastSuccessfulScanTime", value)
}

// GetLastSuccessfulScanTime gets the value of LastSuccessfulScanTime for the instance
func (instance *MDM_Update) GetPropertyLastSuccessfulScanTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastSuccessfulScanTime")
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
func (instance *MDM_Update) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Update) GetPropertyParentID() (value string, err error) {
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
