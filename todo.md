## Script
- [ ] chore: Dedupe adding alias to script (make script idempotent)
## CLI
- [ ] feat: Add ability to save config
  - [ ] spike: Figure out where the config can be saved - dot file?
  - [ ] spike: yml file? json? what else? 
  - [ ] Allow editing config when passing arg to cli e.g. `--edit-config -e`
- [ ] feat: Allow opening with code (vscode) or idea (intelliji) Blocked by config feat
- [ ] feat: Allow setting array of folders to search for repos
- [ ] feat: Close cli when opening IDE
- [ ] feat: indent items to show subfolders - allow for opening/closing of folders?

## Done
- [x] feat: Recursively display git repos
- [x] feat: Ignore folders from the list of ignores
- [x] chore: Measure how long `findRepos` runs
- [x] feat: Pressing enter on a list item opens the folder
- [x] feat: Write script to add cli as an alias to shell
- [x] feat: Add readme to document how to install cli