// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// Win32_PerfFormattedData_Counters_DCLocatorNetlogon struct
type Win32_PerfFormattedData_Counters_DCLocatorNetlogon struct { 
	*Win32_PerfFormattedData

	// 
	CacheHitsPersec uint32

	// 
	CacheMissesPersec uint32

	// 
	DNSQueryFailuresPersec uint32

	// 
	FlagsDSRequiredW2KRequestsPersec uint32

	// 
	FlagsDSRequiredWS2008RequestsPersec uint32

	// 
	FlagsDSRequiredWS2012R2RequestsPersec uint32

	// 
	FlagsDSRequiredWS2012RequestsPersec uint32

	// 
	FlagsDSRequiredWS2016RequestsPersec uint32

	// 
	FlagsForceRediscoveryRequestsPersec uint32

	// 
	FlagsGCRequiredRequestsPersec uint32

	// 
	FlagsKDCRequiredRequestsPersec uint32

	// 
	FlagsKeyListSupportRequiredRequestsPersec uint32

	// 
	FlagsPDCRequiredRequestsPersec uint32

	// 
	FlagsTimeServerRequiredRequestsPersec uint32

	// 
	FlagsTryNextClosestSiteRequestsPersec uint32

	// 
	FlagsWritableRequiredRequestsPersec uint32

	// 
	PingsMailslotPingsSentPersec uint32

	// 
	PingsUDPPingsSentPersec uint32

	// 
	RequestsAverageFailureLatencysecs uint32

	// 
	RequestsAverageSuccessLatencysecs uint32

	// 
	RequestsFailuresPersec uint32

	// 
	RequestsRejectedPersec uint32

	// 
	RequestsSuccessesPersec uint32

	// 
	RequestsTotalActive uint32

	// 
	UDPPortsOpened uint32
}

	func NewWin32_PerfFormattedData_Counters_DCLocatorNetlogonEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon, err error) {tmp, err := NewWin32_PerfFormattedDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_Counters_DCLocatorNetlogon {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

	func NewWin32_PerfFormattedData_Counters_DCLocatorNetlogonEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon, err error) {tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_Counters_DCLocatorNetlogon {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

// SetCacheHitsPersec sets the value of CacheHitsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyCacheHitsPersec(value uint32) (err error) { 
    return instance.SetProperty("CacheHitsPersec", (value))
}

// GetCacheHitsPersec gets the value of CacheHitsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyCacheHitsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CacheHitsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetCacheMissesPersec sets the value of CacheMissesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyCacheMissesPersec(value uint32) (err error) { 
    return instance.SetProperty("CacheMissesPersec", (value))
}

// GetCacheMissesPersec gets the value of CacheMissesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyCacheMissesPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CacheMissesPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetDNSQueryFailuresPersec sets the value of DNSQueryFailuresPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyDNSQueryFailuresPersec(value uint32) (err error) { 
    return instance.SetProperty("DNSQueryFailuresPersec", (value))
}

// GetDNSQueryFailuresPersec gets the value of DNSQueryFailuresPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyDNSQueryFailuresPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DNSQueryFailuresPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFlagsDSRequiredW2KRequestsPersec sets the value of FlagsDSRequiredW2KRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyFlagsDSRequiredW2KRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("FlagsDSRequiredW2KRequestsPersec", (value))
}

// GetFlagsDSRequiredW2KRequestsPersec gets the value of FlagsDSRequiredW2KRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyFlagsDSRequiredW2KRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FlagsDSRequiredW2KRequestsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFlagsDSRequiredWS2008RequestsPersec sets the value of FlagsDSRequiredWS2008RequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyFlagsDSRequiredWS2008RequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("FlagsDSRequiredWS2008RequestsPersec", (value))
}

// GetFlagsDSRequiredWS2008RequestsPersec gets the value of FlagsDSRequiredWS2008RequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyFlagsDSRequiredWS2008RequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FlagsDSRequiredWS2008RequestsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFlagsDSRequiredWS2012R2RequestsPersec sets the value of FlagsDSRequiredWS2012R2RequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyFlagsDSRequiredWS2012R2RequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("FlagsDSRequiredWS2012R2RequestsPersec", (value))
}

// GetFlagsDSRequiredWS2012R2RequestsPersec gets the value of FlagsDSRequiredWS2012R2RequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyFlagsDSRequiredWS2012R2RequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FlagsDSRequiredWS2012R2RequestsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFlagsDSRequiredWS2012RequestsPersec sets the value of FlagsDSRequiredWS2012RequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyFlagsDSRequiredWS2012RequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("FlagsDSRequiredWS2012RequestsPersec", (value))
}

// GetFlagsDSRequiredWS2012RequestsPersec gets the value of FlagsDSRequiredWS2012RequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyFlagsDSRequiredWS2012RequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FlagsDSRequiredWS2012RequestsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFlagsDSRequiredWS2016RequestsPersec sets the value of FlagsDSRequiredWS2016RequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyFlagsDSRequiredWS2016RequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("FlagsDSRequiredWS2016RequestsPersec", (value))
}

// GetFlagsDSRequiredWS2016RequestsPersec gets the value of FlagsDSRequiredWS2016RequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyFlagsDSRequiredWS2016RequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FlagsDSRequiredWS2016RequestsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFlagsForceRediscoveryRequestsPersec sets the value of FlagsForceRediscoveryRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyFlagsForceRediscoveryRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("FlagsForceRediscoveryRequestsPersec", (value))
}

// GetFlagsForceRediscoveryRequestsPersec gets the value of FlagsForceRediscoveryRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyFlagsForceRediscoveryRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FlagsForceRediscoveryRequestsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFlagsGCRequiredRequestsPersec sets the value of FlagsGCRequiredRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyFlagsGCRequiredRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("FlagsGCRequiredRequestsPersec", (value))
}

// GetFlagsGCRequiredRequestsPersec gets the value of FlagsGCRequiredRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyFlagsGCRequiredRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FlagsGCRequiredRequestsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFlagsKDCRequiredRequestsPersec sets the value of FlagsKDCRequiredRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyFlagsKDCRequiredRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("FlagsKDCRequiredRequestsPersec", (value))
}

// GetFlagsKDCRequiredRequestsPersec gets the value of FlagsKDCRequiredRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyFlagsKDCRequiredRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FlagsKDCRequiredRequestsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFlagsKeyListSupportRequiredRequestsPersec sets the value of FlagsKeyListSupportRequiredRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyFlagsKeyListSupportRequiredRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("FlagsKeyListSupportRequiredRequestsPersec", (value))
}

// GetFlagsKeyListSupportRequiredRequestsPersec gets the value of FlagsKeyListSupportRequiredRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyFlagsKeyListSupportRequiredRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FlagsKeyListSupportRequiredRequestsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFlagsPDCRequiredRequestsPersec sets the value of FlagsPDCRequiredRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyFlagsPDCRequiredRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("FlagsPDCRequiredRequestsPersec", (value))
}

// GetFlagsPDCRequiredRequestsPersec gets the value of FlagsPDCRequiredRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyFlagsPDCRequiredRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FlagsPDCRequiredRequestsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFlagsTimeServerRequiredRequestsPersec sets the value of FlagsTimeServerRequiredRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyFlagsTimeServerRequiredRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("FlagsTimeServerRequiredRequestsPersec", (value))
}

// GetFlagsTimeServerRequiredRequestsPersec gets the value of FlagsTimeServerRequiredRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyFlagsTimeServerRequiredRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FlagsTimeServerRequiredRequestsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFlagsTryNextClosestSiteRequestsPersec sets the value of FlagsTryNextClosestSiteRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyFlagsTryNextClosestSiteRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("FlagsTryNextClosestSiteRequestsPersec", (value))
}

// GetFlagsTryNextClosestSiteRequestsPersec gets the value of FlagsTryNextClosestSiteRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyFlagsTryNextClosestSiteRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FlagsTryNextClosestSiteRequestsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFlagsWritableRequiredRequestsPersec sets the value of FlagsWritableRequiredRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyFlagsWritableRequiredRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("FlagsWritableRequiredRequestsPersec", (value))
}

// GetFlagsWritableRequiredRequestsPersec gets the value of FlagsWritableRequiredRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyFlagsWritableRequiredRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FlagsWritableRequiredRequestsPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetPingsMailslotPingsSentPersec sets the value of PingsMailslotPingsSentPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyPingsMailslotPingsSentPersec(value uint32) (err error) { 
    return instance.SetProperty("PingsMailslotPingsSentPersec", (value))
}

// GetPingsMailslotPingsSentPersec gets the value of PingsMailslotPingsSentPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyPingsMailslotPingsSentPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PingsMailslotPingsSentPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetPingsUDPPingsSentPersec sets the value of PingsUDPPingsSentPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyPingsUDPPingsSentPersec(value uint32) (err error) { 
    return instance.SetProperty("PingsUDPPingsSentPersec", (value))
}

// GetPingsUDPPingsSentPersec gets the value of PingsUDPPingsSentPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyPingsUDPPingsSentPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PingsUDPPingsSentPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetRequestsAverageFailureLatencysecs sets the value of RequestsAverageFailureLatencysecs for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyRequestsAverageFailureLatencysecs(value uint32) (err error) { 
    return instance.SetProperty("RequestsAverageFailureLatencysecs", (value))
}

// GetRequestsAverageFailureLatencysecs gets the value of RequestsAverageFailureLatencysecs for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyRequestsAverageFailureLatencysecs()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RequestsAverageFailureLatencysecs")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetRequestsAverageSuccessLatencysecs sets the value of RequestsAverageSuccessLatencysecs for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyRequestsAverageSuccessLatencysecs(value uint32) (err error) { 
    return instance.SetProperty("RequestsAverageSuccessLatencysecs", (value))
}

// GetRequestsAverageSuccessLatencysecs gets the value of RequestsAverageSuccessLatencysecs for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyRequestsAverageSuccessLatencysecs()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RequestsAverageSuccessLatencysecs")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetRequestsFailuresPersec sets the value of RequestsFailuresPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyRequestsFailuresPersec(value uint32) (err error) { 
    return instance.SetProperty("RequestsFailuresPersec", (value))
}

// GetRequestsFailuresPersec gets the value of RequestsFailuresPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyRequestsFailuresPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RequestsFailuresPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetRequestsRejectedPersec sets the value of RequestsRejectedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyRequestsRejectedPersec(value uint32) (err error) { 
    return instance.SetProperty("RequestsRejectedPersec", (value))
}

// GetRequestsRejectedPersec gets the value of RequestsRejectedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyRequestsRejectedPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RequestsRejectedPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetRequestsSuccessesPersec sets the value of RequestsSuccessesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyRequestsSuccessesPersec(value uint32) (err error) { 
    return instance.SetProperty("RequestsSuccessesPersec", (value))
}

// GetRequestsSuccessesPersec gets the value of RequestsSuccessesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyRequestsSuccessesPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RequestsSuccessesPersec")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetRequestsTotalActive sets the value of RequestsTotalActive for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyRequestsTotalActive(value uint32) (err error) { 
    return instance.SetProperty("RequestsTotalActive", (value))
}

// GetRequestsTotalActive gets the value of RequestsTotalActive for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyRequestsTotalActive()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RequestsTotalActive")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetUDPPortsOpened sets the value of UDPPortsOpened for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) SetPropertyUDPPortsOpened(value uint32) (err error) { 
    return instance.SetProperty("UDPPortsOpened", (value))
}

// GetUDPPortsOpened gets the value of UDPPortsOpened for the instance
func (instance *Win32_PerfFormattedData_Counters_DCLocatorNetlogon) GetPropertyUDPPortsOpened()(value uint32, err error) { 
    retValue, err := instance.GetProperty("UDPPortsOpened")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

