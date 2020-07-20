package main

//import (
//	"bytes"
//	"testing"
//	"os"
//	"fmt"
//	"log"
//	"net/http"
//	"github.com/tensorflow/tensorflow/tensorflow/go"
//	"io"
//)
//
//func Test_checkUrl(t *testing.T) {
//	t.Run("happy path", func(t *testing.T) {
//	})
//}
//
//
//func TestCheckArgs(t *testing.T){
//	// assert := assert.New(t)
//	// fmt.Printf("%t",checkArgs([]string{"first","https://cdn.pixabay.com/photo/2019/07/30/05/53/dog-4372036__340.jpg"}))
//	// assert.True(checkArgs([]string{"first","https://cdn.pixabay.com/photo/2019/07/30/05/53/dog-4372036__340.jpg"}))
//	// assert.False(checkArgs([]string{"first"}))
//	var tests = []struct {
//		name string
//       args []string
//       want bool
//   }{
//       {"not enough info",[]string{"1"}, false},
//       {"alright arguments",[]string{"1","www.google.com"}, true},
//   }
//    for _, tt := range tests {
//       testname := fmt.Sprintf(tt.name)
//       t.Run(testname, func(t *testing.T) {
//           ans := checkArgs(tt.args)
//           if ans != tt.want {
//               t.Errorf("got %t, want %t", ans, tt.want)
//           }
//       })
//   }
//}
//
//func imageToTensorTestHelp(url string) (*tensorflow.Tensor, error){
//response, e := http.Get("https://cdn.pixabay.com/photo/2019/07/30/05/53/dog-4372036__340.jpg")
//	if e != nil {
//		log.Fatalf("unable to get image from url: %v", e)
//	}
//	defer response.Body.Close()
//	var buf bytes.Buffer
//	io.Copy(&buf, response)
//
//	tensor, _ := tensorflow.NewTensor(buf.String())
//	graph, input, output, _ := getNormalizedGraph()
//	session, _ := tensorflow.NewSession(graph, nil)
//	normalized, err := session.Run(
//		map[tensorflow.Output]*tensorflow.Tensor{
//			input: tensor,
//		},
//		[]tensorflow.Output{
//			output,
//		},
//		nil)
//	return normalized[0], err
//
//}
//
//func TestImageToTensor(t *testing.T){
//	var tests = []struct {
//		name string
//       url string
//       want *tensorflow.Tensor
//       wantError error
//   }{
//       {"dog image","https://cdn.pixabay.com/photo/2019/07/30/05/53/dog-4372036__340.jpg",
//        imageToTensorTestHelp("https://cdn.pixabay.com/photo/2019/07/30/05/53/dog-4372036__340.jpg"),nil},
//   }
//    for _, tt := range tests {
//       testname := fmt.Sprintf(tt.name)
//       t.Run(testname, func(t *testing.T) {
//       	response, e := http.Get("https://cdn.pixabay.com/photo/2019/07/30/05/53/dog-4372036__340.jpg")
//			if e != nil {
//				log.Fatalf("unable to get image from url: %v", e)
//			}
//			defer response.Body.Close()
//           ans,err := imageToTensor(response.Body)
//           if (ans != tt.want || err != tt.wantError) {
//               t.Errorf("did not get the correct tensor")
//           }
//       })
//   }
//
//}
//
//
//func TestMain(m *testing.M) {
//	os.Args[0]="/go/src/imgrecognition/imgrecognition"
//	os.Args[1]="https://cdn.pixabay.com/photo/2019/07/30/05/53/dog-4372036__340.jpg"
//
//}


// func TestLabels(t *testing. T){
// 	os.Args[0]="/go/src/imgrecognition/imgrecognition"
// 	os.Args[1]="https://cdn.pixabay.com/photo/2019/07/30/05/53/dog-4372036__340.jpg"

// }