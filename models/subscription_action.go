package models

type SubscriptionAction struct {
	Id            string `json:"id,omitempty"`
	Type          string `json:"type,omitempty"`
	EffectiveDate string `json:"effective_date,omitempty"`
	NewPlanId     string `json:"new_plan_id,omitempty"`
}
