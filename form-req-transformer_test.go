package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server *httptest.Server
	//Test Data TV
	userJsonTVClaim = `{"event_id":"GQQPPJJtfd","event_type":"form_response","form_response":{"form_id":"krCLZ1","token":"4","submitted_at":"2017-11-06T10:25:13Z","hidden":{"email":"xxxxx","name":"xxxxx","phone":"xxxxx","policy":"xxxxx"},"definition":{"id":"krCLZ1","title":"TV claim","fields":[{"id":"XdVgAjnucXRP","title":"Was there a theft or a deliberate event by a 3rd party?","type":"yes_no","allow_multiple_selections":false,"allow_other_choice":false},{"id":"AZbSqcXTKlED","title":"Please enter a  crime/loss reference number.","type":"short_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"63241151","title":"Please upload a detailed image of the damage","type":"file_upload","allow_multiple_selections":false,"allow_other_choice":false},{"id":"63241244","title":"Please upload any photos of proof of purchase / ownership / box / receipt","type":"file_upload","allow_multiple_selections":false,"allow_other_choice":false},{"id":"63241572","title":"Are you aware of anything else relevant to your claim that you would like to advise us of at this stage?","type":"long_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"63241542","title":"If you have any other insurance or warranties covering your TV, please advise us of the company name.","type":"short_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"63241330","title":"When did the incident happen?","type":"date","allow_multiple_selections":false,"allow_other_choice":false},{"id":"QYabj3uz2gs9","title":"In as much detail as possible, please provide a full written account of what has happened to your TV, including where it happened.","type":"long_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"WwzqHPb0K9Wv","title":"Make of the TV","type":"short_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"u2dzBoYFjbRA","title":"Serial number of the TV","type":"short_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"63241939","title":"We have made the following assumptions about your property, you and anyone living with you","type":"legal","allow_multiple_selections":false,"allow_other_choice":false},{"id":"lGCVB9tse6Re","title":"Model number of the TV","type":"short_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"XiKbgOrX1OAp","title":"Do you still have the TV in your possession?","type":"yes_no","allow_multiple_selections":false,"allow_other_choice":false},{"id":"63391165","title":"What was the purchase price of the TV?","type":"number","allow_multiple_selections":false,"allow_other_choice":false}]},"answers":[{"type":"boolean","boolean":true,"field":{"id":"XdVgAjnucXRP","type":"yes_no"}},{"type":"text","text":"DU00000000","field":{"id":"AZbSqcXTKlED","type":"short_text"}},{"type":"file_url","file_url":"https://admin.typeform.com/form/results/file/download/krCLZ1/63241151/97750857a9d3-22228398_888332207989782_5676320169463975335_n.jpg","field":{"id":"63241151","type":"file_upload"}},{"type":"file_url","file_url":"https://admin.typeform.com/form/results/file/download/krCLZ1/63241244/e6cc84865de7-22405371_888332421323094_6861338905885136899_n.jpg","field":{"id":"63241244","type":"file_upload"}},{"type":"text","text":"None","field":{"id":"63241572","type":"long_text"}},{"type":"text","text":"None","field":{"id":"63241542","type":"short_text"}},{"type":"date","date":"2017-10-01","field":{"id":"63241330","type":"date"}},{"type":"text","text":"what happened","field":{"id":"QYabj3uz2gs9","type":"long_text"}},{"type":"text","text":"LG","field":{"id":"WwzqHPb0K9Wv","type":"short_text"}},{"type":"text","text":"SerialNo","field":{"id":"u2dzBoYFjbRA","type":"short_text"}},{"type":"boolean","boolean":true,"field":{"id":"63241939","type":"legal"}},{"type":"text","text":"ModelNo","field":{"id":"lGCVB9tse6Re","type":"short_text"}},{"type":"boolean","boolean":true,"field":{"id":"XiKbgOrX1OAp","type":"yes_no"}},{"type":"number","number":20000,"field":{"id":"63391165","type":"number"}}]}}`
	//Test Data Strom
	userJsonStormClaim = `{"event_id":"SChin7eteE","event_type":"form_response","form_response":{"form_id":"H8mm3s","token":"6","submitted_at":"2017-11-06T14:13:27Z","hidden":{"email":"mectors@gmail.com","name":"Maarten","phone":"12345678","policy":"12345678"},"definition":{"id":"H8mm3s","title":"Storm surge claim","fields":[{"id":"j79cNctIvogK","title":"If it is safe and possible to do so, please provide images of the damage to both the outside and the inside of your home.","type":"file_upload","allow_multiple_selections":false,"allow_other_choice":false},{"id":"kaRKsWGqupVP","title":"Are you still have possession of the damage items (i.e. damaged guttering)?","type":"yes_no","allow_multiple_selections":false,"allow_other_choice":false},{"id":"63920548","title":"If there has been any recent maintenance carried out on your home, please describe it","type":"long_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"NBKeGLDeG1pa","title":"If you have any other insurance or warranties covering your home, please advise us of the company name.","type":"short_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"cWMGSX3JIi78","title":"We have made the following assumptions about your property, you and anyone living with you","type":"legal","allow_multiple_selections":false,"allow_other_choice":false},{"id":"nAz5fZvtiuLO","title":"When did the incident happen?","type":"date","allow_multiple_selections":false,"allow_other_choice":false},{"id":"ZSaL9YKNdYHe","title":"In as much detail as possible, please use the text box below to describe the full extent of the damage to your home and how you discovered it.","type":"long_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"63920509","title":"Please describe the details of the condition of your home prior to discovering the damage","type":"long_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"hhmzTJtsEobN","title":"Are you aware of anything else relevant to your claim that you would like to advise us of at this stage?","type":"long_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"64049754","title":"Where did the incident happen? (City/town name)","type":"short_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"63907334","title":"Would you like to upload more images?","type":"yes_no","allow_multiple_selections":false,"allow_other_choice":false}]},"answers":[{"type":"file_url","file_url":"https://admin.typeform.com/form/results/file/download/H8mm3s/j79cNctIvogK/aaee97a02b39-1.cow_on_a_roof.jpg","field":{"id":"j79cNctIvogK","type":"file_upload"}},{"type":"boolean","boolean":true,"field":{"id":"kaRKsWGqupVP","type":"yes_no"}},{"type":"text","text":"no","field":{"id":"63920548","type":"long_text"}},{"type":"text","text":"no","field":{"id":"NBKeGLDeG1pa","type":"short_text"}},{"type":"boolean","boolean":true,"field":{"id":"cWMGSX3JIi78","type":"legal"}},{"type":"date","date":"2017-11-02","field":{"id":"nAz5fZvtiuLO","type":"date"}},{"type":"text","text":"A cow was blown on top of my roof and destroyed tiles.","field":{"id":"ZSaL9YKNdYHe","type":"long_text"}},{"type":"text","text":"perfect","field":{"id":"63920509","type":"long_text"}},{"type":"text","text":"nothing","field":{"id":"hhmzTJtsEobN","type":"long_text"}},{"type":"text","text":"london","field":{"id":"64049754","type":"short_text"}},{"type":"boolean","boolean":false,"field":{"id":"63907334","type":"yes_no"}}]}}`
)

func TestHandler(t *testing.T) {

	//Convert string to reader
	readerTV := strings.NewReader(userJsonTVClaim)
	//Create request with JSON body
	reqTV, err := http.NewRequest("POST", "", readerTV)
	//Create request with JSON body
	reqStorm, err := http.NewRequest("POST", "", strings.NewReader(userJsonStormClaim))
	// empty request
	reqEmpty, err := http.NewRequest("POST", "", strings.NewReader(""))

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	// ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	//TEST CASES
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{"Status200TV", args{rr, reqTV}},
		{"Status200Storm", args{rr, reqStorm}},
		{"Status200Storm", args{rr, reqEmpty}},
	}
	for _, tt := range tests {
		// call ServeHTTP method
		// directly and pass Request and ResponseRecorder.
		handler := http.HandlerFunc(Handler)
		handler.ServeHTTP(tt.args.w, tt.args.r)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
			t.Errorf("content type header does not match: got %v want %v",
				ctype, "application/json")
		}

		//check response content
		res, err := ioutil.ReadAll(rr.Body)
		if err != nil {
			t.Error(err) //Something is wrong while read res
		}

		got := TranformedData{}
		err = json.Unmarshal(res, &got)

		if err != nil && got.TicketDetails.Ticket.Subject != "" {
			t.Errorf("%q. compute weather risk() = %v, want %v", tt.name, got, "non empty")
		}

	}

}
