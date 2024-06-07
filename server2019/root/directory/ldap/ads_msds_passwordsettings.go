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

// ads_msds_passwordsettings struct
type ads_msds_passwordsettings struct { 
	*ds_top

	// 
	DS_msDS_LockoutDuration int64

	// 
	DS_msDS_LockoutObservationWindow int64

	// 
	DS_msDS_LockoutThreshold int32

	// 
	DS_msDS_MaximumPasswordAge int64

	// 
	DS_msDS_MinimumPasswordAge int64

	// 
	DS_msDS_MinimumPasswordLength int32

	// 
	DS_msDS_PasswordComplexityEnabled bool

	// 
	DS_msDS_PasswordHistoryLength int32

	// 
	DS_msDS_PasswordReversibleEncryptionEnabled bool

	// 
	DS_msDS_PasswordSettingsPrecedence int32

	// 
	DS_msDS_PSOAppliesTo []string
}

	func Newads_msds_passwordsettingsEx1(instance *cim.WmiInstance) (newInstance *ads_msds_passwordsettings, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_msds_passwordsettings {
	ds_top: tmp,
	}
	return
	}
	

	func Newads_msds_passwordsettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_msds_passwordsettings, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_msds_passwordsettings {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_msDS_LockoutDuration sets the value of DS_msDS_LockoutDuration for the instance
func (instance *ads_msds_passwordsettings) SetPropertyDS_msDS_LockoutDuration(value int64) (err error) { 
    return instance.SetProperty("DS_msDS_LockoutDuration", (value))
}

// GetDS_msDS_LockoutDuration gets the value of DS_msDS_LockoutDuration for the instance
func (instance *ads_msds_passwordsettings) GetPropertyDS_msDS_LockoutDuration()(value int64, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_LockoutDuration")
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

// SetDS_msDS_LockoutObservationWindow sets the value of DS_msDS_LockoutObservationWindow for the instance
func (instance *ads_msds_passwordsettings) SetPropertyDS_msDS_LockoutObservationWindow(value int64) (err error) { 
    return instance.SetProperty("DS_msDS_LockoutObservationWindow", (value))
}

// GetDS_msDS_LockoutObservationWindow gets the value of DS_msDS_LockoutObservationWindow for the instance
func (instance *ads_msds_passwordsettings) GetPropertyDS_msDS_LockoutObservationWindow()(value int64, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_LockoutObservationWindow")
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

// SetDS_msDS_LockoutThreshold sets the value of DS_msDS_LockoutThreshold for the instance
func (instance *ads_msds_passwordsettings) SetPropertyDS_msDS_LockoutThreshold(value int32) (err error) { 
    return instance.SetProperty("DS_msDS_LockoutThreshold", (value))
}

// GetDS_msDS_LockoutThreshold gets the value of DS_msDS_LockoutThreshold for the instance
func (instance *ads_msds_passwordsettings) GetPropertyDS_msDS_LockoutThreshold()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_LockoutThreshold")
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

// SetDS_msDS_MaximumPasswordAge sets the value of DS_msDS_MaximumPasswordAge for the instance
func (instance *ads_msds_passwordsettings) SetPropertyDS_msDS_MaximumPasswordAge(value int64) (err error) { 
    return instance.SetProperty("DS_msDS_MaximumPasswordAge", (value))
}

// GetDS_msDS_MaximumPasswordAge gets the value of DS_msDS_MaximumPasswordAge for the instance
func (instance *ads_msds_passwordsettings) GetPropertyDS_msDS_MaximumPasswordAge()(value int64, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_MaximumPasswordAge")
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

// SetDS_msDS_MinimumPasswordAge sets the value of DS_msDS_MinimumPasswordAge for the instance
func (instance *ads_msds_passwordsettings) SetPropertyDS_msDS_MinimumPasswordAge(value int64) (err error) { 
    return instance.SetProperty("DS_msDS_MinimumPasswordAge", (value))
}

// GetDS_msDS_MinimumPasswordAge gets the value of DS_msDS_MinimumPasswordAge for the instance
func (instance *ads_msds_passwordsettings) GetPropertyDS_msDS_MinimumPasswordAge()(value int64, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_MinimumPasswordAge")
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

// SetDS_msDS_MinimumPasswordLength sets the value of DS_msDS_MinimumPasswordLength for the instance
func (instance *ads_msds_passwordsettings) SetPropertyDS_msDS_MinimumPasswordLength(value int32) (err error) { 
    return instance.SetProperty("DS_msDS_MinimumPasswordLength", (value))
}

// GetDS_msDS_MinimumPasswordLength gets the value of DS_msDS_MinimumPasswordLength for the instance
func (instance *ads_msds_passwordsettings) GetPropertyDS_msDS_MinimumPasswordLength()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_MinimumPasswordLength")
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

// SetDS_msDS_PasswordComplexityEnabled sets the value of DS_msDS_PasswordComplexityEnabled for the instance
func (instance *ads_msds_passwordsettings) SetPropertyDS_msDS_PasswordComplexityEnabled(value bool) (err error) { 
    return instance.SetProperty("DS_msDS_PasswordComplexityEnabled", (value))
}

// GetDS_msDS_PasswordComplexityEnabled gets the value of DS_msDS_PasswordComplexityEnabled for the instance
func (instance *ads_msds_passwordsettings) GetPropertyDS_msDS_PasswordComplexityEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_PasswordComplexityEnabled")
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

// SetDS_msDS_PasswordHistoryLength sets the value of DS_msDS_PasswordHistoryLength for the instance
func (instance *ads_msds_passwordsettings) SetPropertyDS_msDS_PasswordHistoryLength(value int32) (err error) { 
    return instance.SetProperty("DS_msDS_PasswordHistoryLength", (value))
}

// GetDS_msDS_PasswordHistoryLength gets the value of DS_msDS_PasswordHistoryLength for the instance
func (instance *ads_msds_passwordsettings) GetPropertyDS_msDS_PasswordHistoryLength()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_PasswordHistoryLength")
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

// SetDS_msDS_PasswordReversibleEncryptionEnabled sets the value of DS_msDS_PasswordReversibleEncryptionEnabled for the instance
func (instance *ads_msds_passwordsettings) SetPropertyDS_msDS_PasswordReversibleEncryptionEnabled(value bool) (err error) { 
    return instance.SetProperty("DS_msDS_PasswordReversibleEncryptionEnabled", (value))
}

// GetDS_msDS_PasswordReversibleEncryptionEnabled gets the value of DS_msDS_PasswordReversibleEncryptionEnabled for the instance
func (instance *ads_msds_passwordsettings) GetPropertyDS_msDS_PasswordReversibleEncryptionEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_PasswordReversibleEncryptionEnabled")
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

// SetDS_msDS_PasswordSettingsPrecedence sets the value of DS_msDS_PasswordSettingsPrecedence for the instance
func (instance *ads_msds_passwordsettings) SetPropertyDS_msDS_PasswordSettingsPrecedence(value int32) (err error) { 
    return instance.SetProperty("DS_msDS_PasswordSettingsPrecedence", (value))
}

// GetDS_msDS_PasswordSettingsPrecedence gets the value of DS_msDS_PasswordSettingsPrecedence for the instance
func (instance *ads_msds_passwordsettings) GetPropertyDS_msDS_PasswordSettingsPrecedence()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_PasswordSettingsPrecedence")
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

// SetDS_msDS_PSOAppliesTo sets the value of DS_msDS_PSOAppliesTo for the instance
func (instance *ads_msds_passwordsettings) SetPropertyDS_msDS_PSOAppliesTo(value []string) (err error) { 
    return instance.SetProperty("DS_msDS_PSOAppliesTo", (value))
}

// GetDS_msDS_PSOAppliesTo gets the value of DS_msDS_PSOAppliesTo for the instance
func (instance *ads_msds_passwordsettings) GetPropertyDS_msDS_PSOAppliesTo()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_PSOAppliesTo")
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

