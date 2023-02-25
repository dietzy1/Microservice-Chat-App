package repository

/* func (a *Db) Register(ctx context.Context, creds core.Credentials) error {
	collection := a.mClient.Database("credentials").Collection("private")
	_, err := collection.InsertOne(ctx, creds)
	if err != nil {
		return err
	}
	return nil //idk if this is correct
}

// Performs a check against the database if username is taken
func (a *Db) CheckUsername(ctx context.Context, username string) bool {
	collection := a.mClient.Database("credentials").Collection("private")
	err := collection.FindOne(ctx, bson.M{"username": username})
	return err == nil
}

func (a *Db) CheckPassword(ctx context.Context, uuid string) (string, error) {
	var creds core.Credentials
	collection := a.mClient.Database("credentials").Collection("private")
	err := collection.FindOne(ctx, bson.M{"uuid": uuid}).Decode(&creds)
	if err != nil {
		return "", err
	}
	return creds.Password, nil
}

func (a *Db) Delete(ctx context.Context, creds core.Credentials) error {
	collection := a.mClient.Database("credentials").Collection("private")
	_, err := collection.DeleteOne(ctx, bson.M{"username": creds.Username})
	if err != nil {
		return err
	}
	return nil
}

// Performs a check of input username and password against the database and
// if check is correct then username is changed
func (a *Db) ChangeUsername(ctx context.Context, creds core.Credentials) error {
	collection := a.mClient.Database("credentials").Collection("private")
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": creds.Uuid}, bson.M{"$set": bson.M{"username": creds.Username}})
	if err != nil {
		return err
	}
	return nil
}

// Performs a check of input username and password against the database and
// if check is correct then password is changed
func (a *Db) ChangePassword(ctx context.Context, creds core.Credentials) error {
	collection := a.mClient.Database("credentials").Collection("private")
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": creds.Uuid}, bson.M{"$set": bson.M{"password": creds.Password}})
	if err != nil {
		return err
	}

	return nil
} */
