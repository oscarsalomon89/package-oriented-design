package memorydb

type Book struct {
	ID       uint
	AuthorID uint
	Title    string
	Price    float64
}

var db []Book = []Book{
	{
		ID:       1,
		Title:    "Dune",
		Price:    1965,
		AuthorID: 1,
	},
	{
		ID:       2,
		Title:    "Cita con Rama",
		Price:    1974,
		AuthorID: 1,
	},
	{
		ID:       3,
		Title:    "Un guijarro en el cielo",
		Price:    500,
		AuthorID: 2,
	},
}

type LocalDB struct {
	storage []Book
}

func New() *LocalDB {
	return &LocalDB{storage: db}
}

func (db *LocalDB) SaveItem(item Book) {
	lastElement := db.storage[len(db.storage)-1]
	item.ID = lastElement.ID + 1

	db.storage = append(db.storage, item)
}

func (db *LocalDB) GetAll() []Book {
	return db.storage
}
