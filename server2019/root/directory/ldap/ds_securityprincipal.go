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

// ds_securityprincipal struct
type ds_securityprincipal struct { 
	*ds_top

	// 
	DS_accountNameHistory []string

	// 
	DS_altSecurityIdentities []string

	// 
	DS_msDS_KeyVersionNumber int32

	// 
	DS_msds_tokenGroupNames []string

	// 
	DS_msds_tokenGroupNamesGlobalAndUniversal []string

	// 
	DS_msds_tokenGroupNamesNoGCAcceptable []string

	// 
	DS_objectSid Uint8Array

	// 
	DS_rid int32

	// 
	DS_sAMAccountName string

	// 
	DS_sAMAccountType int32

	// 
	DS_securityIdentifier Uint8Array

	// 
	DS_sIDHistory []Uint8Array

	// 
	DS_supplementalCredentials []Uint8Array

	// 
	DS_tokenGroups []Uint8Array

	// 
	DS_tokenGroupsGlobalAndUniversal []Uint8Array

	// 
	DS_tokenGroupsNoGCAcceptable []Uint8Array
}

	func Newds_securityprincipalEx1(instance *cim.WmiInstance) (newInstance *ds_securityprincipal, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ds_securityprincipal {
	ds_top: tmp,
	}
	return
	}
	

	func Newds_securityprincipalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ds_securityprincipal, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ds_securityprincipal {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_accountNameHistory sets the value of DS_accountNameHistory for the instance
func (instance *ds_securityprincipal) SetPropertyDS_accountNameHistory(value []string) (err error) { 
    return instance.SetProperty("DS_accountNameHistory", (value))
}

// GetDS_accountNameHistory gets the value of DS_accountNameHistory for the instance
func (instance *ds_securityprincipal) GetPropertyDS_accountNameHistory()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_accountNameHistory")
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

// SetDS_altSecurityIdentities sets the value of DS_altSecurityIdentities for the instance
func (instance *ds_securityprincipal) SetPropertyDS_altSecurityIdentities(value []string) (err error) { 
    return instance.SetProperty("DS_altSecurityIdentities", (value))
}

// GetDS_altSecurityIdentities gets the value of DS_altSecurityIdentities for the instance
func (instance *ds_securityprincipal) GetPropertyDS_altSecurityIdentities()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_altSecurityIdentities")
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

// SetDS_msDS_KeyVersionNumber sets the value of DS_msDS_KeyVersionNumber for the instance
func (instance *ds_securityprincipal) SetPropertyDS_msDS_KeyVersionNumber(value int32) (err error) { 
    return instance.SetProperty("DS_msDS_KeyVersionNumber", (value))
}

// GetDS_msDS_KeyVersionNumber gets the value of DS_msDS_KeyVersionNumber for the instance
func (instance *ds_securityprincipal) GetPropertyDS_msDS_KeyVersionNumber()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_KeyVersionNumber")
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

// SetDS_msds_tokenGroupNames sets the value of DS_msds_tokenGroupNames for the instance
func (instance *ds_securityprincipal) SetPropertyDS_msds_tokenGroupNames(value []string) (err error) { 
    return instance.SetProperty("DS_msds_tokenGroupNames", (value))
}

// GetDS_msds_tokenGroupNames gets the value of DS_msds_tokenGroupNames for the instance
func (instance *ds_securityprincipal) GetPropertyDS_msds_tokenGroupNames()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_msds_tokenGroupNames")
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

// SetDS_msds_tokenGroupNamesGlobalAndUniversal sets the value of DS_msds_tokenGroupNamesGlobalAndUniversal for the instance
func (instance *ds_securityprincipal) SetPropertyDS_msds_tokenGroupNamesGlobalAndUniversal(value []string) (err error) { 
    return instance.SetProperty("DS_msds_tokenGroupNamesGlobalAndUniversal", (value))
}

// GetDS_msds_tokenGroupNamesGlobalAndUniversal gets the value of DS_msds_tokenGroupNamesGlobalAndUniversal for the instance
func (instance *ds_securityprincipal) GetPropertyDS_msds_tokenGroupNamesGlobalAndUniversal()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_msds_tokenGroupNamesGlobalAndUniversal")
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

// SetDS_msds_tokenGroupNamesNoGCAcceptable sets the value of DS_msds_tokenGroupNamesNoGCAcceptable for the instance
func (instance *ds_securityprincipal) SetPropertyDS_msds_tokenGroupNamesNoGCAcceptable(value []string) (err error) { 
    return instance.SetProperty("DS_msds_tokenGroupNamesNoGCAcceptable", (value))
}

// GetDS_msds_tokenGroupNamesNoGCAcceptable gets the value of DS_msds_tokenGroupNamesNoGCAcceptable for the instance
func (instance *ds_securityprincipal) GetPropertyDS_msds_tokenGroupNamesNoGCAcceptable()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_msds_tokenGroupNamesNoGCAcceptable")
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

// SetDS_objectSid sets the value of DS_objectSid for the instance
func (instance *ds_securityprincipal) SetPropertyDS_objectSid(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_objectSid", (value))
}

// GetDS_objectSid gets the value of DS_objectSid for the instance
func (instance *ds_securityprincipal) GetPropertyDS_objectSid()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_objectSid")
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

// SetDS_rid sets the value of DS_rid for the instance
func (instance *ds_securityprincipal) SetPropertyDS_rid(value int32) (err error) { 
    return instance.SetProperty("DS_rid", (value))
}

// GetDS_rid gets the value of DS_rid for the instance
func (instance *ds_securityprincipal) GetPropertyDS_rid()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_rid")
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

// SetDS_sAMAccountName sets the value of DS_sAMAccountName for the instance
func (instance *ds_securityprincipal) SetPropertyDS_sAMAccountName(value string) (err error) { 
    return instance.SetProperty("DS_sAMAccountName", (value))
}

// GetDS_sAMAccountName gets the value of DS_sAMAccountName for the instance
func (instance *ds_securityprincipal) GetPropertyDS_sAMAccountName()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_sAMAccountName")
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

// SetDS_sAMAccountType sets the value of DS_sAMAccountType for the instance
func (instance *ds_securityprincipal) SetPropertyDS_sAMAccountType(value int32) (err error) { 
    return instance.SetProperty("DS_sAMAccountType", (value))
}

// GetDS_sAMAccountType gets the value of DS_sAMAccountType for the instance
func (instance *ds_securityprincipal) GetPropertyDS_sAMAccountType()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_sAMAccountType")
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

// SetDS_securityIdentifier sets the value of DS_securityIdentifier for the instance
func (instance *ds_securityprincipal) SetPropertyDS_securityIdentifier(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_securityIdentifier", (value))
}

// GetDS_securityIdentifier gets the value of DS_securityIdentifier for the instance
func (instance *ds_securityprincipal) GetPropertyDS_securityIdentifier()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_securityIdentifier")
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

// SetDS_sIDHistory sets the value of DS_sIDHistory for the instance
func (instance *ds_securityprincipal) SetPropertyDS_sIDHistory(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_sIDHistory", (value))
}

// GetDS_sIDHistory gets the value of DS_sIDHistory for the instance
func (instance *ds_securityprincipal) GetPropertyDS_sIDHistory()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_sIDHistory")
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

// SetDS_supplementalCredentials sets the value of DS_supplementalCredentials for the instance
func (instance *ds_securityprincipal) SetPropertyDS_supplementalCredentials(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_supplementalCredentials", (value))
}

// GetDS_supplementalCredentials gets the value of DS_supplementalCredentials for the instance
func (instance *ds_securityprincipal) GetPropertyDS_supplementalCredentials()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_supplementalCredentials")
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

// SetDS_tokenGroups sets the value of DS_tokenGroups for the instance
func (instance *ds_securityprincipal) SetPropertyDS_tokenGroups(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_tokenGroups", (value))
}

// GetDS_tokenGroups gets the value of DS_tokenGroups for the instance
func (instance *ds_securityprincipal) GetPropertyDS_tokenGroups()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_tokenGroups")
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

// SetDS_tokenGroupsGlobalAndUniversal sets the value of DS_tokenGroupsGlobalAndUniversal for the instance
func (instance *ds_securityprincipal) SetPropertyDS_tokenGroupsGlobalAndUniversal(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_tokenGroupsGlobalAndUniversal", (value))
}

// GetDS_tokenGroupsGlobalAndUniversal gets the value of DS_tokenGroupsGlobalAndUniversal for the instance
func (instance *ds_securityprincipal) GetPropertyDS_tokenGroupsGlobalAndUniversal()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_tokenGroupsGlobalAndUniversal")
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

// SetDS_tokenGroupsNoGCAcceptable sets the value of DS_tokenGroupsNoGCAcceptable for the instance
func (instance *ds_securityprincipal) SetPropertyDS_tokenGroupsNoGCAcceptable(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_tokenGroupsNoGCAcceptable", (value))
}

// GetDS_tokenGroupsNoGCAcceptable gets the value of DS_tokenGroupsNoGCAcceptable for the instance
func (instance *ds_securityprincipal) GetPropertyDS_tokenGroupsNoGCAcceptable()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_tokenGroupsNoGCAcceptable")
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

