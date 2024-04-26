package productmap

const rootDirectory = ""

type ProductMap struct {
	Directory
}

func NewProductMap() ProductMap {
	return ProductMap{
		Directory: newDirectory(rootDirectory, nil),
	}
}
