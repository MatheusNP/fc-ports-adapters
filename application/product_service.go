package application

type ProductService struct {
	persistence ProductPersistenceInterface
}

func NewProductService(persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{
		persistence: persistence,
	}
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	return s.persistence.Get(id)
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price

	if _, err := product.IsValid(); err != nil {
		return &Product{}, err
	}

	return s.persistence.Save(product)
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	if err := product.Enable(); err != nil {
		return &Product{}, err
	}

	return s.persistence.Save(product)
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	if err := product.Disable(); err != nil {
		return &Product{}, err
	}

	return s.persistence.Save(product)
}
