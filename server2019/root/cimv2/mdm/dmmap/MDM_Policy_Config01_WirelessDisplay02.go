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

// MDM_Policy_Config01_WirelessDisplay02 struct
type MDM_Policy_Config01_WirelessDisplay02 struct {
	cim.WmiInstance

	//
	AllowMdnsAdvertisement int32

	//
	AllowMdnsDiscovery int32

	//
	AllowProjectionFromPC int32

	//
	AllowProjectionFromPCOverInfrastructure int32

	//
	AllowProjectionToPC int32

	//
	AllowProjectionToPCOverInfrastructure int32

	//
	AllowUserInputFromWirelessDisplayReceiver int32

	//
	InstanceID string

	//
	ParentID string

	//
	RequirePinForPairing int32
}

// SetAllowMdnsAdvertisement sets the value of AllowMdnsAdvertisement for the instance
func (instance *MDM_Policy_Config01_WirelessDisplay02) SetPropertyAllowMdnsAdvertisement(value int32) (err error) {
	return instance.SetProperty("AllowMdnsAdvertisement", value)
}

// GetAllowMdnsAdvertisement gets the value of AllowMdnsAdvertisement for the instance
func (instance *MDM_Policy_Config01_WirelessDisplay02) GetPropertyAllowMdnsAdvertisement() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowMdnsAdvertisement")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowMdnsDiscovery sets the value of AllowMdnsDiscovery for the instance
func (instance *MDM_Policy_Config01_WirelessDisplay02) SetPropertyAllowMdnsDiscovery(value int32) (err error) {
	return instance.SetProperty("AllowMdnsDiscovery", value)
}

// GetAllowMdnsDiscovery gets the value of AllowMdnsDiscovery for the instance
func (instance *MDM_Policy_Config01_WirelessDisplay02) GetPropertyAllowMdnsDiscovery() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowMdnsDiscovery")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowProjectionFromPC sets the value of AllowProjectionFromPC for the instance
func (instance *MDM_Policy_Config01_WirelessDisplay02) SetPropertyAllowProjectionFromPC(value int32) (err error) {
	return instance.SetProperty("AllowProjectionFromPC", value)
}

// GetAllowProjectionFromPC gets the value of AllowProjectionFromPC for the instance
func (instance *MDM_Policy_Config01_WirelessDisplay02) GetPropertyAllowProjectionFromPC() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowProjectionFromPC")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowProjectionFromPCOverInfrastructure sets the value of AllowProjectionFromPCOverInfrastructure for the instance
func (instance *MDM_Policy_Config01_WirelessDisplay02) SetPropertyAllowProjectionFromPCOverInfrastructure(value int32) (err error) {
	return instance.SetProperty("AllowProjectionFromPCOverInfrastructure", value)
}

// GetAllowProjectionFromPCOverInfrastructure gets the value of AllowProjectionFromPCOverInfrastructure for the instance
func (instance *MDM_Policy_Config01_WirelessDisplay02) GetPropertyAllowProjectionFromPCOverInfrastructure() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowProjectionFromPCOverInfrastructure")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowProjectionToPC sets the value of AllowProjectionToPC for the instance
func (instance *MDM_Policy_Config01_WirelessDisplay02) SetPropertyAllowProjectionToPC(value int32) (err error) {
	return instance.SetProperty("AllowProjectionToPC", value)
}

// GetAllowProjectionToPC gets the value of AllowProjectionToPC for the instance
func (instance *MDM_Policy_Config01_WirelessDisplay02) GetPropertyAllowProjectionToPC() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowProjectionToPC")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowProjectionToPCOverInfrastructure sets the value of AllowProjectionToPCOverInfrastructure for the instance
func (instance *MDM_Policy_Config01_WirelessDisplay02) SetPropertyAllowProjectionToPCOverInfrastructure(value int32) (err error) {
	return instance.SetProperty("AllowProjectionToPCOverInfrastructure", value)
}

// GetAllowProjectionToPCOverInfrastructure gets the value of AllowProjectionToPCOverInfrastructure for the instance
func (instance *MDM_Policy_Config01_WirelessDisplay02) GetPropertyAllowProjectionToPCOverInfrastructure() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowProjectionToPCOverInfrastructure")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowUserInputFromWirelessDisplayReceiver sets the value of AllowUserInputFromWirelessDisplayReceiver for the instance
func (instance *MDM_Policy_Config01_WirelessDisplay02) SetPropertyAllowUserInputFromWirelessDisplayReceiver(value int32) (err error) {
	return instance.SetProperty("AllowUserInputFromWirelessDisplayReceiver", value)
}

// GetAllowUserInputFromWirelessDisplayReceiver gets the value of AllowUserInputFromWirelessDisplayReceiver for the instance
func (instance *MDM_Policy_Config01_WirelessDisplay02) GetPropertyAllowUserInputFromWirelessDisplayReceiver() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowUserInputFromWirelessDisplayReceiver")
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
func (instance *MDM_Policy_Config01_WirelessDisplay02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_WirelessDisplay02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_WirelessDisplay02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_WirelessDisplay02) GetPropertyParentID() (value string, err error) {
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

// SetRequirePinForPairing sets the value of RequirePinForPairing for the instance
func (instance *MDM_Policy_Config01_WirelessDisplay02) SetPropertyRequirePinForPairing(value int32) (err error) {
	return instance.SetProperty("RequirePinForPairing", value)
}

// GetRequirePinForPairing gets the value of RequirePinForPairing for the instance
func (instance *MDM_Policy_Config01_WirelessDisplay02) GetPropertyRequirePinForPairing() (value int32, err error) {
	retValue, err := instance.GetProperty("RequirePinForPairing")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
