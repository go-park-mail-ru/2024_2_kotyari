package promocodes

import (
	"fmt"
	"log/slog"
	"net"
)

func (p *PromoCodesApp) Run() error {
	go func() {
		err := p.reader.Read()
		if err != nil {
			p.log.Error("[PromoCodesApp.Run] Error reading messages",
				slog.String("error", err.Error()))
		}
	}()

	listener, err := net.Listen("tcp",
		fmt.Sprintf("%s:%s", p.grpcConf.Address, p.grpcConf.Port))
	if err != nil {
		p.log.Error("[PromoCodesApp.Run] Failed to start grpc server",
			slog.String("error", err.Error()))

		return err
	}

	p.log.Info("[PromoCodesApp.Run] Started listening",
		slog.String("server-addr", fmt.Sprintf("%s:%s", p.grpcConf.Address, p.grpcConf.Port)))

	if err = p.server.Serve(listener); err != nil {
		p.log.Error("[PromoCodesApp.Run] Failed to start listen",
			slog.String("error", err.Error()))

		return err
	}
	return nil
}
