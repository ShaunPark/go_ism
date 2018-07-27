package gid

type MySQLGIDChecker struct {
	GidChecker
}

func (gidChecker *MySQLGIDChecker) CheckGID(gid string) bool {
	return false
}
