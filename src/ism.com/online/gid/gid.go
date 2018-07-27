package gid

type GidCheckerInterface struct {
}

type GidChecker interface {
	CheckGID(gid string) bool
}
