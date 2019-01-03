package vm

// BaseViewModel struct
type BaseViewModel struct {
	Title string
}

// setTitle func
func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
}
