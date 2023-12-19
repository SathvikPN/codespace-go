# codespace-go
A vscode setup for solving computational problems configured for Go with file based IO operations running on Linux or MacOS

Set Keybindings
---
> vscode-settings > command-palette > preferences: Open Keyboard Shortcuts (JSON) ==> keybindings.json

```json
[
    {
        "key": "shift+enter",
        "command": "workbench.action.tasks.runTask",
        "args": "Gorun"
    }
]
```