// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_iSCSITargetPortal struct
type MSFT_iSCSITargetPortal struct {
	cim.WmiInstance

	//
	InitiatorInstanceName string

	//
	InitiatorPortalAddress string

	//
	IsDataDigest bool

	//
	IsHeaderDigest bool

	//
	TargetPortalAddress string

	//
	TargetPortalPortNumber uint16
}

// SetInitiatorInstanceName sets the value of InitiatorInstanceName for the instance
func (instance *MSFT_iSCSITargetPortal) SetPropertyInitiatorInstanceName(value string) (err error) {
	return instance.SetProperty("InitiatorInstanceName", value)
}

// GetInitiatorInstanceName gets the value of InitiatorInstanceName for the instance
func (instance *MSFT_iSCSITargetPortal) GetPropertyInitiatorInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InitiatorInstanceName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInitiatorPortalAddress sets the value of InitiatorPortalAddress for the instance
func (instance *MSFT_iSCSITargetPortal) SetPropertyInitiatorPortalAddress(value string) (err error) {
	return instance.SetProperty("InitiatorPortalAddress", value)
}

// GetInitiatorPortalAddress gets the value of InitiatorPortalAddress for the instance
func (instance *MSFT_iSCSITargetPortal) GetPropertyInitiatorPortalAddress() (value string, err error) {
	retValue, err := instance.GetProperty("InitiatorPortalAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsDataDigest sets the value of IsDataDigest for the instance
func (instance *MSFT_iSCSITargetPortal) SetPropertyIsDataDigest(value bool) (err error) {
	return instance.SetProperty("IsDataDigest", value)
}

// GetIsDataDigest gets the value of IsDataDigest for the instance
func (instance *MSFT_iSCSITargetPortal) GetPropertyIsDataDigest() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDataDigest")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsHeaderDigest sets the value of IsHeaderDigest for the instance
func (instance *MSFT_iSCSITargetPortal) SetPropertyIsHeaderDigest(value bool) (err error) {
	return instance.SetProperty("IsHeaderDigest", value)
}

// GetIsHeaderDigest gets the value of IsHeaderDigest for the instance
func (instance *MSFT_iSCSITargetPortal) GetPropertyIsHeaderDigest() (value bool, err error) {
	retValue, err := instance.GetProperty("IsHeaderDigest")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTargetPortalAddress sets the value of TargetPortalAddress for the instance
func (instance *MSFT_iSCSITargetPortal) SetPropertyTargetPortalAddress(value string) (err error) {
	return instance.SetProperty("TargetPortalAddress", value)
}

// GetTargetPortalAddress gets the value of TargetPortalAddress for the instance
func (instance *MSFT_iSCSITargetPortal) GetPropertyTargetPortalAddress() (value string, err error) {
	retValue, err := instance.GetProperty("TargetPortalAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTargetPortalPortNumber sets the value of TargetPortalPortNumber for the instance
func (instance *MSFT_iSCSITargetPortal) SetPropertyTargetPortalPortNumber(value uint16) (err error) {
	return instance.SetProperty("TargetPortalPortNumber", value)
}

// GetTargetPortalPortNumber gets the value of TargetPortalPortNumber for the instance
func (instance *MSFT_iSCSITargetPortal) GetPropertyTargetPortalPortNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("TargetPortalPortNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="AuthenticationType" type="string "></param>
// <param name="ChapSecret" type="string "></param>
// <param name="ChapUsername" type="string "></param>
// <param name="InitiatorInstanceName" type="string "></param>
// <param name="InitiatorPortalAddress" type="string "></param>
// <param name="IsDataDigest" type="bool "></param>
// <param name="IsHeaderDigest" type="bool "></param>
// <param name="TargetPortalAddress" type="string "></param>
// <param name="TargetPortalPortNumber" type="uint16 "></param>

// <param name="CreatedTargetPortal" type="MSFT_iSCSITargetPortal "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_iSCSITargetPortal) New( /* IN */ TargetPortalAddress string,
	/* IN */ TargetPortalPortNumber uint16,
	/* IN */ InitiatorInstanceName string,
	/* IN */ InitiatorPortalAddress string,
	/* IN */ AuthenticationType string,
	/* IN */ ChapUsername string,
	/* IN */ ChapSecret string,
	/* IN */ IsHeaderDigest bool,
	/* IN */ IsDataDigest bool,
	/* OUT */ CreatedTargetPortal MSFT_iSCSITargetPortal) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("New", TargetPortalAddress, TargetPortalPortNumber, InitiatorInstanceName, InitiatorPortalAddress, AuthenticationType, ChapUsername, ChapSecret, IsHeaderDigest, IsDataDigest)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="InitiatorInstanceName" type="string "></param>
// <param name="InitiatorPortalAddress" type="string "></param>
// <param name="TargetPortalAddress" type="string "></param>
// <param name="TargetPortalPortNumber" type="uint16 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_iSCSITargetPortal) Remove( /* IN */ InitiatorInstanceName string,
	/* IN */ InitiatorPortalAddress string,
	/* IN */ TargetPortalPortNumber uint16,
	/* IN */ TargetPortalAddress string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Remove", InitiatorInstanceName, InitiatorPortalAddress, TargetPortalPortNumber, TargetPortalAddress)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="InitiatorInstanceName" type="string "></param>
// <param name="InitiatorPortalAddress" type="string "></param>
// <param name="TargetPortalAddress" type="string "></param>
// <param name="TargetPortalPortNumber" type="uint16 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_iSCSITargetPortal) Update( /* IN */ InitiatorInstanceName string,
	/* IN */ InitiatorPortalAddress string,
	/* IN */ TargetPortalAddress string,
	/* IN */ TargetPortalPortNumber uint16) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Update", InitiatorInstanceName, InitiatorPortalAddress, TargetPortalAddress, TargetPortalPortNumber)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
