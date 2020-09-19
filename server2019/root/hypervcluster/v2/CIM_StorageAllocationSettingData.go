// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_StorageAllocationSettingData struct
type CIM_StorageAllocationSettingData struct {
	*CIM_ResourceAllocationSettingData

	// Access describes whether the allocated storage extent is 1 (readable), 2 (writeable), or 3 (both).
	///NOTE: This property is a copy of the CIM_StorageExtent.Access property, except deprecated values. See the description of the CIM_StorageExtent.Access property for details.
	Access StorageAllocationSettingData_Access

	// A unique identifier for the host extent. The identified host extent is used for the storage resource allocation.
	///NOTE: This property is a copy of the CIM_StorageExtent.Name property. See the description of CIM_StorageExtent.Name property for details.
	HostExtentName string

	// The HostExtentNameFormat property identifies the format that is used for the value of the HostExtentName property.
	///NOTE: This property is a copy of the CIM_StorageExtent.NameFormat property, excluding deprecated values. See the description of CIM_StorageExtent.NameFormat class for details.
	///7 = Serial Number/Vendor/Model (SNVM) SNVM is 3 strings representing the vendor name, product name within the vendor namespace, and the serial number within the model namespace. Strings are delimited with a '+'. Spaces may be included and are significant. The serial number is the text representation of the serial number in hexadecimal upper case. This represents the vendor and model ID from SCSI Inquiry data; the vendor field MUST be 8 characters wide and the product field MUST be 16 characters wide. For example,
	///'ACME____+SUPER DISK______+124437458' (_ is a space character)
	///9 = NAA as a generic format. See
	///http://standards.ieee.org/regauth/oui/tutorials/fibrecomp_id.html. Formatted as 16 or 32 unseparated uppercase hex characters (2 per binary byte). For example '21000020372D3C73'
	///10 = EUI as a generic format (EUI64) See
	///http://standards.ieee.org/regauth/oui/tutorials/EUI64.html.
	///Formatted as 16 unseparated uppercase hex characters (2 per binary byte)
	///11 = T10 vendor identifier format as returned by SCSI Inquiry VPD page 83, identifier type 1. See T10 SPC-3 specification. This is the 8-byte ASCII vendor ID from the T10 registry followed by a vendor specific ASCII identifier; spaces are permitted. For non SCSI volumes, 'SNVM' may be the most appropriate choice. 12 = OS Device Name (for LogicalDisks). See LogicalDisk Name description for details.
	HostExtentNameFormat StorageAllocationSettingData_HostExtentNameFormat

	// If the host extent is a SCSI volume, then the preferred source for SCSI volume names is SCSI VPD Page 83 responses.
	///NOTE: This property is a copy of the CIM_StorageExtent.NameNamespace property. See the description of CIM_StorageExtent.NameNamespace class for details.
	///Page 83 returns a list of identifiers for various device elements. The metadata for each identifier includes an Association field, identifiers with association of 0 apply to volumes. Page 83 supports several namespaces specified in the Type field in the identifier metadata. See SCSI SPC-3 specification.
	///2 = VPD Page 83, Type 3 NAA (NameFormat SHOULD be NAA)
	///3 = VPD Page 83, Type 2 EUI64 (NameFormat EUI)
	///4 = VPD Page 83, Type 1 T10 Vendor Identification
	///(NameFormat T10)
	///Less preferred volume namespaces from other interfaces:
	///5 = VPD page 80, Serial number (NameFormat SHOULD be Other)
	///6 = FC NodeWWN (NameFormat SHOULD be NAA or EUI)
	///7 = Serial Number/Vendor/Model (NameFormat SHOULD be SNVM) cThe preferred namespace for LogigicalDisk names is platform specific device namespace; see LogigicalDIsk Description.
	///8 = OS Device Namespace.
	HostExtentNameNamespace StorageAllocationSettingData_HostExtentNameNamespace

	// The HostExtentStartingAddress property identifies the starting address on the host storage extent identified by the value of the HostExtentName property that is used for the allocation of the virtual storage extent.
	///A value of NULL indicates that there is no direct mapping of the virtual storage extent onto the referenced host storage extent.
	///NOTE: This property is a copy of the CIM_BasedOn.StartingAddess property. See the description of CIM_BasedOn association for details.
	HostExtentStartingAddress uint64

	// Size in bytes of the blocks that are allocated at the host as the result of this storage resource allocation or storage resource allocation request. If the block size is variable, then the maximum block size in bytes should be specified. If the block size is unknown or if a block concept does not apply, then the value 1 shall be used.
	///NOTE: This property is a copy of the CIM_StorageExtent.BlockSize property. See the description of the CIM_StorageExtent.BlockSize property for details.
	HostResourceBlockSize uint64

	// A string describing the format of the HostExtentName property if the value of the HostExtentNameFormat property is 1 (Other).
	OtherHostExtentNameFormat string

	// A string describing the namespace of the HostExtentName property if the value of the HostExtentNameNamespace matches 1 (Other).
	OtherHostExtentNameNamespace string

	// Size in bytes of the blocks that are presented to the consumer as the result of this storage resource allocation or storage resource allocation request. If the block size is variable, then the maximum block size in bytes should be specified. If the block size is unknown or if a block concept does not apply, then the value 1 shall be used.
	///NOTE: The use of 1 (and not 0) to indicate that the blocksize is unknown still allows the use of the VirtualQuantity property to convey the size in blocks of size 1).
	///NOTE: This property is a copy of the CIM_StorageExtent.BlockSize property. See the description of the CIM_StorageExtent.BlockSize property for details.
	VirtualResourceBlockSize uint64
}

func NewCIM_StorageAllocationSettingDataEx1(instance *cim.WmiInstance) (newInstance *CIM_StorageAllocationSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_StorageAllocationSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func NewCIM_StorageAllocationSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_StorageAllocationSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_StorageAllocationSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

// SetAccess sets the value of Access for the instance
func (instance *CIM_StorageAllocationSettingData) SetPropertyAccess(value StorageAllocationSettingData_Access) (err error) {
	return instance.SetProperty("Access", (value))
}

// GetAccess gets the value of Access for the instance
func (instance *CIM_StorageAllocationSettingData) GetPropertyAccess() (value StorageAllocationSettingData_Access, err error) {
	retValue, err := instance.GetProperty("Access")
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

	value = StorageAllocationSettingData_Access(valuetmp)

	return
}

// SetHostExtentName sets the value of HostExtentName for the instance
func (instance *CIM_StorageAllocationSettingData) SetPropertyHostExtentName(value string) (err error) {
	return instance.SetProperty("HostExtentName", (value))
}

// GetHostExtentName gets the value of HostExtentName for the instance
func (instance *CIM_StorageAllocationSettingData) GetPropertyHostExtentName() (value string, err error) {
	retValue, err := instance.GetProperty("HostExtentName")
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

// SetHostExtentNameFormat sets the value of HostExtentNameFormat for the instance
func (instance *CIM_StorageAllocationSettingData) SetPropertyHostExtentNameFormat(value StorageAllocationSettingData_HostExtentNameFormat) (err error) {
	return instance.SetProperty("HostExtentNameFormat", (value))
}

// GetHostExtentNameFormat gets the value of HostExtentNameFormat for the instance
func (instance *CIM_StorageAllocationSettingData) GetPropertyHostExtentNameFormat() (value StorageAllocationSettingData_HostExtentNameFormat, err error) {
	retValue, err := instance.GetProperty("HostExtentNameFormat")
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

	value = StorageAllocationSettingData_HostExtentNameFormat(valuetmp)

	return
}

// SetHostExtentNameNamespace sets the value of HostExtentNameNamespace for the instance
func (instance *CIM_StorageAllocationSettingData) SetPropertyHostExtentNameNamespace(value StorageAllocationSettingData_HostExtentNameNamespace) (err error) {
	return instance.SetProperty("HostExtentNameNamespace", (value))
}

// GetHostExtentNameNamespace gets the value of HostExtentNameNamespace for the instance
func (instance *CIM_StorageAllocationSettingData) GetPropertyHostExtentNameNamespace() (value StorageAllocationSettingData_HostExtentNameNamespace, err error) {
	retValue, err := instance.GetProperty("HostExtentNameNamespace")
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

	value = StorageAllocationSettingData_HostExtentNameNamespace(valuetmp)

	return
}

// SetHostExtentStartingAddress sets the value of HostExtentStartingAddress for the instance
func (instance *CIM_StorageAllocationSettingData) SetPropertyHostExtentStartingAddress(value uint64) (err error) {
	return instance.SetProperty("HostExtentStartingAddress", (value))
}

// GetHostExtentStartingAddress gets the value of HostExtentStartingAddress for the instance
func (instance *CIM_StorageAllocationSettingData) GetPropertyHostExtentStartingAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("HostExtentStartingAddress")
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

// SetHostResourceBlockSize sets the value of HostResourceBlockSize for the instance
func (instance *CIM_StorageAllocationSettingData) SetPropertyHostResourceBlockSize(value uint64) (err error) {
	return instance.SetProperty("HostResourceBlockSize", (value))
}

// GetHostResourceBlockSize gets the value of HostResourceBlockSize for the instance
func (instance *CIM_StorageAllocationSettingData) GetPropertyHostResourceBlockSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("HostResourceBlockSize")
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

// SetOtherHostExtentNameFormat sets the value of OtherHostExtentNameFormat for the instance
func (instance *CIM_StorageAllocationSettingData) SetPropertyOtherHostExtentNameFormat(value string) (err error) {
	return instance.SetProperty("OtherHostExtentNameFormat", (value))
}

// GetOtherHostExtentNameFormat gets the value of OtherHostExtentNameFormat for the instance
func (instance *CIM_StorageAllocationSettingData) GetPropertyOtherHostExtentNameFormat() (value string, err error) {
	retValue, err := instance.GetProperty("OtherHostExtentNameFormat")
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

// SetOtherHostExtentNameNamespace sets the value of OtherHostExtentNameNamespace for the instance
func (instance *CIM_StorageAllocationSettingData) SetPropertyOtherHostExtentNameNamespace(value string) (err error) {
	return instance.SetProperty("OtherHostExtentNameNamespace", (value))
}

// GetOtherHostExtentNameNamespace gets the value of OtherHostExtentNameNamespace for the instance
func (instance *CIM_StorageAllocationSettingData) GetPropertyOtherHostExtentNameNamespace() (value string, err error) {
	retValue, err := instance.GetProperty("OtherHostExtentNameNamespace")
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

// SetVirtualResourceBlockSize sets the value of VirtualResourceBlockSize for the instance
func (instance *CIM_StorageAllocationSettingData) SetPropertyVirtualResourceBlockSize(value uint64) (err error) {
	return instance.SetProperty("VirtualResourceBlockSize", (value))
}

// GetVirtualResourceBlockSize gets the value of VirtualResourceBlockSize for the instance
func (instance *CIM_StorageAllocationSettingData) GetPropertyVirtualResourceBlockSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualResourceBlockSize")
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
