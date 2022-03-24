package models

type BookInfo struct {
	BookName          string
	NumOfPages        int
	NumOfBooksinStock int
	Price             int
	StockCode         string
	ISBN              string
	Author            struct {
		AuthorName string
	}
}

type BookList []BookInfo
