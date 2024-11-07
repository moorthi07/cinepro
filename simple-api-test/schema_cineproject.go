package api

type CineProject struct {
	Description	string	`json:"description"`
	Id	int64	`json:"id"`
	Title	string	`json:"title"`
}

// Checks if all of the required fields for CineProject are set
// and validates all of the constraints for the object.
func (obj *CineProject) Validate() error {
	if obj == nil {
		return nil
	}
	return nil
}

