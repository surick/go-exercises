package vm

// BaseViewModel struct
type BaseViewModel struct {
	Title string
	CurrentUser string
}

// setTitle func
func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
}

// setCurrentUser func
func (v *BaseViewModel) SetCurrentUser(username string) {
	v.CurrentUser = username
}
