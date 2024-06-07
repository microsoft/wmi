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

// ads_addresstemplate struct
type ads_addresstemplate struct { 
	*ads_displaytemplate

	// 
	DS_addressSyntax Uint8Array

	// 
	DS_addressType string

	// 
	DS_perMsgDialogDisplayTable Uint8Array

	// 
	DS_perRecipDialogDisplayTable Uint8Array

	// 
	DS_proxyGenerationEnabled bool
}

	func Newads_addresstemplateEx1(instance *cim.WmiInstance) (newInstance *ads_addresstemplate, err error) {tmp, err := Newads_displaytemplateEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_addresstemplate {
	ads_displaytemplate: tmp,
	}
	return
	}
	

	func Newads_addresstemplateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_addresstemplate, err error) {tmp, err := Newads_displaytemplateEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_addresstemplate {
	ads_displaytemplate: tmp,
	}
	return
	}
	

// SetDS_addressSyntax sets the value of DS_addressSyntax for the instance
func (instance *ads_addresstemplate) SetPropertyDS_addressSyntax(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_addressSyntax", (value))
}

// GetDS_addressSyntax gets the value of DS_addressSyntax for the instance
func (instance *ads_addresstemplate) GetPropertyDS_addressSyntax()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_addressSyntax")
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

// SetDS_addressType sets the value of DS_addressType for the instance
func (instance *ads_addresstemplate) SetPropertyDS_addressType(value string) (err error) { 
    return instance.SetProperty("DS_addressType", (value))
}

// GetDS_addressType gets the value of DS_addressType for the instance
func (instance *ads_addresstemplate) GetPropertyDS_addressType()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_addressType")
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

// SetDS_perMsgDialogDisplayTable sets the value of DS_perMsgDialogDisplayTable for the instance
func (instance *ads_addresstemplate) SetPropertyDS_perMsgDialogDisplayTable(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_perMsgDialogDisplayTable", (value))
}

// GetDS_perMsgDialogDisplayTable gets the value of DS_perMsgDialogDisplayTable for the instance
func (instance *ads_addresstemplate) GetPropertyDS_perMsgDialogDisplayTable()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_perMsgDialogDisplayTable")
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

// SetDS_perRecipDialogDisplayTable sets the value of DS_perRecipDialogDisplayTable for the instance
func (instance *ads_addresstemplate) SetPropertyDS_perRecipDialogDisplayTable(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_perRecipDialogDisplayTable", (value))
}

// GetDS_perRecipDialogDisplayTable gets the value of DS_perRecipDialogDisplayTable for the instance
func (instance *ads_addresstemplate) GetPropertyDS_perRecipDialogDisplayTable()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_perRecipDialogDisplayTable")
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

// SetDS_proxyGenerationEnabled sets the value of DS_proxyGenerationEnabled for the instance
func (instance *ads_addresstemplate) SetPropertyDS_proxyGenerationEnabled(value bool) (err error) { 
    return instance.SetProperty("DS_proxyGenerationEnabled", (value))
}

// GetDS_proxyGenerationEnabled gets the value of DS_proxyGenerationEnabled for the instance
func (instance *ads_addresstemplate) GetPropertyDS_proxyGenerationEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("DS_proxyGenerationEnabled")
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

