// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MS_SMHBA_SASPHYSTATISTICS struct
type MS_SMHBA_SASPHYSTATISTICS struct {
	*cim.WmiInstance

	//
	InvalidDwordCount int64

	//
	LossofDwordSyncCount int64

	//
	PhyResetProblemCount int64

	//
	RunningDisparityErrorCount int64

	//
	RxFrames int64

	//
	RxWords int64

	//
	SecondsSinceLastReset int64

	//
	TxFrames int64

	//
	TxWords int64
}

func NewMS_SMHBA_SASPHYSTATISTICSEx1(instance *cim.WmiInstance) (newInstance *MS_SMHBA_SASPHYSTATISTICS, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MS_SMHBA_SASPHYSTATISTICS{
		WmiInstance: tmp,
	}
	return
}

func NewMS_SMHBA_SASPHYSTATISTICSEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MS_SMHBA_SASPHYSTATISTICS, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MS_SMHBA_SASPHYSTATISTICS{
		WmiInstance: tmp,
	}
	return
}

// SetInvalidDwordCount sets the value of InvalidDwordCount for the instance
func (instance *MS_SMHBA_SASPHYSTATISTICS) SetPropertyInvalidDwordCount(value int64) (err error) {
	return instance.SetProperty("InvalidDwordCount", (value))
}

// GetInvalidDwordCount gets the value of InvalidDwordCount for the instance
func (instance *MS_SMHBA_SASPHYSTATISTICS) GetPropertyInvalidDwordCount() (value int64, err error) {
	retValue, err := instance.GetProperty("InvalidDwordCount")
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

// SetLossofDwordSyncCount sets the value of LossofDwordSyncCount for the instance
func (instance *MS_SMHBA_SASPHYSTATISTICS) SetPropertyLossofDwordSyncCount(value int64) (err error) {
	return instance.SetProperty("LossofDwordSyncCount", (value))
}

// GetLossofDwordSyncCount gets the value of LossofDwordSyncCount for the instance
func (instance *MS_SMHBA_SASPHYSTATISTICS) GetPropertyLossofDwordSyncCount() (value int64, err error) {
	retValue, err := instance.GetProperty("LossofDwordSyncCount")
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

// SetPhyResetProblemCount sets the value of PhyResetProblemCount for the instance
func (instance *MS_SMHBA_SASPHYSTATISTICS) SetPropertyPhyResetProblemCount(value int64) (err error) {
	return instance.SetProperty("PhyResetProblemCount", (value))
}

// GetPhyResetProblemCount gets the value of PhyResetProblemCount for the instance
func (instance *MS_SMHBA_SASPHYSTATISTICS) GetPropertyPhyResetProblemCount() (value int64, err error) {
	retValue, err := instance.GetProperty("PhyResetProblemCount")
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

// SetRunningDisparityErrorCount sets the value of RunningDisparityErrorCount for the instance
func (instance *MS_SMHBA_SASPHYSTATISTICS) SetPropertyRunningDisparityErrorCount(value int64) (err error) {
	return instance.SetProperty("RunningDisparityErrorCount", (value))
}

// GetRunningDisparityErrorCount gets the value of RunningDisparityErrorCount for the instance
func (instance *MS_SMHBA_SASPHYSTATISTICS) GetPropertyRunningDisparityErrorCount() (value int64, err error) {
	retValue, err := instance.GetProperty("RunningDisparityErrorCount")
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

// SetRxFrames sets the value of RxFrames for the instance
func (instance *MS_SMHBA_SASPHYSTATISTICS) SetPropertyRxFrames(value int64) (err error) {
	return instance.SetProperty("RxFrames", (value))
}

// GetRxFrames gets the value of RxFrames for the instance
func (instance *MS_SMHBA_SASPHYSTATISTICS) GetPropertyRxFrames() (value int64, err error) {
	retValue, err := instance.GetProperty("RxFrames")
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

// SetRxWords sets the value of RxWords for the instance
func (instance *MS_SMHBA_SASPHYSTATISTICS) SetPropertyRxWords(value int64) (err error) {
	return instance.SetProperty("RxWords", (value))
}

// GetRxWords gets the value of RxWords for the instance
func (instance *MS_SMHBA_SASPHYSTATISTICS) GetPropertyRxWords() (value int64, err error) {
	retValue, err := instance.GetProperty("RxWords")
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

// SetSecondsSinceLastReset sets the value of SecondsSinceLastReset for the instance
func (instance *MS_SMHBA_SASPHYSTATISTICS) SetPropertySecondsSinceLastReset(value int64) (err error) {
	return instance.SetProperty("SecondsSinceLastReset", (value))
}

// GetSecondsSinceLastReset gets the value of SecondsSinceLastReset for the instance
func (instance *MS_SMHBA_SASPHYSTATISTICS) GetPropertySecondsSinceLastReset() (value int64, err error) {
	retValue, err := instance.GetProperty("SecondsSinceLastReset")
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

// SetTxFrames sets the value of TxFrames for the instance
func (instance *MS_SMHBA_SASPHYSTATISTICS) SetPropertyTxFrames(value int64) (err error) {
	return instance.SetProperty("TxFrames", (value))
}

// GetTxFrames gets the value of TxFrames for the instance
func (instance *MS_SMHBA_SASPHYSTATISTICS) GetPropertyTxFrames() (value int64, err error) {
	retValue, err := instance.GetProperty("TxFrames")
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

// SetTxWords sets the value of TxWords for the instance
func (instance *MS_SMHBA_SASPHYSTATISTICS) SetPropertyTxWords(value int64) (err error) {
	return instance.SetProperty("TxWords", (value))
}

// GetTxWords gets the value of TxWords for the instance
func (instance *MS_SMHBA_SASPHYSTATISTICS) GetPropertyTxWords() (value int64, err error) {
	retValue, err := instance.GetProperty("TxWords")
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
