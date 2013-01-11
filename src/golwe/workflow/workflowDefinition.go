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
	Childrens []*Node
	Name      string

}

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
	Name         string `xml:"name,attr"`
	Workflow     string `xml:"workflow,attr"`
	Goto         Goto `xml:"goto"`
}

type Wait struct {
	Name         string `xml:"name,attr"`
	Goto         Goto `xml:"goto"`
}

type WaitFor struct {
	Name         string `xml:"name,attr"`
	Goto         Goto `xml:"goto"`
	Duration 	int `xml:"duration,attr,omitempty"`
}


type WorkflowDefinition struct {
	XMLName		xml.Name	`xml:"workflow"`
	Name       	string		`xml:"name,attr"`
	Start      	Start 		`xml:"start"`
	Task       	[]Task		`xml:"task"`
	Fork	  	 []Fork		`xml:"fork"`
	End			[]End		`xml:"end"`
	Wait	   	[]Wait		`xml:"wait"`
	WaitFor		[]WaitFor	`xml:"waitfor"`
	Call		   []Call	`xml:"call"`
}

func ParseDefinition(wfDefPath string) *WorkflowDefinition {

	wfDef := WorkflowDefinition{}
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

func BuildWorkflowGraph(workflowDefinition *WorkflowDefinition) *Node {
	root := new(Node)
	root.Name = "start"
	root.Childrens = make([]*Node, 1)
	root.Childrens[0] = buildNodeGraph(root, workflowDefinition.Start.Destination, workflowDefinition)
	return root
}

// TODO mutualise code!!
func buildNodeGraph(root *Node, nodeName string, workflowDefinition *WorkflowDefinition) *Node {


	for i := range workflowDefinition.Task {
		if nodeName == workflowDefinition.Task[i].Name {
			node := new(Node)
			node.Name = nodeName
			node.Childrens = make([]*Node, len(workflowDefinition.Task[i].Goto))
			for k := range workflowDefinition.Task[i].Goto {
				//fmt.Printf("children of node %v : %v", nodeName, workflowDefinition.Task[i].Goto[k])
				node.Childrens[k] = buildNodeGraph(node, workflowDefinition.Task[i].Goto[k].Destination, workflowDefinition)
			}
			return node
		}
	}

	for i := range workflowDefinition.Fork {
		if nodeName == workflowDefinition.Fork[i].Name {
			node := new(Node)
			node.Name = nodeName
			node.Childrens = make([]*Node, len(workflowDefinition.Fork[i].Goto))
			for k := range workflowDefinition.Fork[i].Goto {
				//fmt.Printf("children of node %v : %v", nodeName, workflowDefinition.Fork[i].Goto[k])
				node.Childrens[k] = buildNodeGraph(node, workflowDefinition.Fork[i].Goto[k].Destination, workflowDefinition)
			}
			return node
		}
	}

	for i := range workflowDefinition.Wait {
		if nodeName == workflowDefinition.Wait[i].Name {
			node := new(Node)
			node.Name = nodeName
			node.Childrens = make([]*Node, 1)
			//fmt.Printf("children of node %v : %v", nodeName, workflowDefinition.Wait[i].Goto)
			node.Childrens[0] = buildNodeGraph(node, workflowDefinition.Wait[i].Goto.Destination, workflowDefinition)
			return node
		}
	}

	for i := range workflowDefinition.Call { // TODO gestion du branchement vers le sous workflow
		if nodeName == workflowDefinition.Call[i].Name {
			node := new(Node)
			node.Name = nodeName
			node.Childrens = make([]*Node, 1)
			//fmt.Printf("children of node %v : %v", nodeName, workflowDefinition.Call[i].Goto)
			node.Childrens[0] = buildNodeGraph(node, workflowDefinition.Call[i].Goto.Destination, workflowDefinition)
			return node
		}
	}

	for i := range workflowDefinition.WaitFor {
		if nodeName == workflowDefinition.WaitFor[i].Name {
			node := new(Node)
			node.Name = nodeName
			node.Childrens = make([]*Node, 1)
			//fmt.Printf("children of node %v : %v", nodeName, workflowDefinition.WaitFor[i].Goto)
			node.Childrens[0] = buildNodeGraph(node, workflowDefinition.WaitFor[i].Goto.Destination, workflowDefinition)
			return node
		}
	}

	for i := range workflowDefinition.End {
		if nodeName == workflowDefinition.End[i].Name {
			node := new(Node)
			node.Name = nodeName
			return node
		}
	}

	panic("Node name:"+nodeName+", does not exist")

	toto := new (Node)
	toto.Name="Aie"
	return toto
}

func ToString(node *Node) string {
	// TODO Buffer de tableau de lignes pour préparer les lignes de branches à imprimmer
	if (len(node.Childrens) == 0) {
		return "end"
	}else {
		res :=node.Name+"->"

		if (len(node.Childrens)>1 ){
			res+="{"
		}

		for i := range node.Childrens {
			res += ToString(node.Childrens[i])
		}
		if (len(node.Childrens)>1 ){
			res+="}"
		}
		return res
	}
	return ""
}
