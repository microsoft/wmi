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

// ads_msds_deviceregistrationservice struct
type ads_msds_deviceregistrationservice struct { 
	*ds_top

	// 
	DS_msDS_CloudIsEnabled bool

	// 
	DS_msDS_CloudIssuerPublicCertificates []Uint8Array

	// 
	DS_msDS_DeviceLocation string

	// 
	DS_msDS_IsEnabled bool

	// 
	DS_msDS_IssuerCertificates []Uint8Array

	// 
	DS_msDS_IssuerPublicCertificates []Uint8Array

	// 
	DS_msDS_MaximumRegistrationInactivityPeriod int32

	// 
	DS_msDS_RegistrationQuota int32
}

	func Newads_msds_deviceregistrationserviceEx1(instance *cim.WmiInstance) (newInstance *ads_msds_deviceregistrationservice, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_msds_deviceregistrationservice {
	ds_top: tmp,
	}
	return
	}
	

	func Newads_msds_deviceregistrationserviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_msds_deviceregistrationservice, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_msds_deviceregistrationservice {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_msDS_CloudIsEnabled sets the value of DS_msDS_CloudIsEnabled for the instance
func (instance *ads_msds_deviceregistrationservice) SetPropertyDS_msDS_CloudIsEnabled(value bool) (err error) { 
    return instance.SetProperty("DS_msDS_CloudIsEnabled", (value))
}

// GetDS_msDS_CloudIsEnabled gets the value of DS_msDS_CloudIsEnabled for the instance
func (instance *ads_msds_deviceregistrationservice) GetPropertyDS_msDS_CloudIsEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_CloudIsEnabled")
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

// SetDS_msDS_CloudIssuerPublicCertificates sets the value of DS_msDS_CloudIssuerPublicCertificates for the instance
func (instance *ads_msds_deviceregistrationservice) SetPropertyDS_msDS_CloudIssuerPublicCertificates(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_msDS_CloudIssuerPublicCertificates", (value))
}

// GetDS_msDS_CloudIssuerPublicCertificates gets the value of DS_msDS_CloudIssuerPublicCertificates for the instance
func (instance *ads_msds_deviceregistrationservice) GetPropertyDS_msDS_CloudIssuerPublicCertificates()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_CloudIssuerPublicCertificates")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(Uint8Array); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, Uint8Array(valuetmp))
    }

    return
}

// SetDS_msDS_DeviceLocation sets the value of DS_msDS_DeviceLocation for the instance
func (instance *ads_msds_deviceregistrationservice) SetPropertyDS_msDS_DeviceLocation(value string) (err error) { 
    return instance.SetProperty("DS_msDS_DeviceLocation", (value))
}

// GetDS_msDS_DeviceLocation gets the value of DS_msDS_DeviceLocation for the instance
func (instance *ads_msds_deviceregistrationservice) GetPropertyDS_msDS_DeviceLocation()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_DeviceLocation")
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

// SetDS_msDS_IsEnabled sets the value of DS_msDS_IsEnabled for the instance
func (instance *ads_msds_deviceregistrationservice) SetPropertyDS_msDS_IsEnabled(value bool) (err error) { 
    return instance.SetProperty("DS_msDS_IsEnabled", (value))
}

// GetDS_msDS_IsEnabled gets the value of DS_msDS_IsEnabled for the instance
func (instance *ads_msds_deviceregistrationservice) GetPropertyDS_msDS_IsEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_IsEnabled")
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

// SetDS_msDS_IssuerCertificates sets the value of DS_msDS_IssuerCertificates for the instance
func (instance *ads_msds_deviceregistrationservice) SetPropertyDS_msDS_IssuerCertificates(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_msDS_IssuerCertificates", (value))
}

// GetDS_msDS_IssuerCertificates gets the value of DS_msDS_IssuerCertificates for the instance
func (instance *ads_msds_deviceregistrationservice) GetPropertyDS_msDS_IssuerCertificates()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_IssuerCertificates")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(Uint8Array); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, Uint8Array(valuetmp))
    }

    return
}

// SetDS_msDS_IssuerPublicCertificates sets the value of DS_msDS_IssuerPublicCertificates for the instance
func (instance *ads_msds_deviceregistrationservice) SetPropertyDS_msDS_IssuerPublicCertificates(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_msDS_IssuerPublicCertificates", (value))
}

// GetDS_msDS_IssuerPublicCertificates gets the value of DS_msDS_IssuerPublicCertificates for the instance
func (instance *ads_msds_deviceregistrationservice) GetPropertyDS_msDS_IssuerPublicCertificates()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_IssuerPublicCertificates")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(Uint8Array); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, Uint8Array(valuetmp))
    }

    return
}

// SetDS_msDS_MaximumRegistrationInactivityPeriod sets the value of DS_msDS_MaximumRegistrationInactivityPeriod for the instance
func (instance *ads_msds_deviceregistrationservice) SetPropertyDS_msDS_MaximumRegistrationInactivityPeriod(value int32) (err error) { 
    return instance.SetProperty("DS_msDS_MaximumRegistrationInactivityPeriod", (value))
}

// GetDS_msDS_MaximumRegistrationInactivityPeriod gets the value of DS_msDS_MaximumRegistrationInactivityPeriod for the instance
func (instance *ads_msds_deviceregistrationservice) GetPropertyDS_msDS_MaximumRegistrationInactivityPeriod()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_MaximumRegistrationInactivityPeriod")
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

// SetDS_msDS_RegistrationQuota sets the value of DS_msDS_RegistrationQuota for the instance
func (instance *ads_msds_deviceregistrationservice) SetPropertyDS_msDS_RegistrationQuota(value int32) (err error) { 
    return instance.SetProperty("DS_msDS_RegistrationQuota", (value))
}

// GetDS_msDS_RegistrationQuota gets the value of DS_msDS_RegistrationQuota for the instance
func (instance *ads_msds_deviceregistrationservice) GetPropertyDS_msDS_RegistrationQuota()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_RegistrationQuota")
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

