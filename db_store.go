package common

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceType int

const (
	Product ServiceType = iota
	Planning
	Subscription
	Auth
	Customer
)

type DBStore struct {
	mux sync.Mutex
	dbs map[string]*gorm.DB
}

func NewStore() *DBStore {
	return &DBStore{
		mux: sync.Mutex{},
		dbs: make(map[string]*gorm.DB),
	}
}

func (s *DBStore) GetDB(c context.Context, service ServiceType) (*gorm.DB, error) {
	db, exists := s.dbs[c.Value("tenant_id").(string)]
	if exists {
		return db, errors.New("")
	}

	return s.CreateDB(c, service)
}

func (s *DBStore) CreateDB(c context.Context, srv ServiceType) (*gorm.DB, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	db, exists := s.dbs[c.Value("tenant_id").(string)]
	if exists {
		return db, nil
	}

	var addr string
	var db_type string
	switch srv {
	case Auth:
		// addr = os.Getenv("DB_ADDRESS_AUTH_SERVICE")
		// db_type = os.Getenv("DB_TYPE_AUTH_SERVICE")
		addr = "127.0.0.1:3306"
		db_type = "mysql"
	case Product:
		// addr = os.Getenv("DB_ADDRESS_PRODUCT_SERVICE")
		// db_type = os.Getenv("DB_TYPE_PRODUCT_SERVICE")
		addr = "127.0.0.1:3306"
		db_type = "mysql"
	case Planning:
		// addr = os.Getenv("DB_ADDRESS_PLANNING_SERVICE")
		// db_type = os.Getenv("DB_TYPE_PLANNING_SERVICE")
		addr = "127.0.0.1:3306"
		db_type = "mysql"
	case Subscription:
		// addr = os.Getenv("DB_ADDRESS_SUBSCRIPTION_SERVICE")
		// db_type = os.Getenv("DB_TYPE_SUBSCRIPTION_SERVICE")
		addr = "127.0.0.1:3306"
		db_type = "mysql"

	case Customer:
		// addr = os.Getenv("DB_ADDRESS_CUSTOMER_SERVICE")
		// db_type = os.Getenv("DB_TYPE_CUSTOMER_SERVICE")
		addr = "127.0.0.1:3306"
		db_type = "mysql"
	default:
		return nil, errors.New("")
	}

	tenant_id := c.Value("tenant_id").(string)
	user := c.Value("db_user").(string)
	pass := c.Value("db_pass").(string)

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, addr, tenant_id)

	switch db_type {
	case "mysql":
		return gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}
	return db, errors.New("")
}
