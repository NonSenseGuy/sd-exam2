package store

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/nonsenseguy/sd-exam2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type pg struct {
	db *gorm.DB
}

func NewPostgresConnection(conn string) IStore {
	db, err := gorm.Open(postgres.Open(conn),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "", log.LstdFlags),
				logger.Config{
					LogLevel: logger.Info,
					Colorful: true,
				},
			),
		})
	if err != nil {
		panic("Unable to connect to database: " + err.Error())
	}

	if err := db.AutoMigrate(&models.FraudataItem{}); err != nil {
		panic("Unable to migrate database: " + err.Error())
	}

	return &pg{db: db}
}

func (conn *pg) Get(ctx context.Context, in *models.IDRequest) (*models.FraudataItem, error) {
	item := &models.FraudataItem{}

	err := conn.db.WithContext(ctx).Take(item, "id = ?", in.ID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return item, err
}

func (conn *pg) List(ctx context.Context, in *models.ListRequest) ([]*models.FraudataItem, error) {
	if in.Limit == 0 || in.Limit > models.MaxListLimit {
		in.Limit = models.MaxListLimit
	}

	query := conn.db.WithContext(ctx).Limit(in.Limit)

	list := make([]*models.FraudataItem, 0, in.Limit)
	err := query.Order("id").Find(&list).Error

	return list, err
}

func (conn *pg) Report(ctx context.Context, in *models.ReportRequest) error {
	if in.FraudataItem == nil {
		return errors.New("unexpected report input")
	}

	in.FraudataItem.ID = GenerateUniqueID()
	in.FraudataItem.CreatedOn = conn.db.NowFunc()

	return conn.db.WithContext(ctx).Create(in.FraudataItem).Error
}

func (conn *pg) Update(ctx context.Context, in *models.ReportRequest) error {
	item := &models.FraudataItem{
		ID:            in.FraudataItem.ID,
		Name:          in.FraudataItem.Name,
		IsReported:    in.FraudataItem.IsReported,
		ReportReasons: in.FraudataItem.ReportReasons,
		UpdatedOn:     conn.db.NowFunc(),
	}

	return conn.db.WithContext(ctx).Model(item).Select("name", "is_reported", "report_reasons", "updated_on").Updates(item).Error
}

func (conn *pg) Delete(ctx context.Context, in *models.IDRequest) error {
	item := &models.FraudataItem{ID: in.ID}

	return conn.db.WithContext(ctx).Model(item).Delete(item).Error
}
