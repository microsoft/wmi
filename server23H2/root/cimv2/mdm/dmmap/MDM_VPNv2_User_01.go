// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_VPNv2_User_01 struct
type MDM_VPNv2_User_01 struct {
	*cim.WmiInstance

	//
	AlwaysOn bool

	//
	AlwaysOnActive bool

	//
	ByPassForLocal bool

	//
	DataEncryption string

	//
	DisableAdvancedOptionsEditButton bool

	//
	DisableDisconnectButton bool

	//
	DisableIKEv2Fragmentation bool

	//
	DnsSuffix string

	//
	EdpModeId string

	//
	InstanceID string

	//
	IPv4InterfaceMetric int32

	//
	IPv6InterfaceMetric int32

	//
	NetworkOutageTime int32

	//
	ParentID string

	//
	PrivateNetwork bool

	//
	ProfileXML string

	//
	RegisterDNS bool

	//
	RememberCredentials bool

	//
	RequireVpnClientAppUI bool

	//
	TrustedNetworkDetection string

	//
	UseRasCredentials bool
}

func NewMDM_VPNv2_User_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_VPNv2_User_01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_User_01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_VPNv2_User_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_VPNv2_User_01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_User_01{
		WmiInstance: tmp,
	}
	return
}

// SetAlwaysOn sets the value of AlwaysOn for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyAlwaysOn(value bool) (err error) {
	return instance.SetProperty("AlwaysOn", (value))
}

// GetAlwaysOn gets the value of AlwaysOn for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyAlwaysOn() (value bool, err error) {
	retValue, err := instance.GetProperty("AlwaysOn")
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

// SetAlwaysOnActive sets the value of AlwaysOnActive for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyAlwaysOnActive(value bool) (err error) {
	return instance.SetProperty("AlwaysOnActive", (value))
}

// GetAlwaysOnActive gets the value of AlwaysOnActive for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyAlwaysOnActive() (value bool, err error) {
	retValue, err := instance.GetProperty("AlwaysOnActive")
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

// SetByPassForLocal sets the value of ByPassForLocal for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyByPassForLocal(value bool) (err error) {
	return instance.SetProperty("ByPassForLocal", (value))
}

// GetByPassForLocal gets the value of ByPassForLocal for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyByPassForLocal() (value bool, err error) {
	retValue, err := instance.GetProperty("ByPassForLocal")
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

// SetDataEncryption sets the value of DataEncryption for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyDataEncryption(value string) (err error) {
	return instance.SetProperty("DataEncryption", (value))
}

// GetDataEncryption gets the value of DataEncryption for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyDataEncryption() (value string, err error) {
	retValue, err := instance.GetProperty("DataEncryption")
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

// SetDisableAdvancedOptionsEditButton sets the value of DisableAdvancedOptionsEditButton for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyDisableAdvancedOptionsEditButton(value bool) (err error) {
	return instance.SetProperty("DisableAdvancedOptionsEditButton", (value))
}

// GetDisableAdvancedOptionsEditButton gets the value of DisableAdvancedOptionsEditButton for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyDisableAdvancedOptionsEditButton() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableAdvancedOptionsEditButton")
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

// SetDisableDisconnectButton sets the value of DisableDisconnectButton for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyDisableDisconnectButton(value bool) (err error) {
	return instance.SetProperty("DisableDisconnectButton", (value))
}

// GetDisableDisconnectButton gets the value of DisableDisconnectButton for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyDisableDisconnectButton() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableDisconnectButton")
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

// SetDisableIKEv2Fragmentation sets the value of DisableIKEv2Fragmentation for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyDisableIKEv2Fragmentation(value bool) (err error) {
	return instance.SetProperty("DisableIKEv2Fragmentation", (value))
}

// GetDisableIKEv2Fragmentation gets the value of DisableIKEv2Fragmentation for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyDisableIKEv2Fragmentation() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableIKEv2Fragmentation")
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

// SetDnsSuffix sets the value of DnsSuffix for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyDnsSuffix(value string) (err error) {
	return instance.SetProperty("DnsSuffix", (value))
}

// GetDnsSuffix gets the value of DnsSuffix for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyDnsSuffix() (value string, err error) {
	retValue, err := instance.GetProperty("DnsSuffix")
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

// SetEdpModeId sets the value of EdpModeId for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyEdpModeId(value string) (err error) {
	return instance.SetProperty("EdpModeId", (value))
}

// GetEdpModeId gets the value of EdpModeId for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyEdpModeId() (value string, err error) {
	retValue, err := instance.GetProperty("EdpModeId")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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

// SetIPv4InterfaceMetric sets the value of IPv4InterfaceMetric for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyIPv4InterfaceMetric(value int32) (err error) {
	return instance.SetProperty("IPv4InterfaceMetric", (value))
}

// GetIPv4InterfaceMetric gets the value of IPv4InterfaceMetric for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyIPv4InterfaceMetric() (value int32, err error) {
	retValue, err := instance.GetProperty("IPv4InterfaceMetric")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetIPv6InterfaceMetric sets the value of IPv6InterfaceMetric for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyIPv6InterfaceMetric(value int32) (err error) {
	return instance.SetProperty("IPv6InterfaceMetric", (value))
}

// GetIPv6InterfaceMetric gets the value of IPv6InterfaceMetric for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyIPv6InterfaceMetric() (value int32, err error) {
	retValue, err := instance.GetProperty("IPv6InterfaceMetric")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetNetworkOutageTime sets the value of NetworkOutageTime for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyNetworkOutageTime(value int32) (err error) {
	return instance.SetProperty("NetworkOutageTime", (value))
}

// GetNetworkOutageTime gets the value of NetworkOutageTime for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyNetworkOutageTime() (value int32, err error) {
	retValue, err := instance.GetProperty("NetworkOutageTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
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

// SetPrivateNetwork sets the value of PrivateNetwork for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyPrivateNetwork(value bool) (err error) {
	return instance.SetProperty("PrivateNetwork", (value))
}

// GetPrivateNetwork gets the value of PrivateNetwork for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyPrivateNetwork() (value bool, err error) {
	retValue, err := instance.GetProperty("PrivateNetwork")
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

// SetProfileXML sets the value of ProfileXML for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyProfileXML(value string) (err error) {
	return instance.SetProperty("ProfileXML", (value))
}

// GetProfileXML gets the value of ProfileXML for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyProfileXML() (value string, err error) {
	retValue, err := instance.GetProperty("ProfileXML")
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

// SetRegisterDNS sets the value of RegisterDNS for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyRegisterDNS(value bool) (err error) {
	return instance.SetProperty("RegisterDNS", (value))
}

// GetRegisterDNS gets the value of RegisterDNS for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyRegisterDNS() (value bool, err error) {
	retValue, err := instance.GetProperty("RegisterDNS")
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

// SetRememberCredentials sets the value of RememberCredentials for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyRememberCredentials(value bool) (err error) {
	return instance.SetProperty("RememberCredentials", (value))
}

// GetRememberCredentials gets the value of RememberCredentials for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyRememberCredentials() (value bool, err error) {
	retValue, err := instance.GetProperty("RememberCredentials")
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

// SetRequireVpnClientAppUI sets the value of RequireVpnClientAppUI for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyRequireVpnClientAppUI(value bool) (err error) {
	return instance.SetProperty("RequireVpnClientAppUI", (value))
}

// GetRequireVpnClientAppUI gets the value of RequireVpnClientAppUI for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyRequireVpnClientAppUI() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireVpnClientAppUI")
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

// SetTrustedNetworkDetection sets the value of TrustedNetworkDetection for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyTrustedNetworkDetection(value string) (err error) {
	return instance.SetProperty("TrustedNetworkDetection", (value))
}

// GetTrustedNetworkDetection gets the value of TrustedNetworkDetection for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyTrustedNetworkDetection() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedNetworkDetection")
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

// SetUseRasCredentials sets the value of UseRasCredentials for the instance
func (instance *MDM_VPNv2_User_01) SetPropertyUseRasCredentials(value bool) (err error) {
	return instance.SetProperty("UseRasCredentials", (value))
}

// GetUseRasCredentials gets the value of UseRasCredentials for the instance
func (instance *MDM_VPNv2_User_01) GetPropertyUseRasCredentials() (value bool, err error) {
	retValue, err := instance.GetProperty("UseRasCredentials")
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
