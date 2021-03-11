package rest

// Handler層を変えるだけ、例えばCLIにも簡単に対応可能

import (
	"creanArchitecture/usecase"
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// BookHandler: Bookにおける Handlerのインターフェース
type BookHandler interface {
	Index(http.ResponseWriter, *http.Request, httprouter.Params)
}

type bookHandler struct {
	bookUseCase usecase.BookUseCase
}

// request : 本 APIのリクエストパラメータ
//　-> こんな感じでリクエストも受け取れますが、今回は使いません
type request struct {
	Begin uint `query:begin`
	Limit uint `query:limit`
}

/*
	bookField: response 内で使用するBookを表す構造体
	-> ドメインモデルBookにHTTPの関心事である JSONタグを付与したくないために
		簡略化のためにJSONタグを付与した
		ドメインモデルを流用するプロジェクトもしばしば見かける
*/
type bookField struct {
	Id      int64     `json:"id"`
	Title   string    `json:"title"`
	Author  string    `json:"author"`
	IssueAt time.Time `json:"issue_at"`
}

// response: 本 APIのレスポンス
type response struct {
	Books []*bookField `json: "books"`
}

// NewBookUseCase: Bookデータに関する Handlerを生成
func NewBookHandler(bu usecase.BookUseCase) BookHandler {
	return &bookHandler{
		bookUseCase: bu,
	}
}

// BookIndex: GET /books -> Book データ一覧を返す
func (bh bookHandler) Index(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {

	ctx := r.Context()

	// ユースケースの呼出
	books, err := bh.bookUseCase.GetAll(ctx)
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
	}
	// 取得したドメインモデルを responseに変換
	res := make([]*response, 0, len(books))
	bk := make([]*bookField, 0, len(books))

	for _, book := range books {
		bk = append(bk, &bookField{
			Id:      book.ID,
			Title:   book.Title,
			Author:  book.Author,
			IssueAt: book.IssuedAt,
		})
	}
	/*
		a := &response{
			Books: []*bookField{
				&bookField{
					Id:      123456,
					Title:   "book.Title",
					Author:  "book.Author",
					IssueAt: time.Time{},
				},
			},
		}
	*/
	res = append(res, &response{
		Books: bk,
	})
	// クライアントにレスポンスを追加
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
