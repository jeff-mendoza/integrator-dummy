package model

type WebhookRequest struct {
	ID             string                      `json:"id" binding:"required"`
	CreatedAt      string                      `json:"created_at"`
	Amount         int64                       `json:"amount" binding:"required"`
	Payment        PaymentInfo                 `json:"payment"`
	AdditionalInfo AdditionalPaymentIntentInfo `json:"additional_info"`
	State          string                      `json:"state"`
}

type PaymentInfo struct {
	ID    int64  `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
}

type AdditionalPaymentIntentInfo struct {
	ExternalReference string `json:"external_reference,omitempty"`
}


type CreateConfigRequest struct {
	CallerID string `json:"caller_id"`
	Token    string `json:"token"`
}

type CreateConfigResponse struct {
	CallerID string `json:"caller_id"`
	Token    string `json:"token"`
}
