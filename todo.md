## Script
- [ ] chore: Dedupe adding alias to script (make script idempotent)

## CLI
- [ ] feat: Close cli when opening IDE?
- [ ] feat: Indent items to show subfolders - allow for opening/closing of folders?
- [ ] feat: Allow adding additional editors through flags

## Done
- [x] feat: Remove unavailable IDEs on startup
- [x] allow config through flags - add folder path flags
- [x] chore: ignore dot files when looking for repos
- [x] feat: allow default ide, and also selection of ide, switch button?
- [x] feat: Recursively display git repos
- [x] feat: Ignore folders from the list of ignores
- [x] chore: Measure how long `findRepos` runs
- [x] feat: Pressing enter on a list item opens the folder
- [x] feat: Write script to add cli as an alias to shell
- [x] feat: Add readme to document how to install cli

## Testing
- [ ] Test the folder flags
- [ ] Test repos being found and listed
- [ ] Test being able to open a an item with the current editor cmd

## Cancelled
- [-] feat: Allow setting array of folders to search for repos <!-- superseded -->
- [-] feat: Allow opening with code (vscode) or idea (intelliji) Blocked by config feat
- [-] feat: Add ability to save config <!-- Using cli flags instead ->
  - [-] spike: Figure out where the config can be saved - dot file?
  - [-] spike: yml file? json? what else? 
  - [-] Allow editing config when passing arg to cli e.g. `--edit-config -e`