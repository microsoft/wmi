// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_IEFavoriteItem struct
type RSOP_IEFavoriteItem struct {
	RSOP_IEFavoriteOrLinkItem

	//
	folderItem bool

	//
	parentPath string

	//
	rsopID string

	//
	rsopPrecedence uint32

	//
	shortName string
}

// SetfolderItem sets the value of folderItem for the instance
func (instance *RSOP_IEFavoriteItem) SetPropertyfolderItem(value bool) (err error) {
	return instance.SetProperty("folderItem", value)
}

// GetfolderItem gets the value of folderItem for the instance
func (instance *RSOP_IEFavoriteItem) GetPropertyfolderItem() (value bool, err error) {
	retValue, err := instance.GetProperty("folderItem")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetparentPath sets the value of parentPath for the instance
func (instance *RSOP_IEFavoriteItem) SetPropertyparentPath(value string) (err error) {
	return instance.SetProperty("parentPath", value)
}

// GetparentPath gets the value of parentPath for the instance
func (instance *RSOP_IEFavoriteItem) GetPropertyparentPath() (value string, err error) {
	retValue, err := instance.GetProperty("parentPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetrsopID sets the value of rsopID for the instance
func (instance *RSOP_IEFavoriteItem) SetPropertyrsopID(value string) (err error) {
	return instance.SetProperty("rsopID", value)
}

// GetrsopID gets the value of rsopID for the instance
func (instance *RSOP_IEFavoriteItem) GetPropertyrsopID() (value string, err error) {
	retValue, err := instance.GetProperty("rsopID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetrsopPrecedence sets the value of rsopPrecedence for the instance
func (instance *RSOP_IEFavoriteItem) SetPropertyrsopPrecedence(value uint32) (err error) {
	return instance.SetProperty("rsopPrecedence", value)
}

// GetrsopPrecedence gets the value of rsopPrecedence for the instance
func (instance *RSOP_IEFavoriteItem) GetPropertyrsopPrecedence() (value uint32, err error) {
	retValue, err := instance.GetProperty("rsopPrecedence")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetshortName sets the value of shortName for the instance
func (instance *RSOP_IEFavoriteItem) SetPropertyshortName(value string) (err error) {
	return instance.SetProperty("shortName", value)
}

// GetshortName gets the value of shortName for the instance
func (instance *RSOP_IEFavoriteItem) GetPropertyshortName() (value string, err error) {
	retValue, err := instance.GetProperty("shortName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
