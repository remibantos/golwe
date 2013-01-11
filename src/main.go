/**
 * Created with IntelliJ IDEA.
 * User: remi
 * Date: 06/01/13
 * Time: 17:33
 */
package main

import ("golwe/workflow"
 "fmt")

const (
	workflowPath = "src/resources/workflowsample.xml";
)

func main() {

	wfDef := workflow.ParseDefinition(workflowPath)

	fmt.Printf("XMLName: %#v\n", wfDef.XMLName)
	fmt.Printf("Workflow name: %v\n", wfDef.Name)
	fmt.Printf("Start: %v\n", wfDef.Start)
	fmt.Printf("End: %v\n", wfDef.End)
	fmt.Printf("Task nodes: %v\n", wfDef.Task)
	fmt.Printf("Fork nodes: %v\n", wfDef.Fork)
	fmt.Printf("Wait nodes: %v\n", wfDef.Wait)
	fmt.Printf("WaitFor nodes: %v\n", wfDef.WaitFor)

	fmt.Println("--------------------")

	wf := workflow.BuildWorkflowGraph(wfDef);

	fmt.Printf("Generated workflow graph: %v\n",workflow.ToString(wf))

}
