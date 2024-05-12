![build badge](https://github.com/david-gomes5/parakeet/actions/workflows/go.yml/badge.svg)

# Parakeet

This is a golang cli to quickly be able to, list, filter and open repos, on your favourite IDE with ease.

Using BubbleTea as the cli framework.

> [!WARNING]
> STILL HEAVILY UNDER CONSTRUCTION, USE AT OWN RISK

## How to install?
> [!NOTE]
> Changes to the installation script still needs to be made to ensure we're creating an alias that adds folder path flags. Atleast one folder path flag is **REQUIRED**

Run the `add-cli-as-alias.sh` shell script
```bash
$ sh ./add-cli-as-alias.sh
```

Finally, you'll have to manually edit the `alias` created from the script to add the `-f` with paths pointing to your git repositories (This will be uneeded in later versions as installation should help with configuring cli)
## How to update the cli?
Build the cli with the following command:
```bash
$ make build-cli
```

## Flags

| Flag | Type                 | Description                                                                                           |
| ---- | -------------------- | ----------------------------------------------------------------------------------------------------- |
| `-f` | `string`, `[]string` | Use to point to folder directories with git repositories e.g. `parakeet -f "/Users/davidgomes/repos"` |