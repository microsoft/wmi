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

// MDM_Policy_Result01_DeliveryOptimization02 struct
type MDM_Policy_Result01_DeliveryOptimization02 struct {
	*cim.WmiInstance

	//
	DOAbsoluteMaxCacheSize int32

	//
	DOAllowVPNPeerCaching int32

	//
	DOCacheHost string

	//
	DOCacheHostSource int32

	//
	DODelayBackgroundDownloadFromHttp int32

	//
	DODelayCacheServerFallbackBackground int32

	//
	DODelayCacheServerFallbackForeground int32

	//
	DODelayForegroundDownloadFromHttp int32

	//
	DODownloadMode int32

	//
	DOGroupId string

	//
	DOGroupIdSource int32

	//
	DOMaxBackgroundDownloadBandwidth int32

	//
	DOMaxCacheAge int32

	//
	DOMaxCacheSize int32

	//
	DOMaxForegroundDownloadBandwidth int32

	//
	DOMinBackgroundQos int32

	//
	DOMinBatteryPercentageAllowedToUpload int32

	//
	DOMinDiskSizeAllowedToPeer int32

	//
	DOMinFileSizeToCache int32

	//
	DOMinRAMAllowedToPeer int32

	//
	DOModifyCacheDrive string

	//
	DOMonthlyUploadDataCap int32

	//
	DOPercentageMaxBackgroundBandwidth int32

	//
	DOPercentageMaxForegroundBandwidth int32

	//
	DORestrictPeerSelectionBy int32

	//
	DOSetHoursToLimitBackgroundDownloadBandwidth string

	//
	DOSetHoursToLimitForegroundDownloadBandwidth string

	//
	InstanceID string

	//
	ParentID string
}

func NewMDM_Policy_Result01_DeliveryOptimization02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_DeliveryOptimization02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_DeliveryOptimization02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_DeliveryOptimization02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_DeliveryOptimization02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_DeliveryOptimization02{
		WmiInstance: tmp,
	}
	return
}

// SetDOAbsoluteMaxCacheSize sets the value of DOAbsoluteMaxCacheSize for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOAbsoluteMaxCacheSize(value int32) (err error) {
	return instance.SetProperty("DOAbsoluteMaxCacheSize", (value))
}

// GetDOAbsoluteMaxCacheSize gets the value of DOAbsoluteMaxCacheSize for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOAbsoluteMaxCacheSize() (value int32, err error) {
	retValue, err := instance.GetProperty("DOAbsoluteMaxCacheSize")
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

// SetDOAllowVPNPeerCaching sets the value of DOAllowVPNPeerCaching for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOAllowVPNPeerCaching(value int32) (err error) {
	return instance.SetProperty("DOAllowVPNPeerCaching", (value))
}

// GetDOAllowVPNPeerCaching gets the value of DOAllowVPNPeerCaching for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOAllowVPNPeerCaching() (value int32, err error) {
	retValue, err := instance.GetProperty("DOAllowVPNPeerCaching")
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

// SetDOCacheHost sets the value of DOCacheHost for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOCacheHost(value string) (err error) {
	return instance.SetProperty("DOCacheHost", (value))
}

// GetDOCacheHost gets the value of DOCacheHost for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOCacheHost() (value string, err error) {
	retValue, err := instance.GetProperty("DOCacheHost")
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

// SetDOCacheHostSource sets the value of DOCacheHostSource for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOCacheHostSource(value int32) (err error) {
	return instance.SetProperty("DOCacheHostSource", (value))
}

// GetDOCacheHostSource gets the value of DOCacheHostSource for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOCacheHostSource() (value int32, err error) {
	retValue, err := instance.GetProperty("DOCacheHostSource")
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

// SetDODelayBackgroundDownloadFromHttp sets the value of DODelayBackgroundDownloadFromHttp for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDODelayBackgroundDownloadFromHttp(value int32) (err error) {
	return instance.SetProperty("DODelayBackgroundDownloadFromHttp", (value))
}

// GetDODelayBackgroundDownloadFromHttp gets the value of DODelayBackgroundDownloadFromHttp for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDODelayBackgroundDownloadFromHttp() (value int32, err error) {
	retValue, err := instance.GetProperty("DODelayBackgroundDownloadFromHttp")
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

// SetDODelayCacheServerFallbackBackground sets the value of DODelayCacheServerFallbackBackground for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDODelayCacheServerFallbackBackground(value int32) (err error) {
	return instance.SetProperty("DODelayCacheServerFallbackBackground", (value))
}

// GetDODelayCacheServerFallbackBackground gets the value of DODelayCacheServerFallbackBackground for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDODelayCacheServerFallbackBackground() (value int32, err error) {
	retValue, err := instance.GetProperty("DODelayCacheServerFallbackBackground")
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

// SetDODelayCacheServerFallbackForeground sets the value of DODelayCacheServerFallbackForeground for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDODelayCacheServerFallbackForeground(value int32) (err error) {
	return instance.SetProperty("DODelayCacheServerFallbackForeground", (value))
}

// GetDODelayCacheServerFallbackForeground gets the value of DODelayCacheServerFallbackForeground for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDODelayCacheServerFallbackForeground() (value int32, err error) {
	retValue, err := instance.GetProperty("DODelayCacheServerFallbackForeground")
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

// SetDODelayForegroundDownloadFromHttp sets the value of DODelayForegroundDownloadFromHttp for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDODelayForegroundDownloadFromHttp(value int32) (err error) {
	return instance.SetProperty("DODelayForegroundDownloadFromHttp", (value))
}

// GetDODelayForegroundDownloadFromHttp gets the value of DODelayForegroundDownloadFromHttp for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDODelayForegroundDownloadFromHttp() (value int32, err error) {
	retValue, err := instance.GetProperty("DODelayForegroundDownloadFromHttp")
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

// SetDODownloadMode sets the value of DODownloadMode for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDODownloadMode(value int32) (err error) {
	return instance.SetProperty("DODownloadMode", (value))
}

// GetDODownloadMode gets the value of DODownloadMode for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDODownloadMode() (value int32, err error) {
	retValue, err := instance.GetProperty("DODownloadMode")
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

// SetDOGroupId sets the value of DOGroupId for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOGroupId(value string) (err error) {
	return instance.SetProperty("DOGroupId", (value))
}

// GetDOGroupId gets the value of DOGroupId for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOGroupId() (value string, err error) {
	retValue, err := instance.GetProperty("DOGroupId")
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

// SetDOGroupIdSource sets the value of DOGroupIdSource for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOGroupIdSource(value int32) (err error) {
	return instance.SetProperty("DOGroupIdSource", (value))
}

// GetDOGroupIdSource gets the value of DOGroupIdSource for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOGroupIdSource() (value int32, err error) {
	retValue, err := instance.GetProperty("DOGroupIdSource")
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

// SetDOMaxBackgroundDownloadBandwidth sets the value of DOMaxBackgroundDownloadBandwidth for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOMaxBackgroundDownloadBandwidth(value int32) (err error) {
	return instance.SetProperty("DOMaxBackgroundDownloadBandwidth", (value))
}

// GetDOMaxBackgroundDownloadBandwidth gets the value of DOMaxBackgroundDownloadBandwidth for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOMaxBackgroundDownloadBandwidth() (value int32, err error) {
	retValue, err := instance.GetProperty("DOMaxBackgroundDownloadBandwidth")
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

// SetDOMaxCacheAge sets the value of DOMaxCacheAge for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOMaxCacheAge(value int32) (err error) {
	return instance.SetProperty("DOMaxCacheAge", (value))
}

// GetDOMaxCacheAge gets the value of DOMaxCacheAge for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOMaxCacheAge() (value int32, err error) {
	retValue, err := instance.GetProperty("DOMaxCacheAge")
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

// SetDOMaxCacheSize sets the value of DOMaxCacheSize for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOMaxCacheSize(value int32) (err error) {
	return instance.SetProperty("DOMaxCacheSize", (value))
}

// GetDOMaxCacheSize gets the value of DOMaxCacheSize for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOMaxCacheSize() (value int32, err error) {
	retValue, err := instance.GetProperty("DOMaxCacheSize")
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

// SetDOMaxForegroundDownloadBandwidth sets the value of DOMaxForegroundDownloadBandwidth for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOMaxForegroundDownloadBandwidth(value int32) (err error) {
	return instance.SetProperty("DOMaxForegroundDownloadBandwidth", (value))
}

// GetDOMaxForegroundDownloadBandwidth gets the value of DOMaxForegroundDownloadBandwidth for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOMaxForegroundDownloadBandwidth() (value int32, err error) {
	retValue, err := instance.GetProperty("DOMaxForegroundDownloadBandwidth")
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

// SetDOMinBackgroundQos sets the value of DOMinBackgroundQos for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOMinBackgroundQos(value int32) (err error) {
	return instance.SetProperty("DOMinBackgroundQos", (value))
}

// GetDOMinBackgroundQos gets the value of DOMinBackgroundQos for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOMinBackgroundQos() (value int32, err error) {
	retValue, err := instance.GetProperty("DOMinBackgroundQos")
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

// SetDOMinBatteryPercentageAllowedToUpload sets the value of DOMinBatteryPercentageAllowedToUpload for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOMinBatteryPercentageAllowedToUpload(value int32) (err error) {
	return instance.SetProperty("DOMinBatteryPercentageAllowedToUpload", (value))
}

// GetDOMinBatteryPercentageAllowedToUpload gets the value of DOMinBatteryPercentageAllowedToUpload for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOMinBatteryPercentageAllowedToUpload() (value int32, err error) {
	retValue, err := instance.GetProperty("DOMinBatteryPercentageAllowedToUpload")
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

// SetDOMinDiskSizeAllowedToPeer sets the value of DOMinDiskSizeAllowedToPeer for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOMinDiskSizeAllowedToPeer(value int32) (err error) {
	return instance.SetProperty("DOMinDiskSizeAllowedToPeer", (value))
}

// GetDOMinDiskSizeAllowedToPeer gets the value of DOMinDiskSizeAllowedToPeer for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOMinDiskSizeAllowedToPeer() (value int32, err error) {
	retValue, err := instance.GetProperty("DOMinDiskSizeAllowedToPeer")
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

// SetDOMinFileSizeToCache sets the value of DOMinFileSizeToCache for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOMinFileSizeToCache(value int32) (err error) {
	return instance.SetProperty("DOMinFileSizeToCache", (value))
}

// GetDOMinFileSizeToCache gets the value of DOMinFileSizeToCache for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOMinFileSizeToCache() (value int32, err error) {
	retValue, err := instance.GetProperty("DOMinFileSizeToCache")
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

// SetDOMinRAMAllowedToPeer sets the value of DOMinRAMAllowedToPeer for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOMinRAMAllowedToPeer(value int32) (err error) {
	return instance.SetProperty("DOMinRAMAllowedToPeer", (value))
}

// GetDOMinRAMAllowedToPeer gets the value of DOMinRAMAllowedToPeer for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOMinRAMAllowedToPeer() (value int32, err error) {
	retValue, err := instance.GetProperty("DOMinRAMAllowedToPeer")
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

// SetDOModifyCacheDrive sets the value of DOModifyCacheDrive for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOModifyCacheDrive(value string) (err error) {
	return instance.SetProperty("DOModifyCacheDrive", (value))
}

// GetDOModifyCacheDrive gets the value of DOModifyCacheDrive for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOModifyCacheDrive() (value string, err error) {
	retValue, err := instance.GetProperty("DOModifyCacheDrive")
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

// SetDOMonthlyUploadDataCap sets the value of DOMonthlyUploadDataCap for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOMonthlyUploadDataCap(value int32) (err error) {
	return instance.SetProperty("DOMonthlyUploadDataCap", (value))
}

// GetDOMonthlyUploadDataCap gets the value of DOMonthlyUploadDataCap for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOMonthlyUploadDataCap() (value int32, err error) {
	retValue, err := instance.GetProperty("DOMonthlyUploadDataCap")
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

// SetDOPercentageMaxBackgroundBandwidth sets the value of DOPercentageMaxBackgroundBandwidth for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOPercentageMaxBackgroundBandwidth(value int32) (err error) {
	return instance.SetProperty("DOPercentageMaxBackgroundBandwidth", (value))
}

// GetDOPercentageMaxBackgroundBandwidth gets the value of DOPercentageMaxBackgroundBandwidth for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOPercentageMaxBackgroundBandwidth() (value int32, err error) {
	retValue, err := instance.GetProperty("DOPercentageMaxBackgroundBandwidth")
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

// SetDOPercentageMaxForegroundBandwidth sets the value of DOPercentageMaxForegroundBandwidth for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOPercentageMaxForegroundBandwidth(value int32) (err error) {
	return instance.SetProperty("DOPercentageMaxForegroundBandwidth", (value))
}

// GetDOPercentageMaxForegroundBandwidth gets the value of DOPercentageMaxForegroundBandwidth for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOPercentageMaxForegroundBandwidth() (value int32, err error) {
	retValue, err := instance.GetProperty("DOPercentageMaxForegroundBandwidth")
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

// SetDORestrictPeerSelectionBy sets the value of DORestrictPeerSelectionBy for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDORestrictPeerSelectionBy(value int32) (err error) {
	return instance.SetProperty("DORestrictPeerSelectionBy", (value))
}

// GetDORestrictPeerSelectionBy gets the value of DORestrictPeerSelectionBy for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDORestrictPeerSelectionBy() (value int32, err error) {
	retValue, err := instance.GetProperty("DORestrictPeerSelectionBy")
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

// SetDOSetHoursToLimitBackgroundDownloadBandwidth sets the value of DOSetHoursToLimitBackgroundDownloadBandwidth for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOSetHoursToLimitBackgroundDownloadBandwidth(value string) (err error) {
	return instance.SetProperty("DOSetHoursToLimitBackgroundDownloadBandwidth", (value))
}

// GetDOSetHoursToLimitBackgroundDownloadBandwidth gets the value of DOSetHoursToLimitBackgroundDownloadBandwidth for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOSetHoursToLimitBackgroundDownloadBandwidth() (value string, err error) {
	retValue, err := instance.GetProperty("DOSetHoursToLimitBackgroundDownloadBandwidth")
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

// SetDOSetHoursToLimitForegroundDownloadBandwidth sets the value of DOSetHoursToLimitForegroundDownloadBandwidth for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyDOSetHoursToLimitForegroundDownloadBandwidth(value string) (err error) {
	return instance.SetProperty("DOSetHoursToLimitForegroundDownloadBandwidth", (value))
}

// GetDOSetHoursToLimitForegroundDownloadBandwidth gets the value of DOSetHoursToLimitForegroundDownloadBandwidth for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyDOSetHoursToLimitForegroundDownloadBandwidth() (value string, err error) {
	retValue, err := instance.GetProperty("DOSetHoursToLimitForegroundDownloadBandwidth")
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
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyInstanceID() (value string, err error) {
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_DeliveryOptimization02) GetPropertyParentID() (value string, err error) {
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
