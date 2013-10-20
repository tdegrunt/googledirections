package googledirections

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// Things to avoid
const (
	AvoidToll = "tolls"
	AvoidHighways = "highways"
)

// Travel Modes
const (
	TravelModeDriving = "driving"
	TravelModeWalking = "walking"
	TravelModeBicycling = "bicycling"
	TravelModeTransit = "transit"
	
)

// Location represents a latitude and longitude combination
type Location struct {
	Lat float64
	Lng float64
}

// Bounds
type Bounds struct {
	Northeast Location
	Soutwest  Location
}

// Step represents a step in a leg
type Step struct {
	Distance struct {
		Text  string
		Value int
	}
	Duration struct {
		Text  string
		Value int
	}
	EndLocation      Location `json:"end_location"`
	StartLocation    Location `json:"start_location"`
	HtmlInstructions string   `json:"html_instructions"`
	Polyline         struct {
		Points string
	}
	TravelMode string `json:"travel_mode"`
}

// Step represents a leg in a direction
type Leg struct {
	Distance struct {
		Text  string
		Value int
	}
	Duration struct {
		Text  string
		Value int
	}
	EndAddress    string   `json:"end_address"`
	EndLocation   Location `json:"end_location"`
	StartAddress  string   `json:"start_address"`
	StartLocation Location `json:"start_location"`
	Steps         []Step
}

// Route represents a route
type Route struct {
	Bounds           Bounds
	Copyrights       string
	Legs             []Leg
	OverviewPolyline struct {
		Points string
	} `json:"overview_polyline"`
	Summary string
}

// Represents a complete directions result
type Directions struct {
	origin string
	destination string
	baseURL string
	language string
	alternative bool
	sensor bool
	mode string
	avoid string
	// Possible routes
	Routes []Route
	// Status of the response
	Status string
}

// Creates a new google directions client
func NewDirections(origin, destination string) (*Directions, error) {
	var d Directions

	d.baseURL = "http://maps.googleapis.com/maps/api/directions/json"
	d.language = "en"
	d.alternative = false
	d.sensor = false
	d.mode = "driving"

	d.origin = origin
	d.destination = destination
	
	return &d, nil
}

// Sets the travel mode (driving/walking/bicycling/transit)
func (d *Directions)SetTravelMode(mode string) {
	d.mode = mode
}

// Sets what to avoid (tolls/highways)
func (d *Directions)SetAvoid(avoid string) {
	d.avoid = avoid
}

// Sets alternatives
func (d *Directions)SetAllowAlternatives(allow bool) {
	d.alternative = allow
}

// Retrieves the directions results from Google
func (d *Directions)Get() error {
	v := url.Values{}
	v.Set("origin", d.origin)
	v.Add("destination", d.destination)
	v.Add("sensor", strconv.FormatBool(d.sensor))
	v.Add("mode", d.mode)
	v.Add("language", d.language)
	v.Add("alternative", strconv.FormatBool(d.alternative))
	v.Add("avoid", d.avoid)
	//v.Add("ie", "UTF8")

	url := d.baseURL + "?" + v.Encode()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&d)
	if err != nil {
		return err
	}
	return nil
}

// Returns the distance for the first route and it's first leg 
func (d *Directions)GetDistance() int {
	return d.Routes[0].Legs[0].Distance.Value
}
