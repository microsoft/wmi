// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_SCSIProtocolController struct
type CIM_SCSIProtocolController struct {
	*CIM_ProtocolController

	// The NameFormat property identifies how the Name of the SCSIProtocolController is selected.
	///For Fibre Channel, the NameFormat is 'FC Port WWN'.
	///For iSCSI, Name can use any of the 3 iSCSI formats (iqn, eui, naa) which include the iSCSI format as as a prefix in the name, so they are not ambiguous.
	NameFormat SCSIProtocolController_NameFormat

	// A string describing how the ProtocolController is identified when the NameFormat is "Other".
	OtherNameFormat string
}

func NewCIM_SCSIProtocolControllerEx1(instance *cim.WmiInstance) (newInstance *CIM_SCSIProtocolController, err error) {
	tmp, err := NewCIM_ProtocolControllerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_SCSIProtocolController{
		CIM_ProtocolController: tmp,
	}
	return
}

func NewCIM_SCSIProtocolControllerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_SCSIProtocolController, err error) {
	tmp, err := NewCIM_ProtocolControllerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_SCSIProtocolController{
		CIM_ProtocolController: tmp,
	}
	return
}

// SetNameFormat sets the value of NameFormat for the instance
func (instance *CIM_SCSIProtocolController) SetPropertyNameFormat(value SCSIProtocolController_NameFormat) (err error) {
	return instance.SetProperty("NameFormat", (value))
}

// GetNameFormat gets the value of NameFormat for the instance
func (instance *CIM_SCSIProtocolController) GetPropertyNameFormat() (value SCSIProtocolController_NameFormat, err error) {
	retValue, err := instance.GetProperty("NameFormat")
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

	value = SCSIProtocolController_NameFormat(valuetmp)

	return
}

// SetOtherNameFormat sets the value of OtherNameFormat for the instance
func (instance *CIM_SCSIProtocolController) SetPropertyOtherNameFormat(value string) (err error) {
	return instance.SetProperty("OtherNameFormat", (value))
}

// GetOtherNameFormat gets the value of OtherNameFormat for the instance
func (instance *CIM_SCSIProtocolController) GetPropertyOtherNameFormat() (value string, err error) {
	retValue, err := instance.GetProperty("OtherNameFormat")
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
