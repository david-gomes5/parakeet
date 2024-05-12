package utils

type Editor struct {
	name string
	cmd  string
}

type CurrentEditor struct {
	editors     []Editor
	editorIndex int
}

func (e *CurrentEditor) incrementCurrentEditorIndex() {
	e.editorIndex = (e.editorIndex + 1) % len(e.editors)
}

func (e *CurrentEditor) getCurrentEditor() Editor {
	return e.editors[e.editorIndex]
}

var (
	currentEditor = CurrentEditor{
		// What if we want to add more editors?
		// TODO: implement reading from a config file?
		editors: []Editor{
			{
				name: "Visual Studio Code",
				cmd:  "code",
			},
			{
				name: "IntelliJ IDEA",
				cmd:  "idea",
			},
		},
	}
)
