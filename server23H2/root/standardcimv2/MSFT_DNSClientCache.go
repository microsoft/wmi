// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_DNSClientCache struct
type MSFT_DNSClientCache struct {
	*CIM_ManagedElement

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

func NewMSFT_DNSClientCacheEx1(instance *cim.WmiInstance) (newInstance *MSFT_DNSClientCache, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_DNSClientCache{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMSFT_DNSClientCacheEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DNSClientCache, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DNSClientCache{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetData sets the value of Data for the instance
func (instance *MSFT_DNSClientCache) SetPropertyData(value string) (err error) {
	return instance.SetProperty("Data", (value))
}

// GetData gets the value of Data for the instance
func (instance *MSFT_DNSClientCache) GetPropertyData() (value string, err error) {
	retValue, err := instance.GetProperty("Data")
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

// SetDataLength sets the value of DataLength for the instance
func (instance *MSFT_DNSClientCache) SetPropertyDataLength(value uint16) (err error) {
	return instance.SetProperty("DataLength", (value))
}

// GetDataLength gets the value of DataLength for the instance
func (instance *MSFT_DNSClientCache) GetPropertyDataLength() (value uint16, err error) {
	retValue, err := instance.GetProperty("DataLength")
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

// SetEntry sets the value of Entry for the instance
func (instance *MSFT_DNSClientCache) SetPropertyEntry(value string) (err error) {
	return instance.SetProperty("Entry", (value))
}

// GetEntry gets the value of Entry for the instance
func (instance *MSFT_DNSClientCache) GetPropertyEntry() (value string, err error) {
	retValue, err := instance.GetProperty("Entry")
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

// SetName sets the value of Name for the instance
func (instance *MSFT_DNSClientCache) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_DNSClientCache) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetSection sets the value of Section for the instance
func (instance *MSFT_DNSClientCache) SetPropertySection(value DNSClientCache_Section) (err error) {
	return instance.SetProperty("Section", (value))
}

// GetSection gets the value of Section for the instance
func (instance *MSFT_DNSClientCache) GetPropertySection() (value DNSClientCache_Section, err error) {
	retValue, err := instance.GetProperty("Section")
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

	value = DNSClientCache_Section(valuetmp)

	return
}

// SetStatus sets the value of Status for the instance
func (instance *MSFT_DNSClientCache) SetPropertyStatus(value DNSClientCache_Status) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_DNSClientCache) GetPropertyStatus() (value DNSClientCache_Status, err error) {
	retValue, err := instance.GetProperty("Status")
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

	value = DNSClientCache_Status(valuetmp)

	return
}

// SetTimeToLive sets the value of TimeToLive for the instance
func (instance *MSFT_DNSClientCache) SetPropertyTimeToLive(value uint32) (err error) {
	return instance.SetProperty("TimeToLive", (value))
}

// GetTimeToLive gets the value of TimeToLive for the instance
func (instance *MSFT_DNSClientCache) GetPropertyTimeToLive() (value uint32, err error) {
	retValue, err := instance.GetProperty("TimeToLive")
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

// SetType sets the value of Type for the instance
func (instance *MSFT_DNSClientCache) SetPropertyType(value DNSClientCache_Type) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSFT_DNSClientCache) GetPropertyType() (value DNSClientCache_Type, err error) {
	retValue, err := instance.GetProperty("Type")
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

	value = DNSClientCache_Type(valuetmp)

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
