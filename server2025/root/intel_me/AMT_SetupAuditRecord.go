// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.Intel_ME
//////////////////////////////////////////////
package intel_me

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// AMT_SetupAuditRecord struct
type AMT_SetupAuditRecord struct {
	*CIM_LogEntry

	// Indicates whether the certificate chain includes additional CA that were not recorded
	AdditionalCaSerialNums bool

	// The serial number of the first intermediate CA in the certification chain
	CACertificateSerial1 []uint8

	// The serial number of the second intermediate CA in the certification chain
	CACertificateSerial2 []uint8

	// The serial number of the third intermediate CA in the certification chain
	CACertificateSerial3 []uint8

	// Indicates whether a previous provisioning attempt was initiated by the host
	HostInitiated bool

	// Indicates whether the certificate timestamp verification was successful.
	IsTimeValid bool

	// Specifies the TLS algorithm used for provisioning
	ProvisioningTLSMode SetupAuditRecord_ProvisioningTLSMode

	// Common Name specified in the SSL certificate, typically this is the FQDN of the provisioning server
	ProvServerFQDN string

	// The Provisioning Server IP address
	ProvServerIP string

	// Specifies whether the OOB service was pre-configured with DNS suffix information to verify the SSL certificate used for configuration
	SecureDNS bool

	// The stored Hash value of the root certificate used to validate the provisioning server. More information on the root CA can be retrieved by querying the AMT_ProvisioningCertificateHash instance
	SelectedHashData []uint8

	// Specifies the hash algorithm used to generate seed values for HMAC computations performed in the registration operation
	SelectedHashType SetupAuditRecord_SelectedHashType
}

func NewAMT_SetupAuditRecordEx1(instance *cim.WmiInstance) (newInstance *AMT_SetupAuditRecord, err error) {
	tmp, err := NewCIM_LogEntryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AMT_SetupAuditRecord{
		CIM_LogEntry: tmp,
	}
	return
}

func NewAMT_SetupAuditRecordEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AMT_SetupAuditRecord, err error) {
	tmp, err := NewCIM_LogEntryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AMT_SetupAuditRecord{
		CIM_LogEntry: tmp,
	}
	return
}

// SetAdditionalCaSerialNums sets the value of AdditionalCaSerialNums for the instance
func (instance *AMT_SetupAuditRecord) SetPropertyAdditionalCaSerialNums(value bool) (err error) {
	return instance.SetProperty("AdditionalCaSerialNums", (value))
}

// GetAdditionalCaSerialNums gets the value of AdditionalCaSerialNums for the instance
func (instance *AMT_SetupAuditRecord) GetPropertyAdditionalCaSerialNums() (value bool, err error) {
	retValue, err := instance.GetProperty("AdditionalCaSerialNums")
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

// SetCACertificateSerial1 sets the value of CACertificateSerial1 for the instance
func (instance *AMT_SetupAuditRecord) SetPropertyCACertificateSerial1(value []uint8) (err error) {
	return instance.SetProperty("CACertificateSerial1", (value))
}

// GetCACertificateSerial1 gets the value of CACertificateSerial1 for the instance
func (instance *AMT_SetupAuditRecord) GetPropertyCACertificateSerial1() (value []uint8, err error) {
	retValue, err := instance.GetProperty("CACertificateSerial1")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetCACertificateSerial2 sets the value of CACertificateSerial2 for the instance
func (instance *AMT_SetupAuditRecord) SetPropertyCACertificateSerial2(value []uint8) (err error) {
	return instance.SetProperty("CACertificateSerial2", (value))
}

// GetCACertificateSerial2 gets the value of CACertificateSerial2 for the instance
func (instance *AMT_SetupAuditRecord) GetPropertyCACertificateSerial2() (value []uint8, err error) {
	retValue, err := instance.GetProperty("CACertificateSerial2")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetCACertificateSerial3 sets the value of CACertificateSerial3 for the instance
func (instance *AMT_SetupAuditRecord) SetPropertyCACertificateSerial3(value []uint8) (err error) {
	return instance.SetProperty("CACertificateSerial3", (value))
}

// GetCACertificateSerial3 gets the value of CACertificateSerial3 for the instance
func (instance *AMT_SetupAuditRecord) GetPropertyCACertificateSerial3() (value []uint8, err error) {
	retValue, err := instance.GetProperty("CACertificateSerial3")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetHostInitiated sets the value of HostInitiated for the instance
func (instance *AMT_SetupAuditRecord) SetPropertyHostInitiated(value bool) (err error) {
	return instance.SetProperty("HostInitiated", (value))
}

// GetHostInitiated gets the value of HostInitiated for the instance
func (instance *AMT_SetupAuditRecord) GetPropertyHostInitiated() (value bool, err error) {
	retValue, err := instance.GetProperty("HostInitiated")
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

// SetIsTimeValid sets the value of IsTimeValid for the instance
func (instance *AMT_SetupAuditRecord) SetPropertyIsTimeValid(value bool) (err error) {
	return instance.SetProperty("IsTimeValid", (value))
}

// GetIsTimeValid gets the value of IsTimeValid for the instance
func (instance *AMT_SetupAuditRecord) GetPropertyIsTimeValid() (value bool, err error) {
	retValue, err := instance.GetProperty("IsTimeValid")
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

// SetProvisioningTLSMode sets the value of ProvisioningTLSMode for the instance
func (instance *AMT_SetupAuditRecord) SetPropertyProvisioningTLSMode(value SetupAuditRecord_ProvisioningTLSMode) (err error) {
	return instance.SetProperty("ProvisioningTLSMode", (value))
}

// GetProvisioningTLSMode gets the value of ProvisioningTLSMode for the instance
func (instance *AMT_SetupAuditRecord) GetPropertyProvisioningTLSMode() (value SetupAuditRecord_ProvisioningTLSMode, err error) {
	retValue, err := instance.GetProperty("ProvisioningTLSMode")
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

	value = SetupAuditRecord_ProvisioningTLSMode(valuetmp)

	return
}

// SetProvServerFQDN sets the value of ProvServerFQDN for the instance
func (instance *AMT_SetupAuditRecord) SetPropertyProvServerFQDN(value string) (err error) {
	return instance.SetProperty("ProvServerFQDN", (value))
}

// GetProvServerFQDN gets the value of ProvServerFQDN for the instance
func (instance *AMT_SetupAuditRecord) GetPropertyProvServerFQDN() (value string, err error) {
	retValue, err := instance.GetProperty("ProvServerFQDN")
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

// SetProvServerIP sets the value of ProvServerIP for the instance
func (instance *AMT_SetupAuditRecord) SetPropertyProvServerIP(value string) (err error) {
	return instance.SetProperty("ProvServerIP", (value))
}

// GetProvServerIP gets the value of ProvServerIP for the instance
func (instance *AMT_SetupAuditRecord) GetPropertyProvServerIP() (value string, err error) {
	retValue, err := instance.GetProperty("ProvServerIP")
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

// SetSecureDNS sets the value of SecureDNS for the instance
func (instance *AMT_SetupAuditRecord) SetPropertySecureDNS(value bool) (err error) {
	return instance.SetProperty("SecureDNS", (value))
}

// GetSecureDNS gets the value of SecureDNS for the instance
func (instance *AMT_SetupAuditRecord) GetPropertySecureDNS() (value bool, err error) {
	retValue, err := instance.GetProperty("SecureDNS")
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

// SetSelectedHashData sets the value of SelectedHashData for the instance
func (instance *AMT_SetupAuditRecord) SetPropertySelectedHashData(value []uint8) (err error) {
	return instance.SetProperty("SelectedHashData", (value))
}

// GetSelectedHashData gets the value of SelectedHashData for the instance
func (instance *AMT_SetupAuditRecord) GetPropertySelectedHashData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("SelectedHashData")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetSelectedHashType sets the value of SelectedHashType for the instance
func (instance *AMT_SetupAuditRecord) SetPropertySelectedHashType(value SetupAuditRecord_SelectedHashType) (err error) {
	return instance.SetProperty("SelectedHashType", (value))
}

// GetSelectedHashType gets the value of SelectedHashType for the instance
func (instance *AMT_SetupAuditRecord) GetPropertySelectedHashType() (value SetupAuditRecord_SelectedHashType, err error) {
	retValue, err := instance.GetProperty("SelectedHashType")
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

	value = SetupAuditRecord_SelectedHashType(valuetmp)

	return
}
