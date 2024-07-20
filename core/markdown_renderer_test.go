package core

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessMarkdown(t *testing.T) {
	input := `---
title: Test Document
author: John Doe
---

# Heading 1

This is a paragraph.

## Heading 2

- List item 1
- List item 2

[Link](https://example.com)
`

	expected := `---
author: John Doe
title: Test Document
---

# Heading 1

This is a paragraph.

## Heading 2

- List item 1
- List item 2

[Link](https://example.com)
`

	ctx := context.Background()
	ctx = ContextWithShowAST(ctx, false)
	output, err := ProcessMarkdown(ctx, []byte(input))
	assert.NoError(t, err)
	assert.Equal(t, expected, string(output))
}
