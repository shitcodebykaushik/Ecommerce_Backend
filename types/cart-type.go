package types
import "go.mongodb.org/mongo-driver/bson/primitive"
//Cart model "my_cart"

type cart struct {
	ID primitive.ObjectID  `json:"_id" bson:"_id,omitempty"`
	UserID primitive.ObjectID `json:"_id" bsom:"user_id" bson:"user_id"`
	Product primitive.ObjectID  `json:"product" bson:"product"`
	Checktout bool  `json:"chekctout" bson:"checkout"`
}

type CartClient struct {
	UserID string  `json:"user_id" bson:"user_id"`
	ProductID string  `json:"product_id" bson:"product_id"`
}