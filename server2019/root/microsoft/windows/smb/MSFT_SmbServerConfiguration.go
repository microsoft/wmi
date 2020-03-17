// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_SmbServerConfiguration struct
type MSFT_SmbServerConfiguration struct {
	cim.WmiInstance

	//
	AnnounceComment string

	//
	AnnounceServer bool

	//
	AsynchronousCredits uint32

	//
	AuditSmb1Access bool

	//
	AutoDisconnectTimeout uint32

	//
	AutoShareServer bool

	//
	AutoShareWorkstation bool

	//
	CachedOpenLimit uint32

	//
	DurableHandleV2TimeoutInSeconds uint32

	//
	EnableAuthenticateUserSharing bool

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
	EnableStrictNameChecking bool

	//
	EncryptData bool

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
	RequireSecuritySignature bool

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

// SetAnnounceComment sets the value of AnnounceComment for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyAnnounceComment(value string) (err error) {
	return instance.SetProperty("AnnounceComment", value)
}

// GetAnnounceComment gets the value of AnnounceComment for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyAnnounceComment() (value string, err error) {
	retValue, err := instance.GetProperty("AnnounceComment")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAnnounceServer sets the value of AnnounceServer for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyAnnounceServer(value bool) (err error) {
	return instance.SetProperty("AnnounceServer", value)
}

// GetAnnounceServer gets the value of AnnounceServer for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyAnnounceServer() (value bool, err error) {
	retValue, err := instance.GetProperty("AnnounceServer")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAsynchronousCredits sets the value of AsynchronousCredits for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyAsynchronousCredits(value uint32) (err error) {
	return instance.SetProperty("AsynchronousCredits", value)
}

// GetAsynchronousCredits gets the value of AsynchronousCredits for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyAsynchronousCredits() (value uint32, err error) {
	retValue, err := instance.GetProperty("AsynchronousCredits")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAuditSmb1Access sets the value of AuditSmb1Access for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyAuditSmb1Access(value bool) (err error) {
	return instance.SetProperty("AuditSmb1Access", value)
}

// GetAuditSmb1Access gets the value of AuditSmb1Access for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyAuditSmb1Access() (value bool, err error) {
	retValue, err := instance.GetProperty("AuditSmb1Access")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutoDisconnectTimeout sets the value of AutoDisconnectTimeout for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyAutoDisconnectTimeout(value uint32) (err error) {
	return instance.SetProperty("AutoDisconnectTimeout", value)
}

// GetAutoDisconnectTimeout gets the value of AutoDisconnectTimeout for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyAutoDisconnectTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("AutoDisconnectTimeout")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutoShareServer sets the value of AutoShareServer for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyAutoShareServer(value bool) (err error) {
	return instance.SetProperty("AutoShareServer", value)
}

// GetAutoShareServer gets the value of AutoShareServer for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyAutoShareServer() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoShareServer")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutoShareWorkstation sets the value of AutoShareWorkstation for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyAutoShareWorkstation(value bool) (err error) {
	return instance.SetProperty("AutoShareWorkstation", value)
}

// GetAutoShareWorkstation gets the value of AutoShareWorkstation for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyAutoShareWorkstation() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoShareWorkstation")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCachedOpenLimit sets the value of CachedOpenLimit for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyCachedOpenLimit(value uint32) (err error) {
	return instance.SetProperty("CachedOpenLimit", value)
}

// GetCachedOpenLimit gets the value of CachedOpenLimit for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyCachedOpenLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("CachedOpenLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDurableHandleV2TimeoutInSeconds sets the value of DurableHandleV2TimeoutInSeconds for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyDurableHandleV2TimeoutInSeconds(value uint32) (err error) {
	return instance.SetProperty("DurableHandleV2TimeoutInSeconds", value)
}

// GetDurableHandleV2TimeoutInSeconds gets the value of DurableHandleV2TimeoutInSeconds for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyDurableHandleV2TimeoutInSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("DurableHandleV2TimeoutInSeconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableAuthenticateUserSharing sets the value of EnableAuthenticateUserSharing for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableAuthenticateUserSharing(value bool) (err error) {
	return instance.SetProperty("EnableAuthenticateUserSharing", value)
}

// GetEnableAuthenticateUserSharing gets the value of EnableAuthenticateUserSharing for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableAuthenticateUserSharing() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableAuthenticateUserSharing")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableDownlevelTimewarp sets the value of EnableDownlevelTimewarp for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableDownlevelTimewarp(value bool) (err error) {
	return instance.SetProperty("EnableDownlevelTimewarp", value)
}

// GetEnableDownlevelTimewarp gets the value of EnableDownlevelTimewarp for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableDownlevelTimewarp() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableDownlevelTimewarp")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableForcedLogoff sets the value of EnableForcedLogoff for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableForcedLogoff(value bool) (err error) {
	return instance.SetProperty("EnableForcedLogoff", value)
}

// GetEnableForcedLogoff gets the value of EnableForcedLogoff for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableForcedLogoff() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableForcedLogoff")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableLeasing sets the value of EnableLeasing for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableLeasing(value bool) (err error) {
	return instance.SetProperty("EnableLeasing", value)
}

// GetEnableLeasing gets the value of EnableLeasing for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableLeasing() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableLeasing")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableMultiChannel sets the value of EnableMultiChannel for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableMultiChannel(value bool) (err error) {
	return instance.SetProperty("EnableMultiChannel", value)
}

// GetEnableMultiChannel gets the value of EnableMultiChannel for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableMultiChannel() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableMultiChannel")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableOplocks sets the value of EnableOplocks for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableOplocks(value bool) (err error) {
	return instance.SetProperty("EnableOplocks", value)
}

// GetEnableOplocks gets the value of EnableOplocks for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableOplocks() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableOplocks")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableSecuritySignature sets the value of EnableSecuritySignature for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableSecuritySignature(value bool) (err error) {
	return instance.SetProperty("EnableSecuritySignature", value)
}

// GetEnableSecuritySignature gets the value of EnableSecuritySignature for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableSecuritySignature() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableSecuritySignature")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableSMB1Protocol sets the value of EnableSMB1Protocol for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableSMB1Protocol(value bool) (err error) {
	return instance.SetProperty("EnableSMB1Protocol", value)
}

// GetEnableSMB1Protocol gets the value of EnableSMB1Protocol for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableSMB1Protocol() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableSMB1Protocol")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableSMB2Protocol sets the value of EnableSMB2Protocol for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableSMB2Protocol(value bool) (err error) {
	return instance.SetProperty("EnableSMB2Protocol", value)
}

// GetEnableSMB2Protocol gets the value of EnableSMB2Protocol for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableSMB2Protocol() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableSMB2Protocol")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableStrictNameChecking sets the value of EnableStrictNameChecking for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEnableStrictNameChecking(value bool) (err error) {
	return instance.SetProperty("EnableStrictNameChecking", value)
}

// GetEnableStrictNameChecking gets the value of EnableStrictNameChecking for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEnableStrictNameChecking() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableStrictNameChecking")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEncryptData sets the value of EncryptData for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyEncryptData(value bool) (err error) {
	return instance.SetProperty("EncryptData", value)
}

// GetEncryptData gets the value of EncryptData for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyEncryptData() (value bool, err error) {
	retValue, err := instance.GetProperty("EncryptData")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIrpStackSize sets the value of IrpStackSize for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyIrpStackSize(value uint32) (err error) {
	return instance.SetProperty("IrpStackSize", value)
}

// GetIrpStackSize gets the value of IrpStackSize for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyIrpStackSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("IrpStackSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKeepAliveTime sets the value of KeepAliveTime for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyKeepAliveTime(value uint32) (err error) {
	return instance.SetProperty("KeepAliveTime", value)
}

// GetKeepAliveTime gets the value of KeepAliveTime for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyKeepAliveTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("KeepAliveTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxChannelPerSession sets the value of MaxChannelPerSession for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyMaxChannelPerSession(value uint32) (err error) {
	return instance.SetProperty("MaxChannelPerSession", value)
}

// GetMaxChannelPerSession gets the value of MaxChannelPerSession for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyMaxChannelPerSession() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxChannelPerSession")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxMpxCount sets the value of MaxMpxCount for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyMaxMpxCount(value uint32) (err error) {
	return instance.SetProperty("MaxMpxCount", value)
}

// GetMaxMpxCount gets the value of MaxMpxCount for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyMaxMpxCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxMpxCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxSessionPerConnection sets the value of MaxSessionPerConnection for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyMaxSessionPerConnection(value uint32) (err error) {
	return instance.SetProperty("MaxSessionPerConnection", value)
}

// GetMaxSessionPerConnection gets the value of MaxSessionPerConnection for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyMaxSessionPerConnection() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxSessionPerConnection")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxThreadsPerQueue sets the value of MaxThreadsPerQueue for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyMaxThreadsPerQueue(value uint32) (err error) {
	return instance.SetProperty("MaxThreadsPerQueue", value)
}

// GetMaxThreadsPerQueue gets the value of MaxThreadsPerQueue for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyMaxThreadsPerQueue() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxThreadsPerQueue")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxWorkItems sets the value of MaxWorkItems for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyMaxWorkItems(value uint32) (err error) {
	return instance.SetProperty("MaxWorkItems", value)
}

// GetMaxWorkItems gets the value of MaxWorkItems for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyMaxWorkItems() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxWorkItems")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNullSessionPipes sets the value of NullSessionPipes for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyNullSessionPipes(value string) (err error) {
	return instance.SetProperty("NullSessionPipes", value)
}

// GetNullSessionPipes gets the value of NullSessionPipes for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyNullSessionPipes() (value string, err error) {
	retValue, err := instance.GetProperty("NullSessionPipes")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNullSessionShares sets the value of NullSessionShares for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyNullSessionShares(value string) (err error) {
	return instance.SetProperty("NullSessionShares", value)
}

// GetNullSessionShares gets the value of NullSessionShares for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyNullSessionShares() (value string, err error) {
	retValue, err := instance.GetProperty("NullSessionShares")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOplockBreakWait sets the value of OplockBreakWait for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyOplockBreakWait(value uint32) (err error) {
	return instance.SetProperty("OplockBreakWait", value)
}

// GetOplockBreakWait gets the value of OplockBreakWait for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyOplockBreakWait() (value uint32, err error) {
	retValue, err := instance.GetProperty("OplockBreakWait")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPendingClientTimeoutInSeconds sets the value of PendingClientTimeoutInSeconds for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyPendingClientTimeoutInSeconds(value uint32) (err error) {
	return instance.SetProperty("PendingClientTimeoutInSeconds", value)
}

// GetPendingClientTimeoutInSeconds gets the value of PendingClientTimeoutInSeconds for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyPendingClientTimeoutInSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("PendingClientTimeoutInSeconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRejectUnencryptedAccess sets the value of RejectUnencryptedAccess for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyRejectUnencryptedAccess(value bool) (err error) {
	return instance.SetProperty("RejectUnencryptedAccess", value)
}

// GetRejectUnencryptedAccess gets the value of RejectUnencryptedAccess for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyRejectUnencryptedAccess() (value bool, err error) {
	retValue, err := instance.GetProperty("RejectUnencryptedAccess")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequireSecuritySignature sets the value of RequireSecuritySignature for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyRequireSecuritySignature(value bool) (err error) {
	return instance.SetProperty("RequireSecuritySignature", value)
}

// GetRequireSecuritySignature gets the value of RequireSecuritySignature for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyRequireSecuritySignature() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireSecuritySignature")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerHidden sets the value of ServerHidden for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyServerHidden(value bool) (err error) {
	return instance.SetProperty("ServerHidden", value)
}

// GetServerHidden gets the value of ServerHidden for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyServerHidden() (value bool, err error) {
	retValue, err := instance.GetProperty("ServerHidden")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSmb2CreditsMax sets the value of Smb2CreditsMax for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertySmb2CreditsMax(value uint32) (err error) {
	return instance.SetProperty("Smb2CreditsMax", value)
}

// GetSmb2CreditsMax gets the value of Smb2CreditsMax for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertySmb2CreditsMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("Smb2CreditsMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSmb2CreditsMin sets the value of Smb2CreditsMin for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertySmb2CreditsMin(value uint32) (err error) {
	return instance.SetProperty("Smb2CreditsMin", value)
}

// GetSmb2CreditsMin gets the value of Smb2CreditsMin for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertySmb2CreditsMin() (value uint32, err error) {
	retValue, err := instance.GetProperty("Smb2CreditsMin")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSmbServerNameHardeningLevel sets the value of SmbServerNameHardeningLevel for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertySmbServerNameHardeningLevel(value uint32) (err error) {
	return instance.SetProperty("SmbServerNameHardeningLevel", value)
}

// GetSmbServerNameHardeningLevel gets the value of SmbServerNameHardeningLevel for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertySmbServerNameHardeningLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("SmbServerNameHardeningLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTreatHostAsStableStorage sets the value of TreatHostAsStableStorage for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyTreatHostAsStableStorage(value bool) (err error) {
	return instance.SetProperty("TreatHostAsStableStorage", value)
}

// GetTreatHostAsStableStorage gets the value of TreatHostAsStableStorage for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyTreatHostAsStableStorage() (value bool, err error) {
	retValue, err := instance.GetProperty("TreatHostAsStableStorage")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValidateAliasNotCircular sets the value of ValidateAliasNotCircular for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyValidateAliasNotCircular(value bool) (err error) {
	return instance.SetProperty("ValidateAliasNotCircular", value)
}

// GetValidateAliasNotCircular gets the value of ValidateAliasNotCircular for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyValidateAliasNotCircular() (value bool, err error) {
	retValue, err := instance.GetProperty("ValidateAliasNotCircular")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValidateShareScope sets the value of ValidateShareScope for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyValidateShareScope(value bool) (err error) {
	return instance.SetProperty("ValidateShareScope", value)
}

// GetValidateShareScope gets the value of ValidateShareScope for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyValidateShareScope() (value bool, err error) {
	retValue, err := instance.GetProperty("ValidateShareScope")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValidateShareScopeNotAliased sets the value of ValidateShareScopeNotAliased for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyValidateShareScopeNotAliased(value bool) (err error) {
	return instance.SetProperty("ValidateShareScopeNotAliased", value)
}

// GetValidateShareScopeNotAliased gets the value of ValidateShareScopeNotAliased for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyValidateShareScopeNotAliased() (value bool, err error) {
	retValue, err := instance.GetProperty("ValidateShareScopeNotAliased")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValidateTargetName sets the value of ValidateTargetName for the instance
func (instance *MSFT_SmbServerConfiguration) SetPropertyValidateTargetName(value bool) (err error) {
	return instance.SetProperty("ValidateTargetName", value)
}

// GetValidateTargetName gets the value of ValidateTargetName for the instance
func (instance *MSFT_SmbServerConfiguration) GetPropertyValidateTargetName() (value bool, err error) {
	retValue, err := instance.GetProperty("ValidateTargetName")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
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

// <param name="AnnounceComment" type="string "></param>
// <param name="AnnounceServer" type="bool "></param>
// <param name="AsynchronousCredits" type="uint32 "></param>
// <param name="AuditSmb1Access" type="bool "></param>
// <param name="AutoDisconnectTimeout" type="uint32 "></param>
// <param name="AutoShareServer" type="bool "></param>
// <param name="AutoShareWorkstation" type="bool "></param>
// <param name="CachedOpenLimit" type="uint32 "></param>
// <param name="DurableHandleV2TimeoutInSeconds" type="uint32 "></param>
// <param name="EnableAuthenticateUserSharing" type="bool "></param>
// <param name="EnableDownlevelTimewarp" type="bool "></param>
// <param name="EnableForcedLogoff" type="bool "></param>
// <param name="EnableLeasing" type="bool "></param>
// <param name="EnableMultiChannel" type="bool "></param>
// <param name="EnableOplocks" type="bool "></param>
// <param name="EnableSecuritySignature" type="bool "></param>
// <param name="EnableSMB1Protocol" type="bool "></param>
// <param name="EnableSMB2Protocol" type="bool "></param>
// <param name="EnableStrictNameChecking" type="bool "></param>
// <param name="EncryptData" type="bool "></param>
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
// <param name="RequireSecuritySignature" type="bool "></param>
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
	/* IN */ AutoDisconnectTimeout uint32,
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
	/* IN */ AuditSmb1Access bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetConfiguration", AnnounceServer, AsynchronousCredits, AutoShareServer, AutoShareWorkstation, CachedOpenLimit, AnnounceComment, EnableDownlevelTimewarp, EnableLeasing, EnableMultiChannel, EnableStrictNameChecking, AutoDisconnectTimeout, DurableHandleV2TimeoutInSeconds, EnableAuthenticateUserSharing, EnableForcedLogoff, EnableOplocks, EnableSecuritySignature, ServerHidden, IrpStackSize, KeepAliveTime, MaxChannelPerSession, MaxMpxCount, MaxSessionPerConnection, MaxThreadsPerQueue, MaxWorkItems, NullSessionPipes, NullSessionShares, OplockBreakWait, PendingClientTimeoutInSeconds, RequireSecuritySignature, EnableSMB1Protocol, EnableSMB2Protocol, Smb2CreditsMax, Smb2CreditsMin, SmbServerNameHardeningLevel, TreatHostAsStableStorage, ValidateAliasNotCircular, ValidateShareScope, ValidateShareScopeNotAliased, ValidateTargetName, EncryptData, RejectUnencryptedAccess, AuditSmb1Access)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
