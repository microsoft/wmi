// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_PartitionableGpu struct
type Msvm_PartitionableGpu struct {
	*CIM_ComputerSystem

	//
	AvailableCompute uint64

	//
	AvailableDecode uint64

	//
	AvailableEncode uint64

	//
	AvailableVRAM uint64

	//
	MaxPartitionCompute uint64

	//
	MaxPartitionDecode uint64

	//
	MaxPartitionEncode uint64

	//
	MaxPartitionVRAM uint64

	//
	MinPartitionCompute uint64

	//
	MinPartitionDecode uint64

	//
	MinPartitionEncode uint64

	//
	MinPartitionVRAM uint64

	//
	OptimalPartitionCompute uint64

	//
	OptimalPartitionDecode uint64

	//
	OptimalPartitionEncode uint64

	//
	OptimalPartitionVRAM uint64

	//
	PartitionCount uint16

	//
	TotalCompute uint64

	//
	TotalDecode uint64

	//
	TotalEncode uint64

	//
	TotalVRAM uint64

	//
	ValidPartitionCounts []uint16
}

func NewMsvm_PartitionableGpuEx1(instance *cim.WmiInstance) (newInstance *Msvm_PartitionableGpu, err error) {
	tmp, err := NewCIM_ComputerSystemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_PartitionableGpu{
		CIM_ComputerSystem: tmp,
	}
	return
}

func NewMsvm_PartitionableGpuEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_PartitionableGpu, err error) {
	tmp, err := NewCIM_ComputerSystemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_PartitionableGpu{
		CIM_ComputerSystem: tmp,
	}
	return
}

// SetAvailableCompute sets the value of AvailableCompute for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyAvailableCompute(value uint64) (err error) {
	return instance.SetProperty("AvailableCompute", (value))
}

// GetAvailableCompute gets the value of AvailableCompute for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyAvailableCompute() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvailableCompute")
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

// SetAvailableDecode sets the value of AvailableDecode for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyAvailableDecode(value uint64) (err error) {
	return instance.SetProperty("AvailableDecode", (value))
}

// GetAvailableDecode gets the value of AvailableDecode for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyAvailableDecode() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvailableDecode")
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

// SetAvailableEncode sets the value of AvailableEncode for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyAvailableEncode(value uint64) (err error) {
	return instance.SetProperty("AvailableEncode", (value))
}

// GetAvailableEncode gets the value of AvailableEncode for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyAvailableEncode() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvailableEncode")
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

// SetAvailableVRAM sets the value of AvailableVRAM for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyAvailableVRAM(value uint64) (err error) {
	return instance.SetProperty("AvailableVRAM", (value))
}

// GetAvailableVRAM gets the value of AvailableVRAM for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyAvailableVRAM() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvailableVRAM")
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

// SetMaxPartitionCompute sets the value of MaxPartitionCompute for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyMaxPartitionCompute(value uint64) (err error) {
	return instance.SetProperty("MaxPartitionCompute", (value))
}

// GetMaxPartitionCompute gets the value of MaxPartitionCompute for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyMaxPartitionCompute() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxPartitionCompute")
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

// SetMaxPartitionDecode sets the value of MaxPartitionDecode for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyMaxPartitionDecode(value uint64) (err error) {
	return instance.SetProperty("MaxPartitionDecode", (value))
}

// GetMaxPartitionDecode gets the value of MaxPartitionDecode for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyMaxPartitionDecode() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxPartitionDecode")
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

// SetMaxPartitionEncode sets the value of MaxPartitionEncode for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyMaxPartitionEncode(value uint64) (err error) {
	return instance.SetProperty("MaxPartitionEncode", (value))
}

// GetMaxPartitionEncode gets the value of MaxPartitionEncode for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyMaxPartitionEncode() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxPartitionEncode")
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

// SetMaxPartitionVRAM sets the value of MaxPartitionVRAM for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyMaxPartitionVRAM(value uint64) (err error) {
	return instance.SetProperty("MaxPartitionVRAM", (value))
}

// GetMaxPartitionVRAM gets the value of MaxPartitionVRAM for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyMaxPartitionVRAM() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxPartitionVRAM")
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

// SetMinPartitionCompute sets the value of MinPartitionCompute for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyMinPartitionCompute(value uint64) (err error) {
	return instance.SetProperty("MinPartitionCompute", (value))
}

// GetMinPartitionCompute gets the value of MinPartitionCompute for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyMinPartitionCompute() (value uint64, err error) {
	retValue, err := instance.GetProperty("MinPartitionCompute")
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

// SetMinPartitionDecode sets the value of MinPartitionDecode for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyMinPartitionDecode(value uint64) (err error) {
	return instance.SetProperty("MinPartitionDecode", (value))
}

// GetMinPartitionDecode gets the value of MinPartitionDecode for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyMinPartitionDecode() (value uint64, err error) {
	retValue, err := instance.GetProperty("MinPartitionDecode")
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

// SetMinPartitionEncode sets the value of MinPartitionEncode for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyMinPartitionEncode(value uint64) (err error) {
	return instance.SetProperty("MinPartitionEncode", (value))
}

// GetMinPartitionEncode gets the value of MinPartitionEncode for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyMinPartitionEncode() (value uint64, err error) {
	retValue, err := instance.GetProperty("MinPartitionEncode")
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

// SetMinPartitionVRAM sets the value of MinPartitionVRAM for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyMinPartitionVRAM(value uint64) (err error) {
	return instance.SetProperty("MinPartitionVRAM", (value))
}

// GetMinPartitionVRAM gets the value of MinPartitionVRAM for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyMinPartitionVRAM() (value uint64, err error) {
	retValue, err := instance.GetProperty("MinPartitionVRAM")
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

// SetOptimalPartitionCompute sets the value of OptimalPartitionCompute for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyOptimalPartitionCompute(value uint64) (err error) {
	return instance.SetProperty("OptimalPartitionCompute", (value))
}

// GetOptimalPartitionCompute gets the value of OptimalPartitionCompute for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyOptimalPartitionCompute() (value uint64, err error) {
	retValue, err := instance.GetProperty("OptimalPartitionCompute")
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

// SetOptimalPartitionDecode sets the value of OptimalPartitionDecode for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyOptimalPartitionDecode(value uint64) (err error) {
	return instance.SetProperty("OptimalPartitionDecode", (value))
}

// GetOptimalPartitionDecode gets the value of OptimalPartitionDecode for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyOptimalPartitionDecode() (value uint64, err error) {
	retValue, err := instance.GetProperty("OptimalPartitionDecode")
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

// SetOptimalPartitionEncode sets the value of OptimalPartitionEncode for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyOptimalPartitionEncode(value uint64) (err error) {
	return instance.SetProperty("OptimalPartitionEncode", (value))
}

// GetOptimalPartitionEncode gets the value of OptimalPartitionEncode for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyOptimalPartitionEncode() (value uint64, err error) {
	retValue, err := instance.GetProperty("OptimalPartitionEncode")
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

// SetOptimalPartitionVRAM sets the value of OptimalPartitionVRAM for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyOptimalPartitionVRAM(value uint64) (err error) {
	return instance.SetProperty("OptimalPartitionVRAM", (value))
}

// GetOptimalPartitionVRAM gets the value of OptimalPartitionVRAM for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyOptimalPartitionVRAM() (value uint64, err error) {
	retValue, err := instance.GetProperty("OptimalPartitionVRAM")
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

// SetPartitionCount sets the value of PartitionCount for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyPartitionCount(value uint16) (err error) {
	return instance.SetProperty("PartitionCount", (value))
}

// GetPartitionCount gets the value of PartitionCount for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyPartitionCount() (value uint16, err error) {
	retValue, err := instance.GetProperty("PartitionCount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetTotalCompute sets the value of TotalCompute for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyTotalCompute(value uint64) (err error) {
	return instance.SetProperty("TotalCompute", (value))
}

// GetTotalCompute gets the value of TotalCompute for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyTotalCompute() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalCompute")
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

// SetTotalDecode sets the value of TotalDecode for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyTotalDecode(value uint64) (err error) {
	return instance.SetProperty("TotalDecode", (value))
}

// GetTotalDecode gets the value of TotalDecode for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyTotalDecode() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalDecode")
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

// SetTotalEncode sets the value of TotalEncode for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyTotalEncode(value uint64) (err error) {
	return instance.SetProperty("TotalEncode", (value))
}

// GetTotalEncode gets the value of TotalEncode for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyTotalEncode() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalEncode")
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

// SetTotalVRAM sets the value of TotalVRAM for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyTotalVRAM(value uint64) (err error) {
	return instance.SetProperty("TotalVRAM", (value))
}

// GetTotalVRAM gets the value of TotalVRAM for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyTotalVRAM() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalVRAM")
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

// SetValidPartitionCounts sets the value of ValidPartitionCounts for the instance
func (instance *Msvm_PartitionableGpu) SetPropertyValidPartitionCounts(value []uint16) (err error) {
	return instance.SetProperty("ValidPartitionCounts", (value))
}

// GetValidPartitionCounts gets the value of ValidPartitionCounts for the instance
func (instance *Msvm_PartitionableGpu) GetPropertyValidPartitionCounts() (value []uint16, err error) {
	retValue, err := instance.GetProperty("ValidPartitionCounts")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}
