/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package object

// Library 是一个抽象的图书馆，提供了最基本的内容
type Library interface {
	// BorrowBook 是让 Student 借走 Book
	// 如果书已被借走、学生已经借过书了、书不存在于数据库中、学生不存在于中会返回 error
	BorrowBook(int, int) error
	// ListBooks 将列出图书馆中所有的书籍以及它们的持有者
	ListBooks() []*BookWithHeldStudent
	// ListStudents 将列出数据库中所有的学生以及它们持有的图书
	ListStudents() []*StudentWithHeldBook
	// FindStudentById 用于通过 id 查找学生
	FindStudentById(int) (*StudentWithHeldBook, bool)
	// FindBookById 用于通过 id 查找学生
	FindBookById(int) (*BookWithHeldStudent, bool)
	// ReleaseBook 将将 Book 从一名学生手中归还
	// 如果 Book 不存在，则返回 error
	ReleaseBook(int) error
}
