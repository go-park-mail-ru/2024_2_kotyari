package address

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (ar *AddressStore) UpdateUsersAddress(ctx context.Context, addressID uint32, addressModel model.Address) error {
	const query = `
		WITH upsert_address AS (
			INSERT INTO addresses (user_id, city, street, house, flat)
			VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT (user_id)
			DO UPDATE SET 
				city = EXCLUDED.city,
				street = EXCLUDED.street,
				house = EXCLUDED.house,
				flat = EXCLUDED.flat
			RETURNING user_id, city
		)
			UPDATE users
			SET city = upsert_address.city
			FROM upsert_address
			WHERE users.id = upsert_address.user_id;
	`

	_, err := ar.Db.Exec(ctx, query,
		addressID,
		addressModel.City,
		addressModel.Street,
		addressModel.House,
		*addressModel.Flat)

	if err != nil {
		fmt.Println(addressModel, addressID)
		ar.Log.Error("[ AddressStore.UpdateUsersAddress ]Ошибка при обновлении адреса", slog.String("error", err.Error()))
		return err
	}

	return nil
}
