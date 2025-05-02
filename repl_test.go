package main

import (
    "reflect"
    "testing"
)


func TestCleanInput(t *testing.T) {
    tests := []struct {
        input    string
        expected []string
    }{
        {"Hello World", []string{"hello", "world"}},
        {"  This is a test.  ", []string{"this", "is", "a", "test."}},
        {"MIXED Case Letters", []string{"mixed", "case", "letters"}},
        {"Special   Characters!@#", []string{"special", "characters!@#"}},
        {"", []string{}}, // Empty string
    }

    for _, test := range tests {
        result := cleanInput(test.input)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("cleanInput(%q) = %v; want %v", test.input, result, test.expected)
        }
    }
}
