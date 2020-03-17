// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_FileSystem struct
type CIM_FileSystem struct {
	CIM_LogicalElement

	//
	AvailableSpace uint64

	//
	BlockSize uint64

	//
	CasePreserved bool

	//
	CaseSensitive bool

	//
	CodeSet []uint16

	//
	CompressionMethod string

	//
	CreationClassName string

	//
	CSCreationClassName string

	//
	CSName string

	//
	EncryptionMethod string

	//
	FileSystemSize uint64

	//
	MaxFileNameLength uint32

	//
	ReadOnly bool

	//
	Root string
}

// SetAvailableSpace sets the value of AvailableSpace for the instance
func (instance *CIM_FileSystem) SetPropertyAvailableSpace(value uint64) (err error) {
	return instance.SetProperty("AvailableSpace", value)
}

// GetAvailableSpace gets the value of AvailableSpace for the instance
func (instance *CIM_FileSystem) GetPropertyAvailableSpace() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvailableSpace")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBlockSize sets the value of BlockSize for the instance
func (instance *CIM_FileSystem) SetPropertyBlockSize(value uint64) (err error) {
	return instance.SetProperty("BlockSize", value)
}

// GetBlockSize gets the value of BlockSize for the instance
func (instance *CIM_FileSystem) GetPropertyBlockSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("BlockSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCasePreserved sets the value of CasePreserved for the instance
func (instance *CIM_FileSystem) SetPropertyCasePreserved(value bool) (err error) {
	return instance.SetProperty("CasePreserved", value)
}

// GetCasePreserved gets the value of CasePreserved for the instance
func (instance *CIM_FileSystem) GetPropertyCasePreserved() (value bool, err error) {
	retValue, err := instance.GetProperty("CasePreserved")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCaseSensitive sets the value of CaseSensitive for the instance
func (instance *CIM_FileSystem) SetPropertyCaseSensitive(value bool) (err error) {
	return instance.SetProperty("CaseSensitive", value)
}

// GetCaseSensitive gets the value of CaseSensitive for the instance
func (instance *CIM_FileSystem) GetPropertyCaseSensitive() (value bool, err error) {
	retValue, err := instance.GetProperty("CaseSensitive")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCodeSet sets the value of CodeSet for the instance
func (instance *CIM_FileSystem) SetPropertyCodeSet(value []uint16) (err error) {
	return instance.SetProperty("CodeSet", value)
}

// GetCodeSet gets the value of CodeSet for the instance
func (instance *CIM_FileSystem) GetPropertyCodeSet() (value []uint16, err error) {
	retValue, err := instance.GetProperty("CodeSet")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompressionMethod sets the value of CompressionMethod for the instance
func (instance *CIM_FileSystem) SetPropertyCompressionMethod(value string) (err error) {
	return instance.SetProperty("CompressionMethod", value)
}

// GetCompressionMethod gets the value of CompressionMethod for the instance
func (instance *CIM_FileSystem) GetPropertyCompressionMethod() (value string, err error) {
	retValue, err := instance.GetProperty("CompressionMethod")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *CIM_FileSystem) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", value)
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *CIM_FileSystem) GetPropertyCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCSCreationClassName sets the value of CSCreationClassName for the instance
func (instance *CIM_FileSystem) SetPropertyCSCreationClassName(value string) (err error) {
	return instance.SetProperty("CSCreationClassName", value)
}

// GetCSCreationClassName gets the value of CSCreationClassName for the instance
func (instance *CIM_FileSystem) GetPropertyCSCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CSCreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCSName sets the value of CSName for the instance
func (instance *CIM_FileSystem) SetPropertyCSName(value string) (err error) {
	return instance.SetProperty("CSName", value)
}

// GetCSName gets the value of CSName for the instance
func (instance *CIM_FileSystem) GetPropertyCSName() (value string, err error) {
	retValue, err := instance.GetProperty("CSName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEncryptionMethod sets the value of EncryptionMethod for the instance
func (instance *CIM_FileSystem) SetPropertyEncryptionMethod(value string) (err error) {
	return instance.SetProperty("EncryptionMethod", value)
}

// GetEncryptionMethod gets the value of EncryptionMethod for the instance
func (instance *CIM_FileSystem) GetPropertyEncryptionMethod() (value string, err error) {
	retValue, err := instance.GetProperty("EncryptionMethod")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileSystemSize sets the value of FileSystemSize for the instance
func (instance *CIM_FileSystem) SetPropertyFileSystemSize(value uint64) (err error) {
	return instance.SetProperty("FileSystemSize", value)
}

// GetFileSystemSize gets the value of FileSystemSize for the instance
func (instance *CIM_FileSystem) GetPropertyFileSystemSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileSystemSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxFileNameLength sets the value of MaxFileNameLength for the instance
func (instance *CIM_FileSystem) SetPropertyMaxFileNameLength(value uint32) (err error) {
	return instance.SetProperty("MaxFileNameLength", value)
}

// GetMaxFileNameLength gets the value of MaxFileNameLength for the instance
func (instance *CIM_FileSystem) GetPropertyMaxFileNameLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxFileNameLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadOnly sets the value of ReadOnly for the instance
func (instance *CIM_FileSystem) SetPropertyReadOnly(value bool) (err error) {
	return instance.SetProperty("ReadOnly", value)
}

// GetReadOnly gets the value of ReadOnly for the instance
func (instance *CIM_FileSystem) GetPropertyReadOnly() (value bool, err error) {
	retValue, err := instance.GetProperty("ReadOnly")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoot sets the value of Root for the instance
func (instance *CIM_FileSystem) SetPropertyRoot(value string) (err error) {
	return instance.SetProperty("Root", value)
}

// GetRoot gets the value of Root for the instance
func (instance *CIM_FileSystem) GetPropertyRoot() (value string, err error) {
	retValue, err := instance.GetProperty("Root")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
