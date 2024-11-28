package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestTest1 checks the behavior of the Test1 handler
func TestTest1(t *testing.T) {
	// 定義測試案例
	tests := []struct {
		name         string
		requestBody  map[string]interface{}
		expectedCode int
		expectedBody map[string]interface{}
	}{
		{
			name: "Valid Array Input",
			requestBody: map[string]interface{}{
				"Array": []interface{}{"a", "b", "c", "d", "e"},
			},
			expectedCode: http.StatusOK,
			expectedBody: map[string]interface{}{
				"value": "a",
				"next": map[string]interface{}{
					"value": "b",
					"next": map[string]interface{}{
						"value": "c",
						"next": map[string]interface{}{
							"value": "d",
							"next": map[string]interface{}{
								"value": "e",
							},
						},
					},
				},
			},
		},
		{
			name:         "Missing Array Field",
			requestBody:  map[string]interface{}{},
			expectedCode: http.StatusBadRequest,
			expectedBody: nil,
		},
		{
			name:         "Invalid JSON",
			requestBody:  nil, // nil 表示請求體為空
			expectedCode: http.StatusBadRequest,
			expectedBody: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// 準備請求體
			var body []byte
			if test.requestBody != nil {
				body, _ = json.Marshal(test.requestBody)
			}

			// 模擬 HTTP 請求
			req := httptest.NewRequest("POST", "/test-1", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			// 呼叫目標處理器
			Test1(rec, req)

			// 檢查 HTTP 狀態碼
			if rec.Code != test.expectedCode {
				t.Errorf("expected status code %d, got %d", test.expectedCode, rec.Code)
			}

			// 如果有期望的響應體，進一步檢查
			if test.expectedBody != nil {
				var actualBody map[string]interface{}
				err := json.Unmarshal(rec.Body.Bytes(), &actualBody)
				if err != nil {
					t.Fatalf("unable to parse response body: %v", err)
				}

				if !compareJSON(actualBody, test.expectedBody) {
					t.Errorf("expected body %v, got %v", test.expectedBody, actualBody)
				}
			}
		})
	}
}

// compareJSON compares two JSON objects
func compareJSON(a, b map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for key, valA := range a {
		valB, exists := b[key]
		if !exists {
			return false
		}
		switch valA := valA.(type) {
		case map[string]interface{}:
			valB, ok := valB.(map[string]interface{})
			if !ok {
				return false
			}
			if !compareJSON(valA, valB) {
				return false
			}
		default:
			if valA != valB {
				return false
			}
		}
	}
	return true
}

// TestLinkedList verifies that Test1 correctly converts an array to a linked list.
func TestLinkedList(t *testing.T) {
	// Test input and expected output
	input := map[string]interface{}{
		"Array": []interface{}{"a", "b", "c", "d", "e"},
	}
	expectedOutput := &Node{
		Value: "a",
		Next: &Node{
			Value: "b",
			Next: &Node{
				Value: "c",
				Next: &Node{
					Value: "d",
					Next: &Node{
						Value: "e",
					},
				},
			},
		},
	}

	// Marshal the input into JSON for the request body
	requestBody, _ := json.Marshal(input)

	// Create a POST request to the handler
	req := httptest.NewRequest("POST", "/test-1", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Call the handler
	Test1(rec, req)

	// Verify the status code
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	// Parse the response body into a Node struct
	var actualOutput Node
	err := json.Unmarshal(rec.Body.Bytes(), &actualOutput)
	if err != nil {
		t.Fatalf("error unmarshalling response body: %v", err)
	}

	// Compare the actual linked list to the expected linked list
	if !compareLinkedLists(&actualOutput, expectedOutput) {
		t.Errorf("linked list mismatch. Expected: %+v, Got: %+v", expectedOutput, actualOutput)
	}
}

// compareLinkedLists recursively compares two linked lists for equality
func compareLinkedLists(a, b *Node) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || a.Value != b.Value {
		return false
	}
	return compareLinkedLists(a.Next, b.Next)
}
