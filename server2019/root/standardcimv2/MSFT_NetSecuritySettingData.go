// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetSecuritySettingData struct
type MSFT_NetSecuritySettingData struct {
	MSFT_NetSettingData

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

// SetAllowIPsecThroughNAT sets the value of AllowIPsecThroughNAT for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyAllowIPsecThroughNAT(value uint16) (err error) {
	return instance.SetProperty("AllowIPsecThroughNAT", value)
}

// GetAllowIPsecThroughNAT gets the value of AllowIPsecThroughNAT for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyAllowIPsecThroughNAT() (value uint16, err error) {
	retValue, err := instance.GetProperty("AllowIPsecThroughNAT")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCertValidationLevel sets the value of CertValidationLevel for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyCertValidationLevel(value uint16) (err error) {
	return instance.SetProperty("CertValidationLevel", value)
}

// GetCertValidationLevel gets the value of CertValidationLevel for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyCertValidationLevel() (value uint16, err error) {
	retValue, err := instance.GetProperty("CertValidationLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnablePacketQueuing sets the value of EnablePacketQueuing for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyEnablePacketQueuing(value uint16) (err error) {
	return instance.SetProperty("EnablePacketQueuing", value)
}

// GetEnablePacketQueuing gets the value of EnablePacketQueuing for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyEnablePacketQueuing() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnablePacketQueuing")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableStatefulFtp sets the value of EnableStatefulFtp for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyEnableStatefulFtp(value uint16) (err error) {
	return instance.SetProperty("EnableStatefulFtp", value)
}

// GetEnableStatefulFtp gets the value of EnableStatefulFtp for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyEnableStatefulFtp() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnableStatefulFtp")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableStatefulPptp sets the value of EnableStatefulPptp for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyEnableStatefulPptp(value uint16) (err error) {
	return instance.SetProperty("EnableStatefulPptp", value)
}

// GetEnableStatefulPptp gets the value of EnableStatefulPptp for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyEnableStatefulPptp() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnableStatefulPptp")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExemptions sets the value of Exemptions for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyExemptions(value uint32) (err error) {
	return instance.SetProperty("Exemptions", value)
}

// GetExemptions gets the value of Exemptions for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyExemptions() (value uint32, err error) {
	retValue, err := instance.GetProperty("Exemptions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKeyEncoding sets the value of KeyEncoding for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyKeyEncoding(value uint16) (err error) {
	return instance.SetProperty("KeyEncoding", value)
}

// GetKeyEncoding gets the value of KeyEncoding for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyKeyEncoding() (value uint16, err error) {
	retValue, err := instance.GetProperty("KeyEncoding")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxSAIdleTimeSeconds sets the value of MaxSAIdleTimeSeconds for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyMaxSAIdleTimeSeconds(value uint32) (err error) {
	return instance.SetProperty("MaxSAIdleTimeSeconds", value)
}

// GetMaxSAIdleTimeSeconds gets the value of MaxSAIdleTimeSeconds for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyMaxSAIdleTimeSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxSAIdleTimeSeconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProfile sets the value of Profile for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyProfile(value uint16) (err error) {
	return instance.SetProperty("Profile", value)
}

// GetProfile gets the value of Profile for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyProfile() (value uint16, err error) {
	retValue, err := instance.GetProperty("Profile")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteMachineTransportAuthorizationList sets the value of RemoteMachineTransportAuthorizationList for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyRemoteMachineTransportAuthorizationList(value string) (err error) {
	return instance.SetProperty("RemoteMachineTransportAuthorizationList", value)
}

// GetRemoteMachineTransportAuthorizationList gets the value of RemoteMachineTransportAuthorizationList for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyRemoteMachineTransportAuthorizationList() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteMachineTransportAuthorizationList")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteMachineTunnelAuthorizationList sets the value of RemoteMachineTunnelAuthorizationList for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyRemoteMachineTunnelAuthorizationList(value string) (err error) {
	return instance.SetProperty("RemoteMachineTunnelAuthorizationList", value)
}

// GetRemoteMachineTunnelAuthorizationList gets the value of RemoteMachineTunnelAuthorizationList for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyRemoteMachineTunnelAuthorizationList() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteMachineTunnelAuthorizationList")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteUserTransportAuthorizationList sets the value of RemoteUserTransportAuthorizationList for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyRemoteUserTransportAuthorizationList(value string) (err error) {
	return instance.SetProperty("RemoteUserTransportAuthorizationList", value)
}

// GetRemoteUserTransportAuthorizationList gets the value of RemoteUserTransportAuthorizationList for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyRemoteUserTransportAuthorizationList() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteUserTransportAuthorizationList")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteUserTunnelAuthorizationList sets the value of RemoteUserTunnelAuthorizationList for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyRemoteUserTunnelAuthorizationList(value string) (err error) {
	return instance.SetProperty("RemoteUserTunnelAuthorizationList", value)
}

// GetRemoteUserTunnelAuthorizationList gets the value of RemoteUserTunnelAuthorizationList for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyRemoteUserTunnelAuthorizationList() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteUserTunnelAuthorizationList")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequireFullAuthSupport sets the value of RequireFullAuthSupport for the instance
func (instance *MSFT_NetSecuritySettingData) SetPropertyRequireFullAuthSupport(value uint16) (err error) {
	return instance.SetProperty("RequireFullAuthSupport", value)
}

// GetRequireFullAuthSupport gets the value of RequireFullAuthSupport for the instance
func (instance *MSFT_NetSecuritySettingData) GetPropertyRequireFullAuthSupport() (value uint16, err error) {
	retValue, err := instance.GetProperty("RequireFullAuthSupport")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
