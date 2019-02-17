package changelog

// Data is the information that is used to create a
// changelog
type Data struct {
	Version         string
	Date            string
	Fixes           []string
	Features        []string
	BreakingChanges []string
	Contributors    []string
}
