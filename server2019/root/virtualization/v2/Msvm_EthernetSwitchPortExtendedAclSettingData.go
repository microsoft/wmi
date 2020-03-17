// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_EthernetSwitchPortExtendedAclSettingData struct
type Msvm_EthernetSwitchPortExtendedAclSettingData struct {
	Msvm_EthernetSwitchPortFeatureSettingData

	//
	Action uint8

	//
	Direction uint8

	//
	IdleSessionTimeout uint16

	//
	IsolationID uint32

	//
	LocalIPAddress string

	//
	LocalPort string

	//
	Name string

	//
	Protocol string

	//
	RemoteIPAddress string

	//
	RemotePort string

	//
	Stateful bool

	//
	Weight uint16
}

// SetAction sets the value of Action for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyAction(value uint8) (err error) {
	return instance.SetProperty("Action", value)
}

// GetAction gets the value of Action for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyAction() (value uint8, err error) {
	retValue, err := instance.GetProperty("Action")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDirection sets the value of Direction for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyDirection(value uint8) (err error) {
	return instance.SetProperty("Direction", value)
}

// GetDirection gets the value of Direction for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyDirection() (value uint8, err error) {
	retValue, err := instance.GetProperty("Direction")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIdleSessionTimeout sets the value of IdleSessionTimeout for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyIdleSessionTimeout(value uint16) (err error) {
	return instance.SetProperty("IdleSessionTimeout", value)
}

// GetIdleSessionTimeout gets the value of IdleSessionTimeout for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyIdleSessionTimeout() (value uint16, err error) {
	retValue, err := instance.GetProperty("IdleSessionTimeout")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsolationID sets the value of IsolationID for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyIsolationID(value uint32) (err error) {
	return instance.SetProperty("IsolationID", value)
}

// GetIsolationID gets the value of IsolationID for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyIsolationID() (value uint32, err error) {
	retValue, err := instance.GetProperty("IsolationID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalIPAddress sets the value of LocalIPAddress for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyLocalIPAddress(value string) (err error) {
	return instance.SetProperty("LocalIPAddress", value)
}

// GetLocalIPAddress gets the value of LocalIPAddress for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyLocalIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("LocalIPAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalPort sets the value of LocalPort for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyLocalPort(value string) (err error) {
	return instance.SetProperty("LocalPort", value)
}

// GetLocalPort gets the value of LocalPort for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyLocalPort() (value string, err error) {
	retValue, err := instance.GetProperty("LocalPort")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtocol sets the value of Protocol for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyProtocol(value string) (err error) {
	return instance.SetProperty("Protocol", value)
}

// GetProtocol gets the value of Protocol for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyProtocol() (value string, err error) {
	retValue, err := instance.GetProperty("Protocol")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteIPAddress sets the value of RemoteIPAddress for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyRemoteIPAddress(value string) (err error) {
	return instance.SetProperty("RemoteIPAddress", value)
}

// GetRemoteIPAddress gets the value of RemoteIPAddress for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyRemoteIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteIPAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemotePort sets the value of RemotePort for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyRemotePort(value string) (err error) {
	return instance.SetProperty("RemotePort", value)
}

// GetRemotePort gets the value of RemotePort for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyRemotePort() (value string, err error) {
	retValue, err := instance.GetProperty("RemotePort")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStateful sets the value of Stateful for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyStateful(value bool) (err error) {
	return instance.SetProperty("Stateful", value)
}

// GetStateful gets the value of Stateful for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyStateful() (value bool, err error) {
	retValue, err := instance.GetProperty("Stateful")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWeight sets the value of Weight for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyWeight(value uint16) (err error) {
	return instance.SetProperty("Weight", value)
}

// GetWeight gets the value of Weight for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyWeight() (value uint16, err error) {
	retValue, err := instance.GetProperty("Weight")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities")
}
