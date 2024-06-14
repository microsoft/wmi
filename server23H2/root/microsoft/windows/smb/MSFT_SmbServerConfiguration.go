// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_SmbServerConfiguration struct
type MSFT_SmbServerConfiguration struct {
	*cim.WmiInstance

	//
	AnnounceComment string

	//
	AnnounceServer bool

	//
	AsynchronousCredits uint32

	//
	AuditClientCertificateAccess bool

	//
	AuditSmb1Access bool

	//
	AutoDisconnectTimeoutInMinutesV1 uint32

	//
	AutoDisconnectTimeoutInSecondsV2 uint32

	//
	AutoShareServer bool

	//
	AutoShareWorkstation bool

	//
	CachedOpenLimit uint32

	//
	DisableCompression bool

	//
	DisableSmbEncryptionOnSecureConnection bool

	//
	DurableHandleV2TimeoutInSeconds uint32

	//
	EnableAuthenticateUserSharing bool

	//
	EnableDirectoryHandleLeasing bool

	//
	EnableDownlevelTimewarp bool

	//
	EnableForcedLogoff bool

	//
	EnableLeasing bool

	//
	EnableMultiChannel bool

	//
	EnableOplocks bool

	//
	EnableSecuritySignature bool

	//
	EnableSMB1Protocol bool

	//
	EnableSMB2Protocol bool

	//
	EnableSMBQUIC bool

	//
	EnableStrictNameChecking bool

	//
	EncryptData bool

	//
	EncryptionCiphers string

	//
	InvalidAuthenticationDelayTimeInMs uint32

	//
	IrpStackSize uint32

	//
	KeepAliveTime uint32

	//
	MaxChannelPerSession uint32

	//
	MaxMpxCount uint32

	//
	MaxSessionPerConnection uint32

	//
	MaxThreadsPerQueue uint32

	//
	MaxWorkItems uint32

	//
	NullSessionPipes string

	//
	NullSessionShares string

	//
	OplockBreakWait uint32

	//
	PendingClientTimeoutInSeconds uint32

	//
	RejectUnencryptedAccess bool

	//
	RequestCompression bool

	//
	RequireSecuritySignature bool

	//
	RestrictNamedpipeAccessViaQuic bool

	//
	ServerHidden bool

	//
	Smb2CreditsMax uint32

	//
	Smb2CreditsMin uint32

	//
	SmbServerNameHardeningLevel uint32

	//
	TreatHostAsStableStorage bool

	//
	ValidateAliasNotCircular bool

	//
	ValidateShareScope bool

	//
	ValidateShareScopeNotAliased bool

	//
	ValidateTargetName bool
}

func NewMSFT_SmbServerConfigurationEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbServerConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbServerConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_SmbServerConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SmbServerConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbServerConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetAnnounceComment sets the value of AnnounceComment for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyAnnounceComment(value string) (err error) {
	return instance.SetProperty("AnnounceComment", (value))
}

// GetAnnounceComment gets the value of AnnounceComment for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyAnnounceComment() (value string, err error) {
	retValue, err := instance.GetProperty("AnnounceComment")
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

// SetAnnounceServer sets the value of AnnounceServer for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyAnnounceServer(value bool) (err error) {
	return instance.SetProperty("AnnounceServer", (value))
}

// GetAnnounceServer gets the value of AnnounceServer for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyAnnounceServer() (value bool, err error) {
	retValue, err := instance.GetProperty("AnnounceServer")
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

// SetAsynchronousCredits sets the value of AsynchronousCredits for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyAsynchronousCredits(value uint32) (err error) {
	return instance.SetProperty("AsynchronousCredits", (value))
}

// GetAsynchronousCredits gets the value of AsynchronousCredits for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyAsynchronousCredits() (value uint32, err error) {
	retValue, err := instance.GetProperty("AsynchronousCredits")
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

// SetAuditClientCertificateAccess sets the value of AuditClientCertificateAccess for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyAuditClientCertificateAccess(value bool) (err error) {
	return instance.SetProperty("AuditClientCertificateAccess", (value))
}

// GetAuditClientCertificateAccess gets the value of AuditClientCertificateAccess for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyAuditClientCertificateAccess() (value bool, err error) {
	retValue, err := instance.GetProperty("AuditClientCertificateAccess")
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

// SetAuditSmb1Access sets the value of AuditSmb1Access for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyAuditSmb1Access(value bool) (err error) {
	return instance.SetProperty("AuditSmb1Access", (value))
}

// GetAuditSmb1Access gets the value of AuditSmb1Access for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyAuditSmb1Access() (value bool, err error) {
	retValue, err := instance.GetProperty("AuditSmb1Access")
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

// SetAutoDisconnectTimeoutInMinutesV1 sets the value of AutoDisconnectTimeoutInMinutesV1 for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyAutoDisconnectTimeoutInMinutesV1(value uint32) (err error) {
	return instance.SetProperty("AutoDisconnectTimeoutInMinutesV1", (value))
}

// GetAutoDisconnectTimeoutInMinutesV1 gets the value of AutoDisconnectTimeoutInMinutesV1 for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyAutoDisconnectTimeoutInMinutesV1() (value uint32, err error) {
	retValue, err := instance.GetProperty("AutoDisconnectTimeoutInMinutesV1")
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

// SetAutoDisconnectTimeoutInSecondsV2 sets the value of AutoDisconnectTimeoutInSecondsV2 for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyAutoDisconnectTimeoutInSecondsV2(value uint32) (err error) {
	return instance.SetProperty("AutoDisconnectTimeoutInSecondsV2", (value))
}

// GetAutoDisconnectTimeoutInSecondsV2 gets the value of AutoDisconnectTimeoutInSecondsV2 for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyAutoDisconnectTimeoutInSecondsV2() (value uint32, err error) {
	retValue, err := instance.GetProperty("AutoDisconnectTimeoutInSecondsV2")
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

// SetAutoShareServer sets the value of AutoShareServer for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyAutoShareServer(value bool) (err error) {
	return instance.SetProperty("AutoShareServer", (value))
}

// GetAutoShareServer gets the value of AutoShareServer for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyAutoShareServer() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoShareServer")
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

// SetAutoShareWorkstation sets the value of AutoShareWorkstation for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyAutoShareWorkstation(value bool) (err error) {
	return instance.SetProperty("AutoShareWorkstation", (value))
}

// GetAutoShareWorkstation gets the value of AutoShareWorkstation for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyAutoShareWorkstation() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoShareWorkstation")
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

// SetCachedOpenLimit sets the value of CachedOpenLimit for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyCachedOpenLimit(value uint32) (err error) {
	return instance.SetProperty("CachedOpenLimit", (value))
}

// GetCachedOpenLimit gets the value of CachedOpenLimit for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyCachedOpenLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("CachedOpenLimit")
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

// SetDisableCompression sets the value of DisableCompression for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyDisableCompression(value bool) (err error) {
	return instance.SetProperty("DisableCompression", (value))
}

// GetDisableCompression gets the value of DisableCompression for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyDisableCompression() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableCompression")
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

// SetDisableSmbEncryptionOnSecureConnection sets the value of DisableSmbEncryptionOnSecureConnection for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyDisableSmbEncryptionOnSecureConnection(value bool) (err error) {
	return instance.SetProperty("DisableSmbEncryptionOnSecureConnection", (value))
}

// GetDisableSmbEncryptionOnSecureConnection gets the value of DisableSmbEncryptionOnSecureConnection for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyDisableSmbEncryptionOnSecureConnection() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableSmbEncryptionOnSecureConnection")
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

// SetDurableHandleV2TimeoutInSeconds sets the value of DurableHandleV2TimeoutInSeconds for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyDurableHandleV2TimeoutInSeconds(value uint32) (err error) {
	return instance.SetProperty("DurableHandleV2TimeoutInSeconds", (value))
}

// GetDurableHandleV2TimeoutInSeconds gets the value of DurableHandleV2TimeoutInSeconds for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyDurableHandleV2TimeoutInSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("DurableHandleV2TimeoutInSeconds")
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

// SetEnableAuthenticateUserSharing sets the value of EnableAuthenticateUserSharing for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableAuthenticateUserSharing(value bool) (err error) {
	return instance.SetProperty("EnableAuthenticateUserSharing", (value))
}

// GetEnableAuthenticateUserSharing gets the value of EnableAuthenticateUserSharing for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableAuthenticateUserSharing() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableAuthenticateUserSharing")
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

// SetEnableDirectoryHandleLeasing sets the value of EnableDirectoryHandleLeasing for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableDirectoryHandleLeasing(value bool) (err error) {
	return instance.SetProperty("EnableDirectoryHandleLeasing", (value))
}

// GetEnableDirectoryHandleLeasing gets the value of EnableDirectoryHandleLeasing for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableDirectoryHandleLeasing() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableDirectoryHandleLeasing")
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

// SetEnableDownlevelTimewarp sets the value of EnableDownlevelTimewarp for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableDownlevelTimewarp(value bool) (err error) {
	return instance.SetProperty("EnableDownlevelTimewarp", (value))
}

// GetEnableDownlevelTimewarp gets the value of EnableDownlevelTimewarp for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableDownlevelTimewarp() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableDownlevelTimewarp")
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

// SetEnableForcedLogoff sets the value of EnableForcedLogoff for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableForcedLogoff(value bool) (err error) {
	return instance.SetProperty("EnableForcedLogoff", (value))
}

// GetEnableForcedLogoff gets the value of EnableForcedLogoff for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableForcedLogoff() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableForcedLogoff")
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

// SetEnableLeasing sets the value of EnableLeasing for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableLeasing(value bool) (err error) {
	return instance.SetProperty("EnableLeasing", (value))
}

// GetEnableLeasing gets the value of EnableLeasing for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableLeasing() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableLeasing")
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

// SetEnableMultiChannel sets the value of EnableMultiChannel for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableMultiChannel(value bool) (err error) {
	return instance.SetProperty("EnableMultiChannel", (value))
}

// GetEnableMultiChannel gets the value of EnableMultiChannel for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableMultiChannel() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableMultiChannel")
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

// SetEnableOplocks sets the value of EnableOplocks for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableOplocks(value bool) (err error) {
	return instance.SetProperty("EnableOplocks", (value))
}

// GetEnableOplocks gets the value of EnableOplocks for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableOplocks() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableOplocks")
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

// SetEnableSecuritySignature sets the value of EnableSecuritySignature for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableSecuritySignature(value bool) (err error) {
	return instance.SetProperty("EnableSecuritySignature", (value))
}

// GetEnableSecuritySignature gets the value of EnableSecuritySignature for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableSecuritySignature() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableSecuritySignature")
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

// SetEnableSMB1Protocol sets the value of EnableSMB1Protocol for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableSMB1Protocol(value bool) (err error) {
	return instance.SetProperty("EnableSMB1Protocol", (value))
}

// GetEnableSMB1Protocol gets the value of EnableSMB1Protocol for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableSMB1Protocol() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableSMB1Protocol")
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

// SetEnableSMB2Protocol sets the value of EnableSMB2Protocol for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableSMB2Protocol(value bool) (err error) {
	return instance.SetProperty("EnableSMB2Protocol", (value))
}

// GetEnableSMB2Protocol gets the value of EnableSMB2Protocol for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableSMB2Protocol() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableSMB2Protocol")
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

// SetEnableSMBQUIC sets the value of EnableSMBQUIC for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableSMBQUIC(value bool) (err error) {
	return instance.SetProperty("EnableSMBQUIC", (value))
}

// GetEnableSMBQUIC gets the value of EnableSMBQUIC for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableSMBQUIC() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableSMBQUIC")
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

// SetEnableStrictNameChecking sets the value of EnableStrictNameChecking for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableStrictNameChecking(value bool) (err error) {
	return instance.SetProperty("EnableStrictNameChecking", (value))
}

// GetEnableStrictNameChecking gets the value of EnableStrictNameChecking for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableStrictNameChecking() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableStrictNameChecking")
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

// SetEncryptData sets the value of EncryptData for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEncryptData(value bool) (err error) {
	return instance.SetProperty("EncryptData", (value))
}

// GetEncryptData gets the value of EncryptData for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEncryptData() (value bool, err error) {
	retValue, err := instance.GetProperty("EncryptData")
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

// SetEncryptionCiphers sets the value of EncryptionCiphers for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEncryptionCiphers(value string) (err error) {
	return instance.SetProperty("EncryptionCiphers", (value))
}

// GetEncryptionCiphers gets the value of EncryptionCiphers for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEncryptionCiphers() (value string, err error) {
	retValue, err := instance.GetProperty("EncryptionCiphers")
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

// SetInvalidAuthenticationDelayTimeInMs sets the value of InvalidAuthenticationDelayTimeInMs for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyInvalidAuthenticationDelayTimeInMs(value uint32) (err error) {
	return instance.SetProperty("InvalidAuthenticationDelayTimeInMs", (value))
}

// GetInvalidAuthenticationDelayTimeInMs gets the value of InvalidAuthenticationDelayTimeInMs for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyInvalidAuthenticationDelayTimeInMs() (value uint32, err error) {
	retValue, err := instance.GetProperty("InvalidAuthenticationDelayTimeInMs")
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

// SetIrpStackSize sets the value of IrpStackSize for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyIrpStackSize(value uint32) (err error) {
	return instance.SetProperty("IrpStackSize", (value))
}

// GetIrpStackSize gets the value of IrpStackSize for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyIrpStackSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("IrpStackSize")
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

// SetKeepAliveTime sets the value of KeepAliveTime for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyKeepAliveTime(value uint32) (err error) {
	return instance.SetProperty("KeepAliveTime", (value))
}

// GetKeepAliveTime gets the value of KeepAliveTime for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyKeepAliveTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("KeepAliveTime")
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

// SetMaxChannelPerSession sets the value of MaxChannelPerSession for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyMaxChannelPerSession(value uint32) (err error) {
	return instance.SetProperty("MaxChannelPerSession", (value))
}

// GetMaxChannelPerSession gets the value of MaxChannelPerSession for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyMaxChannelPerSession() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxChannelPerSession")
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

// SetMaxMpxCount sets the value of MaxMpxCount for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyMaxMpxCount(value uint32) (err error) {
	return instance.SetProperty("MaxMpxCount", (value))
}

// GetMaxMpxCount gets the value of MaxMpxCount for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyMaxMpxCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxMpxCount")
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

// SetMaxSessionPerConnection sets the value of MaxSessionPerConnection for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyMaxSessionPerConnection(value uint32) (err error) {
	return instance.SetProperty("MaxSessionPerConnection", (value))
}

// GetMaxSessionPerConnection gets the value of MaxSessionPerConnection for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyMaxSessionPerConnection() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxSessionPerConnection")
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

// SetMaxThreadsPerQueue sets the value of MaxThreadsPerQueue for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyMaxThreadsPerQueue(value uint32) (err error) {
	return instance.SetProperty("MaxThreadsPerQueue", (value))
}

// GetMaxThreadsPerQueue gets the value of MaxThreadsPerQueue for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyMaxThreadsPerQueue() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxThreadsPerQueue")
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

// SetMaxWorkItems sets the value of MaxWorkItems for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyMaxWorkItems(value uint32) (err error) {
	return instance.SetProperty("MaxWorkItems", (value))
}

// GetMaxWorkItems gets the value of MaxWorkItems for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyMaxWorkItems() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxWorkItems")
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

// SetNullSessionPipes sets the value of NullSessionPipes for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyNullSessionPipes(value string) (err error) {
	return instance.SetProperty("NullSessionPipes", (value))
}

// GetNullSessionPipes gets the value of NullSessionPipes for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyNullSessionPipes() (value string, err error) {
	retValue, err := instance.GetProperty("NullSessionPipes")
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

// SetNullSessionShares sets the value of NullSessionShares for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyNullSessionShares(value string) (err error) {
	return instance.SetProperty("NullSessionShares", (value))
}

// GetNullSessionShares gets the value of NullSessionShares for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyNullSessionShares() (value string, err error) {
	retValue, err := instance.GetProperty("NullSessionShares")
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

// SetOplockBreakWait sets the value of OplockBreakWait for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyOplockBreakWait(value uint32) (err error) {
	return instance.SetProperty("OplockBreakWait", (value))
}

// GetOplockBreakWait gets the value of OplockBreakWait for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyOplockBreakWait() (value uint32, err error) {
	retValue, err := instance.GetProperty("OplockBreakWait")
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

// SetPendingClientTimeoutInSeconds sets the value of PendingClientTimeoutInSeconds for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyPendingClientTimeoutInSeconds(value uint32) (err error) {
	return instance.SetProperty("PendingClientTimeoutInSeconds", (value))
}

// GetPendingClientTimeoutInSeconds gets the value of PendingClientTimeoutInSeconds for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyPendingClientTimeoutInSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("PendingClientTimeoutInSeconds")
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

// SetRejectUnencryptedAccess sets the value of RejectUnencryptedAccess for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyRejectUnencryptedAccess(value bool) (err error) {
	return instance.SetProperty("RejectUnencryptedAccess", (value))
}

// GetRejectUnencryptedAccess gets the value of RejectUnencryptedAccess for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyRejectUnencryptedAccess() (value bool, err error) {
	retValue, err := instance.GetProperty("RejectUnencryptedAccess")
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

// SetRequestCompression sets the value of RequestCompression for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyRequestCompression(value bool) (err error) {
	return instance.SetProperty("RequestCompression", (value))
}

// GetRequestCompression gets the value of RequestCompression for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyRequestCompression() (value bool, err error) {
	retValue, err := instance.GetProperty("RequestCompression")
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

// SetRequireSecuritySignature sets the value of RequireSecuritySignature for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyRequireSecuritySignature(value bool) (err error) {
	return instance.SetProperty("RequireSecuritySignature", (value))
}

// GetRequireSecuritySignature gets the value of RequireSecuritySignature for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyRequireSecuritySignature() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireSecuritySignature")
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

// SetRestrictNamedpipeAccessViaQuic sets the value of RestrictNamedpipeAccessViaQuic for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyRestrictNamedpipeAccessViaQuic(value bool) (err error) {
	return instance.SetProperty("RestrictNamedpipeAccessViaQuic", (value))
}

// GetRestrictNamedpipeAccessViaQuic gets the value of RestrictNamedpipeAccessViaQuic for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyRestrictNamedpipeAccessViaQuic() (value bool, err error) {
	retValue, err := instance.GetProperty("RestrictNamedpipeAccessViaQuic")
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

// SetServerHidden sets the value of ServerHidden for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyServerHidden(value bool) (err error) {
	return instance.SetProperty("ServerHidden", (value))
}

// GetServerHidden gets the value of ServerHidden for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyServerHidden() (value bool, err error) {
	retValue, err := instance.GetProperty("ServerHidden")
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

// SetSmb2CreditsMax sets the value of Smb2CreditsMax for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertySmb2CreditsMax(value uint32) (err error) {
	return instance.SetProperty("Smb2CreditsMax", (value))
}

// GetSmb2CreditsMax gets the value of Smb2CreditsMax for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertySmb2CreditsMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("Smb2CreditsMax")
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

// SetSmb2CreditsMin sets the value of Smb2CreditsMin for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertySmb2CreditsMin(value uint32) (err error) {
	return instance.SetProperty("Smb2CreditsMin", (value))
}

// GetSmb2CreditsMin gets the value of Smb2CreditsMin for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertySmb2CreditsMin() (value uint32, err error) {
	retValue, err := instance.GetProperty("Smb2CreditsMin")
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

// SetSmbServerNameHardeningLevel sets the value of SmbServerNameHardeningLevel for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertySmbServerNameHardeningLevel(value uint32) (err error) {
	return instance.SetProperty("SmbServerNameHardeningLevel", (value))
}

// GetSmbServerNameHardeningLevel gets the value of SmbServerNameHardeningLevel for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertySmbServerNameHardeningLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("SmbServerNameHardeningLevel")
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

// SetTreatHostAsStableStorage sets the value of TreatHostAsStableStorage for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyTreatHostAsStableStorage(value bool) (err error) {
	return instance.SetProperty("TreatHostAsStableStorage", (value))
}

// GetTreatHostAsStableStorage gets the value of TreatHostAsStableStorage for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyTreatHostAsStableStorage() (value bool, err error) {
	retValue, err := instance.GetProperty("TreatHostAsStableStorage")
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

// SetValidateAliasNotCircular sets the value of ValidateAliasNotCircular for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyValidateAliasNotCircular(value bool) (err error) {
	return instance.SetProperty("ValidateAliasNotCircular", (value))
}

// GetValidateAliasNotCircular gets the value of ValidateAliasNotCircular for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyValidateAliasNotCircular() (value bool, err error) {
	retValue, err := instance.GetProperty("ValidateAliasNotCircular")
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

// SetValidateShareScope sets the value of ValidateShareScope for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyValidateShareScope(value bool) (err error) {
	return instance.SetProperty("ValidateShareScope", (value))
}

// GetValidateShareScope gets the value of ValidateShareScope for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyValidateShareScope() (value bool, err error) {
	retValue, err := instance.GetProperty("ValidateShareScope")
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

// SetValidateShareScopeNotAliased sets the value of ValidateShareScopeNotAliased for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyValidateShareScopeNotAliased(value bool) (err error) {
	return instance.SetProperty("ValidateShareScopeNotAliased", (value))
}

// GetValidateShareScopeNotAliased gets the value of ValidateShareScopeNotAliased for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyValidateShareScopeNotAliased() (value bool, err error) {
	retValue, err := instance.GetProperty("ValidateShareScopeNotAliased")
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

// SetValidateTargetName sets the value of ValidateTargetName for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyValidateTargetName(value bool) (err error) {
	return instance.SetProperty("ValidateTargetName", (value))
}

// GetValidateTargetName gets the value of ValidateTargetName for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyValidateTargetName() (value bool, err error) {
	retValue, err := instance.GetProperty("ValidateTargetName")
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

//

// <param name="Output" type="MSFT_SmbServerConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbServerConfiguration) GetConfiguration( /* OUT */ Output MSFT_SmbServerConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetConfiguration")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="All" type="bool "></param>
// <param name="AnnounceComment" type="bool "></param>
// <param name="AnnounceServer" type="bool "></param>
// <param name="AsynchronousCredits" type="bool "></param>
// <param name="AuditClientCertificateAccess" type="bool "></param>
// <param name="AuditSmb1Access" type="bool "></param>
// <param name="AutoShareServer" type="bool "></param>
// <param name="AutoShareWorkstation" type="bool "></param>
// <param name="CachedOpenLimit" type="bool "></param>
// <param name="DisableCompression" type="bool "></param>
// <param name="DisableSmbEncryptionOnSecureConnection" type="bool "></param>
// <param name="DurableHandleV2TimeoutInSeconds" type="bool "></param>
// <param name="EnableDirectoryHandleLeasing" type="bool "></param>
// <param name="EnableDownlevelTimewarp" type="bool "></param>
// <param name="EnableLeasing" type="bool "></param>
// <param name="EnableMultiChannel" type="bool "></param>
// <param name="EnableOplocks" type="bool "></param>
// <param name="EnableSMB2Protocol" type="bool "></param>
// <param name="EnableSMBQUIC" type="bool "></param>
// <param name="EnableStrictNameChecking" type="bool "></param>
// <param name="EncryptData" type="bool "></param>
// <param name="EncryptionCiphers" type="bool "></param>
// <param name="IrpStackSize" type="bool "></param>
// <param name="KeepAliveTime" type="bool "></param>
// <param name="MaxChannelPerSession" type="bool "></param>
// <param name="MaxMpxCount" type="bool "></param>
// <param name="MaxSessionPerConnection" type="bool "></param>
// <param name="MaxThreadsPerQueue" type="bool "></param>
// <param name="MaxWorkItems" type="bool "></param>
// <param name="NullSessionShares" type="bool "></param>
// <param name="OplockBreakWait" type="bool "></param>
// <param name="PendingClientTimeoutInSeconds" type="bool "></param>
// <param name="RejectUnencryptedAccess" type="bool "></param>
// <param name="RequestCompression" type="bool "></param>
// <param name="RestrictNamedpipeAccessViaQuic" type="bool "></param>
// <param name="ServerHidden" type="bool "></param>
// <param name="Smb2CreditsMax" type="bool "></param>
// <param name="Smb2CreditsMin" type="bool "></param>
// <param name="SmbServerNameHardeningLevel" type="bool "></param>
// <param name="TreatHostAsStableStorage" type="bool "></param>
// <param name="ValidateAliasNotCircular" type="bool "></param>
// <param name="ValidateShareScope" type="bool "></param>
// <param name="ValidateShareScopeNotAliased" type="bool "></param>
// <param name="ValidateTargetName" type="bool "></param>

// <param name="Output" type="MSFT_SmbServerConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbServerConfiguration) ResetConfiguration( /* OUT */ Output MSFT_SmbServerConfiguration,
	/* OPTIONAL IN */ All bool,
	/* OPTIONAL IN */ AnnounceServer bool,
	/* OPTIONAL IN */ AsynchronousCredits bool,
	/* OPTIONAL IN */ AutoShareServer bool,
	/* OPTIONAL IN */ AutoShareWorkstation bool,
	/* OPTIONAL IN */ CachedOpenLimit bool,
	/* OPTIONAL IN */ AnnounceComment bool,
	/* OPTIONAL IN */ EnableDownlevelTimewarp bool,
	/* OPTIONAL IN */ EnableLeasing bool,
	/* OPTIONAL IN */ EnableMultiChannel bool,
	/* OPTIONAL IN */ EnableStrictNameChecking bool,
	/* OPTIONAL IN */ DisableCompression bool,
	/* OPTIONAL IN */ DurableHandleV2TimeoutInSeconds bool,
	/* OPTIONAL IN */ EnableOplocks bool,
	/* OPTIONAL IN */ ServerHidden bool,
	/* OPTIONAL IN */ IrpStackSize bool,
	/* OPTIONAL IN */ KeepAliveTime bool,
	/* OPTIONAL IN */ MaxChannelPerSession bool,
	/* OPTIONAL IN */ MaxMpxCount bool,
	/* OPTIONAL IN */ MaxSessionPerConnection bool,
	/* OPTIONAL IN */ MaxThreadsPerQueue bool,
	/* OPTIONAL IN */ MaxWorkItems bool,
	/* OPTIONAL IN */ NullSessionShares bool,
	/* OPTIONAL IN */ OplockBreakWait bool,
	/* OPTIONAL IN */ PendingClientTimeoutInSeconds bool,
	/* OPTIONAL IN */ EnableSMB2Protocol bool,
	/* OPTIONAL IN */ Smb2CreditsMax bool,
	/* OPTIONAL IN */ Smb2CreditsMin bool,
	/* OPTIONAL IN */ SmbServerNameHardeningLevel bool,
	/* OPTIONAL IN */ TreatHostAsStableStorage bool,
	/* OPTIONAL IN */ ValidateAliasNotCircular bool,
	/* OPTIONAL IN */ ValidateShareScope bool,
	/* OPTIONAL IN */ ValidateShareScopeNotAliased bool,
	/* OPTIONAL IN */ ValidateTargetName bool,
	/* OPTIONAL IN */ EncryptData bool,
	/* OPTIONAL IN */ RejectUnencryptedAccess bool,
	/* OPTIONAL IN */ AuditSmb1Access bool,
	/* OPTIONAL IN */ DisableSmbEncryptionOnSecureConnection bool,
	/* OPTIONAL IN */ RequestCompression bool,
	/* OPTIONAL IN */ RestrictNamedpipeAccessViaQuic bool,
	/* OPTIONAL IN */ EnableSMBQUIC bool,
	/* OPTIONAL IN */ EnableDirectoryHandleLeasing bool,
	/* OPTIONAL IN */ EncryptionCiphers bool,
	/* OPTIONAL IN */ AuditClientCertificateAccess bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ResetConfiguration", All, AnnounceServer, AsynchronousCredits, AutoShareServer, AutoShareWorkstation, CachedOpenLimit, AnnounceComment, EnableDownlevelTimewarp, EnableLeasing, EnableMultiChannel, EnableStrictNameChecking, DisableCompression, DurableHandleV2TimeoutInSeconds, EnableOplocks, ServerHidden, IrpStackSize, KeepAliveTime, MaxChannelPerSession, MaxMpxCount, MaxSessionPerConnection, MaxThreadsPerQueue, MaxWorkItems, NullSessionShares, OplockBreakWait, PendingClientTimeoutInSeconds, EnableSMB2Protocol, Smb2CreditsMax, Smb2CreditsMin, SmbServerNameHardeningLevel, TreatHostAsStableStorage, ValidateAliasNotCircular, ValidateShareScope, ValidateShareScopeNotAliased, ValidateTargetName, EncryptData, RejectUnencryptedAccess, AuditSmb1Access, DisableSmbEncryptionOnSecureConnection, RequestCompression, RestrictNamedpipeAccessViaQuic, EnableSMBQUIC, EnableDirectoryHandleLeasing, EncryptionCiphers, AuditClientCertificateAccess)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AnnounceComment" type="string "></param>
// <param name="AnnounceServer" type="bool "></param>
// <param name="AsynchronousCredits" type="uint32 "></param>
// <param name="AuditClientCertificateAccess" type="bool "></param>
// <param name="AuditSmb1Access" type="bool "></param>
// <param name="AutoDisconnectTimeoutInMinutesV1" type="uint32 "></param>
// <param name="AutoDisconnectTimeoutInSecondsV2" type="uint32 "></param>
// <param name="AutoShareServer" type="bool "></param>
// <param name="AutoShareWorkstation" type="bool "></param>
// <param name="CachedOpenLimit" type="uint32 "></param>
// <param name="DisableCompression" type="bool "></param>
// <param name="DisableSmbEncryptionOnSecureConnection" type="bool "></param>
// <param name="DurableHandleV2TimeoutInSeconds" type="uint32 "></param>
// <param name="EnableAuthenticateUserSharing" type="bool "></param>
// <param name="EnableDirectoryHandleLeasing" type="bool "></param>
// <param name="EnableDownlevelTimewarp" type="bool "></param>
// <param name="EnableForcedLogoff" type="bool "></param>
// <param name="EnableLeasing" type="bool "></param>
// <param name="EnableMultiChannel" type="bool "></param>
// <param name="EnableOplocks" type="bool "></param>
// <param name="EnableSecuritySignature" type="bool "></param>
// <param name="EnableSMB1Protocol" type="bool "></param>
// <param name="EnableSMB2Protocol" type="bool "></param>
// <param name="EnableSMBQUIC" type="bool "></param>
// <param name="EnableStrictNameChecking" type="bool "></param>
// <param name="EncryptData" type="bool "></param>
// <param name="EncryptionCiphers" type="string "></param>
// <param name="InvalidAuthenticationDelayTimeInMs" type="uint32 "></param>
// <param name="IrpStackSize" type="uint32 "></param>
// <param name="KeepAliveTime" type="uint32 "></param>
// <param name="MaxChannelPerSession" type="uint32 "></param>
// <param name="MaxMpxCount" type="uint32 "></param>
// <param name="MaxSessionPerConnection" type="uint32 "></param>
// <param name="MaxThreadsPerQueue" type="uint32 "></param>
// <param name="MaxWorkItems" type="uint32 "></param>
// <param name="NullSessionPipes" type="string "></param>
// <param name="NullSessionShares" type="string "></param>
// <param name="OplockBreakWait" type="uint32 "></param>
// <param name="PendingClientTimeoutInSeconds" type="uint32 "></param>
// <param name="RejectUnencryptedAccess" type="bool "></param>
// <param name="RequestCompression" type="bool "></param>
// <param name="RequireSecuritySignature" type="bool "></param>
// <param name="RestrictNamedpipeAccessViaQuic" type="bool "></param>
// <param name="ServerHidden" type="bool "></param>
// <param name="Smb2CreditsMax" type="uint32 "></param>
// <param name="Smb2CreditsMin" type="uint32 "></param>
// <param name="SmbServerNameHardeningLevel" type="uint32 "></param>
// <param name="TreatHostAsStableStorage" type="bool "></param>
// <param name="ValidateAliasNotCircular" type="bool "></param>
// <param name="ValidateShareScope" type="bool "></param>
// <param name="ValidateShareScopeNotAliased" type="bool "></param>
// <param name="ValidateTargetName" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbServerConfiguration) SetConfiguration( /* IN */ AnnounceServer bool,
	/* IN */ AsynchronousCredits uint32,
	/* IN */ AutoShareServer bool,
	/* IN */ AutoShareWorkstation bool,
	/* IN */ CachedOpenLimit uint32,
	/* IN */ AnnounceComment string,
	/* IN */ EnableDownlevelTimewarp bool,
	/* IN */ EnableLeasing bool,
	/* IN */ EnableMultiChannel bool,
	/* IN */ EnableStrictNameChecking bool,
	/* IN */ AutoDisconnectTimeoutInMinutesV1 uint32,
	/* IN */ AutoDisconnectTimeoutInSecondsV2 uint32,
	/* IN */ DisableCompression bool,
	/* IN */ DurableHandleV2TimeoutInSeconds uint32,
	/* IN */ EnableAuthenticateUserSharing bool,
	/* IN */ EnableForcedLogoff bool,
	/* IN */ EnableOplocks bool,
	/* IN */ EnableSecuritySignature bool,
	/* IN */ ServerHidden bool,
	/* IN */ IrpStackSize uint32,
	/* IN */ KeepAliveTime uint32,
	/* IN */ MaxChannelPerSession uint32,
	/* IN */ MaxMpxCount uint32,
	/* IN */ MaxSessionPerConnection uint32,
	/* IN */ MaxThreadsPerQueue uint32,
	/* IN */ MaxWorkItems uint32,
	/* IN */ NullSessionPipes string,
	/* IN */ NullSessionShares string,
	/* IN */ OplockBreakWait uint32,
	/* IN */ PendingClientTimeoutInSeconds uint32,
	/* IN */ RequestCompression bool,
	/* IN */ RequireSecuritySignature bool,
	/* IN */ EnableSMB1Protocol bool,
	/* IN */ EnableSMB2Protocol bool,
	/* IN */ Smb2CreditsMax uint32,
	/* IN */ Smb2CreditsMin uint32,
	/* IN */ SmbServerNameHardeningLevel uint32,
	/* IN */ TreatHostAsStableStorage bool,
	/* IN */ ValidateAliasNotCircular bool,
	/* IN */ ValidateShareScope bool,
	/* IN */ ValidateShareScopeNotAliased bool,
	/* IN */ ValidateTargetName bool,
	/* IN */ EncryptData bool,
	/* IN */ RejectUnencryptedAccess bool,
	/* IN */ AuditSmb1Access bool,
	/* IN */ DisableSmbEncryptionOnSecureConnection bool,
	/* IN */ RestrictNamedpipeAccessViaQuic bool,
	/* IN */ EnableSMBQUIC bool,
	/* IN */ EnableDirectoryHandleLeasing bool,
	/* IN */ EncryptionCiphers string,
	/* IN */ InvalidAuthenticationDelayTimeInMs uint32,
	/* IN */ AuditClientCertificateAccess bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetConfiguration", AnnounceServer, AsynchronousCredits, AutoShareServer, AutoShareWorkstation, CachedOpenLimit, AnnounceComment, EnableDownlevelTimewarp, EnableLeasing, EnableMultiChannel, EnableStrictNameChecking, AutoDisconnectTimeoutInMinutesV1, AutoDisconnectTimeoutInSecondsV2, DisableCompression, DurableHandleV2TimeoutInSeconds, EnableAuthenticateUserSharing, EnableForcedLogoff, EnableOplocks, EnableSecuritySignature, ServerHidden, IrpStackSize, KeepAliveTime, MaxChannelPerSession, MaxMpxCount, MaxSessionPerConnection, MaxThreadsPerQueue, MaxWorkItems, NullSessionPipes, NullSessionShares, OplockBreakWait, PendingClientTimeoutInSeconds, RequestCompression, RequireSecuritySignature, EnableSMB1Protocol, EnableSMB2Protocol, Smb2CreditsMax, Smb2CreditsMin, SmbServerNameHardeningLevel, TreatHostAsStableStorage, ValidateAliasNotCircular, ValidateShareScope, ValidateShareScopeNotAliased, ValidateTargetName, EncryptData, RejectUnencryptedAccess, AuditSmb1Access, DisableSmbEncryptionOnSecureConnection, RestrictNamedpipeAccessViaQuic, EnableSMBQUIC, EnableDirectoryHandleLeasing, EncryptionCiphers, InvalidAuthenticationDelayTimeInMs, AuditClientCertificateAccess)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
