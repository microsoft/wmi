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

// MDM_Policy_User_Result01_AttachmentManager02 struct
type MDM_Policy_User_Result01_AttachmentManager02 struct {
	cim.WmiInstance

	//
	DoNotPreserveZoneInformation string

	//
	HideZoneInfoMechanism string

	//
	InstanceID string

	//
	NotifyAntivirusPrograms string

	//
	ParentID string
}

// SetDoNotPreserveZoneInformation sets the value of DoNotPreserveZoneInformation for the instance
func (instance *MDM_Policy_User_Result01_AttachmentManager02) SetPropertyDoNotPreserveZoneInformation(value string) (err error) {
	return instance.SetProperty("DoNotPreserveZoneInformation", value)
}

// GetDoNotPreserveZoneInformation gets the value of DoNotPreserveZoneInformation for the instance
func (instance *MDM_Policy_User_Result01_AttachmentManager02) GetPropertyDoNotPreserveZoneInformation() (value string, err error) {
	retValue, err := instance.GetProperty("DoNotPreserveZoneInformation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHideZoneInfoMechanism sets the value of HideZoneInfoMechanism for the instance
func (instance *MDM_Policy_User_Result01_AttachmentManager02) SetPropertyHideZoneInfoMechanism(value string) (err error) {
	return instance.SetProperty("HideZoneInfoMechanism", value)
}

// GetHideZoneInfoMechanism gets the value of HideZoneInfoMechanism for the instance
func (instance *MDM_Policy_User_Result01_AttachmentManager02) GetPropertyHideZoneInfoMechanism() (value string, err error) {
	retValue, err := instance.GetProperty("HideZoneInfoMechanism")
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
func (instance *MDM_Policy_User_Result01_AttachmentManager02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Result01_AttachmentManager02) GetPropertyInstanceID() (value string, err error) {
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

// SetNotifyAntivirusPrograms sets the value of NotifyAntivirusPrograms for the instance
func (instance *MDM_Policy_User_Result01_AttachmentManager02) SetPropertyNotifyAntivirusPrograms(value string) (err error) {
	return instance.SetProperty("NotifyAntivirusPrograms", value)
}

// GetNotifyAntivirusPrograms gets the value of NotifyAntivirusPrograms for the instance
func (instance *MDM_Policy_User_Result01_AttachmentManager02) GetPropertyNotifyAntivirusPrograms() (value string, err error) {
	retValue, err := instance.GetProperty("NotifyAntivirusPrograms")
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
func (instance *MDM_Policy_User_Result01_AttachmentManager02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_User_Result01_AttachmentManager02) GetPropertyParentID() (value string, err error) {
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
