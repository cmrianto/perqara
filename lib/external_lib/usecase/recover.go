package usecase

import (
	"context"
	"log"
)

func (e *usecase) CatchRecover(ctx context.Context) {
	if r := recover(); r != nil {
		log.Println(r)
	}
}
