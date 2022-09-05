package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/widgets"
)

func main() {
	// needs to be called once before you can start using QML/Quick
	widgets.NewQApplication(len(os.Args), os.Args)

	// use the material style
	// the other inbuild styles are:
	// Default, Fusion, Imagine, Universal
	quickcontrols2.QQuickStyle_SetStyle("Material")

	// create the quick view
	// with a minimum size of 800*600
	// set the window title to "Basic Window"
	// and let the root item of the view resize itself to the size of the view automatically
	view := quick.NewQQuickView(nil)
	view.SetMinimumSize(core.NewQSize2(800, 600))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetTitle("Basic Window")

	// load the local qml file
	view.SetSource(core.QUrl_FromLocalFile("cmd/qml/main.qml"))

	// make the view visible
	view.Show()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	widgets.QApplication_Exec()
}