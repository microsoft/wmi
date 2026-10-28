// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// FileOperation struct
type FileOperation struct {
	*FileTrace

	//
	AccessToken uint32

	//
	CreateOnExisting FileOperation_CreateOnExisting

	//
	FileName string

	//
	FileObject uint64

	//
	IsDirectory FileOperation_IsDirectory

	//
	IsFastIO uint8

	//
	IsPagingIO uint8

	//
	LastAccessTime int64

	//
	MinorOperation uint8

	//
	Operation FileOperation_Operation

	//
	OperationalParameters []uint8

	//
	ParametersLength uint32

	//
	PreviousValue []uint8

	//
	PreviousValueLength uint32

	//
	ProcessCreateTime int64

	//
	ProcessId uint32

	//
	ResultData []uint8

	//
	ResultLength uint32

	//
	SequenceNumber uint32

	//
	SessionId uint32

	//
	SidLength uint32

	//
	StartTime int64

	//
	Status uint32

	//
	UserSID interface{}

	//
	VolumeDosName string

	//
	VolumeGuidName string

	//
	VolumeName string

	//
	WindowStation uint64
}

func NewFileOperationEx1(instance *cim.WmiInstance) (newInstance *FileOperation, err error) {
	tmp, err := NewFileTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FileOperation{
		FileTrace: tmp,
	}
	return
}

func NewFileOperationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FileOperation, err error) {
	tmp, err := NewFileTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FileOperation{
		FileTrace: tmp,
	}
	return
}

// SetAccessToken sets the value of AccessToken for the instance
func (instance *FileOperation) SetPropertyAccessToken(value uint32) (err error) {
	return instance.SetProperty("AccessToken", (value))
}

// GetAccessToken gets the value of AccessToken for the instance
func (instance *FileOperation) GetPropertyAccessToken() (value uint32, err error) {
	retValue, err := instance.GetProperty("AccessToken")
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

// SetCreateOnExisting sets the value of CreateOnExisting for the instance
func (instance *FileOperation) SetPropertyCreateOnExisting(value FileOperation_CreateOnExisting) (err error) {
	return instance.SetProperty("CreateOnExisting", (value))
}

// GetCreateOnExisting gets the value of CreateOnExisting for the instance
func (instance *FileOperation) GetPropertyCreateOnExisting() (value FileOperation_CreateOnExisting, err error) {
	retValue, err := instance.GetProperty("CreateOnExisting")
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

	value = FileOperation_CreateOnExisting(valuetmp)

	return
}

// SetFileName sets the value of FileName for the instance
func (instance *FileOperation) SetPropertyFileName(value string) (err error) {
	return instance.SetProperty("FileName", (value))
}

// GetFileName gets the value of FileName for the instance
func (instance *FileOperation) GetPropertyFileName() (value string, err error) {
	retValue, err := instance.GetProperty("FileName")
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

// SetFileObject sets the value of FileObject for the instance
func (instance *FileOperation) SetPropertyFileObject(value uint64) (err error) {
	return instance.SetProperty("FileObject", (value))
}

// GetFileObject gets the value of FileObject for the instance
func (instance *FileOperation) GetPropertyFileObject() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileObject")
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

// SetIsDirectory sets the value of IsDirectory for the instance
func (instance *FileOperation) SetPropertyIsDirectory(value FileOperation_IsDirectory) (err error) {
	return instance.SetProperty("IsDirectory", (value))
}

// GetIsDirectory gets the value of IsDirectory for the instance
func (instance *FileOperation) GetPropertyIsDirectory() (value FileOperation_IsDirectory, err error) {
	retValue, err := instance.GetProperty("IsDirectory")
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

	value = FileOperation_IsDirectory(valuetmp)

	return
}

// SetIsFastIO sets the value of IsFastIO for the instance
func (instance *FileOperation) SetPropertyIsFastIO(value uint8) (err error) {
	return instance.SetProperty("IsFastIO", (value))
}

// GetIsFastIO gets the value of IsFastIO for the instance
func (instance *FileOperation) GetPropertyIsFastIO() (value uint8, err error) {
	retValue, err := instance.GetProperty("IsFastIO")
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

// SetIsPagingIO sets the value of IsPagingIO for the instance
func (instance *FileOperation) SetPropertyIsPagingIO(value uint8) (err error) {
	return instance.SetProperty("IsPagingIO", (value))
}

// GetIsPagingIO gets the value of IsPagingIO for the instance
func (instance *FileOperation) GetPropertyIsPagingIO() (value uint8, err error) {
	retValue, err := instance.GetProperty("IsPagingIO")
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

// SetLastAccessTime sets the value of LastAccessTime for the instance
func (instance *FileOperation) SetPropertyLastAccessTime(value int64) (err error) {
	return instance.SetProperty("LastAccessTime", (value))
}

// GetLastAccessTime gets the value of LastAccessTime for the instance
func (instance *FileOperation) GetPropertyLastAccessTime() (value int64, err error) {
	retValue, err := instance.GetProperty("LastAccessTime")
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

// SetMinorOperation sets the value of MinorOperation for the instance
func (instance *FileOperation) SetPropertyMinorOperation(value uint8) (err error) {
	return instance.SetProperty("MinorOperation", (value))
}

// GetMinorOperation gets the value of MinorOperation for the instance
func (instance *FileOperation) GetPropertyMinorOperation() (value uint8, err error) {
	retValue, err := instance.GetProperty("MinorOperation")
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

// SetOperation sets the value of Operation for the instance
func (instance *FileOperation) SetPropertyOperation(value FileOperation_Operation) (err error) {
	return instance.SetProperty("Operation", (value))
}

// GetOperation gets the value of Operation for the instance
func (instance *FileOperation) GetPropertyOperation() (value FileOperation_Operation, err error) {
	retValue, err := instance.GetProperty("Operation")
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

	value = FileOperation_Operation(valuetmp)

	return
}

// SetOperationalParameters sets the value of OperationalParameters for the instance
func (instance *FileOperation) SetPropertyOperationalParameters(value []uint8) (err error) {
	return instance.SetProperty("OperationalParameters", (value))
}

// GetOperationalParameters gets the value of OperationalParameters for the instance
func (instance *FileOperation) GetPropertyOperationalParameters() (value []uint8, err error) {
	retValue, err := instance.GetProperty("OperationalParameters")
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

// SetParametersLength sets the value of ParametersLength for the instance
func (instance *FileOperation) SetPropertyParametersLength(value uint32) (err error) {
	return instance.SetProperty("ParametersLength", (value))
}

// GetParametersLength gets the value of ParametersLength for the instance
func (instance *FileOperation) GetPropertyParametersLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("ParametersLength")
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

// SetPreviousValue sets the value of PreviousValue for the instance
func (instance *FileOperation) SetPropertyPreviousValue(value []uint8) (err error) {
	return instance.SetProperty("PreviousValue", (value))
}

// GetPreviousValue gets the value of PreviousValue for the instance
func (instance *FileOperation) GetPropertyPreviousValue() (value []uint8, err error) {
	retValue, err := instance.GetProperty("PreviousValue")
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

// SetPreviousValueLength sets the value of PreviousValueLength for the instance
func (instance *FileOperation) SetPropertyPreviousValueLength(value uint32) (err error) {
	return instance.SetProperty("PreviousValueLength", (value))
}

// GetPreviousValueLength gets the value of PreviousValueLength for the instance
func (instance *FileOperation) GetPropertyPreviousValueLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("PreviousValueLength")
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

// SetProcessCreateTime sets the value of ProcessCreateTime for the instance
func (instance *FileOperation) SetPropertyProcessCreateTime(value int64) (err error) {
	return instance.SetProperty("ProcessCreateTime", (value))
}

// GetProcessCreateTime gets the value of ProcessCreateTime for the instance
func (instance *FileOperation) GetPropertyProcessCreateTime() (value int64, err error) {
	retValue, err := instance.GetProperty("ProcessCreateTime")
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *FileOperation) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *FileOperation) GetPropertyProcessId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessId")
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

// SetResultData sets the value of ResultData for the instance
func (instance *FileOperation) SetPropertyResultData(value []uint8) (err error) {
	return instance.SetProperty("ResultData", (value))
}

// GetResultData gets the value of ResultData for the instance
func (instance *FileOperation) GetPropertyResultData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("ResultData")
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

// SetResultLength sets the value of ResultLength for the instance
func (instance *FileOperation) SetPropertyResultLength(value uint32) (err error) {
	return instance.SetProperty("ResultLength", (value))
}

// GetResultLength gets the value of ResultLength for the instance
func (instance *FileOperation) GetPropertyResultLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResultLength")
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

// SetSequenceNumber sets the value of SequenceNumber for the instance
func (instance *FileOperation) SetPropertySequenceNumber(value uint32) (err error) {
	return instance.SetProperty("SequenceNumber", (value))
}

// GetSequenceNumber gets the value of SequenceNumber for the instance
func (instance *FileOperation) GetPropertySequenceNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("SequenceNumber")
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

// SetSessionId sets the value of SessionId for the instance
func (instance *FileOperation) SetPropertySessionId(value uint32) (err error) {
	return instance.SetProperty("SessionId", (value))
}

// GetSessionId gets the value of SessionId for the instance
func (instance *FileOperation) GetPropertySessionId() (value uint32, err error) {
	retValue, err := instance.GetProperty("SessionId")
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

// SetSidLength sets the value of SidLength for the instance
func (instance *FileOperation) SetPropertySidLength(value uint32) (err error) {
	return instance.SetProperty("SidLength", (value))
}

// GetSidLength gets the value of SidLength for the instance
func (instance *FileOperation) GetPropertySidLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("SidLength")
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

// SetStartTime sets the value of StartTime for the instance
func (instance *FileOperation) SetPropertyStartTime(value int64) (err error) {
	return instance.SetProperty("StartTime", (value))
}

// GetStartTime gets the value of StartTime for the instance
func (instance *FileOperation) GetPropertyStartTime() (value int64, err error) {
	retValue, err := instance.GetProperty("StartTime")
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

// SetStatus sets the value of Status for the instance
func (instance *FileOperation) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *FileOperation) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
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

// SetUserSID sets the value of UserSID for the instance
func (instance *FileOperation) SetPropertyUserSID(value interface{}) (err error) {
	return instance.SetProperty("UserSID", (value))
}

// GetUserSID gets the value of UserSID for the instance
func (instance *FileOperation) GetPropertyUserSID() (value interface{}, err error) {
	retValue, err := instance.GetProperty("UserSID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetVolumeDosName sets the value of VolumeDosName for the instance
func (instance *FileOperation) SetPropertyVolumeDosName(value string) (err error) {
	return instance.SetProperty("VolumeDosName", (value))
}

// GetVolumeDosName gets the value of VolumeDosName for the instance
func (instance *FileOperation) GetPropertyVolumeDosName() (value string, err error) {
	retValue, err := instance.GetProperty("VolumeDosName")
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

// SetVolumeGuidName sets the value of VolumeGuidName for the instance
func (instance *FileOperation) SetPropertyVolumeGuidName(value string) (err error) {
	return instance.SetProperty("VolumeGuidName", (value))
}

// GetVolumeGuidName gets the value of VolumeGuidName for the instance
func (instance *FileOperation) GetPropertyVolumeGuidName() (value string, err error) {
	retValue, err := instance.GetProperty("VolumeGuidName")
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

// SetVolumeName sets the value of VolumeName for the instance
func (instance *FileOperation) SetPropertyVolumeName(value string) (err error) {
	return instance.SetProperty("VolumeName", (value))
}

// GetVolumeName gets the value of VolumeName for the instance
func (instance *FileOperation) GetPropertyVolumeName() (value string, err error) {
	retValue, err := instance.GetProperty("VolumeName")
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

// SetWindowStation sets the value of WindowStation for the instance
func (instance *FileOperation) SetPropertyWindowStation(value uint64) (err error) {
	return instance.SetProperty("WindowStation", (value))
}

// GetWindowStation gets the value of WindowStation for the instance
func (instance *FileOperation) GetPropertyWindowStation() (value uint64, err error) {
	retValue, err := instance.GetProperty("WindowStation")
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
