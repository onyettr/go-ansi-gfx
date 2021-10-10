package main

import "fmt"

const FG_BLACK  =  30    /*!< Foreground BLACK  */
const FG_RED    =  31    /*!< Foreground RED    */
const FG_GREEN  =  32    /*!< Foreground GREEN  */
const FG_YELLOW =  33    /*!< Foreground YELLOW */
const FG_BLUE   =  34    /*!< Foreground BLUE   */
const FG_MAGENTA=  35    /*!< Foreground MAGENTA*/
const FG_CYAN   =  36    /*!< Foreground CYAN   */
const FG_WHITE  =  37    /*!< Foreground WHITE  */

const BG_BLACK   = 40    /*!< Background BLACK  */
const BG_RED     = 41    /*!< Background RED    */
const BG_GREEN   = 42    /*!< Background GREEN  */
const BG_YELLOW  = 43    /*!< Background YELLOW */
const BG_BLUE    = 44    /*!< Background BLUE   */
const BG_MAGENTA = 45    /*!< Background MAGENTA*/
const BG_CYAN    = 46    /*!< Background CYAN   */
const BG_WHITE   = 47    /*!< Background WHITE  */

func ScreenErase() {
   fmt.Printf("\033[2J")
}

func  ScreenGotoXY(x int, y int, colour int) {
  fmt.Printf("\033[%d;%df\033[%dm",y,x,colour);
}

func ScreenSetColour(colour int) {
  fmt.Printf("\033[0;%dm", colour);
}

func ScreenSetBackGroundColour(colour int) {
  fmt.Printf("\033[0;%dm", colour);
}

func ScreenReset() {
  fmt.Printf("\033[0m");
}

func main() {
	var i int
	var str = "hello"
	
	fmt.Println("Hello world!")

	i = 100
	fmt.Printf("Type %T %d\n", i, i)
	fmt.Printf("Type %T %s\n", str, str)

	ScreenSetColour(FG_RED)
	fmt.Printf("Type %T %s\n", str, str)	
	/*	ScreenGotoXY(10, 20, FG_BLUE) */
	ScreenSetColour(FG_BLUE)
	fmt.Printf("Type %T %s\n", str, str)	
	ScreenSetColour(FG_GREEN)
	fmt.Printf("Type %T %s\n", str, str)	

	/*	ScreenErase() */
}
