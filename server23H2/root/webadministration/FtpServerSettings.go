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

// FtpServerSettings struct
type FtpServerSettings struct {
	*EmbeddedObject

	//
	AllowUTF8 bool

	//
	Connections FtpConnections

	//
	CustomFeatures FtpCustomFeatures

	//
	DirectoryBrowse FtpDirectoryBrowse

	//
	FileHandling FtpFileHandling

	//
	FirewallSupport FtpFirewallSupport

	//
	LastStartupStatus uint32

	//
	LogFile FtpLogFile

	//
	Messages FtpMessages

	//
	Security FtpSecurity

	//
	ServerAutoStart bool

	//
	Sessions FtpSessionsSettings

	//
	UserIsolation FtpUserIsolation
}

func NewFtpServerSettingsEx1(instance *cim.WmiInstance) (newInstance *FtpServerSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpServerSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpServerSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpServerSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpServerSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetAllowUTF8 sets the value of AllowUTF8 for the instance
func (instance *FtpServerSettings) SetPropertyAllowUTF8(value bool) (err error) {
	return instance.SetProperty("AllowUTF8", (value))
}

// GetAllowUTF8 gets the value of AllowUTF8 for the instance
func (instance *FtpServerSettings) GetPropertyAllowUTF8() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowUTF8")
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

// SetConnections sets the value of Connections for the instance
func (instance *FtpServerSettings) SetPropertyConnections(value FtpConnections) (err error) {
	return instance.SetProperty("Connections", (value))
}

// GetConnections gets the value of Connections for the instance
func (instance *FtpServerSettings) GetPropertyConnections() (value FtpConnections, err error) {
	retValue, err := instance.GetProperty("Connections")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpConnections)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpConnections is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpConnections(valuetmp)

	return
}

// SetCustomFeatures sets the value of CustomFeatures for the instance
func (instance *FtpServerSettings) SetPropertyCustomFeatures(value FtpCustomFeatures) (err error) {
	return instance.SetProperty("CustomFeatures", (value))
}

// GetCustomFeatures gets the value of CustomFeatures for the instance
func (instance *FtpServerSettings) GetPropertyCustomFeatures() (value FtpCustomFeatures, err error) {
	retValue, err := instance.GetProperty("CustomFeatures")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpCustomFeatures)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpCustomFeatures is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpCustomFeatures(valuetmp)

	return
}

// SetDirectoryBrowse sets the value of DirectoryBrowse for the instance
func (instance *FtpServerSettings) SetPropertyDirectoryBrowse(value FtpDirectoryBrowse) (err error) {
	return instance.SetProperty("DirectoryBrowse", (value))
}

// GetDirectoryBrowse gets the value of DirectoryBrowse for the instance
func (instance *FtpServerSettings) GetPropertyDirectoryBrowse() (value FtpDirectoryBrowse, err error) {
	retValue, err := instance.GetProperty("DirectoryBrowse")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpDirectoryBrowse)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpDirectoryBrowse is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpDirectoryBrowse(valuetmp)

	return
}

// SetFileHandling sets the value of FileHandling for the instance
func (instance *FtpServerSettings) SetPropertyFileHandling(value FtpFileHandling) (err error) {
	return instance.SetProperty("FileHandling", (value))
}

// GetFileHandling gets the value of FileHandling for the instance
func (instance *FtpServerSettings) GetPropertyFileHandling() (value FtpFileHandling, err error) {
	retValue, err := instance.GetProperty("FileHandling")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpFileHandling)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpFileHandling is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpFileHandling(valuetmp)

	return
}

// SetFirewallSupport sets the value of FirewallSupport for the instance
func (instance *FtpServerSettings) SetPropertyFirewallSupport(value FtpFirewallSupport) (err error) {
	return instance.SetProperty("FirewallSupport", (value))
}

// GetFirewallSupport gets the value of FirewallSupport for the instance
func (instance *FtpServerSettings) GetPropertyFirewallSupport() (value FtpFirewallSupport, err error) {
	retValue, err := instance.GetProperty("FirewallSupport")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpFirewallSupport)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpFirewallSupport is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpFirewallSupport(valuetmp)

	return
}

// SetLastStartupStatus sets the value of LastStartupStatus for the instance
func (instance *FtpServerSettings) SetPropertyLastStartupStatus(value uint32) (err error) {
	return instance.SetProperty("LastStartupStatus", (value))
}

// GetLastStartupStatus gets the value of LastStartupStatus for the instance
func (instance *FtpServerSettings) GetPropertyLastStartupStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastStartupStatus")
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

// SetLogFile sets the value of LogFile for the instance
func (instance *FtpServerSettings) SetPropertyLogFile(value FtpLogFile) (err error) {
	return instance.SetProperty("LogFile", (value))
}

// GetLogFile gets the value of LogFile for the instance
func (instance *FtpServerSettings) GetPropertyLogFile() (value FtpLogFile, err error) {
	retValue, err := instance.GetProperty("LogFile")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpLogFile)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpLogFile is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpLogFile(valuetmp)

	return
}

// SetMessages sets the value of Messages for the instance
func (instance *FtpServerSettings) SetPropertyMessages(value FtpMessages) (err error) {
	return instance.SetProperty("Messages", (value))
}

// GetMessages gets the value of Messages for the instance
func (instance *FtpServerSettings) GetPropertyMessages() (value FtpMessages, err error) {
	retValue, err := instance.GetProperty("Messages")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpMessages)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpMessages is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpMessages(valuetmp)

	return
}

// SetSecurity sets the value of Security for the instance
func (instance *FtpServerSettings) SetPropertySecurity(value FtpSecurity) (err error) {
	return instance.SetProperty("Security", (value))
}

// GetSecurity gets the value of Security for the instance
func (instance *FtpServerSettings) GetPropertySecurity() (value FtpSecurity, err error) {
	retValue, err := instance.GetProperty("Security")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpSecurity)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpSecurity is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpSecurity(valuetmp)

	return
}

// SetServerAutoStart sets the value of ServerAutoStart for the instance
func (instance *FtpServerSettings) SetPropertyServerAutoStart(value bool) (err error) {
	return instance.SetProperty("ServerAutoStart", (value))
}

// GetServerAutoStart gets the value of ServerAutoStart for the instance
func (instance *FtpServerSettings) GetPropertyServerAutoStart() (value bool, err error) {
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

// SetSessions sets the value of Sessions for the instance
func (instance *FtpServerSettings) SetPropertySessions(value FtpSessionsSettings) (err error) {
	return instance.SetProperty("Sessions", (value))
}

// GetSessions gets the value of Sessions for the instance
func (instance *FtpServerSettings) GetPropertySessions() (value FtpSessionsSettings, err error) {
	retValue, err := instance.GetProperty("Sessions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpSessionsSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpSessionsSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpSessionsSettings(valuetmp)

	return
}

// SetUserIsolation sets the value of UserIsolation for the instance
func (instance *FtpServerSettings) SetPropertyUserIsolation(value FtpUserIsolation) (err error) {
	return instance.SetProperty("UserIsolation", (value))
}

// GetUserIsolation gets the value of UserIsolation for the instance
func (instance *FtpServerSettings) GetPropertyUserIsolation() (value FtpUserIsolation, err error) {
	retValue, err := instance.GetProperty("UserIsolation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpUserIsolation)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpUserIsolation is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpUserIsolation(valuetmp)

	return
}
