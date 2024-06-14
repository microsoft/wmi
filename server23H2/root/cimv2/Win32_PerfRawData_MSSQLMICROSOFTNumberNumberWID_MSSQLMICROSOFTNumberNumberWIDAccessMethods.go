// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods struct
type Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods struct {
	*Win32_PerfRawData

	//
	AUcleanupbatchesPersec uint64

	//
	AUcleanupsPersec uint64

	//
	ByreferenceLobCreateCount uint64

	//
	ByreferenceLobUseCount uint64

	//
	CountLobReadahead uint64

	//
	CountPullInRow uint64

	//
	CountPushOffRow uint64

	//
	DeferreddroppedAUs uint64

	//
	DeferredDroppedrowsets uint64

	//
	DroppedrowsetcleanupsPersec uint64

	//
	DroppedrowsetsskippedPersec uint64

	//
	ExtentDeallocationsPersec uint64

	//
	ExtentsAllocatedPersec uint64

	//
	FailedAUcleanupbatchesPersec uint64

	//
	Failedleafpagecookie uint64

	//
	Failedtreepagecookie uint64

	//
	ForwardedRecordsPersec uint64

	//
	FreeSpacePageFetchesPersec uint64

	//
	FreeSpaceScansPersec uint64

	//
	FullScansPersec uint64

	//
	IndexSearchesPersec uint64

	//
	InSysXactwaitsPersec uint64

	//
	LobHandleCreateCount uint64

	//
	LobHandleDestroyCount uint64

	//
	LobSSProviderCreateCount uint64

	//
	LobSSProviderDestroyCount uint64

	//
	LobSSProviderTruncationCount uint64

	//
	MixedpageallocationsPersec uint64

	//
	PagecompressionattemptsPersec uint64

	//
	PageDeallocationsPersec uint64

	//
	PagesAllocatedPersec uint64

	//
	PagescompressedPersec uint64

	//
	PageSplitsPersec uint64

	//
	ProbeScansPersec uint64

	//
	RangeScansPersec uint64

	//
	ScanPointRevalidationsPersec uint64

	//
	SkippedGhostedRecordsPersec uint64

	//
	TableLockEscalationsPersec uint64

	//
	Usedleafpagecookie uint64

	//
	Usedtreepagecookie uint64

	//
	WorkfilesCreatedPersec uint64

	//
	WorktablesCreatedPersec uint64

	//
	WorktablesFromCacheRatio uint64

	//
	WorktablesFromCacheRatio_Base uint64
}

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethodsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethodsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetAUcleanupbatchesPersec sets the value of AUcleanupbatchesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyAUcleanupbatchesPersec(value uint64) (err error) {
	return instance.SetProperty("AUcleanupbatchesPersec", (value))
}

// GetAUcleanupbatchesPersec gets the value of AUcleanupbatchesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyAUcleanupbatchesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("AUcleanupbatchesPersec")
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

// SetAUcleanupsPersec sets the value of AUcleanupsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyAUcleanupsPersec(value uint64) (err error) {
	return instance.SetProperty("AUcleanupsPersec", (value))
}

// GetAUcleanupsPersec gets the value of AUcleanupsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyAUcleanupsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("AUcleanupsPersec")
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

// SetByreferenceLobCreateCount sets the value of ByreferenceLobCreateCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyByreferenceLobCreateCount(value uint64) (err error) {
	return instance.SetProperty("ByreferenceLobCreateCount", (value))
}

// GetByreferenceLobCreateCount gets the value of ByreferenceLobCreateCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyByreferenceLobCreateCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("ByreferenceLobCreateCount")
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

// SetByreferenceLobUseCount sets the value of ByreferenceLobUseCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyByreferenceLobUseCount(value uint64) (err error) {
	return instance.SetProperty("ByreferenceLobUseCount", (value))
}

// GetByreferenceLobUseCount gets the value of ByreferenceLobUseCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyByreferenceLobUseCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("ByreferenceLobUseCount")
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

// SetCountLobReadahead sets the value of CountLobReadahead for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyCountLobReadahead(value uint64) (err error) {
	return instance.SetProperty("CountLobReadahead", (value))
}

// GetCountLobReadahead gets the value of CountLobReadahead for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyCountLobReadahead() (value uint64, err error) {
	retValue, err := instance.GetProperty("CountLobReadahead")
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

// SetCountPullInRow sets the value of CountPullInRow for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyCountPullInRow(value uint64) (err error) {
	return instance.SetProperty("CountPullInRow", (value))
}

// GetCountPullInRow gets the value of CountPullInRow for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyCountPullInRow() (value uint64, err error) {
	retValue, err := instance.GetProperty("CountPullInRow")
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

// SetCountPushOffRow sets the value of CountPushOffRow for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyCountPushOffRow(value uint64) (err error) {
	return instance.SetProperty("CountPushOffRow", (value))
}

// GetCountPushOffRow gets the value of CountPushOffRow for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyCountPushOffRow() (value uint64, err error) {
	retValue, err := instance.GetProperty("CountPushOffRow")
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

// SetDeferreddroppedAUs sets the value of DeferreddroppedAUs for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyDeferreddroppedAUs(value uint64) (err error) {
	return instance.SetProperty("DeferreddroppedAUs", (value))
}

// GetDeferreddroppedAUs gets the value of DeferreddroppedAUs for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyDeferreddroppedAUs() (value uint64, err error) {
	retValue, err := instance.GetProperty("DeferreddroppedAUs")
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

// SetDeferredDroppedrowsets sets the value of DeferredDroppedrowsets for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyDeferredDroppedrowsets(value uint64) (err error) {
	return instance.SetProperty("DeferredDroppedrowsets", (value))
}

// GetDeferredDroppedrowsets gets the value of DeferredDroppedrowsets for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyDeferredDroppedrowsets() (value uint64, err error) {
	retValue, err := instance.GetProperty("DeferredDroppedrowsets")
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

// SetDroppedrowsetcleanupsPersec sets the value of DroppedrowsetcleanupsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyDroppedrowsetcleanupsPersec(value uint64) (err error) {
	return instance.SetProperty("DroppedrowsetcleanupsPersec", (value))
}

// GetDroppedrowsetcleanupsPersec gets the value of DroppedrowsetcleanupsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyDroppedrowsetcleanupsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DroppedrowsetcleanupsPersec")
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

// SetDroppedrowsetsskippedPersec sets the value of DroppedrowsetsskippedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyDroppedrowsetsskippedPersec(value uint64) (err error) {
	return instance.SetProperty("DroppedrowsetsskippedPersec", (value))
}

// GetDroppedrowsetsskippedPersec gets the value of DroppedrowsetsskippedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyDroppedrowsetsskippedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DroppedrowsetsskippedPersec")
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

// SetExtentDeallocationsPersec sets the value of ExtentDeallocationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyExtentDeallocationsPersec(value uint64) (err error) {
	return instance.SetProperty("ExtentDeallocationsPersec", (value))
}

// GetExtentDeallocationsPersec gets the value of ExtentDeallocationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyExtentDeallocationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ExtentDeallocationsPersec")
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

// SetExtentsAllocatedPersec sets the value of ExtentsAllocatedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyExtentsAllocatedPersec(value uint64) (err error) {
	return instance.SetProperty("ExtentsAllocatedPersec", (value))
}

// GetExtentsAllocatedPersec gets the value of ExtentsAllocatedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyExtentsAllocatedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ExtentsAllocatedPersec")
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

// SetFailedAUcleanupbatchesPersec sets the value of FailedAUcleanupbatchesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyFailedAUcleanupbatchesPersec(value uint64) (err error) {
	return instance.SetProperty("FailedAUcleanupbatchesPersec", (value))
}

// GetFailedAUcleanupbatchesPersec gets the value of FailedAUcleanupbatchesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyFailedAUcleanupbatchesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FailedAUcleanupbatchesPersec")
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

// SetFailedleafpagecookie sets the value of Failedleafpagecookie for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyFailedleafpagecookie(value uint64) (err error) {
	return instance.SetProperty("Failedleafpagecookie", (value))
}

// GetFailedleafpagecookie gets the value of Failedleafpagecookie for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyFailedleafpagecookie() (value uint64, err error) {
	retValue, err := instance.GetProperty("Failedleafpagecookie")
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

// SetFailedtreepagecookie sets the value of Failedtreepagecookie for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyFailedtreepagecookie(value uint64) (err error) {
	return instance.SetProperty("Failedtreepagecookie", (value))
}

// GetFailedtreepagecookie gets the value of Failedtreepagecookie for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyFailedtreepagecookie() (value uint64, err error) {
	retValue, err := instance.GetProperty("Failedtreepagecookie")
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

// SetForwardedRecordsPersec sets the value of ForwardedRecordsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyForwardedRecordsPersec(value uint64) (err error) {
	return instance.SetProperty("ForwardedRecordsPersec", (value))
}

// GetForwardedRecordsPersec gets the value of ForwardedRecordsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyForwardedRecordsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ForwardedRecordsPersec")
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

// SetFreeSpacePageFetchesPersec sets the value of FreeSpacePageFetchesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyFreeSpacePageFetchesPersec(value uint64) (err error) {
	return instance.SetProperty("FreeSpacePageFetchesPersec", (value))
}

// GetFreeSpacePageFetchesPersec gets the value of FreeSpacePageFetchesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyFreeSpacePageFetchesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FreeSpacePageFetchesPersec")
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

// SetFreeSpaceScansPersec sets the value of FreeSpaceScansPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyFreeSpaceScansPersec(value uint64) (err error) {
	return instance.SetProperty("FreeSpaceScansPersec", (value))
}

// GetFreeSpaceScansPersec gets the value of FreeSpaceScansPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyFreeSpaceScansPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FreeSpaceScansPersec")
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

// SetFullScansPersec sets the value of FullScansPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyFullScansPersec(value uint64) (err error) {
	return instance.SetProperty("FullScansPersec", (value))
}

// GetFullScansPersec gets the value of FullScansPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyFullScansPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FullScansPersec")
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

// SetIndexSearchesPersec sets the value of IndexSearchesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyIndexSearchesPersec(value uint64) (err error) {
	return instance.SetProperty("IndexSearchesPersec", (value))
}

// GetIndexSearchesPersec gets the value of IndexSearchesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyIndexSearchesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IndexSearchesPersec")
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

// SetInSysXactwaitsPersec sets the value of InSysXactwaitsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyInSysXactwaitsPersec(value uint64) (err error) {
	return instance.SetProperty("InSysXactwaitsPersec", (value))
}

// GetInSysXactwaitsPersec gets the value of InSysXactwaitsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyInSysXactwaitsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("InSysXactwaitsPersec")
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

// SetLobHandleCreateCount sets the value of LobHandleCreateCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyLobHandleCreateCount(value uint64) (err error) {
	return instance.SetProperty("LobHandleCreateCount", (value))
}

// GetLobHandleCreateCount gets the value of LobHandleCreateCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyLobHandleCreateCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("LobHandleCreateCount")
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

// SetLobHandleDestroyCount sets the value of LobHandleDestroyCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyLobHandleDestroyCount(value uint64) (err error) {
	return instance.SetProperty("LobHandleDestroyCount", (value))
}

// GetLobHandleDestroyCount gets the value of LobHandleDestroyCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyLobHandleDestroyCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("LobHandleDestroyCount")
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

// SetLobSSProviderCreateCount sets the value of LobSSProviderCreateCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyLobSSProviderCreateCount(value uint64) (err error) {
	return instance.SetProperty("LobSSProviderCreateCount", (value))
}

// GetLobSSProviderCreateCount gets the value of LobSSProviderCreateCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyLobSSProviderCreateCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("LobSSProviderCreateCount")
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

// SetLobSSProviderDestroyCount sets the value of LobSSProviderDestroyCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyLobSSProviderDestroyCount(value uint64) (err error) {
	return instance.SetProperty("LobSSProviderDestroyCount", (value))
}

// GetLobSSProviderDestroyCount gets the value of LobSSProviderDestroyCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyLobSSProviderDestroyCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("LobSSProviderDestroyCount")
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

// SetLobSSProviderTruncationCount sets the value of LobSSProviderTruncationCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyLobSSProviderTruncationCount(value uint64) (err error) {
	return instance.SetProperty("LobSSProviderTruncationCount", (value))
}

// GetLobSSProviderTruncationCount gets the value of LobSSProviderTruncationCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyLobSSProviderTruncationCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("LobSSProviderTruncationCount")
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

// SetMixedpageallocationsPersec sets the value of MixedpageallocationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyMixedpageallocationsPersec(value uint64) (err error) {
	return instance.SetProperty("MixedpageallocationsPersec", (value))
}

// GetMixedpageallocationsPersec gets the value of MixedpageallocationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyMixedpageallocationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MixedpageallocationsPersec")
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

// SetPagecompressionattemptsPersec sets the value of PagecompressionattemptsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyPagecompressionattemptsPersec(value uint64) (err error) {
	return instance.SetProperty("PagecompressionattemptsPersec", (value))
}

// GetPagecompressionattemptsPersec gets the value of PagecompressionattemptsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyPagecompressionattemptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PagecompressionattemptsPersec")
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

// SetPageDeallocationsPersec sets the value of PageDeallocationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyPageDeallocationsPersec(value uint64) (err error) {
	return instance.SetProperty("PageDeallocationsPersec", (value))
}

// GetPageDeallocationsPersec gets the value of PageDeallocationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyPageDeallocationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageDeallocationsPersec")
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

// SetPagesAllocatedPersec sets the value of PagesAllocatedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyPagesAllocatedPersec(value uint64) (err error) {
	return instance.SetProperty("PagesAllocatedPersec", (value))
}

// GetPagesAllocatedPersec gets the value of PagesAllocatedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyPagesAllocatedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PagesAllocatedPersec")
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

// SetPagescompressedPersec sets the value of PagescompressedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyPagescompressedPersec(value uint64) (err error) {
	return instance.SetProperty("PagescompressedPersec", (value))
}

// GetPagescompressedPersec gets the value of PagescompressedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyPagescompressedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PagescompressedPersec")
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

// SetPageSplitsPersec sets the value of PageSplitsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyPageSplitsPersec(value uint64) (err error) {
	return instance.SetProperty("PageSplitsPersec", (value))
}

// GetPageSplitsPersec gets the value of PageSplitsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyPageSplitsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageSplitsPersec")
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

// SetProbeScansPersec sets the value of ProbeScansPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyProbeScansPersec(value uint64) (err error) {
	return instance.SetProperty("ProbeScansPersec", (value))
}

// GetProbeScansPersec gets the value of ProbeScansPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyProbeScansPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ProbeScansPersec")
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

// SetRangeScansPersec sets the value of RangeScansPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyRangeScansPersec(value uint64) (err error) {
	return instance.SetProperty("RangeScansPersec", (value))
}

// GetRangeScansPersec gets the value of RangeScansPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyRangeScansPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("RangeScansPersec")
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

// SetScanPointRevalidationsPersec sets the value of ScanPointRevalidationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyScanPointRevalidationsPersec(value uint64) (err error) {
	return instance.SetProperty("ScanPointRevalidationsPersec", (value))
}

// GetScanPointRevalidationsPersec gets the value of ScanPointRevalidationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyScanPointRevalidationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ScanPointRevalidationsPersec")
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

// SetSkippedGhostedRecordsPersec sets the value of SkippedGhostedRecordsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertySkippedGhostedRecordsPersec(value uint64) (err error) {
	return instance.SetProperty("SkippedGhostedRecordsPersec", (value))
}

// GetSkippedGhostedRecordsPersec gets the value of SkippedGhostedRecordsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertySkippedGhostedRecordsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SkippedGhostedRecordsPersec")
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

// SetTableLockEscalationsPersec sets the value of TableLockEscalationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyTableLockEscalationsPersec(value uint64) (err error) {
	return instance.SetProperty("TableLockEscalationsPersec", (value))
}

// GetTableLockEscalationsPersec gets the value of TableLockEscalationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyTableLockEscalationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TableLockEscalationsPersec")
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

// SetUsedleafpagecookie sets the value of Usedleafpagecookie for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyUsedleafpagecookie(value uint64) (err error) {
	return instance.SetProperty("Usedleafpagecookie", (value))
}

// GetUsedleafpagecookie gets the value of Usedleafpagecookie for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyUsedleafpagecookie() (value uint64, err error) {
	retValue, err := instance.GetProperty("Usedleafpagecookie")
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

// SetUsedtreepagecookie sets the value of Usedtreepagecookie for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyUsedtreepagecookie(value uint64) (err error) {
	return instance.SetProperty("Usedtreepagecookie", (value))
}

// GetUsedtreepagecookie gets the value of Usedtreepagecookie for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyUsedtreepagecookie() (value uint64, err error) {
	retValue, err := instance.GetProperty("Usedtreepagecookie")
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

// SetWorkfilesCreatedPersec sets the value of WorkfilesCreatedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyWorkfilesCreatedPersec(value uint64) (err error) {
	return instance.SetProperty("WorkfilesCreatedPersec", (value))
}

// GetWorkfilesCreatedPersec gets the value of WorkfilesCreatedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyWorkfilesCreatedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WorkfilesCreatedPersec")
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

// SetWorktablesCreatedPersec sets the value of WorktablesCreatedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyWorktablesCreatedPersec(value uint64) (err error) {
	return instance.SetProperty("WorktablesCreatedPersec", (value))
}

// GetWorktablesCreatedPersec gets the value of WorktablesCreatedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyWorktablesCreatedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WorktablesCreatedPersec")
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

// SetWorktablesFromCacheRatio sets the value of WorktablesFromCacheRatio for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyWorktablesFromCacheRatio(value uint64) (err error) {
	return instance.SetProperty("WorktablesFromCacheRatio", (value))
}

// GetWorktablesFromCacheRatio gets the value of WorktablesFromCacheRatio for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyWorktablesFromCacheRatio() (value uint64, err error) {
	retValue, err := instance.GetProperty("WorktablesFromCacheRatio")
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

// SetWorktablesFromCacheRatio_Base sets the value of WorktablesFromCacheRatio_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) SetPropertyWorktablesFromCacheRatio_Base(value uint64) (err error) {
	return instance.SetProperty("WorktablesFromCacheRatio_Base", (value))
}

// GetWorktablesFromCacheRatio_Base gets the value of WorktablesFromCacheRatio_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDAccessMethods) GetPropertyWorktablesFromCacheRatio_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("WorktablesFromCacheRatio_Base")
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
