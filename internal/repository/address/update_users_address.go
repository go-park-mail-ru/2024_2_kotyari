package address

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (ar *AddressStore) UpdateUsersAddress(ctx context.Context, userID uint32, addressModel model.Addresses) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	ar.Log.Info("[AddressStore.UpdateUsersAddress] Started executing", slog.Any("request-id", requestID))

	const query = `
		INSERT INTO addresses (user_id, address)
		VALUES ($1, $2)
		ON CONFLICT (user_id)
		DO UPDATE SET 
			address = EXCLUDED.address;
	`

	_, err = ar.Db.Exec(ctx, query,
		userID,
		addressModel.Address)

	if err != nil {
		ar.Log.Error("[ AddressStore.UpdateUsersAddress ]Ошибка при обновлении адреса", slog.String("error", err.Error()))

		return err
	}

	return nil
}
