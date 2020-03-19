// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_RemoteServiceAccessPoint struct
type CIM_RemoteServiceAccessPoint struct {
	*CIM_ServiceAccessPoint

	// 733
	AccessContext RemoteServiceAccessPoint_AccessContext

	// 707
	AccessInfo string

	// 708
	InfoFormat RemoteServiceAccessPoint_InfoFormat

	// 743
	OtherAccessContext string

	// 732
	OtherInfoFormatDescription string
}

func NewCIM_RemoteServiceAccessPointEx1(instance *cim.WmiInstance) (newInstance *CIM_RemoteServiceAccessPoint, err error) {
	tmp, err := NewCIM_ServiceAccessPointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_RemoteServiceAccessPoint{
		CIM_ServiceAccessPoint: tmp,
	}
	return
}

func NewCIM_RemoteServiceAccessPointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_RemoteServiceAccessPoint, err error) {
	tmp, err := NewCIM_ServiceAccessPointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_RemoteServiceAccessPoint{
		CIM_ServiceAccessPoint: tmp,
	}
	return
}

// SetAccessContext sets the value of AccessContext for the instance
func (instance *CIM_RemoteServiceAccessPoint) SetPropertyAccessContext(value RemoteServiceAccessPoint_AccessContext) (err error) {
	return instance.SetProperty("AccessContext", value)
}

// GetAccessContext gets the value of AccessContext for the instance
func (instance *CIM_RemoteServiceAccessPoint) GetPropertyAccessContext() (value RemoteServiceAccessPoint_AccessContext, err error) {
	retValue, err := instance.GetProperty("AccessContext")
	if err != nil {
		return
	}
	value, ok := retValue.(RemoteServiceAccessPoint_AccessContext)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAccessInfo sets the value of AccessInfo for the instance
func (instance *CIM_RemoteServiceAccessPoint) SetPropertyAccessInfo(value string) (err error) {
	return instance.SetProperty("AccessInfo", value)
}

// GetAccessInfo gets the value of AccessInfo for the instance
func (instance *CIM_RemoteServiceAccessPoint) GetPropertyAccessInfo() (value string, err error) {
	retValue, err := instance.GetProperty("AccessInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInfoFormat sets the value of InfoFormat for the instance
func (instance *CIM_RemoteServiceAccessPoint) SetPropertyInfoFormat(value RemoteServiceAccessPoint_InfoFormat) (err error) {
	return instance.SetProperty("InfoFormat", value)
}

// GetInfoFormat gets the value of InfoFormat for the instance
func (instance *CIM_RemoteServiceAccessPoint) GetPropertyInfoFormat() (value RemoteServiceAccessPoint_InfoFormat, err error) {
	retValue, err := instance.GetProperty("InfoFormat")
	if err != nil {
		return
	}
	value, ok := retValue.(RemoteServiceAccessPoint_InfoFormat)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherAccessContext sets the value of OtherAccessContext for the instance
func (instance *CIM_RemoteServiceAccessPoint) SetPropertyOtherAccessContext(value string) (err error) {
	return instance.SetProperty("OtherAccessContext", value)
}

// GetOtherAccessContext gets the value of OtherAccessContext for the instance
func (instance *CIM_RemoteServiceAccessPoint) GetPropertyOtherAccessContext() (value string, err error) {
	retValue, err := instance.GetProperty("OtherAccessContext")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherInfoFormatDescription sets the value of OtherInfoFormatDescription for the instance
func (instance *CIM_RemoteServiceAccessPoint) SetPropertyOtherInfoFormatDescription(value string) (err error) {
	return instance.SetProperty("OtherInfoFormatDescription", value)
}

// GetOtherInfoFormatDescription gets the value of OtherInfoFormatDescription for the instance
func (instance *CIM_RemoteServiceAccessPoint) GetPropertyOtherInfoFormatDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherInfoFormatDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
