// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.directory.LDAP
//
// ////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_acssubnet struct
type ads_acssubnet struct {
	*ds_top

	//
	DS_aCSAllocableRSVPBandwidth int64

	//
	DS_aCSCacheTimeout int32

	//
	DS_aCSDSBMDeadTime int32

	//
	DS_aCSDSBMPriority int32

	//
	DS_aCSDSBMRefresh int32

	//
	DS_aCSEnableACSService bool

	//
	DS_aCSEnableRSVPAccounting bool

	//
	DS_aCSEnableRSVPMessageLogging bool

	//
	DS_aCSEventLogLevel int32

	//
	DS_aCSMaxDurationPerFlow int32

	//
	DS_aCSMaxNoOfAccountFiles int32

	//
	DS_aCSMaxNoOfLogFiles int32

	//
	DS_aCSMaxPeakBandwidth int64

	//
	DS_aCSMaxPeakBandwidthPerFlow int64

	//
	DS_aCSMaxSizeOfRSVPAccountFile int32

	//
	DS_aCSMaxSizeOfRSVPLogFile int32

	//
	DS_aCSMaxTokenRatePerFlow int64

	//
	DS_aCSNonReservedMaxSDUSize int64

	//
	DS_aCSNonReservedMinPolicedSize int64

	//
	DS_aCSNonReservedPeakRate int64

	//
	DS_aCSNonReservedTokenSize int64

	//
	DS_aCSNonReservedTxLimit int64

	//
	DS_aCSNonReservedTxSize int64

	//
	DS_aCSRSVPAccountFilesLocation string

	//
	DS_aCSRSVPLogFilesLocation string

	//
	DS_aCSServerList []string
}

func Newads_acssubnetEx1(instance *cim.WmiInstance) (newInstance *ads_acssubnet, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_acssubnet{
		ds_top: tmp,
	}
	return
}

func Newads_acssubnetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_acssubnet, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_acssubnet{
		ds_top: tmp,
	}
	return
}

// SetDS_aCSAllocableRSVPBandwidth sets the value of DS_aCSAllocableRSVPBandwidth for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSAllocableRSVPBandwidth(value int64) (err error) {
	return instance.SetProperty("DS_aCSAllocableRSVPBandwidth", (value))
}

// GetDS_aCSAllocableRSVPBandwidth gets the value of DS_aCSAllocableRSVPBandwidth for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSAllocableRSVPBandwidth() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSAllocableRSVPBandwidth")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSCacheTimeout sets the value of DS_aCSCacheTimeout for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSCacheTimeout(value int32) (err error) {
	return instance.SetProperty("DS_aCSCacheTimeout", (value))
}

// GetDS_aCSCacheTimeout gets the value of DS_aCSCacheTimeout for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSCacheTimeout() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_aCSCacheTimeout")
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

// SetDS_aCSDSBMDeadTime sets the value of DS_aCSDSBMDeadTime for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSDSBMDeadTime(value int32) (err error) {
	return instance.SetProperty("DS_aCSDSBMDeadTime", (value))
}

// GetDS_aCSDSBMDeadTime gets the value of DS_aCSDSBMDeadTime for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSDSBMDeadTime() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_aCSDSBMDeadTime")
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

// SetDS_aCSDSBMPriority sets the value of DS_aCSDSBMPriority for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSDSBMPriority(value int32) (err error) {
	return instance.SetProperty("DS_aCSDSBMPriority", (value))
}

// GetDS_aCSDSBMPriority gets the value of DS_aCSDSBMPriority for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSDSBMPriority() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_aCSDSBMPriority")
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

// SetDS_aCSDSBMRefresh sets the value of DS_aCSDSBMRefresh for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSDSBMRefresh(value int32) (err error) {
	return instance.SetProperty("DS_aCSDSBMRefresh", (value))
}

// GetDS_aCSDSBMRefresh gets the value of DS_aCSDSBMRefresh for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSDSBMRefresh() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_aCSDSBMRefresh")
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

// SetDS_aCSEnableACSService sets the value of DS_aCSEnableACSService for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSEnableACSService(value bool) (err error) {
	return instance.SetProperty("DS_aCSEnableACSService", (value))
}

// GetDS_aCSEnableACSService gets the value of DS_aCSEnableACSService for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSEnableACSService() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_aCSEnableACSService")
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

// SetDS_aCSEnableRSVPAccounting sets the value of DS_aCSEnableRSVPAccounting for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSEnableRSVPAccounting(value bool) (err error) {
	return instance.SetProperty("DS_aCSEnableRSVPAccounting", (value))
}

// GetDS_aCSEnableRSVPAccounting gets the value of DS_aCSEnableRSVPAccounting for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSEnableRSVPAccounting() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_aCSEnableRSVPAccounting")
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

// SetDS_aCSEnableRSVPMessageLogging sets the value of DS_aCSEnableRSVPMessageLogging for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSEnableRSVPMessageLogging(value bool) (err error) {
	return instance.SetProperty("DS_aCSEnableRSVPMessageLogging", (value))
}

// GetDS_aCSEnableRSVPMessageLogging gets the value of DS_aCSEnableRSVPMessageLogging for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSEnableRSVPMessageLogging() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_aCSEnableRSVPMessageLogging")
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

// SetDS_aCSEventLogLevel sets the value of DS_aCSEventLogLevel for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSEventLogLevel(value int32) (err error) {
	return instance.SetProperty("DS_aCSEventLogLevel", (value))
}

// GetDS_aCSEventLogLevel gets the value of DS_aCSEventLogLevel for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSEventLogLevel() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_aCSEventLogLevel")
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

// SetDS_aCSMaxDurationPerFlow sets the value of DS_aCSMaxDurationPerFlow for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSMaxDurationPerFlow(value int32) (err error) {
	return instance.SetProperty("DS_aCSMaxDurationPerFlow", (value))
}

// GetDS_aCSMaxDurationPerFlow gets the value of DS_aCSMaxDurationPerFlow for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSMaxDurationPerFlow() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_aCSMaxDurationPerFlow")
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

// SetDS_aCSMaxNoOfAccountFiles sets the value of DS_aCSMaxNoOfAccountFiles for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSMaxNoOfAccountFiles(value int32) (err error) {
	return instance.SetProperty("DS_aCSMaxNoOfAccountFiles", (value))
}

// GetDS_aCSMaxNoOfAccountFiles gets the value of DS_aCSMaxNoOfAccountFiles for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSMaxNoOfAccountFiles() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_aCSMaxNoOfAccountFiles")
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

// SetDS_aCSMaxNoOfLogFiles sets the value of DS_aCSMaxNoOfLogFiles for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSMaxNoOfLogFiles(value int32) (err error) {
	return instance.SetProperty("DS_aCSMaxNoOfLogFiles", (value))
}

// GetDS_aCSMaxNoOfLogFiles gets the value of DS_aCSMaxNoOfLogFiles for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSMaxNoOfLogFiles() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_aCSMaxNoOfLogFiles")
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

// SetDS_aCSMaxPeakBandwidth sets the value of DS_aCSMaxPeakBandwidth for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSMaxPeakBandwidth(value int64) (err error) {
	return instance.SetProperty("DS_aCSMaxPeakBandwidth", (value))
}

// GetDS_aCSMaxPeakBandwidth gets the value of DS_aCSMaxPeakBandwidth for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSMaxPeakBandwidth() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSMaxPeakBandwidth")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSMaxPeakBandwidthPerFlow sets the value of DS_aCSMaxPeakBandwidthPerFlow for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSMaxPeakBandwidthPerFlow(value int64) (err error) {
	return instance.SetProperty("DS_aCSMaxPeakBandwidthPerFlow", (value))
}

// GetDS_aCSMaxPeakBandwidthPerFlow gets the value of DS_aCSMaxPeakBandwidthPerFlow for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSMaxPeakBandwidthPerFlow() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSMaxPeakBandwidthPerFlow")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSMaxSizeOfRSVPAccountFile sets the value of DS_aCSMaxSizeOfRSVPAccountFile for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSMaxSizeOfRSVPAccountFile(value int32) (err error) {
	return instance.SetProperty("DS_aCSMaxSizeOfRSVPAccountFile", (value))
}

// GetDS_aCSMaxSizeOfRSVPAccountFile gets the value of DS_aCSMaxSizeOfRSVPAccountFile for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSMaxSizeOfRSVPAccountFile() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_aCSMaxSizeOfRSVPAccountFile")
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

// SetDS_aCSMaxSizeOfRSVPLogFile sets the value of DS_aCSMaxSizeOfRSVPLogFile for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSMaxSizeOfRSVPLogFile(value int32) (err error) {
	return instance.SetProperty("DS_aCSMaxSizeOfRSVPLogFile", (value))
}

// GetDS_aCSMaxSizeOfRSVPLogFile gets the value of DS_aCSMaxSizeOfRSVPLogFile for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSMaxSizeOfRSVPLogFile() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_aCSMaxSizeOfRSVPLogFile")
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

// SetDS_aCSMaxTokenRatePerFlow sets the value of DS_aCSMaxTokenRatePerFlow for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSMaxTokenRatePerFlow(value int64) (err error) {
	return instance.SetProperty("DS_aCSMaxTokenRatePerFlow", (value))
}

// GetDS_aCSMaxTokenRatePerFlow gets the value of DS_aCSMaxTokenRatePerFlow for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSMaxTokenRatePerFlow() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSMaxTokenRatePerFlow")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSNonReservedMaxSDUSize sets the value of DS_aCSNonReservedMaxSDUSize for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSNonReservedMaxSDUSize(value int64) (err error) {
	return instance.SetProperty("DS_aCSNonReservedMaxSDUSize", (value))
}

// GetDS_aCSNonReservedMaxSDUSize gets the value of DS_aCSNonReservedMaxSDUSize for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSNonReservedMaxSDUSize() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSNonReservedMaxSDUSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSNonReservedMinPolicedSize sets the value of DS_aCSNonReservedMinPolicedSize for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSNonReservedMinPolicedSize(value int64) (err error) {
	return instance.SetProperty("DS_aCSNonReservedMinPolicedSize", (value))
}

// GetDS_aCSNonReservedMinPolicedSize gets the value of DS_aCSNonReservedMinPolicedSize for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSNonReservedMinPolicedSize() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSNonReservedMinPolicedSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSNonReservedPeakRate sets the value of DS_aCSNonReservedPeakRate for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSNonReservedPeakRate(value int64) (err error) {
	return instance.SetProperty("DS_aCSNonReservedPeakRate", (value))
}

// GetDS_aCSNonReservedPeakRate gets the value of DS_aCSNonReservedPeakRate for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSNonReservedPeakRate() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSNonReservedPeakRate")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSNonReservedTokenSize sets the value of DS_aCSNonReservedTokenSize for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSNonReservedTokenSize(value int64) (err error) {
	return instance.SetProperty("DS_aCSNonReservedTokenSize", (value))
}

// GetDS_aCSNonReservedTokenSize gets the value of DS_aCSNonReservedTokenSize for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSNonReservedTokenSize() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSNonReservedTokenSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSNonReservedTxLimit sets the value of DS_aCSNonReservedTxLimit for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSNonReservedTxLimit(value int64) (err error) {
	return instance.SetProperty("DS_aCSNonReservedTxLimit", (value))
}

// GetDS_aCSNonReservedTxLimit gets the value of DS_aCSNonReservedTxLimit for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSNonReservedTxLimit() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSNonReservedTxLimit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSNonReservedTxSize sets the value of DS_aCSNonReservedTxSize for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSNonReservedTxSize(value int64) (err error) {
	return instance.SetProperty("DS_aCSNonReservedTxSize", (value))
}

// GetDS_aCSNonReservedTxSize gets the value of DS_aCSNonReservedTxSize for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSNonReservedTxSize() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSNonReservedTxSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSRSVPAccountFilesLocation sets the value of DS_aCSRSVPAccountFilesLocation for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSRSVPAccountFilesLocation(value string) (err error) {
	return instance.SetProperty("DS_aCSRSVPAccountFilesLocation", (value))
}

// GetDS_aCSRSVPAccountFilesLocation gets the value of DS_aCSRSVPAccountFilesLocation for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSRSVPAccountFilesLocation() (value string, err error) {
	retValue, err := instance.GetProperty("DS_aCSRSVPAccountFilesLocation")
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

// SetDS_aCSRSVPLogFilesLocation sets the value of DS_aCSRSVPLogFilesLocation for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSRSVPLogFilesLocation(value string) (err error) {
	return instance.SetProperty("DS_aCSRSVPLogFilesLocation", (value))
}

// GetDS_aCSRSVPLogFilesLocation gets the value of DS_aCSRSVPLogFilesLocation for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSRSVPLogFilesLocation() (value string, err error) {
	retValue, err := instance.GetProperty("DS_aCSRSVPLogFilesLocation")
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

// SetDS_aCSServerList sets the value of DS_aCSServerList for the instance
func (instance *ads_acssubnet) SetPropertyDS_aCSServerList(value []string) (err error) {
	return instance.SetProperty("DS_aCSServerList", (value))
}

// GetDS_aCSServerList gets the value of DS_aCSServerList for the instance
func (instance *ads_acssubnet) GetPropertyDS_aCSServerList() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_aCSServerList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
