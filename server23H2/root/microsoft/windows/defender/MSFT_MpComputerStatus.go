// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.Defender
//////////////////////////////////////////////
package defender

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_MpComputerStatus struct
type MSFT_MpComputerStatus struct {
	*BaseStatus

	//
	AMEngineVersion string

	//
	AMProductVersion string

	//
	AMRunningMode string

	//
	AMServiceEnabled bool

	//
	AMServiceVersion string

	//
	AntispywareEnabled bool

	//
	AntispywareSignatureAge uint32

	//
	AntispywareSignatureLastUpdated string

	//
	AntispywareSignatureVersion string

	//
	AntivirusEnabled bool

	//
	AntivirusSignatureAge uint32

	//
	AntivirusSignatureLastUpdated string

	//
	AntivirusSignatureVersion string

	//
	BehaviorMonitorEnabled bool

	//
	ComputerID string

	//
	ComputerState uint32

	//
	DefenderSignaturesOutOfDate bool

	//
	DeviceControlDefaultEnforcement string

	//
	DeviceControlPoliciesLastUpdated string

	//
	DeviceControlState string

	//
	FullScanAge uint32

	//
	FullScanEndTime string

	//
	FullScanOverdue bool

	//
	FullScanRequired bool

	//
	FullScanSignatureVersion string

	//
	FullScanStartTime string

	//
	InitializationProgress string

	//
	IoavProtectionEnabled bool

	//
	IsTamperProtected bool

	//
	IsVirtualMachine bool

	//
	LastFullScanSource uint8

	//
	LastQuickScanSource uint8

	//
	NISEnabled bool

	//
	NISEngineVersion string

	//
	NISSignatureAge uint32

	//
	NISSignatureLastUpdated string

	//
	NISSignatureVersion string

	//
	OnAccessProtectionEnabled bool

	//
	ProductStatus uint32

	//
	QuickScanAge uint32

	//
	QuickScanEndTime string

	//
	QuickScanOverdue bool

	//
	QuickScanSignatureVersion string

	//
	QuickScanStartTime string

	//
	RealTimeProtectionEnabled bool

	//
	RealTimeScanDirection uint8

	//
	RebootRequired bool

	//
	SmartAppControlExpiration string

	//
	SmartAppControlState string

	//
	TamperProtectionSource string

	//
	TDTCapable string

	//
	TDTMode string

	//
	TDTSiloType string

	//
	TDTStatus string

	//
	TDTTelemetry string

	//
	TroubleShootingDailyMaxQuota string

	//
	TroubleShootingDailyQuotaLeft string

	//
	TroubleShootingEndTime string

	//
	TroubleShootingExpirationLeft string

	//
	TroubleShootingMode string

	//
	TroubleShootingModeSource string

	//
	TroubleShootingQuotaResetTime string

	//
	TroubleShootingStartTime string
}

func NewMSFT_MpComputerStatusEx1(instance *cim.WmiInstance) (newInstance *MSFT_MpComputerStatus, err error) {
	tmp, err := NewBaseStatusEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MpComputerStatus{
		BaseStatus: tmp,
	}
	return
}

func NewMSFT_MpComputerStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MpComputerStatus, err error) {
	tmp, err := NewBaseStatusEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MpComputerStatus{
		BaseStatus: tmp,
	}
	return
}

// SetAMEngineVersion sets the value of AMEngineVersion for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAMEngineVersion(value string) (err error) {
	return instance.SetProperty("AMEngineVersion", (value))
}

// GetAMEngineVersion gets the value of AMEngineVersion for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAMEngineVersion() (value string, err error) {
	retValue, err := instance.GetProperty("AMEngineVersion")
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

// SetAMProductVersion sets the value of AMProductVersion for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAMProductVersion(value string) (err error) {
	return instance.SetProperty("AMProductVersion", (value))
}

// GetAMProductVersion gets the value of AMProductVersion for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAMProductVersion() (value string, err error) {
	retValue, err := instance.GetProperty("AMProductVersion")
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

// SetAMRunningMode sets the value of AMRunningMode for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAMRunningMode(value string) (err error) {
	return instance.SetProperty("AMRunningMode", (value))
}

// GetAMRunningMode gets the value of AMRunningMode for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAMRunningMode() (value string, err error) {
	retValue, err := instance.GetProperty("AMRunningMode")
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

// SetAMServiceEnabled sets the value of AMServiceEnabled for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAMServiceEnabled(value bool) (err error) {
	return instance.SetProperty("AMServiceEnabled", (value))
}

// GetAMServiceEnabled gets the value of AMServiceEnabled for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAMServiceEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("AMServiceEnabled")
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

// SetAMServiceVersion sets the value of AMServiceVersion for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAMServiceVersion(value string) (err error) {
	return instance.SetProperty("AMServiceVersion", (value))
}

// GetAMServiceVersion gets the value of AMServiceVersion for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAMServiceVersion() (value string, err error) {
	retValue, err := instance.GetProperty("AMServiceVersion")
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

// SetAntispywareEnabled sets the value of AntispywareEnabled for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAntispywareEnabled(value bool) (err error) {
	return instance.SetProperty("AntispywareEnabled", (value))
}

// GetAntispywareEnabled gets the value of AntispywareEnabled for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAntispywareEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("AntispywareEnabled")
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

// SetAntispywareSignatureAge sets the value of AntispywareSignatureAge for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAntispywareSignatureAge(value uint32) (err error) {
	return instance.SetProperty("AntispywareSignatureAge", (value))
}

// GetAntispywareSignatureAge gets the value of AntispywareSignatureAge for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAntispywareSignatureAge() (value uint32, err error) {
	retValue, err := instance.GetProperty("AntispywareSignatureAge")
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

// SetAntispywareSignatureLastUpdated sets the value of AntispywareSignatureLastUpdated for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAntispywareSignatureLastUpdated(value string) (err error) {
	return instance.SetProperty("AntispywareSignatureLastUpdated", (value))
}

// GetAntispywareSignatureLastUpdated gets the value of AntispywareSignatureLastUpdated for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAntispywareSignatureLastUpdated() (value string, err error) {
	retValue, err := instance.GetProperty("AntispywareSignatureLastUpdated")
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

// SetAntispywareSignatureVersion sets the value of AntispywareSignatureVersion for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAntispywareSignatureVersion(value string) (err error) {
	return instance.SetProperty("AntispywareSignatureVersion", (value))
}

// GetAntispywareSignatureVersion gets the value of AntispywareSignatureVersion for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAntispywareSignatureVersion() (value string, err error) {
	retValue, err := instance.GetProperty("AntispywareSignatureVersion")
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

// SetAntivirusEnabled sets the value of AntivirusEnabled for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAntivirusEnabled(value bool) (err error) {
	return instance.SetProperty("AntivirusEnabled", (value))
}

// GetAntivirusEnabled gets the value of AntivirusEnabled for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAntivirusEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("AntivirusEnabled")
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

// SetAntivirusSignatureAge sets the value of AntivirusSignatureAge for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAntivirusSignatureAge(value uint32) (err error) {
	return instance.SetProperty("AntivirusSignatureAge", (value))
}

// GetAntivirusSignatureAge gets the value of AntivirusSignatureAge for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAntivirusSignatureAge() (value uint32, err error) {
	retValue, err := instance.GetProperty("AntivirusSignatureAge")
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

// SetAntivirusSignatureLastUpdated sets the value of AntivirusSignatureLastUpdated for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAntivirusSignatureLastUpdated(value string) (err error) {
	return instance.SetProperty("AntivirusSignatureLastUpdated", (value))
}

// GetAntivirusSignatureLastUpdated gets the value of AntivirusSignatureLastUpdated for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAntivirusSignatureLastUpdated() (value string, err error) {
	retValue, err := instance.GetProperty("AntivirusSignatureLastUpdated")
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

// SetAntivirusSignatureVersion sets the value of AntivirusSignatureVersion for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAntivirusSignatureVersion(value string) (err error) {
	return instance.SetProperty("AntivirusSignatureVersion", (value))
}

// GetAntivirusSignatureVersion gets the value of AntivirusSignatureVersion for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAntivirusSignatureVersion() (value string, err error) {
	retValue, err := instance.GetProperty("AntivirusSignatureVersion")
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

// SetBehaviorMonitorEnabled sets the value of BehaviorMonitorEnabled for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyBehaviorMonitorEnabled(value bool) (err error) {
	return instance.SetProperty("BehaviorMonitorEnabled", (value))
}

// GetBehaviorMonitorEnabled gets the value of BehaviorMonitorEnabled for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyBehaviorMonitorEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("BehaviorMonitorEnabled")
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

// SetComputerID sets the value of ComputerID for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyComputerID(value string) (err error) {
	return instance.SetProperty("ComputerID", (value))
}

// GetComputerID gets the value of ComputerID for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyComputerID() (value string, err error) {
	retValue, err := instance.GetProperty("ComputerID")
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

// SetComputerState sets the value of ComputerState for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyComputerState(value uint32) (err error) {
	return instance.SetProperty("ComputerState", (value))
}

// GetComputerState gets the value of ComputerState for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyComputerState() (value uint32, err error) {
	retValue, err := instance.GetProperty("ComputerState")
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

// SetDefenderSignaturesOutOfDate sets the value of DefenderSignaturesOutOfDate for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyDefenderSignaturesOutOfDate(value bool) (err error) {
	return instance.SetProperty("DefenderSignaturesOutOfDate", (value))
}

// GetDefenderSignaturesOutOfDate gets the value of DefenderSignaturesOutOfDate for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyDefenderSignaturesOutOfDate() (value bool, err error) {
	retValue, err := instance.GetProperty("DefenderSignaturesOutOfDate")
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

// SetDeviceControlDefaultEnforcement sets the value of DeviceControlDefaultEnforcement for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyDeviceControlDefaultEnforcement(value string) (err error) {
	return instance.SetProperty("DeviceControlDefaultEnforcement", (value))
}

// GetDeviceControlDefaultEnforcement gets the value of DeviceControlDefaultEnforcement for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyDeviceControlDefaultEnforcement() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceControlDefaultEnforcement")
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

// SetDeviceControlPoliciesLastUpdated sets the value of DeviceControlPoliciesLastUpdated for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyDeviceControlPoliciesLastUpdated(value string) (err error) {
	return instance.SetProperty("DeviceControlPoliciesLastUpdated", (value))
}

// GetDeviceControlPoliciesLastUpdated gets the value of DeviceControlPoliciesLastUpdated for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyDeviceControlPoliciesLastUpdated() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceControlPoliciesLastUpdated")
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

// SetDeviceControlState sets the value of DeviceControlState for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyDeviceControlState(value string) (err error) {
	return instance.SetProperty("DeviceControlState", (value))
}

// GetDeviceControlState gets the value of DeviceControlState for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyDeviceControlState() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceControlState")
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

// SetFullScanAge sets the value of FullScanAge for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyFullScanAge(value uint32) (err error) {
	return instance.SetProperty("FullScanAge", (value))
}

// GetFullScanAge gets the value of FullScanAge for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyFullScanAge() (value uint32, err error) {
	retValue, err := instance.GetProperty("FullScanAge")
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

// SetFullScanEndTime sets the value of FullScanEndTime for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyFullScanEndTime(value string) (err error) {
	return instance.SetProperty("FullScanEndTime", (value))
}

// GetFullScanEndTime gets the value of FullScanEndTime for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyFullScanEndTime() (value string, err error) {
	retValue, err := instance.GetProperty("FullScanEndTime")
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

// SetFullScanOverdue sets the value of FullScanOverdue for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyFullScanOverdue(value bool) (err error) {
	return instance.SetProperty("FullScanOverdue", (value))
}

// GetFullScanOverdue gets the value of FullScanOverdue for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyFullScanOverdue() (value bool, err error) {
	retValue, err := instance.GetProperty("FullScanOverdue")
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

// SetFullScanRequired sets the value of FullScanRequired for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyFullScanRequired(value bool) (err error) {
	return instance.SetProperty("FullScanRequired", (value))
}

// GetFullScanRequired gets the value of FullScanRequired for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyFullScanRequired() (value bool, err error) {
	retValue, err := instance.GetProperty("FullScanRequired")
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

// SetFullScanSignatureVersion sets the value of FullScanSignatureVersion for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyFullScanSignatureVersion(value string) (err error) {
	return instance.SetProperty("FullScanSignatureVersion", (value))
}

// GetFullScanSignatureVersion gets the value of FullScanSignatureVersion for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyFullScanSignatureVersion() (value string, err error) {
	retValue, err := instance.GetProperty("FullScanSignatureVersion")
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

// SetFullScanStartTime sets the value of FullScanStartTime for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyFullScanStartTime(value string) (err error) {
	return instance.SetProperty("FullScanStartTime", (value))
}

// GetFullScanStartTime gets the value of FullScanStartTime for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyFullScanStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("FullScanStartTime")
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

// SetInitializationProgress sets the value of InitializationProgress for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyInitializationProgress(value string) (err error) {
	return instance.SetProperty("InitializationProgress", (value))
}

// GetInitializationProgress gets the value of InitializationProgress for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyInitializationProgress() (value string, err error) {
	retValue, err := instance.GetProperty("InitializationProgress")
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

// SetIoavProtectionEnabled sets the value of IoavProtectionEnabled for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyIoavProtectionEnabled(value bool) (err error) {
	return instance.SetProperty("IoavProtectionEnabled", (value))
}

// GetIoavProtectionEnabled gets the value of IoavProtectionEnabled for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyIoavProtectionEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IoavProtectionEnabled")
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

// SetIsTamperProtected sets the value of IsTamperProtected for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyIsTamperProtected(value bool) (err error) {
	return instance.SetProperty("IsTamperProtected", (value))
}

// GetIsTamperProtected gets the value of IsTamperProtected for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyIsTamperProtected() (value bool, err error) {
	retValue, err := instance.GetProperty("IsTamperProtected")
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

// SetIsVirtualMachine sets the value of IsVirtualMachine for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyIsVirtualMachine(value bool) (err error) {
	return instance.SetProperty("IsVirtualMachine", (value))
}

// GetIsVirtualMachine gets the value of IsVirtualMachine for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyIsVirtualMachine() (value bool, err error) {
	retValue, err := instance.GetProperty("IsVirtualMachine")
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

// SetLastFullScanSource sets the value of LastFullScanSource for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyLastFullScanSource(value uint8) (err error) {
	return instance.SetProperty("LastFullScanSource", (value))
}

// GetLastFullScanSource gets the value of LastFullScanSource for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyLastFullScanSource() (value uint8, err error) {
	retValue, err := instance.GetProperty("LastFullScanSource")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetLastQuickScanSource sets the value of LastQuickScanSource for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyLastQuickScanSource(value uint8) (err error) {
	return instance.SetProperty("LastQuickScanSource", (value))
}

// GetLastQuickScanSource gets the value of LastQuickScanSource for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyLastQuickScanSource() (value uint8, err error) {
	retValue, err := instance.GetProperty("LastQuickScanSource")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetNISEnabled sets the value of NISEnabled for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyNISEnabled(value bool) (err error) {
	return instance.SetProperty("NISEnabled", (value))
}

// GetNISEnabled gets the value of NISEnabled for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyNISEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("NISEnabled")
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

// SetNISEngineVersion sets the value of NISEngineVersion for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyNISEngineVersion(value string) (err error) {
	return instance.SetProperty("NISEngineVersion", (value))
}

// GetNISEngineVersion gets the value of NISEngineVersion for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyNISEngineVersion() (value string, err error) {
	retValue, err := instance.GetProperty("NISEngineVersion")
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

// SetNISSignatureAge sets the value of NISSignatureAge for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyNISSignatureAge(value uint32) (err error) {
	return instance.SetProperty("NISSignatureAge", (value))
}

// GetNISSignatureAge gets the value of NISSignatureAge for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyNISSignatureAge() (value uint32, err error) {
	retValue, err := instance.GetProperty("NISSignatureAge")
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

// SetNISSignatureLastUpdated sets the value of NISSignatureLastUpdated for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyNISSignatureLastUpdated(value string) (err error) {
	return instance.SetProperty("NISSignatureLastUpdated", (value))
}

// GetNISSignatureLastUpdated gets the value of NISSignatureLastUpdated for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyNISSignatureLastUpdated() (value string, err error) {
	retValue, err := instance.GetProperty("NISSignatureLastUpdated")
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

// SetNISSignatureVersion sets the value of NISSignatureVersion for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyNISSignatureVersion(value string) (err error) {
	return instance.SetProperty("NISSignatureVersion", (value))
}

// GetNISSignatureVersion gets the value of NISSignatureVersion for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyNISSignatureVersion() (value string, err error) {
	retValue, err := instance.GetProperty("NISSignatureVersion")
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

// SetOnAccessProtectionEnabled sets the value of OnAccessProtectionEnabled for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyOnAccessProtectionEnabled(value bool) (err error) {
	return instance.SetProperty("OnAccessProtectionEnabled", (value))
}

// GetOnAccessProtectionEnabled gets the value of OnAccessProtectionEnabled for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyOnAccessProtectionEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("OnAccessProtectionEnabled")
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

// SetProductStatus sets the value of ProductStatus for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyProductStatus(value uint32) (err error) {
	return instance.SetProperty("ProductStatus", (value))
}

// GetProductStatus gets the value of ProductStatus for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyProductStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProductStatus")
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

// SetQuickScanAge sets the value of QuickScanAge for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyQuickScanAge(value uint32) (err error) {
	return instance.SetProperty("QuickScanAge", (value))
}

// GetQuickScanAge gets the value of QuickScanAge for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyQuickScanAge() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuickScanAge")
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

// SetQuickScanEndTime sets the value of QuickScanEndTime for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyQuickScanEndTime(value string) (err error) {
	return instance.SetProperty("QuickScanEndTime", (value))
}

// GetQuickScanEndTime gets the value of QuickScanEndTime for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyQuickScanEndTime() (value string, err error) {
	retValue, err := instance.GetProperty("QuickScanEndTime")
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

// SetQuickScanOverdue sets the value of QuickScanOverdue for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyQuickScanOverdue(value bool) (err error) {
	return instance.SetProperty("QuickScanOverdue", (value))
}

// GetQuickScanOverdue gets the value of QuickScanOverdue for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyQuickScanOverdue() (value bool, err error) {
	retValue, err := instance.GetProperty("QuickScanOverdue")
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

// SetQuickScanSignatureVersion sets the value of QuickScanSignatureVersion for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyQuickScanSignatureVersion(value string) (err error) {
	return instance.SetProperty("QuickScanSignatureVersion", (value))
}

// GetQuickScanSignatureVersion gets the value of QuickScanSignatureVersion for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyQuickScanSignatureVersion() (value string, err error) {
	retValue, err := instance.GetProperty("QuickScanSignatureVersion")
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

// SetQuickScanStartTime sets the value of QuickScanStartTime for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyQuickScanStartTime(value string) (err error) {
	return instance.SetProperty("QuickScanStartTime", (value))
}

// GetQuickScanStartTime gets the value of QuickScanStartTime for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyQuickScanStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("QuickScanStartTime")
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

// SetRealTimeProtectionEnabled sets the value of RealTimeProtectionEnabled for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyRealTimeProtectionEnabled(value bool) (err error) {
	return instance.SetProperty("RealTimeProtectionEnabled", (value))
}

// GetRealTimeProtectionEnabled gets the value of RealTimeProtectionEnabled for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyRealTimeProtectionEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("RealTimeProtectionEnabled")
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

// SetRealTimeScanDirection sets the value of RealTimeScanDirection for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyRealTimeScanDirection(value uint8) (err error) {
	return instance.SetProperty("RealTimeScanDirection", (value))
}

// GetRealTimeScanDirection gets the value of RealTimeScanDirection for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyRealTimeScanDirection() (value uint8, err error) {
	retValue, err := instance.GetProperty("RealTimeScanDirection")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetRebootRequired sets the value of RebootRequired for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyRebootRequired(value bool) (err error) {
	return instance.SetProperty("RebootRequired", (value))
}

// GetRebootRequired gets the value of RebootRequired for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyRebootRequired() (value bool, err error) {
	retValue, err := instance.GetProperty("RebootRequired")
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

// SetSmartAppControlExpiration sets the value of SmartAppControlExpiration for the instance
func (instance *MSFT_MpComputerStatus) SetPropertySmartAppControlExpiration(value string) (err error) {
	return instance.SetProperty("SmartAppControlExpiration", (value))
}

// GetSmartAppControlExpiration gets the value of SmartAppControlExpiration for the instance
func (instance *MSFT_MpComputerStatus) GetPropertySmartAppControlExpiration() (value string, err error) {
	retValue, err := instance.GetProperty("SmartAppControlExpiration")
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

// SetSmartAppControlState sets the value of SmartAppControlState for the instance
func (instance *MSFT_MpComputerStatus) SetPropertySmartAppControlState(value string) (err error) {
	return instance.SetProperty("SmartAppControlState", (value))
}

// GetSmartAppControlState gets the value of SmartAppControlState for the instance
func (instance *MSFT_MpComputerStatus) GetPropertySmartAppControlState() (value string, err error) {
	retValue, err := instance.GetProperty("SmartAppControlState")
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

// SetTamperProtectionSource sets the value of TamperProtectionSource for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyTamperProtectionSource(value string) (err error) {
	return instance.SetProperty("TamperProtectionSource", (value))
}

// GetTamperProtectionSource gets the value of TamperProtectionSource for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyTamperProtectionSource() (value string, err error) {
	retValue, err := instance.GetProperty("TamperProtectionSource")
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

// SetTDTCapable sets the value of TDTCapable for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyTDTCapable(value string) (err error) {
	return instance.SetProperty("TDTCapable", (value))
}

// GetTDTCapable gets the value of TDTCapable for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyTDTCapable() (value string, err error) {
	retValue, err := instance.GetProperty("TDTCapable")
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

// SetTDTMode sets the value of TDTMode for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyTDTMode(value string) (err error) {
	return instance.SetProperty("TDTMode", (value))
}

// GetTDTMode gets the value of TDTMode for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyTDTMode() (value string, err error) {
	retValue, err := instance.GetProperty("TDTMode")
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

// SetTDTSiloType sets the value of TDTSiloType for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyTDTSiloType(value string) (err error) {
	return instance.SetProperty("TDTSiloType", (value))
}

// GetTDTSiloType gets the value of TDTSiloType for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyTDTSiloType() (value string, err error) {
	retValue, err := instance.GetProperty("TDTSiloType")
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

// SetTDTStatus sets the value of TDTStatus for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyTDTStatus(value string) (err error) {
	return instance.SetProperty("TDTStatus", (value))
}

// GetTDTStatus gets the value of TDTStatus for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyTDTStatus() (value string, err error) {
	retValue, err := instance.GetProperty("TDTStatus")
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

// SetTDTTelemetry sets the value of TDTTelemetry for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyTDTTelemetry(value string) (err error) {
	return instance.SetProperty("TDTTelemetry", (value))
}

// GetTDTTelemetry gets the value of TDTTelemetry for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyTDTTelemetry() (value string, err error) {
	retValue, err := instance.GetProperty("TDTTelemetry")
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

// SetTroubleShootingDailyMaxQuota sets the value of TroubleShootingDailyMaxQuota for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyTroubleShootingDailyMaxQuota(value string) (err error) {
	return instance.SetProperty("TroubleShootingDailyMaxQuota", (value))
}

// GetTroubleShootingDailyMaxQuota gets the value of TroubleShootingDailyMaxQuota for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyTroubleShootingDailyMaxQuota() (value string, err error) {
	retValue, err := instance.GetProperty("TroubleShootingDailyMaxQuota")
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

// SetTroubleShootingDailyQuotaLeft sets the value of TroubleShootingDailyQuotaLeft for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyTroubleShootingDailyQuotaLeft(value string) (err error) {
	return instance.SetProperty("TroubleShootingDailyQuotaLeft", (value))
}

// GetTroubleShootingDailyQuotaLeft gets the value of TroubleShootingDailyQuotaLeft for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyTroubleShootingDailyQuotaLeft() (value string, err error) {
	retValue, err := instance.GetProperty("TroubleShootingDailyQuotaLeft")
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

// SetTroubleShootingEndTime sets the value of TroubleShootingEndTime for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyTroubleShootingEndTime(value string) (err error) {
	return instance.SetProperty("TroubleShootingEndTime", (value))
}

// GetTroubleShootingEndTime gets the value of TroubleShootingEndTime for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyTroubleShootingEndTime() (value string, err error) {
	retValue, err := instance.GetProperty("TroubleShootingEndTime")
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

// SetTroubleShootingExpirationLeft sets the value of TroubleShootingExpirationLeft for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyTroubleShootingExpirationLeft(value string) (err error) {
	return instance.SetProperty("TroubleShootingExpirationLeft", (value))
}

// GetTroubleShootingExpirationLeft gets the value of TroubleShootingExpirationLeft for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyTroubleShootingExpirationLeft() (value string, err error) {
	retValue, err := instance.GetProperty("TroubleShootingExpirationLeft")
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

// SetTroubleShootingMode sets the value of TroubleShootingMode for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyTroubleShootingMode(value string) (err error) {
	return instance.SetProperty("TroubleShootingMode", (value))
}

// GetTroubleShootingMode gets the value of TroubleShootingMode for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyTroubleShootingMode() (value string, err error) {
	retValue, err := instance.GetProperty("TroubleShootingMode")
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

// SetTroubleShootingModeSource sets the value of TroubleShootingModeSource for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyTroubleShootingModeSource(value string) (err error) {
	return instance.SetProperty("TroubleShootingModeSource", (value))
}

// GetTroubleShootingModeSource gets the value of TroubleShootingModeSource for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyTroubleShootingModeSource() (value string, err error) {
	retValue, err := instance.GetProperty("TroubleShootingModeSource")
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

// SetTroubleShootingQuotaResetTime sets the value of TroubleShootingQuotaResetTime for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyTroubleShootingQuotaResetTime(value string) (err error) {
	return instance.SetProperty("TroubleShootingQuotaResetTime", (value))
}

// GetTroubleShootingQuotaResetTime gets the value of TroubleShootingQuotaResetTime for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyTroubleShootingQuotaResetTime() (value string, err error) {
	retValue, err := instance.GetProperty("TroubleShootingQuotaResetTime")
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

// SetTroubleShootingStartTime sets the value of TroubleShootingStartTime for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyTroubleShootingStartTime(value string) (err error) {
	return instance.SetProperty("TroubleShootingStartTime", (value))
}

// GetTroubleShootingStartTime gets the value of TroubleShootingStartTime for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyTroubleShootingStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("TroubleShootingStartTime")
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
