// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.Intel_ME
//////////////////////////////////////////////
package intel_me

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// AMT_EthernetPortSettings struct
type AMT_EthernetPortSettings struct {
	*CIM_SettingData

	// Default Gateway in a string format. For example: 10.12.232.1
	DefaultGateway string

	// Indicates whether DHCP is in use. When this parameter is set to true these parameters will be ignored: SubnetMask, DefaultGAteway, IPAddress, PrimaryDNS, SecondaryDNS.
	DHCPEnabled bool

	// String representation of IP address. Get operation - reports the acquired IP address (whether in static or DHCP mode). Put operation - sets the IP address (in static mode only).
	IPAddress string

	// Indicates whether the network link is up
	LinkIsUp bool

	// The MAC address used by Intel(R) AMT in a string format. For Example: 01-02-3f-b0-99-99. (This property can only be read and can't be changed.)
	MACAddress string

	// Primary DNS in a string format. For example: 10.12.232.1
	PrimaryDNS string

	// Secondary DNS in a string format. For example: 10.12.232.1
	SecondaryDNS string

	// Subnet mask in a string format. For example: 255.255.0.0
	SubnetMask string
}

func NewAMT_EthernetPortSettingsEx1(instance *cim.WmiInstance) (newInstance *AMT_EthernetPortSettings, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AMT_EthernetPortSettings{
		CIM_SettingData: tmp,
	}
	return
}

func NewAMT_EthernetPortSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AMT_EthernetPortSettings, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AMT_EthernetPortSettings{
		CIM_SettingData: tmp,
	}
	return
}

// SetDefaultGateway sets the value of DefaultGateway for the instance
func (instance *AMT_EthernetPortSettings) SetPropertyDefaultGateway(value string) (err error) {
	return instance.SetProperty("DefaultGateway", (value))
}

// GetDefaultGateway gets the value of DefaultGateway for the instance
func (instance *AMT_EthernetPortSettings) GetPropertyDefaultGateway() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultGateway")
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

// SetDHCPEnabled sets the value of DHCPEnabled for the instance
func (instance *AMT_EthernetPortSettings) SetPropertyDHCPEnabled(value bool) (err error) {
	return instance.SetProperty("DHCPEnabled", (value))
}

// GetDHCPEnabled gets the value of DHCPEnabled for the instance
func (instance *AMT_EthernetPortSettings) GetPropertyDHCPEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DHCPEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetIPAddress sets the value of IPAddress for the instance
func (instance *AMT_EthernetPortSettings) SetPropertyIPAddress(value string) (err error) {
	return instance.SetProperty("IPAddress", (value))
}

// GetIPAddress gets the value of IPAddress for the instance
func (instance *AMT_EthernetPortSettings) GetPropertyIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("IPAddress")
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

// SetLinkIsUp sets the value of LinkIsUp for the instance
func (instance *AMT_EthernetPortSettings) SetPropertyLinkIsUp(value bool) (err error) {
	return instance.SetProperty("LinkIsUp", (value))
}

// GetLinkIsUp gets the value of LinkIsUp for the instance
func (instance *AMT_EthernetPortSettings) GetPropertyLinkIsUp() (value bool, err error) {
	retValue, err := instance.GetProperty("LinkIsUp")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetMACAddress sets the value of MACAddress for the instance
func (instance *AMT_EthernetPortSettings) SetPropertyMACAddress(value string) (err error) {
	return instance.SetProperty("MACAddress", (value))
}

// GetMACAddress gets the value of MACAddress for the instance
func (instance *AMT_EthernetPortSettings) GetPropertyMACAddress() (value string, err error) {
	retValue, err := instance.GetProperty("MACAddress")
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

// SetPrimaryDNS sets the value of PrimaryDNS for the instance
func (instance *AMT_EthernetPortSettings) SetPropertyPrimaryDNS(value string) (err error) {
	return instance.SetProperty("PrimaryDNS", (value))
}

// GetPrimaryDNS gets the value of PrimaryDNS for the instance
func (instance *AMT_EthernetPortSettings) GetPropertyPrimaryDNS() (value string, err error) {
	retValue, err := instance.GetProperty("PrimaryDNS")
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

// SetSecondaryDNS sets the value of SecondaryDNS for the instance
func (instance *AMT_EthernetPortSettings) SetPropertySecondaryDNS(value string) (err error) {
	return instance.SetProperty("SecondaryDNS", (value))
}

// GetSecondaryDNS gets the value of SecondaryDNS for the instance
func (instance *AMT_EthernetPortSettings) GetPropertySecondaryDNS() (value string, err error) {
	retValue, err := instance.GetProperty("SecondaryDNS")
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

// SetSubnetMask sets the value of SubnetMask for the instance
func (instance *AMT_EthernetPortSettings) SetPropertySubnetMask(value string) (err error) {
	return instance.SetProperty("SubnetMask", (value))
}

// GetSubnetMask gets the value of SubnetMask for the instance
func (instance *AMT_EthernetPortSettings) GetPropertySubnetMask() (value string, err error) {
	retValue, err := instance.GetProperty("SubnetMask")
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
