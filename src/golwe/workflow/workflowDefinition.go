/**
 * Created with IntelliJ IDEA.
 * User: remi
 * Date: 06/01/13
 * Time: 23:13
 * To change this template use File | Settings | File Templates.
 */
package workflow
import ("fmt"
	"encoding/xml"
	"io/ioutil")

type Start struct {
	Destination string `xml:"destination,attr"`
}

type End struct {
	Name         string `xml:"name,attr"`
}

type Goto struct{
	Destination string `xml:"destination,attr"`
}

type Task struct {
	Name         string `xml:"name,attr"`
	Processor    string `xml:"processor,attr"`
	Goto         []Goto `xml:"goto"`
}

type Fork struct {
	Name         string `xml:"name,attr"`
	Goto         []Goto `xml:"goto"`
}

type Call struct {
	Workflow     string `xml:"workflow,attr"`
}

type Wait struct {
	Name         string `xml:"name,attr"`
	Goto         Goto `xml:"goto"`
}

type WaitFor struct {
	Name        string `xml:"name,attr"`
	Goto        Goto `xml:"goto"`
	Duration 	int `xml:"duration,attr,omitempty"`
}


type Workflow struct {
	XMLName		xml.Name	`xml:"workflow"`
	Name       	string		`xml:"name,attr"`
	Start      	Start 		`xml:"start"`
	Task       	[]Task		`xml:"task"`
	Fork	  	[]Fork		`xml:"fork"`
	End			[]End		`xml:"end"`
	Wait	   	[]Wait		`xml:"wait"`
	WaitFor		[]WaitFor	`xml:"waitfor"`
}

func Extract(wfDefPath string ) *Workflow {

	wfDef := Workflow{}
	wfDefXML, err := ioutil.ReadFile(wfDefPath)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	err = xml.Unmarshal([]byte(wfDefXML), &wfDef)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	return &wfDef
}
