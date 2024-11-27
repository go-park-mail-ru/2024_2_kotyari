package profile_service

func (app *ProfilesApp) Stop() {
	app.gRPCServer.GracefulStop()
	app.log.Info("gRPC server stopped gracefully")
}
