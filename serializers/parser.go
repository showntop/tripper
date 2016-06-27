package serializers

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

var (
	ErrBadJSONAPIStructTag = errors.New("Bad jsonapi struct tag format")
	ErrBadJSONAPIID        = errors.New("id should be either string or int")
)

func parseRootToNode(model interface{}, included *map[string]*Node, sideload bool) (node *Node, err error) {
	node = new(Node)
	err = parseToNode(model, node, included, sideload)
	return
}

//parse the model to a node struct for json api format spec
func parseToNode(model interface{}, node *Node, included *map[string]*Node, sideload bool) (err error) {
	modelValue := reflect.ValueOf(model).Elem()
	for i := 0; i < modelValue.NumField(); i++ {

		structField := modelValue.Type().Field(i) //field
		fieldValue := modelValue.Field(i)         //value
		tag := structField.Tag.Get("jsonapi")     //tag
		if tag == "" {
			continue
		}

		fmt.Println(structField.Name + ":" + "" + ":" + tag)

		args := strings.Split(tag, ",")

		if len(args) < 1 {
			err = ErrBadJSONAPIStructTag
			break
		}
		node.Type = getStructType(model)

		usage := args[0]

		switch usage {
		case "inline":
			// if structField.Type.Kind() == reflect.Struct {
			parseToNode(fieldValue.Addr().Interface(), node, included, sideload)
			// }
		case "primary":
			id := fieldValue.Interface()
			// idType := id.(type)
			// _ = idType
			switch nId := id.(type) {
			case bson.ObjectId:
				node.Id = nId.Hex()
			case string:
				node.Id = nId
			case int:
				node.Id = strconv.Itoa(nId)
			case int64:
				node.Id = strconv.FormatInt(nId, 10)
			case uint64:
				node.Id = strconv.FormatUint(nId, 10)
			default:
				err = ErrBadJSONAPIID
				break
			}
		case "attr":
			var omitEmpty bool

			if len(args) > 2 {
				omitEmpty = args[2] == "omitempty"
			}

			if node.Attributes == nil {
				node.Attributes = make(map[string]interface{})
			}

			if fieldValue.Type() == reflect.TypeOf(time.Time{}) {
				t := fieldValue.Interface().(time.Time)

				if t.IsZero() {
					continue
				}

				node.Attributes[args[1]] = t.Unix()
			} else if fieldValue.Type() == reflect.TypeOf(new(time.Time)) {
				// A time pointer may be nil
				if fieldValue.IsNil() {
					if omitEmpty {
						continue
					}

					node.Attributes[args[1]] = nil
				} else {
					tm := fieldValue.Interface().(*time.Time)

					if tm.IsZero() && omitEmpty {
						continue
					}

					node.Attributes[args[1]] = tm.Unix()
				}
			} else {
				strAttr, ok := fieldValue.Interface().(string)

				if ok && strAttr == "" && omitEmpty {
					continue
				} else if ok {
					node.Attributes[args[1]] = strAttr
				} else {
					node.Attributes[args[1]] = fieldValue.Interface()
				}
			}
		case "relation":
			isSlice := fieldValue.Type().Kind() == reflect.Slice

			if (isSlice && fieldValue.Len() < 1) || (!isSlice && fieldValue.IsNil()) {
				continue
			}

			if node.Relationships == nil {
				node.Relationships = make(map[string]interface{})
			}

			if isSlice {
				relationship, err := parseToRelationships(args[1], fieldValue, included, sideload)

				if err == nil {
					d := relationship.Data
					if sideload {
						var shallowNodes []*Node

						for _, n := range d {
							// appendIncluded(included, n)
							shallowNodes = append(shallowNodes, toShallowNode(n))
						}

						node.Relationships[args[1]] = &RelationshipManyNode{Data: shallowNodes}
					} else {
						node.Relationships[args[1]] = relationship
					}
				} else {
					err = err
					break
				}
			} else {
				relationship, err := parseRootToNode(fieldValue.Interface(), included, sideload)
				if err == nil {
					if sideload {
						// appendIncluded(included, relationship)
						node.Relationships[args[1]] = &RelationshipOneNode{Data: toShallowNode(relationship)}
					} else {
						node.Relationships[args[1]] = &RelationshipOneNode{Data: relationship}
					}
				} else {
					err = err
					break
				}
			}
		default:
			err = ErrBadJSONAPIStructTag
		}
	}
	return
}

func parseToRelationships(relationName string, models reflect.Value, included *map[string]*Node, sideload bool) (*RelationshipManyNode, error) {
	var nodes []*Node

	if models.Len() == 0 {
		nodes = make([]*Node, 0)
	}

	for i := 0; i < models.Len(); i++ {
		n := models.Index(i).Interface()
		node, err := parseRootToNode(n, included, sideload)
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, node)
	}

	return &RelationshipManyNode{Data: nodes}, nil
}

func toShallowNode(node *Node) *Node {
	return &Node{
		Id:   node.Id,
		Type: node.Type,
	}
}
