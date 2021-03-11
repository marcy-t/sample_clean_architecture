package usecase

import (
	"context"
	"creanArchitecture/domain/model"
	"creanArchitecture/domain/repository"
)

// BookUseCase: BookにおけるUseCaseのインターフェース
type BookUseCase interface {
	GetAll(context.Context) ([]*model.Book, error)
}

type bookUseCase struct {
	bookRepository repository.BookRepository
}

// NewBookUseCase: Bookデータに関する UseCaseを生成
func NewBookUseCase(br repository.BookRepository) BookUseCase {
	return &bookUseCase{
		bookRepository: br,
	}
}

// GetAll : Book データを全件取得するためのユースケース
// -> 本システムではあまりユースケース層の恩恵を受けれないが、
// もう少し大きなシステムになってくると
// 「ドメインモデルの調整者」としての役割が見えてくる
func (bu bookUseCase) GetAll(ctx context.Context) (book []*model.Book, err error) {
	// Persistence(Repository)
	books, err := bu.bookRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}
