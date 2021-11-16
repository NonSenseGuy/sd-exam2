package store

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/nonsenseguy/sd-exam2/models"
)

type IStore interface {
	Get(ctx context.Context, in *models.IDRequest) (*models.FraudataItem, error)
	List(ctx context.Context, in *models.ListRequest) ([]*models.FraudataItem, error)
	Report(ctx context.Context, in *models.ReportRequest) error
	Update(ctx context.Context, in *models.ReportRequest) error
	Delete(ctx context.Context, in *models.IDRequest) error
}

func init() {
	rand.Seed(time.Now().UTC().Unix())
}

func GenerateUniqueID() string {
	word := []byte("0987654321")
	rand.Shuffle(len(word), func(i, j int) {
		word[i], word[j] = word[j], word[i]
	})

	now := time.Now().UTC()

	return fmt.Sprintf("%010v-%010v-%s", now.Unix(), now.Nanosecond(), string(word))
}
