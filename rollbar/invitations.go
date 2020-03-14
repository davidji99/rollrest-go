package rollbar

import "github.com/davidji99/simpleresty"

// InvitationService handles communication with the invitation related
// methods of the Rollbar API.
//
// Rollbar API docs: N/A
type InvitationService service

// InvitationResponse represents a response after inviting an user.
type InvitationResponse struct {
	ErrorCount *int        `json:"err,omitempty"`
	Result     *Invitation `json:"result,omitempty"`
}

// InvitationListResponse represents a response of all invitations.
type InvitationListResponse struct {
	ErrorCount *int          `json:"err,omitempty"`
	Result     []*Invitation `json:"result,omitempty"`
}

// Invitation represents an invitation in Rollbar (usually an user's invitation to a team).
type Invitation struct {
	ID           *int64  `json:"id,omitempty"`
	FromUserID   *int64  `json:"from_user_id,omitempty"`
	TeamID       *int64  `json:"team_id,omitempty"`
	ToEmail      *string `json:"to_email,omitempty"`
	Status       *string `json:"status,omitempty"`
	DateCreated  *int64  `json:"date_created,omitempty"`
	DateRedeemed *int64  `json:"date_redeemed,omitempty"`
}

// Get a invitation.
//
// Rollbar API docs: https://explorer.docs.rollbar.com/#operation/get-invitation
func (i *InvitationService) Get(inviteID int) (*InvitationResponse, *simpleresty.Response, error) {
	var result *InvitationResponse
	urlStr := i.client.http.RequestURL("/invite/%d", inviteID)

	// Set the correct authentication header
	i.client.setAuthTokenHeader(i.client.accountAccessToken)

	// Execute the request
	response, getErr := i.client.http.Get(urlStr, &result, nil)

	return result, response, getErr
}

// Cancel a invitation.
//
// Rollbar API docs: https://explorer.docs.rollbar.com/#operation/cancel-invitation
func (i *InvitationService) Cancel(inviteID int) (*simpleresty.Response, error) {
	urlStr := i.client.http.RequestURL("/invite/%d", inviteID)

	// Set the correct authentication header
	i.client.setAuthTokenHeader(i.client.accountAccessToken)

	// Execute the request
	response, getErr := i.client.http.Delete(urlStr, nil, nil)

	return response, getErr
}
