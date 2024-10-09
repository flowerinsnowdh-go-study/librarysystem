/*
 * Copyright (c) 2024 flowerinsnow
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package dao

import (
	"database/sql"
	"github.com/flowerinsnowdh-go-study/librarysystem/object"
)

// SimpleLibrary 是 object.Library 的简单实现
type SimpleLibrary struct {
	DB *sql.DB
}

func (l *SimpleLibrary) BorrowBook(bookId int, studentId int) error {
	if book, ok := l.FindBookById(bookId); !ok {
		return &NoSuchBookError{
			Id: bookId,
		}
	} else if student, ok := l.FindStudentById(studentId); !ok {
		return &NoSuchStudentError{
			Id: studentId,
		}
	} else if book.Borrower != nil { // 一本书不能同时被两名同学借阅
		return &BookAlreadyBorrowedError{
			Book:     book.Instance,
			Borrower: book.Borrower,
		}
	} else if student.Borrowed != nil { // 一名同学不能同时借两本书
		return &StudentAlreadyBorrowedError{
			Student:  student.Instance,
			Borrowed: student.Borrowed,
		}
	} else {
		if _, err := l.DB.Exec(
			"UPDATE `book` SET `borrowed_by` = ? WHERE `id` = ?",
			student.Instance.Id, book.Instance.Id,
		); err != nil {
			panic(err)
		}
		return nil
	}
}

func (l *SimpleLibrary) ListBooks() []*object.BookWithHeldStudent {
	if rows, err := l.DB.Query(
		"SELECT `book`.`id`, `book`.`name`, `student`.`id`, `student`.`name` FROM `book` LEFT JOIN `student` ON `student`.`id` = `book`.`borrowed_by`",
	); err != nil {
		panic(err)
	} else {
		defer func() {
			_ = rows.Close()
		}()
		var result []*object.BookWithHeldStudent = make([]*object.BookWithHeldStudent, 0)
		for rows.Next() {
			var bookId int
			var bookName string
			var studentId sql.NullInt32
			var studentName sql.NullString
			_ = rows.Scan(&bookId, &bookName, &studentId, &studentName)

			var book *object.Book = &object.Book{
				Id:   bookId,
				Name: bookName,
			}
			var student *object.Student = object.NewStudent(&studentId, &studentName) // nullable
			result = append(result, &object.BookWithHeldStudent{
				Instance: book,
				Borrower: student,
			})
		}
		return result
	}
}

func (l *SimpleLibrary) ListStudents() []*object.StudentWithHeldBook {
	if rows, err := l.DB.Query(
		"SELECT `student`.`id`, `student`.`name`, `book`.`id`, `book`.`name` FROM `student` LEFT JOIN `book` ON `student`.`id` = `book`.`borrowed_by`",
	); err != nil {
		panic(err)
	} else {
		defer func() {
			_ = rows.Close()
		}()
		var result []*object.StudentWithHeldBook = make([]*object.StudentWithHeldBook, 0)
		for rows.Next() {
			var studentId int
			var studentName string
			var bookId sql.NullInt32
			var bookName sql.NullString
			_ = rows.Scan(&studentId, &studentName, &bookId, &bookName)

			var student *object.Student = &object.Student{
				Id:   studentId,
				Name: studentName,
			}
			var book *object.Book = object.NewBook(&bookId, &bookName) // nullable
			result = append(result, &object.StudentWithHeldBook{
				Instance: student,
				Borrowed: book,
			})
		}
		return result
	}
}

func (l *SimpleLibrary) FindStudentById(id int) (*object.StudentWithHeldBook, bool) {
	if rows, err := l.DB.Query(
		"SELECT `student`.`name`, `book`.`id`, `book`.`name` FROM `student` LEFT JOIN `book` ON `book`.`borrowed_by` = `student`.`id` WHERE `student`.`id` = ?",
		id,
	); err != nil {
		panic(err)
	} else {
		defer func() {
			_ = rows.Close()
		}()
		if !rows.Next() {
			return nil, false
		}

		var studentName string
		var bookId sql.NullInt32
		var bookName sql.NullString

		if err := rows.Scan(&studentName, &bookId, &bookName); err != nil {
			panic(err)
			return nil, false
		}

		var student *object.Student = &object.Student{
			Id:   id,
			Name: studentName,
		}

		var book *object.Book = object.NewBook(&bookId, &bookName) // nullable

		return &object.StudentWithHeldBook{
			Instance: student,
			Borrowed: book,
		}, true
	}
}

func (l *SimpleLibrary) FindBookById(id int) (*object.BookWithHeldStudent, bool) {
	if rows, err := l.DB.Query(
		"SELECT `book`.`name`, `student`.`id`, `student`.`name` FROM `book` LEFT JOIN `student` ON `book`.`borrowed_by` = `student`.`id` WHERE `book`.`id` = ?",
		id,
	); err != nil {
		panic(err)
	} else {
		defer func() {
			_ = rows.Close()
		}()
		if !rows.Next() {
			return nil, false
		}

		var bookName string
		var studentId sql.NullInt32
		var studentName sql.NullString

		if err := rows.Scan(&bookName, &studentId, &studentName); err != nil {
			panic(err)
			return nil, false
		}

		var book *object.Book = &object.Book{
			Id:   id,
			Name: bookName,
		}

		var student *object.Student = object.NewStudent(&studentId, &studentName) // nullable

		return &object.BookWithHeldStudent{
			Instance: book,
			Borrower: student,
		}, true
	}
}

func (l *SimpleLibrary) ReleaseBook(bookId int) error {
	if book, ok := l.FindBookById(bookId); !ok {
		return &NoSuchBookError{
			Id: bookId,
		}
	} else if book.Borrower != nil {
		return &BookNotBorrowedError{
			Id: bookId,
		}
	} else {
		if _, err := l.DB.Exec(
			"UPDATE `book` SET `borrowed_by` = NULL WHERE `id` = ?",
			book.Instance.Id,
		); err != nil {
			panic(err)
		}
		return nil
	}
}
