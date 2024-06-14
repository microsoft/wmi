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

// HttpCompressionSection struct
type HttpCompressionSection struct {
	*ConfigurationSectionWithCollection

	//
	CacheControlHeader string

	//
	Directory string

	//
	DoDiskSpaceLimiting bool

	//
	DynamicCompressionBufferLimit uint32

	//
	DynamicCompressionDisableCpuUsage uint32

	//
	DynamicCompressionEnableCpuUsage uint32

	//
	DynamicTypes DynamicTypeSettings

	//
	ExpiresHeader string

	//
	HttpCompression []HttpCompressionSchemeElement

	//
	MaxDiskSpaceUsage uint32

	//
	MinFileSizeForComp uint32

	//
	NoCompressionForHttp10 bool

	//
	NoCompressionForProxies bool

	//
	NoCompressionForRange bool

	//
	SendCacheHeaders bool

	//
	StaticCompressionDisableCpuUsage uint32

	//
	StaticCompressionEnableCpuUsage uint32

	//
	StaticCompressionIgnoreHitFrequency bool

	//
	StaticTypes StaticTypeSettings
}

func NewHttpCompressionSectionEx1(instance *cim.WmiInstance) (newInstance *HttpCompressionSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HttpCompressionSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewHttpCompressionSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HttpCompressionSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HttpCompressionSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetCacheControlHeader sets the value of CacheControlHeader for the instance
func (instance *HttpCompressionSection) SetPropertyCacheControlHeader(value string) (err error) {
	return instance.SetProperty("CacheControlHeader", (value))
}

// GetCacheControlHeader gets the value of CacheControlHeader for the instance
func (instance *HttpCompressionSection) GetPropertyCacheControlHeader() (value string, err error) {
	retValue, err := instance.GetProperty("CacheControlHeader")
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
func (instance *HttpCompressionSection) SetPropertyDirectory(value string) (err error) {
	return instance.SetProperty("Directory", (value))
}

// GetDirectory gets the value of Directory for the instance
func (instance *HttpCompressionSection) GetPropertyDirectory() (value string, err error) {
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

// SetDoDiskSpaceLimiting sets the value of DoDiskSpaceLimiting for the instance
func (instance *HttpCompressionSection) SetPropertyDoDiskSpaceLimiting(value bool) (err error) {
	return instance.SetProperty("DoDiskSpaceLimiting", (value))
}

// GetDoDiskSpaceLimiting gets the value of DoDiskSpaceLimiting for the instance
func (instance *HttpCompressionSection) GetPropertyDoDiskSpaceLimiting() (value bool, err error) {
	retValue, err := instance.GetProperty("DoDiskSpaceLimiting")
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

// SetDynamicCompressionBufferLimit sets the value of DynamicCompressionBufferLimit for the instance
func (instance *HttpCompressionSection) SetPropertyDynamicCompressionBufferLimit(value uint32) (err error) {
	return instance.SetProperty("DynamicCompressionBufferLimit", (value))
}

// GetDynamicCompressionBufferLimit gets the value of DynamicCompressionBufferLimit for the instance
func (instance *HttpCompressionSection) GetPropertyDynamicCompressionBufferLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("DynamicCompressionBufferLimit")
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

// SetDynamicCompressionDisableCpuUsage sets the value of DynamicCompressionDisableCpuUsage for the instance
func (instance *HttpCompressionSection) SetPropertyDynamicCompressionDisableCpuUsage(value uint32) (err error) {
	return instance.SetProperty("DynamicCompressionDisableCpuUsage", (value))
}

// GetDynamicCompressionDisableCpuUsage gets the value of DynamicCompressionDisableCpuUsage for the instance
func (instance *HttpCompressionSection) GetPropertyDynamicCompressionDisableCpuUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("DynamicCompressionDisableCpuUsage")
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

// SetDynamicCompressionEnableCpuUsage sets the value of DynamicCompressionEnableCpuUsage for the instance
func (instance *HttpCompressionSection) SetPropertyDynamicCompressionEnableCpuUsage(value uint32) (err error) {
	return instance.SetProperty("DynamicCompressionEnableCpuUsage", (value))
}

// GetDynamicCompressionEnableCpuUsage gets the value of DynamicCompressionEnableCpuUsage for the instance
func (instance *HttpCompressionSection) GetPropertyDynamicCompressionEnableCpuUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("DynamicCompressionEnableCpuUsage")
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

// SetDynamicTypes sets the value of DynamicTypes for the instance
func (instance *HttpCompressionSection) SetPropertyDynamicTypes(value DynamicTypeSettings) (err error) {
	return instance.SetProperty("DynamicTypes", (value))
}

// GetDynamicTypes gets the value of DynamicTypes for the instance
func (instance *HttpCompressionSection) GetPropertyDynamicTypes() (value DynamicTypeSettings, err error) {
	retValue, err := instance.GetProperty("DynamicTypes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(DynamicTypeSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " DynamicTypeSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = DynamicTypeSettings(valuetmp)

	return
}

// SetExpiresHeader sets the value of ExpiresHeader for the instance
func (instance *HttpCompressionSection) SetPropertyExpiresHeader(value string) (err error) {
	return instance.SetProperty("ExpiresHeader", (value))
}

// GetExpiresHeader gets the value of ExpiresHeader for the instance
func (instance *HttpCompressionSection) GetPropertyExpiresHeader() (value string, err error) {
	retValue, err := instance.GetProperty("ExpiresHeader")
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

// SetHttpCompression sets the value of HttpCompression for the instance
func (instance *HttpCompressionSection) SetPropertyHttpCompression(value []HttpCompressionSchemeElement) (err error) {
	return instance.SetProperty("HttpCompression", (value))
}

// GetHttpCompression gets the value of HttpCompression for the instance
func (instance *HttpCompressionSection) GetPropertyHttpCompression() (value []HttpCompressionSchemeElement, err error) {
	retValue, err := instance.GetProperty("HttpCompression")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(HttpCompressionSchemeElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " HttpCompressionSchemeElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, HttpCompressionSchemeElement(valuetmp))
	}

	return
}

// SetMaxDiskSpaceUsage sets the value of MaxDiskSpaceUsage for the instance
func (instance *HttpCompressionSection) SetPropertyMaxDiskSpaceUsage(value uint32) (err error) {
	return instance.SetProperty("MaxDiskSpaceUsage", (value))
}

// GetMaxDiskSpaceUsage gets the value of MaxDiskSpaceUsage for the instance
func (instance *HttpCompressionSection) GetPropertyMaxDiskSpaceUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxDiskSpaceUsage")
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

// SetMinFileSizeForComp sets the value of MinFileSizeForComp for the instance
func (instance *HttpCompressionSection) SetPropertyMinFileSizeForComp(value uint32) (err error) {
	return instance.SetProperty("MinFileSizeForComp", (value))
}

// GetMinFileSizeForComp gets the value of MinFileSizeForComp for the instance
func (instance *HttpCompressionSection) GetPropertyMinFileSizeForComp() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinFileSizeForComp")
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

// SetNoCompressionForHttp10 sets the value of NoCompressionForHttp10 for the instance
func (instance *HttpCompressionSection) SetPropertyNoCompressionForHttp10(value bool) (err error) {
	return instance.SetProperty("NoCompressionForHttp10", (value))
}

// GetNoCompressionForHttp10 gets the value of NoCompressionForHttp10 for the instance
func (instance *HttpCompressionSection) GetPropertyNoCompressionForHttp10() (value bool, err error) {
	retValue, err := instance.GetProperty("NoCompressionForHttp10")
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

// SetNoCompressionForProxies sets the value of NoCompressionForProxies for the instance
func (instance *HttpCompressionSection) SetPropertyNoCompressionForProxies(value bool) (err error) {
	return instance.SetProperty("NoCompressionForProxies", (value))
}

// GetNoCompressionForProxies gets the value of NoCompressionForProxies for the instance
func (instance *HttpCompressionSection) GetPropertyNoCompressionForProxies() (value bool, err error) {
	retValue, err := instance.GetProperty("NoCompressionForProxies")
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

// SetNoCompressionForRange sets the value of NoCompressionForRange for the instance
func (instance *HttpCompressionSection) SetPropertyNoCompressionForRange(value bool) (err error) {
	return instance.SetProperty("NoCompressionForRange", (value))
}

// GetNoCompressionForRange gets the value of NoCompressionForRange for the instance
func (instance *HttpCompressionSection) GetPropertyNoCompressionForRange() (value bool, err error) {
	retValue, err := instance.GetProperty("NoCompressionForRange")
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

// SetSendCacheHeaders sets the value of SendCacheHeaders for the instance
func (instance *HttpCompressionSection) SetPropertySendCacheHeaders(value bool) (err error) {
	return instance.SetProperty("SendCacheHeaders", (value))
}

// GetSendCacheHeaders gets the value of SendCacheHeaders for the instance
func (instance *HttpCompressionSection) GetPropertySendCacheHeaders() (value bool, err error) {
	retValue, err := instance.GetProperty("SendCacheHeaders")
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

// SetStaticCompressionDisableCpuUsage sets the value of StaticCompressionDisableCpuUsage for the instance
func (instance *HttpCompressionSection) SetPropertyStaticCompressionDisableCpuUsage(value uint32) (err error) {
	return instance.SetProperty("StaticCompressionDisableCpuUsage", (value))
}

// GetStaticCompressionDisableCpuUsage gets the value of StaticCompressionDisableCpuUsage for the instance
func (instance *HttpCompressionSection) GetPropertyStaticCompressionDisableCpuUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("StaticCompressionDisableCpuUsage")
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

// SetStaticCompressionEnableCpuUsage sets the value of StaticCompressionEnableCpuUsage for the instance
func (instance *HttpCompressionSection) SetPropertyStaticCompressionEnableCpuUsage(value uint32) (err error) {
	return instance.SetProperty("StaticCompressionEnableCpuUsage", (value))
}

// GetStaticCompressionEnableCpuUsage gets the value of StaticCompressionEnableCpuUsage for the instance
func (instance *HttpCompressionSection) GetPropertyStaticCompressionEnableCpuUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("StaticCompressionEnableCpuUsage")
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

// SetStaticCompressionIgnoreHitFrequency sets the value of StaticCompressionIgnoreHitFrequency for the instance
func (instance *HttpCompressionSection) SetPropertyStaticCompressionIgnoreHitFrequency(value bool) (err error) {
	return instance.SetProperty("StaticCompressionIgnoreHitFrequency", (value))
}

// GetStaticCompressionIgnoreHitFrequency gets the value of StaticCompressionIgnoreHitFrequency for the instance
func (instance *HttpCompressionSection) GetPropertyStaticCompressionIgnoreHitFrequency() (value bool, err error) {
	retValue, err := instance.GetProperty("StaticCompressionIgnoreHitFrequency")
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

// SetStaticTypes sets the value of StaticTypes for the instance
func (instance *HttpCompressionSection) SetPropertyStaticTypes(value StaticTypeSettings) (err error) {
	return instance.SetProperty("StaticTypes", (value))
}

// GetStaticTypes gets the value of StaticTypes for the instance
func (instance *HttpCompressionSection) GetPropertyStaticTypes() (value StaticTypeSettings, err error) {
	retValue, err := instance.GetProperty("StaticTypes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(StaticTypeSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " StaticTypeSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = StaticTypeSettings(valuetmp)

	return
}
