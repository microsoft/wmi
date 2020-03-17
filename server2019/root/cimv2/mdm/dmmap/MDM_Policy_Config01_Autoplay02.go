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

// MDM_Policy_Config01_Autoplay02 struct
type MDM_Policy_Config01_Autoplay02 struct {
	cim.WmiInstance

	//
	DisallowAutoplayForNonVolumeDevices string

	//
	InstanceID string

	//
	ParentID string

	//
	SetDefaultAutoRunBehavior string

	//
	TurnOffAutoPlay string
}

// SetDisallowAutoplayForNonVolumeDevices sets the value of DisallowAutoplayForNonVolumeDevices for the instance
func (instance *MDM_Policy_Config01_Autoplay02) SetPropertyDisallowAutoplayForNonVolumeDevices(value string) (err error) {
	return instance.SetProperty("DisallowAutoplayForNonVolumeDevices", value)
}

// GetDisallowAutoplayForNonVolumeDevices gets the value of DisallowAutoplayForNonVolumeDevices for the instance
func (instance *MDM_Policy_Config01_Autoplay02) GetPropertyDisallowAutoplayForNonVolumeDevices() (value string, err error) {
	retValue, err := instance.GetProperty("DisallowAutoplayForNonVolumeDevices")
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
func (instance *MDM_Policy_Config01_Autoplay02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Autoplay02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_Autoplay02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Autoplay02) GetPropertyParentID() (value string, err error) {
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

// SetSetDefaultAutoRunBehavior sets the value of SetDefaultAutoRunBehavior for the instance
func (instance *MDM_Policy_Config01_Autoplay02) SetPropertySetDefaultAutoRunBehavior(value string) (err error) {
	return instance.SetProperty("SetDefaultAutoRunBehavior", value)
}

// GetSetDefaultAutoRunBehavior gets the value of SetDefaultAutoRunBehavior for the instance
func (instance *MDM_Policy_Config01_Autoplay02) GetPropertySetDefaultAutoRunBehavior() (value string, err error) {
	retValue, err := instance.GetProperty("SetDefaultAutoRunBehavior")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTurnOffAutoPlay sets the value of TurnOffAutoPlay for the instance
func (instance *MDM_Policy_Config01_Autoplay02) SetPropertyTurnOffAutoPlay(value string) (err error) {
	return instance.SetProperty("TurnOffAutoPlay", value)
}

// GetTurnOffAutoPlay gets the value of TurnOffAutoPlay for the instance
func (instance *MDM_Policy_Config01_Autoplay02) GetPropertyTurnOffAutoPlay() (value string, err error) {
	retValue, err := instance.GetProperty("TurnOffAutoPlay")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
