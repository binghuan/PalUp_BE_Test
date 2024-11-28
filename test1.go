package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// Node represents a node in the linked list
type Node struct {
	Value interface{} `json:"value"`
	Next  *Node       `json:"next,omitempty"`
}

// Parse JSON body and convert array to linked list
func Test1(w http.ResponseWriter, r *http.Request) {
	type Input struct {
		Array []interface{} `json:"Array"`
	}

	var input Input
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &input)
	if err != nil || input.Array == nil {
		http.Error(w, "Invalid JSON or missing 'Array' field", http.StatusBadRequest)
		return
	}

	// Create a linked list from the array
	var head, current *Node
	for _, val := range input.Array {
		node := &Node{Value: val}
		if head == nil {
			head = node
			current = head
		} else {
			current.Next = node
			current = node
		}
	}

	// Convert linked list back to JSON for response
	response, err := json.Marshal(head)
	if err != nil {
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
