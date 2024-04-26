package productmap

const rootDirectory = ""

type ProductMap struct {
	directory
}

func NewProductMap() ProductMap {
	return ProductMap{
		directory: newDirectory(rootDirectory, nil),
	}
}
