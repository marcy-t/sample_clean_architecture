package repository

import (
	"context"
	"creanArchitecture/domain/model"
)

/*
	今回、リポジトリでやることを簡単に言うと、
	DBやKVSなどで行うCRUD処理の定義です
	ただし、Domain層には技術的関心ごとを持ち込まないというルール
	ここではinterfaceを定義するだけ

	実装は後述するinfraで行います
	(infra層は技術的内容を行う層)
*/

// BookRepository: BookにおけるRepositoryのインターフェイス
// -> 依存性逆転の法則により infra層はdomain層（本インターフェース）に依存
type BookRepository interface {
	GetAll(context.Context) ([]*model.Book, error)
}
