package core

import (
	"os"
	"os/exec"
	"testing"
)

func TestFullMarkdownRendering(t *testing.T) {
	input := `---
filetype: product
test:
  - this and that
  - test2
x:
  "y":
    - a
    - b
---

# Hello, World!

This is a **test** file for our Goldmark AST roundtrip.

## Features

1. Lists
2. *Italic*
3. **Bold**

> Blockquotes are supported too.

| Column 1 | Column 2 |
|----------|----------|
| Cell 1   | Cell 2   |

- [ ] Task 1
- [x] Task 2

Here's some ` + "`inline code`" + ` and a code block:

` + "```go" + `
func main() {
 fmt.Println("Hello, World!")
}
` + "```" + `

![Seaweed Salad photo](https://static.spotapps.co/spots/a4/3ebb855c2348c68c7b94a4956d9662/full)

---

[OpenAI](https://www.openai.com)

~~strikethrough~~

1. First item
 - Subitem
 - Another subitem
2. Second item

Term
: Definition

Here's a sentence with a footnote.[^1]

[^1]: This is the footnote.

:smile: :heart: :thumbsup:

When $a \ne 0$, there are two solutions to $(ax^2 + bx + c = 0)$ and they are $$ x = {-b \pm \sqrt{b^2-4ac} \over 2a} $$
`

	output, err := ProcessMarkdown([]byte(input), false)
	if err != nil {
		t.Fatalf("Failed to process markdown: %v", err)
	}

	// Write input and output to temporary files
	tmpInput := "tmp_input.md"
	tmpOutput := "tmp_output.md"
	defer os.Remove(tmpInput)
	defer os.Remove(tmpOutput)

	err = os.WriteFile(tmpInput, []byte(input), 0o644)
	if err != nil {
		t.Fatalf("Failed to write input to temporary file: %v", err)
	}

	err = os.WriteFile(tmpOutput, output, 0o644)
	if err != nil {
		t.Fatalf("Failed to write output to temporary file: %v", err)
	}

	// Run diff command
	cmd := exec.Command("diff", "--unified", "--ignore-blank-lines", "--ignore-all-space", tmpInput, tmpOutput)
	diff, err := cmd.CombinedOutput()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if exitError.ExitCode() != 1 { // diff returns 1 if files are different, which we expect
				t.Errorf("diff command failed with exit code %d: %s", exitError.ExitCode(), string(diff))
			}
		} else {
			t.Errorf("Failed to run diff command: %v", err)
		}
	}

	if len(diff) > 0 {
		t.Errorf("Differences found between input and output:\n%s", string(diff))
	}
}
