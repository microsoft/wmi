// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_SmbMapping struct
type MSFT_SmbMapping struct {
	*cim.WmiInstance

	//
	CompressNetworkTraffic int32

	//
	GlobalMapping bool

	//
	LocalPath string

	//
	RemotePath string

	//
	RequireIntegrity bool

	//
	RequirePrivacy bool

	//
	SkipCertificateCheck bool

	//
	Status SmbMapping_Status

	//
	TransportType SmbMapping_TransportType

	//
	UseWriteThrough bool
}

func NewMSFT_SmbMappingEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbMapping, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbMapping{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_SmbMappingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SmbMapping, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbMapping{
		WmiInstance: tmp,
	}
	return
}

// SetCompressNetworkTraffic sets the value of CompressNetworkTraffic for the instance
func (instance *MSFT_SmbMapping) SetPropertyCompressNetworkTraffic(value int32) (err error) {
	return instance.SetProperty("CompressNetworkTraffic", (value))
}

// GetCompressNetworkTraffic gets the value of CompressNetworkTraffic for the instance
func (instance *MSFT_SmbMapping) GetPropertyCompressNetworkTraffic() (value int32, err error) {
	retValue, err := instance.GetProperty("CompressNetworkTraffic")
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

// SetGlobalMapping sets the value of GlobalMapping for the instance
func (instance *MSFT_SmbMapping) SetPropertyGlobalMapping(value bool) (err error) {
	return instance.SetProperty("GlobalMapping", (value))
}

// GetGlobalMapping gets the value of GlobalMapping for the instance
func (instance *MSFT_SmbMapping) GetPropertyGlobalMapping() (value bool, err error) {
	retValue, err := instance.GetProperty("GlobalMapping")
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

// SetLocalPath sets the value of LocalPath for the instance
func (instance *MSFT_SmbMapping) SetPropertyLocalPath(value string) (err error) {
	return instance.SetProperty("LocalPath", (value))
}

// GetLocalPath gets the value of LocalPath for the instance
func (instance *MSFT_SmbMapping) GetPropertyLocalPath() (value string, err error) {
	retValue, err := instance.GetProperty("LocalPath")
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

// SetRemotePath sets the value of RemotePath for the instance
func (instance *MSFT_SmbMapping) SetPropertyRemotePath(value string) (err error) {
	return instance.SetProperty("RemotePath", (value))
}

// GetRemotePath gets the value of RemotePath for the instance
func (instance *MSFT_SmbMapping) GetPropertyRemotePath() (value string, err error) {
	retValue, err := instance.GetProperty("RemotePath")
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

// SetRequireIntegrity sets the value of RequireIntegrity for the instance
func (instance *MSFT_SmbMapping) SetPropertyRequireIntegrity(value bool) (err error) {
	return instance.SetProperty("RequireIntegrity", (value))
}

// GetRequireIntegrity gets the value of RequireIntegrity for the instance
func (instance *MSFT_SmbMapping) GetPropertyRequireIntegrity() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireIntegrity")
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

// SetRequirePrivacy sets the value of RequirePrivacy for the instance
func (instance *MSFT_SmbMapping) SetPropertyRequirePrivacy(value bool) (err error) {
	return instance.SetProperty("RequirePrivacy", (value))
}

// GetRequirePrivacy gets the value of RequirePrivacy for the instance
func (instance *MSFT_SmbMapping) GetPropertyRequirePrivacy() (value bool, err error) {
	retValue, err := instance.GetProperty("RequirePrivacy")
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

// SetSkipCertificateCheck sets the value of SkipCertificateCheck for the instance
func (instance *MSFT_SmbMapping) SetPropertySkipCertificateCheck(value bool) (err error) {
	return instance.SetProperty("SkipCertificateCheck", (value))
}

// GetSkipCertificateCheck gets the value of SkipCertificateCheck for the instance
func (instance *MSFT_SmbMapping) GetPropertySkipCertificateCheck() (value bool, err error) {
	retValue, err := instance.GetProperty("SkipCertificateCheck")
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

// SetStatus sets the value of Status for the instance
func (instance *MSFT_SmbMapping) SetPropertyStatus(value SmbMapping_Status) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_SmbMapping) GetPropertyStatus() (value SmbMapping_Status, err error) {
	retValue, err := instance.GetProperty("Status")
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

	value = SmbMapping_Status(valuetmp)

	return
}

// SetTransportType sets the value of TransportType for the instance
func (instance *MSFT_SmbMapping) SetPropertyTransportType(value SmbMapping_TransportType) (err error) {
	return instance.SetProperty("TransportType", (value))
}

// GetTransportType gets the value of TransportType for the instance
func (instance *MSFT_SmbMapping) GetPropertyTransportType() (value SmbMapping_TransportType, err error) {
	retValue, err := instance.GetProperty("TransportType")
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

	value = SmbMapping_TransportType(valuetmp)

	return
}

// SetUseWriteThrough sets the value of UseWriteThrough for the instance
func (instance *MSFT_SmbMapping) SetPropertyUseWriteThrough(value bool) (err error) {
	return instance.SetProperty("UseWriteThrough", (value))
}

// GetUseWriteThrough gets the value of UseWriteThrough for the instance
func (instance *MSFT_SmbMapping) GetPropertyUseWriteThrough() (value bool, err error) {
	retValue, err := instance.GetProperty("UseWriteThrough")
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

//

// <param name="Force" type="bool "></param>
// <param name="GlobalMapping" type="bool "></param>
// <param name="UpdateProfile" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbMapping) Remove( /* IN */ UpdateProfile bool,
	/* IN */ Force bool,
	/* IN */ GlobalMapping bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Remove", UpdateProfile, Force, GlobalMapping)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="CompressNetworkTraffic" type="bool "></param>
// <param name="GlobalMapping" type="bool "></param>
// <param name="HomeFolder" type="bool "></param>
// <param name="LocalPath" type="string "></param>
// <param name="Password" type="string "></param>
// <param name="Persistent" type="bool "></param>
// <param name="RemotePath" type="string "></param>
// <param name="RequireIntegrity" type="bool "></param>
// <param name="RequirePrivacy" type="bool "></param>
// <param name="SaveCredentials" type="bool "></param>
// <param name="SkipCertificateCheck" type="bool "></param>
// <param name="TransportType" type="uint32 "></param>
// <param name="UserName" type="string "></param>
// <param name="UseWriteThrough" type="bool "></param>

// <param name="CreatedMapping" type="MSFT_SmbMapping "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbMapping) Create( /* IN */ LocalPath string,
	/* IN */ RemotePath string,
	/* IN */ UserName string,
	/* IN */ Password string,
	/* IN */ Persistent bool,
	/* IN */ SaveCredentials bool,
	/* IN */ HomeFolder bool,
	/* OUT */ CreatedMapping MSFT_SmbMapping,
	/* OPTIONAL IN */ RequireIntegrity bool,
	/* OPTIONAL IN */ RequirePrivacy bool,
	/* OPTIONAL IN */ UseWriteThrough bool,
	/* OPTIONAL IN */ TransportType uint32,
	/* OPTIONAL IN */ SkipCertificateCheck bool,
	/* OPTIONAL IN */ CompressNetworkTraffic bool,
	/* OPTIONAL IN */ GlobalMapping bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Create", LocalPath, RemotePath, UserName, Password, Persistent, SaveCredentials, HomeFolder, RequireIntegrity, RequirePrivacy, UseWriteThrough, TransportType, SkipCertificateCheck, CompressNetworkTraffic, GlobalMapping)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
