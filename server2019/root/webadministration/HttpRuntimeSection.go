// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// HttpRuntimeSection struct
type HttpRuntimeSection struct { 
	*ConfigurationSection

	// 
	ApartmentThreading bool

	// 
	AppRequestQueueLimit int32

	// 
	AsyncPreloadMode int32

	// 
	DelayNotificationTimeout string

	// 
	Enable bool

	// 
	EnableHeaderChecking bool

	// 
	EnableKernelOutputCache bool

	// 
	EnableVersionHeader bool

	// 
	EncoderType string

	// 
	ExecutionTimeout string

	// 
	MaxQueryStringLength int32

	// 
	MaxRequestLength int32

	// 
	MaxUrlLength int32

	// 
	MaxWaitChangeNotification int32

	// 
	MinFreeThreads int32

	// 
	MinLocalRequestFreeThreads int32

	// 
	RelaxedUrlToFileSystemMapping bool

	// 
	RequestLengthDiskThreshold int32

	// 
	RequestPathInvalidCharacters string

	// 
	RequestValidationMode string

	// 
	RequestValidationType string

	// 
	RequireRootedSaveAsPath bool

	// 
	SendCacheControlHeader bool

	// 
	ShutdownTimeout string

	// 
	UseFullyQualifiedRedirectUrl bool

	// 
	WaitChangeNotification int32
}

	func NewHttpRuntimeSectionEx1(instance *cim.WmiInstance) (newInstance *HttpRuntimeSection, err error) {tmp, err := NewConfigurationSectionEx1(instance)
		
	if err != nil { return }
	newInstance = &HttpRuntimeSection {
	ConfigurationSection: tmp,
	}
	return
	}
	

	func NewHttpRuntimeSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HttpRuntimeSection, err error) {tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HttpRuntimeSection {
	ConfigurationSection: tmp,
	}
	return
	}
	

// SetApartmentThreading sets the value of ApartmentThreading for the instance
func (instance *HttpRuntimeSection) SetPropertyApartmentThreading(value bool) (err error) { 
    return instance.SetProperty("ApartmentThreading", (value))
}

// GetApartmentThreading gets the value of ApartmentThreading for the instance
func (instance *HttpRuntimeSection) GetPropertyApartmentThreading()(value bool, err error) { 
    retValue, err := instance.GetProperty("ApartmentThreading")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetAppRequestQueueLimit sets the value of AppRequestQueueLimit for the instance
func (instance *HttpRuntimeSection) SetPropertyAppRequestQueueLimit(value int32) (err error) { 
    return instance.SetProperty("AppRequestQueueLimit", (value))
}

// GetAppRequestQueueLimit gets the value of AppRequestQueueLimit for the instance
func (instance *HttpRuntimeSection) GetPropertyAppRequestQueueLimit()(value int32, err error) { 
    retValue, err := instance.GetProperty("AppRequestQueueLimit")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetAsyncPreloadMode sets the value of AsyncPreloadMode for the instance
func (instance *HttpRuntimeSection) SetPropertyAsyncPreloadMode(value int32) (err error) { 
    return instance.SetProperty("AsyncPreloadMode", (value))
}

// GetAsyncPreloadMode gets the value of AsyncPreloadMode for the instance
func (instance *HttpRuntimeSection) GetPropertyAsyncPreloadMode()(value int32, err error) { 
    retValue, err := instance.GetProperty("AsyncPreloadMode")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetDelayNotificationTimeout sets the value of DelayNotificationTimeout for the instance
func (instance *HttpRuntimeSection) SetPropertyDelayNotificationTimeout(value string) (err error) { 
    return instance.SetProperty("DelayNotificationTimeout", (value))
}

// GetDelayNotificationTimeout gets the value of DelayNotificationTimeout for the instance
func (instance *HttpRuntimeSection) GetPropertyDelayNotificationTimeout()(value string, err error) { 
    retValue, err := instance.GetProperty("DelayNotificationTimeout")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetEnable sets the value of Enable for the instance
func (instance *HttpRuntimeSection) SetPropertyEnable(value bool) (err error) { 
    return instance.SetProperty("Enable", (value))
}

// GetEnable gets the value of Enable for the instance
func (instance *HttpRuntimeSection) GetPropertyEnable()(value bool, err error) { 
    retValue, err := instance.GetProperty("Enable")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetEnableHeaderChecking sets the value of EnableHeaderChecking for the instance
func (instance *HttpRuntimeSection) SetPropertyEnableHeaderChecking(value bool) (err error) { 
    return instance.SetProperty("EnableHeaderChecking", (value))
}

// GetEnableHeaderChecking gets the value of EnableHeaderChecking for the instance
func (instance *HttpRuntimeSection) GetPropertyEnableHeaderChecking()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnableHeaderChecking")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetEnableKernelOutputCache sets the value of EnableKernelOutputCache for the instance
func (instance *HttpRuntimeSection) SetPropertyEnableKernelOutputCache(value bool) (err error) { 
    return instance.SetProperty("EnableKernelOutputCache", (value))
}

// GetEnableKernelOutputCache gets the value of EnableKernelOutputCache for the instance
func (instance *HttpRuntimeSection) GetPropertyEnableKernelOutputCache()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnableKernelOutputCache")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetEnableVersionHeader sets the value of EnableVersionHeader for the instance
func (instance *HttpRuntimeSection) SetPropertyEnableVersionHeader(value bool) (err error) { 
    return instance.SetProperty("EnableVersionHeader", (value))
}

// GetEnableVersionHeader gets the value of EnableVersionHeader for the instance
func (instance *HttpRuntimeSection) GetPropertyEnableVersionHeader()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnableVersionHeader")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetEncoderType sets the value of EncoderType for the instance
func (instance *HttpRuntimeSection) SetPropertyEncoderType(value string) (err error) { 
    return instance.SetProperty("EncoderType", (value))
}

// GetEncoderType gets the value of EncoderType for the instance
func (instance *HttpRuntimeSection) GetPropertyEncoderType()(value string, err error) { 
    retValue, err := instance.GetProperty("EncoderType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetExecutionTimeout sets the value of ExecutionTimeout for the instance
func (instance *HttpRuntimeSection) SetPropertyExecutionTimeout(value string) (err error) { 
    return instance.SetProperty("ExecutionTimeout", (value))
}

// GetExecutionTimeout gets the value of ExecutionTimeout for the instance
func (instance *HttpRuntimeSection) GetPropertyExecutionTimeout()(value string, err error) { 
    retValue, err := instance.GetProperty("ExecutionTimeout")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetMaxQueryStringLength sets the value of MaxQueryStringLength for the instance
func (instance *HttpRuntimeSection) SetPropertyMaxQueryStringLength(value int32) (err error) { 
    return instance.SetProperty("MaxQueryStringLength", (value))
}

// GetMaxQueryStringLength gets the value of MaxQueryStringLength for the instance
func (instance *HttpRuntimeSection) GetPropertyMaxQueryStringLength()(value int32, err error) { 
    retValue, err := instance.GetProperty("MaxQueryStringLength")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetMaxRequestLength sets the value of MaxRequestLength for the instance
func (instance *HttpRuntimeSection) SetPropertyMaxRequestLength(value int32) (err error) { 
    return instance.SetProperty("MaxRequestLength", (value))
}

// GetMaxRequestLength gets the value of MaxRequestLength for the instance
func (instance *HttpRuntimeSection) GetPropertyMaxRequestLength()(value int32, err error) { 
    retValue, err := instance.GetProperty("MaxRequestLength")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetMaxUrlLength sets the value of MaxUrlLength for the instance
func (instance *HttpRuntimeSection) SetPropertyMaxUrlLength(value int32) (err error) { 
    return instance.SetProperty("MaxUrlLength", (value))
}

// GetMaxUrlLength gets the value of MaxUrlLength for the instance
func (instance *HttpRuntimeSection) GetPropertyMaxUrlLength()(value int32, err error) { 
    retValue, err := instance.GetProperty("MaxUrlLength")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetMaxWaitChangeNotification sets the value of MaxWaitChangeNotification for the instance
func (instance *HttpRuntimeSection) SetPropertyMaxWaitChangeNotification(value int32) (err error) { 
    return instance.SetProperty("MaxWaitChangeNotification", (value))
}

// GetMaxWaitChangeNotification gets the value of MaxWaitChangeNotification for the instance
func (instance *HttpRuntimeSection) GetPropertyMaxWaitChangeNotification()(value int32, err error) { 
    retValue, err := instance.GetProperty("MaxWaitChangeNotification")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetMinFreeThreads sets the value of MinFreeThreads for the instance
func (instance *HttpRuntimeSection) SetPropertyMinFreeThreads(value int32) (err error) { 
    return instance.SetProperty("MinFreeThreads", (value))
}

// GetMinFreeThreads gets the value of MinFreeThreads for the instance
func (instance *HttpRuntimeSection) GetPropertyMinFreeThreads()(value int32, err error) { 
    retValue, err := instance.GetProperty("MinFreeThreads")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetMinLocalRequestFreeThreads sets the value of MinLocalRequestFreeThreads for the instance
func (instance *HttpRuntimeSection) SetPropertyMinLocalRequestFreeThreads(value int32) (err error) { 
    return instance.SetProperty("MinLocalRequestFreeThreads", (value))
}

// GetMinLocalRequestFreeThreads gets the value of MinLocalRequestFreeThreads for the instance
func (instance *HttpRuntimeSection) GetPropertyMinLocalRequestFreeThreads()(value int32, err error) { 
    retValue, err := instance.GetProperty("MinLocalRequestFreeThreads")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetRelaxedUrlToFileSystemMapping sets the value of RelaxedUrlToFileSystemMapping for the instance
func (instance *HttpRuntimeSection) SetPropertyRelaxedUrlToFileSystemMapping(value bool) (err error) { 
    return instance.SetProperty("RelaxedUrlToFileSystemMapping", (value))
}

// GetRelaxedUrlToFileSystemMapping gets the value of RelaxedUrlToFileSystemMapping for the instance
func (instance *HttpRuntimeSection) GetPropertyRelaxedUrlToFileSystemMapping()(value bool, err error) { 
    retValue, err := instance.GetProperty("RelaxedUrlToFileSystemMapping")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetRequestLengthDiskThreshold sets the value of RequestLengthDiskThreshold for the instance
func (instance *HttpRuntimeSection) SetPropertyRequestLengthDiskThreshold(value int32) (err error) { 
    return instance.SetProperty("RequestLengthDiskThreshold", (value))
}

// GetRequestLengthDiskThreshold gets the value of RequestLengthDiskThreshold for the instance
func (instance *HttpRuntimeSection) GetPropertyRequestLengthDiskThreshold()(value int32, err error) { 
    retValue, err := instance.GetProperty("RequestLengthDiskThreshold")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetRequestPathInvalidCharacters sets the value of RequestPathInvalidCharacters for the instance
func (instance *HttpRuntimeSection) SetPropertyRequestPathInvalidCharacters(value string) (err error) { 
    return instance.SetProperty("RequestPathInvalidCharacters", (value))
}

// GetRequestPathInvalidCharacters gets the value of RequestPathInvalidCharacters for the instance
func (instance *HttpRuntimeSection) GetPropertyRequestPathInvalidCharacters()(value string, err error) { 
    retValue, err := instance.GetProperty("RequestPathInvalidCharacters")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetRequestValidationMode sets the value of RequestValidationMode for the instance
func (instance *HttpRuntimeSection) SetPropertyRequestValidationMode(value string) (err error) { 
    return instance.SetProperty("RequestValidationMode", (value))
}

// GetRequestValidationMode gets the value of RequestValidationMode for the instance
func (instance *HttpRuntimeSection) GetPropertyRequestValidationMode()(value string, err error) { 
    retValue, err := instance.GetProperty("RequestValidationMode")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetRequestValidationType sets the value of RequestValidationType for the instance
func (instance *HttpRuntimeSection) SetPropertyRequestValidationType(value string) (err error) { 
    return instance.SetProperty("RequestValidationType", (value))
}

// GetRequestValidationType gets the value of RequestValidationType for the instance
func (instance *HttpRuntimeSection) GetPropertyRequestValidationType()(value string, err error) { 
    retValue, err := instance.GetProperty("RequestValidationType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetRequireRootedSaveAsPath sets the value of RequireRootedSaveAsPath for the instance
func (instance *HttpRuntimeSection) SetPropertyRequireRootedSaveAsPath(value bool) (err error) { 
    return instance.SetProperty("RequireRootedSaveAsPath", (value))
}

// GetRequireRootedSaveAsPath gets the value of RequireRootedSaveAsPath for the instance
func (instance *HttpRuntimeSection) GetPropertyRequireRootedSaveAsPath()(value bool, err error) { 
    retValue, err := instance.GetProperty("RequireRootedSaveAsPath")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetSendCacheControlHeader sets the value of SendCacheControlHeader for the instance
func (instance *HttpRuntimeSection) SetPropertySendCacheControlHeader(value bool) (err error) { 
    return instance.SetProperty("SendCacheControlHeader", (value))
}

// GetSendCacheControlHeader gets the value of SendCacheControlHeader for the instance
func (instance *HttpRuntimeSection) GetPropertySendCacheControlHeader()(value bool, err error) { 
    retValue, err := instance.GetProperty("SendCacheControlHeader")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetShutdownTimeout sets the value of ShutdownTimeout for the instance
func (instance *HttpRuntimeSection) SetPropertyShutdownTimeout(value string) (err error) { 
    return instance.SetProperty("ShutdownTimeout", (value))
}

// GetShutdownTimeout gets the value of ShutdownTimeout for the instance
func (instance *HttpRuntimeSection) GetPropertyShutdownTimeout()(value string, err error) { 
    retValue, err := instance.GetProperty("ShutdownTimeout")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetUseFullyQualifiedRedirectUrl sets the value of UseFullyQualifiedRedirectUrl for the instance
func (instance *HttpRuntimeSection) SetPropertyUseFullyQualifiedRedirectUrl(value bool) (err error) { 
    return instance.SetProperty("UseFullyQualifiedRedirectUrl", (value))
}

// GetUseFullyQualifiedRedirectUrl gets the value of UseFullyQualifiedRedirectUrl for the instance
func (instance *HttpRuntimeSection) GetPropertyUseFullyQualifiedRedirectUrl()(value bool, err error) { 
    retValue, err := instance.GetProperty("UseFullyQualifiedRedirectUrl")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetWaitChangeNotification sets the value of WaitChangeNotification for the instance
func (instance *HttpRuntimeSection) SetPropertyWaitChangeNotification(value int32) (err error) { 
    return instance.SetProperty("WaitChangeNotification", (value))
}

// GetWaitChangeNotification gets the value of WaitChangeNotification for the instance
func (instance *HttpRuntimeSection) GetPropertyWaitChangeNotification()(value int32, err error) { 
    retValue, err := instance.GetProperty("WaitChangeNotification")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

