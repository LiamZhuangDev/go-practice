// Polymorphism means the same interface can represent multiple concrete types, and the same code can work with all of them transparently.
// The same code can take many forms of behavior, depending on the concrete type behind it.
// "One interface, many implementations."

package interface_example

type Database interface {
	Connect() error
	Query(q string) (interface{}, error)
	Close() error
}

type MySQL struct {
	Connection string
}

func (db *MySQL) Connect() error {
	return nil
}

func (db *MySQL) Query(q string) (interface{}, error) {
	return []string{"Nothing in MySQL"}, nil
}

func (db *MySQL) Close() error {
	return nil
}

type PostgreSQL struct {
	Connection string
}

func (db *PostgreSQL) Connect() error {
	return nil
}

func (db *PostgreSQL) Query(q string) (interface{}, error) {
	return []string{"Nothing in PostgreSQL"}, nil
}

func (db *PostgreSQL) Close() error {
	return nil
}

// Demonstrates polymorphism with the Database interface
func ExecuteQuery(db Database, query string) (interface{}, error) {
	if err := db.Connect(); err != nil {
		return nil, err
	}
	defer db.Close()

	if result, err := db.Query(query); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
