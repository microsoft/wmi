// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_DiagnosticRecord struct
type CIM_DiagnosticRecord struct {
	CIM_RecordForLog

	//
	CreationTimeStamp string

	//
	ExpirationDate string

	//
	ManagedElementName string

	//
	OtherRecordTypeDescription string

	//
	RecordType DiagnosticRecord_RecordType

	//
	ServiceName string
}

// SetCreationTimeStamp sets the value of CreationTimeStamp for the instance
func (instance *CIM_DiagnosticRecord) SetPropertyCreationTimeStamp(value string) (err error) {
	return instance.SetProperty("CreationTimeStamp", value)
}

// GetCreationTimeStamp gets the value of CreationTimeStamp for the instance
func (instance *CIM_DiagnosticRecord) GetPropertyCreationTimeStamp() (value string, err error) {
	retValue, err := instance.GetProperty("CreationTimeStamp")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExpirationDate sets the value of ExpirationDate for the instance
func (instance *CIM_DiagnosticRecord) SetPropertyExpirationDate(value string) (err error) {
	return instance.SetProperty("ExpirationDate", value)
}

// GetExpirationDate gets the value of ExpirationDate for the instance
func (instance *CIM_DiagnosticRecord) GetPropertyExpirationDate() (value string, err error) {
	retValue, err := instance.GetProperty("ExpirationDate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetManagedElementName sets the value of ManagedElementName for the instance
func (instance *CIM_DiagnosticRecord) SetPropertyManagedElementName(value string) (err error) {
	return instance.SetProperty("ManagedElementName", value)
}

// GetManagedElementName gets the value of ManagedElementName for the instance
func (instance *CIM_DiagnosticRecord) GetPropertyManagedElementName() (value string, err error) {
	retValue, err := instance.GetProperty("ManagedElementName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherRecordTypeDescription sets the value of OtherRecordTypeDescription for the instance
func (instance *CIM_DiagnosticRecord) SetPropertyOtherRecordTypeDescription(value string) (err error) {
	return instance.SetProperty("OtherRecordTypeDescription", value)
}

// GetOtherRecordTypeDescription gets the value of OtherRecordTypeDescription for the instance
func (instance *CIM_DiagnosticRecord) GetPropertyOtherRecordTypeDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherRecordTypeDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecordType sets the value of RecordType for the instance
func (instance *CIM_DiagnosticRecord) SetPropertyRecordType(value DiagnosticRecord_RecordType) (err error) {
	return instance.SetProperty("RecordType", value)
}

// GetRecordType gets the value of RecordType for the instance
func (instance *CIM_DiagnosticRecord) GetPropertyRecordType() (value DiagnosticRecord_RecordType, err error) {
	retValue, err := instance.GetProperty("RecordType")
	if err != nil {
		return
	}
	value, ok := retValue.(DiagnosticRecord_RecordType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServiceName sets the value of ServiceName for the instance
func (instance *CIM_DiagnosticRecord) SetPropertyServiceName(value string) (err error) {
	return instance.SetProperty("ServiceName", value)
}

// GetServiceName gets the value of ServiceName for the instance
func (instance *CIM_DiagnosticRecord) GetPropertyServiceName() (value string, err error) {
	retValue, err := instance.GetProperty("ServiceName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
