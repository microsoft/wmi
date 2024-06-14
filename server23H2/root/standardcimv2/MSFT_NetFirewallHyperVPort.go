// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetFirewallHyperVPort struct
type MSFT_NetFirewallHyperVPort struct {
	*CIM_ManagedElement

	//
	Constrained uint16

	//
	InterfaceGuid string

	//
	NetworkType uint16

	//
	PartitionGuid string

	//
	PortName string

	//
	Profile uint16

	//
	SwitchName string

	//
	VMCreatorId string
}

func NewMSFT_NetFirewallHyperVPortEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetFirewallHyperVPort, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallHyperVPort{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMSFT_NetFirewallHyperVPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetFirewallHyperVPort, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallHyperVPort{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetConstrained sets the value of Constrained for the instance
func (instance *MSFT_NetFirewallHyperVPort) SetPropertyConstrained(value uint16) (err error) {
	return instance.SetProperty("Constrained", (value))
}

// GetConstrained gets the value of Constrained for the instance
func (instance *MSFT_NetFirewallHyperVPort) GetPropertyConstrained() (value uint16, err error) {
	retValue, err := instance.GetProperty("Constrained")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetInterfaceGuid sets the value of InterfaceGuid for the instance
func (instance *MSFT_NetFirewallHyperVPort) SetPropertyInterfaceGuid(value string) (err error) {
	return instance.SetProperty("InterfaceGuid", (value))
}

// GetInterfaceGuid gets the value of InterfaceGuid for the instance
func (instance *MSFT_NetFirewallHyperVPort) GetPropertyInterfaceGuid() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceGuid")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetNetworkType sets the value of NetworkType for the instance
func (instance *MSFT_NetFirewallHyperVPort) SetPropertyNetworkType(value uint16) (err error) {
	return instance.SetProperty("NetworkType", (value))
}

// GetNetworkType gets the value of NetworkType for the instance
func (instance *MSFT_NetFirewallHyperVPort) GetPropertyNetworkType() (value uint16, err error) {
	retValue, err := instance.GetProperty("NetworkType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetPartitionGuid sets the value of PartitionGuid for the instance
func (instance *MSFT_NetFirewallHyperVPort) SetPropertyPartitionGuid(value string) (err error) {
	return instance.SetProperty("PartitionGuid", (value))
}

// GetPartitionGuid gets the value of PartitionGuid for the instance
func (instance *MSFT_NetFirewallHyperVPort) GetPropertyPartitionGuid() (value string, err error) {
	retValue, err := instance.GetProperty("PartitionGuid")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetPortName sets the value of PortName for the instance
func (instance *MSFT_NetFirewallHyperVPort) SetPropertyPortName(value string) (err error) {
	return instance.SetProperty("PortName", (value))
}

// GetPortName gets the value of PortName for the instance
func (instance *MSFT_NetFirewallHyperVPort) GetPropertyPortName() (value string, err error) {
	retValue, err := instance.GetProperty("PortName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetProfile sets the value of Profile for the instance
func (instance *MSFT_NetFirewallHyperVPort) SetPropertyProfile(value uint16) (err error) {
	return instance.SetProperty("Profile", (value))
}

// GetProfile gets the value of Profile for the instance
func (instance *MSFT_NetFirewallHyperVPort) GetPropertyProfile() (value uint16, err error) {
	retValue, err := instance.GetProperty("Profile")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetSwitchName sets the value of SwitchName for the instance
func (instance *MSFT_NetFirewallHyperVPort) SetPropertySwitchName(value string) (err error) {
	return instance.SetProperty("SwitchName", (value))
}

// GetSwitchName gets the value of SwitchName for the instance
func (instance *MSFT_NetFirewallHyperVPort) GetPropertySwitchName() (value string, err error) {
	retValue, err := instance.GetProperty("SwitchName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetVMCreatorId sets the value of VMCreatorId for the instance
func (instance *MSFT_NetFirewallHyperVPort) SetPropertyVMCreatorId(value string) (err error) {
	return instance.SetProperty("VMCreatorId", (value))
}

// GetVMCreatorId gets the value of VMCreatorId for the instance
func (instance *MSFT_NetFirewallHyperVPort) GetPropertyVMCreatorId() (value string, err error) {
	retValue, err := instance.GetProperty("VMCreatorId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
