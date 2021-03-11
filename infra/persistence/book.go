package persistence

import (
	"context"
	"creanArchitecture/domain/model"
	"creanArchitecture/domain/repository"
	"time"
)

/*
	Infra層
	先ほど述べた通り、infra層は技術的関心を扱う層です
	ここではさっき定義したrepositoryの処理を実装します
*/

/*
	repositoryという名前にしたいがDomain配下のrepositoryという
	パッケージ名が被ってしまうため persistenceで代替
*/

type bookPersistence struct{}

// NewBookPersistence: Bookデータに関する Persistenceを生成
func NewBookPersistence() repository.BookRepository {
	return &bookPersistence{}
}

// GetAll: DBからBookデータを全権取得(BookRepository インターフェースのGetAl()を実装したもの)
// -> 本来はDBからデータを取得するが、簡略化のために省略

/*
	実際にはDBにアクセスし、データを持ってくるようにします
	一旦ここではモックデータを返すようにしておきます
	また、「Persistence」の中身を見るとRepositoryのインターフェースを返していると思います
	本当はRepositoryという名前はDomain層とInfra層でパッケージ名が被っているためやむなく
*/
/*
	ここで疑問
	どこでインターフェースと関連づけているの？という疑問が生まれる
	答えは、NewBookPersistence()です
	この関数のの戻り値はインターフェースです。
	したがって「return &bookPersistence{}」がインターフェースを満たしてないとエラーになります
	このようにインターフェースを満たしているか否かを判別します
	NewBookPersistence()をどこで使うかは後述します
*/

/*
	依存関係を見ていきます
	上述した通り、Infra層はDomain層のインターフェースを満たすように作られているので
	Domain層に依存します
*/
func (bq bookPersistence) GetAll(context.Context) ([]*model.Book, error) {
	book1 := model.Book{}
	book1.ID = 1
	book1.Title = "DDDがわかる本"
	book1.Author = "太朗くん"
	book1.IssuedAt = time.Now().Add(-24 * time.Hour)

	book2 := model.Book{}
	book2.ID = 2
	book2.Title = "レイヤードアーキテクチャがわかる本"
	book2.IssuedAt = time.Now().Add(-24 * time.Hour)

	return []*model.Book{
		&book1,
		&book2,
	}, nil
}
