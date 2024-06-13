// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFC_HBAPortStatistics struct
type MSFC_HBAPortStatistics struct {
	*cim.WmiInstance

	//
	DumpedFrames int64

	//
	ErrorFrames int64

	//
	InvalidCRCCount int64

	//
	InvalidTxWordCount int64

	//
	LinkFailureCount int64

	//
	LIPCount int64

	//
	LossOfSignalCount int64

	//
	LossOfSyncCount int64

	//
	NOSCount int64

	//
	PrimitiveSeqProtocolErrCount int64

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

func NewMSFC_HBAPortStatisticsEx1(instance *cim.WmiInstance) (newInstance *MSFC_HBAPortStatistics, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFC_HBAPortStatistics{
		WmiInstance: tmp,
	}
	return
}

func NewMSFC_HBAPortStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFC_HBAPortStatistics, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFC_HBAPortStatistics{
		WmiInstance: tmp,
	}
	return
}

// SetDumpedFrames sets the value of DumpedFrames for the instance
func (instance *MSFC_HBAPortStatistics) SetPropertyDumpedFrames(value int64) (err error) {
	return instance.SetProperty("DumpedFrames", (value))
}

// GetDumpedFrames gets the value of DumpedFrames for the instance
func (instance *MSFC_HBAPortStatistics) GetPropertyDumpedFrames() (value int64, err error) {
	retValue, err := instance.GetProperty("DumpedFrames")
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

// SetErrorFrames sets the value of ErrorFrames for the instance
func (instance *MSFC_HBAPortStatistics) SetPropertyErrorFrames(value int64) (err error) {
	return instance.SetProperty("ErrorFrames", (value))
}

// GetErrorFrames gets the value of ErrorFrames for the instance
func (instance *MSFC_HBAPortStatistics) GetPropertyErrorFrames() (value int64, err error) {
	retValue, err := instance.GetProperty("ErrorFrames")
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

// SetInvalidCRCCount sets the value of InvalidCRCCount for the instance
func (instance *MSFC_HBAPortStatistics) SetPropertyInvalidCRCCount(value int64) (err error) {
	return instance.SetProperty("InvalidCRCCount", (value))
}

// GetInvalidCRCCount gets the value of InvalidCRCCount for the instance
func (instance *MSFC_HBAPortStatistics) GetPropertyInvalidCRCCount() (value int64, err error) {
	retValue, err := instance.GetProperty("InvalidCRCCount")
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

// SetInvalidTxWordCount sets the value of InvalidTxWordCount for the instance
func (instance *MSFC_HBAPortStatistics) SetPropertyInvalidTxWordCount(value int64) (err error) {
	return instance.SetProperty("InvalidTxWordCount", (value))
}

// GetInvalidTxWordCount gets the value of InvalidTxWordCount for the instance
func (instance *MSFC_HBAPortStatistics) GetPropertyInvalidTxWordCount() (value int64, err error) {
	retValue, err := instance.GetProperty("InvalidTxWordCount")
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

// SetLinkFailureCount sets the value of LinkFailureCount for the instance
func (instance *MSFC_HBAPortStatistics) SetPropertyLinkFailureCount(value int64) (err error) {
	return instance.SetProperty("LinkFailureCount", (value))
}

// GetLinkFailureCount gets the value of LinkFailureCount for the instance
func (instance *MSFC_HBAPortStatistics) GetPropertyLinkFailureCount() (value int64, err error) {
	retValue, err := instance.GetProperty("LinkFailureCount")
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

// SetLIPCount sets the value of LIPCount for the instance
func (instance *MSFC_HBAPortStatistics) SetPropertyLIPCount(value int64) (err error) {
	return instance.SetProperty("LIPCount", (value))
}

// GetLIPCount gets the value of LIPCount for the instance
func (instance *MSFC_HBAPortStatistics) GetPropertyLIPCount() (value int64, err error) {
	retValue, err := instance.GetProperty("LIPCount")
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

// SetLossOfSignalCount sets the value of LossOfSignalCount for the instance
func (instance *MSFC_HBAPortStatistics) SetPropertyLossOfSignalCount(value int64) (err error) {
	return instance.SetProperty("LossOfSignalCount", (value))
}

// GetLossOfSignalCount gets the value of LossOfSignalCount for the instance
func (instance *MSFC_HBAPortStatistics) GetPropertyLossOfSignalCount() (value int64, err error) {
	retValue, err := instance.GetProperty("LossOfSignalCount")
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

// SetLossOfSyncCount sets the value of LossOfSyncCount for the instance
func (instance *MSFC_HBAPortStatistics) SetPropertyLossOfSyncCount(value int64) (err error) {
	return instance.SetProperty("LossOfSyncCount", (value))
}

// GetLossOfSyncCount gets the value of LossOfSyncCount for the instance
func (instance *MSFC_HBAPortStatistics) GetPropertyLossOfSyncCount() (value int64, err error) {
	retValue, err := instance.GetProperty("LossOfSyncCount")
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

// SetNOSCount sets the value of NOSCount for the instance
func (instance *MSFC_HBAPortStatistics) SetPropertyNOSCount(value int64) (err error) {
	return instance.SetProperty("NOSCount", (value))
}

// GetNOSCount gets the value of NOSCount for the instance
func (instance *MSFC_HBAPortStatistics) GetPropertyNOSCount() (value int64, err error) {
	retValue, err := instance.GetProperty("NOSCount")
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

// SetPrimitiveSeqProtocolErrCount sets the value of PrimitiveSeqProtocolErrCount for the instance
func (instance *MSFC_HBAPortStatistics) SetPropertyPrimitiveSeqProtocolErrCount(value int64) (err error) {
	return instance.SetProperty("PrimitiveSeqProtocolErrCount", (value))
}

// GetPrimitiveSeqProtocolErrCount gets the value of PrimitiveSeqProtocolErrCount for the instance
func (instance *MSFC_HBAPortStatistics) GetPropertyPrimitiveSeqProtocolErrCount() (value int64, err error) {
	retValue, err := instance.GetProperty("PrimitiveSeqProtocolErrCount")
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
func (instance *MSFC_HBAPortStatistics) SetPropertyRxFrames(value int64) (err error) {
	return instance.SetProperty("RxFrames", (value))
}

// GetRxFrames gets the value of RxFrames for the instance
func (instance *MSFC_HBAPortStatistics) GetPropertyRxFrames() (value int64, err error) {
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
func (instance *MSFC_HBAPortStatistics) SetPropertyRxWords(value int64) (err error) {
	return instance.SetProperty("RxWords", (value))
}

// GetRxWords gets the value of RxWords for the instance
func (instance *MSFC_HBAPortStatistics) GetPropertyRxWords() (value int64, err error) {
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
func (instance *MSFC_HBAPortStatistics) SetPropertySecondsSinceLastReset(value int64) (err error) {
	return instance.SetProperty("SecondsSinceLastReset", (value))
}

// GetSecondsSinceLastReset gets the value of SecondsSinceLastReset for the instance
func (instance *MSFC_HBAPortStatistics) GetPropertySecondsSinceLastReset() (value int64, err error) {
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
func (instance *MSFC_HBAPortStatistics) SetPropertyTxFrames(value int64) (err error) {
	return instance.SetProperty("TxFrames", (value))
}

// GetTxFrames gets the value of TxFrames for the instance
func (instance *MSFC_HBAPortStatistics) GetPropertyTxFrames() (value int64, err error) {
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
func (instance *MSFC_HBAPortStatistics) SetPropertyTxWords(value int64) (err error) {
	return instance.SetProperty("TxWords", (value))
}

// GetTxWords gets the value of TxWords for the instance
func (instance *MSFC_HBAPortStatistics) GetPropertyTxWords() (value int64, err error) {
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
