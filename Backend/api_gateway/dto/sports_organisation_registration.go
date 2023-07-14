package dto

import (
	application_pb "common/proto/application/generated"
	auth_pb "common/proto/auth/generated"
)

type SportsOrganisationRegistration struct {
	Account            auth_pb.Account
	SportsOrganisation application_pb.SportsOrganisation
}

type SportsOrganisationRegistrationResponse struct {
	AccountID            string
	SportsOrganisationID string
}
