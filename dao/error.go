/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package dao

import (
	"github.com/flowerinsnowdh-go-study/librarysystem/object"
	"strconv"
)

// BookAlreadyBorrowedError 代表一本书已被一位同学借阅，不允许再借阅给其他同学的错误
type BookAlreadyBorrowedError struct {
	// Book 代表这本书
	Book *object.Book
	// Borrower 代表借给的学生
	Borrower *object.Student
}

func (err *BookAlreadyBorrowedError) Error() string {
	return "book " + err.Book.String() + " is already borrowed by " + err.Borrower.String()
}

// StudentAlreadyBorrowedError 代表一名同学已经借阅了一本书，不允许借阅第二本书的错误
type StudentAlreadyBorrowedError struct {
	// Student 代表这名学生
	Student *object.Student
	// Borrowed 代表借走的书
	Borrowed *object.Book
}

func (err *StudentAlreadyBorrowedError) Error() string {
	return "student " + err.Student.String() + " is already borrowed book " + err.Borrowed.String()
}

type NoSuchBookError struct {
	Id int
}

func (e *NoSuchBookError) Error() string {
	return "no such book: " + strconv.Itoa(e.Id)
}

type NoSuchStudentError struct {
	Id int
}

func (e *NoSuchStudentError) Error() string {
	return "no such student: " + strconv.Itoa(e.Id)
}

// BookNotBorrowedError 代表一本书未被借阅
type BookNotBorrowedError struct {
	Id int
}

func (e *BookNotBorrowedError) Error() string {
	return "book not borrowed"
}
