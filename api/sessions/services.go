package sessions

import "github.com/ariel17/efimeral/api/apierrors"

// NewSession builds a new container instance and returns its data associated
// or an error if it fails.
func NewSession(sr *SessionDistribution) (*Session, *apierrors.APIError) {
	return nil, nil
}

// FetchSession returns the session data for given session ID, or an error if
// it fails to fetch.
func FetchSession(id string) (*Session, *apierrors.APIError) {
	return nil, nil
}

// RemoveSession terminates and removes all data about given session ID. It
// returns its data at the delete moment or an error if the operation fails.
func RemoveSession(id string) (*Session, *apierrors.APIError) {
	return nil, nil
}
