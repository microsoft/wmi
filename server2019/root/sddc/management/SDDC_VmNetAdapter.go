// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.SDDC.Management
//////////////////////////////////////////////
package management

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// SDDC_VmNetAdapter struct
type SDDC_VmNetAdapter struct {
	cim.WmiInstance

	//
	Acls []string

	//
	Id string

	//
	IpAddresses []string

	//
	IsDynamicMacAddressEnabled bool

	//
	IsMacAddressSpoofingAllowed bool

	//
	IsManagementOS bool

	//
	Name string

	//
	PhysicalAddress string

	//
	Status uint16

	//
	SwitchName string

	//
	VLanId string

	//
	VmId string

	//
	VmName string
}

// SetAcls sets the value of Acls for the instance
func (instance *SDDC_VmNetAdapter) SetPropertyAcls(value []string) (err error) {
	return instance.SetProperty("Acls", value)
}

// GetAcls gets the value of Acls for the instance
func (instance *SDDC_VmNetAdapter) GetPropertyAcls() (value []string, err error) {
	retValue, err := instance.GetProperty("Acls")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *SDDC_VmNetAdapter) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *SDDC_VmNetAdapter) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIpAddresses sets the value of IpAddresses for the instance
func (instance *SDDC_VmNetAdapter) SetPropertyIpAddresses(value []string) (err error) {
	return instance.SetProperty("IpAddresses", value)
}

// GetIpAddresses gets the value of IpAddresses for the instance
func (instance *SDDC_VmNetAdapter) GetPropertyIpAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("IpAddresses")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsDynamicMacAddressEnabled sets the value of IsDynamicMacAddressEnabled for the instance
func (instance *SDDC_VmNetAdapter) SetPropertyIsDynamicMacAddressEnabled(value bool) (err error) {
	return instance.SetProperty("IsDynamicMacAddressEnabled", value)
}

// GetIsDynamicMacAddressEnabled gets the value of IsDynamicMacAddressEnabled for the instance
func (instance *SDDC_VmNetAdapter) GetPropertyIsDynamicMacAddressEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDynamicMacAddressEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsMacAddressSpoofingAllowed sets the value of IsMacAddressSpoofingAllowed for the instance
func (instance *SDDC_VmNetAdapter) SetPropertyIsMacAddressSpoofingAllowed(value bool) (err error) {
	return instance.SetProperty("IsMacAddressSpoofingAllowed", value)
}

// GetIsMacAddressSpoofingAllowed gets the value of IsMacAddressSpoofingAllowed for the instance
func (instance *SDDC_VmNetAdapter) GetPropertyIsMacAddressSpoofingAllowed() (value bool, err error) {
	retValue, err := instance.GetProperty("IsMacAddressSpoofingAllowed")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsManagementOS sets the value of IsManagementOS for the instance
func (instance *SDDC_VmNetAdapter) SetPropertyIsManagementOS(value bool) (err error) {
	return instance.SetProperty("IsManagementOS", value)
}

// GetIsManagementOS gets the value of IsManagementOS for the instance
func (instance *SDDC_VmNetAdapter) GetPropertyIsManagementOS() (value bool, err error) {
	retValue, err := instance.GetProperty("IsManagementOS")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *SDDC_VmNetAdapter) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *SDDC_VmNetAdapter) GetPropertyName() (value string, err error) {
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

// SetPhysicalAddress sets the value of PhysicalAddress for the instance
func (instance *SDDC_VmNetAdapter) SetPropertyPhysicalAddress(value string) (err error) {
	return instance.SetProperty("PhysicalAddress", value)
}

// GetPhysicalAddress gets the value of PhysicalAddress for the instance
func (instance *SDDC_VmNetAdapter) GetPropertyPhysicalAddress() (value string, err error) {
	retValue, err := instance.GetProperty("PhysicalAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *SDDC_VmNetAdapter) SetPropertyStatus(value uint16) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *SDDC_VmNetAdapter) GetPropertyStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSwitchName sets the value of SwitchName for the instance
func (instance *SDDC_VmNetAdapter) SetPropertySwitchName(value string) (err error) {
	return instance.SetProperty("SwitchName", value)
}

// GetSwitchName gets the value of SwitchName for the instance
func (instance *SDDC_VmNetAdapter) GetPropertySwitchName() (value string, err error) {
	retValue, err := instance.GetProperty("SwitchName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVLanId sets the value of VLanId for the instance
func (instance *SDDC_VmNetAdapter) SetPropertyVLanId(value string) (err error) {
	return instance.SetProperty("VLanId", value)
}

// GetVLanId gets the value of VLanId for the instance
func (instance *SDDC_VmNetAdapter) GetPropertyVLanId() (value string, err error) {
	retValue, err := instance.GetProperty("VLanId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVmId sets the value of VmId for the instance
func (instance *SDDC_VmNetAdapter) SetPropertyVmId(value string) (err error) {
	return instance.SetProperty("VmId", value)
}

// GetVmId gets the value of VmId for the instance
func (instance *SDDC_VmNetAdapter) GetPropertyVmId() (value string, err error) {
	retValue, err := instance.GetProperty("VmId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVmName sets the value of VmName for the instance
func (instance *SDDC_VmNetAdapter) SetPropertyVmName(value string) (err error) {
	return instance.SetProperty("VmName", value)
}

// GetVmName gets the value of VmName for the instance
func (instance *SDDC_VmNetAdapter) GetPropertyVmName() (value string, err error) {
	retValue, err := instance.GetProperty("VmName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
