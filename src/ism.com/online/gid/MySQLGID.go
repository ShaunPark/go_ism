package gid

import "ism.com/common/db"

type MySQLGIDChecker struct {
	GidChecker
}

func (gidChecker *MySQLGIDChecker) CheckGID(gid string) bool {
	dbConn := db.GetDatabase()

	stmtIns, err := dbConn.Prepare("INSERT INTO ISM_GID VALUES( ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	_, err = stmtIns.Exec(gid) // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return false
}
