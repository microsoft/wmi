// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// XmlDictionaryReaderQuotas struct
type XmlDictionaryReaderQuotas struct {
	*cim.WmiInstance

	// The maximum allowed array length.
	MaxArrayLength int32

	// The maximum allowed bytes returned per read.
	MaxBytesPerRead int32

	// The maximum nested node depth per read.
	MaxDepth int32

	// The maximum characters allowed in a table name.
	MaxNameTableCharCount int32

	// The maximum characters allowed in XML element content.
	MaxStringContentLength int32
}

func NewXmlDictionaryReaderQuotasEx1(instance *cim.WmiInstance) (newInstance *XmlDictionaryReaderQuotas, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &XmlDictionaryReaderQuotas{
		WmiInstance: tmp,
	}
	return
}

func NewXmlDictionaryReaderQuotasEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *XmlDictionaryReaderQuotas, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &XmlDictionaryReaderQuotas{
		WmiInstance: tmp,
	}
	return
}

// SetMaxArrayLength sets the value of MaxArrayLength for the instance
func (instance *XmlDictionaryReaderQuotas) SetPropertyMaxArrayLength(value int32) (err error) {
	return instance.SetProperty("MaxArrayLength", (value))
}

// GetMaxArrayLength gets the value of MaxArrayLength for the instance
func (instance *XmlDictionaryReaderQuotas) GetPropertyMaxArrayLength() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxArrayLength")
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

	value = int32(valuetmp)

	return
}

// SetMaxBytesPerRead sets the value of MaxBytesPerRead for the instance
func (instance *XmlDictionaryReaderQuotas) SetPropertyMaxBytesPerRead(value int32) (err error) {
	return instance.SetProperty("MaxBytesPerRead", (value))
}

// GetMaxBytesPerRead gets the value of MaxBytesPerRead for the instance
func (instance *XmlDictionaryReaderQuotas) GetPropertyMaxBytesPerRead() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxBytesPerRead")
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

	value = int32(valuetmp)

	return
}

// SetMaxDepth sets the value of MaxDepth for the instance
func (instance *XmlDictionaryReaderQuotas) SetPropertyMaxDepth(value int32) (err error) {
	return instance.SetProperty("MaxDepth", (value))
}

// GetMaxDepth gets the value of MaxDepth for the instance
func (instance *XmlDictionaryReaderQuotas) GetPropertyMaxDepth() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxDepth")
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

	value = int32(valuetmp)

	return
}

// SetMaxNameTableCharCount sets the value of MaxNameTableCharCount for the instance
func (instance *XmlDictionaryReaderQuotas) SetPropertyMaxNameTableCharCount(value int32) (err error) {
	return instance.SetProperty("MaxNameTableCharCount", (value))
}

// GetMaxNameTableCharCount gets the value of MaxNameTableCharCount for the instance
func (instance *XmlDictionaryReaderQuotas) GetPropertyMaxNameTableCharCount() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxNameTableCharCount")
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

	value = int32(valuetmp)

	return
}

// SetMaxStringContentLength sets the value of MaxStringContentLength for the instance
func (instance *XmlDictionaryReaderQuotas) SetPropertyMaxStringContentLength(value int32) (err error) {
	return instance.SetProperty("MaxStringContentLength", (value))
}

// GetMaxStringContentLength gets the value of MaxStringContentLength for the instance
func (instance *XmlDictionaryReaderQuotas) GetPropertyMaxStringContentLength() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxStringContentLength")
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

	value = int32(valuetmp)

	return
}
