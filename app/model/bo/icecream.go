package bo

type Icecream struct {
	tableName             struct{} `sql:"icecream_store"`
	Id                    int64    `pg:",pk" schema:"id" conform:"trim,omitempty" json:"id,omitempty" sql:"id"`
	Name                  string   `pg:"name"`
	ImageClosed           string   `pg:"image_closed"`
	ImageOpen             string   `pg:"image_open"`
	Description           string   `pg:"description"`
	Story                 string   `pg:"story"`
	AllergyInfo           string   `pg:"allergy_info"`
	SourcingValues        string   `pg:"sourcing_values"`
	Ingredients           string   `pg:"ingredients"`
	DietaryCertifications string   `pg:"dietary_certifications"`
	ProductId             string   `pg:"product_id"`
	//11
}

func NewIcecream() *Icecream {
	return &Icecream{}
}
