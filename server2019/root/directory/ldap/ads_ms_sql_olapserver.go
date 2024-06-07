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

// ads_ms_sql_olapserver struct
type ads_ms_sql_olapserver struct { 
	*ads_serviceconnectionpoint

	// 
	DS_mS_SQL_Build int32

	// 
	DS_mS_SQL_Contact string

	// 
	DS_mS_SQL_InformationURL string

	// 
	DS_mS_SQL_Keywords []string

	// 
	DS_mS_SQL_Language string

	// 
	DS_mS_SQL_Name string

	// 
	DS_mS_SQL_PublicationURL string

	// 
	DS_mS_SQL_RegisteredOwner string

	// 
	DS_mS_SQL_ServiceAccount string

	// 
	DS_mS_SQL_Status int64

	// 
	DS_mS_SQL_Version string
}

	func Newads_ms_sql_olapserverEx1(instance *cim.WmiInstance) (newInstance *ads_ms_sql_olapserver, err error) {tmp, err := Newads_serviceconnectionpointEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_ms_sql_olapserver {
	ads_serviceconnectionpoint: tmp,
	}
	return
	}
	

	func Newads_ms_sql_olapserverEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_ms_sql_olapserver, err error) {tmp, err := Newads_serviceconnectionpointEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_ms_sql_olapserver {
	ads_serviceconnectionpoint: tmp,
	}
	return
	}
	

// SetDS_mS_SQL_Build sets the value of DS_mS_SQL_Build for the instance
func (instance *ads_ms_sql_olapserver) SetPropertyDS_mS_SQL_Build(value int32) (err error) { 
    return instance.SetProperty("DS_mS_SQL_Build", (value))
}

// GetDS_mS_SQL_Build gets the value of DS_mS_SQL_Build for the instance
func (instance *ads_ms_sql_olapserver) GetPropertyDS_mS_SQL_Build()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_mS_SQL_Build")
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

// SetDS_mS_SQL_Contact sets the value of DS_mS_SQL_Contact for the instance
func (instance *ads_ms_sql_olapserver) SetPropertyDS_mS_SQL_Contact(value string) (err error) { 
    return instance.SetProperty("DS_mS_SQL_Contact", (value))
}

// GetDS_mS_SQL_Contact gets the value of DS_mS_SQL_Contact for the instance
func (instance *ads_ms_sql_olapserver) GetPropertyDS_mS_SQL_Contact()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_mS_SQL_Contact")
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

// SetDS_mS_SQL_InformationURL sets the value of DS_mS_SQL_InformationURL for the instance
func (instance *ads_ms_sql_olapserver) SetPropertyDS_mS_SQL_InformationURL(value string) (err error) { 
    return instance.SetProperty("DS_mS_SQL_InformationURL", (value))
}

// GetDS_mS_SQL_InformationURL gets the value of DS_mS_SQL_InformationURL for the instance
func (instance *ads_ms_sql_olapserver) GetPropertyDS_mS_SQL_InformationURL()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_mS_SQL_InformationURL")
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

// SetDS_mS_SQL_Keywords sets the value of DS_mS_SQL_Keywords for the instance
func (instance *ads_ms_sql_olapserver) SetPropertyDS_mS_SQL_Keywords(value []string) (err error) { 
    return instance.SetProperty("DS_mS_SQL_Keywords", (value))
}

// GetDS_mS_SQL_Keywords gets the value of DS_mS_SQL_Keywords for the instance
func (instance *ads_ms_sql_olapserver) GetPropertyDS_mS_SQL_Keywords()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_mS_SQL_Keywords")
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

// SetDS_mS_SQL_Language sets the value of DS_mS_SQL_Language for the instance
func (instance *ads_ms_sql_olapserver) SetPropertyDS_mS_SQL_Language(value string) (err error) { 
    return instance.SetProperty("DS_mS_SQL_Language", (value))
}

// GetDS_mS_SQL_Language gets the value of DS_mS_SQL_Language for the instance
func (instance *ads_ms_sql_olapserver) GetPropertyDS_mS_SQL_Language()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_mS_SQL_Language")
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

// SetDS_mS_SQL_Name sets the value of DS_mS_SQL_Name for the instance
func (instance *ads_ms_sql_olapserver) SetPropertyDS_mS_SQL_Name(value string) (err error) { 
    return instance.SetProperty("DS_mS_SQL_Name", (value))
}

// GetDS_mS_SQL_Name gets the value of DS_mS_SQL_Name for the instance
func (instance *ads_ms_sql_olapserver) GetPropertyDS_mS_SQL_Name()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_mS_SQL_Name")
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

// SetDS_mS_SQL_PublicationURL sets the value of DS_mS_SQL_PublicationURL for the instance
func (instance *ads_ms_sql_olapserver) SetPropertyDS_mS_SQL_PublicationURL(value string) (err error) { 
    return instance.SetProperty("DS_mS_SQL_PublicationURL", (value))
}

// GetDS_mS_SQL_PublicationURL gets the value of DS_mS_SQL_PublicationURL for the instance
func (instance *ads_ms_sql_olapserver) GetPropertyDS_mS_SQL_PublicationURL()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_mS_SQL_PublicationURL")
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

// SetDS_mS_SQL_RegisteredOwner sets the value of DS_mS_SQL_RegisteredOwner for the instance
func (instance *ads_ms_sql_olapserver) SetPropertyDS_mS_SQL_RegisteredOwner(value string) (err error) { 
    return instance.SetProperty("DS_mS_SQL_RegisteredOwner", (value))
}

// GetDS_mS_SQL_RegisteredOwner gets the value of DS_mS_SQL_RegisteredOwner for the instance
func (instance *ads_ms_sql_olapserver) GetPropertyDS_mS_SQL_RegisteredOwner()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_mS_SQL_RegisteredOwner")
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

// SetDS_mS_SQL_ServiceAccount sets the value of DS_mS_SQL_ServiceAccount for the instance
func (instance *ads_ms_sql_olapserver) SetPropertyDS_mS_SQL_ServiceAccount(value string) (err error) { 
    return instance.SetProperty("DS_mS_SQL_ServiceAccount", (value))
}

// GetDS_mS_SQL_ServiceAccount gets the value of DS_mS_SQL_ServiceAccount for the instance
func (instance *ads_ms_sql_olapserver) GetPropertyDS_mS_SQL_ServiceAccount()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_mS_SQL_ServiceAccount")
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

// SetDS_mS_SQL_Status sets the value of DS_mS_SQL_Status for the instance
func (instance *ads_ms_sql_olapserver) SetPropertyDS_mS_SQL_Status(value int64) (err error) { 
    return instance.SetProperty("DS_mS_SQL_Status", (value))
}

// GetDS_mS_SQL_Status gets the value of DS_mS_SQL_Status for the instance
func (instance *ads_ms_sql_olapserver) GetPropertyDS_mS_SQL_Status()(value int64, err error) { 
    retValue, err := instance.GetProperty("DS_mS_SQL_Status")
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

// SetDS_mS_SQL_Version sets the value of DS_mS_SQL_Version for the instance
func (instance *ads_ms_sql_olapserver) SetPropertyDS_mS_SQL_Version(value string) (err error) { 
    return instance.SetProperty("DS_mS_SQL_Version", (value))
}

// GetDS_mS_SQL_Version gets the value of DS_mS_SQL_Version for the instance
func (instance *ads_ms_sql_olapserver) GetPropertyDS_mS_SQL_Version()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_mS_SQL_Version")
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

