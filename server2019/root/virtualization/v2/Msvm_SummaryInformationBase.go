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

// Msvm_SummaryInformationBase struct
type Msvm_SummaryInformationBase struct {
	*CIM_View

	//
	CreationTime string

	//
	EnabledState uint16

	//
	EnhancedSessionModeState uint16

	//
	HealthState uint16

	//
	HostComputerSystemName string

	//
	Name string

	//
	Notes string

	//
	NumberOfProcessors uint16

	//
	OperationalStatus []uint16

	//
	OtherEnabledState string

	//
	StatusDescriptions []string

	//
	UpTime uint64

	//
	Version string

	//
	VirtualSwitchNames []string

	//
	VirtualSystemSubType SummaryInformationBase_VirtualSystemSubType
}

func NewMsvm_SummaryInformationBaseEx1(instance *cim.WmiInstance) (newInstance *Msvm_SummaryInformationBase, err error) {
	tmp, err := NewCIM_ViewEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SummaryInformationBase{
		CIM_View: tmp,
	}
	return
}

func NewMsvm_SummaryInformationBaseEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SummaryInformationBase, err error) {
	tmp, err := NewCIM_ViewEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SummaryInformationBase{
		CIM_View: tmp,
	}
	return
}

// SetCreationTime sets the value of CreationTime for the instance
func (instance *Msvm_SummaryInformationBase) SetPropertyCreationTime(value string) (err error) {
	return instance.SetProperty("CreationTime", value)
}

// GetCreationTime gets the value of CreationTime for the instance
func (instance *Msvm_SummaryInformationBase) GetPropertyCreationTime() (value string, err error) {
	retValue, err := instance.GetProperty("CreationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabledState sets the value of EnabledState for the instance
func (instance *Msvm_SummaryInformationBase) SetPropertyEnabledState(value uint16) (err error) {
	return instance.SetProperty("EnabledState", value)
}

// GetEnabledState gets the value of EnabledState for the instance
func (instance *Msvm_SummaryInformationBase) GetPropertyEnabledState() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnabledState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnhancedSessionModeState sets the value of EnhancedSessionModeState for the instance
func (instance *Msvm_SummaryInformationBase) SetPropertyEnhancedSessionModeState(value uint16) (err error) {
	return instance.SetProperty("EnhancedSessionModeState", value)
}

// GetEnhancedSessionModeState gets the value of EnhancedSessionModeState for the instance
func (instance *Msvm_SummaryInformationBase) GetPropertyEnhancedSessionModeState() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnhancedSessionModeState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHealthState sets the value of HealthState for the instance
func (instance *Msvm_SummaryInformationBase) SetPropertyHealthState(value uint16) (err error) {
	return instance.SetProperty("HealthState", value)
}

// GetHealthState gets the value of HealthState for the instance
func (instance *Msvm_SummaryInformationBase) GetPropertyHealthState() (value uint16, err error) {
	retValue, err := instance.GetProperty("HealthState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostComputerSystemName sets the value of HostComputerSystemName for the instance
func (instance *Msvm_SummaryInformationBase) SetPropertyHostComputerSystemName(value string) (err error) {
	return instance.SetProperty("HostComputerSystemName", value)
}

// GetHostComputerSystemName gets the value of HostComputerSystemName for the instance
func (instance *Msvm_SummaryInformationBase) GetPropertyHostComputerSystemName() (value string, err error) {
	retValue, err := instance.GetProperty("HostComputerSystemName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *Msvm_SummaryInformationBase) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *Msvm_SummaryInformationBase) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNotes sets the value of Notes for the instance
func (instance *Msvm_SummaryInformationBase) SetPropertyNotes(value string) (err error) {
	return instance.SetProperty("Notes", value)
}

// GetNotes gets the value of Notes for the instance
func (instance *Msvm_SummaryInformationBase) GetPropertyNotes() (value string, err error) {
	retValue, err := instance.GetProperty("Notes")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfProcessors sets the value of NumberOfProcessors for the instance
func (instance *Msvm_SummaryInformationBase) SetPropertyNumberOfProcessors(value uint16) (err error) {
	return instance.SetProperty("NumberOfProcessors", value)
}

// GetNumberOfProcessors gets the value of NumberOfProcessors for the instance
func (instance *Msvm_SummaryInformationBase) GetPropertyNumberOfProcessors() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfProcessors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *Msvm_SummaryInformationBase) SetPropertyOperationalStatus(value []uint16) (err error) {
	return instance.SetProperty("OperationalStatus", value)
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *Msvm_SummaryInformationBase) GetPropertyOperationalStatus() (value []uint16, err error) {
	retValue, err := instance.GetProperty("OperationalStatus")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherEnabledState sets the value of OtherEnabledState for the instance
func (instance *Msvm_SummaryInformationBase) SetPropertyOtherEnabledState(value string) (err error) {
	return instance.SetProperty("OtherEnabledState", value)
}

// GetOtherEnabledState gets the value of OtherEnabledState for the instance
func (instance *Msvm_SummaryInformationBase) GetPropertyOtherEnabledState() (value string, err error) {
	retValue, err := instance.GetProperty("OtherEnabledState")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatusDescriptions sets the value of StatusDescriptions for the instance
func (instance *Msvm_SummaryInformationBase) SetPropertyStatusDescriptions(value []string) (err error) {
	return instance.SetProperty("StatusDescriptions", value)
}

// GetStatusDescriptions gets the value of StatusDescriptions for the instance
func (instance *Msvm_SummaryInformationBase) GetPropertyStatusDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("StatusDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUpTime sets the value of UpTime for the instance
func (instance *Msvm_SummaryInformationBase) SetPropertyUpTime(value uint64) (err error) {
	return instance.SetProperty("UpTime", value)
}

// GetUpTime gets the value of UpTime for the instance
func (instance *Msvm_SummaryInformationBase) GetPropertyUpTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("UpTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVersion sets the value of Version for the instance
func (instance *Msvm_SummaryInformationBase) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *Msvm_SummaryInformationBase) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualSwitchNames sets the value of VirtualSwitchNames for the instance
func (instance *Msvm_SummaryInformationBase) SetPropertyVirtualSwitchNames(value []string) (err error) {
	return instance.SetProperty("VirtualSwitchNames", value)
}

// GetVirtualSwitchNames gets the value of VirtualSwitchNames for the instance
func (instance *Msvm_SummaryInformationBase) GetPropertyVirtualSwitchNames() (value []string, err error) {
	retValue, err := instance.GetProperty("VirtualSwitchNames")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualSystemSubType sets the value of VirtualSystemSubType for the instance
func (instance *Msvm_SummaryInformationBase) SetPropertyVirtualSystemSubType(value SummaryInformationBase_VirtualSystemSubType) (err error) {
	return instance.SetProperty("VirtualSystemSubType", value)
}

// GetVirtualSystemSubType gets the value of VirtualSystemSubType for the instance
func (instance *Msvm_SummaryInformationBase) GetPropertyVirtualSystemSubType() (value SummaryInformationBase_VirtualSystemSubType, err error) {
	retValue, err := instance.GetProperty("VirtualSystemSubType")
	if err != nil {
		return
	}
	value, ok := retValue.(SummaryInformationBase_VirtualSystemSubType)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_SummaryInformationBase) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}
