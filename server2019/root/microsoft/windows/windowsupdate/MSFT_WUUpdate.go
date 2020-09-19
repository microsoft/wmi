// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.WindowsUpdate
//////////////////////////////////////////////
package windowsupdate

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_WUUpdate struct
type MSFT_WUUpdate struct {
	*cim.WmiInstance

	//
	Description string

	//
	KBArticleID string

	//
	MsrcSeverity string

	//
	RevisionNumber uint32

	//
	Title string

	//
	UpdateID string
}

func NewMSFT_WUUpdateEx1(instance *cim.WmiInstance) (newInstance *MSFT_WUUpdate, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_WUUpdate{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_WUUpdateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_WUUpdate, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_WUUpdate{
		WmiInstance: tmp,
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_WUUpdate) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_WUUpdate) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
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

// SetKBArticleID sets the value of KBArticleID for the instance
func (instance *MSFT_WUUpdate) SetPropertyKBArticleID(value string) (err error) {
	return instance.SetProperty("KBArticleID", (value))
}

// GetKBArticleID gets the value of KBArticleID for the instance
func (instance *MSFT_WUUpdate) GetPropertyKBArticleID() (value string, err error) {
	retValue, err := instance.GetProperty("KBArticleID")
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

// SetMsrcSeverity sets the value of MsrcSeverity for the instance
func (instance *MSFT_WUUpdate) SetPropertyMsrcSeverity(value string) (err error) {
	return instance.SetProperty("MsrcSeverity", (value))
}

// GetMsrcSeverity gets the value of MsrcSeverity for the instance
func (instance *MSFT_WUUpdate) GetPropertyMsrcSeverity() (value string, err error) {
	retValue, err := instance.GetProperty("MsrcSeverity")
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

// SetRevisionNumber sets the value of RevisionNumber for the instance
func (instance *MSFT_WUUpdate) SetPropertyRevisionNumber(value uint32) (err error) {
	return instance.SetProperty("RevisionNumber", (value))
}

// GetRevisionNumber gets the value of RevisionNumber for the instance
func (instance *MSFT_WUUpdate) GetPropertyRevisionNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("RevisionNumber")
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

// SetTitle sets the value of Title for the instance
func (instance *MSFT_WUUpdate) SetPropertyTitle(value string) (err error) {
	return instance.SetProperty("Title", (value))
}

// GetTitle gets the value of Title for the instance
func (instance *MSFT_WUUpdate) GetPropertyTitle() (value string, err error) {
	retValue, err := instance.GetProperty("Title")
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

// SetUpdateID sets the value of UpdateID for the instance
func (instance *MSFT_WUUpdate) SetPropertyUpdateID(value string) (err error) {
	return instance.SetProperty("UpdateID", (value))
}

// GetUpdateID gets the value of UpdateID for the instance
func (instance *MSFT_WUUpdate) GetPropertyUpdateID() (value string, err error) {
	retValue, err := instance.GetProperty("UpdateID")
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
