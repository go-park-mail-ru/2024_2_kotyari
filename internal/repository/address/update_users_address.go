package address

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (ar *AddressStore) UpdateUsersAddress(ctx context.Context, addressID uint32, addressModel model.Address) error {

	const query = `
		INSERT INTO addresses (user_id, city, street, house, flat)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (user_id)
    	DO UPDATE SET city = EXCLUDED.city, 
				  street = EXCLUDED.street, 
				  house = EXCLUDED.house, 
				  flat = EXCLUDED.flat;
	`

	_, err := ar.db.Exec(ctx, query,
		addressModel.City,
		addressModel.Street,
		addressModel.House,
		addressModel.Flat,
		addressID)

	if err != nil {
		ar.log.Error("[ AddressStore.UpdateUsersAddress ]Ошибка при обновлении адреса", slog.String("error", err.Error()))
		return err
	}
	ar.log.Info("[ AddressStore.UpdateUsersAddress ]Обновление адреса")
	return nil
}
