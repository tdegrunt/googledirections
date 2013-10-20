package googledirections

import (
	"testing"
)

func TestSimple(t *testing.T) {

	directions, err := NewDirections("Amsterdam, The Netherlands", "Haarlem, The Netherlands")
	if err != nil {
		t.Error("Expected nil")
	} 

	err = directions.Get()
	if err != nil {
		t.Error("Expected nil")
	} 
	
	if directions.GetDistance() != 20806 {
		t.Errorf("Expected distance to be 20806, found %d", directions.GetDistance())
	}
	
}

func TestNotFound(t *testing.T) {

	directions, err := NewDirections("abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz")
	if err != nil {
		t.Error("Expected nil")
	} 

	err = directions.Get()
	if err != nil {
		t.Error("Expected nil")
	} 
	
	if directions.Status != "NOT_FOUND" {
		t.Errorf("Expected status to be NOT_FOUND, got: %d", directions.Status)
	}
	
}

func TestZeroResults(t *testing.T) {

	directions, err := NewDirections("Amsterdam, The Netherlands", "New York, NY, USA")
	if err != nil {
		t.Error("Expected nil")
	} 

	err = directions.Get()
	if err != nil {
		t.Error("Expected nil")
	} 
	
	if directions.Status != "ZERO_RESULTS" {
		t.Errorf("Expected status to be ZERO_RESULTS, got: %d", directions.Status)
	}
	
}

func TestFrench(t *testing.T) {

	directions, err := NewDirections("Créteil, France", "Asnières-sur-Oise, France")
	if err != nil {
		t.Error("Expected nil")
	} 

	err = directions.Get()
	if err != nil {
		t.Error("Expected nil")
	} 
	
	if directions.Status != "OK" {
		t.Errorf("Expected status to be OK, got: %d", directions.Status)
	}

	if directions.GetDistance() != 53398 {
		t.Errorf("Expected distance to be 53398, found %d", directions.GetDistance())
	}	
}