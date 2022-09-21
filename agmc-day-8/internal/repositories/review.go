package repositories

// func (r *repositories) CreateReview(review models.Review) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	qry := r.mg.Database(utils.GoDotEnvVariable("MONGODB_NAME")).Collection(utils.GoDotEnvVariable("MONGODB_COLLECTION_REVIEWS"))

// 	_, err := qry.InsertOne(ctx, review)
// 	return err
// }

// func (r *repositories) GetReviews() (*[]models.Review, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	var reviews []models.Review
// 	qry := r.mg.Database(utils.GoDotEnvVariable("MONGODB_NAME")).Collection(utils.GoDotEnvVariable("MONGODB_COLLECTION_REVIEWS"))

// 	results, err := qry.Find(ctx, bson.M{})
// 	if err != nil {
// 		return &reviews, err
// 	}

// 	//reading from the db in an optimal way
// 	defer results.Close(ctx)
// 	for results.Next(ctx) {
// 		var review models.Review
// 		if err = results.Decode(&review); err != nil {
// 			return &reviews, err
// 		}

// 		reviews = append(reviews, review)
// 	}

// 	return &reviews, err
// }

// func (r *repositories) UpdateReview(review primitive.M, id string) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	objId, _ := primitive.ObjectIDFromHex(id)
// 	qry := r.mg.Database(utils.GoDotEnvVariable("MONGODB_NAME")).Collection(utils.GoDotEnvVariable("MONGODB_COLLECTION_REVIEWS"))

// 	result, err := qry.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": review})
// 	if err != nil {
// 		return err
// 	}

// 	var updateReview models.Review
// 	if result.MatchedCount == 1 {
// 		err := qry.FindOne(ctx, bson.M{"id": objId}).Decode(&updateReview)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return err
// }

// func (r *repositories) DeleteReview(id string) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	objId, _ := primitive.ObjectIDFromHex(id)
// 	qry := r.mg.Database(utils.GoDotEnvVariable("MONGODB_NAME")).Collection(utils.GoDotEnvVariable("MONGODB_COLLECTION_REVIEWS"))

// 	result, err := qry.DeleteOne(ctx, bson.M{"id": objId})

// 	if result.DeletedCount < 1 {
// 		return err
// 	}

// 	return err
// }

// func (r *repositories) GetReview(id string) (*models.Review, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	objId, _ := primitive.ObjectIDFromHex(id)
// 	qry := r.mg.Database(utils.GoDotEnvVariable("MONGODB_NAME")).Collection(utils.GoDotEnvVariable("MONGODB_COLLECTION_REVIEWS"))

// 	var review models.Review
// 	err := qry.FindOne(ctx, bson.M{"id": objId}).Decode(&review)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &review, nil
// }
