// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.Defender
//////////////////////////////////////////////
package defender

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_MpThreatDetection struct
type MSFT_MpThreatDetection struct {
	*BaseStatus

	//
	ActionSuccess bool

	//
	AdditionalActionsBitMask uint32

	//
	AMProductVersion string

	//
	CleaningActionID uint8

	//
	CurrentThreatExecutionStatusID uint8

	//
	DetectionID string

	//
	DetectionSourceTypeID uint8

	//
	DomainUser string

	//
	InitialDetectionTime string

	//
	LastThreatStatusChangeTime string

	//
	ProcessName string

	//
	RemediationTime string

	//
	Resources []string

	//
	ThreatID int64

	//
	ThreatStatusErrorCode int32

	//
	ThreatStatusID uint8
}

func NewMSFT_MpThreatDetectionEx1(instance *cim.WmiInstance) (newInstance *MSFT_MpThreatDetection, err error) {
	tmp, err := NewBaseStatusEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MpThreatDetection{
		BaseStatus: tmp,
	}
	return
}

func NewMSFT_MpThreatDetectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MpThreatDetection, err error) {
	tmp, err := NewBaseStatusEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MpThreatDetection{
		BaseStatus: tmp,
	}
	return
}

// SetActionSuccess sets the value of ActionSuccess for the instance
func (instance *MSFT_MpThreatDetection) SetPropertyActionSuccess(value bool) (err error) {
	return instance.SetProperty("ActionSuccess", (value))
}

// GetActionSuccess gets the value of ActionSuccess for the instance
func (instance *MSFT_MpThreatDetection) GetPropertyActionSuccess() (value bool, err error) {
	retValue, err := instance.GetProperty("ActionSuccess")
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

// SetAdditionalActionsBitMask sets the value of AdditionalActionsBitMask for the instance
func (instance *MSFT_MpThreatDetection) SetPropertyAdditionalActionsBitMask(value uint32) (err error) {
	return instance.SetProperty("AdditionalActionsBitMask", (value))
}

// GetAdditionalActionsBitMask gets the value of AdditionalActionsBitMask for the instance
func (instance *MSFT_MpThreatDetection) GetPropertyAdditionalActionsBitMask() (value uint32, err error) {
	retValue, err := instance.GetProperty("AdditionalActionsBitMask")
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

// SetAMProductVersion sets the value of AMProductVersion for the instance
func (instance *MSFT_MpThreatDetection) SetPropertyAMProductVersion(value string) (err error) {
	return instance.SetProperty("AMProductVersion", (value))
}

// GetAMProductVersion gets the value of AMProductVersion for the instance
func (instance *MSFT_MpThreatDetection) GetPropertyAMProductVersion() (value string, err error) {
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

// SetCleaningActionID sets the value of CleaningActionID for the instance
func (instance *MSFT_MpThreatDetection) SetPropertyCleaningActionID(value uint8) (err error) {
	return instance.SetProperty("CleaningActionID", (value))
}

// GetCleaningActionID gets the value of CleaningActionID for the instance
func (instance *MSFT_MpThreatDetection) GetPropertyCleaningActionID() (value uint8, err error) {
	retValue, err := instance.GetProperty("CleaningActionID")
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

// SetCurrentThreatExecutionStatusID sets the value of CurrentThreatExecutionStatusID for the instance
func (instance *MSFT_MpThreatDetection) SetPropertyCurrentThreatExecutionStatusID(value uint8) (err error) {
	return instance.SetProperty("CurrentThreatExecutionStatusID", (value))
}

// GetCurrentThreatExecutionStatusID gets the value of CurrentThreatExecutionStatusID for the instance
func (instance *MSFT_MpThreatDetection) GetPropertyCurrentThreatExecutionStatusID() (value uint8, err error) {
	retValue, err := instance.GetProperty("CurrentThreatExecutionStatusID")
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

// SetDetectionID sets the value of DetectionID for the instance
func (instance *MSFT_MpThreatDetection) SetPropertyDetectionID(value string) (err error) {
	return instance.SetProperty("DetectionID", (value))
}

// GetDetectionID gets the value of DetectionID for the instance
func (instance *MSFT_MpThreatDetection) GetPropertyDetectionID() (value string, err error) {
	retValue, err := instance.GetProperty("DetectionID")
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

// SetDetectionSourceTypeID sets the value of DetectionSourceTypeID for the instance
func (instance *MSFT_MpThreatDetection) SetPropertyDetectionSourceTypeID(value uint8) (err error) {
	return instance.SetProperty("DetectionSourceTypeID", (value))
}

// GetDetectionSourceTypeID gets the value of DetectionSourceTypeID for the instance
func (instance *MSFT_MpThreatDetection) GetPropertyDetectionSourceTypeID() (value uint8, err error) {
	retValue, err := instance.GetProperty("DetectionSourceTypeID")
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

// SetDomainUser sets the value of DomainUser for the instance
func (instance *MSFT_MpThreatDetection) SetPropertyDomainUser(value string) (err error) {
	return instance.SetProperty("DomainUser", (value))
}

// GetDomainUser gets the value of DomainUser for the instance
func (instance *MSFT_MpThreatDetection) GetPropertyDomainUser() (value string, err error) {
	retValue, err := instance.GetProperty("DomainUser")
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

// SetInitialDetectionTime sets the value of InitialDetectionTime for the instance
func (instance *MSFT_MpThreatDetection) SetPropertyInitialDetectionTime(value string) (err error) {
	return instance.SetProperty("InitialDetectionTime", (value))
}

// GetInitialDetectionTime gets the value of InitialDetectionTime for the instance
func (instance *MSFT_MpThreatDetection) GetPropertyInitialDetectionTime() (value string, err error) {
	retValue, err := instance.GetProperty("InitialDetectionTime")
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

// SetLastThreatStatusChangeTime sets the value of LastThreatStatusChangeTime for the instance
func (instance *MSFT_MpThreatDetection) SetPropertyLastThreatStatusChangeTime(value string) (err error) {
	return instance.SetProperty("LastThreatStatusChangeTime", (value))
}

// GetLastThreatStatusChangeTime gets the value of LastThreatStatusChangeTime for the instance
func (instance *MSFT_MpThreatDetection) GetPropertyLastThreatStatusChangeTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastThreatStatusChangeTime")
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

// SetProcessName sets the value of ProcessName for the instance
func (instance *MSFT_MpThreatDetection) SetPropertyProcessName(value string) (err error) {
	return instance.SetProperty("ProcessName", (value))
}

// GetProcessName gets the value of ProcessName for the instance
func (instance *MSFT_MpThreatDetection) GetPropertyProcessName() (value string, err error) {
	retValue, err := instance.GetProperty("ProcessName")
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

// SetRemediationTime sets the value of RemediationTime for the instance
func (instance *MSFT_MpThreatDetection) SetPropertyRemediationTime(value string) (err error) {
	return instance.SetProperty("RemediationTime", (value))
}

// GetRemediationTime gets the value of RemediationTime for the instance
func (instance *MSFT_MpThreatDetection) GetPropertyRemediationTime() (value string, err error) {
	retValue, err := instance.GetProperty("RemediationTime")
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

// SetResources sets the value of Resources for the instance
func (instance *MSFT_MpThreatDetection) SetPropertyResources(value []string) (err error) {
	return instance.SetProperty("Resources", (value))
}

// GetResources gets the value of Resources for the instance
func (instance *MSFT_MpThreatDetection) GetPropertyResources() (value []string, err error) {
	retValue, err := instance.GetProperty("Resources")
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

// SetThreatID sets the value of ThreatID for the instance
func (instance *MSFT_MpThreatDetection) SetPropertyThreatID(value int64) (err error) {
	return instance.SetProperty("ThreatID", (value))
}

// GetThreatID gets the value of ThreatID for the instance
func (instance *MSFT_MpThreatDetection) GetPropertyThreatID() (value int64, err error) {
	retValue, err := instance.GetProperty("ThreatID")
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

// SetThreatStatusErrorCode sets the value of ThreatStatusErrorCode for the instance
func (instance *MSFT_MpThreatDetection) SetPropertyThreatStatusErrorCode(value int32) (err error) {
	return instance.SetProperty("ThreatStatusErrorCode", (value))
}

// GetThreatStatusErrorCode gets the value of ThreatStatusErrorCode for the instance
func (instance *MSFT_MpThreatDetection) GetPropertyThreatStatusErrorCode() (value int32, err error) {
	retValue, err := instance.GetProperty("ThreatStatusErrorCode")
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

// SetThreatStatusID sets the value of ThreatStatusID for the instance
func (instance *MSFT_MpThreatDetection) SetPropertyThreatStatusID(value uint8) (err error) {
	return instance.SetProperty("ThreatStatusID", (value))
}

// GetThreatStatusID gets the value of ThreatStatusID for the instance
func (instance *MSFT_MpThreatDetection) GetPropertyThreatStatusID() (value uint8, err error) {
	retValue, err := instance.GetProperty("ThreatStatusID")
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
