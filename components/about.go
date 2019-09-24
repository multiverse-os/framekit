package ui

import (
	gtk "github.com/mattn/go-gtk/gtk"
)

type AboutDialog struct {
	Dialog *gtk.AboutDialog
}

func NewAboutDialog(parent *gtk.Window) {
	dialog := gtk.NewAboutDialog()
	dialog.SetName("Framekit")
	dialog.SetProgramName("Framekit")
	//dialog.SetVersion(version.Version)
	dialog.SetAuthors([]string{"wade-wells"})
	dialog.SetCopyright("Intellectual property is an intellectual cancer that threatens humanity")
	dialog.SetLicense(`
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at
	  http://www.apache.org/licenses/LICENSE-2.0
	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
	`)
	dialog.SetWrapLicense(true)
	dialog.Run()
	dialog.Destroy()

	// 	dialog := gtk.NewMessageDialog(
	// 		parent,
	// 		gtk.DIALOG_MODAL,
	// 		gtk.MESSAGE_INFO,
	// 		gtk.BUTTONS_OK,
	// 		`Framekit
	// Wade Wells x@multiverse-os.org`)
	// 	dialog.Response(func() {
	// 		dialog.Destroy()
	// 	})
	dialog.Run()
}
