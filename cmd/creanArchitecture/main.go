package main

import (
	handler "creanArchitecture/handler/rest"
	"creanArchitecture/infra/persistence"
	"creanArchitecture/usecase"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/*
	順序
 	構造体定義
	実行関数定義

	interface定義
	interfaceで間接的に呼び出しもとの関数
*/

/*
	レイヤードアーキテクチャにおける各層の依存関係
	↓ 依存の方向

	Handler層
	   ↓
	Usecase層
	   ↓
	Domain層
	   ↑
	Infra層
*/

/*
	Handler, UseCase, Repositoryの三つを使用
	main.goにて、これらの依存関係を定義してやることで利用可能な状態にします
	（DIは今回は使用してない）
	この時利用するのが各層に用意されているNewXxx()という関数です
	NewXxx()を使用して、HandlerやUseCase,Repositoryを作成し
	メソッド実行できるようにします。
*/

/*
	55行目から57行目の処理
	ここで、各層のNewXxx()の処理を使って依存関係を定義しています。
	DIライブラリを使うことでよりスマートにかけると思いますが、
	愚直にやるならこんな感じです。
*/

func main() {
	// 依存関係を注入（DI までは行きませんが注入っぽいことをしている）
	// DI ライブラリを使えば、もっとスマートになるはず
	bookPersistence := persistence.NewBookPersistence()
	bookUseCase := usecase.NewBookUseCase(bookPersistence)
	bookHandler := handler.NewBookHandler(bookUseCase)

	// ルーティングの設定
	router := httprouter.New()
	router.GET("/api/v1/books", bookHandler.Index)
	// サーバ起動
	fmt.Println("========================")
	fmt.Println("Server Start >> http://localhost:3000")
	fmt.Println("========================")
	log.Fatal(http.ListenAndServe(":3000", router))

}
