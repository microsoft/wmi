// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_IEToolbarButton struct
type RSOP_IEToolbarButton struct {
	cim.WmiInstance

	//
	actionPath string

	//
	buttonOrder int32

	//
	caption string

	//
	hotIconPath string

	//
	iconPath string

	//
	rsopID string

	//
	rsopPrecedence uint32

	//
	showOnToolbarByDefault bool
}

// SetactionPath sets the value of actionPath for the instance
func (instance *RSOP_IEToolbarButton) SetPropertyactionPath(value string) (err error) {
	return instance.SetProperty("actionPath", value)
}

// GetactionPath gets the value of actionPath for the instance
func (instance *RSOP_IEToolbarButton) GetPropertyactionPath() (value string, err error) {
	retValue, err := instance.GetProperty("actionPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetbuttonOrder sets the value of buttonOrder for the instance
func (instance *RSOP_IEToolbarButton) SetPropertybuttonOrder(value int32) (err error) {
	return instance.SetProperty("buttonOrder", value)
}

// GetbuttonOrder gets the value of buttonOrder for the instance
func (instance *RSOP_IEToolbarButton) GetPropertybuttonOrder() (value int32, err error) {
	retValue, err := instance.GetProperty("buttonOrder")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setcaption sets the value of caption for the instance
func (instance *RSOP_IEToolbarButton) SetPropertycaption(value string) (err error) {
	return instance.SetProperty("caption", value)
}

// Getcaption gets the value of caption for the instance
func (instance *RSOP_IEToolbarButton) GetPropertycaption() (value string, err error) {
	retValue, err := instance.GetProperty("caption")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SethotIconPath sets the value of hotIconPath for the instance
func (instance *RSOP_IEToolbarButton) SetPropertyhotIconPath(value string) (err error) {
	return instance.SetProperty("hotIconPath", value)
}

// GethotIconPath gets the value of hotIconPath for the instance
func (instance *RSOP_IEToolbarButton) GetPropertyhotIconPath() (value string, err error) {
	retValue, err := instance.GetProperty("hotIconPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SeticonPath sets the value of iconPath for the instance
func (instance *RSOP_IEToolbarButton) SetPropertyiconPath(value string) (err error) {
	return instance.SetProperty("iconPath", value)
}

// GeticonPath gets the value of iconPath for the instance
func (instance *RSOP_IEToolbarButton) GetPropertyiconPath() (value string, err error) {
	retValue, err := instance.GetProperty("iconPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetrsopID sets the value of rsopID for the instance
func (instance *RSOP_IEToolbarButton) SetPropertyrsopID(value string) (err error) {
	return instance.SetProperty("rsopID", value)
}

// GetrsopID gets the value of rsopID for the instance
func (instance *RSOP_IEToolbarButton) GetPropertyrsopID() (value string, err error) {
	retValue, err := instance.GetProperty("rsopID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetrsopPrecedence sets the value of rsopPrecedence for the instance
func (instance *RSOP_IEToolbarButton) SetPropertyrsopPrecedence(value uint32) (err error) {
	return instance.SetProperty("rsopPrecedence", value)
}

// GetrsopPrecedence gets the value of rsopPrecedence for the instance
func (instance *RSOP_IEToolbarButton) GetPropertyrsopPrecedence() (value uint32, err error) {
	retValue, err := instance.GetProperty("rsopPrecedence")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetshowOnToolbarByDefault sets the value of showOnToolbarByDefault for the instance
func (instance *RSOP_IEToolbarButton) SetPropertyshowOnToolbarByDefault(value bool) (err error) {
	return instance.SetProperty("showOnToolbarByDefault", value)
}

// GetshowOnToolbarByDefault gets the value of showOnToolbarByDefault for the instance
func (instance *RSOP_IEToolbarButton) GetPropertyshowOnToolbarByDefault() (value bool, err error) {
	retValue, err := instance.GetProperty("showOnToolbarByDefault")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
