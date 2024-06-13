// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
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

// MSFT_SmbClientConfiguration struct
type MSFT_SmbClientConfiguration struct {
	*cim.WmiInstance

	//
	CompressibilitySamplingSize uint64

	//
	CompressibleThreshold uint64

	//
	ConnectionCountPerRssNetworkInterface uint32

	//
	DirectoryCacheEntriesMax uint32

	//
	DirectoryCacheEntrySizeMax uint32

	//
	DirectoryCacheLifetime uint32

	//
	DisableCompression bool

	//
	DormantFileLimit uint32

	//
	EnableBandwidthThrottling bool

	//
	EnableByteRangeLockingOnReadOnlyFiles bool

	//
	EnableCompressibilitySampling bool

	//
	EnableInsecureGuestLogons bool

	//
	EnableLargeMtu bool

	//
	EnableLoadBalanceScaleOut bool

	//
	EnableMailslots bool

	//
	EnableMultiChannel bool

	//
	EnableSecuritySignature bool

	//
	EncryptionCiphers string

	//
	ExtendedSessionTimeout uint32

	//
	FileInfoCacheEntriesMax uint32

	//
	FileInfoCacheLifetime uint32

	//
	FileNotFoundCacheEntriesMax uint32

	//
	FileNotFoundCacheLifetime uint32

	//
	ForceSMBEncryptionOverQuic bool

	//
	InvalidAuthenticationCacheLifetime uint32

	//
	KeepConn uint32

	//
	MaxCmds uint32

	//
	MaximumConnectionCountPerServer uint32

	//
	OplocksDisabled bool

	//
	RequestCompression bool

	//
	RequireSecuritySignature bool

	//
	SessionTimeout uint32

	//
	SkipCertificateCheck bool

	//
	UseOpportunisticLocking bool

	//
	WindowSizeThreshold uint32
}

func NewMSFT_SmbClientConfigurationEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbClientConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbClientConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_SmbClientConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SmbClientConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbClientConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetCompressibilitySamplingSize sets the value of CompressibilitySamplingSize for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyCompressibilitySamplingSize(value uint64) (err error) {
	return instance.SetProperty("CompressibilitySamplingSize", (value))
}

// GetCompressibilitySamplingSize gets the value of CompressibilitySamplingSize for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyCompressibilitySamplingSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("CompressibilitySamplingSize")
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

// SetCompressibleThreshold sets the value of CompressibleThreshold for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyCompressibleThreshold(value uint64) (err error) {
	return instance.SetProperty("CompressibleThreshold", (value))
}

// GetCompressibleThreshold gets the value of CompressibleThreshold for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyCompressibleThreshold() (value uint64, err error) {
	retValue, err := instance.GetProperty("CompressibleThreshold")
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

// SetConnectionCountPerRssNetworkInterface sets the value of ConnectionCountPerRssNetworkInterface for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyConnectionCountPerRssNetworkInterface(value uint32) (err error) {
	return instance.SetProperty("ConnectionCountPerRssNetworkInterface", (value))
}

// GetConnectionCountPerRssNetworkInterface gets the value of ConnectionCountPerRssNetworkInterface for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyConnectionCountPerRssNetworkInterface() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectionCountPerRssNetworkInterface")
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

// SetDirectoryCacheEntriesMax sets the value of DirectoryCacheEntriesMax for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyDirectoryCacheEntriesMax(value uint32) (err error) {
	return instance.SetProperty("DirectoryCacheEntriesMax", (value))
}

// GetDirectoryCacheEntriesMax gets the value of DirectoryCacheEntriesMax for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyDirectoryCacheEntriesMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("DirectoryCacheEntriesMax")
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

// SetDirectoryCacheEntrySizeMax sets the value of DirectoryCacheEntrySizeMax for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyDirectoryCacheEntrySizeMax(value uint32) (err error) {
	return instance.SetProperty("DirectoryCacheEntrySizeMax", (value))
}

// GetDirectoryCacheEntrySizeMax gets the value of DirectoryCacheEntrySizeMax for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyDirectoryCacheEntrySizeMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("DirectoryCacheEntrySizeMax")
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

// SetDirectoryCacheLifetime sets the value of DirectoryCacheLifetime for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyDirectoryCacheLifetime(value uint32) (err error) {
	return instance.SetProperty("DirectoryCacheLifetime", (value))
}

// GetDirectoryCacheLifetime gets the value of DirectoryCacheLifetime for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyDirectoryCacheLifetime() (value uint32, err error) {
	retValue, err := instance.GetProperty("DirectoryCacheLifetime")
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
func (instance *MSFT_SmbClientConfiguration) SetPropertyDisableCompression(value bool) (err error) {
	return instance.SetProperty("DisableCompression", (value))
}

// GetDisableCompression gets the value of DisableCompression for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyDisableCompression() (value bool, err error) {
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

// SetDormantFileLimit sets the value of DormantFileLimit for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyDormantFileLimit(value uint32) (err error) {
	return instance.SetProperty("DormantFileLimit", (value))
}

// GetDormantFileLimit gets the value of DormantFileLimit for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyDormantFileLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("DormantFileLimit")
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

// SetEnableBandwidthThrottling sets the value of EnableBandwidthThrottling for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyEnableBandwidthThrottling(value bool) (err error) {
	return instance.SetProperty("EnableBandwidthThrottling", (value))
}

// GetEnableBandwidthThrottling gets the value of EnableBandwidthThrottling for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyEnableBandwidthThrottling() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableBandwidthThrottling")
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

// SetEnableByteRangeLockingOnReadOnlyFiles sets the value of EnableByteRangeLockingOnReadOnlyFiles for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyEnableByteRangeLockingOnReadOnlyFiles(value bool) (err error) {
	return instance.SetProperty("EnableByteRangeLockingOnReadOnlyFiles", (value))
}

// GetEnableByteRangeLockingOnReadOnlyFiles gets the value of EnableByteRangeLockingOnReadOnlyFiles for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyEnableByteRangeLockingOnReadOnlyFiles() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableByteRangeLockingOnReadOnlyFiles")
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

// SetEnableCompressibilitySampling sets the value of EnableCompressibilitySampling for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyEnableCompressibilitySampling(value bool) (err error) {
	return instance.SetProperty("EnableCompressibilitySampling", (value))
}

// GetEnableCompressibilitySampling gets the value of EnableCompressibilitySampling for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyEnableCompressibilitySampling() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableCompressibilitySampling")
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

// SetEnableInsecureGuestLogons sets the value of EnableInsecureGuestLogons for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyEnableInsecureGuestLogons(value bool) (err error) {
	return instance.SetProperty("EnableInsecureGuestLogons", (value))
}

// GetEnableInsecureGuestLogons gets the value of EnableInsecureGuestLogons for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyEnableInsecureGuestLogons() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableInsecureGuestLogons")
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

// SetEnableLargeMtu sets the value of EnableLargeMtu for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyEnableLargeMtu(value bool) (err error) {
	return instance.SetProperty("EnableLargeMtu", (value))
}

// GetEnableLargeMtu gets the value of EnableLargeMtu for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyEnableLargeMtu() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableLargeMtu")
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

// SetEnableLoadBalanceScaleOut sets the value of EnableLoadBalanceScaleOut for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyEnableLoadBalanceScaleOut(value bool) (err error) {
	return instance.SetProperty("EnableLoadBalanceScaleOut", (value))
}

// GetEnableLoadBalanceScaleOut gets the value of EnableLoadBalanceScaleOut for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyEnableLoadBalanceScaleOut() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableLoadBalanceScaleOut")
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

// SetEnableMailslots sets the value of EnableMailslots for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyEnableMailslots(value bool) (err error) {
	return instance.SetProperty("EnableMailslots", (value))
}

// GetEnableMailslots gets the value of EnableMailslots for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyEnableMailslots() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableMailslots")
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
func (instance *MSFT_SmbClientConfiguration) SetPropertyEnableMultiChannel(value bool) (err error) {
	return instance.SetProperty("EnableMultiChannel", (value))
}

// GetEnableMultiChannel gets the value of EnableMultiChannel for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyEnableMultiChannel() (value bool, err error) {
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

// SetEnableSecuritySignature sets the value of EnableSecuritySignature for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyEnableSecuritySignature(value bool) (err error) {
	return instance.SetProperty("EnableSecuritySignature", (value))
}

// GetEnableSecuritySignature gets the value of EnableSecuritySignature for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyEnableSecuritySignature() (value bool, err error) {
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

// SetEncryptionCiphers sets the value of EncryptionCiphers for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyEncryptionCiphers(value string) (err error) {
	return instance.SetProperty("EncryptionCiphers", (value))
}

// GetEncryptionCiphers gets the value of EncryptionCiphers for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyEncryptionCiphers() (value string, err error) {
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

// SetExtendedSessionTimeout sets the value of ExtendedSessionTimeout for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyExtendedSessionTimeout(value uint32) (err error) {
	return instance.SetProperty("ExtendedSessionTimeout", (value))
}

// GetExtendedSessionTimeout gets the value of ExtendedSessionTimeout for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyExtendedSessionTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExtendedSessionTimeout")
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

// SetFileInfoCacheEntriesMax sets the value of FileInfoCacheEntriesMax for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyFileInfoCacheEntriesMax(value uint32) (err error) {
	return instance.SetProperty("FileInfoCacheEntriesMax", (value))
}

// GetFileInfoCacheEntriesMax gets the value of FileInfoCacheEntriesMax for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyFileInfoCacheEntriesMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileInfoCacheEntriesMax")
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

// SetFileInfoCacheLifetime sets the value of FileInfoCacheLifetime for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyFileInfoCacheLifetime(value uint32) (err error) {
	return instance.SetProperty("FileInfoCacheLifetime", (value))
}

// GetFileInfoCacheLifetime gets the value of FileInfoCacheLifetime for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyFileInfoCacheLifetime() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileInfoCacheLifetime")
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

// SetFileNotFoundCacheEntriesMax sets the value of FileNotFoundCacheEntriesMax for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyFileNotFoundCacheEntriesMax(value uint32) (err error) {
	return instance.SetProperty("FileNotFoundCacheEntriesMax", (value))
}

// GetFileNotFoundCacheEntriesMax gets the value of FileNotFoundCacheEntriesMax for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyFileNotFoundCacheEntriesMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileNotFoundCacheEntriesMax")
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

// SetFileNotFoundCacheLifetime sets the value of FileNotFoundCacheLifetime for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyFileNotFoundCacheLifetime(value uint32) (err error) {
	return instance.SetProperty("FileNotFoundCacheLifetime", (value))
}

// GetFileNotFoundCacheLifetime gets the value of FileNotFoundCacheLifetime for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyFileNotFoundCacheLifetime() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileNotFoundCacheLifetime")
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

// SetForceSMBEncryptionOverQuic sets the value of ForceSMBEncryptionOverQuic for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyForceSMBEncryptionOverQuic(value bool) (err error) {
	return instance.SetProperty("ForceSMBEncryptionOverQuic", (value))
}

// GetForceSMBEncryptionOverQuic gets the value of ForceSMBEncryptionOverQuic for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyForceSMBEncryptionOverQuic() (value bool, err error) {
	retValue, err := instance.GetProperty("ForceSMBEncryptionOverQuic")
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

// SetInvalidAuthenticationCacheLifetime sets the value of InvalidAuthenticationCacheLifetime for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyInvalidAuthenticationCacheLifetime(value uint32) (err error) {
	return instance.SetProperty("InvalidAuthenticationCacheLifetime", (value))
}

// GetInvalidAuthenticationCacheLifetime gets the value of InvalidAuthenticationCacheLifetime for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyInvalidAuthenticationCacheLifetime() (value uint32, err error) {
	retValue, err := instance.GetProperty("InvalidAuthenticationCacheLifetime")
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

// SetKeepConn sets the value of KeepConn for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyKeepConn(value uint32) (err error) {
	return instance.SetProperty("KeepConn", (value))
}

// GetKeepConn gets the value of KeepConn for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyKeepConn() (value uint32, err error) {
	retValue, err := instance.GetProperty("KeepConn")
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

// SetMaxCmds sets the value of MaxCmds for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyMaxCmds(value uint32) (err error) {
	return instance.SetProperty("MaxCmds", (value))
}

// GetMaxCmds gets the value of MaxCmds for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyMaxCmds() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxCmds")
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

// SetMaximumConnectionCountPerServer sets the value of MaximumConnectionCountPerServer for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyMaximumConnectionCountPerServer(value uint32) (err error) {
	return instance.SetProperty("MaximumConnectionCountPerServer", (value))
}

// GetMaximumConnectionCountPerServer gets the value of MaximumConnectionCountPerServer for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyMaximumConnectionCountPerServer() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumConnectionCountPerServer")
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

// SetOplocksDisabled sets the value of OplocksDisabled for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyOplocksDisabled(value bool) (err error) {
	return instance.SetProperty("OplocksDisabled", (value))
}

// GetOplocksDisabled gets the value of OplocksDisabled for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyOplocksDisabled() (value bool, err error) {
	retValue, err := instance.GetProperty("OplocksDisabled")
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
func (instance *MSFT_SmbClientConfiguration) SetPropertyRequestCompression(value bool) (err error) {
	return instance.SetProperty("RequestCompression", (value))
}

// GetRequestCompression gets the value of RequestCompression for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyRequestCompression() (value bool, err error) {
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
func (instance *MSFT_SmbClientConfiguration) SetPropertyRequireSecuritySignature(value bool) (err error) {
	return instance.SetProperty("RequireSecuritySignature", (value))
}

// GetRequireSecuritySignature gets the value of RequireSecuritySignature for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyRequireSecuritySignature() (value bool, err error) {
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

// SetSessionTimeout sets the value of SessionTimeout for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertySessionTimeout(value uint32) (err error) {
	return instance.SetProperty("SessionTimeout", (value))
}

// GetSessionTimeout gets the value of SessionTimeout for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertySessionTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("SessionTimeout")
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

// SetSkipCertificateCheck sets the value of SkipCertificateCheck for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertySkipCertificateCheck(value bool) (err error) {
	return instance.SetProperty("SkipCertificateCheck", (value))
}

// GetSkipCertificateCheck gets the value of SkipCertificateCheck for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertySkipCertificateCheck() (value bool, err error) {
	retValue, err := instance.GetProperty("SkipCertificateCheck")
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

// SetUseOpportunisticLocking sets the value of UseOpportunisticLocking for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyUseOpportunisticLocking(value bool) (err error) {
	return instance.SetProperty("UseOpportunisticLocking", (value))
}

// GetUseOpportunisticLocking gets the value of UseOpportunisticLocking for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyUseOpportunisticLocking() (value bool, err error) {
	retValue, err := instance.GetProperty("UseOpportunisticLocking")
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

// SetWindowSizeThreshold sets the value of WindowSizeThreshold for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyWindowSizeThreshold(value uint32) (err error) {
	return instance.SetProperty("WindowSizeThreshold", (value))
}

// GetWindowSizeThreshold gets the value of WindowSizeThreshold for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyWindowSizeThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("WindowSizeThreshold")
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

//

// <param name="Output" type="MSFT_SmbClientConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbClientConfiguration) GetConfiguration( /* OUT */ Output MSFT_SmbClientConfiguration) (result uint32, err error) {
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
// <param name="CompressibilitySampling" type="bool "></param>
// <param name="ConnectionCountPerRssNetworkInterface" type="bool "></param>
// <param name="DirectoryCacheEntriesMax" type="bool "></param>
// <param name="DirectoryCacheEntrySizeMax" type="bool "></param>
// <param name="DirectoryCacheLifetime" type="bool "></param>
// <param name="DisableCompression" type="bool "></param>
// <param name="DormantFileLimit" type="bool "></param>
// <param name="EnableBandwidthThrottling" type="bool "></param>
// <param name="EnableByteRangeLockingOnReadOnlyFiles" type="bool "></param>
// <param name="EnableLargeMtu" type="bool "></param>
// <param name="EnableLoadBalanceScaleOut" type="bool "></param>
// <param name="EnableMailslots" type="bool "></param>
// <param name="EnableMultiChannel" type="bool "></param>
// <param name="EncryptionCiphers" type="bool "></param>
// <param name="ExtendedSessionTimeout" type="bool "></param>
// <param name="FileInfoCacheEntriesMax" type="bool "></param>
// <param name="FileInfoCacheLifetime" type="bool "></param>
// <param name="FileNotFoundCacheEntriesMax" type="bool "></param>
// <param name="FileNotFoundCacheLifetime" type="bool "></param>
// <param name="ForceSMBEncryptionOverQuic" type="bool "></param>
// <param name="InvalidAuthenticationCacheLifetime" type="bool "></param>
// <param name="KeepConn" type="bool "></param>
// <param name="MaxCmds" type="bool "></param>
// <param name="MaximumConnectionCountPerServer" type="bool "></param>
// <param name="OplocksDisabled" type="bool "></param>
// <param name="RequestCompression" type="bool "></param>
// <param name="SessionTimeout" type="bool "></param>
// <param name="SkipCertificateCheck" type="bool "></param>
// <param name="UseOpportunisticLocking" type="bool "></param>
// <param name="WindowSizeThreshold" type="bool "></param>

// <param name="Output" type="MSFT_SmbClientConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbClientConfiguration) ResetConfiguration( /* OUT */ Output MSFT_SmbClientConfiguration,
	/* OPTIONAL IN */ All bool,
	/* OPTIONAL IN */ ConnectionCountPerRssNetworkInterface bool,
	/* OPTIONAL IN */ DirectoryCacheEntriesMax bool,
	/* OPTIONAL IN */ DirectoryCacheEntrySizeMax bool,
	/* OPTIONAL IN */ DirectoryCacheLifetime bool,
	/* OPTIONAL IN */ EnableBandwidthThrottling bool,
	/* OPTIONAL IN */ EnableByteRangeLockingOnReadOnlyFiles bool,
	/* OPTIONAL IN */ EnableLargeMtu bool,
	/* OPTIONAL IN */ EnableMailslots bool,
	/* OPTIONAL IN */ EnableMultiChannel bool,
	/* OPTIONAL IN */ DormantFileLimit bool,
	/* OPTIONAL IN */ ExtendedSessionTimeout bool,
	/* OPTIONAL IN */ FileInfoCacheEntriesMax bool,
	/* OPTIONAL IN */ FileInfoCacheLifetime bool,
	/* OPTIONAL IN */ FileNotFoundCacheEntriesMax bool,
	/* OPTIONAL IN */ FileNotFoundCacheLifetime bool,
	/* OPTIONAL IN */ KeepConn bool,
	/* OPTIONAL IN */ MaxCmds bool,
	/* OPTIONAL IN */ MaximumConnectionCountPerServer bool,
	/* OPTIONAL IN */ OplocksDisabled bool,
	/* OPTIONAL IN */ SessionTimeout bool,
	/* OPTIONAL IN */ UseOpportunisticLocking bool,
	/* OPTIONAL IN */ WindowSizeThreshold bool,
	/* OPTIONAL IN */ EnableLoadBalanceScaleOut bool,
	/* OPTIONAL IN */ ForceSMBEncryptionOverQuic bool,
	/* OPTIONAL IN */ SkipCertificateCheck bool,
	/* OPTIONAL IN */ RequestCompression bool,
	/* OPTIONAL IN */ DisableCompression bool,
	/* OPTIONAL IN */ CompressibilitySampling bool,
	/* OPTIONAL IN */ EncryptionCiphers bool,
	/* OPTIONAL IN */ InvalidAuthenticationCacheLifetime bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ResetConfiguration", All, ConnectionCountPerRssNetworkInterface, DirectoryCacheEntriesMax, DirectoryCacheEntrySizeMax, DirectoryCacheLifetime, EnableBandwidthThrottling, EnableByteRangeLockingOnReadOnlyFiles, EnableLargeMtu, EnableMailslots, EnableMultiChannel, DormantFileLimit, ExtendedSessionTimeout, FileInfoCacheEntriesMax, FileInfoCacheLifetime, FileNotFoundCacheEntriesMax, FileNotFoundCacheLifetime, KeepConn, MaxCmds, MaximumConnectionCountPerServer, OplocksDisabled, SessionTimeout, UseOpportunisticLocking, WindowSizeThreshold, EnableLoadBalanceScaleOut, ForceSMBEncryptionOverQuic, SkipCertificateCheck, RequestCompression, DisableCompression, CompressibilitySampling, EncryptionCiphers, InvalidAuthenticationCacheLifetime)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CompressibilitySamplingSize" type="uint64 "></param>
// <param name="CompressibleThreshold" type="uint64 "></param>
// <param name="ConnectionCountPerRssNetworkInterface" type="uint32 "></param>
// <param name="DirectoryCacheEntriesMax" type="uint32 "></param>
// <param name="DirectoryCacheEntrySizeMax" type="uint32 "></param>
// <param name="DirectoryCacheLifetime" type="uint32 "></param>
// <param name="DisableCompression" type="bool "></param>
// <param name="DormantFileLimit" type="uint32 "></param>
// <param name="EnableBandwidthThrottling" type="bool "></param>
// <param name="EnableByteRangeLockingOnReadOnlyFiles" type="bool "></param>
// <param name="EnableCompressibilitySampling" type="bool "></param>
// <param name="EnableInsecureGuestLogons" type="bool "></param>
// <param name="EnableLargeMtu" type="bool "></param>
// <param name="EnableLoadBalanceScaleOut" type="bool "></param>
// <param name="EnableMailslots" type="bool "></param>
// <param name="EnableMultiChannel" type="bool "></param>
// <param name="EnableSecuritySignature" type="bool "></param>
// <param name="EncryptionCiphers" type="string "></param>
// <param name="ExtendedSessionTimeout" type="uint32 "></param>
// <param name="FileInfoCacheEntriesMax" type="uint32 "></param>
// <param name="FileInfoCacheLifetime" type="uint32 "></param>
// <param name="FileNotFoundCacheEntriesMax" type="uint32 "></param>
// <param name="FileNotFoundCacheLifetime" type="uint32 "></param>
// <param name="ForceSMBEncryptionOverQuic" type="bool "></param>
// <param name="InvalidAuthenticationCacheLifetime" type="uint32 "></param>
// <param name="KeepConn" type="uint32 "></param>
// <param name="MaxCmds" type="uint32 "></param>
// <param name="MaximumConnectionCountPerServer" type="uint32 "></param>
// <param name="OplocksDisabled" type="bool "></param>
// <param name="RequestCompression" type="bool "></param>
// <param name="RequireSecuritySignature" type="bool "></param>
// <param name="SessionTimeout" type="uint32 "></param>
// <param name="SkipCertificateCheck" type="bool "></param>
// <param name="UseOpportunisticLocking" type="bool "></param>
// <param name="WindowSizeThreshold" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbClientConfiguration) SetConfiguration( /* IN */ ConnectionCountPerRssNetworkInterface uint32,
	/* IN */ DirectoryCacheEntriesMax uint32,
	/* IN */ DirectoryCacheEntrySizeMax uint32,
	/* IN */ DirectoryCacheLifetime uint32,
	/* IN */ EnableBandwidthThrottling bool,
	/* IN */ EnableByteRangeLockingOnReadOnlyFiles bool,
	/* IN */ EnableLargeMtu bool,
	/* IN */ EnableMailslots bool,
	/* IN */ EnableMultiChannel bool,
	/* IN */ DormantFileLimit uint32,
	/* IN */ EnableSecuritySignature bool,
	/* IN */ ExtendedSessionTimeout uint32,
	/* IN */ FileInfoCacheEntriesMax uint32,
	/* IN */ FileInfoCacheLifetime uint32,
	/* IN */ FileNotFoundCacheEntriesMax uint32,
	/* IN */ FileNotFoundCacheLifetime uint32,
	/* IN */ KeepConn uint32,
	/* IN */ MaxCmds uint32,
	/* IN */ MaximumConnectionCountPerServer uint32,
	/* IN */ OplocksDisabled bool,
	/* IN */ RequireSecuritySignature bool,
	/* IN */ SessionTimeout uint32,
	/* IN */ UseOpportunisticLocking bool,
	/* IN */ WindowSizeThreshold uint32,
	/* IN */ EnableLoadBalanceScaleOut bool,
	/* IN */ EnableInsecureGuestLogons bool,
	/* IN */ ForceSMBEncryptionOverQuic bool,
	/* IN */ SkipCertificateCheck bool,
	/* IN */ RequestCompression bool,
	/* IN */ DisableCompression bool,
	/* IN */ EnableCompressibilitySampling bool,
	/* IN */ CompressibilitySamplingSize uint64,
	/* IN */ CompressibleThreshold uint64,
	/* IN */ EncryptionCiphers string,
	/* IN */ InvalidAuthenticationCacheLifetime uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetConfiguration", ConnectionCountPerRssNetworkInterface, DirectoryCacheEntriesMax, DirectoryCacheEntrySizeMax, DirectoryCacheLifetime, EnableBandwidthThrottling, EnableByteRangeLockingOnReadOnlyFiles, EnableLargeMtu, EnableMailslots, EnableMultiChannel, DormantFileLimit, EnableSecuritySignature, ExtendedSessionTimeout, FileInfoCacheEntriesMax, FileInfoCacheLifetime, FileNotFoundCacheEntriesMax, FileNotFoundCacheLifetime, KeepConn, MaxCmds, MaximumConnectionCountPerServer, OplocksDisabled, RequireSecuritySignature, SessionTimeout, UseOpportunisticLocking, WindowSizeThreshold, EnableLoadBalanceScaleOut, EnableInsecureGuestLogons, ForceSMBEncryptionOverQuic, SkipCertificateCheck, RequestCompression, DisableCompression, EnableCompressibilitySampling, CompressibilitySamplingSize, CompressibleThreshold, EncryptionCiphers, InvalidAuthenticationCacheLifetime)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
