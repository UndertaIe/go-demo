package gomicro

import micro "go-micro.dev/v4"

func (Gomicro) Demo() {
	// create a new service
	service := micro.NewService(
		micro.Name("helloworld"),
	)

	// initialise flags
	service.Init()

	// start the service
	service.Run()
}
