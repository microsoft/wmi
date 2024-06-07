// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// ads_server struct
type ads_server struct { 
	*ds_top

	// 
	DS_bridgeheadTransportList []string

	// 
	DS_dNSHostName string

	// 
	DS_mailAddress string

	// 
	DS_managedBy string

	// 
	DS_msDS_isGC bool

	// 
	DS_msDS_isRODC bool

	// 
	DS_msDS_IsUserCachableAtRodc int32

	// 
	DS_msDS_SiteName string

	// 
	DS_serialNumber []string

	// 
	DS_serverReference string
}

	func Newads_serverEx1(instance *cim.WmiInstance) (newInstance *ads_server, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_server {
	ds_top: tmp,
	}
	return
	}
	

	func Newads_serverEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_server, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_server {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_bridgeheadTransportList sets the value of DS_bridgeheadTransportList for the instance
func (instance *ads_server) SetPropertyDS_bridgeheadTransportList(value []string) (err error) { 
    return instance.SetProperty("DS_bridgeheadTransportList", (value))
}

// GetDS_bridgeheadTransportList gets the value of DS_bridgeheadTransportList for the instance
func (instance *ads_server) GetPropertyDS_bridgeheadTransportList()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_bridgeheadTransportList")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

// SetDS_dNSHostName sets the value of DS_dNSHostName for the instance
func (instance *ads_server) SetPropertyDS_dNSHostName(value string) (err error) { 
    return instance.SetProperty("DS_dNSHostName", (value))
}

// GetDS_dNSHostName gets the value of DS_dNSHostName for the instance
func (instance *ads_server) GetPropertyDS_dNSHostName()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_dNSHostName")
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

// SetDS_mailAddress sets the value of DS_mailAddress for the instance
func (instance *ads_server) SetPropertyDS_mailAddress(value string) (err error) { 
    return instance.SetProperty("DS_mailAddress", (value))
}

// GetDS_mailAddress gets the value of DS_mailAddress for the instance
func (instance *ads_server) GetPropertyDS_mailAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_mailAddress")
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

// SetDS_managedBy sets the value of DS_managedBy for the instance
func (instance *ads_server) SetPropertyDS_managedBy(value string) (err error) { 
    return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_server) GetPropertyDS_managedBy()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_managedBy")
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

// SetDS_msDS_isGC sets the value of DS_msDS_isGC for the instance
func (instance *ads_server) SetPropertyDS_msDS_isGC(value bool) (err error) { 
    return instance.SetProperty("DS_msDS_isGC", (value))
}

// GetDS_msDS_isGC gets the value of DS_msDS_isGC for the instance
func (instance *ads_server) GetPropertyDS_msDS_isGC()(value bool, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_isGC")
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

// SetDS_msDS_isRODC sets the value of DS_msDS_isRODC for the instance
func (instance *ads_server) SetPropertyDS_msDS_isRODC(value bool) (err error) { 
    return instance.SetProperty("DS_msDS_isRODC", (value))
}

// GetDS_msDS_isRODC gets the value of DS_msDS_isRODC for the instance
func (instance *ads_server) GetPropertyDS_msDS_isRODC()(value bool, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_isRODC")
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

// SetDS_msDS_IsUserCachableAtRodc sets the value of DS_msDS_IsUserCachableAtRodc for the instance
func (instance *ads_server) SetPropertyDS_msDS_IsUserCachableAtRodc(value int32) (err error) { 
    return instance.SetProperty("DS_msDS_IsUserCachableAtRodc", (value))
}

// GetDS_msDS_IsUserCachableAtRodc gets the value of DS_msDS_IsUserCachableAtRodc for the instance
func (instance *ads_server) GetPropertyDS_msDS_IsUserCachableAtRodc()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_IsUserCachableAtRodc")
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

// SetDS_msDS_SiteName sets the value of DS_msDS_SiteName for the instance
func (instance *ads_server) SetPropertyDS_msDS_SiteName(value string) (err error) { 
    return instance.SetProperty("DS_msDS_SiteName", (value))
}

// GetDS_msDS_SiteName gets the value of DS_msDS_SiteName for the instance
func (instance *ads_server) GetPropertyDS_msDS_SiteName()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_SiteName")
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

// SetDS_serialNumber sets the value of DS_serialNumber for the instance
func (instance *ads_server) SetPropertyDS_serialNumber(value []string) (err error) { 
    return instance.SetProperty("DS_serialNumber", (value))
}

// GetDS_serialNumber gets the value of DS_serialNumber for the instance
func (instance *ads_server) GetPropertyDS_serialNumber()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_serialNumber")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

// SetDS_serverReference sets the value of DS_serverReference for the instance
func (instance *ads_server) SetPropertyDS_serverReference(value string) (err error) { 
    return instance.SetProperty("DS_serverReference", (value))
}

// GetDS_serverReference gets the value of DS_serverReference for the instance
func (instance *ads_server) GetPropertyDS_serverReference()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_serverReference")
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

