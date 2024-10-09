/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package object

import (
	"database/sql"
	"strconv"
)

type Book struct {
	Id   int
	Name string
}

func (b *Book) String() string {
	return b.Name + "(" + strconv.Itoa(b.Id) + ")"
}

// NewBook 返回一个新的 Book 指针
// 如果无效返回空
func NewBook(id *sql.NullInt32, name *sql.NullString) *Book {
	if !id.Valid || !name.Valid {
		return nil
	}
	return &Book{
		Id:   int(id.Int32),
		Name: name.String,
	}
}

type BookWithHeldStudent struct {
	Instance *Book
	Borrower *Student
}
