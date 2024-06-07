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

// ads_organizationalunit struct
type ads_organizationalunit struct { 
	*ds_top

	// 
	DS_businessCategory []string

	// 
	DS_c string

	// 
	DS_co string

	// 
	DS_countryCode int32

	// 
	DS_defaultGroup string

	// 
	DS_desktopProfile string

	// 
	DS_destinationIndicator []string

	// 
	DS_facsimileTelephoneNumber string

	// 
	DS_gPLink string

	// 
	DS_gPOptions int32

	// 
	DS_internationalISDNNumber []string

	// 
	DS_l string

	// 
	DS_managedBy string

	// 
	DS_msCOM_UserPartitionSetLink string

	// 
	DS_ou []string

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
	DS_thumbnailLogo Uint8Array

	// 
	DS_uPNSuffixes []string

	// 
	DS_userPassword []Uint8Array

	// 
	DS_x121Address []string
}

	func Newads_organizationalunitEx1(instance *cim.WmiInstance) (newInstance *ads_organizationalunit, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_organizationalunit {
	ds_top: tmp,
	}
	return
	}
	

	func Newads_organizationalunitEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_organizationalunit, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_organizationalunit {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_businessCategory sets the value of DS_businessCategory for the instance
func (instance *ads_organizationalunit) SetPropertyDS_businessCategory(value []string) (err error) { 
    return instance.SetProperty("DS_businessCategory", (value))
}

// GetDS_businessCategory gets the value of DS_businessCategory for the instance
func (instance *ads_organizationalunit) GetPropertyDS_businessCategory()(value []string, err error) { 
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

// SetDS_c sets the value of DS_c for the instance
func (instance *ads_organizationalunit) SetPropertyDS_c(value string) (err error) { 
    return instance.SetProperty("DS_c", (value))
}

// GetDS_c gets the value of DS_c for the instance
func (instance *ads_organizationalunit) GetPropertyDS_c()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_c")
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

// SetDS_co sets the value of DS_co for the instance
func (instance *ads_organizationalunit) SetPropertyDS_co(value string) (err error) { 
    return instance.SetProperty("DS_co", (value))
}

// GetDS_co gets the value of DS_co for the instance
func (instance *ads_organizationalunit) GetPropertyDS_co()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_co")
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

// SetDS_countryCode sets the value of DS_countryCode for the instance
func (instance *ads_organizationalunit) SetPropertyDS_countryCode(value int32) (err error) { 
    return instance.SetProperty("DS_countryCode", (value))
}

// GetDS_countryCode gets the value of DS_countryCode for the instance
func (instance *ads_organizationalunit) GetPropertyDS_countryCode()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_countryCode")
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

// SetDS_defaultGroup sets the value of DS_defaultGroup for the instance
func (instance *ads_organizationalunit) SetPropertyDS_defaultGroup(value string) (err error) { 
    return instance.SetProperty("DS_defaultGroup", (value))
}

// GetDS_defaultGroup gets the value of DS_defaultGroup for the instance
func (instance *ads_organizationalunit) GetPropertyDS_defaultGroup()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_defaultGroup")
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

// SetDS_desktopProfile sets the value of DS_desktopProfile for the instance
func (instance *ads_organizationalunit) SetPropertyDS_desktopProfile(value string) (err error) { 
    return instance.SetProperty("DS_desktopProfile", (value))
}

// GetDS_desktopProfile gets the value of DS_desktopProfile for the instance
func (instance *ads_organizationalunit) GetPropertyDS_desktopProfile()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_desktopProfile")
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

// SetDS_destinationIndicator sets the value of DS_destinationIndicator for the instance
func (instance *ads_organizationalunit) SetPropertyDS_destinationIndicator(value []string) (err error) { 
    return instance.SetProperty("DS_destinationIndicator", (value))
}

// GetDS_destinationIndicator gets the value of DS_destinationIndicator for the instance
func (instance *ads_organizationalunit) GetPropertyDS_destinationIndicator()(value []string, err error) { 
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
func (instance *ads_organizationalunit) SetPropertyDS_facsimileTelephoneNumber(value string) (err error) { 
    return instance.SetProperty("DS_facsimileTelephoneNumber", (value))
}

// GetDS_facsimileTelephoneNumber gets the value of DS_facsimileTelephoneNumber for the instance
func (instance *ads_organizationalunit) GetPropertyDS_facsimileTelephoneNumber()(value string, err error) { 
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

// SetDS_gPLink sets the value of DS_gPLink for the instance
func (instance *ads_organizationalunit) SetPropertyDS_gPLink(value string) (err error) { 
    return instance.SetProperty("DS_gPLink", (value))
}

// GetDS_gPLink gets the value of DS_gPLink for the instance
func (instance *ads_organizationalunit) GetPropertyDS_gPLink()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_gPLink")
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

// SetDS_gPOptions sets the value of DS_gPOptions for the instance
func (instance *ads_organizationalunit) SetPropertyDS_gPOptions(value int32) (err error) { 
    return instance.SetProperty("DS_gPOptions", (value))
}

// GetDS_gPOptions gets the value of DS_gPOptions for the instance
func (instance *ads_organizationalunit) GetPropertyDS_gPOptions()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_gPOptions")
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

// SetDS_internationalISDNNumber sets the value of DS_internationalISDNNumber for the instance
func (instance *ads_organizationalunit) SetPropertyDS_internationalISDNNumber(value []string) (err error) { 
    return instance.SetProperty("DS_internationalISDNNumber", (value))
}

// GetDS_internationalISDNNumber gets the value of DS_internationalISDNNumber for the instance
func (instance *ads_organizationalunit) GetPropertyDS_internationalISDNNumber()(value []string, err error) { 
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
func (instance *ads_organizationalunit) SetPropertyDS_l(value string) (err error) { 
    return instance.SetProperty("DS_l", (value))
}

// GetDS_l gets the value of DS_l for the instance
func (instance *ads_organizationalunit) GetPropertyDS_l()(value string, err error) { 
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

// SetDS_managedBy sets the value of DS_managedBy for the instance
func (instance *ads_organizationalunit) SetPropertyDS_managedBy(value string) (err error) { 
    return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_organizationalunit) GetPropertyDS_managedBy()(value string, err error) { 
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

// SetDS_msCOM_UserPartitionSetLink sets the value of DS_msCOM_UserPartitionSetLink for the instance
func (instance *ads_organizationalunit) SetPropertyDS_msCOM_UserPartitionSetLink(value string) (err error) { 
    return instance.SetProperty("DS_msCOM_UserPartitionSetLink", (value))
}

// GetDS_msCOM_UserPartitionSetLink gets the value of DS_msCOM_UserPartitionSetLink for the instance
func (instance *ads_organizationalunit) GetPropertyDS_msCOM_UserPartitionSetLink()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msCOM_UserPartitionSetLink")
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

// SetDS_ou sets the value of DS_ou for the instance
func (instance *ads_organizationalunit) SetPropertyDS_ou(value []string) (err error) { 
    return instance.SetProperty("DS_ou", (value))
}

// GetDS_ou gets the value of DS_ou for the instance
func (instance *ads_organizationalunit) GetPropertyDS_ou()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_ou")
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
func (instance *ads_organizationalunit) SetPropertyDS_physicalDeliveryOfficeName(value string) (err error) { 
    return instance.SetProperty("DS_physicalDeliveryOfficeName", (value))
}

// GetDS_physicalDeliveryOfficeName gets the value of DS_physicalDeliveryOfficeName for the instance
func (instance *ads_organizationalunit) GetPropertyDS_physicalDeliveryOfficeName()(value string, err error) { 
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
func (instance *ads_organizationalunit) SetPropertyDS_postalAddress(value []string) (err error) { 
    return instance.SetProperty("DS_postalAddress", (value))
}

// GetDS_postalAddress gets the value of DS_postalAddress for the instance
func (instance *ads_organizationalunit) GetPropertyDS_postalAddress()(value []string, err error) { 
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
func (instance *ads_organizationalunit) SetPropertyDS_postalCode(value string) (err error) { 
    return instance.SetProperty("DS_postalCode", (value))
}

// GetDS_postalCode gets the value of DS_postalCode for the instance
func (instance *ads_organizationalunit) GetPropertyDS_postalCode()(value string, err error) { 
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
func (instance *ads_organizationalunit) SetPropertyDS_postOfficeBox(value []string) (err error) { 
    return instance.SetProperty("DS_postOfficeBox", (value))
}

// GetDS_postOfficeBox gets the value of DS_postOfficeBox for the instance
func (instance *ads_organizationalunit) GetPropertyDS_postOfficeBox()(value []string, err error) { 
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
func (instance *ads_organizationalunit) SetPropertyDS_preferredDeliveryMethod(value []int32) (err error) { 
    return instance.SetProperty("DS_preferredDeliveryMethod", (value))
}

// GetDS_preferredDeliveryMethod gets the value of DS_preferredDeliveryMethod for the instance
func (instance *ads_organizationalunit) GetPropertyDS_preferredDeliveryMethod()(value []int32, err error) { 
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
func (instance *ads_organizationalunit) SetPropertyDS_registeredAddress(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_registeredAddress", (value))
}

// GetDS_registeredAddress gets the value of DS_registeredAddress for the instance
func (instance *ads_organizationalunit) GetPropertyDS_registeredAddress()(value []Uint8Array, err error) { 
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
func (instance *ads_organizationalunit) SetPropertyDS_searchGuide(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_searchGuide", (value))
}

// GetDS_searchGuide gets the value of DS_searchGuide for the instance
func (instance *ads_organizationalunit) GetPropertyDS_searchGuide()(value []Uint8Array, err error) { 
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
func (instance *ads_organizationalunit) SetPropertyDS_seeAlso(value []string) (err error) { 
    return instance.SetProperty("DS_seeAlso", (value))
}

// GetDS_seeAlso gets the value of DS_seeAlso for the instance
func (instance *ads_organizationalunit) GetPropertyDS_seeAlso()(value []string, err error) { 
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
func (instance *ads_organizationalunit) SetPropertyDS_st(value string) (err error) { 
    return instance.SetProperty("DS_st", (value))
}

// GetDS_st gets the value of DS_st for the instance
func (instance *ads_organizationalunit) GetPropertyDS_st()(value string, err error) { 
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
func (instance *ads_organizationalunit) SetPropertyDS_street(value string) (err error) { 
    return instance.SetProperty("DS_street", (value))
}

// GetDS_street gets the value of DS_street for the instance
func (instance *ads_organizationalunit) GetPropertyDS_street()(value string, err error) { 
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
func (instance *ads_organizationalunit) SetPropertyDS_telephoneNumber(value string) (err error) { 
    return instance.SetProperty("DS_telephoneNumber", (value))
}

// GetDS_telephoneNumber gets the value of DS_telephoneNumber for the instance
func (instance *ads_organizationalunit) GetPropertyDS_telephoneNumber()(value string, err error) { 
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
func (instance *ads_organizationalunit) SetPropertyDS_teletexTerminalIdentifier(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_teletexTerminalIdentifier", (value))
}

// GetDS_teletexTerminalIdentifier gets the value of DS_teletexTerminalIdentifier for the instance
func (instance *ads_organizationalunit) GetPropertyDS_teletexTerminalIdentifier()(value []Uint8Array, err error) { 
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
func (instance *ads_organizationalunit) SetPropertyDS_telexNumber(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_telexNumber", (value))
}

// GetDS_telexNumber gets the value of DS_telexNumber for the instance
func (instance *ads_organizationalunit) GetPropertyDS_telexNumber()(value []Uint8Array, err error) { 
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

// SetDS_thumbnailLogo sets the value of DS_thumbnailLogo for the instance
func (instance *ads_organizationalunit) SetPropertyDS_thumbnailLogo(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_thumbnailLogo", (value))
}

// GetDS_thumbnailLogo gets the value of DS_thumbnailLogo for the instance
func (instance *ads_organizationalunit) GetPropertyDS_thumbnailLogo()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_thumbnailLogo")
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

// SetDS_uPNSuffixes sets the value of DS_uPNSuffixes for the instance
func (instance *ads_organizationalunit) SetPropertyDS_uPNSuffixes(value []string) (err error) { 
    return instance.SetProperty("DS_uPNSuffixes", (value))
}

// GetDS_uPNSuffixes gets the value of DS_uPNSuffixes for the instance
func (instance *ads_organizationalunit) GetPropertyDS_uPNSuffixes()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_uPNSuffixes")
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

// SetDS_userPassword sets the value of DS_userPassword for the instance
func (instance *ads_organizationalunit) SetPropertyDS_userPassword(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_userPassword", (value))
}

// GetDS_userPassword gets the value of DS_userPassword for the instance
func (instance *ads_organizationalunit) GetPropertyDS_userPassword()(value []Uint8Array, err error) { 
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
func (instance *ads_organizationalunit) SetPropertyDS_x121Address(value []string) (err error) { 
    return instance.SetProperty("DS_x121Address", (value))
}

// GetDS_x121Address gets the value of DS_x121Address for the instance
func (instance *ads_organizationalunit) GetPropertyDS_x121Address()(value []string, err error) { 
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

