// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DeliveryOptimization
//////////////////////////////////////////////
package deliveryoptimization

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DeliveryOptimizationFile struct
type MSFT_DeliveryOptimizationFile struct {
	cim.WmiInstance

	//
	BytesFromCacheServer uint64

	//
	BytesFromGroupPeers uint64

	//
	BytesFromHttp uint64

	//
	BytesFromInternetPeers uint64

	//
	BytesFromLanPeers uint64

	//
	BytesToGroupPeers uint64

	//
	BytesToInternetPeers uint64

	//
	BytesToLanPeers uint64

	//
	CacheHost string

	//
	CacheServerConnectionCount uint32

	//
	DownloadDurationMsecs uint64

	//
	DownloadMode DeliveryOptimizationFile_DownloadMode

	//
	ExpireOn string

	//
	FileId string

	//
	FileSize uint64

	//
	FileSizeInCache uint64

	//
	GroupConnectionCount uint32

	//
	HttpConnectionCount uint32

	//
	InternetConnectionCount uint32

	//
	IsBackground bool

	//
	IsPinned bool

	//
	LanConnectionCount uint32

	//
	PeerCount uint32

	//
	PredefinedCallerApplication string

	//
	SourceURL string

	//
	Status DeliveryOptimizationFile_Status

	//
	TotalBytesDownloaded uint64
}

// SetBytesFromCacheServer sets the value of BytesFromCacheServer for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyBytesFromCacheServer(value uint64) (err error) {
	return instance.SetProperty("BytesFromCacheServer", value)
}

// GetBytesFromCacheServer gets the value of BytesFromCacheServer for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyBytesFromCacheServer() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesFromCacheServer")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesFromGroupPeers sets the value of BytesFromGroupPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyBytesFromGroupPeers(value uint64) (err error) {
	return instance.SetProperty("BytesFromGroupPeers", value)
}

// GetBytesFromGroupPeers gets the value of BytesFromGroupPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyBytesFromGroupPeers() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesFromGroupPeers")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesFromHttp sets the value of BytesFromHttp for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyBytesFromHttp(value uint64) (err error) {
	return instance.SetProperty("BytesFromHttp", value)
}

// GetBytesFromHttp gets the value of BytesFromHttp for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyBytesFromHttp() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesFromHttp")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesFromInternetPeers sets the value of BytesFromInternetPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyBytesFromInternetPeers(value uint64) (err error) {
	return instance.SetProperty("BytesFromInternetPeers", value)
}

// GetBytesFromInternetPeers gets the value of BytesFromInternetPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyBytesFromInternetPeers() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesFromInternetPeers")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesFromLanPeers sets the value of BytesFromLanPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyBytesFromLanPeers(value uint64) (err error) {
	return instance.SetProperty("BytesFromLanPeers", value)
}

// GetBytesFromLanPeers gets the value of BytesFromLanPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyBytesFromLanPeers() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesFromLanPeers")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesToGroupPeers sets the value of BytesToGroupPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyBytesToGroupPeers(value uint64) (err error) {
	return instance.SetProperty("BytesToGroupPeers", value)
}

// GetBytesToGroupPeers gets the value of BytesToGroupPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyBytesToGroupPeers() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesToGroupPeers")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesToInternetPeers sets the value of BytesToInternetPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyBytesToInternetPeers(value uint64) (err error) {
	return instance.SetProperty("BytesToInternetPeers", value)
}

// GetBytesToInternetPeers gets the value of BytesToInternetPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyBytesToInternetPeers() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesToInternetPeers")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesToLanPeers sets the value of BytesToLanPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyBytesToLanPeers(value uint64) (err error) {
	return instance.SetProperty("BytesToLanPeers", value)
}

// GetBytesToLanPeers gets the value of BytesToLanPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyBytesToLanPeers() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesToLanPeers")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheHost sets the value of CacheHost for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyCacheHost(value string) (err error) {
	return instance.SetProperty("CacheHost", value)
}

// GetCacheHost gets the value of CacheHost for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyCacheHost() (value string, err error) {
	retValue, err := instance.GetProperty("CacheHost")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheServerConnectionCount sets the value of CacheServerConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyCacheServerConnectionCount(value uint32) (err error) {
	return instance.SetProperty("CacheServerConnectionCount", value)
}

// GetCacheServerConnectionCount gets the value of CacheServerConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyCacheServerConnectionCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("CacheServerConnectionCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDownloadDurationMsecs sets the value of DownloadDurationMsecs for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyDownloadDurationMsecs(value uint64) (err error) {
	return instance.SetProperty("DownloadDurationMsecs", value)
}

// GetDownloadDurationMsecs gets the value of DownloadDurationMsecs for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyDownloadDurationMsecs() (value uint64, err error) {
	retValue, err := instance.GetProperty("DownloadDurationMsecs")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDownloadMode sets the value of DownloadMode for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyDownloadMode(value DeliveryOptimizationFile_DownloadMode) (err error) {
	return instance.SetProperty("DownloadMode", value)
}

// GetDownloadMode gets the value of DownloadMode for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyDownloadMode() (value DeliveryOptimizationFile_DownloadMode, err error) {
	retValue, err := instance.GetProperty("DownloadMode")
	if err != nil {
		return
	}
	value, ok := retValue.(DeliveryOptimizationFile_DownloadMode)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExpireOn sets the value of ExpireOn for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyExpireOn(value string) (err error) {
	return instance.SetProperty("ExpireOn", value)
}

// GetExpireOn gets the value of ExpireOn for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyExpireOn() (value string, err error) {
	retValue, err := instance.GetProperty("ExpireOn")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileId sets the value of FileId for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyFileId(value string) (err error) {
	return instance.SetProperty("FileId", value)
}

// GetFileId gets the value of FileId for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyFileId() (value string, err error) {
	retValue, err := instance.GetProperty("FileId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileSize sets the value of FileSize for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyFileSize(value uint64) (err error) {
	return instance.SetProperty("FileSize", value)
}

// GetFileSize gets the value of FileSize for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyFileSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileSizeInCache sets the value of FileSizeInCache for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyFileSizeInCache(value uint64) (err error) {
	return instance.SetProperty("FileSizeInCache", value)
}

// GetFileSizeInCache gets the value of FileSizeInCache for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyFileSizeInCache() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileSizeInCache")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGroupConnectionCount sets the value of GroupConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyGroupConnectionCount(value uint32) (err error) {
	return instance.SetProperty("GroupConnectionCount", value)
}

// GetGroupConnectionCount gets the value of GroupConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyGroupConnectionCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("GroupConnectionCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHttpConnectionCount sets the value of HttpConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyHttpConnectionCount(value uint32) (err error) {
	return instance.SetProperty("HttpConnectionCount", value)
}

// GetHttpConnectionCount gets the value of HttpConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyHttpConnectionCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("HttpConnectionCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternetConnectionCount sets the value of InternetConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyInternetConnectionCount(value uint32) (err error) {
	return instance.SetProperty("InternetConnectionCount", value)
}

// GetInternetConnectionCount gets the value of InternetConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyInternetConnectionCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("InternetConnectionCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsBackground sets the value of IsBackground for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyIsBackground(value bool) (err error) {
	return instance.SetProperty("IsBackground", value)
}

// GetIsBackground gets the value of IsBackground for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyIsBackground() (value bool, err error) {
	retValue, err := instance.GetProperty("IsBackground")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsPinned sets the value of IsPinned for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyIsPinned(value bool) (err error) {
	return instance.SetProperty("IsPinned", value)
}

// GetIsPinned gets the value of IsPinned for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyIsPinned() (value bool, err error) {
	retValue, err := instance.GetProperty("IsPinned")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLanConnectionCount sets the value of LanConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyLanConnectionCount(value uint32) (err error) {
	return instance.SetProperty("LanConnectionCount", value)
}

// GetLanConnectionCount gets the value of LanConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyLanConnectionCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("LanConnectionCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPeerCount sets the value of PeerCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyPeerCount(value uint32) (err error) {
	return instance.SetProperty("PeerCount", value)
}

// GetPeerCount gets the value of PeerCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyPeerCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("PeerCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPredefinedCallerApplication sets the value of PredefinedCallerApplication for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyPredefinedCallerApplication(value string) (err error) {
	return instance.SetProperty("PredefinedCallerApplication", value)
}

// GetPredefinedCallerApplication gets the value of PredefinedCallerApplication for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyPredefinedCallerApplication() (value string, err error) {
	retValue, err := instance.GetProperty("PredefinedCallerApplication")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceURL sets the value of SourceURL for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertySourceURL(value string) (err error) {
	return instance.SetProperty("SourceURL", value)
}

// GetSourceURL gets the value of SourceURL for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertySourceURL() (value string, err error) {
	retValue, err := instance.GetProperty("SourceURL")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyStatus(value DeliveryOptimizationFile_Status) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyStatus() (value DeliveryOptimizationFile_Status, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(DeliveryOptimizationFile_Status)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalBytesDownloaded sets the value of TotalBytesDownloaded for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyTotalBytesDownloaded(value uint64) (err error) {
	return instance.SetProperty("TotalBytesDownloaded", value)
}

// GetTotalBytesDownloaded gets the value of TotalBytesDownloaded for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyTotalBytesDownloaded() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalBytesDownloaded")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
