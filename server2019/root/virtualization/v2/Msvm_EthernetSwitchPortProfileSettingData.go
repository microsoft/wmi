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

// Msvm_EthernetSwitchPortProfileSettingData struct
type Msvm_EthernetSwitchPortProfileSettingData struct { 
	*Msvm_EthernetSwitchPortFeatureSettingData

	// 
	CdnLabelId uint32

	// 
	CdnLabelString string

	// 
	NetCfgInstanceId string

	// 
	PciBusNumber uint32

	// 
	PciDeviceNumber uint32

	// 
	PciFunctionNumber uint32

	// 
	PciSegmentNumber uint32

	// 
	ProfileData uint32

	// 
	ProfileId string

	// 
	ProfileName string

	// 
	VendorId string

	// 
	VendorName string
}

	func NewMsvm_EthernetSwitchPortProfileSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchPortProfileSettingData, err error) {tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_EthernetSwitchPortProfileSettingData {
	Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
	}
	

	func NewMsvm_EthernetSwitchPortProfileSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_EthernetSwitchPortProfileSettingData, err error) {tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_EthernetSwitchPortProfileSettingData {
	Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
	}
	

// SetCdnLabelId sets the value of CdnLabelId for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) SetPropertyCdnLabelId(value uint32) (err error) { 
    return instance.SetProperty("CdnLabelId", (value))
}

// GetCdnLabelId gets the value of CdnLabelId for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) GetPropertyCdnLabelId()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CdnLabelId")
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

// SetCdnLabelString sets the value of CdnLabelString for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) SetPropertyCdnLabelString(value string) (err error) { 
    return instance.SetProperty("CdnLabelString", (value))
}

// GetCdnLabelString gets the value of CdnLabelString for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) GetPropertyCdnLabelString()(value string, err error) { 
    retValue, err := instance.GetProperty("CdnLabelString")
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

// SetNetCfgInstanceId sets the value of NetCfgInstanceId for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) SetPropertyNetCfgInstanceId(value string) (err error) { 
    return instance.SetProperty("NetCfgInstanceId", (value))
}

// GetNetCfgInstanceId gets the value of NetCfgInstanceId for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) GetPropertyNetCfgInstanceId()(value string, err error) { 
    retValue, err := instance.GetProperty("NetCfgInstanceId")
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

// SetPciBusNumber sets the value of PciBusNumber for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) SetPropertyPciBusNumber(value uint32) (err error) { 
    return instance.SetProperty("PciBusNumber", (value))
}

// GetPciBusNumber gets the value of PciBusNumber for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) GetPropertyPciBusNumber()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PciBusNumber")
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

// SetPciDeviceNumber sets the value of PciDeviceNumber for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) SetPropertyPciDeviceNumber(value uint32) (err error) { 
    return instance.SetProperty("PciDeviceNumber", (value))
}

// GetPciDeviceNumber gets the value of PciDeviceNumber for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) GetPropertyPciDeviceNumber()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PciDeviceNumber")
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

// SetPciFunctionNumber sets the value of PciFunctionNumber for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) SetPropertyPciFunctionNumber(value uint32) (err error) { 
    return instance.SetProperty("PciFunctionNumber", (value))
}

// GetPciFunctionNumber gets the value of PciFunctionNumber for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) GetPropertyPciFunctionNumber()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PciFunctionNumber")
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

// SetPciSegmentNumber sets the value of PciSegmentNumber for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) SetPropertyPciSegmentNumber(value uint32) (err error) { 
    return instance.SetProperty("PciSegmentNumber", (value))
}

// GetPciSegmentNumber gets the value of PciSegmentNumber for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) GetPropertyPciSegmentNumber()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PciSegmentNumber")
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

// SetProfileData sets the value of ProfileData for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) SetPropertyProfileData(value uint32) (err error) { 
    return instance.SetProperty("ProfileData", (value))
}

// GetProfileData gets the value of ProfileData for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) GetPropertyProfileData()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ProfileData")
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

// SetProfileId sets the value of ProfileId for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) SetPropertyProfileId(value string) (err error) { 
    return instance.SetProperty("ProfileId", (value))
}

// GetProfileId gets the value of ProfileId for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) GetPropertyProfileId()(value string, err error) { 
    retValue, err := instance.GetProperty("ProfileId")
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

// SetProfileName sets the value of ProfileName for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) SetPropertyProfileName(value string) (err error) { 
    return instance.SetProperty("ProfileName", (value))
}

// GetProfileName gets the value of ProfileName for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) GetPropertyProfileName()(value string, err error) { 
    retValue, err := instance.GetProperty("ProfileName")
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

// SetVendorId sets the value of VendorId for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) SetPropertyVendorId(value string) (err error) { 
    return instance.SetProperty("VendorId", (value))
}

// GetVendorId gets the value of VendorId for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) GetPropertyVendorId()(value string, err error) { 
    retValue, err := instance.GetProperty("VendorId")
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

// SetVendorName sets the value of VendorName for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) SetPropertyVendorName(value string) (err error) { 
    return instance.SetProperty("VendorName", (value))
}

// GetVendorName gets the value of VendorName for the instance
func (instance *Msvm_EthernetSwitchPortProfileSettingData) GetPropertyVendorName()(value string, err error) { 
    retValue, err := instance.GetProperty("VendorName")
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
func  (instance* Msvm_EthernetSwitchPortProfileSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities"); 
	}
	

