// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.Intel_ME
//////////////////////////////////////////////
package intel_me

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_RecordForLog struct
type CIM_RecordForLog struct {
	*CIM_ManagedElement

	// A locale indicates a particular geographical, political, or cultural region. The Locale specifies the language used in creating the RecordForLog data. If the Locale property is empty, it is assumed that the default locale is en_US (English).
	///The locale string consists of three sub-strings, separated by underscores:
	///- The first sub-string is the language code, as specified in ISO639.
	///- The second sub-string is the country code, as specified in ISO3166.
	///- The third sub-string is a variant, which is vendor specific.
	///For example, US English appears as: "en_US_WIN", where the "WIN" variant would specify a Windows browser-specific collation (if one exists). Since the variant is not standardized, it is not commonly used and generally is limited to easily recognizable values ("WIN", "UNIX", "EURO", etc.) used in standard environments. The language and country codes are required; the variant may be empty.
	Locale string

	// A string containing LogRecord data.
	///If the corresponding RecordFormat property is <empty>, or cannot be parsed according to the recommended format, RecordData should be interpreted as a free-form string. If the RecordFormat property contains parseable format information (as recommended in the RecordFormat Description qualifier), the RecordData string SHOULD be parsed in accordance with this format. In this case, RecordData SHOULD begin with the delimiter character and this character SHOULD be used to separate substrings in the manner described. The RecordData string can then be parsed by the data consumer and appropriately typed.
	RecordData string

	// A string describing the data structure of the information in the property, RecordData. If the RecordFormat string is <empty>, RecordData should be interpreted as a free-form string.
	///
	///To describe the data structure of RecordData, the RecordFormat string should be constructed as follows:
	///- The first character is a delimiter character and is used to parse the remainder of the string into sub-strings.
	///- Each sub-string is separated by the delimiter character and should be in the form of a CIM property declaration (i.e., datatype and property name). This set of declarations may be used to interpret the similarly delimited RecordData property.
	///For example, using a '*' delimiter, RecordFormat = "*string ThisDay*uint32 ThisYear*datetime SomeTime"
	///may be used to interpret: RecordData = "*This is Friday*2002*20020807141000.000000-300".
	RecordFormat string
}

func NewCIM_RecordForLogEx1(instance *cim.WmiInstance) (newInstance *CIM_RecordForLog, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_RecordForLog{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewCIM_RecordForLogEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_RecordForLog, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_RecordForLog{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetLocale sets the value of Locale for the instance
func (instance *CIM_RecordForLog) SetPropertyLocale(value string) (err error) {
	return instance.SetProperty("Locale", (value))
}

// GetLocale gets the value of Locale for the instance
func (instance *CIM_RecordForLog) GetPropertyLocale() (value string, err error) {
	retValue, err := instance.GetProperty("Locale")
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

// SetRecordData sets the value of RecordData for the instance
func (instance *CIM_RecordForLog) SetPropertyRecordData(value string) (err error) {
	return instance.SetProperty("RecordData", (value))
}

// GetRecordData gets the value of RecordData for the instance
func (instance *CIM_RecordForLog) GetPropertyRecordData() (value string, err error) {
	retValue, err := instance.GetProperty("RecordData")
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

// SetRecordFormat sets the value of RecordFormat for the instance
func (instance *CIM_RecordForLog) SetPropertyRecordFormat(value string) (err error) {
	return instance.SetProperty("RecordFormat", (value))
}

// GetRecordFormat gets the value of RecordFormat for the instance
func (instance *CIM_RecordForLog) GetPropertyRecordFormat() (value string, err error) {
	retValue, err := instance.GetProperty("RecordFormat")
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
