// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_Memory struct
type CIM_Memory struct {
	*CIM_StorageExtent

	// An array of octets holding additional error information. An example is ECC Syndrome or the return of the check bits if a CRC-based ErrorMethodology is used. In the latter case, if a single bit error is recognized and the CRC algorithm is known, it is possible to determine the exact bit that failed. This type of data (ECC Syndrome, Check Bit or Parity Bit data, or other vendor supplied information) is included in this field. If the ErrorInfo property is equal to 3, "OK", then AdditionalErrorData has no meaning.
	AdditionalErrorData []uint8

	// Boolean indicating that the most recent error was correctable. If the ErrorInfo property is equal to 3, "OK", then this property has no meaning.
	CorrectableError bool

	// The ending address, referenced by an application or operating system and mapped by a memory controller, for this Memory object. The ending address is specified in KBytes.
	EndingAddress uint64

	// An integer enumeration indicating the memory access operation that caused the last error. The type of error is described by the ErrorInfo property. If the ErrorInfo property is equal to 3, "OK", then this property has no meaning.
	ErrorAccess Memory_ErrorAccess

	// Specifies the address of the last memory error. The type of error is described by the ErrorInfo property. If the ErrorInfo property is equal to 3, "OK", then this property has no meaning.
	ErrorAddress uint64

	// Data captured during the last erroneous mebmory access. The data occupies the first n octets of the array necessary to hold the number of bits specified by the ErrorTransferSize property. If ErrorTransferSize is 0, then this property has no meaning.
	ErrorData []uint8

	// The ordering for data stored in the ErrorData property. "Least Significant Byte First" (value=1) or "Most Significant Byte First" (2) can be specified. If ErrorTransferSize is 0, then this property has no meaning.
	ErrorDataOrder Memory_ErrorDataOrder

	// An integer enumeration describing the type of error that occurred most recently. For example, single (value=6) or double bit errors (7) can be specified using this property. The values, 12-14, are undefined in the CIM Schema since in DMI, they mix the semantics of the type of error and whether it was correctable or not. The latter is indicated in the property, CorrectableError.
	ErrorInfo Memory_ErrorInfo

	// Specifies the range, in bytes, to which the last error can be resolved. For example, if error addresses are resolved to bit 11 (ie, on a typical page basis), then errors can be resolved to 4K boundaries and this property is set to 4000. If the ErrorInfo property is equal to 3, "OK", then this property has no meaning.
	ErrorResolution uint64

	// The time that the last memory error occurred. The type of error is described by the ErrorInfo property. If the Error Info property is equal to 3, "OK", then this property has no meaning.
	ErrorTime string

	// The size of the data transfer in bits that caused the last error. 0 indicates no error. If the ErrorInfo property is equal to 3, "OK", then this property should be set to 0.
	ErrorTransferSize uint32

	// Free form string providing more information if the Error Type property is set to 1, "Other". If not set to 1, this string has no meaning.
	OtherErrorDescription string

	// The beginning address, referenced by an application or operating system and mapped by a memory controller, for this Memory object. The starting address is specified in KBytes.
	StartingAddress uint64

	// Boolean indicating whether the address information in the property, ErrorAddress, is a system-level address (TRUE) or a physical address (FALSE). If the ErrorInfo property is equal to 3, "OK", then this property has no meaning.
	SystemLevelAddress bool

	// Volatile is a property that indicates whether this memory is volatile or not.
	Volatile bool
}

func NewCIM_MemoryEx1(instance *cim.WmiInstance) (newInstance *CIM_Memory, err error) {
	tmp, err := NewCIM_StorageExtentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_Memory{
		CIM_StorageExtent: tmp,
	}
	return
}

func NewCIM_MemoryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Memory, err error) {
	tmp, err := NewCIM_StorageExtentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Memory{
		CIM_StorageExtent: tmp,
	}
	return
}

// SetAdditionalErrorData sets the value of AdditionalErrorData for the instance
func (instance *CIM_Memory) SetPropertyAdditionalErrorData(value []uint8) (err error) {
	return instance.SetProperty("AdditionalErrorData", (value))
}

// GetAdditionalErrorData gets the value of AdditionalErrorData for the instance
func (instance *CIM_Memory) GetPropertyAdditionalErrorData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("AdditionalErrorData")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetCorrectableError sets the value of CorrectableError for the instance
func (instance *CIM_Memory) SetPropertyCorrectableError(value bool) (err error) {
	return instance.SetProperty("CorrectableError", (value))
}

// GetCorrectableError gets the value of CorrectableError for the instance
func (instance *CIM_Memory) GetPropertyCorrectableError() (value bool, err error) {
	retValue, err := instance.GetProperty("CorrectableError")
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

// SetEndingAddress sets the value of EndingAddress for the instance
func (instance *CIM_Memory) SetPropertyEndingAddress(value uint64) (err error) {
	return instance.SetProperty("EndingAddress", (value))
}

// GetEndingAddress gets the value of EndingAddress for the instance
func (instance *CIM_Memory) GetPropertyEndingAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("EndingAddress")
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

// SetErrorAccess sets the value of ErrorAccess for the instance
func (instance *CIM_Memory) SetPropertyErrorAccess(value Memory_ErrorAccess) (err error) {
	return instance.SetProperty("ErrorAccess", (value))
}

// GetErrorAccess gets the value of ErrorAccess for the instance
func (instance *CIM_Memory) GetPropertyErrorAccess() (value Memory_ErrorAccess, err error) {
	retValue, err := instance.GetProperty("ErrorAccess")
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

	value = Memory_ErrorAccess(valuetmp)

	return
}

// SetErrorAddress sets the value of ErrorAddress for the instance
func (instance *CIM_Memory) SetPropertyErrorAddress(value uint64) (err error) {
	return instance.SetProperty("ErrorAddress", (value))
}

// GetErrorAddress gets the value of ErrorAddress for the instance
func (instance *CIM_Memory) GetPropertyErrorAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("ErrorAddress")
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

// SetErrorData sets the value of ErrorData for the instance
func (instance *CIM_Memory) SetPropertyErrorData(value []uint8) (err error) {
	return instance.SetProperty("ErrorData", (value))
}

// GetErrorData gets the value of ErrorData for the instance
func (instance *CIM_Memory) GetPropertyErrorData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("ErrorData")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetErrorDataOrder sets the value of ErrorDataOrder for the instance
func (instance *CIM_Memory) SetPropertyErrorDataOrder(value Memory_ErrorDataOrder) (err error) {
	return instance.SetProperty("ErrorDataOrder", (value))
}

// GetErrorDataOrder gets the value of ErrorDataOrder for the instance
func (instance *CIM_Memory) GetPropertyErrorDataOrder() (value Memory_ErrorDataOrder, err error) {
	retValue, err := instance.GetProperty("ErrorDataOrder")
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

	value = Memory_ErrorDataOrder(valuetmp)

	return
}

// SetErrorInfo sets the value of ErrorInfo for the instance
func (instance *CIM_Memory) SetPropertyErrorInfo(value Memory_ErrorInfo) (err error) {
	return instance.SetProperty("ErrorInfo", (value))
}

// GetErrorInfo gets the value of ErrorInfo for the instance
func (instance *CIM_Memory) GetPropertyErrorInfo() (value Memory_ErrorInfo, err error) {
	retValue, err := instance.GetProperty("ErrorInfo")
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

	value = Memory_ErrorInfo(valuetmp)

	return
}

// SetErrorResolution sets the value of ErrorResolution for the instance
func (instance *CIM_Memory) SetPropertyErrorResolution(value uint64) (err error) {
	return instance.SetProperty("ErrorResolution", (value))
}

// GetErrorResolution gets the value of ErrorResolution for the instance
func (instance *CIM_Memory) GetPropertyErrorResolution() (value uint64, err error) {
	retValue, err := instance.GetProperty("ErrorResolution")
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

// SetErrorTime sets the value of ErrorTime for the instance
func (instance *CIM_Memory) SetPropertyErrorTime(value string) (err error) {
	return instance.SetProperty("ErrorTime", (value))
}

// GetErrorTime gets the value of ErrorTime for the instance
func (instance *CIM_Memory) GetPropertyErrorTime() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorTime")
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

// SetErrorTransferSize sets the value of ErrorTransferSize for the instance
func (instance *CIM_Memory) SetPropertyErrorTransferSize(value uint32) (err error) {
	return instance.SetProperty("ErrorTransferSize", (value))
}

// GetErrorTransferSize gets the value of ErrorTransferSize for the instance
func (instance *CIM_Memory) GetPropertyErrorTransferSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("ErrorTransferSize")
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

// SetOtherErrorDescription sets the value of OtherErrorDescription for the instance
func (instance *CIM_Memory) SetPropertyOtherErrorDescription(value string) (err error) {
	return instance.SetProperty("OtherErrorDescription", (value))
}

// GetOtherErrorDescription gets the value of OtherErrorDescription for the instance
func (instance *CIM_Memory) GetPropertyOtherErrorDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherErrorDescription")
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

// SetStartingAddress sets the value of StartingAddress for the instance
func (instance *CIM_Memory) SetPropertyStartingAddress(value uint64) (err error) {
	return instance.SetProperty("StartingAddress", (value))
}

// GetStartingAddress gets the value of StartingAddress for the instance
func (instance *CIM_Memory) GetPropertyStartingAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("StartingAddress")
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

// SetSystemLevelAddress sets the value of SystemLevelAddress for the instance
func (instance *CIM_Memory) SetPropertySystemLevelAddress(value bool) (err error) {
	return instance.SetProperty("SystemLevelAddress", (value))
}

// GetSystemLevelAddress gets the value of SystemLevelAddress for the instance
func (instance *CIM_Memory) GetPropertySystemLevelAddress() (value bool, err error) {
	retValue, err := instance.GetProperty("SystemLevelAddress")
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

// Setvolatile sets the value of volatile for the instance
func (instance *CIM_Memory) SetPropertyvolatile(value bool) (err error) {
	return instance.SetProperty("volatile", (value))
}

// Getvolatile gets the value of volatile for the instance
func (instance *CIM_Memory) GetPropertyvolatile() (value bool, err error) {
	retValue, err := instance.GetProperty("volatile")
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
