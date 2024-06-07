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

// ads_msmqsitelink struct
type ads_msmqsitelink struct { 
	*ds_top

	// 
	DS_mSMQCost int32

	// 
	DS_mSMQSite1 string

	// 
	DS_mSMQSite2 string

	// 
	DS_mSMQSiteGates []string

	// 
	DS_mSMQSiteGatesMig []string
}

	func Newads_msmqsitelinkEx1(instance *cim.WmiInstance) (newInstance *ads_msmqsitelink, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_msmqsitelink {
	ds_top: tmp,
	}
	return
	}
	

	func Newads_msmqsitelinkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_msmqsitelink, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_msmqsitelink {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_mSMQCost sets the value of DS_mSMQCost for the instance
func (instance *ads_msmqsitelink) SetPropertyDS_mSMQCost(value int32) (err error) { 
    return instance.SetProperty("DS_mSMQCost", (value))
}

// GetDS_mSMQCost gets the value of DS_mSMQCost for the instance
func (instance *ads_msmqsitelink) GetPropertyDS_mSMQCost()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_mSMQCost")
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

// SetDS_mSMQSite1 sets the value of DS_mSMQSite1 for the instance
func (instance *ads_msmqsitelink) SetPropertyDS_mSMQSite1(value string) (err error) { 
    return instance.SetProperty("DS_mSMQSite1", (value))
}

// GetDS_mSMQSite1 gets the value of DS_mSMQSite1 for the instance
func (instance *ads_msmqsitelink) GetPropertyDS_mSMQSite1()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_mSMQSite1")
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

// SetDS_mSMQSite2 sets the value of DS_mSMQSite2 for the instance
func (instance *ads_msmqsitelink) SetPropertyDS_mSMQSite2(value string) (err error) { 
    return instance.SetProperty("DS_mSMQSite2", (value))
}

// GetDS_mSMQSite2 gets the value of DS_mSMQSite2 for the instance
func (instance *ads_msmqsitelink) GetPropertyDS_mSMQSite2()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_mSMQSite2")
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

// SetDS_mSMQSiteGates sets the value of DS_mSMQSiteGates for the instance
func (instance *ads_msmqsitelink) SetPropertyDS_mSMQSiteGates(value []string) (err error) { 
    return instance.SetProperty("DS_mSMQSiteGates", (value))
}

// GetDS_mSMQSiteGates gets the value of DS_mSMQSiteGates for the instance
func (instance *ads_msmqsitelink) GetPropertyDS_mSMQSiteGates()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_mSMQSiteGates")
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

// SetDS_mSMQSiteGatesMig sets the value of DS_mSMQSiteGatesMig for the instance
func (instance *ads_msmqsitelink) SetPropertyDS_mSMQSiteGatesMig(value []string) (err error) { 
    return instance.SetProperty("DS_mSMQSiteGatesMig", (value))
}

// GetDS_mSMQSiteGatesMig gets the value of DS_mSMQSiteGatesMig for the instance
func (instance *ads_msmqsitelink) GetPropertyDS_mSMQSiteGatesMig()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_mSMQSiteGatesMig")
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

