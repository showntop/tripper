package serializers

type ObjectPayload struct {
	Data *Node `json:"data"`
	// Included []*Node            `json:"included,omitempty"`
	// Links    *map[string]string `json:"links,omitempty"`
}

type ListPayload struct {
	Data []*Node `json:"data"`
	// Included []*Node            `json:"included,omitempty"`
	// Links    *map[string]string `json:"links,omitempty"`
}

type Node struct {
	Type          string                 `json:"type"`
	Id            string                 `json:"id"`
	ClientId      string                 `json:"client-id,omitempty"`
	Attributes    map[string]interface{} `json:"attributes,omitempty"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
}

type RelationshipOneNode struct {
	Data *Node `json:"data"`
	// Links *map[string]string `json:"links,omitempty"`
}

type RelationshipManyNode struct {
	Data []*Node `json:"data"`
	// Links *map[string]string `json:"links,omitempty"`
}

// type ErrorNode struct {
// 	Status string `json:"status,omitempty"`
// 	Source string `json:"source,omitempty"`
// 	Title  string `json:"title,omitempty"`
// 	Detail string `json:"detail,omitempty"`
// }

// type ErrorPayload struct {
// 	Errors []*ErrorNode `json:"errors"`
// }
