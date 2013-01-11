/**
 * Created with IntelliJ IDEA.
 * User: remi
 * Date: 06/01/13
 * Time: 23:13
 */
package workflow

import ("fmt"
	/*"image"
	"image/color"
	"image/draw"
	"os"
	"image/png"*/
	"strconv"
)

func DrawWorkflow(wf *Node) {
	/*fmt.Printf("Generating image for %v", wf.Name)
	m := image.NewRGBA(image.Rect(0, 0, 300, 200))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	f, err := os.Create(wf.Name+".png")

	if err != nil {
		panic(err)
	}

	png.Encode(f, m)
	*/

	for i := range wf.Childrens {
		//DrawWorkflow(wf.Childrens[i])
		fmt.Printf("\ncoucou_"+strconv.Itoa(i))
		DrawWorkflow(wf.Childrens[i])
	}

}

