// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RemoteAccessConfiguration struct
type RemoteAccessConfiguration struct {
	*cim.WmiInstance

	//
	BGPPeerConfig []RemoteAccessBgpPeerConfig

	//
	GlobalIpv4TransportInfo []uint8

	//
	GlobalIpv6TransportInfo []uint8

	//
	RadiusAccountingStatus string

	//
	RadiusServerConfigs []RemoteAccessRadiusConfig

	//
	RemoteAccessCapacityKbps uint64

	//
	RoutingDomainConfig []RemoteAccessRoutingDomainConfiguration

	//
	S2SInterfaceConfiguration []RemoteAccessS2SConfiguration

	//
	ServerAuthConfig RemoteAccessServerAuthConfiguration

	//
	ServerIpSecConfiguration interface{}

	//
	Version uint32

	//
	VPNAuthType string

	//
	VSIDs []VSIDConfiguration
}

func NewRemoteAccessConfigurationEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RemoteAccessConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewRemoteAccessConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetBGPPeerConfig sets the value of BGPPeerConfig for the instance
func (instance *RemoteAccessConfiguration) SetPropertyBGPPeerConfig(value []RemoteAccessBgpPeerConfig) (err error) {
	return instance.SetProperty("BGPPeerConfig", (value))
}

// GetBGPPeerConfig gets the value of BGPPeerConfig for the instance
func (instance *RemoteAccessConfiguration) GetPropertyBGPPeerConfig() (value []RemoteAccessBgpPeerConfig, err error) {
	retValue, err := instance.GetProperty("BGPPeerConfig")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(RemoteAccessBgpPeerConfig)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " RemoteAccessBgpPeerConfig is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, RemoteAccessBgpPeerConfig(valuetmp))
	}

	return
}

// SetGlobalIpv4TransportInfo sets the value of GlobalIpv4TransportInfo for the instance
func (instance *RemoteAccessConfiguration) SetPropertyGlobalIpv4TransportInfo(value []uint8) (err error) {
	return instance.SetProperty("GlobalIpv4TransportInfo", (value))
}

// GetGlobalIpv4TransportInfo gets the value of GlobalIpv4TransportInfo for the instance
func (instance *RemoteAccessConfiguration) GetPropertyGlobalIpv4TransportInfo() (value []uint8, err error) {
	retValue, err := instance.GetProperty("GlobalIpv4TransportInfo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetGlobalIpv6TransportInfo sets the value of GlobalIpv6TransportInfo for the instance
func (instance *RemoteAccessConfiguration) SetPropertyGlobalIpv6TransportInfo(value []uint8) (err error) {
	return instance.SetProperty("GlobalIpv6TransportInfo", (value))
}

// GetGlobalIpv6TransportInfo gets the value of GlobalIpv6TransportInfo for the instance
func (instance *RemoteAccessConfiguration) GetPropertyGlobalIpv6TransportInfo() (value []uint8, err error) {
	retValue, err := instance.GetProperty("GlobalIpv6TransportInfo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetRadiusAccountingStatus sets the value of RadiusAccountingStatus for the instance
func (instance *RemoteAccessConfiguration) SetPropertyRadiusAccountingStatus(value string) (err error) {
	return instance.SetProperty("RadiusAccountingStatus", (value))
}

// GetRadiusAccountingStatus gets the value of RadiusAccountingStatus for the instance
func (instance *RemoteAccessConfiguration) GetPropertyRadiusAccountingStatus() (value string, err error) {
	retValue, err := instance.GetProperty("RadiusAccountingStatus")
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

// SetRadiusServerConfigs sets the value of RadiusServerConfigs for the instance
func (instance *RemoteAccessConfiguration) SetPropertyRadiusServerConfigs(value []RemoteAccessRadiusConfig) (err error) {
	return instance.SetProperty("RadiusServerConfigs", (value))
}

// GetRadiusServerConfigs gets the value of RadiusServerConfigs for the instance
func (instance *RemoteAccessConfiguration) GetPropertyRadiusServerConfigs() (value []RemoteAccessRadiusConfig, err error) {
	retValue, err := instance.GetProperty("RadiusServerConfigs")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(RemoteAccessRadiusConfig)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " RemoteAccessRadiusConfig is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, RemoteAccessRadiusConfig(valuetmp))
	}

	return
}

// SetRemoteAccessCapacityKbps sets the value of RemoteAccessCapacityKbps for the instance
func (instance *RemoteAccessConfiguration) SetPropertyRemoteAccessCapacityKbps(value uint64) (err error) {
	return instance.SetProperty("RemoteAccessCapacityKbps", (value))
}

// GetRemoteAccessCapacityKbps gets the value of RemoteAccessCapacityKbps for the instance
func (instance *RemoteAccessConfiguration) GetPropertyRemoteAccessCapacityKbps() (value uint64, err error) {
	retValue, err := instance.GetProperty("RemoteAccessCapacityKbps")
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

// SetRoutingDomainConfig sets the value of RoutingDomainConfig for the instance
func (instance *RemoteAccessConfiguration) SetPropertyRoutingDomainConfig(value []RemoteAccessRoutingDomainConfiguration) (err error) {
	return instance.SetProperty("RoutingDomainConfig", (value))
}

// GetRoutingDomainConfig gets the value of RoutingDomainConfig for the instance
func (instance *RemoteAccessConfiguration) GetPropertyRoutingDomainConfig() (value []RemoteAccessRoutingDomainConfiguration, err error) {
	retValue, err := instance.GetProperty("RoutingDomainConfig")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(RemoteAccessRoutingDomainConfiguration)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " RemoteAccessRoutingDomainConfiguration is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, RemoteAccessRoutingDomainConfiguration(valuetmp))
	}

	return
}

// SetS2SInterfaceConfiguration sets the value of S2SInterfaceConfiguration for the instance
func (instance *RemoteAccessConfiguration) SetPropertyS2SInterfaceConfiguration(value []RemoteAccessS2SConfiguration) (err error) {
	return instance.SetProperty("S2SInterfaceConfiguration", (value))
}

// GetS2SInterfaceConfiguration gets the value of S2SInterfaceConfiguration for the instance
func (instance *RemoteAccessConfiguration) GetPropertyS2SInterfaceConfiguration() (value []RemoteAccessS2SConfiguration, err error) {
	retValue, err := instance.GetProperty("S2SInterfaceConfiguration")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(RemoteAccessS2SConfiguration)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " RemoteAccessS2SConfiguration is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, RemoteAccessS2SConfiguration(valuetmp))
	}

	return
}

// SetServerAuthConfig sets the value of ServerAuthConfig for the instance
func (instance *RemoteAccessConfiguration) SetPropertyServerAuthConfig(value RemoteAccessServerAuthConfiguration) (err error) {
	return instance.SetProperty("ServerAuthConfig", (value))
}

// GetServerAuthConfig gets the value of ServerAuthConfig for the instance
func (instance *RemoteAccessConfiguration) GetPropertyServerAuthConfig() (value RemoteAccessServerAuthConfiguration, err error) {
	retValue, err := instance.GetProperty("ServerAuthConfig")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RemoteAccessServerAuthConfiguration)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RemoteAccessServerAuthConfiguration is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RemoteAccessServerAuthConfiguration(valuetmp)

	return
}

// SetServerIpSecConfiguration sets the value of ServerIpSecConfiguration for the instance
func (instance *RemoteAccessConfiguration) SetPropertyServerIpSecConfiguration(value interface{}) (err error) {
	return instance.SetProperty("ServerIpSecConfiguration", (value))
}

// GetServerIpSecConfiguration gets the value of ServerIpSecConfiguration for the instance
func (instance *RemoteAccessConfiguration) GetPropertyServerIpSecConfiguration() (value interface{}, err error) {
	retValue, err := instance.GetProperty("ServerIpSecConfiguration")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetVersion sets the value of Version for the instance
func (instance *RemoteAccessConfiguration) SetPropertyVersion(value uint32) (err error) {
	return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *RemoteAccessConfiguration) GetPropertyVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("Version")
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

// SetVPNAuthType sets the value of VPNAuthType for the instance
func (instance *RemoteAccessConfiguration) SetPropertyVPNAuthType(value string) (err error) {
	return instance.SetProperty("VPNAuthType", (value))
}

// GetVPNAuthType gets the value of VPNAuthType for the instance
func (instance *RemoteAccessConfiguration) GetPropertyVPNAuthType() (value string, err error) {
	retValue, err := instance.GetProperty("VPNAuthType")
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

// SetVSIDs sets the value of VSIDs for the instance
func (instance *RemoteAccessConfiguration) SetPropertyVSIDs(value []VSIDConfiguration) (err error) {
	return instance.SetProperty("VSIDs", (value))
}

// GetVSIDs gets the value of VSIDs for the instance
func (instance *RemoteAccessConfiguration) GetPropertyVSIDs() (value []VSIDConfiguration, err error) {
	retValue, err := instance.GetProperty("VSIDs")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(VSIDConfiguration)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " VSIDConfiguration is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, VSIDConfiguration(valuetmp))
	}

	return
}
