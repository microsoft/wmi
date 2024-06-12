// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
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

// MSiSCSI_BootConfiguration struct
type MSiSCSI_BootConfiguration struct {
	*cim.WmiInstance

	//
	Active bool

	// If TRUE dynamically discover boot device.
	DiscoverBootDevice bool

	// The InitiatorNode specifies the iSCSI name of the initiator node to use for the connection. If empty, then the adapter can choose any initiator node name.
	InitiatorNode string

	//
	InstanceName string

	// Options that affect how login is performed. See ISCSI_LoginOptions
	LoginOptions ISCSI_LoginOptions

	// LUN on target to use as boot device.
	LUN uint64

	// Authentication Password, for CHAP this is the shared secret to use when generating the response to the target challange. This field is a variable length array.
	Password []uint8

	// Size in bytes of Target Password.
	PasswordSize uint32

	// Security flags
	SecurityFlags uint64

	// TargetName specifies the iSCSI target name on which the boot device resides.
	TargetName string

	// Target portal to use for connection to the target.
	TargetPortal ISCSI_TargetPortal

	// **extra fields** Authentication Username, for CHAP this is the CHAP Name (CHAP_N) use when authenticating with the target. NOTE: This field is a variable length array, the field that follows this field starts immediately after the end of this field subject to appropriate padding.
	Username []uint8

	// Size in bytes of Target Username.
	UsernameSize uint32
}

func NewMSiSCSI_BootConfigurationEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_BootConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_BootConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSI_BootConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_BootConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_BootConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_BootConfiguration) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_BootConfiguration) GetPropertyActive() (value bool, err error) {
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

// SetDiscoverBootDevice sets the value of DiscoverBootDevice for the instance
func (instance *MSiSCSI_BootConfiguration) SetPropertyDiscoverBootDevice(value bool) (err error) {
	return instance.SetProperty("DiscoverBootDevice", (value))
}

// GetDiscoverBootDevice gets the value of DiscoverBootDevice for the instance
func (instance *MSiSCSI_BootConfiguration) GetPropertyDiscoverBootDevice() (value bool, err error) {
	retValue, err := instance.GetProperty("DiscoverBootDevice")
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

// SetInitiatorNode sets the value of InitiatorNode for the instance
func (instance *MSiSCSI_BootConfiguration) SetPropertyInitiatorNode(value string) (err error) {
	return instance.SetProperty("InitiatorNode", (value))
}

// GetInitiatorNode gets the value of InitiatorNode for the instance
func (instance *MSiSCSI_BootConfiguration) GetPropertyInitiatorNode() (value string, err error) {
	retValue, err := instance.GetProperty("InitiatorNode")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSiSCSI_BootConfiguration) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_BootConfiguration) GetPropertyInstanceName() (value string, err error) {
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

// SetLoginOptions sets the value of LoginOptions for the instance
func (instance *MSiSCSI_BootConfiguration) SetPropertyLoginOptions(value ISCSI_LoginOptions) (err error) {
	return instance.SetProperty("LoginOptions", (value))
}

// GetLoginOptions gets the value of LoginOptions for the instance
func (instance *MSiSCSI_BootConfiguration) GetPropertyLoginOptions() (value ISCSI_LoginOptions, err error) {
	retValue, err := instance.GetProperty("LoginOptions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ISCSI_LoginOptions)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ISCSI_LoginOptions is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ISCSI_LoginOptions(valuetmp)

	return
}

// SetLUN sets the value of LUN for the instance
func (instance *MSiSCSI_BootConfiguration) SetPropertyLUN(value uint64) (err error) {
	return instance.SetProperty("LUN", (value))
}

// GetLUN gets the value of LUN for the instance
func (instance *MSiSCSI_BootConfiguration) GetPropertyLUN() (value uint64, err error) {
	retValue, err := instance.GetProperty("LUN")
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

// SetPassword sets the value of Password for the instance
func (instance *MSiSCSI_BootConfiguration) SetPropertyPassword(value []uint8) (err error) {
	return instance.SetProperty("Password", (value))
}

// GetPassword gets the value of Password for the instance
func (instance *MSiSCSI_BootConfiguration) GetPropertyPassword() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Password")
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

// SetPasswordSize sets the value of PasswordSize for the instance
func (instance *MSiSCSI_BootConfiguration) SetPropertyPasswordSize(value uint32) (err error) {
	return instance.SetProperty("PasswordSize", (value))
}

// GetPasswordSize gets the value of PasswordSize for the instance
func (instance *MSiSCSI_BootConfiguration) GetPropertyPasswordSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("PasswordSize")
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

// SetSecurityFlags sets the value of SecurityFlags for the instance
func (instance *MSiSCSI_BootConfiguration) SetPropertySecurityFlags(value uint64) (err error) {
	return instance.SetProperty("SecurityFlags", (value))
}

// GetSecurityFlags gets the value of SecurityFlags for the instance
func (instance *MSiSCSI_BootConfiguration) GetPropertySecurityFlags() (value uint64, err error) {
	retValue, err := instance.GetProperty("SecurityFlags")
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

// SetTargetName sets the value of TargetName for the instance
func (instance *MSiSCSI_BootConfiguration) SetPropertyTargetName(value string) (err error) {
	return instance.SetProperty("TargetName", (value))
}

// GetTargetName gets the value of TargetName for the instance
func (instance *MSiSCSI_BootConfiguration) GetPropertyTargetName() (value string, err error) {
	retValue, err := instance.GetProperty("TargetName")
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

// SetTargetPortal sets the value of TargetPortal for the instance
func (instance *MSiSCSI_BootConfiguration) SetPropertyTargetPortal(value ISCSI_TargetPortal) (err error) {
	return instance.SetProperty("TargetPortal", (value))
}

// GetTargetPortal gets the value of TargetPortal for the instance
func (instance *MSiSCSI_BootConfiguration) GetPropertyTargetPortal() (value ISCSI_TargetPortal, err error) {
	retValue, err := instance.GetProperty("TargetPortal")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ISCSI_TargetPortal)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ISCSI_TargetPortal is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ISCSI_TargetPortal(valuetmp)

	return
}

// SetUsername sets the value of Username for the instance
func (instance *MSiSCSI_BootConfiguration) SetPropertyUsername(value []uint8) (err error) {
	return instance.SetProperty("Username", (value))
}

// GetUsername gets the value of Username for the instance
func (instance *MSiSCSI_BootConfiguration) GetPropertyUsername() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Username")
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

// SetUsernameSize sets the value of UsernameSize for the instance
func (instance *MSiSCSI_BootConfiguration) SetPropertyUsernameSize(value uint32) (err error) {
	return instance.SetProperty("UsernameSize", (value))
}

// GetUsernameSize gets the value of UsernameSize for the instance
func (instance *MSiSCSI_BootConfiguration) GetPropertyUsernameSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("UsernameSize")
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
