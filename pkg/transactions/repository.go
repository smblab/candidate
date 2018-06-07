package transactions

import (
	"time"

	"github.com/jinzhu/gorm"
	// adding support for sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Transaction model
type Transaction struct {
	Name            string `gorm:"primary_key"`
	Parent          string `gorm:"index"`
	Message         string
	Amount          float64
	CurrencyCode    string
	TransactionTime time.Time
}

// Repository of transactions
type Repository interface {
	AddTransaction(*Transaction) error
	GetTransaction(string) (*Transaction, error)
	DeleteTransaction(string) error
	ListTransactions(string) ([]Transaction, error)
	Close() error
}

type gormRepository struct {
	db *gorm.DB
}

func (r *gormRepository) AddTransaction(t *Transaction) error {
	if err := r.db.Create(t).Error; err != nil {
		return err
	}
	return nil
}

func (r *gormRepository) GetTransaction(name string) (*Transaction, error) {
	var t Transaction
	if err := r.db.First(&t, "name = ?", name).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *gormRepository) DeleteTransaction(name string) error {
	var t Transaction
	if err := r.db.First(&t, "name = ?", name).Error; err != nil {
		return err
	}
	if err := r.db.Delete(t).Error; err != nil {
		return err
	}
	return nil
}

func (r *gormRepository) ListTransactions(name string) ([]Transaction, error) {
	var transactions []Transaction
	if err := r.db.Find(&transactions, "parent = ?", name).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *gormRepository) Close() error {
	return r.db.Close()
}

// NewRepository initializes a new gorm backed repository
func NewRepository(dbfile string) (Repository, error) {
	db, err := gorm.Open("sqlite3", dbfile)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Transaction{})
	return &gormRepository{
		db: db,
	}, nil
}
