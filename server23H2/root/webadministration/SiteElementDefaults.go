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

// SiteElementDefaults struct
type SiteElementDefaults struct {
	*EmbeddedObject

	//
	FtpServer FtpServerSettings

	//
	Limits SiteLimits

	//
	LogFile SiteLogFile

	//
	ServerAutoStart bool

	//
	TraceFailedRequestsLogging TraceFailedRequestsLogging
}

func NewSiteElementDefaultsEx1(instance *cim.WmiInstance) (newInstance *SiteElementDefaults, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SiteElementDefaults{
		EmbeddedObject: tmp,
	}
	return
}

func NewSiteElementDefaultsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SiteElementDefaults, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SiteElementDefaults{
		EmbeddedObject: tmp,
	}
	return
}

// SetFtpServer sets the value of FtpServer for the instance
func (instance *SiteElementDefaults) SetPropertyFtpServer(value FtpServerSettings) (err error) {
	return instance.SetProperty("FtpServer", (value))
}

// GetFtpServer gets the value of FtpServer for the instance
func (instance *SiteElementDefaults) GetPropertyFtpServer() (value FtpServerSettings, err error) {
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

// SetLimits sets the value of Limits for the instance
func (instance *SiteElementDefaults) SetPropertyLimits(value SiteLimits) (err error) {
	return instance.SetProperty("Limits", (value))
}

// GetLimits gets the value of Limits for the instance
func (instance *SiteElementDefaults) GetPropertyLimits() (value SiteLimits, err error) {
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
func (instance *SiteElementDefaults) SetPropertyLogFile(value SiteLogFile) (err error) {
	return instance.SetProperty("LogFile", (value))
}

// GetLogFile gets the value of LogFile for the instance
func (instance *SiteElementDefaults) GetPropertyLogFile() (value SiteLogFile, err error) {
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

// SetServerAutoStart sets the value of ServerAutoStart for the instance
func (instance *SiteElementDefaults) SetPropertyServerAutoStart(value bool) (err error) {
	return instance.SetProperty("ServerAutoStart", (value))
}

// GetServerAutoStart gets the value of ServerAutoStart for the instance
func (instance *SiteElementDefaults) GetPropertyServerAutoStart() (value bool, err error) {
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
func (instance *SiteElementDefaults) SetPropertyTraceFailedRequestsLogging(value TraceFailedRequestsLogging) (err error) {
	return instance.SetProperty("TraceFailedRequestsLogging", (value))
}

// GetTraceFailedRequestsLogging gets the value of TraceFailedRequestsLogging for the instance
func (instance *SiteElementDefaults) GetPropertyTraceFailedRequestsLogging() (value TraceFailedRequestsLogging, err error) {
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
