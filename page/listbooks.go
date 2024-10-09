package page

import (
    "github.com/flowerinsnowdh-go-study/librarysystem/dao"
    "github.com/flowerinsnowdh-go-study/librarysystem/object"
    "strconv"
)

type ListBooksPageRow struct {
    Id        int
    Name      string
    Student   string
    OptionURL string
    OptionName string
}

func LoadNewBooks(library *dao.SimpleLibrary) []*ListBooksPageRow {
    var list []*object.BookWithHeldStudent = library.ListBooks()
    var rows []*ListBooksPageRow = make([]*ListBooksPageRow, 0, len(list))
    for _, book := range list {
        var student string
        var optionURL string
        var optionName string
        if book.Borrower != nil {
            student = book.Borrower.String()
            optionURL = "/release?book=" + strconv.Itoa(book.Instance.Id)
            optionName = "还书"
        } else {
            student = "无"
            optionURL = "/borrowto"
            optionName = "借给"
        }
        rows = append(rows, &ListBooksPageRow{
            Id:         book.Instance.Id,
            Name:       book.Instance.Name,
            Student:    student,
            OptionURL:  optionURL,
            OptionName: optionName,
        })
    }
    return rows
}
