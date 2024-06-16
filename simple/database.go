package simple

type Database struct {
	Name string
}

type DatabasePostgre Database
type DatabaseMongoDB Database

func NewDatabasePostgre() *DatabasePostgre {
	return (*DatabasePostgre)(&Database{Name: "Postgre"})
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}

type DatabaseRepository struct {
	DatabasePostgre *DatabasePostgre
	DatabaseMongoDB *DatabaseMongoDB
}

func NewDatabaseRepository(postgre *DatabasePostgre, mongo *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePostgre: postgre,
		DatabaseMongoDB: mongo,
	}
}