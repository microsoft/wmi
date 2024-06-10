// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.DeliveryOptimization
//
// ////////////////////////////////////////////
package deliveryoptimization

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_DeliveryOptimizationFile struct
type MSFT_DeliveryOptimizationFile struct {
	*cim.WmiInstance

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
	BytesFromLinkLocalPeers uint64

	//
	BytesToGroupPeers uint64

	//
	BytesToInternetPeers uint64

	//
	BytesToLanPeers uint64

	//
	BytesToLinkLocalPeers uint64

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
	LinkLocalConnectionCount uint32

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

func NewMSFT_DeliveryOptimizationFileEx1(instance *cim.WmiInstance) (newInstance *MSFT_DeliveryOptimizationFile, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_DeliveryOptimizationFile{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_DeliveryOptimizationFileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DeliveryOptimizationFile, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DeliveryOptimizationFile{
		WmiInstance: tmp,
	}
	return
}

// SetBytesFromCacheServer sets the value of BytesFromCacheServer for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyBytesFromCacheServer(value uint64) (err error) {
	return instance.SetProperty("BytesFromCacheServer", (value))
}

// GetBytesFromCacheServer gets the value of BytesFromCacheServer for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyBytesFromCacheServer() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesFromCacheServer")
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

// SetBytesFromGroupPeers sets the value of BytesFromGroupPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyBytesFromGroupPeers(value uint64) (err error) {
	return instance.SetProperty("BytesFromGroupPeers", (value))
}

// GetBytesFromGroupPeers gets the value of BytesFromGroupPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyBytesFromGroupPeers() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesFromGroupPeers")
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

// SetBytesFromHttp sets the value of BytesFromHttp for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyBytesFromHttp(value uint64) (err error) {
	return instance.SetProperty("BytesFromHttp", (value))
}

// GetBytesFromHttp gets the value of BytesFromHttp for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyBytesFromHttp() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesFromHttp")
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

// SetBytesFromInternetPeers sets the value of BytesFromInternetPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyBytesFromInternetPeers(value uint64) (err error) {
	return instance.SetProperty("BytesFromInternetPeers", (value))
}

// GetBytesFromInternetPeers gets the value of BytesFromInternetPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyBytesFromInternetPeers() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesFromInternetPeers")
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

// SetBytesFromLanPeers sets the value of BytesFromLanPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyBytesFromLanPeers(value uint64) (err error) {
	return instance.SetProperty("BytesFromLanPeers", (value))
}

// GetBytesFromLanPeers gets the value of BytesFromLanPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyBytesFromLanPeers() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesFromLanPeers")
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

// SetBytesFromLinkLocalPeers sets the value of BytesFromLinkLocalPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyBytesFromLinkLocalPeers(value uint64) (err error) {
	return instance.SetProperty("BytesFromLinkLocalPeers", (value))
}

// GetBytesFromLinkLocalPeers gets the value of BytesFromLinkLocalPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyBytesFromLinkLocalPeers() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesFromLinkLocalPeers")
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

// SetBytesToGroupPeers sets the value of BytesToGroupPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyBytesToGroupPeers(value uint64) (err error) {
	return instance.SetProperty("BytesToGroupPeers", (value))
}

// GetBytesToGroupPeers gets the value of BytesToGroupPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyBytesToGroupPeers() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesToGroupPeers")
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

// SetBytesToInternetPeers sets the value of BytesToInternetPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyBytesToInternetPeers(value uint64) (err error) {
	return instance.SetProperty("BytesToInternetPeers", (value))
}

// GetBytesToInternetPeers gets the value of BytesToInternetPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyBytesToInternetPeers() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesToInternetPeers")
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

// SetBytesToLanPeers sets the value of BytesToLanPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyBytesToLanPeers(value uint64) (err error) {
	return instance.SetProperty("BytesToLanPeers", (value))
}

// GetBytesToLanPeers gets the value of BytesToLanPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyBytesToLanPeers() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesToLanPeers")
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

// SetBytesToLinkLocalPeers sets the value of BytesToLinkLocalPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyBytesToLinkLocalPeers(value uint64) (err error) {
	return instance.SetProperty("BytesToLinkLocalPeers", (value))
}

// GetBytesToLinkLocalPeers gets the value of BytesToLinkLocalPeers for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyBytesToLinkLocalPeers() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesToLinkLocalPeers")
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

// SetCacheHost sets the value of CacheHost for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyCacheHost(value string) (err error) {
	return instance.SetProperty("CacheHost", (value))
}

// GetCacheHost gets the value of CacheHost for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyCacheHost() (value string, err error) {
	retValue, err := instance.GetProperty("CacheHost")
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

// SetCacheServerConnectionCount sets the value of CacheServerConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyCacheServerConnectionCount(value uint32) (err error) {
	return instance.SetProperty("CacheServerConnectionCount", (value))
}

// GetCacheServerConnectionCount gets the value of CacheServerConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyCacheServerConnectionCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("CacheServerConnectionCount")
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

// SetDownloadDurationMsecs sets the value of DownloadDurationMsecs for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyDownloadDurationMsecs(value uint64) (err error) {
	return instance.SetProperty("DownloadDurationMsecs", (value))
}

// GetDownloadDurationMsecs gets the value of DownloadDurationMsecs for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyDownloadDurationMsecs() (value uint64, err error) {
	retValue, err := instance.GetProperty("DownloadDurationMsecs")
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

// SetDownloadMode sets the value of DownloadMode for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyDownloadMode(value DeliveryOptimizationFile_DownloadMode) (err error) {
	return instance.SetProperty("DownloadMode", (value))
}

// GetDownloadMode gets the value of DownloadMode for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyDownloadMode() (value DeliveryOptimizationFile_DownloadMode, err error) {
	retValue, err := instance.GetProperty("DownloadMode")
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

	value = DeliveryOptimizationFile_DownloadMode(valuetmp)

	return
}

// SetExpireOn sets the value of ExpireOn for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyExpireOn(value string) (err error) {
	return instance.SetProperty("ExpireOn", (value))
}

// GetExpireOn gets the value of ExpireOn for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyExpireOn() (value string, err error) {
	retValue, err := instance.GetProperty("ExpireOn")
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

// SetFileId sets the value of FileId for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyFileId(value string) (err error) {
	return instance.SetProperty("FileId", (value))
}

// GetFileId gets the value of FileId for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyFileId() (value string, err error) {
	retValue, err := instance.GetProperty("FileId")
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

// SetFileSize sets the value of FileSize for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyFileSize(value uint64) (err error) {
	return instance.SetProperty("FileSize", (value))
}

// GetFileSize gets the value of FileSize for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyFileSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileSize")
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

// SetFileSizeInCache sets the value of FileSizeInCache for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyFileSizeInCache(value uint64) (err error) {
	return instance.SetProperty("FileSizeInCache", (value))
}

// GetFileSizeInCache gets the value of FileSizeInCache for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyFileSizeInCache() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileSizeInCache")
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

// SetGroupConnectionCount sets the value of GroupConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyGroupConnectionCount(value uint32) (err error) {
	return instance.SetProperty("GroupConnectionCount", (value))
}

// GetGroupConnectionCount gets the value of GroupConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyGroupConnectionCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("GroupConnectionCount")
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

// SetHttpConnectionCount sets the value of HttpConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyHttpConnectionCount(value uint32) (err error) {
	return instance.SetProperty("HttpConnectionCount", (value))
}

// GetHttpConnectionCount gets the value of HttpConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyHttpConnectionCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("HttpConnectionCount")
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

// SetInternetConnectionCount sets the value of InternetConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyInternetConnectionCount(value uint32) (err error) {
	return instance.SetProperty("InternetConnectionCount", (value))
}

// GetInternetConnectionCount gets the value of InternetConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyInternetConnectionCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("InternetConnectionCount")
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

// SetIsBackground sets the value of IsBackground for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyIsBackground(value bool) (err error) {
	return instance.SetProperty("IsBackground", (value))
}

// GetIsBackground gets the value of IsBackground for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyIsBackground() (value bool, err error) {
	retValue, err := instance.GetProperty("IsBackground")
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

// SetIsPinned sets the value of IsPinned for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyIsPinned(value bool) (err error) {
	return instance.SetProperty("IsPinned", (value))
}

// GetIsPinned gets the value of IsPinned for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyIsPinned() (value bool, err error) {
	retValue, err := instance.GetProperty("IsPinned")
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

// SetLanConnectionCount sets the value of LanConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyLanConnectionCount(value uint32) (err error) {
	return instance.SetProperty("LanConnectionCount", (value))
}

// GetLanConnectionCount gets the value of LanConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyLanConnectionCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("LanConnectionCount")
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

// SetLinkLocalConnectionCount sets the value of LinkLocalConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyLinkLocalConnectionCount(value uint32) (err error) {
	return instance.SetProperty("LinkLocalConnectionCount", (value))
}

// GetLinkLocalConnectionCount gets the value of LinkLocalConnectionCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyLinkLocalConnectionCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("LinkLocalConnectionCount")
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

// SetPeerCount sets the value of PeerCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyPeerCount(value uint32) (err error) {
	return instance.SetProperty("PeerCount", (value))
}

// GetPeerCount gets the value of PeerCount for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyPeerCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("PeerCount")
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

// SetPredefinedCallerApplication sets the value of PredefinedCallerApplication for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyPredefinedCallerApplication(value string) (err error) {
	return instance.SetProperty("PredefinedCallerApplication", (value))
}

// GetPredefinedCallerApplication gets the value of PredefinedCallerApplication for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyPredefinedCallerApplication() (value string, err error) {
	retValue, err := instance.GetProperty("PredefinedCallerApplication")
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

// SetSourceURL sets the value of SourceURL for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertySourceURL(value string) (err error) {
	return instance.SetProperty("SourceURL", (value))
}

// GetSourceURL gets the value of SourceURL for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertySourceURL() (value string, err error) {
	retValue, err := instance.GetProperty("SourceURL")
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

// SetStatus sets the value of Status for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyStatus(value DeliveryOptimizationFile_Status) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyStatus() (value DeliveryOptimizationFile_Status, err error) {
	retValue, err := instance.GetProperty("Status")
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

	value = DeliveryOptimizationFile_Status(valuetmp)

	return
}

// SetTotalBytesDownloaded sets the value of TotalBytesDownloaded for the instance
func (instance *MSFT_DeliveryOptimizationFile) SetPropertyTotalBytesDownloaded(value uint64) (err error) {
	return instance.SetProperty("TotalBytesDownloaded", (value))
}

// GetTotalBytesDownloaded gets the value of TotalBytesDownloaded for the instance
func (instance *MSFT_DeliveryOptimizationFile) GetPropertyTotalBytesDownloaded() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalBytesDownloaded")
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

// 11

// <param name="deletePinned" type="bool "></param>
// <param name="fileId" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DeliveryOptimizationFile) Delete( /* IN */ fileId string,
	/* IN */ deletePinned bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Delete", fileId, deletePinned)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// 12

// <param name="fileId" type="string "></param>
// <param name="pinned" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DeliveryOptimizationFile) SetPinned( /* IN */ fileId string,
	/* IN */ pinned bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetPinned", fileId, pinned)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// 13

// <param name="expiration" type="string "></param>
// <param name="fileId" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DeliveryOptimizationFile) SetExpiration( /* IN */ fileId string,
	/* IN */ expiration string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetExpiration", fileId, expiration)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
