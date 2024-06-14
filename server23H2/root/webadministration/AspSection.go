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

// AspSection struct
type AspSection struct {
	*ConfigurationSection

	//
	AppAllowClientDebug bool

	//
	AppAllowDebugging bool

	//
	BufferingOn bool

	//
	Cache AspCache

	//
	CalcLineNumber bool

	//
	CodePage uint32

	//
	ComPlus AspComPlus

	//
	EnableApplicationRestart bool

	//
	EnableAspHtmlFallback bool

	//
	EnableChunkedEncoding bool

	//
	EnableParentPaths bool

	//
	ErrorsToNTLog bool

	//
	ExceptionCatchEnable bool

	//
	Lcid uint32

	//
	Limits AspLimits

	//
	LogErrorRequests bool

	//
	RunOnEndAnonymously bool

	//
	ScriptErrorMessage string

	//
	ScriptErrorSentToBrowser bool

	//
	ScriptLanguage string

	//
	Session AspSession
}

func NewAspSectionEx1(instance *cim.WmiInstance) (newInstance *AspSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AspSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewAspSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AspSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AspSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetAppAllowClientDebug sets the value of AppAllowClientDebug for the instance
func (instance *AspSection) SetPropertyAppAllowClientDebug(value bool) (err error) {
	return instance.SetProperty("AppAllowClientDebug", (value))
}

// GetAppAllowClientDebug gets the value of AppAllowClientDebug for the instance
func (instance *AspSection) GetPropertyAppAllowClientDebug() (value bool, err error) {
	retValue, err := instance.GetProperty("AppAllowClientDebug")
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

// SetAppAllowDebugging sets the value of AppAllowDebugging for the instance
func (instance *AspSection) SetPropertyAppAllowDebugging(value bool) (err error) {
	return instance.SetProperty("AppAllowDebugging", (value))
}

// GetAppAllowDebugging gets the value of AppAllowDebugging for the instance
func (instance *AspSection) GetPropertyAppAllowDebugging() (value bool, err error) {
	retValue, err := instance.GetProperty("AppAllowDebugging")
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

// SetBufferingOn sets the value of BufferingOn for the instance
func (instance *AspSection) SetPropertyBufferingOn(value bool) (err error) {
	return instance.SetProperty("BufferingOn", (value))
}

// GetBufferingOn gets the value of BufferingOn for the instance
func (instance *AspSection) GetPropertyBufferingOn() (value bool, err error) {
	retValue, err := instance.GetProperty("BufferingOn")
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

// SetCache sets the value of Cache for the instance
func (instance *AspSection) SetPropertyCache(value AspCache) (err error) {
	return instance.SetProperty("Cache", (value))
}

// GetCache gets the value of Cache for the instance
func (instance *AspSection) GetPropertyCache() (value AspCache, err error) {
	retValue, err := instance.GetProperty("Cache")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(AspCache)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " AspCache is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AspCache(valuetmp)

	return
}

// SetCalcLineNumber sets the value of CalcLineNumber for the instance
func (instance *AspSection) SetPropertyCalcLineNumber(value bool) (err error) {
	return instance.SetProperty("CalcLineNumber", (value))
}

// GetCalcLineNumber gets the value of CalcLineNumber for the instance
func (instance *AspSection) GetPropertyCalcLineNumber() (value bool, err error) {
	retValue, err := instance.GetProperty("CalcLineNumber")
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

// SetCodePage sets the value of CodePage for the instance
func (instance *AspSection) SetPropertyCodePage(value uint32) (err error) {
	return instance.SetProperty("CodePage", (value))
}

// GetCodePage gets the value of CodePage for the instance
func (instance *AspSection) GetPropertyCodePage() (value uint32, err error) {
	retValue, err := instance.GetProperty("CodePage")
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

// SetComPlus sets the value of ComPlus for the instance
func (instance *AspSection) SetPropertyComPlus(value AspComPlus) (err error) {
	return instance.SetProperty("ComPlus", (value))
}

// GetComPlus gets the value of ComPlus for the instance
func (instance *AspSection) GetPropertyComPlus() (value AspComPlus, err error) {
	retValue, err := instance.GetProperty("ComPlus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(AspComPlus)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " AspComPlus is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AspComPlus(valuetmp)

	return
}

// SetEnableApplicationRestart sets the value of EnableApplicationRestart for the instance
func (instance *AspSection) SetPropertyEnableApplicationRestart(value bool) (err error) {
	return instance.SetProperty("EnableApplicationRestart", (value))
}

// GetEnableApplicationRestart gets the value of EnableApplicationRestart for the instance
func (instance *AspSection) GetPropertyEnableApplicationRestart() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableApplicationRestart")
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

// SetEnableAspHtmlFallback sets the value of EnableAspHtmlFallback for the instance
func (instance *AspSection) SetPropertyEnableAspHtmlFallback(value bool) (err error) {
	return instance.SetProperty("EnableAspHtmlFallback", (value))
}

// GetEnableAspHtmlFallback gets the value of EnableAspHtmlFallback for the instance
func (instance *AspSection) GetPropertyEnableAspHtmlFallback() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableAspHtmlFallback")
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

// SetEnableChunkedEncoding sets the value of EnableChunkedEncoding for the instance
func (instance *AspSection) SetPropertyEnableChunkedEncoding(value bool) (err error) {
	return instance.SetProperty("EnableChunkedEncoding", (value))
}

// GetEnableChunkedEncoding gets the value of EnableChunkedEncoding for the instance
func (instance *AspSection) GetPropertyEnableChunkedEncoding() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableChunkedEncoding")
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

// SetEnableParentPaths sets the value of EnableParentPaths for the instance
func (instance *AspSection) SetPropertyEnableParentPaths(value bool) (err error) {
	return instance.SetProperty("EnableParentPaths", (value))
}

// GetEnableParentPaths gets the value of EnableParentPaths for the instance
func (instance *AspSection) GetPropertyEnableParentPaths() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableParentPaths")
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

// SetErrorsToNTLog sets the value of ErrorsToNTLog for the instance
func (instance *AspSection) SetPropertyErrorsToNTLog(value bool) (err error) {
	return instance.SetProperty("ErrorsToNTLog", (value))
}

// GetErrorsToNTLog gets the value of ErrorsToNTLog for the instance
func (instance *AspSection) GetPropertyErrorsToNTLog() (value bool, err error) {
	retValue, err := instance.GetProperty("ErrorsToNTLog")
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

// SetExceptionCatchEnable sets the value of ExceptionCatchEnable for the instance
func (instance *AspSection) SetPropertyExceptionCatchEnable(value bool) (err error) {
	return instance.SetProperty("ExceptionCatchEnable", (value))
}

// GetExceptionCatchEnable gets the value of ExceptionCatchEnable for the instance
func (instance *AspSection) GetPropertyExceptionCatchEnable() (value bool, err error) {
	retValue, err := instance.GetProperty("ExceptionCatchEnable")
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

// SetLcid sets the value of Lcid for the instance
func (instance *AspSection) SetPropertyLcid(value uint32) (err error) {
	return instance.SetProperty("Lcid", (value))
}

// GetLcid gets the value of Lcid for the instance
func (instance *AspSection) GetPropertyLcid() (value uint32, err error) {
	retValue, err := instance.GetProperty("Lcid")
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
func (instance *AspSection) SetPropertyLimits(value AspLimits) (err error) {
	return instance.SetProperty("Limits", (value))
}

// GetLimits gets the value of Limits for the instance
func (instance *AspSection) GetPropertyLimits() (value AspLimits, err error) {
	retValue, err := instance.GetProperty("Limits")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(AspLimits)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " AspLimits is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AspLimits(valuetmp)

	return
}

// SetLogErrorRequests sets the value of LogErrorRequests for the instance
func (instance *AspSection) SetPropertyLogErrorRequests(value bool) (err error) {
	return instance.SetProperty("LogErrorRequests", (value))
}

// GetLogErrorRequests gets the value of LogErrorRequests for the instance
func (instance *AspSection) GetPropertyLogErrorRequests() (value bool, err error) {
	retValue, err := instance.GetProperty("LogErrorRequests")
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

// SetRunOnEndAnonymously sets the value of RunOnEndAnonymously for the instance
func (instance *AspSection) SetPropertyRunOnEndAnonymously(value bool) (err error) {
	return instance.SetProperty("RunOnEndAnonymously", (value))
}

// GetRunOnEndAnonymously gets the value of RunOnEndAnonymously for the instance
func (instance *AspSection) GetPropertyRunOnEndAnonymously() (value bool, err error) {
	retValue, err := instance.GetProperty("RunOnEndAnonymously")
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

// SetScriptErrorMessage sets the value of ScriptErrorMessage for the instance
func (instance *AspSection) SetPropertyScriptErrorMessage(value string) (err error) {
	return instance.SetProperty("ScriptErrorMessage", (value))
}

// GetScriptErrorMessage gets the value of ScriptErrorMessage for the instance
func (instance *AspSection) GetPropertyScriptErrorMessage() (value string, err error) {
	retValue, err := instance.GetProperty("ScriptErrorMessage")
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

// SetScriptErrorSentToBrowser sets the value of ScriptErrorSentToBrowser for the instance
func (instance *AspSection) SetPropertyScriptErrorSentToBrowser(value bool) (err error) {
	return instance.SetProperty("ScriptErrorSentToBrowser", (value))
}

// GetScriptErrorSentToBrowser gets the value of ScriptErrorSentToBrowser for the instance
func (instance *AspSection) GetPropertyScriptErrorSentToBrowser() (value bool, err error) {
	retValue, err := instance.GetProperty("ScriptErrorSentToBrowser")
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

// SetScriptLanguage sets the value of ScriptLanguage for the instance
func (instance *AspSection) SetPropertyScriptLanguage(value string) (err error) {
	return instance.SetProperty("ScriptLanguage", (value))
}

// GetScriptLanguage gets the value of ScriptLanguage for the instance
func (instance *AspSection) GetPropertyScriptLanguage() (value string, err error) {
	retValue, err := instance.GetProperty("ScriptLanguage")
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

// SetSession sets the value of Session for the instance
func (instance *AspSection) SetPropertySession(value AspSession) (err error) {
	return instance.SetProperty("Session", (value))
}

// GetSession gets the value of Session for the instance
func (instance *AspSection) GetPropertySession() (value AspSession, err error) {
	retValue, err := instance.GetProperty("Session")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(AspSession)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " AspSession is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AspSession(valuetmp)

	return
}
