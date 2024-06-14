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

// MSFT_NetSecuritySettingData struct
type MSFT_NetSecuritySettingData struct {
	*MSFT_NetSettingData

	//
	AllowIPsecThroughNAT uint16

	//
	CertValidationLevel uint16

	//
	EnablePacketQueuing uint16

	//
	EnableStatefulFtp uint16

	//
	EnableStatefulPptp uint16

	//
	Exemptions uint32

	//
	KeyEncoding uint16

	//
	MaxSAIdleTimeSeconds uint32

	//
	Profile uint16

	//
	RemoteMachineTransportAuthorizationList string

	//
	RemoteMachineTunnelAuthorizationList string

	//
	RemoteUserTransportAuthorizationList string

	//
	RemoteUserTunnelAuthorizationList string

	//
	RequireFullAuthSupport uint16
}

func NewMSFT_NetSecuritySettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetSecuritySettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetSecuritySettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetSecuritySettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetSecuritySettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetSecuritySettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetAllowIPsecThroughNAT sets the value of AllowIPsecThroughNAT for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyAllowIPsecThroughNAT(value uint16) (err error) {
	return instance.SetProperty("AllowIPsecThroughNAT", (value))
}

// GetAllowIPsecThroughNAT gets the value of AllowIPsecThroughNAT for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyAllowIPsecThroughNAT() (value uint16, err error) {
	retValue, err := instance.GetProperty("AllowIPsecThroughNAT")
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

// SetCertValidationLevel sets the value of CertValidationLevel for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyCertValidationLevel(value uint16) (err error) {
	return instance.SetProperty("CertValidationLevel", (value))
}

// GetCertValidationLevel gets the value of CertValidationLevel for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyCertValidationLevel() (value uint16, err error) {
	retValue, err := instance.GetProperty("CertValidationLevel")
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

// SetEnablePacketQueuing sets the value of EnablePacketQueuing for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyEnablePacketQueuing(value uint16) (err error) {
	return instance.SetProperty("EnablePacketQueuing", (value))
}

// GetEnablePacketQueuing gets the value of EnablePacketQueuing for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyEnablePacketQueuing() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnablePacketQueuing")
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

// SetEnableStatefulFtp sets the value of EnableStatefulFtp for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyEnableStatefulFtp(value uint16) (err error) {
	return instance.SetProperty("EnableStatefulFtp", (value))
}

// GetEnableStatefulFtp gets the value of EnableStatefulFtp for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyEnableStatefulFtp() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnableStatefulFtp")
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

// SetEnableStatefulPptp sets the value of EnableStatefulPptp for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyEnableStatefulPptp(value uint16) (err error) {
	return instance.SetProperty("EnableStatefulPptp", (value))
}

// GetEnableStatefulPptp gets the value of EnableStatefulPptp for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyEnableStatefulPptp() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnableStatefulPptp")
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

// SetExemptions sets the value of Exemptions for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyExemptions(value uint32) (err error) {
	return instance.SetProperty("Exemptions", (value))
}

// GetExemptions gets the value of Exemptions for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyExemptions() (value uint32, err error) {
	retValue, err := instance.GetProperty("Exemptions")
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

// SetKeyEncoding sets the value of KeyEncoding for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyKeyEncoding(value uint16) (err error) {
	return instance.SetProperty("KeyEncoding", (value))
}

// GetKeyEncoding gets the value of KeyEncoding for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyKeyEncoding() (value uint16, err error) {
	retValue, err := instance.GetProperty("KeyEncoding")
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

// SetMaxSAIdleTimeSeconds sets the value of MaxSAIdleTimeSeconds for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyMaxSAIdleTimeSeconds(value uint32) (err error) {
	return instance.SetProperty("MaxSAIdleTimeSeconds", (value))
}

// GetMaxSAIdleTimeSeconds gets the value of MaxSAIdleTimeSeconds for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyMaxSAIdleTimeSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxSAIdleTimeSeconds")
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

// SetProfile sets the value of Profile for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyProfile(value uint16) (err error) {
	return instance.SetProperty("Profile", (value))
}

// GetProfile gets the value of Profile for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyProfile() (value uint16, err error) {
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

// SetRemoteMachineTransportAuthorizationList sets the value of RemoteMachineTransportAuthorizationList for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyRemoteMachineTransportAuthorizationList(value string) (err error) {
	return instance.SetProperty("RemoteMachineTransportAuthorizationList", (value))
}

// GetRemoteMachineTransportAuthorizationList gets the value of RemoteMachineTransportAuthorizationList for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyRemoteMachineTransportAuthorizationList() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteMachineTransportAuthorizationList")
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

// SetRemoteMachineTunnelAuthorizationList sets the value of RemoteMachineTunnelAuthorizationList for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyRemoteMachineTunnelAuthorizationList(value string) (err error) {
	return instance.SetProperty("RemoteMachineTunnelAuthorizationList", (value))
}

// GetRemoteMachineTunnelAuthorizationList gets the value of RemoteMachineTunnelAuthorizationList for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyRemoteMachineTunnelAuthorizationList() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteMachineTunnelAuthorizationList")
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

// SetRemoteUserTransportAuthorizationList sets the value of RemoteUserTransportAuthorizationList for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyRemoteUserTransportAuthorizationList(value string) (err error) {
	return instance.SetProperty("RemoteUserTransportAuthorizationList", (value))
}

// GetRemoteUserTransportAuthorizationList gets the value of RemoteUserTransportAuthorizationList for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyRemoteUserTransportAuthorizationList() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteUserTransportAuthorizationList")
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

// SetRemoteUserTunnelAuthorizationList sets the value of RemoteUserTunnelAuthorizationList for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyRemoteUserTunnelAuthorizationList(value string) (err error) {
	return instance.SetProperty("RemoteUserTunnelAuthorizationList", (value))
}

// GetRemoteUserTunnelAuthorizationList gets the value of RemoteUserTunnelAuthorizationList for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyRemoteUserTunnelAuthorizationList() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteUserTunnelAuthorizationList")
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

// SetRequireFullAuthSupport sets the value of RequireFullAuthSupport for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyRequireFullAuthSupport(value uint16) (err error) {
	return instance.SetProperty("RequireFullAuthSupport", (value))
}

// GetRequireFullAuthSupport gets the value of RequireFullAuthSupport for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyRequireFullAuthSupport() (value uint16, err error) {
	retValue, err := instance.GetProperty("RequireFullAuthSupport")
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
