// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SiteLogFile struct
type SiteLogFile struct {
	*EmbeddedObject

	//
	CustomLogPluginClsid string

	//
	Directory string

	//
	Enabled bool

	//
	FlushByEntryCountW3CLog uint32

	//
	LocalTimeRollover bool

	//
	LogExtFileFlags int32

	//
	LogFormat int32

	//
	LogSiteId bool

	//
	LogTargetW3C int32

	//
	MaxLogLineLength uint32

	//
	Period int32

	//
	TruncateSize string
}

func NewSiteLogFileEx1(instance *cim.WmiInstance) (newInstance *SiteLogFile, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SiteLogFile{
		EmbeddedObject: tmp,
	}
	return
}

func NewSiteLogFileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SiteLogFile, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SiteLogFile{
		EmbeddedObject: tmp,
	}
	return
}

// SetCustomLogPluginClsid sets the value of CustomLogPluginClsid for the instance
func (instance *SiteLogFile) SetPropertyCustomLogPluginClsid(value string) (err error) {
	return instance.SetProperty("CustomLogPluginClsid", (value))
}

// GetCustomLogPluginClsid gets the value of CustomLogPluginClsid for the instance
func (instance *SiteLogFile) GetPropertyCustomLogPluginClsid() (value string, err error) {
	retValue, err := instance.GetProperty("CustomLogPluginClsid")
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

// SetDirectory sets the value of Directory for the instance
func (instance *SiteLogFile) SetPropertyDirectory(value string) (err error) {
	return instance.SetProperty("Directory", (value))
}

// GetDirectory gets the value of Directory for the instance
func (instance *SiteLogFile) GetPropertyDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("Directory")
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

// SetEnabled sets the value of Enabled for the instance
func (instance *SiteLogFile) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *SiteLogFile) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
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

// SetFlushByEntryCountW3CLog sets the value of FlushByEntryCountW3CLog for the instance
func (instance *SiteLogFile) SetPropertyFlushByEntryCountW3CLog(value uint32) (err error) {
	return instance.SetProperty("FlushByEntryCountW3CLog", (value))
}

// GetFlushByEntryCountW3CLog gets the value of FlushByEntryCountW3CLog for the instance
func (instance *SiteLogFile) GetPropertyFlushByEntryCountW3CLog() (value uint32, err error) {
	retValue, err := instance.GetProperty("FlushByEntryCountW3CLog")
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

// SetLocalTimeRollover sets the value of LocalTimeRollover for the instance
func (instance *SiteLogFile) SetPropertyLocalTimeRollover(value bool) (err error) {
	return instance.SetProperty("LocalTimeRollover", (value))
}

// GetLocalTimeRollover gets the value of LocalTimeRollover for the instance
func (instance *SiteLogFile) GetPropertyLocalTimeRollover() (value bool, err error) {
	retValue, err := instance.GetProperty("LocalTimeRollover")
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

// SetLogExtFileFlags sets the value of LogExtFileFlags for the instance
func (instance *SiteLogFile) SetPropertyLogExtFileFlags(value int32) (err error) {
	return instance.SetProperty("LogExtFileFlags", (value))
}

// GetLogExtFileFlags gets the value of LogExtFileFlags for the instance
func (instance *SiteLogFile) GetPropertyLogExtFileFlags() (value int32, err error) {
	retValue, err := instance.GetProperty("LogExtFileFlags")
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

	value = int32(valuetmp)

	return
}

// SetLogFormat sets the value of LogFormat for the instance
func (instance *SiteLogFile) SetPropertyLogFormat(value int32) (err error) {
	return instance.SetProperty("LogFormat", (value))
}

// GetLogFormat gets the value of LogFormat for the instance
func (instance *SiteLogFile) GetPropertyLogFormat() (value int32, err error) {
	retValue, err := instance.GetProperty("LogFormat")
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

	value = int32(valuetmp)

	return
}

// SetLogSiteId sets the value of LogSiteId for the instance
func (instance *SiteLogFile) SetPropertyLogSiteId(value bool) (err error) {
	return instance.SetProperty("LogSiteId", (value))
}

// GetLogSiteId gets the value of LogSiteId for the instance
func (instance *SiteLogFile) GetPropertyLogSiteId() (value bool, err error) {
	retValue, err := instance.GetProperty("LogSiteId")
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

// SetLogTargetW3C sets the value of LogTargetW3C for the instance
func (instance *SiteLogFile) SetPropertyLogTargetW3C(value int32) (err error) {
	return instance.SetProperty("LogTargetW3C", (value))
}

// GetLogTargetW3C gets the value of LogTargetW3C for the instance
func (instance *SiteLogFile) GetPropertyLogTargetW3C() (value int32, err error) {
	retValue, err := instance.GetProperty("LogTargetW3C")
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

	value = int32(valuetmp)

	return
}

// SetMaxLogLineLength sets the value of MaxLogLineLength for the instance
func (instance *SiteLogFile) SetPropertyMaxLogLineLength(value uint32) (err error) {
	return instance.SetProperty("MaxLogLineLength", (value))
}

// GetMaxLogLineLength gets the value of MaxLogLineLength for the instance
func (instance *SiteLogFile) GetPropertyMaxLogLineLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxLogLineLength")
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

// SetPeriod sets the value of Period for the instance
func (instance *SiteLogFile) SetPropertyPeriod(value int32) (err error) {
	return instance.SetProperty("Period", (value))
}

// GetPeriod gets the value of Period for the instance
func (instance *SiteLogFile) GetPropertyPeriod() (value int32, err error) {
	retValue, err := instance.GetProperty("Period")
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

	value = int32(valuetmp)

	return
}

// SetTruncateSize sets the value of TruncateSize for the instance
func (instance *SiteLogFile) SetPropertyTruncateSize(value string) (err error) {
	return instance.SetProperty("TruncateSize", (value))
}

// GetTruncateSize gets the value of TruncateSize for the instance
func (instance *SiteLogFile) GetPropertyTruncateSize() (value string, err error) {
	retValue, err := instance.GetProperty("TruncateSize")
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
