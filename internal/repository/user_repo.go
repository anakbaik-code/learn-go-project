package repository

import (
	"context"
	"database/sql"
	"go-dbsqlc/db"
	"go-dbsqlc/internal/domain"
)

type TxBiang interface {
	db.DBTX
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}
type UserRepository interface {
	GetById(ctx context.Context, id int64) (domain.User, error)
	Create(ctx context.Context, user domain.User, addresses []domain.Address) (domain.User, error)
	List(ctx context.Context) ([]domain.User, error)
	Update(ctx context.Context, user domain.User, addresses []domain.Address) error
	Delete(ctx context.Context, id int64) error
	UpdateAvatar(ctx context.Context, user domain.User) error
}

type userRepository struct {
	db      TxBiang
	queries *db.Queries
}

func NewUserRepository(db db.DBTX, queries *db.Queries) UserRepository {
	txDB, ok := db.(TxBiang)
	if !ok {
		panic("database must support transactions (TxBiang)")
	}

	return &userRepository{
		db:      txDB, // Masuk ke struct dengan aman
		queries: queries,
	}
}

func (r *userRepository) GetById(ctx context.Context, id int64) (domain.User, error) {
	u, err := r.queries.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	dbAddress, err := r.queries.GetAddressesByUserID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	// Mapping domain Address
	var domainAddresses []domain.Address
	for _, addr := range dbAddress {
		domainAddresses = append(domainAddresses, domain.Address{
			ID:      int64(addr.ID),
			UserID:  int64(addr.UserID),
			Street:  addr.Street,
			City:    addr.City,
			Country: addr.Country,
		})
	}

	return domain.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		AvatarUrl: u.AvatarUrl.String,
		Addresses: domainAddresses,
	}, nil
}

func (r *userRepository) Create(ctx context.Context, user domain.User, addresses []domain.Address) (domain.User, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.User{}, err
	}
	defer tx.Rollback()
	qtx := r.queries.WithTx(tx)
	result, err := qtx.CreateUser(ctx, db.CreateUserParams{
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		return domain.User{}, err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return domain.User{}, err
	}

	for _, address := range addresses {
		_, err := qtx.CreateUserAddress(ctx, db.CreateUserAddressParams{
			UserID:  int64(userId),
			Street:  address.Street,
			City:    address.City,
			Country: address.Country,
		})
		if err != nil {
			return domain.User{}, err
		}
	}

	if err := tx.Commit(); err != nil {
		return domain.User{}, err
	}

	u, err := r.queries.GetUser(ctx, userId)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Addresses: addresses,
	}, nil
}

func (r *userRepository) List(ctx context.Context) ([]domain.User, error) {
	users, err := r.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	dbAddresses, err := r.queries.GetAllUserAddresses(ctx)
	if err != nil {
		return nil, err
	}

	// mapping
	addressMap := make(map[int64][]domain.Address)
	for _, addr := range dbAddresses {
		uId := int64(addr.UserID)

		addressMap[uId] = append(addressMap[uId], domain.Address{
			ID:      int64(addr.ID),
			UserID:  addr.UserID,
			Street:  addr.Street,
			City:    addr.City,
			Country: addr.Country,
		})
	}

	var result []domain.User
	for _, user := range users {
		result = append(result, domain.User{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Addresses: addressMap[user.ID],
		})
	}
	return result, nil
}

func (r *userRepository) Update(ctx context.Context, user domain.User, addresses []domain.Address) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := r.queries.WithTx(tx)
	if err := qtx.UpdateUser(ctx, db.UpdateUserParams{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}); err != nil {
		return err
	}

	if err := qtx.DeleteAddressesByUserID(ctx, user.ID); err != nil {
		return err
	}

	for _, address := range addresses {
		_, err := qtx.CreateUserAddress(ctx, db.CreateUserAddressParams{
			UserID:  user.ID,
			Street:  address.Street,
			City:    address.City,
			Country: address.Country,
		})
		if err != nil {
			return err
		}
	}
	return tx.Commit()

}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	return r.queries.DeleteUser(ctx, id)
}

func (r *userRepository) UpdateAvatar(ctx context.Context, user domain.User) error {
	arg := db.UpdateAvatarParams{
		ID: user.ID,
		AvatarUrl: sql.NullString{
			String: user.AvatarUrl,
			Valid:  user.AvatarUrl != "",
		},
	}
	return r.queries.UpdateAvatar(ctx, arg)
}
