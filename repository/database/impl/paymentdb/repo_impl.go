package paymentdb

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
	"gorm.io/gorm"
)

type defaultRepo struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &defaultRepo{db: db}
}

func (s *defaultRepo) FindPaymentSchemeById(ctx context.Context, id int) (resp *entity.PaymentScheme, err error) {
	err = s.db.WithContext(ctx).Take(&resp, "id = ?", id).Error
	return
}

func (s *defaultRepo) CreateTransaction(ctx context.Context, req *entity.Transaction) (err error) {
	err = s.db.WithContext(ctx).Create(req).Error
	return
}

func (s *defaultRepo) FindPaymentScheme(ctx context.Context) (resp []entity.PaymentScheme, err error) {
	err = s.db.WithContext(ctx).Find(&resp).Error
	return
}

func (s *defaultRepo) GetListTransactionDasboar(ctx context.Context, param paginate.Pagination) (resp []entity.Transaction, count int64, err error) {
	err = s.db.WithContext(ctx).Model(&entity.Transaction{}).Count(&count).Error
	if err != nil {
		return
	}
	err = s.db.WithContext(ctx).
		Preload("User").Scopes(paginate.Paginate(param.Page, param.Limit)).Find(&resp).Error
	return
}

func (s *defaultRepo) GetInstalmentUser(ctx context.Context, param paginate.Pagination) (resp []entity.Instalment, count int64, err error) {
	query := func(db *gorm.DB) *gorm.DB {
		db.Where("instalments.status = ?", entity.Paid)

		if param.PaymentSchemeId != 0 {
			db.Joins("JOIN transactions ON transactions.id = instalments.transaction_id").
			Joins("JOIN payment_schemes ON transactions.payment_scheme_id = payment_schemes.id").
			Where("payment_schemes.id = ?", param.PaymentSchemeId)
		}
		return db
	}

	err = s.db.WithContext(ctx).Model(&entity.Instalment{}).Scopes(query).Count(&count).Error
	if err != nil {
		return
	}
	err = s.db.WithContext(ctx).
		Scopes(paginate.Paginate(param.Page, param.Limit)).
		Scopes(query).
		Preload("Transaction.PaymentScheme").
		Preload("Transaction.User").
		Preload("OngoingInstalment").
		Find(&resp).Error
	return
}

func (s *defaultRepo) GetTransactionByUserId(ctx context.Context, userId int) (resp []entity.Transaction, err error) {
	err = s.db.WithContext(ctx).Find(&resp, "user_id = ?", userId).Error
	return
}

func (s *defaultRepo) GetListTransactionByUserIdAndStatus(ctx context.Context, userId int, status entity.TranscationStatus, param paginate.Pagination) (resp []entity.Transaction, count int64, err error) {
	query := func (db *gorm.DB) *gorm.DB  {
		return db.Where("user_id = ? AND status = ?", userId, status)
	}

	err = s.db.WithContext(ctx).Model(&entity.Transaction{}).Scopes(query).Count(&count).Error
	if err != nil {
		return
	}

	err = s.db.WithContext(ctx).Scopes(paginate.Paginate(param.Page, param.Limit)).
		Preload("Warehouse.District.Regency").
		Preload("Warehouse.WarehouseImg").
		Scopes(query).
		Find(&resp).Error
	return
}

func (s *defaultRepo) GetListTransactionData(ctx context.Context, param paginate.PaginationTrx) (resp []entity.Transaction, count int64, err error) {
	query := func(db *gorm.DB) *gorm.DB {
		if param.ProvinceId != 0 || param.Search != "" {
			db.Joins("JOIN warehouses ON transactions.warehouse_id = warehouses.id").
				Joins("JOIN districts ON warehouses.district_id = districts.id").
				Joins("JOIN regencies ON districts.regency_id = regencies.id").
				Joins("JOIN provinces ON regencies.province_id = provinces.id").
				Joins("JOIN users ON transactions.user_id = users.id")
			if param.ProvinceId != 0 {
				db.Where("provinces.id = ?", param.ProvinceId)
			}
			if param.Search != "" {
				db.Where("warehouses.name LIKE ?", "%"+param.Search+"%").
					Or("users.username LIKE ?", "%"+param.Search+"%")
			}
		}

		if param.Status != "" {
			db.Where("status = ?", param.Status)
		}

		return db
	}
	err = s.db.WithContext(ctx).Model(&entity.Transaction{}).Scopes(query).Count(&count).Error
	if err != nil {
		return
	}
	err = s.db.WithContext(ctx).Scopes(paginate.Paginate(param.Page, param.Limit)).
		Preload("User").
		Preload("Warehouse.District.Regency.Province").
		Preload("PaymentScheme").
		Scopes(query).
		Find(&resp).Error
	return
}

func (s *defaultRepo) GetTransactionById(ctx context.Context, transactionId string) (resp *entity.Transaction, err error) {
	err = s.db.WithContext(ctx).Preload("PaymentScheme").Preload("User").Preload("Warehouse").Take(&resp, "id = ?", transactionId).Error
	return
}

func (s *defaultRepo) GetTransactionDetailById(ctx context.Context, transactionId string) (resp *entity.Transaction, err error) {
	err = s.db.WithContext(ctx).Preload("PaymentScheme").
		Preload("Warehouse.District.Regency.Province").
		Preload("Warehouse.WarehouseImg").
		Preload("User").
		Take(&resp, "id = ?", transactionId).Error
	return
}

func (s *defaultRepo) BeginTrans(ctx context.Context) *gorm.DB {
	return s.db.Begin()
}

func (s *defaultRepo) CreateInstalment(ctx context.Context, tx *gorm.DB, req *entity.Instalment) (err error) {
	err = tx.WithContext(ctx).Create(req).Error
	return
}

func (s *defaultRepo) UpdateTransaction(ctx context.Context, tx *gorm.DB, req *entity.Transaction) (err error) {
	err = tx.WithContext(ctx).Save(req).Error
	return
}

func (s *defaultRepo) GetListInstalmentByTransactionId(ctx context.Context, transactionId string, param paginate.Pagination) (resp []entity.Instalment, count int64, err error) {
	query := func (db *gorm.DB) *gorm.DB {
		return db.Where("transaction_id = ?", transactionId)
	}

	err = s.db.WithContext(ctx).Model(&entity.Instalment{}).Scopes(query).Count(&count).Error
	if err != nil {
		return
	}

	err = s.db.WithContext(ctx).Scopes(paginate.Paginate(param.Page, param.Limit)).Scopes(query).Find(&resp).Error
	return
}

func (s *defaultRepo) GetListPaymentMethod(ctx context.Context) (resp []entity.PaymentMethod, err error) {
	err = s.db.WithContext(ctx).Find(&resp).Error
	return
}

func (s *defaultRepo) GetPaymentMethodById(ctx context.Context, id int) (resp *entity.PaymentMethod, err error) {
	err = s.db.WithContext(ctx).Take(&resp, "id = ?", id).Error
	return
}

func (s *defaultRepo) GetInstalmentById(ctx context.Context, id int) (resp *entity.Instalment, err error) {
	err = s.db.WithContext(ctx).Preload("OngoingInstalment").Take(&resp, "id = ?", id).Error
	return
} 

func (s *defaultRepo) UpdateInstalment(ctx context.Context, tx *gorm.DB, req *entity.Instalment) (err error) {
	err = tx.WithContext(ctx).Save(req).Error
	return
}

func (s *defaultRepo) CreateOngoingInstalment(ctx context.Context, req *entity.OngoingInstalment) (err error) {
	err = s.db.WithContext(ctx).Create(req).Error
	return
}

func (s *defaultRepo) UpdateOngoingInstalment(ctx context.Context, tx *gorm.DB, req *entity.OngoingInstalment) (err error) {
	err = tx.WithContext(ctx).Save(req).Error
	return
}

func (s *defaultRepo) FindOngoingInstalmentByXpayment(ctx context.Context, xpaymentId string) (resp *entity.OngoingInstalment, err error) {
	err = s.db.WithContext(ctx).Take(&resp, "x_payment = ?", xpaymentId).Error
	return
}
