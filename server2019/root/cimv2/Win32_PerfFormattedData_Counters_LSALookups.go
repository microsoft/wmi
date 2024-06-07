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

// Win32_PerfFormattedData_Counters_LSALookups struct
type Win32_PerfFormattedData_Counters_LSALookups struct { 
	*Win32_PerfFormattedData

	// 
	IsolatedNamesInboundRequestsPersec uint32

	// 
	IsolatedNamesOutboundRequestsPersec uint32

	// 
	NamesCachePercentFull uint32

	// 
	NamesCachePercentHit uint32

	// 
	NamesCompletionTime uint32

	// 
	NamesErrorsPersec uint32

	// 
	NameSIDcacheentriesaddedPersec uint32

	// 
	NameSIDcacheentriespurgedPersec uint32

	// 
	NameSIDCacheSizeMaxEntries uint32

	// 
	NamesInboundRequestsPersec uint32

	// 
	NamesOutboundRequestsPersec uint32

	// 
	NamesPrimaryDomainRequestsPersec uint32

	// 
	NamesPrimaryDomainTime uint32

	// 
	NamesRemoteRequestTime uint32

	// 
	NamesTrustedDomainRequestsPersec uint32

	// 
	NamesTrustedDomainRequestTime uint32

	// 
	NamesUnresolvedPersec uint32

	// 
	NamesXforestRequestsPersec uint32

	// 
	NamesXforestTime uint32

	// 
	SIDsCachePercentFull uint32

	// 
	SIDsCachePercentHit uint32

	// 
	SIDsCompletionTime uint32

	// 
	SIDsErrorsPersec uint32

	// 
	SIDsInboundRequestsPersec uint32

	// 
	SIDsOutboundRequestsPersec uint32

	// 
	SIDsPrimaryDomainRequestsPersec uint32

	// 
	SIDsPrimaryDomainRequestTime uint32

	// 
	SIDsRemoteRequestTime uint32

	// 
	SIDsTrustedDomainRequestsPersec uint32

	// 
	SIDsTrustedDomainRequestTime uint32

	// 
	SIDsUnresolvedPersec uint32

	// 
	SIDsXforestRequestsPersec uint32

	// 
	SIDsXforestRequestTime uint32
}

	func NewWin32_PerfFormattedData_Counters_LSALookupsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Counters_LSALookups, err error) {tmp, err := NewWin32_PerfFormattedDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_Counters_LSALookups {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

	func NewWin32_PerfFormattedData_Counters_LSALookupsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfFormattedData_Counters_LSALookups, err error) {tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_Counters_LSALookups {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

// SetIsolatedNamesInboundRequestsPersec sets the value of IsolatedNamesInboundRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertyIsolatedNamesInboundRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("IsolatedNamesInboundRequestsPersec", (value))
}

// GetIsolatedNamesInboundRequestsPersec gets the value of IsolatedNamesInboundRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertyIsolatedNamesInboundRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IsolatedNamesInboundRequestsPersec")
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

// SetIsolatedNamesOutboundRequestsPersec sets the value of IsolatedNamesOutboundRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertyIsolatedNamesOutboundRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("IsolatedNamesOutboundRequestsPersec", (value))
}

// GetIsolatedNamesOutboundRequestsPersec gets the value of IsolatedNamesOutboundRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertyIsolatedNamesOutboundRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IsolatedNamesOutboundRequestsPersec")
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

// SetNamesCachePercentFull sets the value of NamesCachePercentFull for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertyNamesCachePercentFull(value uint32) (err error) { 
    return instance.SetProperty("NamesCachePercentFull", (value))
}

// GetNamesCachePercentFull gets the value of NamesCachePercentFull for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertyNamesCachePercentFull()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NamesCachePercentFull")
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

// SetNamesCachePercentHit sets the value of NamesCachePercentHit for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertyNamesCachePercentHit(value uint32) (err error) { 
    return instance.SetProperty("NamesCachePercentHit", (value))
}

// GetNamesCachePercentHit gets the value of NamesCachePercentHit for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertyNamesCachePercentHit()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NamesCachePercentHit")
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

// SetNamesCompletionTime sets the value of NamesCompletionTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertyNamesCompletionTime(value uint32) (err error) { 
    return instance.SetProperty("NamesCompletionTime", (value))
}

// GetNamesCompletionTime gets the value of NamesCompletionTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertyNamesCompletionTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NamesCompletionTime")
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

// SetNamesErrorsPersec sets the value of NamesErrorsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertyNamesErrorsPersec(value uint32) (err error) { 
    return instance.SetProperty("NamesErrorsPersec", (value))
}

// GetNamesErrorsPersec gets the value of NamesErrorsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertyNamesErrorsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NamesErrorsPersec")
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

// SetNameSIDcacheentriesaddedPersec sets the value of NameSIDcacheentriesaddedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertyNameSIDcacheentriesaddedPersec(value uint32) (err error) { 
    return instance.SetProperty("NameSIDcacheentriesaddedPersec", (value))
}

// GetNameSIDcacheentriesaddedPersec gets the value of NameSIDcacheentriesaddedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertyNameSIDcacheentriesaddedPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NameSIDcacheentriesaddedPersec")
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

// SetNameSIDcacheentriespurgedPersec sets the value of NameSIDcacheentriespurgedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertyNameSIDcacheentriespurgedPersec(value uint32) (err error) { 
    return instance.SetProperty("NameSIDcacheentriespurgedPersec", (value))
}

// GetNameSIDcacheentriespurgedPersec gets the value of NameSIDcacheentriespurgedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertyNameSIDcacheentriespurgedPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NameSIDcacheentriespurgedPersec")
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

// SetNameSIDCacheSizeMaxEntries sets the value of NameSIDCacheSizeMaxEntries for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertyNameSIDCacheSizeMaxEntries(value uint32) (err error) { 
    return instance.SetProperty("NameSIDCacheSizeMaxEntries", (value))
}

// GetNameSIDCacheSizeMaxEntries gets the value of NameSIDCacheSizeMaxEntries for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertyNameSIDCacheSizeMaxEntries()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NameSIDCacheSizeMaxEntries")
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

// SetNamesInboundRequestsPersec sets the value of NamesInboundRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertyNamesInboundRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("NamesInboundRequestsPersec", (value))
}

// GetNamesInboundRequestsPersec gets the value of NamesInboundRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertyNamesInboundRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NamesInboundRequestsPersec")
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

// SetNamesOutboundRequestsPersec sets the value of NamesOutboundRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertyNamesOutboundRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("NamesOutboundRequestsPersec", (value))
}

// GetNamesOutboundRequestsPersec gets the value of NamesOutboundRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertyNamesOutboundRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NamesOutboundRequestsPersec")
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

// SetNamesPrimaryDomainRequestsPersec sets the value of NamesPrimaryDomainRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertyNamesPrimaryDomainRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("NamesPrimaryDomainRequestsPersec", (value))
}

// GetNamesPrimaryDomainRequestsPersec gets the value of NamesPrimaryDomainRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertyNamesPrimaryDomainRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NamesPrimaryDomainRequestsPersec")
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

// SetNamesPrimaryDomainTime sets the value of NamesPrimaryDomainTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertyNamesPrimaryDomainTime(value uint32) (err error) { 
    return instance.SetProperty("NamesPrimaryDomainTime", (value))
}

// GetNamesPrimaryDomainTime gets the value of NamesPrimaryDomainTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertyNamesPrimaryDomainTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NamesPrimaryDomainTime")
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

// SetNamesRemoteRequestTime sets the value of NamesRemoteRequestTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertyNamesRemoteRequestTime(value uint32) (err error) { 
    return instance.SetProperty("NamesRemoteRequestTime", (value))
}

// GetNamesRemoteRequestTime gets the value of NamesRemoteRequestTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertyNamesRemoteRequestTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NamesRemoteRequestTime")
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

// SetNamesTrustedDomainRequestsPersec sets the value of NamesTrustedDomainRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertyNamesTrustedDomainRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("NamesTrustedDomainRequestsPersec", (value))
}

// GetNamesTrustedDomainRequestsPersec gets the value of NamesTrustedDomainRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertyNamesTrustedDomainRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NamesTrustedDomainRequestsPersec")
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

// SetNamesTrustedDomainRequestTime sets the value of NamesTrustedDomainRequestTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertyNamesTrustedDomainRequestTime(value uint32) (err error) { 
    return instance.SetProperty("NamesTrustedDomainRequestTime", (value))
}

// GetNamesTrustedDomainRequestTime gets the value of NamesTrustedDomainRequestTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertyNamesTrustedDomainRequestTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NamesTrustedDomainRequestTime")
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

// SetNamesUnresolvedPersec sets the value of NamesUnresolvedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertyNamesUnresolvedPersec(value uint32) (err error) { 
    return instance.SetProperty("NamesUnresolvedPersec", (value))
}

// GetNamesUnresolvedPersec gets the value of NamesUnresolvedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertyNamesUnresolvedPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NamesUnresolvedPersec")
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

// SetNamesXforestRequestsPersec sets the value of NamesXforestRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertyNamesXforestRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("NamesXforestRequestsPersec", (value))
}

// GetNamesXforestRequestsPersec gets the value of NamesXforestRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertyNamesXforestRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NamesXforestRequestsPersec")
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

// SetNamesXforestTime sets the value of NamesXforestTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertyNamesXforestTime(value uint32) (err error) { 
    return instance.SetProperty("NamesXforestTime", (value))
}

// GetNamesXforestTime gets the value of NamesXforestTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertyNamesXforestTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NamesXforestTime")
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

// SetSIDsCachePercentFull sets the value of SIDsCachePercentFull for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertySIDsCachePercentFull(value uint32) (err error) { 
    return instance.SetProperty("SIDsCachePercentFull", (value))
}

// GetSIDsCachePercentFull gets the value of SIDsCachePercentFull for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertySIDsCachePercentFull()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SIDsCachePercentFull")
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

// SetSIDsCachePercentHit sets the value of SIDsCachePercentHit for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertySIDsCachePercentHit(value uint32) (err error) { 
    return instance.SetProperty("SIDsCachePercentHit", (value))
}

// GetSIDsCachePercentHit gets the value of SIDsCachePercentHit for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertySIDsCachePercentHit()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SIDsCachePercentHit")
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

// SetSIDsCompletionTime sets the value of SIDsCompletionTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertySIDsCompletionTime(value uint32) (err error) { 
    return instance.SetProperty("SIDsCompletionTime", (value))
}

// GetSIDsCompletionTime gets the value of SIDsCompletionTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertySIDsCompletionTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SIDsCompletionTime")
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

// SetSIDsErrorsPersec sets the value of SIDsErrorsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertySIDsErrorsPersec(value uint32) (err error) { 
    return instance.SetProperty("SIDsErrorsPersec", (value))
}

// GetSIDsErrorsPersec gets the value of SIDsErrorsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertySIDsErrorsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SIDsErrorsPersec")
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

// SetSIDsInboundRequestsPersec sets the value of SIDsInboundRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertySIDsInboundRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("SIDsInboundRequestsPersec", (value))
}

// GetSIDsInboundRequestsPersec gets the value of SIDsInboundRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertySIDsInboundRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SIDsInboundRequestsPersec")
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

// SetSIDsOutboundRequestsPersec sets the value of SIDsOutboundRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertySIDsOutboundRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("SIDsOutboundRequestsPersec", (value))
}

// GetSIDsOutboundRequestsPersec gets the value of SIDsOutboundRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertySIDsOutboundRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SIDsOutboundRequestsPersec")
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

// SetSIDsPrimaryDomainRequestsPersec sets the value of SIDsPrimaryDomainRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertySIDsPrimaryDomainRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("SIDsPrimaryDomainRequestsPersec", (value))
}

// GetSIDsPrimaryDomainRequestsPersec gets the value of SIDsPrimaryDomainRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertySIDsPrimaryDomainRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SIDsPrimaryDomainRequestsPersec")
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

// SetSIDsPrimaryDomainRequestTime sets the value of SIDsPrimaryDomainRequestTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertySIDsPrimaryDomainRequestTime(value uint32) (err error) { 
    return instance.SetProperty("SIDsPrimaryDomainRequestTime", (value))
}

// GetSIDsPrimaryDomainRequestTime gets the value of SIDsPrimaryDomainRequestTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertySIDsPrimaryDomainRequestTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SIDsPrimaryDomainRequestTime")
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

// SetSIDsRemoteRequestTime sets the value of SIDsRemoteRequestTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertySIDsRemoteRequestTime(value uint32) (err error) { 
    return instance.SetProperty("SIDsRemoteRequestTime", (value))
}

// GetSIDsRemoteRequestTime gets the value of SIDsRemoteRequestTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertySIDsRemoteRequestTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SIDsRemoteRequestTime")
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

// SetSIDsTrustedDomainRequestsPersec sets the value of SIDsTrustedDomainRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertySIDsTrustedDomainRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("SIDsTrustedDomainRequestsPersec", (value))
}

// GetSIDsTrustedDomainRequestsPersec gets the value of SIDsTrustedDomainRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertySIDsTrustedDomainRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SIDsTrustedDomainRequestsPersec")
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

// SetSIDsTrustedDomainRequestTime sets the value of SIDsTrustedDomainRequestTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertySIDsTrustedDomainRequestTime(value uint32) (err error) { 
    return instance.SetProperty("SIDsTrustedDomainRequestTime", (value))
}

// GetSIDsTrustedDomainRequestTime gets the value of SIDsTrustedDomainRequestTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertySIDsTrustedDomainRequestTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SIDsTrustedDomainRequestTime")
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

// SetSIDsUnresolvedPersec sets the value of SIDsUnresolvedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertySIDsUnresolvedPersec(value uint32) (err error) { 
    return instance.SetProperty("SIDsUnresolvedPersec", (value))
}

// GetSIDsUnresolvedPersec gets the value of SIDsUnresolvedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertySIDsUnresolvedPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SIDsUnresolvedPersec")
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

// SetSIDsXforestRequestsPersec sets the value of SIDsXforestRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertySIDsXforestRequestsPersec(value uint32) (err error) { 
    return instance.SetProperty("SIDsXforestRequestsPersec", (value))
}

// GetSIDsXforestRequestsPersec gets the value of SIDsXforestRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertySIDsXforestRequestsPersec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SIDsXforestRequestsPersec")
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

// SetSIDsXforestRequestTime sets the value of SIDsXforestRequestTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) SetPropertySIDsXforestRequestTime(value uint32) (err error) { 
    return instance.SetProperty("SIDsXforestRequestTime", (value))
}

// GetSIDsXforestRequestTime gets the value of SIDsXforestRequestTime for the instance
func (instance *Win32_PerfFormattedData_Counters_LSALookups) GetPropertySIDsXforestRequestTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SIDsXforestRequestTime")
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

