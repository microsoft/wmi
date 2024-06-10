// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_VirtualHardDiskSettingData struct
type Msvm_VirtualHardDiskSettingData struct {
	*CIM_SettingData

	//
	BlockSize uint32

	//
	DataAlignment uint64

	//
	Format uint16

	//
	IsPmemCompatible bool

	//
	LogicalSectorSize uint32

	//
	MaxInternalSize uint64

	//
	ParentIdentifier string

	//
	ParentPath string

	//
	ParentTimestamp string

	//
	Path string

	//
	PhysicalSectorSize uint32

	//
	PmemAddressAbstractionType uint16

	//
	Type uint16

	//
	VirtualDiskId string
}

func NewMsvm_VirtualHardDiskSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualHardDiskSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualHardDiskSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_VirtualHardDiskSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualHardDiskSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualHardDiskSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetBlockSize sets the value of BlockSize for the instance
func (instance *Msvm_VirtualHardDiskSettingData) SetPropertyBlockSize(value uint32) (err error) {
	return instance.SetProperty("BlockSize", (value))
}

// GetBlockSize gets the value of BlockSize for the instance
func (instance *Msvm_VirtualHardDiskSettingData) GetPropertyBlockSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("BlockSize")
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

// SetDataAlignment sets the value of DataAlignment for the instance
func (instance *Msvm_VirtualHardDiskSettingData) SetPropertyDataAlignment(value uint64) (err error) {
	return instance.SetProperty("DataAlignment", (value))
}

// GetDataAlignment gets the value of DataAlignment for the instance
func (instance *Msvm_VirtualHardDiskSettingData) GetPropertyDataAlignment() (value uint64, err error) {
	retValue, err := instance.GetProperty("DataAlignment")
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

// SetFormat sets the value of Format for the instance
func (instance *Msvm_VirtualHardDiskSettingData) SetPropertyFormat(value uint16) (err error) {
	return instance.SetProperty("Format", (value))
}

// GetFormat gets the value of Format for the instance
func (instance *Msvm_VirtualHardDiskSettingData) GetPropertyFormat() (value uint16, err error) {
	retValue, err := instance.GetProperty("Format")
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

// SetIsPmemCompatible sets the value of IsPmemCompatible for the instance
func (instance *Msvm_VirtualHardDiskSettingData) SetPropertyIsPmemCompatible(value bool) (err error) {
	return instance.SetProperty("IsPmemCompatible", (value))
}

// GetIsPmemCompatible gets the value of IsPmemCompatible for the instance
func (instance *Msvm_VirtualHardDiskSettingData) GetPropertyIsPmemCompatible() (value bool, err error) {
	retValue, err := instance.GetProperty("IsPmemCompatible")
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

// SetLogicalSectorSize sets the value of LogicalSectorSize for the instance
func (instance *Msvm_VirtualHardDiskSettingData) SetPropertyLogicalSectorSize(value uint32) (err error) {
	return instance.SetProperty("LogicalSectorSize", (value))
}

// GetLogicalSectorSize gets the value of LogicalSectorSize for the instance
func (instance *Msvm_VirtualHardDiskSettingData) GetPropertyLogicalSectorSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogicalSectorSize")
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

// SetMaxInternalSize sets the value of MaxInternalSize for the instance
func (instance *Msvm_VirtualHardDiskSettingData) SetPropertyMaxInternalSize(value uint64) (err error) {
	return instance.SetProperty("MaxInternalSize", (value))
}

// GetMaxInternalSize gets the value of MaxInternalSize for the instance
func (instance *Msvm_VirtualHardDiskSettingData) GetPropertyMaxInternalSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxInternalSize")
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

// SetParentIdentifier sets the value of ParentIdentifier for the instance
func (instance *Msvm_VirtualHardDiskSettingData) SetPropertyParentIdentifier(value string) (err error) {
	return instance.SetProperty("ParentIdentifier", (value))
}

// GetParentIdentifier gets the value of ParentIdentifier for the instance
func (instance *Msvm_VirtualHardDiskSettingData) GetPropertyParentIdentifier() (value string, err error) {
	retValue, err := instance.GetProperty("ParentIdentifier")
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

// SetParentPath sets the value of ParentPath for the instance
func (instance *Msvm_VirtualHardDiskSettingData) SetPropertyParentPath(value string) (err error) {
	return instance.SetProperty("ParentPath", (value))
}

// GetParentPath gets the value of ParentPath for the instance
func (instance *Msvm_VirtualHardDiskSettingData) GetPropertyParentPath() (value string, err error) {
	retValue, err := instance.GetProperty("ParentPath")
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

// SetParentTimestamp sets the value of ParentTimestamp for the instance
func (instance *Msvm_VirtualHardDiskSettingData) SetPropertyParentTimestamp(value string) (err error) {
	return instance.SetProperty("ParentTimestamp", (value))
}

// GetParentTimestamp gets the value of ParentTimestamp for the instance
func (instance *Msvm_VirtualHardDiskSettingData) GetPropertyParentTimestamp() (value string, err error) {
	retValue, err := instance.GetProperty("ParentTimestamp")
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

// SetPath sets the value of Path for the instance
func (instance *Msvm_VirtualHardDiskSettingData) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *Msvm_VirtualHardDiskSettingData) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
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

// SetPhysicalSectorSize sets the value of PhysicalSectorSize for the instance
func (instance *Msvm_VirtualHardDiskSettingData) SetPropertyPhysicalSectorSize(value uint32) (err error) {
	return instance.SetProperty("PhysicalSectorSize", (value))
}

// GetPhysicalSectorSize gets the value of PhysicalSectorSize for the instance
func (instance *Msvm_VirtualHardDiskSettingData) GetPropertyPhysicalSectorSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("PhysicalSectorSize")
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

// SetPmemAddressAbstractionType sets the value of PmemAddressAbstractionType for the instance
func (instance *Msvm_VirtualHardDiskSettingData) SetPropertyPmemAddressAbstractionType(value uint16) (err error) {
	return instance.SetProperty("PmemAddressAbstractionType", (value))
}

// GetPmemAddressAbstractionType gets the value of PmemAddressAbstractionType for the instance
func (instance *Msvm_VirtualHardDiskSettingData) GetPropertyPmemAddressAbstractionType() (value uint16, err error) {
	retValue, err := instance.GetProperty("PmemAddressAbstractionType")
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

// SetType sets the value of Type for the instance
func (instance *Msvm_VirtualHardDiskSettingData) SetPropertyType(value uint16) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *Msvm_VirtualHardDiskSettingData) GetPropertyType() (value uint16, err error) {
	retValue, err := instance.GetProperty("Type")
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

// SetVirtualDiskId sets the value of VirtualDiskId for the instance
func (instance *Msvm_VirtualHardDiskSettingData) SetPropertyVirtualDiskId(value string) (err error) {
	return instance.SetProperty("VirtualDiskId", (value))
}

// GetVirtualDiskId gets the value of VirtualDiskId for the instance
func (instance *Msvm_VirtualHardDiskSettingData) GetPropertyVirtualDiskId() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualDiskId")
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
