package utils

import (
	"os/exec"
)

type Editor struct {
	name string
	cmd  string
}

type CurrentEditor struct {
	editors     []Editor
	editorIndex int
}

func NewCurrentEditor() CurrentEditor {
	currentEditor := CurrentEditor{editors: []Editor{
		{
			name: "Visual Studio Code",
			cmd:  "code",
		},
		{
			name: "IntelliJ IDEA",
			cmd:  "idea",
		},
	}}

	// Check if the editor cmd is available
	for i, editor := range currentEditor.editors {
		isAvailable := currentEditor.checkEditorCmdAvailable(editor.cmd)

		if !isAvailable {
			println(editor.name + " is not available")
			currentEditor.removeEditor(i)
		}
	}

	return currentEditor
}

func (e *CurrentEditor) incrementCurrentEditorIndex() {
	e.editorIndex = (e.editorIndex + 1) % len(e.editors)
}

func (e *CurrentEditor) getCurrentEditor() Editor {
	return e.editors[e.editorIndex]
}

func (e *CurrentEditor) removeEditor(i int) {
	e.editors = append(e.editors[:i], e.editors[i+1:]...)
}

func (e *CurrentEditor) checkEditorCmdAvailable(cmdName string) bool {
	cmd := exec.Command("command", "-v", cmdName)

	if err := cmd.Run(); err != nil {
		return false
	}

	return true
}

var (
	currentEditor = NewCurrentEditor()
)
