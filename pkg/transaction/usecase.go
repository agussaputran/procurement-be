package transaction

import "context"

type ITransactionUsecase interface {
	Store(ctx context.Context, request TransactionData) (*TransactionFetchResponse, error)
}

type TransactionUsecase struct {
	repository ITransactionRepository
}

func NewTransactionUsecase(repository ITransactionRepository) ITransactionUsecase {
	return &TransactionUsecase{
		repository: repository,
	}
}

func (u *TransactionUsecase) Store(ctx context.Context, request TransactionData) (response *TransactionFetchResponse, err error) {
	return
}
