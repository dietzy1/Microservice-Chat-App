package client

/* func NewAccountClient() *accountclientv1.AccountServiceClient {
	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+os.Getenv("AUTH"),
		//"localhost:9000",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := accountclientv1.NewAccountServiceClient(conn)
	return &client
} */
