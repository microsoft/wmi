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

// ads_contact struct
type ads_contact struct { 
	*ads_organizationalperson

	// 
	DS_garbageCollPeriod int32

	// 
	DS_info string

	// 
	DS_labeledURI []string

	// 
	DS_legacyExchangeDN string

	// 
	DS_msDS_ExternalDirectoryObjectId string

	// 
	DS_msDS_GeoCoordinatesAltitude int64

	// 
	DS_msDS_GeoCoordinatesLatitude int64

	// 
	DS_msDS_GeoCoordinatesLongitude int64

	// 
	DS_msDS_preferredDataLocation string

	// 
	DS_msDS_SourceObjectDN string

	// 
	DS_msExchAssistantName string

	// 
	DS_msExchLabeledURI []string

	// 
	DS_notes string

	// 
	DS_secretary []string

	// 
	DS_showInAddressBook []string

	// 
	DS_textEncodedORAddress string

	// 
	DS_userCert Uint8Array

	// 
	DS_userCertificate []Uint8Array

	// 
	DS_userSMIMECertificate []Uint8Array
}

	func Newads_contactEx1(instance *cim.WmiInstance) (newInstance *ads_contact, err error) {tmp, err := Newads_organizationalpersonEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_contact {
	ads_organizationalperson: tmp,
	}
	return
	}
	

	func Newads_contactEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_contact, err error) {tmp, err := Newads_organizationalpersonEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_contact {
	ads_organizationalperson: tmp,
	}
	return
	}
	

// SetDS_garbageCollPeriod sets the value of DS_garbageCollPeriod for the instance
func (instance *ads_contact) SetPropertyDS_garbageCollPeriod(value int32) (err error) { 
    return instance.SetProperty("DS_garbageCollPeriod", (value))
}

// GetDS_garbageCollPeriod gets the value of DS_garbageCollPeriod for the instance
func (instance *ads_contact) GetPropertyDS_garbageCollPeriod()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_garbageCollPeriod")
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

// SetDS_info sets the value of DS_info for the instance
func (instance *ads_contact) SetPropertyDS_info(value string) (err error) { 
    return instance.SetProperty("DS_info", (value))
}

// GetDS_info gets the value of DS_info for the instance
func (instance *ads_contact) GetPropertyDS_info()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_info")
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

// SetDS_labeledURI sets the value of DS_labeledURI for the instance
func (instance *ads_contact) SetPropertyDS_labeledURI(value []string) (err error) { 
    return instance.SetProperty("DS_labeledURI", (value))
}

// GetDS_labeledURI gets the value of DS_labeledURI for the instance
func (instance *ads_contact) GetPropertyDS_labeledURI()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_labeledURI")
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

// SetDS_legacyExchangeDN sets the value of DS_legacyExchangeDN for the instance
func (instance *ads_contact) SetPropertyDS_legacyExchangeDN(value string) (err error) { 
    return instance.SetProperty("DS_legacyExchangeDN", (value))
}

// GetDS_legacyExchangeDN gets the value of DS_legacyExchangeDN for the instance
func (instance *ads_contact) GetPropertyDS_legacyExchangeDN()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_legacyExchangeDN")
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

// SetDS_msDS_ExternalDirectoryObjectId sets the value of DS_msDS_ExternalDirectoryObjectId for the instance
func (instance *ads_contact) SetPropertyDS_msDS_ExternalDirectoryObjectId(value string) (err error) { 
    return instance.SetProperty("DS_msDS_ExternalDirectoryObjectId", (value))
}

// GetDS_msDS_ExternalDirectoryObjectId gets the value of DS_msDS_ExternalDirectoryObjectId for the instance
func (instance *ads_contact) GetPropertyDS_msDS_ExternalDirectoryObjectId()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_ExternalDirectoryObjectId")
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

// SetDS_msDS_GeoCoordinatesAltitude sets the value of DS_msDS_GeoCoordinatesAltitude for the instance
func (instance *ads_contact) SetPropertyDS_msDS_GeoCoordinatesAltitude(value int64) (err error) { 
    return instance.SetProperty("DS_msDS_GeoCoordinatesAltitude", (value))
}

// GetDS_msDS_GeoCoordinatesAltitude gets the value of DS_msDS_GeoCoordinatesAltitude for the instance
func (instance *ads_contact) GetPropertyDS_msDS_GeoCoordinatesAltitude()(value int64, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_GeoCoordinatesAltitude")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int64(valuetmp)

    return
}

// SetDS_msDS_GeoCoordinatesLatitude sets the value of DS_msDS_GeoCoordinatesLatitude for the instance
func (instance *ads_contact) SetPropertyDS_msDS_GeoCoordinatesLatitude(value int64) (err error) { 
    return instance.SetProperty("DS_msDS_GeoCoordinatesLatitude", (value))
}

// GetDS_msDS_GeoCoordinatesLatitude gets the value of DS_msDS_GeoCoordinatesLatitude for the instance
func (instance *ads_contact) GetPropertyDS_msDS_GeoCoordinatesLatitude()(value int64, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_GeoCoordinatesLatitude")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int64(valuetmp)

    return
}

// SetDS_msDS_GeoCoordinatesLongitude sets the value of DS_msDS_GeoCoordinatesLongitude for the instance
func (instance *ads_contact) SetPropertyDS_msDS_GeoCoordinatesLongitude(value int64) (err error) { 
    return instance.SetProperty("DS_msDS_GeoCoordinatesLongitude", (value))
}

// GetDS_msDS_GeoCoordinatesLongitude gets the value of DS_msDS_GeoCoordinatesLongitude for the instance
func (instance *ads_contact) GetPropertyDS_msDS_GeoCoordinatesLongitude()(value int64, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_GeoCoordinatesLongitude")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int64(valuetmp)

    return
}

// SetDS_msDS_preferredDataLocation sets the value of DS_msDS_preferredDataLocation for the instance
func (instance *ads_contact) SetPropertyDS_msDS_preferredDataLocation(value string) (err error) { 
    return instance.SetProperty("DS_msDS_preferredDataLocation", (value))
}

// GetDS_msDS_preferredDataLocation gets the value of DS_msDS_preferredDataLocation for the instance
func (instance *ads_contact) GetPropertyDS_msDS_preferredDataLocation()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_preferredDataLocation")
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

// SetDS_msDS_SourceObjectDN sets the value of DS_msDS_SourceObjectDN for the instance
func (instance *ads_contact) SetPropertyDS_msDS_SourceObjectDN(value string) (err error) { 
    return instance.SetProperty("DS_msDS_SourceObjectDN", (value))
}

// GetDS_msDS_SourceObjectDN gets the value of DS_msDS_SourceObjectDN for the instance
func (instance *ads_contact) GetPropertyDS_msDS_SourceObjectDN()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_SourceObjectDN")
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

// SetDS_msExchAssistantName sets the value of DS_msExchAssistantName for the instance
func (instance *ads_contact) SetPropertyDS_msExchAssistantName(value string) (err error) { 
    return instance.SetProperty("DS_msExchAssistantName", (value))
}

// GetDS_msExchAssistantName gets the value of DS_msExchAssistantName for the instance
func (instance *ads_contact) GetPropertyDS_msExchAssistantName()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msExchAssistantName")
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

// SetDS_msExchLabeledURI sets the value of DS_msExchLabeledURI for the instance
func (instance *ads_contact) SetPropertyDS_msExchLabeledURI(value []string) (err error) { 
    return instance.SetProperty("DS_msExchLabeledURI", (value))
}

// GetDS_msExchLabeledURI gets the value of DS_msExchLabeledURI for the instance
func (instance *ads_contact) GetPropertyDS_msExchLabeledURI()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_msExchLabeledURI")
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

// SetDS_notes sets the value of DS_notes for the instance
func (instance *ads_contact) SetPropertyDS_notes(value string) (err error) { 
    return instance.SetProperty("DS_notes", (value))
}

// GetDS_notes gets the value of DS_notes for the instance
func (instance *ads_contact) GetPropertyDS_notes()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_notes")
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

// SetDS_secretary sets the value of DS_secretary for the instance
func (instance *ads_contact) SetPropertyDS_secretary(value []string) (err error) { 
    return instance.SetProperty("DS_secretary", (value))
}

// GetDS_secretary gets the value of DS_secretary for the instance
func (instance *ads_contact) GetPropertyDS_secretary()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_secretary")
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

// SetDS_showInAddressBook sets the value of DS_showInAddressBook for the instance
func (instance *ads_contact) SetPropertyDS_showInAddressBook(value []string) (err error) { 
    return instance.SetProperty("DS_showInAddressBook", (value))
}

// GetDS_showInAddressBook gets the value of DS_showInAddressBook for the instance
func (instance *ads_contact) GetPropertyDS_showInAddressBook()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_showInAddressBook")
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

// SetDS_textEncodedORAddress sets the value of DS_textEncodedORAddress for the instance
func (instance *ads_contact) SetPropertyDS_textEncodedORAddress(value string) (err error) { 
    return instance.SetProperty("DS_textEncodedORAddress", (value))
}

// GetDS_textEncodedORAddress gets the value of DS_textEncodedORAddress for the instance
func (instance *ads_contact) GetPropertyDS_textEncodedORAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_textEncodedORAddress")
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

// SetDS_userCert sets the value of DS_userCert for the instance
func (instance *ads_contact) SetPropertyDS_userCert(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_userCert", (value))
}

// GetDS_userCert gets the value of DS_userCert for the instance
func (instance *ads_contact) GetPropertyDS_userCert()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_userCert")
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

// SetDS_userCertificate sets the value of DS_userCertificate for the instance
func (instance *ads_contact) SetPropertyDS_userCertificate(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_userCertificate", (value))
}

// GetDS_userCertificate gets the value of DS_userCertificate for the instance
func (instance *ads_contact) GetPropertyDS_userCertificate()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_userCertificate")
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

// SetDS_userSMIMECertificate sets the value of DS_userSMIMECertificate for the instance
func (instance *ads_contact) SetPropertyDS_userSMIMECertificate(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_userSMIMECertificate", (value))
}

// GetDS_userSMIMECertificate gets the value of DS_userSMIMECertificate for the instance
func (instance *ads_contact) GetPropertyDS_userSMIMECertificate()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_userSMIMECertificate")
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

