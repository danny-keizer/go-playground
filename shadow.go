package main

import "fmt"

// This defines the database transaction struct.
type DatabaseTransaction struct {
	database databaseConnector
}

// This defines the database connector interface.
type databaseConnector interface {
	create() string
	read() string
	update() string
	delete() string
}

// This agent speaks to postgresql.
type postgresqlAgent struct{}

func (p *postgresqlAgent) create() string {
	return "Created something inside the postgresql database."
}

func (p *postgresqlAgent) read() string {
	return "Read something inside the postgresql database."
}

func (p *postgresqlAgent) update() string {
	return "Updated something inside the postgresql database."
}

func (p *postgresqlAgent) delete() string {
	return "Deleted something inside the postgresql database."
}

// This agent speaks to mariadb.
type mariadbAgent struct{}

func (p *mariadbAgent) create() string {
	return "Created something inside the mariadb database."
}

func (p *mariadbAgent) read() string {
	return "Read something inside the mariadb database."
}

func (p *mariadbAgent) update() string {
	return "Updated something inside the mariadb database."
}

func (p *mariadbAgent) delete() string {
	return "Deleted something inside the mariadb database."
}

// Boilerplate.

func NewDatabaseTransaction(connector databaseConnector) *DatabaseTransaction {
	return &DatabaseTransaction{connector}
}

func dbCreate(transaction DatabaseTransaction) string {
	return transaction.database.create()
}

func dbRead(transaction DatabaseTransaction) string {
	return transaction.database.read()
}

func dbUpdate(transaction DatabaseTransaction) string {
	return transaction.database.update()
}

func dbDelete(transaction DatabaseTransaction) string {
	return transaction.database.delete()
}

// Then main entry point.
func main() {

	var databaseType string = "postgresql"

	fmt.Print("Select which database to modify (postgresql or mariadb): ")
	fmt.Scan(&databaseType)

	switch databaseType {

	case "postgresql":
		myTransaction := NewDatabaseTransaction(&postgresqlAgent{})
		fmt.Println(dbCreate(*myTransaction))
		fmt.Println(dbRead(*myTransaction))
		fmt.Println(dbUpdate(*myTransaction))
		fmt.Println(dbDelete(*myTransaction))

	case "mariadb":
		myTransaction := NewDatabaseTransaction(&postgresqlAgent{})
		fmt.Println(dbCreate(*myTransaction))
		fmt.Println(dbRead(*myTransaction))
		fmt.Println(dbUpdate(*myTransaction))
		fmt.Println(dbDelete(*myTransaction))
	default:
		fmt.Println("Unsupported database type, the following are supported: postgresql, mariadb.")

	}
}
