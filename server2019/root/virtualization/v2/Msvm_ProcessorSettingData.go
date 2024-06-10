// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_ProcessorSettingData struct
type Msvm_ProcessorSettingData struct {
	*CIM_ResourceAllocationSettingData

	//
	CpuGroupId string

	//
	DisableSpeculationControls bool

	//
	EnableHostResourceProtection bool

	//
	EnablePerfmonLbr bool

	//
	EnablePerfmonPebs bool

	//
	EnablePerfmonPmu bool

	//
	ExposeVirtualizationExtensions bool

	//
	HideHypervisorPresent bool

	//
	HwThreadsPerCore uint64

	//
	LimitCPUID bool

	//
	LimitProcessorFeatures bool

	//
	MaxNumaNodesPerSocket uint64

	//
	MaxProcessorsPerNumaNode uint64
}

func NewMsvm_ProcessorSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_ProcessorSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ProcessorSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func NewMsvm_ProcessorSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ProcessorSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ProcessorSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

// SetCpuGroupId sets the value of CpuGroupId for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyCpuGroupId(value string) (err error) {
	return instance.SetProperty("CpuGroupId", (value))
}

// GetCpuGroupId gets the value of CpuGroupId for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyCpuGroupId() (value string, err error) {
	retValue, err := instance.GetProperty("CpuGroupId")
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

// SetDisableSpeculationControls sets the value of DisableSpeculationControls for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyDisableSpeculationControls(value bool) (err error) {
	return instance.SetProperty("DisableSpeculationControls", (value))
}

// GetDisableSpeculationControls gets the value of DisableSpeculationControls for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyDisableSpeculationControls() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableSpeculationControls")
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

// SetEnableHostResourceProtection sets the value of EnableHostResourceProtection for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyEnableHostResourceProtection(value bool) (err error) {
	return instance.SetProperty("EnableHostResourceProtection", (value))
}

// GetEnableHostResourceProtection gets the value of EnableHostResourceProtection for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyEnableHostResourceProtection() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableHostResourceProtection")
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

// SetEnablePerfmonLbr sets the value of EnablePerfmonLbr for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyEnablePerfmonLbr(value bool) (err error) {
	return instance.SetProperty("EnablePerfmonLbr", (value))
}

// GetEnablePerfmonLbr gets the value of EnablePerfmonLbr for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyEnablePerfmonLbr() (value bool, err error) {
	retValue, err := instance.GetProperty("EnablePerfmonLbr")
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

// SetEnablePerfmonPebs sets the value of EnablePerfmonPebs for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyEnablePerfmonPebs(value bool) (err error) {
	return instance.SetProperty("EnablePerfmonPebs", (value))
}

// GetEnablePerfmonPebs gets the value of EnablePerfmonPebs for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyEnablePerfmonPebs() (value bool, err error) {
	retValue, err := instance.GetProperty("EnablePerfmonPebs")
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

// SetEnablePerfmonPmu sets the value of EnablePerfmonPmu for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyEnablePerfmonPmu(value bool) (err error) {
	return instance.SetProperty("EnablePerfmonPmu", (value))
}

// GetEnablePerfmonPmu gets the value of EnablePerfmonPmu for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyEnablePerfmonPmu() (value bool, err error) {
	retValue, err := instance.GetProperty("EnablePerfmonPmu")
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

// SetExposeVirtualizationExtensions sets the value of ExposeVirtualizationExtensions for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyExposeVirtualizationExtensions(value bool) (err error) {
	return instance.SetProperty("ExposeVirtualizationExtensions", (value))
}

// GetExposeVirtualizationExtensions gets the value of ExposeVirtualizationExtensions for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyExposeVirtualizationExtensions() (value bool, err error) {
	retValue, err := instance.GetProperty("ExposeVirtualizationExtensions")
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

// SetHideHypervisorPresent sets the value of HideHypervisorPresent for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyHideHypervisorPresent(value bool) (err error) {
	return instance.SetProperty("HideHypervisorPresent", (value))
}

// GetHideHypervisorPresent gets the value of HideHypervisorPresent for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyHideHypervisorPresent() (value bool, err error) {
	retValue, err := instance.GetProperty("HideHypervisorPresent")
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

// SetHwThreadsPerCore sets the value of HwThreadsPerCore for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyHwThreadsPerCore(value uint64) (err error) {
	return instance.SetProperty("HwThreadsPerCore", (value))
}

// GetHwThreadsPerCore gets the value of HwThreadsPerCore for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyHwThreadsPerCore() (value uint64, err error) {
	retValue, err := instance.GetProperty("HwThreadsPerCore")
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

// SetLimitCPUID sets the value of LimitCPUID for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyLimitCPUID(value bool) (err error) {
	return instance.SetProperty("LimitCPUID", (value))
}

// GetLimitCPUID gets the value of LimitCPUID for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyLimitCPUID() (value bool, err error) {
	retValue, err := instance.GetProperty("LimitCPUID")
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

// SetLimitProcessorFeatures sets the value of LimitProcessorFeatures for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyLimitProcessorFeatures(value bool) (err error) {
	return instance.SetProperty("LimitProcessorFeatures", (value))
}

// GetLimitProcessorFeatures gets the value of LimitProcessorFeatures for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyLimitProcessorFeatures() (value bool, err error) {
	retValue, err := instance.GetProperty("LimitProcessorFeatures")
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

// SetMaxNumaNodesPerSocket sets the value of MaxNumaNodesPerSocket for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyMaxNumaNodesPerSocket(value uint64) (err error) {
	return instance.SetProperty("MaxNumaNodesPerSocket", (value))
}

// GetMaxNumaNodesPerSocket gets the value of MaxNumaNodesPerSocket for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyMaxNumaNodesPerSocket() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxNumaNodesPerSocket")
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

// SetMaxProcessorsPerNumaNode sets the value of MaxProcessorsPerNumaNode for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyMaxProcessorsPerNumaNode(value uint64) (err error) {
	return instance.SetProperty("MaxProcessorsPerNumaNode", (value))
}

// GetMaxProcessorsPerNumaNode gets the value of MaxProcessorsPerNumaNode for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyMaxProcessorsPerNumaNode() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxProcessorsPerNumaNode")
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
func (instance *Msvm_ProcessorSettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}
