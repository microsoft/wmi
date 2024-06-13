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

// Win32_PerfFormattedData_Counters_StorportUnitTransfers struct
type Win32_PerfFormattedData_Counters_StorportUnitTransfers struct {
	*Win32_PerfFormattedData

	//
	SuccessfulTransfersPersecBucket014K uint64

	//
	SuccessfulTransfersPersecBucket028K uint64

	//
	SuccessfulTransfersPersecBucket0316K uint64

	//
	SuccessfulTransfersPersecBucket0432K uint64

	//
	SuccessfulTransfersPersecBucket0564K uint64

	//
	SuccessfulTransfersPersecBucket06128K uint64

	//
	SuccessfulTransfersPersecBucket07256K uint64

	//
	SuccessfulTransfersPersecBucket081M uint64

	//
	SuccessfulTransfersPersecBucket091M uint64

	//
	TransferBytesAverage uint64

	//
	TransferBytesPersec uint64

	//
	TransferLatency uint64

	//
	TransferLatencyBucket01128us uint64

	//
	TransferLatencyBucket02256us uint64

	//
	TransferLatencyBucket03512us uint64

	//
	TransferLatencyBucket041ms uint64

	//
	TransferLatencyBucket054ms uint64

	//
	TransferLatencyBucket0616ms uint64

	//
	TransferLatencyBucket0764ms uint64

	//
	TransferLatencyBucket08128ms uint64

	//
	TransferLatencyBucket09256ms uint64

	//
	TransferLatencyBucket10512ms uint64

	//
	TransferLatencyBucket111s uint64

	//
	TransferLatencyBucket122s uint64

	//
	TransferLatencyBucket1310s uint64

	//
	TransferLatencyBucket1410s uint64

	//
	TransfersPersec uint64

	//
	TransfersPersecBucket01128us uint64

	//
	TransfersPersecBucket02256us uint64

	//
	TransfersPersecBucket03512us uint64

	//
	TransfersPersecBucket041ms uint64

	//
	TransfersPersecBucket054ms uint64

	//
	TransfersPersecBucket0616ms uint64

	//
	TransfersPersecBucket0764ms uint64

	//
	TransfersPersecBucket08128ms uint64

	//
	TransfersPersecBucket09256ms uint64

	//
	TransfersPersecBucket10512ms uint64

	//
	TransfersPersecBucket111s uint64

	//
	TransfersPersecBucket122s uint64

	//
	TransfersPersecBucket1310s uint64

	//
	TransfersPersecBucket1410s uint64
}

func NewWin32_PerfFormattedData_Counters_StorportUnitTransfersEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Counters_StorportUnitTransfers, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_StorportUnitTransfers{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Counters_StorportUnitTransfersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Counters_StorportUnitTransfers, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_StorportUnitTransfers{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetSuccessfulTransfersPersecBucket014K sets the value of SuccessfulTransfersPersecBucket014K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertySuccessfulTransfersPersecBucket014K(value uint64) (err error) {
	return instance.SetProperty("SuccessfulTransfersPersecBucket014K", (value))
}

// GetSuccessfulTransfersPersecBucket014K gets the value of SuccessfulTransfersPersecBucket014K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertySuccessfulTransfersPersecBucket014K() (value uint64, err error) {
	retValue, err := instance.GetProperty("SuccessfulTransfersPersecBucket014K")
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

// SetSuccessfulTransfersPersecBucket028K sets the value of SuccessfulTransfersPersecBucket028K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertySuccessfulTransfersPersecBucket028K(value uint64) (err error) {
	return instance.SetProperty("SuccessfulTransfersPersecBucket028K", (value))
}

// GetSuccessfulTransfersPersecBucket028K gets the value of SuccessfulTransfersPersecBucket028K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertySuccessfulTransfersPersecBucket028K() (value uint64, err error) {
	retValue, err := instance.GetProperty("SuccessfulTransfersPersecBucket028K")
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

// SetSuccessfulTransfersPersecBucket0316K sets the value of SuccessfulTransfersPersecBucket0316K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertySuccessfulTransfersPersecBucket0316K(value uint64) (err error) {
	return instance.SetProperty("SuccessfulTransfersPersecBucket0316K", (value))
}

// GetSuccessfulTransfersPersecBucket0316K gets the value of SuccessfulTransfersPersecBucket0316K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertySuccessfulTransfersPersecBucket0316K() (value uint64, err error) {
	retValue, err := instance.GetProperty("SuccessfulTransfersPersecBucket0316K")
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

// SetSuccessfulTransfersPersecBucket0432K sets the value of SuccessfulTransfersPersecBucket0432K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertySuccessfulTransfersPersecBucket0432K(value uint64) (err error) {
	return instance.SetProperty("SuccessfulTransfersPersecBucket0432K", (value))
}

// GetSuccessfulTransfersPersecBucket0432K gets the value of SuccessfulTransfersPersecBucket0432K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertySuccessfulTransfersPersecBucket0432K() (value uint64, err error) {
	retValue, err := instance.GetProperty("SuccessfulTransfersPersecBucket0432K")
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

// SetSuccessfulTransfersPersecBucket0564K sets the value of SuccessfulTransfersPersecBucket0564K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertySuccessfulTransfersPersecBucket0564K(value uint64) (err error) {
	return instance.SetProperty("SuccessfulTransfersPersecBucket0564K", (value))
}

// GetSuccessfulTransfersPersecBucket0564K gets the value of SuccessfulTransfersPersecBucket0564K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertySuccessfulTransfersPersecBucket0564K() (value uint64, err error) {
	retValue, err := instance.GetProperty("SuccessfulTransfersPersecBucket0564K")
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

// SetSuccessfulTransfersPersecBucket06128K sets the value of SuccessfulTransfersPersecBucket06128K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertySuccessfulTransfersPersecBucket06128K(value uint64) (err error) {
	return instance.SetProperty("SuccessfulTransfersPersecBucket06128K", (value))
}

// GetSuccessfulTransfersPersecBucket06128K gets the value of SuccessfulTransfersPersecBucket06128K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertySuccessfulTransfersPersecBucket06128K() (value uint64, err error) {
	retValue, err := instance.GetProperty("SuccessfulTransfersPersecBucket06128K")
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

// SetSuccessfulTransfersPersecBucket07256K sets the value of SuccessfulTransfersPersecBucket07256K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertySuccessfulTransfersPersecBucket07256K(value uint64) (err error) {
	return instance.SetProperty("SuccessfulTransfersPersecBucket07256K", (value))
}

// GetSuccessfulTransfersPersecBucket07256K gets the value of SuccessfulTransfersPersecBucket07256K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertySuccessfulTransfersPersecBucket07256K() (value uint64, err error) {
	retValue, err := instance.GetProperty("SuccessfulTransfersPersecBucket07256K")
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

// SetSuccessfulTransfersPersecBucket081M sets the value of SuccessfulTransfersPersecBucket081M for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertySuccessfulTransfersPersecBucket081M(value uint64) (err error) {
	return instance.SetProperty("SuccessfulTransfersPersecBucket081M", (value))
}

// GetSuccessfulTransfersPersecBucket081M gets the value of SuccessfulTransfersPersecBucket081M for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertySuccessfulTransfersPersecBucket081M() (value uint64, err error) {
	retValue, err := instance.GetProperty("SuccessfulTransfersPersecBucket081M")
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

// SetSuccessfulTransfersPersecBucket091M sets the value of SuccessfulTransfersPersecBucket091M for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertySuccessfulTransfersPersecBucket091M(value uint64) (err error) {
	return instance.SetProperty("SuccessfulTransfersPersecBucket091M", (value))
}

// GetSuccessfulTransfersPersecBucket091M gets the value of SuccessfulTransfersPersecBucket091M for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertySuccessfulTransfersPersecBucket091M() (value uint64, err error) {
	retValue, err := instance.GetProperty("SuccessfulTransfersPersecBucket091M")
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

// SetTransferBytesAverage sets the value of TransferBytesAverage for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransferBytesAverage(value uint64) (err error) {
	return instance.SetProperty("TransferBytesAverage", (value))
}

// GetTransferBytesAverage gets the value of TransferBytesAverage for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransferBytesAverage() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransferBytesAverage")
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

// SetTransferBytesPersec sets the value of TransferBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransferBytesPersec(value uint64) (err error) {
	return instance.SetProperty("TransferBytesPersec", (value))
}

// GetTransferBytesPersec gets the value of TransferBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransferBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransferBytesPersec")
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

// SetTransferLatency sets the value of TransferLatency for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransferLatency(value uint64) (err error) {
	return instance.SetProperty("TransferLatency", (value))
}

// GetTransferLatency gets the value of TransferLatency for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransferLatency() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransferLatency")
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

// SetTransferLatencyBucket01128us sets the value of TransferLatencyBucket01128us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransferLatencyBucket01128us(value uint64) (err error) {
	return instance.SetProperty("TransferLatencyBucket01128us", (value))
}

// GetTransferLatencyBucket01128us gets the value of TransferLatencyBucket01128us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransferLatencyBucket01128us() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransferLatencyBucket01128us")
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

// SetTransferLatencyBucket02256us sets the value of TransferLatencyBucket02256us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransferLatencyBucket02256us(value uint64) (err error) {
	return instance.SetProperty("TransferLatencyBucket02256us", (value))
}

// GetTransferLatencyBucket02256us gets the value of TransferLatencyBucket02256us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransferLatencyBucket02256us() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransferLatencyBucket02256us")
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

// SetTransferLatencyBucket03512us sets the value of TransferLatencyBucket03512us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransferLatencyBucket03512us(value uint64) (err error) {
	return instance.SetProperty("TransferLatencyBucket03512us", (value))
}

// GetTransferLatencyBucket03512us gets the value of TransferLatencyBucket03512us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransferLatencyBucket03512us() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransferLatencyBucket03512us")
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

// SetTransferLatencyBucket041ms sets the value of TransferLatencyBucket041ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransferLatencyBucket041ms(value uint64) (err error) {
	return instance.SetProperty("TransferLatencyBucket041ms", (value))
}

// GetTransferLatencyBucket041ms gets the value of TransferLatencyBucket041ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransferLatencyBucket041ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransferLatencyBucket041ms")
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

// SetTransferLatencyBucket054ms sets the value of TransferLatencyBucket054ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransferLatencyBucket054ms(value uint64) (err error) {
	return instance.SetProperty("TransferLatencyBucket054ms", (value))
}

// GetTransferLatencyBucket054ms gets the value of TransferLatencyBucket054ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransferLatencyBucket054ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransferLatencyBucket054ms")
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

// SetTransferLatencyBucket0616ms sets the value of TransferLatencyBucket0616ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransferLatencyBucket0616ms(value uint64) (err error) {
	return instance.SetProperty("TransferLatencyBucket0616ms", (value))
}

// GetTransferLatencyBucket0616ms gets the value of TransferLatencyBucket0616ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransferLatencyBucket0616ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransferLatencyBucket0616ms")
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

// SetTransferLatencyBucket0764ms sets the value of TransferLatencyBucket0764ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransferLatencyBucket0764ms(value uint64) (err error) {
	return instance.SetProperty("TransferLatencyBucket0764ms", (value))
}

// GetTransferLatencyBucket0764ms gets the value of TransferLatencyBucket0764ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransferLatencyBucket0764ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransferLatencyBucket0764ms")
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

// SetTransferLatencyBucket08128ms sets the value of TransferLatencyBucket08128ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransferLatencyBucket08128ms(value uint64) (err error) {
	return instance.SetProperty("TransferLatencyBucket08128ms", (value))
}

// GetTransferLatencyBucket08128ms gets the value of TransferLatencyBucket08128ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransferLatencyBucket08128ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransferLatencyBucket08128ms")
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

// SetTransferLatencyBucket09256ms sets the value of TransferLatencyBucket09256ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransferLatencyBucket09256ms(value uint64) (err error) {
	return instance.SetProperty("TransferLatencyBucket09256ms", (value))
}

// GetTransferLatencyBucket09256ms gets the value of TransferLatencyBucket09256ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransferLatencyBucket09256ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransferLatencyBucket09256ms")
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

// SetTransferLatencyBucket10512ms sets the value of TransferLatencyBucket10512ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransferLatencyBucket10512ms(value uint64) (err error) {
	return instance.SetProperty("TransferLatencyBucket10512ms", (value))
}

// GetTransferLatencyBucket10512ms gets the value of TransferLatencyBucket10512ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransferLatencyBucket10512ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransferLatencyBucket10512ms")
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

// SetTransferLatencyBucket111s sets the value of TransferLatencyBucket111s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransferLatencyBucket111s(value uint64) (err error) {
	return instance.SetProperty("TransferLatencyBucket111s", (value))
}

// GetTransferLatencyBucket111s gets the value of TransferLatencyBucket111s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransferLatencyBucket111s() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransferLatencyBucket111s")
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

// SetTransferLatencyBucket122s sets the value of TransferLatencyBucket122s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransferLatencyBucket122s(value uint64) (err error) {
	return instance.SetProperty("TransferLatencyBucket122s", (value))
}

// GetTransferLatencyBucket122s gets the value of TransferLatencyBucket122s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransferLatencyBucket122s() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransferLatencyBucket122s")
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

// SetTransferLatencyBucket1310s sets the value of TransferLatencyBucket1310s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransferLatencyBucket1310s(value uint64) (err error) {
	return instance.SetProperty("TransferLatencyBucket1310s", (value))
}

// GetTransferLatencyBucket1310s gets the value of TransferLatencyBucket1310s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransferLatencyBucket1310s() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransferLatencyBucket1310s")
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

// SetTransferLatencyBucket1410s sets the value of TransferLatencyBucket1410s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransferLatencyBucket1410s(value uint64) (err error) {
	return instance.SetProperty("TransferLatencyBucket1410s", (value))
}

// GetTransferLatencyBucket1410s gets the value of TransferLatencyBucket1410s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransferLatencyBucket1410s() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransferLatencyBucket1410s")
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

// SetTransfersPersec sets the value of TransfersPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransfersPersec(value uint64) (err error) {
	return instance.SetProperty("TransfersPersec", (value))
}

// GetTransfersPersec gets the value of TransfersPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransfersPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransfersPersec")
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

// SetTransfersPersecBucket01128us sets the value of TransfersPersecBucket01128us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransfersPersecBucket01128us(value uint64) (err error) {
	return instance.SetProperty("TransfersPersecBucket01128us", (value))
}

// GetTransfersPersecBucket01128us gets the value of TransfersPersecBucket01128us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransfersPersecBucket01128us() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransfersPersecBucket01128us")
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

// SetTransfersPersecBucket02256us sets the value of TransfersPersecBucket02256us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransfersPersecBucket02256us(value uint64) (err error) {
	return instance.SetProperty("TransfersPersecBucket02256us", (value))
}

// GetTransfersPersecBucket02256us gets the value of TransfersPersecBucket02256us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransfersPersecBucket02256us() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransfersPersecBucket02256us")
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

// SetTransfersPersecBucket03512us sets the value of TransfersPersecBucket03512us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransfersPersecBucket03512us(value uint64) (err error) {
	return instance.SetProperty("TransfersPersecBucket03512us", (value))
}

// GetTransfersPersecBucket03512us gets the value of TransfersPersecBucket03512us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransfersPersecBucket03512us() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransfersPersecBucket03512us")
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

// SetTransfersPersecBucket041ms sets the value of TransfersPersecBucket041ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransfersPersecBucket041ms(value uint64) (err error) {
	return instance.SetProperty("TransfersPersecBucket041ms", (value))
}

// GetTransfersPersecBucket041ms gets the value of TransfersPersecBucket041ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransfersPersecBucket041ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransfersPersecBucket041ms")
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

// SetTransfersPersecBucket054ms sets the value of TransfersPersecBucket054ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransfersPersecBucket054ms(value uint64) (err error) {
	return instance.SetProperty("TransfersPersecBucket054ms", (value))
}

// GetTransfersPersecBucket054ms gets the value of TransfersPersecBucket054ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransfersPersecBucket054ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransfersPersecBucket054ms")
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

// SetTransfersPersecBucket0616ms sets the value of TransfersPersecBucket0616ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransfersPersecBucket0616ms(value uint64) (err error) {
	return instance.SetProperty("TransfersPersecBucket0616ms", (value))
}

// GetTransfersPersecBucket0616ms gets the value of TransfersPersecBucket0616ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransfersPersecBucket0616ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransfersPersecBucket0616ms")
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

// SetTransfersPersecBucket0764ms sets the value of TransfersPersecBucket0764ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransfersPersecBucket0764ms(value uint64) (err error) {
	return instance.SetProperty("TransfersPersecBucket0764ms", (value))
}

// GetTransfersPersecBucket0764ms gets the value of TransfersPersecBucket0764ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransfersPersecBucket0764ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransfersPersecBucket0764ms")
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

// SetTransfersPersecBucket08128ms sets the value of TransfersPersecBucket08128ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransfersPersecBucket08128ms(value uint64) (err error) {
	return instance.SetProperty("TransfersPersecBucket08128ms", (value))
}

// GetTransfersPersecBucket08128ms gets the value of TransfersPersecBucket08128ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransfersPersecBucket08128ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransfersPersecBucket08128ms")
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

// SetTransfersPersecBucket09256ms sets the value of TransfersPersecBucket09256ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransfersPersecBucket09256ms(value uint64) (err error) {
	return instance.SetProperty("TransfersPersecBucket09256ms", (value))
}

// GetTransfersPersecBucket09256ms gets the value of TransfersPersecBucket09256ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransfersPersecBucket09256ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransfersPersecBucket09256ms")
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

// SetTransfersPersecBucket10512ms sets the value of TransfersPersecBucket10512ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransfersPersecBucket10512ms(value uint64) (err error) {
	return instance.SetProperty("TransfersPersecBucket10512ms", (value))
}

// GetTransfersPersecBucket10512ms gets the value of TransfersPersecBucket10512ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransfersPersecBucket10512ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransfersPersecBucket10512ms")
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

// SetTransfersPersecBucket111s sets the value of TransfersPersecBucket111s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransfersPersecBucket111s(value uint64) (err error) {
	return instance.SetProperty("TransfersPersecBucket111s", (value))
}

// GetTransfersPersecBucket111s gets the value of TransfersPersecBucket111s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransfersPersecBucket111s() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransfersPersecBucket111s")
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

// SetTransfersPersecBucket122s sets the value of TransfersPersecBucket122s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransfersPersecBucket122s(value uint64) (err error) {
	return instance.SetProperty("TransfersPersecBucket122s", (value))
}

// GetTransfersPersecBucket122s gets the value of TransfersPersecBucket122s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransfersPersecBucket122s() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransfersPersecBucket122s")
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

// SetTransfersPersecBucket1310s sets the value of TransfersPersecBucket1310s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransfersPersecBucket1310s(value uint64) (err error) {
	return instance.SetProperty("TransfersPersecBucket1310s", (value))
}

// GetTransfersPersecBucket1310s gets the value of TransfersPersecBucket1310s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransfersPersecBucket1310s() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransfersPersecBucket1310s")
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

// SetTransfersPersecBucket1410s sets the value of TransfersPersecBucket1410s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) SetPropertyTransfersPersecBucket1410s(value uint64) (err error) {
	return instance.SetProperty("TransfersPersecBucket1410s", (value))
}

// GetTransfersPersecBucket1410s gets the value of TransfersPersecBucket1410s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitTransfers) GetPropertyTransfersPersecBucket1410s() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransfersPersecBucket1410s")
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
