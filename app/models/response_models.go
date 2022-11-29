package models

// ClassworkResponse represents a JSON response
// to the Classwork POST request.
type ClassworkResponse struct {
	HTTPError             // Error, if one is attached to the response
	Classwork []Classwork `json:"classwork"` // The resulting classwork
}

// LoginResponse represents a JSON response
// to the Login POST request.
type LoginResponse struct {
	HTTPError // Error, if one is attached to the response
}

// IPRResponse represents a JSON response
// to the IPR POST request.
type IPRResponse struct {
	HTTPError       // Error, if one is attached to the response
	IPR       []IPR `json:"ipr"` // The resulting IPR(s)
}

// ReportCardResponse represents a JSON response
// to the Report Card POST request.
type ReportCardResponse struct {
	HTTPError             // Error, if one is attached to the response
	ReportCard ReportCard `json:"reportCard"` // The resulting report card
}

// ScheduleResponse represents a JSON response
// to the Schedule POST request.
type ScheduleResponse struct {
	HTTPError          // Error, if one is attached to the response
	Schedule  Schedule `json:"schedule"` // The resulting schedule
}

// TranscriptResponse represents a JSON response
// to the Transcript POST request.
type TranscriptResponse struct {
	HTTPError             // Error, if one is attached to the response
	Transcript Transcript `json:"transcript"` // The resulting transcript
}
