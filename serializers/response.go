package serializers

import (
	"encoding/json"
	"log"
)

// func MarshalErrorPayload(errors []interface) ([]byte, error){

// }

func MarshalObjectPayload(model interface{}) ([]byte, error) {
	included := make(map[string]*Node)

	node, err := parseRootToNode(model, &included, true)
	if err != nil {
		return nil, err
	}

	payload := &ObjectPayload{Data: node}

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func MarshalManyPayload(models []interface{}) ([]byte, error) {
	var data []*Node
	included := make(map[string]*Node)

	for i := 0; i < len(models); i++ {
		model := models[i]
		log.Printf("%v", model)
		node, err := parseRootToNode(model, &included, true)
		if err != nil {
			return nil, err
		}
		data = append(data, node)
	}

	if len(models) == 0 {
		data = make([]*Node, 0)
	}

	payload := &ListPayload{
		Data: data,
		// Included: nodeMapValues(&included),
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return b, nil
}
