// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS74E058DA_6E11_46EF_86F3_7C91863856BD.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEProgramSettings struct
type RSOP_IEProgramSettings struct {
	*cim.WmiInstance

	//
	calendarProgram string

	//
	checkIfIEIsDefaultBrowser bool

	//
	contactListProgram string

	//
	emailProgram string

	//
	htmlEditorHKCURegData string

	//
	htmlEditorHKLMRegData string

	//
	htmlEditorProgram string

	//
	internetCallProgram string

	//
	newsgroupsProgram string

	//
	rsopID string

	//
	rsopPrecedence uint32

	//
	useIEForFTP bool
}

func NewRSOP_IEProgramSettingsEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEProgramSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_IEProgramSettings{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_IEProgramSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEProgramSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEProgramSettings{
		WmiInstance: tmp,
	}
	return
}

// SetcalendarProgram sets the value of calendarProgram for the instance
func (instance *RSOP_IEProgramSettings) SetPropertycalendarProgram(value string) (err error) {
	return instance.SetProperty("calendarProgram", (value))
}

// GetcalendarProgram gets the value of calendarProgram for the instance
func (instance *RSOP_IEProgramSettings) GetPropertycalendarProgram() (value string, err error) {
	retValue, err := instance.GetProperty("calendarProgram")
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

// SetcheckIfIEIsDefaultBrowser sets the value of checkIfIEIsDefaultBrowser for the instance
func (instance *RSOP_IEProgramSettings) SetPropertycheckIfIEIsDefaultBrowser(value bool) (err error) {
	return instance.SetProperty("checkIfIEIsDefaultBrowser", (value))
}

// GetcheckIfIEIsDefaultBrowser gets the value of checkIfIEIsDefaultBrowser for the instance
func (instance *RSOP_IEProgramSettings) GetPropertycheckIfIEIsDefaultBrowser() (value bool, err error) {
	retValue, err := instance.GetProperty("checkIfIEIsDefaultBrowser")
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

// SetcontactListProgram sets the value of contactListProgram for the instance
func (instance *RSOP_IEProgramSettings) SetPropertycontactListProgram(value string) (err error) {
	return instance.SetProperty("contactListProgram", (value))
}

// GetcontactListProgram gets the value of contactListProgram for the instance
func (instance *RSOP_IEProgramSettings) GetPropertycontactListProgram() (value string, err error) {
	retValue, err := instance.GetProperty("contactListProgram")
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

// SetemailProgram sets the value of emailProgram for the instance
func (instance *RSOP_IEProgramSettings) SetPropertyemailProgram(value string) (err error) {
	return instance.SetProperty("emailProgram", (value))
}

// GetemailProgram gets the value of emailProgram for the instance
func (instance *RSOP_IEProgramSettings) GetPropertyemailProgram() (value string, err error) {
	retValue, err := instance.GetProperty("emailProgram")
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

// SethtmlEditorHKCURegData sets the value of htmlEditorHKCURegData for the instance
func (instance *RSOP_IEProgramSettings) SetPropertyhtmlEditorHKCURegData(value string) (err error) {
	return instance.SetProperty("htmlEditorHKCURegData", (value))
}

// GethtmlEditorHKCURegData gets the value of htmlEditorHKCURegData for the instance
func (instance *RSOP_IEProgramSettings) GetPropertyhtmlEditorHKCURegData() (value string, err error) {
	retValue, err := instance.GetProperty("htmlEditorHKCURegData")
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

// SethtmlEditorHKLMRegData sets the value of htmlEditorHKLMRegData for the instance
func (instance *RSOP_IEProgramSettings) SetPropertyhtmlEditorHKLMRegData(value string) (err error) {
	return instance.SetProperty("htmlEditorHKLMRegData", (value))
}

// GethtmlEditorHKLMRegData gets the value of htmlEditorHKLMRegData for the instance
func (instance *RSOP_IEProgramSettings) GetPropertyhtmlEditorHKLMRegData() (value string, err error) {
	retValue, err := instance.GetProperty("htmlEditorHKLMRegData")
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

// SethtmlEditorProgram sets the value of htmlEditorProgram for the instance
func (instance *RSOP_IEProgramSettings) SetPropertyhtmlEditorProgram(value string) (err error) {
	return instance.SetProperty("htmlEditorProgram", (value))
}

// GethtmlEditorProgram gets the value of htmlEditorProgram for the instance
func (instance *RSOP_IEProgramSettings) GetPropertyhtmlEditorProgram() (value string, err error) {
	retValue, err := instance.GetProperty("htmlEditorProgram")
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

// SetinternetCallProgram sets the value of internetCallProgram for the instance
func (instance *RSOP_IEProgramSettings) SetPropertyinternetCallProgram(value string) (err error) {
	return instance.SetProperty("internetCallProgram", (value))
}

// GetinternetCallProgram gets the value of internetCallProgram for the instance
func (instance *RSOP_IEProgramSettings) GetPropertyinternetCallProgram() (value string, err error) {
	retValue, err := instance.GetProperty("internetCallProgram")
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

// SetnewsgroupsProgram sets the value of newsgroupsProgram for the instance
func (instance *RSOP_IEProgramSettings) SetPropertynewsgroupsProgram(value string) (err error) {
	return instance.SetProperty("newsgroupsProgram", (value))
}

// GetnewsgroupsProgram gets the value of newsgroupsProgram for the instance
func (instance *RSOP_IEProgramSettings) GetPropertynewsgroupsProgram() (value string, err error) {
	retValue, err := instance.GetProperty("newsgroupsProgram")
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

// SetrsopID sets the value of rsopID for the instance
func (instance *RSOP_IEProgramSettings) SetPropertyrsopID(value string) (err error) {
	return instance.SetProperty("rsopID", (value))
}

// GetrsopID gets the value of rsopID for the instance
func (instance *RSOP_IEProgramSettings) GetPropertyrsopID() (value string, err error) {
	retValue, err := instance.GetProperty("rsopID")
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

// SetrsopPrecedence sets the value of rsopPrecedence for the instance
func (instance *RSOP_IEProgramSettings) SetPropertyrsopPrecedence(value uint32) (err error) {
	return instance.SetProperty("rsopPrecedence", (value))
}

// GetrsopPrecedence gets the value of rsopPrecedence for the instance
func (instance *RSOP_IEProgramSettings) GetPropertyrsopPrecedence() (value uint32, err error) {
	retValue, err := instance.GetProperty("rsopPrecedence")
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

// SetuseIEForFTP sets the value of useIEForFTP for the instance
func (instance *RSOP_IEProgramSettings) SetPropertyuseIEForFTP(value bool) (err error) {
	return instance.SetProperty("useIEForFTP", (value))
}

// GetuseIEForFTP gets the value of useIEForFTP for the instance
func (instance *RSOP_IEProgramSettings) GetPropertyuseIEForFTP() (value bool, err error) {
	retValue, err := instance.GetProperty("useIEForFTP")
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
