  
func Diaglog() {
framebox := gtk.NewVBox(false, 1)
    frame.Add(framebox)
 
    entry := gtk.NewEntry()
    entry.SetText("<Name>")
    framebox.Add(entry)
 
    buttons := gtk.NewHBox(false, 1)
 
    button := gtk.NewButtonWithLabel("Hello me")
    button.Clicked(func() {
        print("button clicked: ", button.GetLabel(), "\n")
        messagedialog := gtk.NewMessageDialog(
            button.GetTopLevelAsWindow(),
            gtk.DIALOG_MODAL,
            gtk.MESSAGE_INFO,
            gtk.BUTTONS_OK,
            "Hello, " + entry.GetText())
        messagedialog.Response(func() {
 
             
            messagedialog.Destroy()
        })
        messagedialog.Run()
    })
    buttons.Add(button)
	}
