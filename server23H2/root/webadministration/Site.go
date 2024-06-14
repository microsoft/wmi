// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Site struct
type Site struct {
	*ConfiguredObject

	//
	ApplicationDefaults ApplicationElementDefaults

	//
	Bindings []BindingElement

	//
	FtpServer FtpServerSettings

	//
	Id uint32

	//
	Limits SiteLimits

	//
	LogFile SiteLogFile

	//
	Name string

	//
	ServerAutoStart bool

	//
	TraceFailedRequestsLogging TraceFailedRequestsLogging

	//
	VirtualDirectoryDefaults VirtualDirectoryElementDefaults
}

func NewSiteEx1(instance *cim.WmiInstance) (newInstance *Site, err error) {
	tmp, err := NewConfiguredObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Site{
		ConfiguredObject: tmp,
	}
	return
}

func NewSiteEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Site, err error) {
	tmp, err := NewConfiguredObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Site{
		ConfiguredObject: tmp,
	}
	return
}

// SetApplicationDefaults sets the value of ApplicationDefaults for the instance
func (instance *Site) SetPropertyApplicationDefaults(value ApplicationElementDefaults) (err error) {
	return instance.SetProperty("ApplicationDefaults", (value))
}

// GetApplicationDefaults gets the value of ApplicationDefaults for the instance
func (instance *Site) GetPropertyApplicationDefaults() (value ApplicationElementDefaults, err error) {
	retValue, err := instance.GetProperty("ApplicationDefaults")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ApplicationElementDefaults)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ApplicationElementDefaults is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ApplicationElementDefaults(valuetmp)

	return
}

// SetBindings sets the value of Bindings for the instance
func (instance *Site) SetPropertyBindings(value []BindingElement) (err error) {
	return instance.SetProperty("Bindings", (value))
}

// GetBindings gets the value of Bindings for the instance
func (instance *Site) GetPropertyBindings() (value []BindingElement, err error) {
	retValue, err := instance.GetProperty("Bindings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(BindingElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " BindingElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, BindingElement(valuetmp))
	}

	return
}

// SetFtpServer sets the value of FtpServer for the instance
func (instance *Site) SetPropertyFtpServer(value FtpServerSettings) (err error) {
	return instance.SetProperty("FtpServer", (value))
}

// GetFtpServer gets the value of FtpServer for the instance
func (instance *Site) GetPropertyFtpServer() (value FtpServerSettings, err error) {
	retValue, err := instance.GetProperty("FtpServer")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpServerSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpServerSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpServerSettings(valuetmp)

	return
}

// SetId sets the value of Id for the instance
func (instance *Site) SetPropertyId(value uint32) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *Site) GetPropertyId() (value uint32, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetLimits sets the value of Limits for the instance
func (instance *Site) SetPropertyLimits(value SiteLimits) (err error) {
	return instance.SetProperty("Limits", (value))
}

// GetLimits gets the value of Limits for the instance
func (instance *Site) GetPropertyLimits() (value SiteLimits, err error) {
	retValue, err := instance.GetProperty("Limits")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(SiteLimits)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " SiteLimits is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = SiteLimits(valuetmp)

	return
}

// SetLogFile sets the value of LogFile for the instance
func (instance *Site) SetPropertyLogFile(value SiteLogFile) (err error) {
	return instance.SetProperty("LogFile", (value))
}

// GetLogFile gets the value of LogFile for the instance
func (instance *Site) GetPropertyLogFile() (value SiteLogFile, err error) {
	retValue, err := instance.GetProperty("LogFile")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(SiteLogFile)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " SiteLogFile is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = SiteLogFile(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *Site) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *Site) GetPropertyName() (value string, err error) {
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

// SetServerAutoStart sets the value of ServerAutoStart for the instance
func (instance *Site) SetPropertyServerAutoStart(value bool) (err error) {
	return instance.SetProperty("ServerAutoStart", (value))
}

// GetServerAutoStart gets the value of ServerAutoStart for the instance
func (instance *Site) GetPropertyServerAutoStart() (value bool, err error) {
	retValue, err := instance.GetProperty("ServerAutoStart")
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

// SetTraceFailedRequestsLogging sets the value of TraceFailedRequestsLogging for the instance
func (instance *Site) SetPropertyTraceFailedRequestsLogging(value TraceFailedRequestsLogging) (err error) {
	return instance.SetProperty("TraceFailedRequestsLogging", (value))
}

// GetTraceFailedRequestsLogging gets the value of TraceFailedRequestsLogging for the instance
func (instance *Site) GetPropertyTraceFailedRequestsLogging() (value TraceFailedRequestsLogging, err error) {
	retValue, err := instance.GetProperty("TraceFailedRequestsLogging")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(TraceFailedRequestsLogging)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " TraceFailedRequestsLogging is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = TraceFailedRequestsLogging(valuetmp)

	return
}

// SetVirtualDirectoryDefaults sets the value of VirtualDirectoryDefaults for the instance
func (instance *Site) SetPropertyVirtualDirectoryDefaults(value VirtualDirectoryElementDefaults) (err error) {
	return instance.SetProperty("VirtualDirectoryDefaults", (value))
}

// GetVirtualDirectoryDefaults gets the value of VirtualDirectoryDefaults for the instance
func (instance *Site) GetPropertyVirtualDirectoryDefaults() (value VirtualDirectoryElementDefaults, err error) {
	retValue, err := instance.GetProperty("VirtualDirectoryDefaults")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(VirtualDirectoryElementDefaults)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " VirtualDirectoryElementDefaults is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VirtualDirectoryElementDefaults(valuetmp)

	return
}

//

// <param name="PropertyName" type="string "></param>
func (instance *Site) RevertToParent( /* OPTIONAL IN */ PropertyName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("RevertToParent", PropertyName)
	if err != nil {
		return
	}
	return

}

//

// <param name="ProtocolName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Site) GetState( /* OPTIONAL IN */ ProtocolName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("GetState", ProtocolName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ProtocolName" type="string "></param>
func (instance *Site) Start( /* OPTIONAL IN */ ProtocolName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("Start", ProtocolName)
	if err != nil {
		return
	}
	return

}

//

// <param name="ProtocolName" type="string "></param>
func (instance *Site) Stop( /* OPTIONAL IN */ ProtocolName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("Stop", ProtocolName)
	if err != nil {
		return
	}
	return

}

//

// <param name="Bindings" type="BindingElement []"></param>
// <param name="Name" type="string "></param>
// <param name="PhysicalPath" type="string "></param>
// <param name="ServerAutoStart" type="bool "></param>
func (instance *Site) Create( /* IN */ Name string,
	/* IN */ Bindings []BindingElement,
	/* OPTIONAL IN */ PhysicalPath string,
	/* OPTIONAL IN */ ServerAutoStart bool) (err error) {
	_, err = instance.InvokeMethodWithReturn("Create", Name, Bindings, PhysicalPath, ServerAutoStart)
	if err != nil {
		return
	}
	return

}
