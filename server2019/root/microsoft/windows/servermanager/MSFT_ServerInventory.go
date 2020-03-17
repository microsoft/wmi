// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ServerInventory struct
type MSFT_ServerInventory struct {
	cim.WmiInstance

	//
	ActivationStatus uint8

	//
	CbsTimestamp string

	//
	CeipSetting uint8

	//
	Description string

	//
	Domain string

	//
	Fqdn string

	//
	IsDomainJoined bool

	//
	IsWerSettingEnforcedByGP bool

	//
	Manufacturer string

	//
	Name string

	//
	ProcessorNames []string

	//
	ProductId string

	//
	TotalPhysicalMemory uint64

	//
	Type uint8

	//
	WerSetting uint8
}

// SetActivationStatus sets the value of ActivationStatus for the instance
func (instance *MSFT_ServerInventory) SetPropertyActivationStatus(value uint8) (err error) {
	return instance.SetProperty("ActivationStatus", value)
}

// GetActivationStatus gets the value of ActivationStatus for the instance
func (instance *MSFT_ServerInventory) GetPropertyActivationStatus() (value uint8, err error) {
	retValue, err := instance.GetProperty("ActivationStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCbsTimestamp sets the value of CbsTimestamp for the instance
func (instance *MSFT_ServerInventory) SetPropertyCbsTimestamp(value string) (err error) {
	return instance.SetProperty("CbsTimestamp", value)
}

// GetCbsTimestamp gets the value of CbsTimestamp for the instance
func (instance *MSFT_ServerInventory) GetPropertyCbsTimestamp() (value string, err error) {
	retValue, err := instance.GetProperty("CbsTimestamp")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCeipSetting sets the value of CeipSetting for the instance
func (instance *MSFT_ServerInventory) SetPropertyCeipSetting(value uint8) (err error) {
	return instance.SetProperty("CeipSetting", value)
}

// GetCeipSetting gets the value of CeipSetting for the instance
func (instance *MSFT_ServerInventory) GetPropertyCeipSetting() (value uint8, err error) {
	retValue, err := instance.GetProperty("CeipSetting")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_ServerInventory) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_ServerInventory) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDomain sets the value of Domain for the instance
func (instance *MSFT_ServerInventory) SetPropertyDomain(value string) (err error) {
	return instance.SetProperty("Domain", value)
}

// GetDomain gets the value of Domain for the instance
func (instance *MSFT_ServerInventory) GetPropertyDomain() (value string, err error) {
	retValue, err := instance.GetProperty("Domain")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFqdn sets the value of Fqdn for the instance
func (instance *MSFT_ServerInventory) SetPropertyFqdn(value string) (err error) {
	return instance.SetProperty("Fqdn", value)
}

// GetFqdn gets the value of Fqdn for the instance
func (instance *MSFT_ServerInventory) GetPropertyFqdn() (value string, err error) {
	retValue, err := instance.GetProperty("Fqdn")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsDomainJoined sets the value of IsDomainJoined for the instance
func (instance *MSFT_ServerInventory) SetPropertyIsDomainJoined(value bool) (err error) {
	return instance.SetProperty("IsDomainJoined", value)
}

// GetIsDomainJoined gets the value of IsDomainJoined for the instance
func (instance *MSFT_ServerInventory) GetPropertyIsDomainJoined() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDomainJoined")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsWerSettingEnforcedByGP sets the value of IsWerSettingEnforcedByGP for the instance
func (instance *MSFT_ServerInventory) SetPropertyIsWerSettingEnforcedByGP(value bool) (err error) {
	return instance.SetProperty("IsWerSettingEnforcedByGP", value)
}

// GetIsWerSettingEnforcedByGP gets the value of IsWerSettingEnforcedByGP for the instance
func (instance *MSFT_ServerInventory) GetPropertyIsWerSettingEnforcedByGP() (value bool, err error) {
	retValue, err := instance.GetProperty("IsWerSettingEnforcedByGP")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *MSFT_ServerInventory) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", value)
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *MSFT_ServerInventory) GetPropertyManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("Manufacturer")
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
func (instance *MSFT_ServerInventory) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_ServerInventory) GetPropertyName() (value string, err error) {
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

// SetProcessorNames sets the value of ProcessorNames for the instance
func (instance *MSFT_ServerInventory) SetPropertyProcessorNames(value []string) (err error) {
	return instance.SetProperty("ProcessorNames", value)
}

// GetProcessorNames gets the value of ProcessorNames for the instance
func (instance *MSFT_ServerInventory) GetPropertyProcessorNames() (value []string, err error) {
	retValue, err := instance.GetProperty("ProcessorNames")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProductId sets the value of ProductId for the instance
func (instance *MSFT_ServerInventory) SetPropertyProductId(value string) (err error) {
	return instance.SetProperty("ProductId", value)
}

// GetProductId gets the value of ProductId for the instance
func (instance *MSFT_ServerInventory) GetPropertyProductId() (value string, err error) {
	retValue, err := instance.GetProperty("ProductId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalPhysicalMemory sets the value of TotalPhysicalMemory for the instance
func (instance *MSFT_ServerInventory) SetPropertyTotalPhysicalMemory(value uint64) (err error) {
	return instance.SetProperty("TotalPhysicalMemory", value)
}

// GetTotalPhysicalMemory gets the value of TotalPhysicalMemory for the instance
func (instance *MSFT_ServerInventory) GetPropertyTotalPhysicalMemory() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalPhysicalMemory")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *MSFT_ServerInventory) SetPropertyType(value uint8) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *MSFT_ServerInventory) GetPropertyType() (value uint8, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWerSetting sets the value of WerSetting for the instance
func (instance *MSFT_ServerInventory) SetPropertyWerSetting(value uint8) (err error) {
	return instance.SetProperty("WerSetting", value)
}

// GetWerSetting gets the value of WerSetting for the instance
func (instance *MSFT_ServerInventory) GetPropertyWerSetting() (value uint8, err error) {
	retValue, err := instance.GetProperty("WerSetting")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}
