package model

import "time"

/*
    Domain層
	システムが扱う業務領域に関するコードを置くところ
	よって「書籍」がどういうものなのかモデルという形で定義
*/

// Book: Bookを表すドメインモデル
type Book struct {
	ID       int64
	Title    string
	Author   string
	IssuedAt time.Time
}
