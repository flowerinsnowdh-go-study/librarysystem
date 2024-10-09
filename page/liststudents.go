package page

import (
    "github.com/flowerinsnowdh-go-study/librarysystem/dao"
    "github.com/flowerinsnowdh-go-study/librarysystem/object"
    "strconv"
)

type ListStudentsPageRow struct {
    Id        int
    Name      string
    Book      string
    OptionURL string
    OptionName string
}

func LoadNewStudents(library *dao.SimpleLibrary) []*ListStudentsPageRow {
    var list []*object.StudentWithHeldBook = library.ListStudents()
    var rows []*ListStudentsPageRow = make([]*ListStudentsPageRow, 0, len(list))
    for _, student := range list {
        var book string
        var optionURL string
        var optionName string
        if student.Borrowed != nil {
            book = student.Borrowed.String()
            optionURL = "/release?book=" + strconv.Itoa(student.Borrowed.Id)
            optionName = "还书"
        } else {
            book = "无"
            optionURL = "/borrownew"
            optionName = "借书"
        }
        rows = append(rows, &ListStudentsPageRow{
            Id:         student.Instance.Id,
            Name:       student.Instance.Name,
            Book:       book,
            OptionURL:  optionURL,
            OptionName: optionName,
        })
    }
    return rows
}
