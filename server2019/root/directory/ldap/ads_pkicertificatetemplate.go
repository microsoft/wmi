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

// ads_pkicertificatetemplate struct
type ads_pkicertificatetemplate struct { 
	*ds_top

	// 
	DS_msPKI_Cert_Template_OID string

	// 
	DS_msPKI_Certificate_Application_Policy []string

	// 
	DS_msPKI_Certificate_Name_Flag int32

	// 
	DS_msPKI_Certificate_Policy []string

	// 
	DS_msPKI_Enrollment_Flag int32

	// 
	DS_msPKI_Minimal_Key_Size int32

	// 
	DS_msPKI_Private_Key_Flag int32

	// 
	DS_msPKI_RA_Application_Policies []string

	// 
	DS_msPKI_RA_Policies []string

	// 
	DS_msPKI_RA_Signature int32

	// 
	DS_msPKI_Supersede_Templates []string

	// 
	DS_msPKI_Template_Minor_Revision int32

	// 
	DS_msPKI_Template_Schema_Version int32

	// 
	DS_pKICriticalExtensions []string

	// 
	DS_pKIDefaultCSPs []string

	// 
	DS_pKIDefaultKeySpec int32

	// 
	DS_pKIEnrollmentAccess []Uint8Array

	// 
	DS_pKIExpirationPeriod Uint8Array

	// 
	DS_pKIExtendedKeyUsage []string

	// 
	DS_pKIKeyUsage Uint8Array

	// 
	DS_pKIMaxIssuingDepth int32

	// 
	DS_pKIOverlapPeriod Uint8Array
}

	func Newads_pkicertificatetemplateEx1(instance *cim.WmiInstance) (newInstance *ads_pkicertificatetemplate, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_pkicertificatetemplate {
	ds_top: tmp,
	}
	return
	}
	

	func Newads_pkicertificatetemplateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_pkicertificatetemplate, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_pkicertificatetemplate {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_msPKI_Cert_Template_OID sets the value of DS_msPKI_Cert_Template_OID for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_msPKI_Cert_Template_OID(value string) (err error) { 
    return instance.SetProperty("DS_msPKI_Cert_Template_OID", (value))
}

// GetDS_msPKI_Cert_Template_OID gets the value of DS_msPKI_Cert_Template_OID for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_msPKI_Cert_Template_OID()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msPKI_Cert_Template_OID")
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

// SetDS_msPKI_Certificate_Application_Policy sets the value of DS_msPKI_Certificate_Application_Policy for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_msPKI_Certificate_Application_Policy(value []string) (err error) { 
    return instance.SetProperty("DS_msPKI_Certificate_Application_Policy", (value))
}

// GetDS_msPKI_Certificate_Application_Policy gets the value of DS_msPKI_Certificate_Application_Policy for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_msPKI_Certificate_Application_Policy()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_msPKI_Certificate_Application_Policy")
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

// SetDS_msPKI_Certificate_Name_Flag sets the value of DS_msPKI_Certificate_Name_Flag for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_msPKI_Certificate_Name_Flag(value int32) (err error) { 
    return instance.SetProperty("DS_msPKI_Certificate_Name_Flag", (value))
}

// GetDS_msPKI_Certificate_Name_Flag gets the value of DS_msPKI_Certificate_Name_Flag for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_msPKI_Certificate_Name_Flag()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msPKI_Certificate_Name_Flag")
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

// SetDS_msPKI_Certificate_Policy sets the value of DS_msPKI_Certificate_Policy for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_msPKI_Certificate_Policy(value []string) (err error) { 
    return instance.SetProperty("DS_msPKI_Certificate_Policy", (value))
}

// GetDS_msPKI_Certificate_Policy gets the value of DS_msPKI_Certificate_Policy for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_msPKI_Certificate_Policy()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_msPKI_Certificate_Policy")
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

// SetDS_msPKI_Enrollment_Flag sets the value of DS_msPKI_Enrollment_Flag for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_msPKI_Enrollment_Flag(value int32) (err error) { 
    return instance.SetProperty("DS_msPKI_Enrollment_Flag", (value))
}

// GetDS_msPKI_Enrollment_Flag gets the value of DS_msPKI_Enrollment_Flag for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_msPKI_Enrollment_Flag()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msPKI_Enrollment_Flag")
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

// SetDS_msPKI_Minimal_Key_Size sets the value of DS_msPKI_Minimal_Key_Size for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_msPKI_Minimal_Key_Size(value int32) (err error) { 
    return instance.SetProperty("DS_msPKI_Minimal_Key_Size", (value))
}

// GetDS_msPKI_Minimal_Key_Size gets the value of DS_msPKI_Minimal_Key_Size for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_msPKI_Minimal_Key_Size()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msPKI_Minimal_Key_Size")
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

// SetDS_msPKI_Private_Key_Flag sets the value of DS_msPKI_Private_Key_Flag for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_msPKI_Private_Key_Flag(value int32) (err error) { 
    return instance.SetProperty("DS_msPKI_Private_Key_Flag", (value))
}

// GetDS_msPKI_Private_Key_Flag gets the value of DS_msPKI_Private_Key_Flag for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_msPKI_Private_Key_Flag()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msPKI_Private_Key_Flag")
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

// SetDS_msPKI_RA_Application_Policies sets the value of DS_msPKI_RA_Application_Policies for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_msPKI_RA_Application_Policies(value []string) (err error) { 
    return instance.SetProperty("DS_msPKI_RA_Application_Policies", (value))
}

// GetDS_msPKI_RA_Application_Policies gets the value of DS_msPKI_RA_Application_Policies for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_msPKI_RA_Application_Policies()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_msPKI_RA_Application_Policies")
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

// SetDS_msPKI_RA_Policies sets the value of DS_msPKI_RA_Policies for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_msPKI_RA_Policies(value []string) (err error) { 
    return instance.SetProperty("DS_msPKI_RA_Policies", (value))
}

// GetDS_msPKI_RA_Policies gets the value of DS_msPKI_RA_Policies for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_msPKI_RA_Policies()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_msPKI_RA_Policies")
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

// SetDS_msPKI_RA_Signature sets the value of DS_msPKI_RA_Signature for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_msPKI_RA_Signature(value int32) (err error) { 
    return instance.SetProperty("DS_msPKI_RA_Signature", (value))
}

// GetDS_msPKI_RA_Signature gets the value of DS_msPKI_RA_Signature for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_msPKI_RA_Signature()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msPKI_RA_Signature")
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

// SetDS_msPKI_Supersede_Templates sets the value of DS_msPKI_Supersede_Templates for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_msPKI_Supersede_Templates(value []string) (err error) { 
    return instance.SetProperty("DS_msPKI_Supersede_Templates", (value))
}

// GetDS_msPKI_Supersede_Templates gets the value of DS_msPKI_Supersede_Templates for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_msPKI_Supersede_Templates()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_msPKI_Supersede_Templates")
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

// SetDS_msPKI_Template_Minor_Revision sets the value of DS_msPKI_Template_Minor_Revision for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_msPKI_Template_Minor_Revision(value int32) (err error) { 
    return instance.SetProperty("DS_msPKI_Template_Minor_Revision", (value))
}

// GetDS_msPKI_Template_Minor_Revision gets the value of DS_msPKI_Template_Minor_Revision for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_msPKI_Template_Minor_Revision()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msPKI_Template_Minor_Revision")
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

// SetDS_msPKI_Template_Schema_Version sets the value of DS_msPKI_Template_Schema_Version for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_msPKI_Template_Schema_Version(value int32) (err error) { 
    return instance.SetProperty("DS_msPKI_Template_Schema_Version", (value))
}

// GetDS_msPKI_Template_Schema_Version gets the value of DS_msPKI_Template_Schema_Version for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_msPKI_Template_Schema_Version()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msPKI_Template_Schema_Version")
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

// SetDS_pKICriticalExtensions sets the value of DS_pKICriticalExtensions for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_pKICriticalExtensions(value []string) (err error) { 
    return instance.SetProperty("DS_pKICriticalExtensions", (value))
}

// GetDS_pKICriticalExtensions gets the value of DS_pKICriticalExtensions for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_pKICriticalExtensions()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_pKICriticalExtensions")
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

// SetDS_pKIDefaultCSPs sets the value of DS_pKIDefaultCSPs for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_pKIDefaultCSPs(value []string) (err error) { 
    return instance.SetProperty("DS_pKIDefaultCSPs", (value))
}

// GetDS_pKIDefaultCSPs gets the value of DS_pKIDefaultCSPs for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_pKIDefaultCSPs()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_pKIDefaultCSPs")
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

// SetDS_pKIDefaultKeySpec sets the value of DS_pKIDefaultKeySpec for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_pKIDefaultKeySpec(value int32) (err error) { 
    return instance.SetProperty("DS_pKIDefaultKeySpec", (value))
}

// GetDS_pKIDefaultKeySpec gets the value of DS_pKIDefaultKeySpec for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_pKIDefaultKeySpec()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_pKIDefaultKeySpec")
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

// SetDS_pKIEnrollmentAccess sets the value of DS_pKIEnrollmentAccess for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_pKIEnrollmentAccess(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_pKIEnrollmentAccess", (value))
}

// GetDS_pKIEnrollmentAccess gets the value of DS_pKIEnrollmentAccess for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_pKIEnrollmentAccess()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_pKIEnrollmentAccess")
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

// SetDS_pKIExpirationPeriod sets the value of DS_pKIExpirationPeriod for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_pKIExpirationPeriod(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_pKIExpirationPeriod", (value))
}

// GetDS_pKIExpirationPeriod gets the value of DS_pKIExpirationPeriod for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_pKIExpirationPeriod()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_pKIExpirationPeriod")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(Uint8Array); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Uint8Array(valuetmp)

    return
}

// SetDS_pKIExtendedKeyUsage sets the value of DS_pKIExtendedKeyUsage for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_pKIExtendedKeyUsage(value []string) (err error) { 
    return instance.SetProperty("DS_pKIExtendedKeyUsage", (value))
}

// GetDS_pKIExtendedKeyUsage gets the value of DS_pKIExtendedKeyUsage for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_pKIExtendedKeyUsage()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_pKIExtendedKeyUsage")
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

// SetDS_pKIKeyUsage sets the value of DS_pKIKeyUsage for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_pKIKeyUsage(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_pKIKeyUsage", (value))
}

// GetDS_pKIKeyUsage gets the value of DS_pKIKeyUsage for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_pKIKeyUsage()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_pKIKeyUsage")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(Uint8Array); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Uint8Array(valuetmp)

    return
}

// SetDS_pKIMaxIssuingDepth sets the value of DS_pKIMaxIssuingDepth for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_pKIMaxIssuingDepth(value int32) (err error) { 
    return instance.SetProperty("DS_pKIMaxIssuingDepth", (value))
}

// GetDS_pKIMaxIssuingDepth gets the value of DS_pKIMaxIssuingDepth for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_pKIMaxIssuingDepth()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_pKIMaxIssuingDepth")
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

// SetDS_pKIOverlapPeriod sets the value of DS_pKIOverlapPeriod for the instance
func (instance *ads_pkicertificatetemplate) SetPropertyDS_pKIOverlapPeriod(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_pKIOverlapPeriod", (value))
}

// GetDS_pKIOverlapPeriod gets the value of DS_pKIOverlapPeriod for the instance
func (instance *ads_pkicertificatetemplate) GetPropertyDS_pKIOverlapPeriod()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_pKIOverlapPeriod")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(Uint8Array); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Uint8Array(valuetmp)

    return
}

