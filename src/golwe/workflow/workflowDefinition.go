/**
 * Created with IntelliJ IDEA.
 * User: remi
 * Date: 06/01/13
 * Time: 23:13
 */
package workflow
import ("fmt"
	"encoding/xml"
	"io/ioutil")

type Node struct{

}

type Start struct {
	Destination string `xml:"destination,attr"`
	Node
}

type End struct {
	Name         string `xml:"name,attr"`
	Node
}

type Goto struct{
	Destination string `xml:"destination,attr"`
	computedNode Node
}

type Task struct {
	Name         string `xml:"name,attr"`
	Processor    string `xml:"processor,attr"`
	Goto         []Goto `xml:"goto"`
	Node
}

type Fork struct {
	Name         string `xml:"name,attr"`
	Goto         []Goto `xml:"goto"`
	Node
}

type Call struct {
	Workflow     string `xml:"workflow,attr"`
	Node
}

type Wait struct {
	Name         string `xml:"name,attr"`
	Goto         Goto `xml:"goto"`
	Node
}

type WaitFor struct {
	Name        string `xml:"name,attr"`
	Goto        Goto `xml:"goto"`
	Duration 	int `xml:"duration,attr,omitempty"`
	Node
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
