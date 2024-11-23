package csat

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (c *CSATService) CreateCSAT(ctx context.Context, csat model.CSAT) error {
	_, err := c.repository.GetCSAT(ctx, csat)
	if err != nil {
		if errors.Is(err, errs.NoUserCSAT) {
			c.log.Info("[CSATService.CreateCSAT] No csat for user, creating")

			err = c.repository.CreateCSAT(ctx, csat)
			if err != nil {
				c.log.Info("[CSATService.CreateCSAT] Error creating csat", slog.String("error", err.Error()))

				return err
			}

			return nil
		}
		c.log.Info("[CSATService.CreateCSAT] Unexpected error", slog.String("error", err.Error()))

		return err
	}

	return errs.UserCSATAlreadyExists
}
