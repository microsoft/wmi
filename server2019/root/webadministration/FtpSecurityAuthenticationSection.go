// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// FtpSecurityAuthenticationSection struct
type FtpSecurityAuthenticationSection struct { 
	*ConfigurationSectionWithCollection

	// 
	DenyByFailure FtpDenyByFailureSettings
}

	func NewFtpSecurityAuthenticationSectionEx1(instance *cim.WmiInstance) (newInstance *FtpSecurityAuthenticationSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &FtpSecurityAuthenticationSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewFtpSecurityAuthenticationSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *FtpSecurityAuthenticationSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &FtpSecurityAuthenticationSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetDenyByFailure sets the value of DenyByFailure for the instance
func (instance *FtpSecurityAuthenticationSection) SetPropertyDenyByFailure(value FtpDenyByFailureSettings) (err error) { 
    return instance.SetProperty("DenyByFailure", (value))
}

// GetDenyByFailure gets the value of DenyByFailure for the instance
func (instance *FtpSecurityAuthenticationSection) GetPropertyDenyByFailure()(value FtpDenyByFailureSettings, err error) { 
    retValue, err := instance.GetProperty("DenyByFailure")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(FtpDenyByFailureSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " FtpDenyByFailureSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = FtpDenyByFailureSettings(valuetmp)

    return
}

