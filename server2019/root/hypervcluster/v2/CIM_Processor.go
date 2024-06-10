// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.HyperVCluster.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_Processor struct
type CIM_Processor struct {
	*CIM_LogicalDevice

	// Processor address width in bits.
	AddressWidth uint16

	// The CPUStatus property that indicates the current status of the Processor. For example, the Processor might be disabled by the user (value=2), or disabled due to a POST error (value=3). Information in this property can be obtained from SMBIOS, the Type 4 structure, and the Status attribute.
	CPUStatus Processor_CPUStatus

	// The current speed (in MHz) of this Processor.
	CurrentClockSpeed uint32

	// Processor data width in bits.
	DataWidth uint16

	// The speed (in MHz) of the external bus interface (also known as the front side bus).
	ExternalBusClockSpeed uint32

	// The Processor family type. For example, values include "Pentium(R) processor with MMX(TM) technology" (value=14) and "68040" (value=96).
	Family Processor_Family

	// Loading of this Processor, averaged over the last minute, in Percent.
	LoadPercentage uint16

	// The maximum speed (in MHz) of this Processor.
	MaxClockSpeed uint32

	// A string that describes the Processor Family type. It is used when the Family property is set to 1 ("Other"). This string should be set to NULL when the Family property is any value other than 1.
	OtherFamilyDescription string

	// A free-form string that describes the role of the Processor, for example, "Central Processor" or "Math Processor".
	Role string

	// Stepping is a free-form string that indicates the revision level of the Processor within the Processor.Family.
	Stepping string

	// A globally unique identifier for the Processor. This identifier can be unique only within a Processor Family.
	UniqueID string

	// CPU socket information that includes data on how this Processor can be upgraded (if upgrades are supported). This property is an integer enumeration.
	UpgradeMethod Processor_UpgradeMethod
}

func NewCIM_ProcessorEx1(instance *cim.WmiInstance) (newInstance *CIM_Processor, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_Processor{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewCIM_ProcessorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Processor, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Processor{
		CIM_LogicalDevice: tmp,
	}
	return
}

// SetAddressWidth sets the value of AddressWidth for the instance
func (instance *CIM_Processor) SetPropertyAddressWidth(value uint16) (err error) {
	return instance.SetProperty("AddressWidth", (value))
}

// GetAddressWidth gets the value of AddressWidth for the instance
func (instance *CIM_Processor) GetPropertyAddressWidth() (value uint16, err error) {
	retValue, err := instance.GetProperty("AddressWidth")
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

// SetCPUStatus sets the value of CPUStatus for the instance
func (instance *CIM_Processor) SetPropertyCPUStatus(value Processor_CPUStatus) (err error) {
	return instance.SetProperty("CPUStatus", (value))
}

// GetCPUStatus gets the value of CPUStatus for the instance
func (instance *CIM_Processor) GetPropertyCPUStatus() (value Processor_CPUStatus, err error) {
	retValue, err := instance.GetProperty("CPUStatus")
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

	value = Processor_CPUStatus(valuetmp)

	return
}

// SetCurrentClockSpeed sets the value of CurrentClockSpeed for the instance
func (instance *CIM_Processor) SetPropertyCurrentClockSpeed(value uint32) (err error) {
	return instance.SetProperty("CurrentClockSpeed", (value))
}

// GetCurrentClockSpeed gets the value of CurrentClockSpeed for the instance
func (instance *CIM_Processor) GetPropertyCurrentClockSpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentClockSpeed")
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

// SetDataWidth sets the value of DataWidth for the instance
func (instance *CIM_Processor) SetPropertyDataWidth(value uint16) (err error) {
	return instance.SetProperty("DataWidth", (value))
}

// GetDataWidth gets the value of DataWidth for the instance
func (instance *CIM_Processor) GetPropertyDataWidth() (value uint16, err error) {
	retValue, err := instance.GetProperty("DataWidth")
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

// SetExternalBusClockSpeed sets the value of ExternalBusClockSpeed for the instance
func (instance *CIM_Processor) SetPropertyExternalBusClockSpeed(value uint32) (err error) {
	return instance.SetProperty("ExternalBusClockSpeed", (value))
}

// GetExternalBusClockSpeed gets the value of ExternalBusClockSpeed for the instance
func (instance *CIM_Processor) GetPropertyExternalBusClockSpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExternalBusClockSpeed")
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

// SetFamily sets the value of Family for the instance
func (instance *CIM_Processor) SetPropertyFamily(value Processor_Family) (err error) {
	return instance.SetProperty("Family", (value))
}

// GetFamily gets the value of Family for the instance
func (instance *CIM_Processor) GetPropertyFamily() (value Processor_Family, err error) {
	retValue, err := instance.GetProperty("Family")
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

	value = Processor_Family(valuetmp)

	return
}

// SetLoadPercentage sets the value of LoadPercentage for the instance
func (instance *CIM_Processor) SetPropertyLoadPercentage(value uint16) (err error) {
	return instance.SetProperty("LoadPercentage", (value))
}

// GetLoadPercentage gets the value of LoadPercentage for the instance
func (instance *CIM_Processor) GetPropertyLoadPercentage() (value uint16, err error) {
	retValue, err := instance.GetProperty("LoadPercentage")
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

// SetMaxClockSpeed sets the value of MaxClockSpeed for the instance
func (instance *CIM_Processor) SetPropertyMaxClockSpeed(value uint32) (err error) {
	return instance.SetProperty("MaxClockSpeed", (value))
}

// GetMaxClockSpeed gets the value of MaxClockSpeed for the instance
func (instance *CIM_Processor) GetPropertyMaxClockSpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxClockSpeed")
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

// SetOtherFamilyDescription sets the value of OtherFamilyDescription for the instance
func (instance *CIM_Processor) SetPropertyOtherFamilyDescription(value string) (err error) {
	return instance.SetProperty("OtherFamilyDescription", (value))
}

// GetOtherFamilyDescription gets the value of OtherFamilyDescription for the instance
func (instance *CIM_Processor) GetPropertyOtherFamilyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherFamilyDescription")
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

// SetRole sets the value of Role for the instance
func (instance *CIM_Processor) SetPropertyRole(value string) (err error) {
	return instance.SetProperty("Role", (value))
}

// GetRole gets the value of Role for the instance
func (instance *CIM_Processor) GetPropertyRole() (value string, err error) {
	retValue, err := instance.GetProperty("Role")
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

// SetStepping sets the value of Stepping for the instance
func (instance *CIM_Processor) SetPropertyStepping(value string) (err error) {
	return instance.SetProperty("Stepping", (value))
}

// GetStepping gets the value of Stepping for the instance
func (instance *CIM_Processor) GetPropertyStepping() (value string, err error) {
	retValue, err := instance.GetProperty("Stepping")
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

// SetUniqueID sets the value of UniqueID for the instance
func (instance *CIM_Processor) SetPropertyUniqueID(value string) (err error) {
	return instance.SetProperty("UniqueID", (value))
}

// GetUniqueID gets the value of UniqueID for the instance
func (instance *CIM_Processor) GetPropertyUniqueID() (value string, err error) {
	retValue, err := instance.GetProperty("UniqueID")
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

// SetUpgradeMethod sets the value of UpgradeMethod for the instance
func (instance *CIM_Processor) SetPropertyUpgradeMethod(value Processor_UpgradeMethod) (err error) {
	return instance.SetProperty("UpgradeMethod", (value))
}

// GetUpgradeMethod gets the value of UpgradeMethod for the instance
func (instance *CIM_Processor) GetPropertyUpgradeMethod() (value Processor_UpgradeMethod, err error) {
	retValue, err := instance.GetProperty("UpgradeMethod")
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

	value = Processor_UpgradeMethod(valuetmp)

	return
}
