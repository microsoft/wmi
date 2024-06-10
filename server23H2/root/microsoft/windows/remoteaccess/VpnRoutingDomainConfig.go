// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// VpnRoutingDomainConfig struct
type VpnRoutingDomainConfig struct {
	*RoutingDomainConfiguration

	//
	AvailabilityStatus string

	//
	DnsIPAddress []string

	//
	EnableQoS uint32

	//
	InterimAccountingPeriodSec uint32

	//
	IPAddressRange []string

	//
	IPv6Prefix string

	//
	MaximumVpnConnections uint32

	//
	NetBiosIPAddress []string

	//
	RxBandwidthKbps uint64

	//
	TenantName []string

	//
	TxBandwidthKbps uint64
}

func NewVpnRoutingDomainConfigEx1(instance *cim.WmiInstance) (newInstance *VpnRoutingDomainConfig, err error) {
	tmp, err := NewRoutingDomainConfigurationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &VpnRoutingDomainConfig{
		RoutingDomainConfiguration: tmp,
	}
	return
}

func NewVpnRoutingDomainConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnRoutingDomainConfig, err error) {
	tmp, err := NewRoutingDomainConfigurationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnRoutingDomainConfig{
		RoutingDomainConfiguration: tmp,
	}
	return
}

// SetAvailabilityStatus sets the value of AvailabilityStatus for the instance
func (instance *VpnRoutingDomainConfig) SetPropertyAvailabilityStatus(value string) (err error) {
	return instance.SetProperty("AvailabilityStatus", (value))
}

// GetAvailabilityStatus gets the value of AvailabilityStatus for the instance
func (instance *VpnRoutingDomainConfig) GetPropertyAvailabilityStatus() (value string, err error) {
	retValue, err := instance.GetProperty("AvailabilityStatus")
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

// SetDnsIPAddress sets the value of DnsIPAddress for the instance
func (instance *VpnRoutingDomainConfig) SetPropertyDnsIPAddress(value []string) (err error) {
	return instance.SetProperty("DnsIPAddress", (value))
}

// GetDnsIPAddress gets the value of DnsIPAddress for the instance
func (instance *VpnRoutingDomainConfig) GetPropertyDnsIPAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("DnsIPAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetEnableQoS sets the value of EnableQoS for the instance
func (instance *VpnRoutingDomainConfig) SetPropertyEnableQoS(value uint32) (err error) {
	return instance.SetProperty("EnableQoS", (value))
}

// GetEnableQoS gets the value of EnableQoS for the instance
func (instance *VpnRoutingDomainConfig) GetPropertyEnableQoS() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnableQoS")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetInterimAccountingPeriodSec sets the value of InterimAccountingPeriodSec for the instance
func (instance *VpnRoutingDomainConfig) SetPropertyInterimAccountingPeriodSec(value uint32) (err error) {
	return instance.SetProperty("InterimAccountingPeriodSec", (value))
}

// GetInterimAccountingPeriodSec gets the value of InterimAccountingPeriodSec for the instance
func (instance *VpnRoutingDomainConfig) GetPropertyInterimAccountingPeriodSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterimAccountingPeriodSec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetIPAddressRange sets the value of IPAddressRange for the instance
func (instance *VpnRoutingDomainConfig) SetPropertyIPAddressRange(value []string) (err error) {
	return instance.SetProperty("IPAddressRange", (value))
}

// GetIPAddressRange gets the value of IPAddressRange for the instance
func (instance *VpnRoutingDomainConfig) GetPropertyIPAddressRange() (value []string, err error) {
	retValue, err := instance.GetProperty("IPAddressRange")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetIPv6Prefix sets the value of IPv6Prefix for the instance
func (instance *VpnRoutingDomainConfig) SetPropertyIPv6Prefix(value string) (err error) {
	return instance.SetProperty("IPv6Prefix", (value))
}

// GetIPv6Prefix gets the value of IPv6Prefix for the instance
func (instance *VpnRoutingDomainConfig) GetPropertyIPv6Prefix() (value string, err error) {
	retValue, err := instance.GetProperty("IPv6Prefix")
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

// SetMaximumVpnConnections sets the value of MaximumVpnConnections for the instance
func (instance *VpnRoutingDomainConfig) SetPropertyMaximumVpnConnections(value uint32) (err error) {
	return instance.SetProperty("MaximumVpnConnections", (value))
}

// GetMaximumVpnConnections gets the value of MaximumVpnConnections for the instance
func (instance *VpnRoutingDomainConfig) GetPropertyMaximumVpnConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumVpnConnections")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetNetBiosIPAddress sets the value of NetBiosIPAddress for the instance
func (instance *VpnRoutingDomainConfig) SetPropertyNetBiosIPAddress(value []string) (err error) {
	return instance.SetProperty("NetBiosIPAddress", (value))
}

// GetNetBiosIPAddress gets the value of NetBiosIPAddress for the instance
func (instance *VpnRoutingDomainConfig) GetPropertyNetBiosIPAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("NetBiosIPAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetRxBandwidthKbps sets the value of RxBandwidthKbps for the instance
func (instance *VpnRoutingDomainConfig) SetPropertyRxBandwidthKbps(value uint64) (err error) {
	return instance.SetProperty("RxBandwidthKbps", (value))
}

// GetRxBandwidthKbps gets the value of RxBandwidthKbps for the instance
func (instance *VpnRoutingDomainConfig) GetPropertyRxBandwidthKbps() (value uint64, err error) {
	retValue, err := instance.GetProperty("RxBandwidthKbps")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTenantName sets the value of TenantName for the instance
func (instance *VpnRoutingDomainConfig) SetPropertyTenantName(value []string) (err error) {
	return instance.SetProperty("TenantName", (value))
}

// GetTenantName gets the value of TenantName for the instance
func (instance *VpnRoutingDomainConfig) GetPropertyTenantName() (value []string, err error) {
	retValue, err := instance.GetProperty("TenantName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetTxBandwidthKbps sets the value of TxBandwidthKbps for the instance
func (instance *VpnRoutingDomainConfig) SetPropertyTxBandwidthKbps(value uint64) (err error) {
	return instance.SetProperty("TxBandwidthKbps", (value))
}

// GetTxBandwidthKbps gets the value of TxBandwidthKbps for the instance
func (instance *VpnRoutingDomainConfig) GetPropertyTxBandwidthKbps() (value uint64, err error) {
	retValue, err := instance.GetProperty("TxBandwidthKbps")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
