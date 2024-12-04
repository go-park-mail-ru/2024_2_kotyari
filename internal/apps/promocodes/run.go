package promocodes

import "log/slog"

func (p *PromoCodesApp) Run() error {
	err := p.reader.Read()
	if err != nil {
		p.log.Error("[PromoCodesApp.Run] Error reading messages",
			slog.String("error", err.Error()))

		return err
	}

	return nil
}
