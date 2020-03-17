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

// MSFT_SmbClientConfiguration struct
type MSFT_SmbClientConfiguration struct {
	cim.WmiInstance

	//
	ConnectionCountPerRssNetworkInterface uint32

	//
	DirectoryCacheEntriesMax uint32

	//
	DirectoryCacheEntrySizeMax uint32

	//
	DirectoryCacheLifetime uint32

	//
	DormantFileLimit uint32

	//
	EnableBandwidthThrottling bool

	//
	EnableByteRangeLockingOnReadOnlyFiles bool

	//
	EnableInsecureGuestLogons bool

	//
	EnableLargeMtu bool

	//
	EnableLoadBalanceScaleOut bool

	//
	EnableMultiChannel bool

	//
	EnableSecuritySignature bool

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
	KeepConn uint32

	//
	MaxCmds uint32

	//
	MaximumConnectionCountPerServer uint32

	//
	OplocksDisabled bool

	//
	RequireSecuritySignature bool

	//
	SessionTimeout uint32

	//
	UseOpportunisticLocking bool

	//
	WindowSizeThreshold uint32
}

// SetConnectionCountPerRssNetworkInterface sets the value of ConnectionCountPerRssNetworkInterface for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyConnectionCountPerRssNetworkInterface(value uint32) (err error) {
	return instance.SetProperty("ConnectionCountPerRssNetworkInterface", value)
}

// GetConnectionCountPerRssNetworkInterface gets the value of ConnectionCountPerRssNetworkInterface for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyConnectionCountPerRssNetworkInterface() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectionCountPerRssNetworkInterface")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDirectoryCacheEntriesMax sets the value of DirectoryCacheEntriesMax for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyDirectoryCacheEntriesMax(value uint32) (err error) {
	return instance.SetProperty("DirectoryCacheEntriesMax", value)
}

// GetDirectoryCacheEntriesMax gets the value of DirectoryCacheEntriesMax for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyDirectoryCacheEntriesMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("DirectoryCacheEntriesMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDirectoryCacheEntrySizeMax sets the value of DirectoryCacheEntrySizeMax for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyDirectoryCacheEntrySizeMax(value uint32) (err error) {
	return instance.SetProperty("DirectoryCacheEntrySizeMax", value)
}

// GetDirectoryCacheEntrySizeMax gets the value of DirectoryCacheEntrySizeMax for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyDirectoryCacheEntrySizeMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("DirectoryCacheEntrySizeMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDirectoryCacheLifetime sets the value of DirectoryCacheLifetime for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyDirectoryCacheLifetime(value uint32) (err error) {
	return instance.SetProperty("DirectoryCacheLifetime", value)
}

// GetDirectoryCacheLifetime gets the value of DirectoryCacheLifetime for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyDirectoryCacheLifetime() (value uint32, err error) {
	retValue, err := instance.GetProperty("DirectoryCacheLifetime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDormantFileLimit sets the value of DormantFileLimit for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyDormantFileLimit(value uint32) (err error) {
	return instance.SetProperty("DormantFileLimit", value)
}

// GetDormantFileLimit gets the value of DormantFileLimit for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyDormantFileLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("DormantFileLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableBandwidthThrottling sets the value of EnableBandwidthThrottling for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyEnableBandwidthThrottling(value bool) (err error) {
	return instance.SetProperty("EnableBandwidthThrottling", value)
}

// GetEnableBandwidthThrottling gets the value of EnableBandwidthThrottling for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyEnableBandwidthThrottling() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableBandwidthThrottling")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableByteRangeLockingOnReadOnlyFiles sets the value of EnableByteRangeLockingOnReadOnlyFiles for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyEnableByteRangeLockingOnReadOnlyFiles(value bool) (err error) {
	return instance.SetProperty("EnableByteRangeLockingOnReadOnlyFiles", value)
}

// GetEnableByteRangeLockingOnReadOnlyFiles gets the value of EnableByteRangeLockingOnReadOnlyFiles for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyEnableByteRangeLockingOnReadOnlyFiles() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableByteRangeLockingOnReadOnlyFiles")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableInsecureGuestLogons sets the value of EnableInsecureGuestLogons for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyEnableInsecureGuestLogons(value bool) (err error) {
	return instance.SetProperty("EnableInsecureGuestLogons", value)
}

// GetEnableInsecureGuestLogons gets the value of EnableInsecureGuestLogons for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyEnableInsecureGuestLogons() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableInsecureGuestLogons")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableLargeMtu sets the value of EnableLargeMtu for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyEnableLargeMtu(value bool) (err error) {
	return instance.SetProperty("EnableLargeMtu", value)
}

// GetEnableLargeMtu gets the value of EnableLargeMtu for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyEnableLargeMtu() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableLargeMtu")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableLoadBalanceScaleOut sets the value of EnableLoadBalanceScaleOut for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyEnableLoadBalanceScaleOut(value bool) (err error) {
	return instance.SetProperty("EnableLoadBalanceScaleOut", value)
}

// GetEnableLoadBalanceScaleOut gets the value of EnableLoadBalanceScaleOut for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyEnableLoadBalanceScaleOut() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableLoadBalanceScaleOut")
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
func (instance *MSFT_SmbClientConfiguration) SetPropertyEnableMultiChannel(value bool) (err error) {
	return instance.SetProperty("EnableMultiChannel", value)
}

// GetEnableMultiChannel gets the value of EnableMultiChannel for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyEnableMultiChannel() (value bool, err error) {
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

// SetEnableSecuritySignature sets the value of EnableSecuritySignature for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyEnableSecuritySignature(value bool) (err error) {
	return instance.SetProperty("EnableSecuritySignature", value)
}

// GetEnableSecuritySignature gets the value of EnableSecuritySignature for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyEnableSecuritySignature() (value bool, err error) {
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

// SetExtendedSessionTimeout sets the value of ExtendedSessionTimeout for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyExtendedSessionTimeout(value uint32) (err error) {
	return instance.SetProperty("ExtendedSessionTimeout", value)
}

// GetExtendedSessionTimeout gets the value of ExtendedSessionTimeout for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyExtendedSessionTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExtendedSessionTimeout")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileInfoCacheEntriesMax sets the value of FileInfoCacheEntriesMax for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyFileInfoCacheEntriesMax(value uint32) (err error) {
	return instance.SetProperty("FileInfoCacheEntriesMax", value)
}

// GetFileInfoCacheEntriesMax gets the value of FileInfoCacheEntriesMax for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyFileInfoCacheEntriesMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileInfoCacheEntriesMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileInfoCacheLifetime sets the value of FileInfoCacheLifetime for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyFileInfoCacheLifetime(value uint32) (err error) {
	return instance.SetProperty("FileInfoCacheLifetime", value)
}

// GetFileInfoCacheLifetime gets the value of FileInfoCacheLifetime for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyFileInfoCacheLifetime() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileInfoCacheLifetime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileNotFoundCacheEntriesMax sets the value of FileNotFoundCacheEntriesMax for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyFileNotFoundCacheEntriesMax(value uint32) (err error) {
	return instance.SetProperty("FileNotFoundCacheEntriesMax", value)
}

// GetFileNotFoundCacheEntriesMax gets the value of FileNotFoundCacheEntriesMax for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyFileNotFoundCacheEntriesMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileNotFoundCacheEntriesMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileNotFoundCacheLifetime sets the value of FileNotFoundCacheLifetime for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyFileNotFoundCacheLifetime(value uint32) (err error) {
	return instance.SetProperty("FileNotFoundCacheLifetime", value)
}

// GetFileNotFoundCacheLifetime gets the value of FileNotFoundCacheLifetime for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyFileNotFoundCacheLifetime() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileNotFoundCacheLifetime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKeepConn sets the value of KeepConn for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyKeepConn(value uint32) (err error) {
	return instance.SetProperty("KeepConn", value)
}

// GetKeepConn gets the value of KeepConn for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyKeepConn() (value uint32, err error) {
	retValue, err := instance.GetProperty("KeepConn")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxCmds sets the value of MaxCmds for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyMaxCmds(value uint32) (err error) {
	return instance.SetProperty("MaxCmds", value)
}

// GetMaxCmds gets the value of MaxCmds for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyMaxCmds() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxCmds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumConnectionCountPerServer sets the value of MaximumConnectionCountPerServer for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyMaximumConnectionCountPerServer(value uint32) (err error) {
	return instance.SetProperty("MaximumConnectionCountPerServer", value)
}

// GetMaximumConnectionCountPerServer gets the value of MaximumConnectionCountPerServer for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyMaximumConnectionCountPerServer() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumConnectionCountPerServer")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOplocksDisabled sets the value of OplocksDisabled for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyOplocksDisabled(value bool) (err error) {
	return instance.SetProperty("OplocksDisabled", value)
}

// GetOplocksDisabled gets the value of OplocksDisabled for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyOplocksDisabled() (value bool, err error) {
	retValue, err := instance.GetProperty("OplocksDisabled")
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
func (instance *MSFT_SmbClientConfiguration) SetPropertyRequireSecuritySignature(value bool) (err error) {
	return instance.SetProperty("RequireSecuritySignature", value)
}

// GetRequireSecuritySignature gets the value of RequireSecuritySignature for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyRequireSecuritySignature() (value bool, err error) {
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

// SetSessionTimeout sets the value of SessionTimeout for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertySessionTimeout(value uint32) (err error) {
	return instance.SetProperty("SessionTimeout", value)
}

// GetSessionTimeout gets the value of SessionTimeout for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertySessionTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("SessionTimeout")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUseOpportunisticLocking sets the value of UseOpportunisticLocking for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyUseOpportunisticLocking(value bool) (err error) {
	return instance.SetProperty("UseOpportunisticLocking", value)
}

// GetUseOpportunisticLocking gets the value of UseOpportunisticLocking for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyUseOpportunisticLocking() (value bool, err error) {
	retValue, err := instance.GetProperty("UseOpportunisticLocking")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWindowSizeThreshold sets the value of WindowSizeThreshold for the instance
func (instance *MSFT_SmbClientConfiguration) SetPropertyWindowSizeThreshold(value uint32) (err error) {
	return instance.SetProperty("WindowSizeThreshold", value)
}

// GetWindowSizeThreshold gets the value of WindowSizeThreshold for the instance
func (instance *MSFT_SmbClientConfiguration) GetPropertyWindowSizeThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("WindowSizeThreshold")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
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

// <param name="ConnectionCountPerRssNetworkInterface" type="uint32 "></param>
// <param name="DirectoryCacheEntriesMax" type="uint32 "></param>
// <param name="DirectoryCacheEntrySizeMax" type="uint32 "></param>
// <param name="DirectoryCacheLifetime" type="uint32 "></param>
// <param name="DormantFileLimit" type="uint32 "></param>
// <param name="EnableBandwidthThrottling" type="bool "></param>
// <param name="EnableByteRangeLockingOnReadOnlyFiles" type="bool "></param>
// <param name="EnableInsecureGuestLogons" type="bool "></param>
// <param name="EnableLargeMtu" type="bool "></param>
// <param name="EnableLoadBalanceScaleOut" type="bool "></param>
// <param name="EnableMultiChannel" type="bool "></param>
// <param name="EnableSecuritySignature" type="bool "></param>
// <param name="ExtendedSessionTimeout" type="uint32 "></param>
// <param name="FileInfoCacheEntriesMax" type="uint32 "></param>
// <param name="FileInfoCacheLifetime" type="uint32 "></param>
// <param name="FileNotFoundCacheEntriesMax" type="uint32 "></param>
// <param name="FileNotFoundCacheLifetime" type="uint32 "></param>
// <param name="KeepConn" type="uint32 "></param>
// <param name="MaxCmds" type="uint32 "></param>
// <param name="MaximumConnectionCountPerServer" type="uint32 "></param>
// <param name="OplocksDisabled" type="bool "></param>
// <param name="RequireSecuritySignature" type="bool "></param>
// <param name="SessionTimeout" type="uint32 "></param>
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
	/* IN */ EnableInsecureGuestLogons bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetConfiguration", ConnectionCountPerRssNetworkInterface, DirectoryCacheEntriesMax, DirectoryCacheEntrySizeMax, DirectoryCacheLifetime, EnableBandwidthThrottling, EnableByteRangeLockingOnReadOnlyFiles, EnableLargeMtu, EnableMultiChannel, DormantFileLimit, EnableSecuritySignature, ExtendedSessionTimeout, FileInfoCacheEntriesMax, FileInfoCacheLifetime, FileNotFoundCacheEntriesMax, FileNotFoundCacheLifetime, KeepConn, MaxCmds, MaximumConnectionCountPerServer, OplocksDisabled, RequireSecuritySignature, SessionTimeout, UseOpportunisticLocking, WindowSizeThreshold, EnableLoadBalanceScaleOut, EnableInsecureGuestLogons)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
