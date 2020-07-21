package main

import (
	"fmt"
	"github.com/tensorflow/tensorflow/tensorflow/go"
	"log"
	"net/http"
	// "os"
	"testing"
)

func Test_checkUrl(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
	})
}

func TestCheckArgs(t *testing.T) {
	// assert := assert.New(t)
	// fmt.Printf("%t",checkArgs([]string{"first","https://cdn.pixabay.com/photo/2019/07/30/05/53/dog-4372036__340.jpg"}))
	// assert.True(checkArgs([]string{"first","https://cdn.pixabay.com/photo/2019/07/30/05/53/dog-4372036__340.jpg"}))
	// assert.False(checkArgs([]string{"first"}))
	var tests = []struct {
		name string
		args []string
		want bool
	}{
		{"not enough info", []string{"1"}, false},
		{"alright arguments", []string{"1", "www.google.com"}, true},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf(tt.name)
		t.Run(testname, func(t *testing.T) {
			ans := checkArgs(tt.args)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}

type emptyValue struct{}

type classForFunctions interface {
	runSessionMock(*tensorflow.Session, *tensorflow.Tensor, tensorflow.Output, tensorflow.Output) ([]*tensorflow.Tensor, error)
	createTensor(value interface{}) (*tensorflow.Tensor, error)
}

type mock struct {
	Tensor  *tensorflow.Tensor
	Tensors []*tensorflow.Tensor
}

func (m mock) runSessionMock(_ *tensorflow.Session, _ *tensorflow.Tensor, _ tensorflow.Output,
	_ tensorflow.Output) ([]*tensorflow.Tensor, error) {
	return m.Tensors, nil
}

func (m mock) createTensorMock(_ interface{}) (*tensorflow.Tensor, error) {
	return m.Tensor, nil
}

func TestImageToTensor(t *testing.T) {
	mockInfo1 := emptyValue{}
	tensor1, _ := tensorflow.NewTensor(mockInfo1)
	mockInfo2 := emptyValue{}
	tensor2, _ := tensorflow.NewTensor(mockInfo2)
	var tensors = []*tensorflow.Tensor{tensor1, tensor2}
	var m = mock{tensor1, tensors}
	var tests = []struct {
		name      string
		url       string
		want      *tensorflow.Tensor
		wantError error
	}{
		{"dog image", "https://cdn.pixabay.com/photo/2019/07/30/05/53/dog-4372036__340.jpg",
			tensor1, nil},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf(tt.name)
		t.Run(testName, func(t *testing.T) {
			response, e := http.Get("https://cdn.pixabay.com/photo/2019/07/30/05/53/dog-4372036__340.jpg")
			if e != nil {
				log.Fatalf("unable to get image from url: %v", e)
			}
			defer response.Body.Close()
			ans, err := imageToTensor(response.Body, m.createTensorMock, m.runSessionMock, tensorflow.NewGraph(), tensorflow.Output{}, tensorflow.Output{})
			if ans != tt.want || err != tt.wantError {
				t.Errorf("did not get the correct tensor")
			}
		})
	}

}

//Labels tests
func TestLabels(t *testing.T) {
	var labels []Label
	labels = append(labels, Label{"hello", 5})
	labels = append(labels, Label{"word", 4})
	labels = append(labels, Label{"this", 5})
	labels = append(labels, Label{"is", 5})
	labels = append(labels, Label{"a", 5})
	var tests = []struct {
		name          string
		labels        []string
		probabilities []float32
		want          Labels
	}{
		{"normal labels", []string{"hello", "world", "this", "is", "a"}, []float32{5, 4, 3, 2, 1}, labels},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf(tt.name)
		t.Run(testname, func(t *testing.T) {
			response, e := http.Get("https://cdn.pixabay.com/photo/2019/07/30/05/53/dog-4372036__340.jpg")
			if e != nil {
				log.Fatalf("unable to get image from url: %v", e)
			}
			defer response.Body.Close()
			ans := getTopFiveLabels(tt.labels, tt.probabilities)
			if ans.Equals(tt.want) {
				t.Errorf("did not get the correct tensor")
			}
		})
	}
}


