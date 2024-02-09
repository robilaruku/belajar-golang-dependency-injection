package simple

type Database struct {
	Name string
}

// alias on golang
type DatabasePostgreSQL Database
type DatabaseMongoDB Database

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{Name: "PostgreSQL"})
}

func NewDatabaseRepository(postgreSQL *DatabasePostgreSQL, mongodb *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePostgreSQL: postgreSQL,
		DatabaseMongoDB: mongodb,
	}
}

