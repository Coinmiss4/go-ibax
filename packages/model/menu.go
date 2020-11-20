/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package model

import "github.com/IBAX-io/go-ibax/packages/converter"

// Menu is model
type Menu struct {
	ecosystem  int64
}

// SetTablePrefix is setting table prefix
func (m *Menu) SetTablePrefix(prefix string) {
	m.ecosystem = converter.StrToInt64(prefix)
}

// TableName returns name of table
func (m Menu) TableName() string {
	if m.ecosystem == 0 {
		m.ecosystem = 1
	}
	return `1_menu`
}

// Get is retrieving model from database
func (m *Menu) Get(name string) (bool, error) {
	return isFound(DBConn.Where("ecosystem=? and name = ?", m.ecosystem, name).First(m))
}
