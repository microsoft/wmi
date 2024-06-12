// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSTapeProblemDeviceError struct
type MSTapeProblemDeviceError struct {
	*MSTapeDriver

	//
	Active bool

	//
	DriveHardwareError bool

	//
	DriveRequiresCleaning bool

	//
	HardError bool

	//
	InstanceName string

	//
	MediaLife bool

	//
	ReadFailure bool

	//
	ReadWarning bool

	//
	ScsiInterfaceError bool

	//
	TapeSnapped bool

	//
	TimetoCleanDrive bool

	//
	UnsupportedFormat bool

	//
	WriteFailure bool

	//
	WriteWarning bool
}

func NewMSTapeProblemDeviceErrorEx1(instance *cim.WmiInstance) (newInstance *MSTapeProblemDeviceError, err error) {
	tmp, err := NewMSTapeDriverEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSTapeProblemDeviceError{
		MSTapeDriver: tmp,
	}
	return
}

func NewMSTapeProblemDeviceErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSTapeProblemDeviceError, err error) {
	tmp, err := NewMSTapeDriverEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSTapeProblemDeviceError{
		MSTapeDriver: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSTapeProblemDeviceError) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSTapeProblemDeviceError) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetDriveHardwareError sets the value of DriveHardwareError for the instance
func (instance *MSTapeProblemDeviceError) SetPropertyDriveHardwareError(value bool) (err error) {
	return instance.SetProperty("DriveHardwareError", (value))
}

// GetDriveHardwareError gets the value of DriveHardwareError for the instance
func (instance *MSTapeProblemDeviceError) GetPropertyDriveHardwareError() (value bool, err error) {
	retValue, err := instance.GetProperty("DriveHardwareError")
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

// SetDriveRequiresCleaning sets the value of DriveRequiresCleaning for the instance
func (instance *MSTapeProblemDeviceError) SetPropertyDriveRequiresCleaning(value bool) (err error) {
	return instance.SetProperty("DriveRequiresCleaning", (value))
}

// GetDriveRequiresCleaning gets the value of DriveRequiresCleaning for the instance
func (instance *MSTapeProblemDeviceError) GetPropertyDriveRequiresCleaning() (value bool, err error) {
	retValue, err := instance.GetProperty("DriveRequiresCleaning")
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

// SetHardError sets the value of HardError for the instance
func (instance *MSTapeProblemDeviceError) SetPropertyHardError(value bool) (err error) {
	return instance.SetProperty("HardError", (value))
}

// GetHardError gets the value of HardError for the instance
func (instance *MSTapeProblemDeviceError) GetPropertyHardError() (value bool, err error) {
	retValue, err := instance.GetProperty("HardError")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSTapeProblemDeviceError) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSTapeProblemDeviceError) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetMediaLife sets the value of MediaLife for the instance
func (instance *MSTapeProblemDeviceError) SetPropertyMediaLife(value bool) (err error) {
	return instance.SetProperty("MediaLife", (value))
}

// GetMediaLife gets the value of MediaLife for the instance
func (instance *MSTapeProblemDeviceError) GetPropertyMediaLife() (value bool, err error) {
	retValue, err := instance.GetProperty("MediaLife")
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

// SetReadFailure sets the value of ReadFailure for the instance
func (instance *MSTapeProblemDeviceError) SetPropertyReadFailure(value bool) (err error) {
	return instance.SetProperty("ReadFailure", (value))
}

// GetReadFailure gets the value of ReadFailure for the instance
func (instance *MSTapeProblemDeviceError) GetPropertyReadFailure() (value bool, err error) {
	retValue, err := instance.GetProperty("ReadFailure")
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

// SetReadWarning sets the value of ReadWarning for the instance
func (instance *MSTapeProblemDeviceError) SetPropertyReadWarning(value bool) (err error) {
	return instance.SetProperty("ReadWarning", (value))
}

// GetReadWarning gets the value of ReadWarning for the instance
func (instance *MSTapeProblemDeviceError) GetPropertyReadWarning() (value bool, err error) {
	retValue, err := instance.GetProperty("ReadWarning")
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

// SetScsiInterfaceError sets the value of ScsiInterfaceError for the instance
func (instance *MSTapeProblemDeviceError) SetPropertyScsiInterfaceError(value bool) (err error) {
	return instance.SetProperty("ScsiInterfaceError", (value))
}

// GetScsiInterfaceError gets the value of ScsiInterfaceError for the instance
func (instance *MSTapeProblemDeviceError) GetPropertyScsiInterfaceError() (value bool, err error) {
	retValue, err := instance.GetProperty("ScsiInterfaceError")
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

// SetTapeSnapped sets the value of TapeSnapped for the instance
func (instance *MSTapeProblemDeviceError) SetPropertyTapeSnapped(value bool) (err error) {
	return instance.SetProperty("TapeSnapped", (value))
}

// GetTapeSnapped gets the value of TapeSnapped for the instance
func (instance *MSTapeProblemDeviceError) GetPropertyTapeSnapped() (value bool, err error) {
	retValue, err := instance.GetProperty("TapeSnapped")
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

// SetTimetoCleanDrive sets the value of TimetoCleanDrive for the instance
func (instance *MSTapeProblemDeviceError) SetPropertyTimetoCleanDrive(value bool) (err error) {
	return instance.SetProperty("TimetoCleanDrive", (value))
}

// GetTimetoCleanDrive gets the value of TimetoCleanDrive for the instance
func (instance *MSTapeProblemDeviceError) GetPropertyTimetoCleanDrive() (value bool, err error) {
	retValue, err := instance.GetProperty("TimetoCleanDrive")
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

// SetUnsupportedFormat sets the value of UnsupportedFormat for the instance
func (instance *MSTapeProblemDeviceError) SetPropertyUnsupportedFormat(value bool) (err error) {
	return instance.SetProperty("UnsupportedFormat", (value))
}

// GetUnsupportedFormat gets the value of UnsupportedFormat for the instance
func (instance *MSTapeProblemDeviceError) GetPropertyUnsupportedFormat() (value bool, err error) {
	retValue, err := instance.GetProperty("UnsupportedFormat")
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

// SetWriteFailure sets the value of WriteFailure for the instance
func (instance *MSTapeProblemDeviceError) SetPropertyWriteFailure(value bool) (err error) {
	return instance.SetProperty("WriteFailure", (value))
}

// GetWriteFailure gets the value of WriteFailure for the instance
func (instance *MSTapeProblemDeviceError) GetPropertyWriteFailure() (value bool, err error) {
	retValue, err := instance.GetProperty("WriteFailure")
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

// SetWriteWarning sets the value of WriteWarning for the instance
func (instance *MSTapeProblemDeviceError) SetPropertyWriteWarning(value bool) (err error) {
	return instance.SetProperty("WriteWarning", (value))
}

// GetWriteWarning gets the value of WriteWarning for the instance
func (instance *MSTapeProblemDeviceError) GetPropertyWriteWarning() (value bool, err error) {
	retValue, err := instance.GetProperty("WriteWarning")
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
