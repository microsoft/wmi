// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_DNSClientCache struct
type MSFT_DNSClientCache struct {
	CIM_ManagedElement

	// 680
	Data string

	// 675
	DataLength uint16

	// 663
	Entry string

	// 664
	Name string

	// 676
	Section DNSClientCache_Section

	// 681
	Status DNSClientCache_Status

	// 674
	TimeToLive uint32

	// 665
	Type DNSClientCache_Type
}

// SetData sets the value of Data for the instance
func (instance *MSFT_DNSClientCache) SetPropertyData(value string) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *MSFT_DNSClientCache) GetPropertyData() (value string, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDataLength sets the value of DataLength for the instance
func (instance *MSFT_DNSClientCache) SetPropertyDataLength(value uint16) (err error) {
	return instance.SetProperty("DataLength", value)
}

// GetDataLength gets the value of DataLength for the instance
func (instance *MSFT_DNSClientCache) GetPropertyDataLength() (value uint16, err error) {
	retValue, err := instance.GetProperty("DataLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEntry sets the value of Entry for the instance
func (instance *MSFT_DNSClientCache) SetPropertyEntry(value string) (err error) {
	return instance.SetProperty("Entry", value)
}

// GetEntry gets the value of Entry for the instance
func (instance *MSFT_DNSClientCache) GetPropertyEntry() (value string, err error) {
	retValue, err := instance.GetProperty("Entry")
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
func (instance *MSFT_DNSClientCache) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_DNSClientCache) GetPropertyName() (value string, err error) {
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

// SetSection sets the value of Section for the instance
func (instance *MSFT_DNSClientCache) SetPropertySection(value DNSClientCache_Section) (err error) {
	return instance.SetProperty("Section", value)
}

// GetSection gets the value of Section for the instance
func (instance *MSFT_DNSClientCache) GetPropertySection() (value DNSClientCache_Section, err error) {
	retValue, err := instance.GetProperty("Section")
	if err != nil {
		return
	}
	value, ok := retValue.(DNSClientCache_Section)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *MSFT_DNSClientCache) SetPropertyStatus(value DNSClientCache_Status) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_DNSClientCache) GetPropertyStatus() (value DNSClientCache_Status, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(DNSClientCache_Status)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeToLive sets the value of TimeToLive for the instance
func (instance *MSFT_DNSClientCache) SetPropertyTimeToLive(value uint32) (err error) {
	return instance.SetProperty("TimeToLive", value)
}

// GetTimeToLive gets the value of TimeToLive for the instance
func (instance *MSFT_DNSClientCache) GetPropertyTimeToLive() (value uint32, err error) {
	retValue, err := instance.GetProperty("TimeToLive")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *MSFT_DNSClientCache) SetPropertyType(value DNSClientCache_Type) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *MSFT_DNSClientCache) GetPropertyType() (value DNSClientCache_Type, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(DNSClientCache_Type)
	if !ok {
		// TODO: Set an error
	}
	return
}

// 684

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DNSClientCache) Clear() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Clear")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
