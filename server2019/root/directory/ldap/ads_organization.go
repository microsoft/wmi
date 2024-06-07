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

// ads_organization struct
type ads_organization struct { 
	*ds_top

	// 
	DS_businessCategory []string

	// 
	DS_destinationIndicator []string

	// 
	DS_facsimileTelephoneNumber string

	// 
	DS_internationalISDNNumber []string

	// 
	DS_l string

	// 
	DS_o []string

	// 
	DS_physicalDeliveryOfficeName string

	// 
	DS_postalAddress []string

	// 
	DS_postalCode string

	// 
	DS_postOfficeBox []string

	// 
	DS_preferredDeliveryMethod []int32

	// 
	DS_registeredAddress []Uint8Array

	// 
	DS_searchGuide []Uint8Array

	// 
	DS_seeAlso []string

	// 
	DS_st string

	// 
	DS_street string

	// 
	DS_telephoneNumber string

	// 
	DS_teletexTerminalIdentifier []Uint8Array

	// 
	DS_telexNumber []Uint8Array

	// 
	DS_userPassword []Uint8Array

	// 
	DS_x121Address []string
}

	func Newads_organizationEx1(instance *cim.WmiInstance) (newInstance *ads_organization, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_organization {
	ds_top: tmp,
	}
	return
	}
	

	func Newads_organizationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_organization, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_organization {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_businessCategory sets the value of DS_businessCategory for the instance
func (instance *ads_organization) SetPropertyDS_businessCategory(value []string) (err error) { 
    return instance.SetProperty("DS_businessCategory", (value))
}

// GetDS_businessCategory gets the value of DS_businessCategory for the instance
func (instance *ads_organization) GetPropertyDS_businessCategory()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_businessCategory")
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

// SetDS_destinationIndicator sets the value of DS_destinationIndicator for the instance
func (instance *ads_organization) SetPropertyDS_destinationIndicator(value []string) (err error) { 
    return instance.SetProperty("DS_destinationIndicator", (value))
}

// GetDS_destinationIndicator gets the value of DS_destinationIndicator for the instance
func (instance *ads_organization) GetPropertyDS_destinationIndicator()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_destinationIndicator")
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

// SetDS_facsimileTelephoneNumber sets the value of DS_facsimileTelephoneNumber for the instance
func (instance *ads_organization) SetPropertyDS_facsimileTelephoneNumber(value string) (err error) { 
    return instance.SetProperty("DS_facsimileTelephoneNumber", (value))
}

// GetDS_facsimileTelephoneNumber gets the value of DS_facsimileTelephoneNumber for the instance
func (instance *ads_organization) GetPropertyDS_facsimileTelephoneNumber()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_facsimileTelephoneNumber")
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

// SetDS_internationalISDNNumber sets the value of DS_internationalISDNNumber for the instance
func (instance *ads_organization) SetPropertyDS_internationalISDNNumber(value []string) (err error) { 
    return instance.SetProperty("DS_internationalISDNNumber", (value))
}

// GetDS_internationalISDNNumber gets the value of DS_internationalISDNNumber for the instance
func (instance *ads_organization) GetPropertyDS_internationalISDNNumber()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_internationalISDNNumber")
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

// SetDS_l sets the value of DS_l for the instance
func (instance *ads_organization) SetPropertyDS_l(value string) (err error) { 
    return instance.SetProperty("DS_l", (value))
}

// GetDS_l gets the value of DS_l for the instance
func (instance *ads_organization) GetPropertyDS_l()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_l")
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

// SetDS_o sets the value of DS_o for the instance
func (instance *ads_organization) SetPropertyDS_o(value []string) (err error) { 
    return instance.SetProperty("DS_o", (value))
}

// GetDS_o gets the value of DS_o for the instance
func (instance *ads_organization) GetPropertyDS_o()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_o")
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

// SetDS_physicalDeliveryOfficeName sets the value of DS_physicalDeliveryOfficeName for the instance
func (instance *ads_organization) SetPropertyDS_physicalDeliveryOfficeName(value string) (err error) { 
    return instance.SetProperty("DS_physicalDeliveryOfficeName", (value))
}

// GetDS_physicalDeliveryOfficeName gets the value of DS_physicalDeliveryOfficeName for the instance
func (instance *ads_organization) GetPropertyDS_physicalDeliveryOfficeName()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_physicalDeliveryOfficeName")
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

// SetDS_postalAddress sets the value of DS_postalAddress for the instance
func (instance *ads_organization) SetPropertyDS_postalAddress(value []string) (err error) { 
    return instance.SetProperty("DS_postalAddress", (value))
}

// GetDS_postalAddress gets the value of DS_postalAddress for the instance
func (instance *ads_organization) GetPropertyDS_postalAddress()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_postalAddress")
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

// SetDS_postalCode sets the value of DS_postalCode for the instance
func (instance *ads_organization) SetPropertyDS_postalCode(value string) (err error) { 
    return instance.SetProperty("DS_postalCode", (value))
}

// GetDS_postalCode gets the value of DS_postalCode for the instance
func (instance *ads_organization) GetPropertyDS_postalCode()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_postalCode")
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

// SetDS_postOfficeBox sets the value of DS_postOfficeBox for the instance
func (instance *ads_organization) SetPropertyDS_postOfficeBox(value []string) (err error) { 
    return instance.SetProperty("DS_postOfficeBox", (value))
}

// GetDS_postOfficeBox gets the value of DS_postOfficeBox for the instance
func (instance *ads_organization) GetPropertyDS_postOfficeBox()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_postOfficeBox")
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

// SetDS_preferredDeliveryMethod sets the value of DS_preferredDeliveryMethod for the instance
func (instance *ads_organization) SetPropertyDS_preferredDeliveryMethod(value []int32) (err error) { 
    return instance.SetProperty("DS_preferredDeliveryMethod", (value))
}

// GetDS_preferredDeliveryMethod gets the value of DS_preferredDeliveryMethod for the instance
func (instance *ads_organization) GetPropertyDS_preferredDeliveryMethod()(value []int32, err error) { 
    retValue, err := instance.GetProperty("DS_preferredDeliveryMethod")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(int32); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, int32(valuetmp))
    }

    return
}

// SetDS_registeredAddress sets the value of DS_registeredAddress for the instance
func (instance *ads_organization) SetPropertyDS_registeredAddress(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_registeredAddress", (value))
}

// GetDS_registeredAddress gets the value of DS_registeredAddress for the instance
func (instance *ads_organization) GetPropertyDS_registeredAddress()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_registeredAddress")
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

// SetDS_searchGuide sets the value of DS_searchGuide for the instance
func (instance *ads_organization) SetPropertyDS_searchGuide(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_searchGuide", (value))
}

// GetDS_searchGuide gets the value of DS_searchGuide for the instance
func (instance *ads_organization) GetPropertyDS_searchGuide()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_searchGuide")
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

// SetDS_seeAlso sets the value of DS_seeAlso for the instance
func (instance *ads_organization) SetPropertyDS_seeAlso(value []string) (err error) { 
    return instance.SetProperty("DS_seeAlso", (value))
}

// GetDS_seeAlso gets the value of DS_seeAlso for the instance
func (instance *ads_organization) GetPropertyDS_seeAlso()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_seeAlso")
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

// SetDS_st sets the value of DS_st for the instance
func (instance *ads_organization) SetPropertyDS_st(value string) (err error) { 
    return instance.SetProperty("DS_st", (value))
}

// GetDS_st gets the value of DS_st for the instance
func (instance *ads_organization) GetPropertyDS_st()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_st")
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

// SetDS_street sets the value of DS_street for the instance
func (instance *ads_organization) SetPropertyDS_street(value string) (err error) { 
    return instance.SetProperty("DS_street", (value))
}

// GetDS_street gets the value of DS_street for the instance
func (instance *ads_organization) GetPropertyDS_street()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_street")
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

// SetDS_telephoneNumber sets the value of DS_telephoneNumber for the instance
func (instance *ads_organization) SetPropertyDS_telephoneNumber(value string) (err error) { 
    return instance.SetProperty("DS_telephoneNumber", (value))
}

// GetDS_telephoneNumber gets the value of DS_telephoneNumber for the instance
func (instance *ads_organization) GetPropertyDS_telephoneNumber()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_telephoneNumber")
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

// SetDS_teletexTerminalIdentifier sets the value of DS_teletexTerminalIdentifier for the instance
func (instance *ads_organization) SetPropertyDS_teletexTerminalIdentifier(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_teletexTerminalIdentifier", (value))
}

// GetDS_teletexTerminalIdentifier gets the value of DS_teletexTerminalIdentifier for the instance
func (instance *ads_organization) GetPropertyDS_teletexTerminalIdentifier()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_teletexTerminalIdentifier")
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

// SetDS_telexNumber sets the value of DS_telexNumber for the instance
func (instance *ads_organization) SetPropertyDS_telexNumber(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_telexNumber", (value))
}

// GetDS_telexNumber gets the value of DS_telexNumber for the instance
func (instance *ads_organization) GetPropertyDS_telexNumber()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_telexNumber")
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

// SetDS_userPassword sets the value of DS_userPassword for the instance
func (instance *ads_organization) SetPropertyDS_userPassword(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_userPassword", (value))
}

// GetDS_userPassword gets the value of DS_userPassword for the instance
func (instance *ads_organization) GetPropertyDS_userPassword()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_userPassword")
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

// SetDS_x121Address sets the value of DS_x121Address for the instance
func (instance *ads_organization) SetPropertyDS_x121Address(value []string) (err error) { 
    return instance.SetProperty("DS_x121Address", (value))
}

// GetDS_x121Address gets the value of DS_x121Address for the instance
func (instance *ads_organization) GetPropertyDS_x121Address()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_x121Address")
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

