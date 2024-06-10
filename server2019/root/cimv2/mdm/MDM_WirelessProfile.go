// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_WirelessProfile struct
type MDM_WirelessProfile struct {
	*cim.WmiInstance

	//
	AutoConnect bool

	//
	ConnectionType uint8

	//
	ConnectToMorePreferedNetwork bool

	//
	ConnectWhenNotBroadcasting bool

	//
	EnableFIPSCompliance bool

	//
	Name string

	//
	OneXAuthenticationMode uint8

	//
	OneXCacheUserData bool

	//
	OneXEAPType uint8

	//
	OneXEAPXml string

	//
	OneXSingleSignOnAllowAdditionalDialogs bool

	//
	OneXSingleSignOnMaxDelay uint32

	//
	OneXSingleSignOnType uint8

	//
	OneXSingleSignOnUserBasedVirtualLAN bool

	//
	PMKCacheMode bool

	//
	PMKCacheSize uint32

	//
	PMKCacheTTL uint32

	//
	PreAuthMode bool

	//
	PreAuthThrottle uint32

	//
	SecurityAuthentication uint8

	//
	SecurityEncryption uint8

	//
	SharedKeyMaterial string

	//
	SharedKeyProtected bool

	//
	SharedKeyType uint8

	//
	SSID string
}

func NewMDM_WirelessProfileEx1(instance *cim.WmiInstance) (newInstance *MDM_WirelessProfile, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_WirelessProfile{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_WirelessProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_WirelessProfile, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_WirelessProfile{
		WmiInstance: tmp,
	}
	return
}

// SetAutoConnect sets the value of AutoConnect for the instance
func (instance *MDM_WirelessProfile) SetPropertyAutoConnect(value bool) (err error) {
	return instance.SetProperty("AutoConnect", (value))
}

// GetAutoConnect gets the value of AutoConnect for the instance
func (instance *MDM_WirelessProfile) GetPropertyAutoConnect() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoConnect")
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

// SetConnectionType sets the value of ConnectionType for the instance
func (instance *MDM_WirelessProfile) SetPropertyConnectionType(value uint8) (err error) {
	return instance.SetProperty("ConnectionType", (value))
}

// GetConnectionType gets the value of ConnectionType for the instance
func (instance *MDM_WirelessProfile) GetPropertyConnectionType() (value uint8, err error) {
	retValue, err := instance.GetProperty("ConnectionType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetConnectToMorePreferedNetwork sets the value of ConnectToMorePreferedNetwork for the instance
func (instance *MDM_WirelessProfile) SetPropertyConnectToMorePreferedNetwork(value bool) (err error) {
	return instance.SetProperty("ConnectToMorePreferedNetwork", (value))
}

// GetConnectToMorePreferedNetwork gets the value of ConnectToMorePreferedNetwork for the instance
func (instance *MDM_WirelessProfile) GetPropertyConnectToMorePreferedNetwork() (value bool, err error) {
	retValue, err := instance.GetProperty("ConnectToMorePreferedNetwork")
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

// SetConnectWhenNotBroadcasting sets the value of ConnectWhenNotBroadcasting for the instance
func (instance *MDM_WirelessProfile) SetPropertyConnectWhenNotBroadcasting(value bool) (err error) {
	return instance.SetProperty("ConnectWhenNotBroadcasting", (value))
}

// GetConnectWhenNotBroadcasting gets the value of ConnectWhenNotBroadcasting for the instance
func (instance *MDM_WirelessProfile) GetPropertyConnectWhenNotBroadcasting() (value bool, err error) {
	retValue, err := instance.GetProperty("ConnectWhenNotBroadcasting")
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

// SetEnableFIPSCompliance sets the value of EnableFIPSCompliance for the instance
func (instance *MDM_WirelessProfile) SetPropertyEnableFIPSCompliance(value bool) (err error) {
	return instance.SetProperty("EnableFIPSCompliance", (value))
}

// GetEnableFIPSCompliance gets the value of EnableFIPSCompliance for the instance
func (instance *MDM_WirelessProfile) GetPropertyEnableFIPSCompliance() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableFIPSCompliance")
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

// SetName sets the value of Name for the instance
func (instance *MDM_WirelessProfile) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MDM_WirelessProfile) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetOneXAuthenticationMode sets the value of OneXAuthenticationMode for the instance
func (instance *MDM_WirelessProfile) SetPropertyOneXAuthenticationMode(value uint8) (err error) {
	return instance.SetProperty("OneXAuthenticationMode", (value))
}

// GetOneXAuthenticationMode gets the value of OneXAuthenticationMode for the instance
func (instance *MDM_WirelessProfile) GetPropertyOneXAuthenticationMode() (value uint8, err error) {
	retValue, err := instance.GetProperty("OneXAuthenticationMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetOneXCacheUserData sets the value of OneXCacheUserData for the instance
func (instance *MDM_WirelessProfile) SetPropertyOneXCacheUserData(value bool) (err error) {
	return instance.SetProperty("OneXCacheUserData", (value))
}

// GetOneXCacheUserData gets the value of OneXCacheUserData for the instance
func (instance *MDM_WirelessProfile) GetPropertyOneXCacheUserData() (value bool, err error) {
	retValue, err := instance.GetProperty("OneXCacheUserData")
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

// SetOneXEAPType sets the value of OneXEAPType for the instance
func (instance *MDM_WirelessProfile) SetPropertyOneXEAPType(value uint8) (err error) {
	return instance.SetProperty("OneXEAPType", (value))
}

// GetOneXEAPType gets the value of OneXEAPType for the instance
func (instance *MDM_WirelessProfile) GetPropertyOneXEAPType() (value uint8, err error) {
	retValue, err := instance.GetProperty("OneXEAPType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetOneXEAPXml sets the value of OneXEAPXml for the instance
func (instance *MDM_WirelessProfile) SetPropertyOneXEAPXml(value string) (err error) {
	return instance.SetProperty("OneXEAPXml", (value))
}

// GetOneXEAPXml gets the value of OneXEAPXml for the instance
func (instance *MDM_WirelessProfile) GetPropertyOneXEAPXml() (value string, err error) {
	retValue, err := instance.GetProperty("OneXEAPXml")
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

// SetOneXSingleSignOnAllowAdditionalDialogs sets the value of OneXSingleSignOnAllowAdditionalDialogs for the instance
func (instance *MDM_WirelessProfile) SetPropertyOneXSingleSignOnAllowAdditionalDialogs(value bool) (err error) {
	return instance.SetProperty("OneXSingleSignOnAllowAdditionalDialogs", (value))
}

// GetOneXSingleSignOnAllowAdditionalDialogs gets the value of OneXSingleSignOnAllowAdditionalDialogs for the instance
func (instance *MDM_WirelessProfile) GetPropertyOneXSingleSignOnAllowAdditionalDialogs() (value bool, err error) {
	retValue, err := instance.GetProperty("OneXSingleSignOnAllowAdditionalDialogs")
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

// SetOneXSingleSignOnMaxDelay sets the value of OneXSingleSignOnMaxDelay for the instance
func (instance *MDM_WirelessProfile) SetPropertyOneXSingleSignOnMaxDelay(value uint32) (err error) {
	return instance.SetProperty("OneXSingleSignOnMaxDelay", (value))
}

// GetOneXSingleSignOnMaxDelay gets the value of OneXSingleSignOnMaxDelay for the instance
func (instance *MDM_WirelessProfile) GetPropertyOneXSingleSignOnMaxDelay() (value uint32, err error) {
	retValue, err := instance.GetProperty("OneXSingleSignOnMaxDelay")
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

// SetOneXSingleSignOnType sets the value of OneXSingleSignOnType for the instance
func (instance *MDM_WirelessProfile) SetPropertyOneXSingleSignOnType(value uint8) (err error) {
	return instance.SetProperty("OneXSingleSignOnType", (value))
}

// GetOneXSingleSignOnType gets the value of OneXSingleSignOnType for the instance
func (instance *MDM_WirelessProfile) GetPropertyOneXSingleSignOnType() (value uint8, err error) {
	retValue, err := instance.GetProperty("OneXSingleSignOnType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetOneXSingleSignOnUserBasedVirtualLAN sets the value of OneXSingleSignOnUserBasedVirtualLAN for the instance
func (instance *MDM_WirelessProfile) SetPropertyOneXSingleSignOnUserBasedVirtualLAN(value bool) (err error) {
	return instance.SetProperty("OneXSingleSignOnUserBasedVirtualLAN", (value))
}

// GetOneXSingleSignOnUserBasedVirtualLAN gets the value of OneXSingleSignOnUserBasedVirtualLAN for the instance
func (instance *MDM_WirelessProfile) GetPropertyOneXSingleSignOnUserBasedVirtualLAN() (value bool, err error) {
	retValue, err := instance.GetProperty("OneXSingleSignOnUserBasedVirtualLAN")
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

// SetPMKCacheMode sets the value of PMKCacheMode for the instance
func (instance *MDM_WirelessProfile) SetPropertyPMKCacheMode(value bool) (err error) {
	return instance.SetProperty("PMKCacheMode", (value))
}

// GetPMKCacheMode gets the value of PMKCacheMode for the instance
func (instance *MDM_WirelessProfile) GetPropertyPMKCacheMode() (value bool, err error) {
	retValue, err := instance.GetProperty("PMKCacheMode")
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

// SetPMKCacheSize sets the value of PMKCacheSize for the instance
func (instance *MDM_WirelessProfile) SetPropertyPMKCacheSize(value uint32) (err error) {
	return instance.SetProperty("PMKCacheSize", (value))
}

// GetPMKCacheSize gets the value of PMKCacheSize for the instance
func (instance *MDM_WirelessProfile) GetPropertyPMKCacheSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("PMKCacheSize")
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

// SetPMKCacheTTL sets the value of PMKCacheTTL for the instance
func (instance *MDM_WirelessProfile) SetPropertyPMKCacheTTL(value uint32) (err error) {
	return instance.SetProperty("PMKCacheTTL", (value))
}

// GetPMKCacheTTL gets the value of PMKCacheTTL for the instance
func (instance *MDM_WirelessProfile) GetPropertyPMKCacheTTL() (value uint32, err error) {
	retValue, err := instance.GetProperty("PMKCacheTTL")
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

// SetPreAuthMode sets the value of PreAuthMode for the instance
func (instance *MDM_WirelessProfile) SetPropertyPreAuthMode(value bool) (err error) {
	return instance.SetProperty("PreAuthMode", (value))
}

// GetPreAuthMode gets the value of PreAuthMode for the instance
func (instance *MDM_WirelessProfile) GetPropertyPreAuthMode() (value bool, err error) {
	retValue, err := instance.GetProperty("PreAuthMode")
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

// SetPreAuthThrottle sets the value of PreAuthThrottle for the instance
func (instance *MDM_WirelessProfile) SetPropertyPreAuthThrottle(value uint32) (err error) {
	return instance.SetProperty("PreAuthThrottle", (value))
}

// GetPreAuthThrottle gets the value of PreAuthThrottle for the instance
func (instance *MDM_WirelessProfile) GetPropertyPreAuthThrottle() (value uint32, err error) {
	retValue, err := instance.GetProperty("PreAuthThrottle")
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

// SetSecurityAuthentication sets the value of SecurityAuthentication for the instance
func (instance *MDM_WirelessProfile) SetPropertySecurityAuthentication(value uint8) (err error) {
	return instance.SetProperty("SecurityAuthentication", (value))
}

// GetSecurityAuthentication gets the value of SecurityAuthentication for the instance
func (instance *MDM_WirelessProfile) GetPropertySecurityAuthentication() (value uint8, err error) {
	retValue, err := instance.GetProperty("SecurityAuthentication")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetSecurityEncryption sets the value of SecurityEncryption for the instance
func (instance *MDM_WirelessProfile) SetPropertySecurityEncryption(value uint8) (err error) {
	return instance.SetProperty("SecurityEncryption", (value))
}

// GetSecurityEncryption gets the value of SecurityEncryption for the instance
func (instance *MDM_WirelessProfile) GetPropertySecurityEncryption() (value uint8, err error) {
	retValue, err := instance.GetProperty("SecurityEncryption")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetSharedKeyMaterial sets the value of SharedKeyMaterial for the instance
func (instance *MDM_WirelessProfile) SetPropertySharedKeyMaterial(value string) (err error) {
	return instance.SetProperty("SharedKeyMaterial", (value))
}

// GetSharedKeyMaterial gets the value of SharedKeyMaterial for the instance
func (instance *MDM_WirelessProfile) GetPropertySharedKeyMaterial() (value string, err error) {
	retValue, err := instance.GetProperty("SharedKeyMaterial")
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

// SetSharedKeyProtected sets the value of SharedKeyProtected for the instance
func (instance *MDM_WirelessProfile) SetPropertySharedKeyProtected(value bool) (err error) {
	return instance.SetProperty("SharedKeyProtected", (value))
}

// GetSharedKeyProtected gets the value of SharedKeyProtected for the instance
func (instance *MDM_WirelessProfile) GetPropertySharedKeyProtected() (value bool, err error) {
	retValue, err := instance.GetProperty("SharedKeyProtected")
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

// SetSharedKeyType sets the value of SharedKeyType for the instance
func (instance *MDM_WirelessProfile) SetPropertySharedKeyType(value uint8) (err error) {
	return instance.SetProperty("SharedKeyType", (value))
}

// GetSharedKeyType gets the value of SharedKeyType for the instance
func (instance *MDM_WirelessProfile) GetPropertySharedKeyType() (value uint8, err error) {
	retValue, err := instance.GetProperty("SharedKeyType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetSSID sets the value of SSID for the instance
func (instance *MDM_WirelessProfile) SetPropertySSID(value string) (err error) {
	return instance.SetProperty("SSID", (value))
}

// GetSSID gets the value of SSID for the instance
func (instance *MDM_WirelessProfile) GetPropertySSID() (value string, err error) {
	retValue, err := instance.GetProperty("SSID")
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
