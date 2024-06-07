// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// Msvm_ProcessorSettingData struct
type Msvm_ProcessorSettingData struct { 
	*CIM_ResourceAllocationSettingData

	// 
	AllowACountMCount bool

	// 
	ApicMode uint8

	// 
	CpuBrandString string

	// 
	CpuGroupId string

	// 
	DisableSpeculationControls bool

	// 
	EnableHostResourceProtection bool

	// 
	EnableLegacyApicMode bool

	// 
	EnablePageShattering uint8

	// 
	EnablePerfmonArchPmu bool

	// 
	EnablePerfmonIpt bool

	// 
	EnablePerfmonLbr bool

	// 
	EnablePerfmonPebs bool

	// 
	EnablePerfmonPmu bool

	// 
	EnableSocketTopology bool

	// 
	EnlightenmentSet string

	// 
	ExposeVirtualizationExtensions bool

	// 
	ExtendedVirtualizationExtensions ProcessorSettingData_ExtendedVirtualizationExtensions

	// 
	HideHypervisorPresent bool

	// 
	HwThreadsPerCore uint64

	// 
	L3CacheWays uint32

	// 
	L3ProcessorDistributionPolicy ProcessorSettingData_L3ProcessorDistributionPolicy

	// 
	LimitCPUID bool

	// 
	LimitProcessorFeatures bool

	// 
	LimitProcessorFeaturesMode ProcessorSettingData_LimitProcessorFeaturesMode

	// 
	MaxClusterCountPerSocket uint32

	// 
	MaxHwIsolatedGuests uint32

	// 
	MaxNumaNodesPerSocket uint64

	// 
	MaxProcessorCountPerL3 uint32

	// 
	MaxProcessorsPerNumaNode uint64

	// 
	PerfCpuFreqCapMhz uint32

	// 
	PhysicalAddressWidth uint32

	// 
	ProcessorFeatureSet string
}

	func NewMsvm_ProcessorSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_ProcessorSettingData, err error) {tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_ProcessorSettingData {
	CIM_ResourceAllocationSettingData: tmp,
	}
	return
	}
	

	func NewMsvm_ProcessorSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_ProcessorSettingData, err error) {tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_ProcessorSettingData {
	CIM_ResourceAllocationSettingData: tmp,
	}
	return
	}
	

// SetAllowACountMCount sets the value of AllowACountMCount for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyAllowACountMCount(value bool) (err error) { 
    return instance.SetProperty("AllowACountMCount", (value))
}

// GetAllowACountMCount gets the value of AllowACountMCount for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyAllowACountMCount()(value bool, err error) { 
    retValue, err := instance.GetProperty("AllowACountMCount")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetApicMode sets the value of ApicMode for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyApicMode(value uint8) (err error) { 
    return instance.SetProperty("ApicMode", (value))
}

// GetApicMode gets the value of ApicMode for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyApicMode()(value uint8, err error) { 
    retValue, err := instance.GetProperty("ApicMode")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetCpuBrandString sets the value of CpuBrandString for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyCpuBrandString(value string) (err error) { 
    return instance.SetProperty("CpuBrandString", (value))
}

// GetCpuBrandString gets the value of CpuBrandString for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyCpuBrandString()(value string, err error) { 
    retValue, err := instance.GetProperty("CpuBrandString")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetCpuGroupId sets the value of CpuGroupId for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyCpuGroupId(value string) (err error) { 
    return instance.SetProperty("CpuGroupId", (value))
}

// GetCpuGroupId gets the value of CpuGroupId for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyCpuGroupId()(value string, err error) { 
    retValue, err := instance.GetProperty("CpuGroupId")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
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
func (instance *Msvm_ProcessorSettingData) GetPropertyDisableSpeculationControls()(value bool, err error) { 
    retValue, err := instance.GetProperty("DisableSpeculationControls")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
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
func (instance *Msvm_ProcessorSettingData) GetPropertyEnableHostResourceProtection()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnableHostResourceProtection")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetEnableLegacyApicMode sets the value of EnableLegacyApicMode for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyEnableLegacyApicMode(value bool) (err error) { 
    return instance.SetProperty("EnableLegacyApicMode", (value))
}

// GetEnableLegacyApicMode gets the value of EnableLegacyApicMode for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyEnableLegacyApicMode()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnableLegacyApicMode")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetEnablePageShattering sets the value of EnablePageShattering for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyEnablePageShattering(value uint8) (err error) { 
    return instance.SetProperty("EnablePageShattering", (value))
}

// GetEnablePageShattering gets the value of EnablePageShattering for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyEnablePageShattering()(value uint8, err error) { 
    retValue, err := instance.GetProperty("EnablePageShattering")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetEnablePerfmonArchPmu sets the value of EnablePerfmonArchPmu for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyEnablePerfmonArchPmu(value bool) (err error) { 
    return instance.SetProperty("EnablePerfmonArchPmu", (value))
}

// GetEnablePerfmonArchPmu gets the value of EnablePerfmonArchPmu for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyEnablePerfmonArchPmu()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnablePerfmonArchPmu")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetEnablePerfmonIpt sets the value of EnablePerfmonIpt for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyEnablePerfmonIpt(value bool) (err error) { 
    return instance.SetProperty("EnablePerfmonIpt", (value))
}

// GetEnablePerfmonIpt gets the value of EnablePerfmonIpt for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyEnablePerfmonIpt()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnablePerfmonIpt")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
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
func (instance *Msvm_ProcessorSettingData) GetPropertyEnablePerfmonLbr()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnablePerfmonLbr")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
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
func (instance *Msvm_ProcessorSettingData) GetPropertyEnablePerfmonPebs()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnablePerfmonPebs")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
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
func (instance *Msvm_ProcessorSettingData) GetPropertyEnablePerfmonPmu()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnablePerfmonPmu")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetEnableSocketTopology sets the value of EnableSocketTopology for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyEnableSocketTopology(value bool) (err error) { 
    return instance.SetProperty("EnableSocketTopology", (value))
}

// GetEnableSocketTopology gets the value of EnableSocketTopology for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyEnableSocketTopology()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnableSocketTopology")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetEnlightenmentSet sets the value of EnlightenmentSet for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyEnlightenmentSet(value string) (err error) { 
    return instance.SetProperty("EnlightenmentSet", (value))
}

// GetEnlightenmentSet gets the value of EnlightenmentSet for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyEnlightenmentSet()(value string, err error) { 
    retValue, err := instance.GetProperty("EnlightenmentSet")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetExposeVirtualizationExtensions sets the value of ExposeVirtualizationExtensions for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyExposeVirtualizationExtensions(value bool) (err error) { 
    return instance.SetProperty("ExposeVirtualizationExtensions", (value))
}

// GetExposeVirtualizationExtensions gets the value of ExposeVirtualizationExtensions for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyExposeVirtualizationExtensions()(value bool, err error) { 
    retValue, err := instance.GetProperty("ExposeVirtualizationExtensions")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetExtendedVirtualizationExtensions sets the value of ExtendedVirtualizationExtensions for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyExtendedVirtualizationExtensions(value ProcessorSettingData_ExtendedVirtualizationExtensions) (err error) { 
    return instance.SetProperty("ExtendedVirtualizationExtensions", (value))
}

// GetExtendedVirtualizationExtensions gets the value of ExtendedVirtualizationExtensions for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyExtendedVirtualizationExtensions()(value ProcessorSettingData_ExtendedVirtualizationExtensions, err error) { 
    retValue, err := instance.GetProperty("ExtendedVirtualizationExtensions")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = ProcessorSettingData_ExtendedVirtualizationExtensions(valuetmp)

    return
}

// SetHideHypervisorPresent sets the value of HideHypervisorPresent for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyHideHypervisorPresent(value bool) (err error) { 
    return instance.SetProperty("HideHypervisorPresent", (value))
}

// GetHideHypervisorPresent gets the value of HideHypervisorPresent for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyHideHypervisorPresent()(value bool, err error) { 
    retValue, err := instance.GetProperty("HideHypervisorPresent")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
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
func (instance *Msvm_ProcessorSettingData) GetPropertyHwThreadsPerCore()(value uint64, err error) { 
    retValue, err := instance.GetProperty("HwThreadsPerCore")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetL3CacheWays sets the value of L3CacheWays for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyL3CacheWays(value uint32) (err error) { 
    return instance.SetProperty("L3CacheWays", (value))
}

// GetL3CacheWays gets the value of L3CacheWays for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyL3CacheWays()(value uint32, err error) { 
    retValue, err := instance.GetProperty("L3CacheWays")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetL3ProcessorDistributionPolicy sets the value of L3ProcessorDistributionPolicy for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyL3ProcessorDistributionPolicy(value ProcessorSettingData_L3ProcessorDistributionPolicy) (err error) { 
    return instance.SetProperty("L3ProcessorDistributionPolicy", (value))
}

// GetL3ProcessorDistributionPolicy gets the value of L3ProcessorDistributionPolicy for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyL3ProcessorDistributionPolicy()(value ProcessorSettingData_L3ProcessorDistributionPolicy, err error) { 
    retValue, err := instance.GetProperty("L3ProcessorDistributionPolicy")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = ProcessorSettingData_L3ProcessorDistributionPolicy(valuetmp)

    return
}

// SetLimitCPUID sets the value of LimitCPUID for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyLimitCPUID(value bool) (err error) { 
    return instance.SetProperty("LimitCPUID", (value))
}

// GetLimitCPUID gets the value of LimitCPUID for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyLimitCPUID()(value bool, err error) { 
    retValue, err := instance.GetProperty("LimitCPUID")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
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
func (instance *Msvm_ProcessorSettingData) GetPropertyLimitProcessorFeatures()(value bool, err error) { 
    retValue, err := instance.GetProperty("LimitProcessorFeatures")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetLimitProcessorFeaturesMode sets the value of LimitProcessorFeaturesMode for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyLimitProcessorFeaturesMode(value ProcessorSettingData_LimitProcessorFeaturesMode) (err error) { 
    return instance.SetProperty("LimitProcessorFeaturesMode", (value))
}

// GetLimitProcessorFeaturesMode gets the value of LimitProcessorFeaturesMode for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyLimitProcessorFeaturesMode()(value ProcessorSettingData_LimitProcessorFeaturesMode, err error) { 
    retValue, err := instance.GetProperty("LimitProcessorFeaturesMode")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = ProcessorSettingData_LimitProcessorFeaturesMode(valuetmp)

    return
}

// SetMaxClusterCountPerSocket sets the value of MaxClusterCountPerSocket for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyMaxClusterCountPerSocket(value uint32) (err error) { 
    return instance.SetProperty("MaxClusterCountPerSocket", (value))
}

// GetMaxClusterCountPerSocket gets the value of MaxClusterCountPerSocket for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyMaxClusterCountPerSocket()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxClusterCountPerSocket")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetMaxHwIsolatedGuests sets the value of MaxHwIsolatedGuests for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyMaxHwIsolatedGuests(value uint32) (err error) { 
    return instance.SetProperty("MaxHwIsolatedGuests", (value))
}

// GetMaxHwIsolatedGuests gets the value of MaxHwIsolatedGuests for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyMaxHwIsolatedGuests()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxHwIsolatedGuests")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetMaxNumaNodesPerSocket sets the value of MaxNumaNodesPerSocket for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyMaxNumaNodesPerSocket(value uint64) (err error) { 
    return instance.SetProperty("MaxNumaNodesPerSocket", (value))
}

// GetMaxNumaNodesPerSocket gets the value of MaxNumaNodesPerSocket for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyMaxNumaNodesPerSocket()(value uint64, err error) { 
    retValue, err := instance.GetProperty("MaxNumaNodesPerSocket")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetMaxProcessorCountPerL3 sets the value of MaxProcessorCountPerL3 for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyMaxProcessorCountPerL3(value uint32) (err error) { 
    return instance.SetProperty("MaxProcessorCountPerL3", (value))
}

// GetMaxProcessorCountPerL3 gets the value of MaxProcessorCountPerL3 for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyMaxProcessorCountPerL3()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxProcessorCountPerL3")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetMaxProcessorsPerNumaNode sets the value of MaxProcessorsPerNumaNode for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyMaxProcessorsPerNumaNode(value uint64) (err error) { 
    return instance.SetProperty("MaxProcessorsPerNumaNode", (value))
}

// GetMaxProcessorsPerNumaNode gets the value of MaxProcessorsPerNumaNode for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyMaxProcessorsPerNumaNode()(value uint64, err error) { 
    retValue, err := instance.GetProperty("MaxProcessorsPerNumaNode")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetPerfCpuFreqCapMhz sets the value of PerfCpuFreqCapMhz for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyPerfCpuFreqCapMhz(value uint32) (err error) { 
    return instance.SetProperty("PerfCpuFreqCapMhz", (value))
}

// GetPerfCpuFreqCapMhz gets the value of PerfCpuFreqCapMhz for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyPerfCpuFreqCapMhz()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PerfCpuFreqCapMhz")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetPhysicalAddressWidth sets the value of PhysicalAddressWidth for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyPhysicalAddressWidth(value uint32) (err error) { 
    return instance.SetProperty("PhysicalAddressWidth", (value))
}

// GetPhysicalAddressWidth gets the value of PhysicalAddressWidth for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyPhysicalAddressWidth()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PhysicalAddressWidth")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetProcessorFeatureSet sets the value of ProcessorFeatureSet for the instance
func (instance *Msvm_ProcessorSettingData) SetPropertyProcessorFeatureSet(value string) (err error) { 
    return instance.SetProperty("ProcessorFeatureSet", (value))
}

// GetProcessorFeatureSet gets the value of ProcessorFeatureSet for the instance
func (instance *Msvm_ProcessorSettingData) GetPropertyProcessorFeatureSet()(value string, err error) { 
    retValue, err := instance.GetProperty("ProcessorFeatureSet")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}
func  (instance* Msvm_ProcessorSettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_AllocationCapabilities"); 
	}
	

