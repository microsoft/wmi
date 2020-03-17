// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// CIM_SCSIProtocolController struct
type CIM_SCSIProtocolController struct {
	CIM_ProtocolController

	// The NameFormat property identifies how the Name of the SCSIProtocolController is selected.
	///For Fibre Channel, the NameFormat is 'FC Port WWN'.
	///For iSCSI, Name can use any of the 3 iSCSI formats (iqn, eui, naa) which include the iSCSI format as as a prefix in the name, so they are not ambiguous.
	NameFormat SCSIProtocolController_NameFormat

	// A string describing how the ProtocolController is identified when the NameFormat is "Other".
	OtherNameFormat string
}

// SetNameFormat sets the value of NameFormat for the instance
func (instance *CIM_SCSIProtocolController) SetPropertyNameFormat(value SCSIProtocolController_NameFormat) (err error) {
	return instance.SetProperty("NameFormat", value)
}

// GetNameFormat gets the value of NameFormat for the instance
func (instance *CIM_SCSIProtocolController) GetPropertyNameFormat() (value SCSIProtocolController_NameFormat, err error) {
	retValue, err := instance.GetProperty("NameFormat")
	if err != nil {
		return
	}
	value, ok := retValue.(SCSIProtocolController_NameFormat)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherNameFormat sets the value of OtherNameFormat for the instance
func (instance *CIM_SCSIProtocolController) SetPropertyOtherNameFormat(value string) (err error) {
	return instance.SetProperty("OtherNameFormat", value)
}

// GetOtherNameFormat gets the value of OtherNameFormat for the instance
func (instance *CIM_SCSIProtocolController) GetPropertyOtherNameFormat() (value string, err error) {
	retValue, err := instance.GetProperty("OtherNameFormat")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
