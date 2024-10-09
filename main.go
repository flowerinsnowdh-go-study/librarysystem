/*
 * Copyright (c) 2024 flowerinsnow
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"github.com/flowerinsnowdh-go-study/librarysystem/dao"
	"github.com/flowerinsnowdh-go-study/librarysystem/object"
    "github.com/flowerinsnowdh-go-study/librarysystem/page"
    _ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var library *dao.SimpleLibrary

func main() {
	const dbFile = "librarysystem.db"
	const initSQLFile = "init.sql"

	var shouldInit bool = false

	if _, err := os.Stat(dbFile); err != nil {
		_ = os.Remove(dbFile)
		shouldInit = true
	}

	var db *sql.DB
	var err error
	db, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed to open database file", dbFile)
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(125)
	}
	defer func() {
		_ = db.Close()
	}()
	library = &dao.SimpleLibrary{
		DB: db,
	}
	if shouldInit {
		if bytes, err := os.ReadFile(initSQLFile); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "failed to open initSQL file", initSQLFile)
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(125)
		} else {
			_, _ = db.Exec(string(bytes))
		}
	}

	//var mux *http.ServeMux = http.NewServeMux()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, "resources/index.html")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	})
	http.HandleFunc("/liststudents", func(w http.ResponseWriter, r *http.Request) {
		var t *template.Template
		var err error
		if t, err = template.ParseFiles("resources/liststudents.html"); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
        var rows = page.LoadNewStudents(library)
		if err = t.Execute(w, rows); err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	})
	http.HandleFunc("/listbooks", func(w http.ResponseWriter, r *http.Request) {
		var t *template.Template
		var err error
		if t, err = template.ParseFiles("resources/listbooks.html"); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var rows = page.LoadNewBooks(library)
		if err = t.Execute(w, rows); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	})

	go func() {
		if err = http.ListenAndServe("[::1]:8080", nil); err != nil {
			panic(err)
		}
	}()

	printHelp()

	var scanner *bufio.Reader = bufio.NewReader(os.Stdin)
	var lineData []byte
	var line string
	var duration time.Duration

	for {
		fmt.Print("> ")
		lineData, _, _ = scanner.ReadLine()
		line = string(lineData)

		duration = time.Since(time.Now())
		if err, quit := command(strings.Split(line, " ")); quit {
			return
		} else if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "failed in "+strconv.FormatInt(duration.Milliseconds(), 10)+"ms")
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
		} else {
			fmt.Println("success in " + strconv.FormatInt(duration.Milliseconds(), 10) + "ms")
		}
	}
}

func command(args []string) (err error, quit bool) {
	switch strings.ToLower(args[0]) {
	case "liststudents":
		var students []*object.StudentWithHeldBook = library.ListStudents()
		fmt.Println("共", len(students), "名学生")
		for i := 0; i < len(students); i++ {
			if students[i].Borrowed != nil {
				fmt.Println(students[i].Instance.String(), "借阅了", students[i].Borrowed.String())
			} else {
				fmt.Println(students[i].Instance)
			}
		}
		return nil, false
	case "listbooks":
		var books []*object.BookWithHeldStudent = library.ListBooks()
		fmt.Println("共", len(books), "本书")
		for i := 0; i < len(books); i++ {
			if books[i].Borrower != nil {
				fmt.Println(books[i].Instance, "被", books[i].Borrower, "借阅")
			} else {
				fmt.Println(books[i].Instance)
			}
		}
		return nil, false
	case "borrow":
		// BEGIN 检查参数
		if len(args) != 3 {
			return errors.New("invalid length of arguments"), false
		}
		var studentId int
		var bookId int
		var e error
		if studentId, e = strconv.Atoi(args[1]); e != nil {
			return errors.New("not a number " + args[1]), false
		}
		if bookId, e = strconv.Atoi(args[2]); e != nil {
			return errors.New("not a number " + args[2]), false
		}
		// END 检查参数
		e = library.BorrowBook(studentId, bookId)
		return e, false
	case "release":
		// BEGIN 检查参数
		if len(args) != 2 {
			return errors.New("invalid length of arguments"), false
		}
		var bookId int
		var e error
		if bookId, e = strconv.Atoi(args[1]); e != nil {
			return errors.New("not a number " + args[1]), false
		}
		e = library.ReleaseBook(bookId)
		// END 检查参数
		return e, false
	case "help", "print":
		printHelp()
		return nil, false
	case "quit", "exit":
		fmt.Println("bye")
		return nil, true
	default:
		return errors.New("unknown command"), false
	}
}

func printHelp() {
	fmt.Println("命令：")
	fmt.Println("liststudents - 列出所有学生")
	fmt.Println("listbooks - 列出所有书")
	fmt.Println("borrow <student id> <book id> - 记录学生借阅了书")
	fmt.Println("release <book id> - 记录书被归还")
	fmt.Println("help - 打印此页面")
	fmt.Println("quit - 退出")
}
