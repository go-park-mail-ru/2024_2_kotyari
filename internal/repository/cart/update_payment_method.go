package cart

import "context"

func (cs *CartsStore) UpdatePaymentMethod(ctx context.Context, userID uint32, method string) error {
	const query = `UPDATE users SET preferred_payment_method = $1 WHERE id = $2`
	_, err := cs.db.Exec(ctx, query, method, userID)
	if err != nil {
		cs.log.Error("[CartsStore.UpdateUserPaymentMethod] Error updating payment method")
		return err
	}
	return nil
}
