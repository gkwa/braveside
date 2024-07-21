# braveside

Purpose:


## example usage

```bash


```

## install braveside


on macos/linux:
```bash

brew install gkwa/homebrew-tools/braveside

```


on windows:

```powershell

TBD

```




## Todo

# Missing GitHub Markdown Widgets

### Subscript text

Example:
```
H<sub>2</sub>O
```
Renders as: H<sub>2</sub>O

### Superscript text

Example:
```
X<sup>2</sup>
```
Renders as: X<sup>2</sup>

### Footnotes

Example:
```
Here's a sentence with a footnote.[^1]

[^1]: This is the footnote.
```
Renders with a superscript number and the footnote at the bottom of the document.

### Color models (HEX, RGB, HSL)

Examples:
```
`#0969DA`
`rgb(9, 105, 218)`
`hsl(212, 92%, 45%)`
```
These render as colored circles in GitHub issues, pull requests, and discussions.

### Section links

Example:
```
[Link to Headers](#headers)
```
This creates a link to a section within the document.

### Task lists

Example:
```
- [x] Completed task
- [ ] Incomplete task
```
Renders as checkboxes that can be interactively checked on GitHub.

### Mentions (people and teams)

Example:
```
@github/support What do you think about these updates?
```
This notifies the mentioned team or individual on GitHub.

### Issue and pull request references

Example:
```
#123
```
This automatically links to issue or pull request number 123 in the current repository.

### External resource references (custom autolinks)

Example (if configured):
```
TICKET-123
```
This could automatically link to a JIRA ticket or other external resource.

### Emoji support

Example:
```
:smile: :heart: :thumbsup:
```
Renders as: 😄 ❤️ 👍

### Alerts

Examples:
```
> [!NOTE]
> This is a note.

> [!WARNING]
> This is a warning.
```
These render as distinctively styled alert boxes.

### HTML comments for hiding content

Example:
```
<!-- This content will not appear in the rendered Markdown -->
```
This text would be hidden in the rendered view.

### Disabling Markdown rendering

This is a GitHub UI feature that allows viewing the raw Markdown source.

### Specifying theme for image display

Example:
```html
<picture>
  <source media="(prefers-color-scheme: dark)" srcset="dark-image.png">
  <source media="(prefers-color-scheme: light)" srcset="light-image.png">
  <img alt="Text describing the image" src="default-image.png">
</picture>
```
This allows different images to be displayed based on the user's theme setting.

### Autolinked references and URLs

Example:
```
Visit https://github.com for more information.
```
GitHub automatically converts this into a clickable link.