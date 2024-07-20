package types


// Product model "products"

type Product struct {
	
	Name string `json:"name" bson:"name"`
	Description string   `json:"description" bson:"desciprtion"`
	Price string  `json:"price" bson:"price"`
	ImageUrl string  `json:"image_url" bson:"image_url"`
	MetaInfo string   `json:"MetaInfo" bson:"MetaInfo"`
	CreatedAt int64    `json:"created_at" bson:"created_at"`
	UpdatedAt int64    `json:"updated_at" bson:"update_at"`
}

type ProductCLient struct {

	Name string  `json:"name" bson: "name"`
	Description string `json:"description" bson:"description"`
	Price float64  `json:"price" bson:"price"`
	ImageUrl string  `json:"image_url" bson:"image_url"`
	MetaInfo map[string]interface {}  `json:"meta_info,omitempty" bson:"meta_info,omitempty"` 
}

type UpdateProduct struct {
	ID string `json:"id"`
	Name string `json:"name,omitempty"`
	Description string `json:"de scription,omitempty"`
	Price float64 `json:"price"`
}

