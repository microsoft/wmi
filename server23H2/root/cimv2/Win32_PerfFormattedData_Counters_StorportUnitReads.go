// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_Counters_StorportUnitReads struct
type Win32_PerfFormattedData_Counters_StorportUnitReads struct {
	*Win32_PerfFormattedData

	//
	ReadBytesAverage uint64

	//
	ReadBytesPersec uint64

	//
	ReadLatency uint64

	//
	ReadLatencyBucket01128us uint64

	//
	ReadLatencyBucket02256us uint64

	//
	ReadLatencyBucket03512us uint64

	//
	ReadLatencyBucket041ms uint64

	//
	ReadLatencyBucket054ms uint64

	//
	ReadLatencyBucket0616ms uint64

	//
	ReadLatencyBucket0764ms uint64

	//
	ReadLatencyBucket08128ms uint64

	//
	ReadLatencyBucket09256ms uint64

	//
	ReadLatencyBucket10512ms uint64

	//
	ReadLatencyBucket111s uint64

	//
	ReadLatencyBucket122s uint64

	//
	ReadLatencyBucket1310s uint64

	//
	ReadLatencyBucket1410s uint64

	//
	ReadsPersec uint64

	//
	ReadsPersecBucket01128us uint64

	//
	ReadsPersecBucket02256us uint64

	//
	ReadsPersecBucket03512us uint64

	//
	ReadsPersecBucket041ms uint64

	//
	ReadsPersecBucket054ms uint64

	//
	ReadsPersecBucket0616ms uint64

	//
	ReadsPersecBucket0764ms uint64

	//
	ReadsPersecBucket08128ms uint64

	//
	ReadsPersecBucket09256ms uint64

	//
	ReadsPersecBucket10512ms uint64

	//
	ReadsPersecBucket111s uint64

	//
	ReadsPersecBucket122s uint64

	//
	ReadsPersecBucket1310s uint64

	//
	ReadsPersecBucket1410s uint64

	//
	SuccessfulReadsPersecBucket014K uint64

	//
	SuccessfulReadsPersecBucket028K uint64

	//
	SuccessfulReadsPersecBucket0316K uint64

	//
	SuccessfulReadsPersecBucket0432K uint64

	//
	SuccessfulReadsPersecBucket0564K uint64

	//
	SuccessfulReadsPersecBucket06128K uint64

	//
	SuccessfulReadsPersecBucket07256K uint64

	//
	SuccessfulReadsPersecBucket081M uint64

	//
	SuccessfulReadsPersecBucket091M uint64
}

func NewWin32_PerfFormattedData_Counters_StorportUnitReadsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Counters_StorportUnitReads, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_StorportUnitReads{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Counters_StorportUnitReadsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Counters_StorportUnitReads, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_StorportUnitReads{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetReadBytesAverage sets the value of ReadBytesAverage for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadBytesAverage(value uint64) (err error) {
	return instance.SetProperty("ReadBytesAverage", (value))
}

// GetReadBytesAverage gets the value of ReadBytesAverage for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadBytesAverage() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadBytesAverage")
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

// SetReadBytesPersec sets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadBytesPersec(value uint64) (err error) {
	return instance.SetProperty("ReadBytesPersec", (value))
}

// GetReadBytesPersec gets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadBytesPersec")
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

// SetReadLatency sets the value of ReadLatency for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadLatency(value uint64) (err error) {
	return instance.SetProperty("ReadLatency", (value))
}

// GetReadLatency gets the value of ReadLatency for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadLatency() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadLatency")
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

// SetReadLatencyBucket01128us sets the value of ReadLatencyBucket01128us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadLatencyBucket01128us(value uint64) (err error) {
	return instance.SetProperty("ReadLatencyBucket01128us", (value))
}

// GetReadLatencyBucket01128us gets the value of ReadLatencyBucket01128us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadLatencyBucket01128us() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadLatencyBucket01128us")
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

// SetReadLatencyBucket02256us sets the value of ReadLatencyBucket02256us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadLatencyBucket02256us(value uint64) (err error) {
	return instance.SetProperty("ReadLatencyBucket02256us", (value))
}

// GetReadLatencyBucket02256us gets the value of ReadLatencyBucket02256us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadLatencyBucket02256us() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadLatencyBucket02256us")
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

// SetReadLatencyBucket03512us sets the value of ReadLatencyBucket03512us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadLatencyBucket03512us(value uint64) (err error) {
	return instance.SetProperty("ReadLatencyBucket03512us", (value))
}

// GetReadLatencyBucket03512us gets the value of ReadLatencyBucket03512us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadLatencyBucket03512us() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadLatencyBucket03512us")
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

// SetReadLatencyBucket041ms sets the value of ReadLatencyBucket041ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadLatencyBucket041ms(value uint64) (err error) {
	return instance.SetProperty("ReadLatencyBucket041ms", (value))
}

// GetReadLatencyBucket041ms gets the value of ReadLatencyBucket041ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadLatencyBucket041ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadLatencyBucket041ms")
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

// SetReadLatencyBucket054ms sets the value of ReadLatencyBucket054ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadLatencyBucket054ms(value uint64) (err error) {
	return instance.SetProperty("ReadLatencyBucket054ms", (value))
}

// GetReadLatencyBucket054ms gets the value of ReadLatencyBucket054ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadLatencyBucket054ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadLatencyBucket054ms")
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

// SetReadLatencyBucket0616ms sets the value of ReadLatencyBucket0616ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadLatencyBucket0616ms(value uint64) (err error) {
	return instance.SetProperty("ReadLatencyBucket0616ms", (value))
}

// GetReadLatencyBucket0616ms gets the value of ReadLatencyBucket0616ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadLatencyBucket0616ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadLatencyBucket0616ms")
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

// SetReadLatencyBucket0764ms sets the value of ReadLatencyBucket0764ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadLatencyBucket0764ms(value uint64) (err error) {
	return instance.SetProperty("ReadLatencyBucket0764ms", (value))
}

// GetReadLatencyBucket0764ms gets the value of ReadLatencyBucket0764ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadLatencyBucket0764ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadLatencyBucket0764ms")
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

// SetReadLatencyBucket08128ms sets the value of ReadLatencyBucket08128ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadLatencyBucket08128ms(value uint64) (err error) {
	return instance.SetProperty("ReadLatencyBucket08128ms", (value))
}

// GetReadLatencyBucket08128ms gets the value of ReadLatencyBucket08128ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadLatencyBucket08128ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadLatencyBucket08128ms")
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

// SetReadLatencyBucket09256ms sets the value of ReadLatencyBucket09256ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadLatencyBucket09256ms(value uint64) (err error) {
	return instance.SetProperty("ReadLatencyBucket09256ms", (value))
}

// GetReadLatencyBucket09256ms gets the value of ReadLatencyBucket09256ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadLatencyBucket09256ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadLatencyBucket09256ms")
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

// SetReadLatencyBucket10512ms sets the value of ReadLatencyBucket10512ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadLatencyBucket10512ms(value uint64) (err error) {
	return instance.SetProperty("ReadLatencyBucket10512ms", (value))
}

// GetReadLatencyBucket10512ms gets the value of ReadLatencyBucket10512ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadLatencyBucket10512ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadLatencyBucket10512ms")
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

// SetReadLatencyBucket111s sets the value of ReadLatencyBucket111s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadLatencyBucket111s(value uint64) (err error) {
	return instance.SetProperty("ReadLatencyBucket111s", (value))
}

// GetReadLatencyBucket111s gets the value of ReadLatencyBucket111s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadLatencyBucket111s() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadLatencyBucket111s")
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

// SetReadLatencyBucket122s sets the value of ReadLatencyBucket122s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadLatencyBucket122s(value uint64) (err error) {
	return instance.SetProperty("ReadLatencyBucket122s", (value))
}

// GetReadLatencyBucket122s gets the value of ReadLatencyBucket122s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadLatencyBucket122s() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadLatencyBucket122s")
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

// SetReadLatencyBucket1310s sets the value of ReadLatencyBucket1310s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadLatencyBucket1310s(value uint64) (err error) {
	return instance.SetProperty("ReadLatencyBucket1310s", (value))
}

// GetReadLatencyBucket1310s gets the value of ReadLatencyBucket1310s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadLatencyBucket1310s() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadLatencyBucket1310s")
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

// SetReadLatencyBucket1410s sets the value of ReadLatencyBucket1410s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadLatencyBucket1410s(value uint64) (err error) {
	return instance.SetProperty("ReadLatencyBucket1410s", (value))
}

// GetReadLatencyBucket1410s gets the value of ReadLatencyBucket1410s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadLatencyBucket1410s() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadLatencyBucket1410s")
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

// SetReadsPersec sets the value of ReadsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadsPersec(value uint64) (err error) {
	return instance.SetProperty("ReadsPersec", (value))
}

// GetReadsPersec gets the value of ReadsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadsPersec")
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

// SetReadsPersecBucket01128us sets the value of ReadsPersecBucket01128us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadsPersecBucket01128us(value uint64) (err error) {
	return instance.SetProperty("ReadsPersecBucket01128us", (value))
}

// GetReadsPersecBucket01128us gets the value of ReadsPersecBucket01128us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadsPersecBucket01128us() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadsPersecBucket01128us")
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

// SetReadsPersecBucket02256us sets the value of ReadsPersecBucket02256us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadsPersecBucket02256us(value uint64) (err error) {
	return instance.SetProperty("ReadsPersecBucket02256us", (value))
}

// GetReadsPersecBucket02256us gets the value of ReadsPersecBucket02256us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadsPersecBucket02256us() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadsPersecBucket02256us")
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

// SetReadsPersecBucket03512us sets the value of ReadsPersecBucket03512us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadsPersecBucket03512us(value uint64) (err error) {
	return instance.SetProperty("ReadsPersecBucket03512us", (value))
}

// GetReadsPersecBucket03512us gets the value of ReadsPersecBucket03512us for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadsPersecBucket03512us() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadsPersecBucket03512us")
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

// SetReadsPersecBucket041ms sets the value of ReadsPersecBucket041ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadsPersecBucket041ms(value uint64) (err error) {
	return instance.SetProperty("ReadsPersecBucket041ms", (value))
}

// GetReadsPersecBucket041ms gets the value of ReadsPersecBucket041ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadsPersecBucket041ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadsPersecBucket041ms")
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

// SetReadsPersecBucket054ms sets the value of ReadsPersecBucket054ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadsPersecBucket054ms(value uint64) (err error) {
	return instance.SetProperty("ReadsPersecBucket054ms", (value))
}

// GetReadsPersecBucket054ms gets the value of ReadsPersecBucket054ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadsPersecBucket054ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadsPersecBucket054ms")
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

// SetReadsPersecBucket0616ms sets the value of ReadsPersecBucket0616ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadsPersecBucket0616ms(value uint64) (err error) {
	return instance.SetProperty("ReadsPersecBucket0616ms", (value))
}

// GetReadsPersecBucket0616ms gets the value of ReadsPersecBucket0616ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadsPersecBucket0616ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadsPersecBucket0616ms")
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

// SetReadsPersecBucket0764ms sets the value of ReadsPersecBucket0764ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadsPersecBucket0764ms(value uint64) (err error) {
	return instance.SetProperty("ReadsPersecBucket0764ms", (value))
}

// GetReadsPersecBucket0764ms gets the value of ReadsPersecBucket0764ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadsPersecBucket0764ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadsPersecBucket0764ms")
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

// SetReadsPersecBucket08128ms sets the value of ReadsPersecBucket08128ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadsPersecBucket08128ms(value uint64) (err error) {
	return instance.SetProperty("ReadsPersecBucket08128ms", (value))
}

// GetReadsPersecBucket08128ms gets the value of ReadsPersecBucket08128ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadsPersecBucket08128ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadsPersecBucket08128ms")
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

// SetReadsPersecBucket09256ms sets the value of ReadsPersecBucket09256ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadsPersecBucket09256ms(value uint64) (err error) {
	return instance.SetProperty("ReadsPersecBucket09256ms", (value))
}

// GetReadsPersecBucket09256ms gets the value of ReadsPersecBucket09256ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadsPersecBucket09256ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadsPersecBucket09256ms")
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

// SetReadsPersecBucket10512ms sets the value of ReadsPersecBucket10512ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadsPersecBucket10512ms(value uint64) (err error) {
	return instance.SetProperty("ReadsPersecBucket10512ms", (value))
}

// GetReadsPersecBucket10512ms gets the value of ReadsPersecBucket10512ms for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadsPersecBucket10512ms() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadsPersecBucket10512ms")
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

// SetReadsPersecBucket111s sets the value of ReadsPersecBucket111s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadsPersecBucket111s(value uint64) (err error) {
	return instance.SetProperty("ReadsPersecBucket111s", (value))
}

// GetReadsPersecBucket111s gets the value of ReadsPersecBucket111s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadsPersecBucket111s() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadsPersecBucket111s")
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

// SetReadsPersecBucket122s sets the value of ReadsPersecBucket122s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadsPersecBucket122s(value uint64) (err error) {
	return instance.SetProperty("ReadsPersecBucket122s", (value))
}

// GetReadsPersecBucket122s gets the value of ReadsPersecBucket122s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadsPersecBucket122s() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadsPersecBucket122s")
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

// SetReadsPersecBucket1310s sets the value of ReadsPersecBucket1310s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadsPersecBucket1310s(value uint64) (err error) {
	return instance.SetProperty("ReadsPersecBucket1310s", (value))
}

// GetReadsPersecBucket1310s gets the value of ReadsPersecBucket1310s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadsPersecBucket1310s() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadsPersecBucket1310s")
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

// SetReadsPersecBucket1410s sets the value of ReadsPersecBucket1410s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertyReadsPersecBucket1410s(value uint64) (err error) {
	return instance.SetProperty("ReadsPersecBucket1410s", (value))
}

// GetReadsPersecBucket1410s gets the value of ReadsPersecBucket1410s for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertyReadsPersecBucket1410s() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadsPersecBucket1410s")
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

// SetSuccessfulReadsPersecBucket014K sets the value of SuccessfulReadsPersecBucket014K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertySuccessfulReadsPersecBucket014K(value uint64) (err error) {
	return instance.SetProperty("SuccessfulReadsPersecBucket014K", (value))
}

// GetSuccessfulReadsPersecBucket014K gets the value of SuccessfulReadsPersecBucket014K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertySuccessfulReadsPersecBucket014K() (value uint64, err error) {
	retValue, err := instance.GetProperty("SuccessfulReadsPersecBucket014K")
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

// SetSuccessfulReadsPersecBucket028K sets the value of SuccessfulReadsPersecBucket028K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertySuccessfulReadsPersecBucket028K(value uint64) (err error) {
	return instance.SetProperty("SuccessfulReadsPersecBucket028K", (value))
}

// GetSuccessfulReadsPersecBucket028K gets the value of SuccessfulReadsPersecBucket028K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertySuccessfulReadsPersecBucket028K() (value uint64, err error) {
	retValue, err := instance.GetProperty("SuccessfulReadsPersecBucket028K")
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

// SetSuccessfulReadsPersecBucket0316K sets the value of SuccessfulReadsPersecBucket0316K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertySuccessfulReadsPersecBucket0316K(value uint64) (err error) {
	return instance.SetProperty("SuccessfulReadsPersecBucket0316K", (value))
}

// GetSuccessfulReadsPersecBucket0316K gets the value of SuccessfulReadsPersecBucket0316K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertySuccessfulReadsPersecBucket0316K() (value uint64, err error) {
	retValue, err := instance.GetProperty("SuccessfulReadsPersecBucket0316K")
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

// SetSuccessfulReadsPersecBucket0432K sets the value of SuccessfulReadsPersecBucket0432K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertySuccessfulReadsPersecBucket0432K(value uint64) (err error) {
	return instance.SetProperty("SuccessfulReadsPersecBucket0432K", (value))
}

// GetSuccessfulReadsPersecBucket0432K gets the value of SuccessfulReadsPersecBucket0432K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertySuccessfulReadsPersecBucket0432K() (value uint64, err error) {
	retValue, err := instance.GetProperty("SuccessfulReadsPersecBucket0432K")
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

// SetSuccessfulReadsPersecBucket0564K sets the value of SuccessfulReadsPersecBucket0564K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertySuccessfulReadsPersecBucket0564K(value uint64) (err error) {
	return instance.SetProperty("SuccessfulReadsPersecBucket0564K", (value))
}

// GetSuccessfulReadsPersecBucket0564K gets the value of SuccessfulReadsPersecBucket0564K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertySuccessfulReadsPersecBucket0564K() (value uint64, err error) {
	retValue, err := instance.GetProperty("SuccessfulReadsPersecBucket0564K")
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

// SetSuccessfulReadsPersecBucket06128K sets the value of SuccessfulReadsPersecBucket06128K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertySuccessfulReadsPersecBucket06128K(value uint64) (err error) {
	return instance.SetProperty("SuccessfulReadsPersecBucket06128K", (value))
}

// GetSuccessfulReadsPersecBucket06128K gets the value of SuccessfulReadsPersecBucket06128K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertySuccessfulReadsPersecBucket06128K() (value uint64, err error) {
	retValue, err := instance.GetProperty("SuccessfulReadsPersecBucket06128K")
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

// SetSuccessfulReadsPersecBucket07256K sets the value of SuccessfulReadsPersecBucket07256K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertySuccessfulReadsPersecBucket07256K(value uint64) (err error) {
	return instance.SetProperty("SuccessfulReadsPersecBucket07256K", (value))
}

// GetSuccessfulReadsPersecBucket07256K gets the value of SuccessfulReadsPersecBucket07256K for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertySuccessfulReadsPersecBucket07256K() (value uint64, err error) {
	retValue, err := instance.GetProperty("SuccessfulReadsPersecBucket07256K")
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

// SetSuccessfulReadsPersecBucket081M sets the value of SuccessfulReadsPersecBucket081M for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertySuccessfulReadsPersecBucket081M(value uint64) (err error) {
	return instance.SetProperty("SuccessfulReadsPersecBucket081M", (value))
}

// GetSuccessfulReadsPersecBucket081M gets the value of SuccessfulReadsPersecBucket081M for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertySuccessfulReadsPersecBucket081M() (value uint64, err error) {
	retValue, err := instance.GetProperty("SuccessfulReadsPersecBucket081M")
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

// SetSuccessfulReadsPersecBucket091M sets the value of SuccessfulReadsPersecBucket091M for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) SetPropertySuccessfulReadsPersecBucket091M(value uint64) (err error) {
	return instance.SetProperty("SuccessfulReadsPersecBucket091M", (value))
}

// GetSuccessfulReadsPersecBucket091M gets the value of SuccessfulReadsPersecBucket091M for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitReads) GetPropertySuccessfulReadsPersecBucket091M() (value uint64, err error) {
	retValue, err := instance.GetProperty("SuccessfulReadsPersecBucket091M")
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
