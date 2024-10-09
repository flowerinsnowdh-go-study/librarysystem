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

type Student struct {
	Id   int
	Name string
}

func (s *Student) String() string {
	return s.Name + "(" + strconv.Itoa(s.Id) + ")"
}

// NewStudent 返回一个 Student 指针
// 如果无效返回空
func NewStudent(id *sql.NullInt32, name *sql.NullString) *Student {
	if !id.Valid || !name.Valid {
		return nil
	}
	return &Student{
		Id:   int(id.Int32),
		Name: name.String,
	}
}

type StudentWithHeldBook struct {
	Instance *Student
	Borrowed *Book
}
