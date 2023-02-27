package repository

// Performs a check of input username and password against the database and
// if check is correct then username is changed
/* func (a *Db) ChangeUsername(ctx context.Context, creds core.Credentials) error {
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
}  */
