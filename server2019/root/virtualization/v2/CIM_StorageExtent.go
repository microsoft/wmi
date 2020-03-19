// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_StorageExtent struct
type CIM_StorageExtent struct {
	*CIM_LogicalDevice

	// Access describes whether the media is readable (value=1), writeable (value=2), or both (value=3). "Unknown" (0) and "Write Once" (4) can also be defined.
	Access StorageExtent_Access

	// Size in bytes of the blocks which form this StorageExtent. If variable block size, then the maximum block size in bytes should be specified. If the block size is unknown or if a block concept is not valid (for example, for AggregateExtents, Memory or LogicalDisks), enter a 1.
	BlockSize uint64

	// The maximum number of blocks, of size BlockSize, which are available for consumption when layering StorageExtents using the BasedOn association. This property only has meaning when this StorageExtent is an Antecedent reference in a BasedOn relationship. For example, a StorageExtent could be composed of 120 blocks. However, the Extent itself may use 20 blocks for redundancy data. If another StorageExtent is BasedOn this Extent, only 100 blocks would be available to it. This information ('100 blocks is available for consumption') is indicated in the ConsumableBlocks property.
	ConsumableBlocks uint64

	// Type of data organization used.
	DataOrganization StorageExtent_DataOrganization

	// Number of complete copies of data currently maintained.
	DataRedundancy uint16

	// Current value for Delta reservation. This is a percentage that specifies the amount of space that should be reserved in a replica for caching changes.
	DeltaReservation uint8

	// ErrorMethodology is a free-form string describing the type of error detection and correction supported by this StorageExtent.
	ErrorMethodology string

	// StorageExtents have additional status information beyond that captured in the OperationalStatus and other properties, inherited from ManagedSystemElement. This additional information (for example, "Protection Disabled", value=9) is captured in the ExtentStatus property.
	///'In-Band Access Granted' says that access to data on an extent is granted to some consumer and is only valid when 'Exported' is also set. It is set as a side effect of PrivilegeManagementService.ChangeAccess or equivalent interfaces.
	///'Imported' indicates that the extent is used in the current system, but known to be managed by some other system. For example, a server imports volumes from a disk array.
	///'Exported' indicates the extent is meant to be used by some comsumer. A disk array's logical units are exported.
	///Intermediate composite extents may be neither imported nor exported.
	ExtentStatus []StorageExtent_ExtentStatus

	// True indicates that the underlying StorageExtent(s) participate in a StorageRedundancyGroup.
	IsBasedOnUnderlyingRedundancy bool

	// The list here applies to all StorageExtent subclasses. Please look at the Description in each subclass for guidelines on the approriate values for that subclass. Note that any of these formats could apply to a CompositeExtent.
	///
	///Note - this property originally touched on two concepts that are now separated into this property and NameNamespace. Values 2,3,4,5,6, and 8 are retained for backwards compatibility but are deprecated in lieu of the corresponding values in CIM_StorageExtent.NameNamespace.
	///
	///For example, the preferred source for SCSI virtual (RAID) disk names is from Inquiry VPD page 83 response, type 3 identifiers. These will have NameFormat set to 'NAA' and NameNamespace to 'VPD83Type3'.
	///
	///Format of the Name property. Values for extents representing SCSI volumes are (per SCSI SPC-3):
	///2 = VPD Page 83, NAA IEEE Registered Extended (VPD83NAA6)
	///(DEPRECATED)
	///3 = VPD Page 83, NAA IEEE Registered (VPD83NAA5)
	///(DEPRECATED)
	///4 = VPD Page 83, (VPD83Type2) (DEPRECATED)
	///5 = VPD Page 83,
	///T10 Vendor Identification (VPD83Type1) (DEPRECATED)
	///6 = VPD Page 83, Vendor Specific (VPD83Type0) (DEPRECATED)
	///7 = Serial Number/Vendor/Model (SNVM) SNVM is 3 strings representing the vendor name, product name within the vendor namespace, and the serial number within the model namespace. Strings are delimited with a '+'. Spaces may be included and are significant. The serial number is the text representation of the serial number in hexadecimal upper case. This represents the vendor and model ID from SCSI Inquiry data; the vendor field MUST be 8 characters wide and the product field MUST be 16 characters wide. For example,
	///'ACME____+SUPER DISK______+124437458' (_ is a space character)
	///8 = Node WWN (for single LUN/controller) (NodeWWN)
	///(DEPRECATED)
	///9 = NAA as a generic format. See
	///http://standards.ieee.org/regauth/oui/tutorials/fibrecomp_id.html. Formatted as 16 or 32 unseparated uppercase hex characters (2 per binary byte). For example '21000020372D3C73'
	///10 = EUI as a generic format (EUI64) See
	///http://standards.ieee.org/regauth/oui/tutorials/EUI64.html.
	///Formatted as 16 unseparated uppercase hex characters (2 per binary byte)
	///11 = T10 vendor identifier format as returned by SCSI Inquiry VPD page 83, identifier type 1. See T10 SPC-3 specification. This is the 8-byte ASCII vendor ID from the T10 registry followed by a vendor specific ASCII identifier; spaces are permitted. For non SCSI volumes, 'SNVM' may be the most appropriate choice. 12 = OS Device Name (for LogicalDisks). See LogicalDisk Name description for details.
	NameFormat StorageExtent_NameFormat

	// The preferred source SCSI for volume names is SCSI VPD Page 83 responses. Page 83 returns a list of identifiers for various device elements. The metadata for each identifier includes an Association field, identifiers with association of 0 apply to volumes. Page 83 supports several namespaces specified in the Type field in the identifier metadata. See SCSI SPC-3 specification.
	///2 = VPD Page 83, Type 3 NAA (NameFormat SHOULD be NAA)
	///3 = VPD Page 83, Type 2 EUI64 (NameFormat EUI)
	///4 = VPD Page 83, Type 1 T10 Vendor Identification
	///(NameFormat T10)
	///Less preferred volume namespaces from other interfaces:
	///5 = VPD page 80, Serial number (NameFormat SHOULD be Other)
	///6 = FC NodeWWN (NameFormat SHOULD be NAA or EUI)
	///7 = Serial Number/Vendor/Model (NameFormat SHOULD be SNVM)
	///The preferred namespace for LogigicalDisk names is platform specific device namespace; see LogigicalDIsk Description.
	///8 = OS Device Namespace.
	NameNamespace StorageExtent_NameNamespace

	// Indicates whether or not there exists no single point of failure.
	NoSinglePointOfFailure bool

	// Total number of logically contiguous blocks, of size Block Size, which form this Extent. The total size of the Extent can be calculated by multiplying BlockSize by NumberOfBlocks. If the BlockSize is 1, this property is the total size of the Extent.
	NumberOfBlocks uint64

	// A string describing the format of the Name property when NameFormat includes the value 1, "Other".
	OtherNameFormat string

	// A string describing the namespace of the Name property when NameNamespace includes the value 1, "Other".
	OtherNameNamespace string

	// How many physical packages can currently fail without data loss. For example, in the storage domain, this might be disk spindles.
	PackageRedundancy uint16

	// If true, "Primordial" indicates that the containing System does not have the ability to create or delete this operational element. This is important because StorageExtents are assembled into higher-level abstractions using the BasedOn association. Although the higher-level abstractions can be created and deleted, the most basic, (i.e. primordial), hardware-based storage entities cannot. They are physically realized as part of the System, or are actually managed by some other System and imported as if they were physically realized. In other words, a Primordial StorageExtent exists in, but is not created by its System and conversely a non-Primordial StorageExtent is created in the context of its System. For StorageVolumes, this property will generally be false. One use of this property is to enable algorithms that aggregate StorageExtent.ConsumableSpace across all, StorageExtents but that also want to distinquish the space that underlies Primordial StoragePools. Since implementations are not required to surface all Component StorageExtents of a StoragePool, this information is not accessible in any other way.
	Primordial bool

	// A free form string describing the media and/or its use.
	Purpose string

	// Boolean set to TRUE if the Storage is sequentially accessed by a MediaAccessDevice. A TapePartition is an example of a sequentially accessed StorageExtent. StorageVolumes, Disk Partitions and LogicalDisks represent randomly accessed Extents.
	SequentialAccess bool
}

func NewCIM_StorageExtentEx1(instance *cim.WmiInstance) (newInstance *CIM_StorageExtent, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_StorageExtent{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewCIM_StorageExtentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_StorageExtent, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_StorageExtent{
		CIM_LogicalDevice: tmp,
	}
	return
}

// SetAccess sets the value of Access for the instance
func (instance *CIM_StorageExtent) SetPropertyAccess(value StorageExtent_Access) (err error) {
	return instance.SetProperty("Access", value)
}

// GetAccess gets the value of Access for the instance
func (instance *CIM_StorageExtent) GetPropertyAccess() (value StorageExtent_Access, err error) {
	retValue, err := instance.GetProperty("Access")
	if err != nil {
		return
	}
	value, ok := retValue.(StorageExtent_Access)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBlockSize sets the value of BlockSize for the instance
func (instance *CIM_StorageExtent) SetPropertyBlockSize(value uint64) (err error) {
	return instance.SetProperty("BlockSize", value)
}

// GetBlockSize gets the value of BlockSize for the instance
func (instance *CIM_StorageExtent) GetPropertyBlockSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("BlockSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConsumableBlocks sets the value of ConsumableBlocks for the instance
func (instance *CIM_StorageExtent) SetPropertyConsumableBlocks(value uint64) (err error) {
	return instance.SetProperty("ConsumableBlocks", value)
}

// GetConsumableBlocks gets the value of ConsumableBlocks for the instance
func (instance *CIM_StorageExtent) GetPropertyConsumableBlocks() (value uint64, err error) {
	retValue, err := instance.GetProperty("ConsumableBlocks")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDataOrganization sets the value of DataOrganization for the instance
func (instance *CIM_StorageExtent) SetPropertyDataOrganization(value StorageExtent_DataOrganization) (err error) {
	return instance.SetProperty("DataOrganization", value)
}

// GetDataOrganization gets the value of DataOrganization for the instance
func (instance *CIM_StorageExtent) GetPropertyDataOrganization() (value StorageExtent_DataOrganization, err error) {
	retValue, err := instance.GetProperty("DataOrganization")
	if err != nil {
		return
	}
	value, ok := retValue.(StorageExtent_DataOrganization)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDataRedundancy sets the value of DataRedundancy for the instance
func (instance *CIM_StorageExtent) SetPropertyDataRedundancy(value uint16) (err error) {
	return instance.SetProperty("DataRedundancy", value)
}

// GetDataRedundancy gets the value of DataRedundancy for the instance
func (instance *CIM_StorageExtent) GetPropertyDataRedundancy() (value uint16, err error) {
	retValue, err := instance.GetProperty("DataRedundancy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeltaReservation sets the value of DeltaReservation for the instance
func (instance *CIM_StorageExtent) SetPropertyDeltaReservation(value uint8) (err error) {
	return instance.SetProperty("DeltaReservation", value)
}

// GetDeltaReservation gets the value of DeltaReservation for the instance
func (instance *CIM_StorageExtent) GetPropertyDeltaReservation() (value uint8, err error) {
	retValue, err := instance.GetProperty("DeltaReservation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorMethodology sets the value of ErrorMethodology for the instance
func (instance *CIM_StorageExtent) SetPropertyErrorMethodology(value string) (err error) {
	return instance.SetProperty("ErrorMethodology", value)
}

// GetErrorMethodology gets the value of ErrorMethodology for the instance
func (instance *CIM_StorageExtent) GetPropertyErrorMethodology() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorMethodology")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExtentStatus sets the value of ExtentStatus for the instance
func (instance *CIM_StorageExtent) SetPropertyExtentStatus(value []StorageExtent_ExtentStatus) (err error) {
	return instance.SetProperty("ExtentStatus", value)
}

// GetExtentStatus gets the value of ExtentStatus for the instance
func (instance *CIM_StorageExtent) GetPropertyExtentStatus() (value []StorageExtent_ExtentStatus, err error) {
	retValue, err := instance.GetProperty("ExtentStatus")
	if err != nil {
		return
	}
	value, ok := retValue.([]StorageExtent_ExtentStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsBasedOnUnderlyingRedundancy sets the value of IsBasedOnUnderlyingRedundancy for the instance
func (instance *CIM_StorageExtent) SetPropertyIsBasedOnUnderlyingRedundancy(value bool) (err error) {
	return instance.SetProperty("IsBasedOnUnderlyingRedundancy", value)
}

// GetIsBasedOnUnderlyingRedundancy gets the value of IsBasedOnUnderlyingRedundancy for the instance
func (instance *CIM_StorageExtent) GetPropertyIsBasedOnUnderlyingRedundancy() (value bool, err error) {
	retValue, err := instance.GetProperty("IsBasedOnUnderlyingRedundancy")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNameFormat sets the value of NameFormat for the instance
func (instance *CIM_StorageExtent) SetPropertyNameFormat(value StorageExtent_NameFormat) (err error) {
	return instance.SetProperty("NameFormat", value)
}

// GetNameFormat gets the value of NameFormat for the instance
func (instance *CIM_StorageExtent) GetPropertyNameFormat() (value StorageExtent_NameFormat, err error) {
	retValue, err := instance.GetProperty("NameFormat")
	if err != nil {
		return
	}
	value, ok := retValue.(StorageExtent_NameFormat)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNameNamespace sets the value of NameNamespace for the instance
func (instance *CIM_StorageExtent) SetPropertyNameNamespace(value StorageExtent_NameNamespace) (err error) {
	return instance.SetProperty("NameNamespace", value)
}

// GetNameNamespace gets the value of NameNamespace for the instance
func (instance *CIM_StorageExtent) GetPropertyNameNamespace() (value StorageExtent_NameNamespace, err error) {
	retValue, err := instance.GetProperty("NameNamespace")
	if err != nil {
		return
	}
	value, ok := retValue.(StorageExtent_NameNamespace)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNoSinglePointOfFailure sets the value of NoSinglePointOfFailure for the instance
func (instance *CIM_StorageExtent) SetPropertyNoSinglePointOfFailure(value bool) (err error) {
	return instance.SetProperty("NoSinglePointOfFailure", value)
}

// GetNoSinglePointOfFailure gets the value of NoSinglePointOfFailure for the instance
func (instance *CIM_StorageExtent) GetPropertyNoSinglePointOfFailure() (value bool, err error) {
	retValue, err := instance.GetProperty("NoSinglePointOfFailure")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfBlocks sets the value of NumberOfBlocks for the instance
func (instance *CIM_StorageExtent) SetPropertyNumberOfBlocks(value uint64) (err error) {
	return instance.SetProperty("NumberOfBlocks", value)
}

// GetNumberOfBlocks gets the value of NumberOfBlocks for the instance
func (instance *CIM_StorageExtent) GetPropertyNumberOfBlocks() (value uint64, err error) {
	retValue, err := instance.GetProperty("NumberOfBlocks")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherNameFormat sets the value of OtherNameFormat for the instance
func (instance *CIM_StorageExtent) SetPropertyOtherNameFormat(value string) (err error) {
	return instance.SetProperty("OtherNameFormat", value)
}

// GetOtherNameFormat gets the value of OtherNameFormat for the instance
func (instance *CIM_StorageExtent) GetPropertyOtherNameFormat() (value string, err error) {
	retValue, err := instance.GetProperty("OtherNameFormat")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherNameNamespace sets the value of OtherNameNamespace for the instance
func (instance *CIM_StorageExtent) SetPropertyOtherNameNamespace(value string) (err error) {
	return instance.SetProperty("OtherNameNamespace", value)
}

// GetOtherNameNamespace gets the value of OtherNameNamespace for the instance
func (instance *CIM_StorageExtent) GetPropertyOtherNameNamespace() (value string, err error) {
	retValue, err := instance.GetProperty("OtherNameNamespace")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPackageRedundancy sets the value of PackageRedundancy for the instance
func (instance *CIM_StorageExtent) SetPropertyPackageRedundancy(value uint16) (err error) {
	return instance.SetProperty("PackageRedundancy", value)
}

// GetPackageRedundancy gets the value of PackageRedundancy for the instance
func (instance *CIM_StorageExtent) GetPropertyPackageRedundancy() (value uint16, err error) {
	retValue, err := instance.GetProperty("PackageRedundancy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrimordial sets the value of Primordial for the instance
func (instance *CIM_StorageExtent) SetPropertyPrimordial(value bool) (err error) {
	return instance.SetProperty("Primordial", value)
}

// GetPrimordial gets the value of Primordial for the instance
func (instance *CIM_StorageExtent) GetPropertyPrimordial() (value bool, err error) {
	retValue, err := instance.GetProperty("Primordial")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPurpose sets the value of Purpose for the instance
func (instance *CIM_StorageExtent) SetPropertyPurpose(value string) (err error) {
	return instance.SetProperty("Purpose", value)
}

// GetPurpose gets the value of Purpose for the instance
func (instance *CIM_StorageExtent) GetPropertyPurpose() (value string, err error) {
	retValue, err := instance.GetProperty("Purpose")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSequentialAccess sets the value of SequentialAccess for the instance
func (instance *CIM_StorageExtent) SetPropertySequentialAccess(value bool) (err error) {
	return instance.SetProperty("SequentialAccess", value)
}

// GetSequentialAccess gets the value of SequentialAccess for the instance
func (instance *CIM_StorageExtent) GetPropertySequentialAccess() (value bool, err error) {
	retValue, err := instance.GetProperty("SequentialAccess")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
