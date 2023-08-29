# golang-bubbletea-tool-template

This is my personal template for creating Golang Bubbletea apps.

## Why?

Reinventing the wheel sucks and I can bake my experience into this and never think about it again.

## Major learning points

- Don't panic in your main model's
  Init! [Until this is fixed your application will hang really badly](https://github.com/charmbracelet/bubbletea/issues/904).
  And even once it is fixed it's better to return `tea.Quit` to shut down cleanly anyway.
    - Instead, each model should do the following:
        - Have a `New()` function to initialize it
        - Contain a private, boolean variable called initialized that gets set in `New()`
        - In the `Init()` function, check if the model is initialized and if not print an error message and
          return `tea.Quit`
