package subscription_domain

// SubscriptionError represents the error format response by OCM on a subscription request
type SubscriptionError struct {
	Kind     string `json:"kind"`
	ID       string `json:"id"`
	Href     string `json:"href"`
	Code     string `json:"code"`
	Reason   string `json:"reason"`
	Error    error  `json:"-"`
	Response []byte `json:"-"`
}

/*
{
    "kind": "Error",
    "id": "400",
    "href": "/api/clusters_mgmt/v1/errors/400",
    "code": "CLUSTERS-MGMT-400",
    "reason": "bad order value 'asdfasdfasdf'",
    "operation_id": "1grla68e486r9gqcbiifalhl9hpljhje"
}
*/
