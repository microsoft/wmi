// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_ms_sql_sqlserver struct
type ads_ms_sql_sqlserver struct {
	*ads_serviceconnectionpoint

	//
	DS_mS_SQL_AppleTalk string

	//
	DS_mS_SQL_Build int32

	//
	DS_mS_SQL_CharacterSet int32

	//
	DS_mS_SQL_Clustered bool

	//
	DS_mS_SQL_Contact string

	//
	DS_mS_SQL_GPSHeight string

	//
	DS_mS_SQL_GPSLatitude string

	//
	DS_mS_SQL_GPSLongitude string

	//
	DS_mS_SQL_InformationURL string

	//
	DS_mS_SQL_Keywords []string

	//
	DS_mS_SQL_LastUpdatedDate string

	//
	DS_mS_SQL_Location string

	//
	DS_mS_SQL_Memory int64

	//
	DS_mS_SQL_MultiProtocol string

	//
	DS_mS_SQL_Name string

	//
	DS_mS_SQL_NamedPipe string

	//
	DS_mS_SQL_RegisteredOwner string

	//
	DS_mS_SQL_ServiceAccount string

	//
	DS_mS_SQL_SortOrder string

	//
	DS_mS_SQL_SPX string

	//
	DS_mS_SQL_Status int64

	//
	DS_mS_SQL_TCPIP string

	//
	DS_mS_SQL_UnicodeSortOrder int32

	//
	DS_mS_SQL_Vines string
}

func Newads_ms_sql_sqlserverEx1(instance *cim.WmiInstance) (newInstance *ads_ms_sql_sqlserver, err error) {
	tmp, err := Newads_serviceconnectionpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ms_sql_sqlserver{
		ads_serviceconnectionpoint: tmp,
	}
	return
}

func Newads_ms_sql_sqlserverEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ms_sql_sqlserver, err error) {
	tmp, err := Newads_serviceconnectionpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ms_sql_sqlserver{
		ads_serviceconnectionpoint: tmp,
	}
	return
}

// SetDS_mS_SQL_AppleTalk sets the value of DS_mS_SQL_AppleTalk for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_AppleTalk(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_AppleTalk", (value))
}

// GetDS_mS_SQL_AppleTalk gets the value of DS_mS_SQL_AppleTalk for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_AppleTalk() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_AppleTalk")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDS_mS_SQL_Build sets the value of DS_mS_SQL_Build for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_Build(value int32) (err error) {
	return instance.SetProperty("DS_mS_SQL_Build", (value))
}

// GetDS_mS_SQL_Build gets the value of DS_mS_SQL_Build for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_Build() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Build")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetDS_mS_SQL_CharacterSet sets the value of DS_mS_SQL_CharacterSet for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_CharacterSet(value int32) (err error) {
	return instance.SetProperty("DS_mS_SQL_CharacterSet", (value))
}

// GetDS_mS_SQL_CharacterSet gets the value of DS_mS_SQL_CharacterSet for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_CharacterSet() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_CharacterSet")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetDS_mS_SQL_Clustered sets the value of DS_mS_SQL_Clustered for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_Clustered(value bool) (err error) {
	return instance.SetProperty("DS_mS_SQL_Clustered", (value))
}

// GetDS_mS_SQL_Clustered gets the value of DS_mS_SQL_Clustered for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_Clustered() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Clustered")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetDS_mS_SQL_Contact sets the value of DS_mS_SQL_Contact for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_Contact(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_Contact", (value))
}

// GetDS_mS_SQL_Contact gets the value of DS_mS_SQL_Contact for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_Contact() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Contact")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDS_mS_SQL_GPSHeight sets the value of DS_mS_SQL_GPSHeight for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_GPSHeight(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_GPSHeight", (value))
}

// GetDS_mS_SQL_GPSHeight gets the value of DS_mS_SQL_GPSHeight for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_GPSHeight() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_GPSHeight")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDS_mS_SQL_GPSLatitude sets the value of DS_mS_SQL_GPSLatitude for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_GPSLatitude(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_GPSLatitude", (value))
}

// GetDS_mS_SQL_GPSLatitude gets the value of DS_mS_SQL_GPSLatitude for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_GPSLatitude() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_GPSLatitude")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDS_mS_SQL_GPSLongitude sets the value of DS_mS_SQL_GPSLongitude for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_GPSLongitude(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_GPSLongitude", (value))
}

// GetDS_mS_SQL_GPSLongitude gets the value of DS_mS_SQL_GPSLongitude for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_GPSLongitude() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_GPSLongitude")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDS_mS_SQL_InformationURL sets the value of DS_mS_SQL_InformationURL for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_InformationURL(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_InformationURL", (value))
}

// GetDS_mS_SQL_InformationURL gets the value of DS_mS_SQL_InformationURL for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_InformationURL() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_InformationURL")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDS_mS_SQL_Keywords sets the value of DS_mS_SQL_Keywords for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_Keywords(value []string) (err error) {
	return instance.SetProperty("DS_mS_SQL_Keywords", (value))
}

// GetDS_mS_SQL_Keywords gets the value of DS_mS_SQL_Keywords for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_Keywords() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Keywords")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetDS_mS_SQL_LastUpdatedDate sets the value of DS_mS_SQL_LastUpdatedDate for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_LastUpdatedDate(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_LastUpdatedDate", (value))
}

// GetDS_mS_SQL_LastUpdatedDate gets the value of DS_mS_SQL_LastUpdatedDate for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_LastUpdatedDate() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_LastUpdatedDate")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDS_mS_SQL_Location sets the value of DS_mS_SQL_Location for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_Location(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_Location", (value))
}

// GetDS_mS_SQL_Location gets the value of DS_mS_SQL_Location for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_Location() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Location")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDS_mS_SQL_Memory sets the value of DS_mS_SQL_Memory for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_Memory(value int64) (err error) {
	return instance.SetProperty("DS_mS_SQL_Memory", (value))
}

// GetDS_mS_SQL_Memory gets the value of DS_mS_SQL_Memory for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_Memory() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Memory")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_mS_SQL_MultiProtocol sets the value of DS_mS_SQL_MultiProtocol for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_MultiProtocol(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_MultiProtocol", (value))
}

// GetDS_mS_SQL_MultiProtocol gets the value of DS_mS_SQL_MultiProtocol for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_MultiProtocol() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_MultiProtocol")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDS_mS_SQL_Name sets the value of DS_mS_SQL_Name for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_Name(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_Name", (value))
}

// GetDS_mS_SQL_Name gets the value of DS_mS_SQL_Name for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_Name() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Name")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDS_mS_SQL_NamedPipe sets the value of DS_mS_SQL_NamedPipe for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_NamedPipe(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_NamedPipe", (value))
}

// GetDS_mS_SQL_NamedPipe gets the value of DS_mS_SQL_NamedPipe for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_NamedPipe() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_NamedPipe")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDS_mS_SQL_RegisteredOwner sets the value of DS_mS_SQL_RegisteredOwner for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_RegisteredOwner(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_RegisteredOwner", (value))
}

// GetDS_mS_SQL_RegisteredOwner gets the value of DS_mS_SQL_RegisteredOwner for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_RegisteredOwner() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_RegisteredOwner")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDS_mS_SQL_ServiceAccount sets the value of DS_mS_SQL_ServiceAccount for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_ServiceAccount(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_ServiceAccount", (value))
}

// GetDS_mS_SQL_ServiceAccount gets the value of DS_mS_SQL_ServiceAccount for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_ServiceAccount() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_ServiceAccount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDS_mS_SQL_SortOrder sets the value of DS_mS_SQL_SortOrder for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_SortOrder(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_SortOrder", (value))
}

// GetDS_mS_SQL_SortOrder gets the value of DS_mS_SQL_SortOrder for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_SortOrder() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_SortOrder")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDS_mS_SQL_SPX sets the value of DS_mS_SQL_SPX for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_SPX(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_SPX", (value))
}

// GetDS_mS_SQL_SPX gets the value of DS_mS_SQL_SPX for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_SPX() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_SPX")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDS_mS_SQL_Status sets the value of DS_mS_SQL_Status for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_Status(value int64) (err error) {
	return instance.SetProperty("DS_mS_SQL_Status", (value))
}

// GetDS_mS_SQL_Status gets the value of DS_mS_SQL_Status for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_Status() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Status")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_mS_SQL_TCPIP sets the value of DS_mS_SQL_TCPIP for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_TCPIP(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_TCPIP", (value))
}

// GetDS_mS_SQL_TCPIP gets the value of DS_mS_SQL_TCPIP for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_TCPIP() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_TCPIP")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDS_mS_SQL_UnicodeSortOrder sets the value of DS_mS_SQL_UnicodeSortOrder for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_UnicodeSortOrder(value int32) (err error) {
	return instance.SetProperty("DS_mS_SQL_UnicodeSortOrder", (value))
}

// GetDS_mS_SQL_UnicodeSortOrder gets the value of DS_mS_SQL_UnicodeSortOrder for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_UnicodeSortOrder() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_UnicodeSortOrder")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetDS_mS_SQL_Vines sets the value of DS_mS_SQL_Vines for the instance
func (instance *ads_ms_sql_sqlserver) SetPropertyDS_mS_SQL_Vines(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_Vines", (value))
}

// GetDS_mS_SQL_Vines gets the value of DS_mS_SQL_Vines for the instance
func (instance *ads_ms_sql_sqlserver) GetPropertyDS_mS_SQL_Vines() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Vines")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
