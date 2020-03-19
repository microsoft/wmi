// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.protectionManagement
//////////////////////////////////////////////
package protectionmanagement

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MpComputerStatus struct
type MSFT_MpComputerStatus struct {
	*BaseStatus

	//
	AMEngineVersion string

	//
	AMProductVersion string

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
	FullScanAge uint32

	//
	FullScanEndTime string

	//
	FullScanStartTime string

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
	QuickScanAge uint32

	//
	QuickScanEndTime string

	//
	QuickScanStartTime string

	//
	RealTimeProtectionEnabled bool

	//
	RealTimeScanDirection uint8
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
	return instance.SetProperty("AMEngineVersion", value)
}

// GetAMEngineVersion gets the value of AMEngineVersion for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAMEngineVersion() (value string, err error) {
	retValue, err := instance.GetProperty("AMEngineVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAMProductVersion sets the value of AMProductVersion for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAMProductVersion(value string) (err error) {
	return instance.SetProperty("AMProductVersion", value)
}

// GetAMProductVersion gets the value of AMProductVersion for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAMProductVersion() (value string, err error) {
	retValue, err := instance.GetProperty("AMProductVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAMServiceEnabled sets the value of AMServiceEnabled for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAMServiceEnabled(value bool) (err error) {
	return instance.SetProperty("AMServiceEnabled", value)
}

// GetAMServiceEnabled gets the value of AMServiceEnabled for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAMServiceEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("AMServiceEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAMServiceVersion sets the value of AMServiceVersion for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAMServiceVersion(value string) (err error) {
	return instance.SetProperty("AMServiceVersion", value)
}

// GetAMServiceVersion gets the value of AMServiceVersion for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAMServiceVersion() (value string, err error) {
	retValue, err := instance.GetProperty("AMServiceVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAntispywareEnabled sets the value of AntispywareEnabled for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAntispywareEnabled(value bool) (err error) {
	return instance.SetProperty("AntispywareEnabled", value)
}

// GetAntispywareEnabled gets the value of AntispywareEnabled for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAntispywareEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("AntispywareEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAntispywareSignatureAge sets the value of AntispywareSignatureAge for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAntispywareSignatureAge(value uint32) (err error) {
	return instance.SetProperty("AntispywareSignatureAge", value)
}

// GetAntispywareSignatureAge gets the value of AntispywareSignatureAge for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAntispywareSignatureAge() (value uint32, err error) {
	retValue, err := instance.GetProperty("AntispywareSignatureAge")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAntispywareSignatureLastUpdated sets the value of AntispywareSignatureLastUpdated for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAntispywareSignatureLastUpdated(value string) (err error) {
	return instance.SetProperty("AntispywareSignatureLastUpdated", value)
}

// GetAntispywareSignatureLastUpdated gets the value of AntispywareSignatureLastUpdated for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAntispywareSignatureLastUpdated() (value string, err error) {
	retValue, err := instance.GetProperty("AntispywareSignatureLastUpdated")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAntispywareSignatureVersion sets the value of AntispywareSignatureVersion for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAntispywareSignatureVersion(value string) (err error) {
	return instance.SetProperty("AntispywareSignatureVersion", value)
}

// GetAntispywareSignatureVersion gets the value of AntispywareSignatureVersion for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAntispywareSignatureVersion() (value string, err error) {
	retValue, err := instance.GetProperty("AntispywareSignatureVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAntivirusEnabled sets the value of AntivirusEnabled for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAntivirusEnabled(value bool) (err error) {
	return instance.SetProperty("AntivirusEnabled", value)
}

// GetAntivirusEnabled gets the value of AntivirusEnabled for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAntivirusEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("AntivirusEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAntivirusSignatureAge sets the value of AntivirusSignatureAge for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAntivirusSignatureAge(value uint32) (err error) {
	return instance.SetProperty("AntivirusSignatureAge", value)
}

// GetAntivirusSignatureAge gets the value of AntivirusSignatureAge for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAntivirusSignatureAge() (value uint32, err error) {
	retValue, err := instance.GetProperty("AntivirusSignatureAge")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAntivirusSignatureLastUpdated sets the value of AntivirusSignatureLastUpdated for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAntivirusSignatureLastUpdated(value string) (err error) {
	return instance.SetProperty("AntivirusSignatureLastUpdated", value)
}

// GetAntivirusSignatureLastUpdated gets the value of AntivirusSignatureLastUpdated for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAntivirusSignatureLastUpdated() (value string, err error) {
	retValue, err := instance.GetProperty("AntivirusSignatureLastUpdated")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAntivirusSignatureVersion sets the value of AntivirusSignatureVersion for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyAntivirusSignatureVersion(value string) (err error) {
	return instance.SetProperty("AntivirusSignatureVersion", value)
}

// GetAntivirusSignatureVersion gets the value of AntivirusSignatureVersion for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyAntivirusSignatureVersion() (value string, err error) {
	retValue, err := instance.GetProperty("AntivirusSignatureVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBehaviorMonitorEnabled sets the value of BehaviorMonitorEnabled for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyBehaviorMonitorEnabled(value bool) (err error) {
	return instance.SetProperty("BehaviorMonitorEnabled", value)
}

// GetBehaviorMonitorEnabled gets the value of BehaviorMonitorEnabled for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyBehaviorMonitorEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("BehaviorMonitorEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetComputerID sets the value of ComputerID for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyComputerID(value string) (err error) {
	return instance.SetProperty("ComputerID", value)
}

// GetComputerID gets the value of ComputerID for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyComputerID() (value string, err error) {
	retValue, err := instance.GetProperty("ComputerID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetComputerState sets the value of ComputerState for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyComputerState(value uint32) (err error) {
	return instance.SetProperty("ComputerState", value)
}

// GetComputerState gets the value of ComputerState for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyComputerState() (value uint32, err error) {
	retValue, err := instance.GetProperty("ComputerState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFullScanAge sets the value of FullScanAge for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyFullScanAge(value uint32) (err error) {
	return instance.SetProperty("FullScanAge", value)
}

// GetFullScanAge gets the value of FullScanAge for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyFullScanAge() (value uint32, err error) {
	retValue, err := instance.GetProperty("FullScanAge")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFullScanEndTime sets the value of FullScanEndTime for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyFullScanEndTime(value string) (err error) {
	return instance.SetProperty("FullScanEndTime", value)
}

// GetFullScanEndTime gets the value of FullScanEndTime for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyFullScanEndTime() (value string, err error) {
	retValue, err := instance.GetProperty("FullScanEndTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFullScanStartTime sets the value of FullScanStartTime for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyFullScanStartTime(value string) (err error) {
	return instance.SetProperty("FullScanStartTime", value)
}

// GetFullScanStartTime gets the value of FullScanStartTime for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyFullScanStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("FullScanStartTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIoavProtectionEnabled sets the value of IoavProtectionEnabled for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyIoavProtectionEnabled(value bool) (err error) {
	return instance.SetProperty("IoavProtectionEnabled", value)
}

// GetIoavProtectionEnabled gets the value of IoavProtectionEnabled for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyIoavProtectionEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IoavProtectionEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsTamperProtected sets the value of IsTamperProtected for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyIsTamperProtected(value bool) (err error) {
	return instance.SetProperty("IsTamperProtected", value)
}

// GetIsTamperProtected gets the value of IsTamperProtected for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyIsTamperProtected() (value bool, err error) {
	retValue, err := instance.GetProperty("IsTamperProtected")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsVirtualMachine sets the value of IsVirtualMachine for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyIsVirtualMachine(value bool) (err error) {
	return instance.SetProperty("IsVirtualMachine", value)
}

// GetIsVirtualMachine gets the value of IsVirtualMachine for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyIsVirtualMachine() (value bool, err error) {
	retValue, err := instance.GetProperty("IsVirtualMachine")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastFullScanSource sets the value of LastFullScanSource for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyLastFullScanSource(value uint8) (err error) {
	return instance.SetProperty("LastFullScanSource", value)
}

// GetLastFullScanSource gets the value of LastFullScanSource for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyLastFullScanSource() (value uint8, err error) {
	retValue, err := instance.GetProperty("LastFullScanSource")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastQuickScanSource sets the value of LastQuickScanSource for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyLastQuickScanSource(value uint8) (err error) {
	return instance.SetProperty("LastQuickScanSource", value)
}

// GetLastQuickScanSource gets the value of LastQuickScanSource for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyLastQuickScanSource() (value uint8, err error) {
	retValue, err := instance.GetProperty("LastQuickScanSource")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNISEnabled sets the value of NISEnabled for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyNISEnabled(value bool) (err error) {
	return instance.SetProperty("NISEnabled", value)
}

// GetNISEnabled gets the value of NISEnabled for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyNISEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("NISEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNISEngineVersion sets the value of NISEngineVersion for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyNISEngineVersion(value string) (err error) {
	return instance.SetProperty("NISEngineVersion", value)
}

// GetNISEngineVersion gets the value of NISEngineVersion for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyNISEngineVersion() (value string, err error) {
	retValue, err := instance.GetProperty("NISEngineVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNISSignatureAge sets the value of NISSignatureAge for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyNISSignatureAge(value uint32) (err error) {
	return instance.SetProperty("NISSignatureAge", value)
}

// GetNISSignatureAge gets the value of NISSignatureAge for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyNISSignatureAge() (value uint32, err error) {
	retValue, err := instance.GetProperty("NISSignatureAge")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNISSignatureLastUpdated sets the value of NISSignatureLastUpdated for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyNISSignatureLastUpdated(value string) (err error) {
	return instance.SetProperty("NISSignatureLastUpdated", value)
}

// GetNISSignatureLastUpdated gets the value of NISSignatureLastUpdated for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyNISSignatureLastUpdated() (value string, err error) {
	retValue, err := instance.GetProperty("NISSignatureLastUpdated")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNISSignatureVersion sets the value of NISSignatureVersion for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyNISSignatureVersion(value string) (err error) {
	return instance.SetProperty("NISSignatureVersion", value)
}

// GetNISSignatureVersion gets the value of NISSignatureVersion for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyNISSignatureVersion() (value string, err error) {
	retValue, err := instance.GetProperty("NISSignatureVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOnAccessProtectionEnabled sets the value of OnAccessProtectionEnabled for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyOnAccessProtectionEnabled(value bool) (err error) {
	return instance.SetProperty("OnAccessProtectionEnabled", value)
}

// GetOnAccessProtectionEnabled gets the value of OnAccessProtectionEnabled for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyOnAccessProtectionEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("OnAccessProtectionEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQuickScanAge sets the value of QuickScanAge for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyQuickScanAge(value uint32) (err error) {
	return instance.SetProperty("QuickScanAge", value)
}

// GetQuickScanAge gets the value of QuickScanAge for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyQuickScanAge() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuickScanAge")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQuickScanEndTime sets the value of QuickScanEndTime for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyQuickScanEndTime(value string) (err error) {
	return instance.SetProperty("QuickScanEndTime", value)
}

// GetQuickScanEndTime gets the value of QuickScanEndTime for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyQuickScanEndTime() (value string, err error) {
	retValue, err := instance.GetProperty("QuickScanEndTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQuickScanStartTime sets the value of QuickScanStartTime for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyQuickScanStartTime(value string) (err error) {
	return instance.SetProperty("QuickScanStartTime", value)
}

// GetQuickScanStartTime gets the value of QuickScanStartTime for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyQuickScanStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("QuickScanStartTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRealTimeProtectionEnabled sets the value of RealTimeProtectionEnabled for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyRealTimeProtectionEnabled(value bool) (err error) {
	return instance.SetProperty("RealTimeProtectionEnabled", value)
}

// GetRealTimeProtectionEnabled gets the value of RealTimeProtectionEnabled for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyRealTimeProtectionEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("RealTimeProtectionEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRealTimeScanDirection sets the value of RealTimeScanDirection for the instance
func (instance *MSFT_MpComputerStatus) SetPropertyRealTimeScanDirection(value uint8) (err error) {
	return instance.SetProperty("RealTimeScanDirection", value)
}

// GetRealTimeScanDirection gets the value of RealTimeScanDirection for the instance
func (instance *MSFT_MpComputerStatus) GetPropertyRealTimeScanDirection() (value uint8, err error) {
	retValue, err := instance.GetProperty("RealTimeScanDirection")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}
