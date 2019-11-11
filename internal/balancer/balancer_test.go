package balancer

import "testing"

func compareCharacterGroup(t **testing.T, actual, expected *CharacterGroup) {
	if actual.Opening != expected.Opening {
		(*t).Fatalf("Opening: Expected %U but got %U.\n", expected.Opening, actual.Opening)
	}

	if actual.Closing != expected.Closing {
		(*t).Fatalf("Closing: Expected %U but got %U.\n", expected.Closing, actual.Closing)
	}
}

func compareCharacterEntry(t **testing.T, actual, expected *CharacterEntry) {
	if actual.Char != expected.Char {
		(*t).Fatalf("Rune: Expected %U but got %U.\n", expected.Char, actual.Char)
	}

	compareCharacterGroup(t, actual.Group, expected.Group)
}

func compareCharacterGroupStack(t **testing.T, actual, expected *CharacterGroupStack) {
	if len(actual.array) != len(expected.array) {
		(*t).Fatalf("Stack: Expected array to be %v but got %v.\n", expected.array, actual.array)
	}

	for i, v := range actual.array {
		if expected.array[i] != v {
			(*t).Fatalf("Stack: Expected array to be %v but got %v.\n", expected.array, actual.array)
		}
	}
}

func compareBoolean(t **testing.T, actual, expected bool) {
	if actual != expected {
		(*t).Fatalf("Expected %t but got %t.\n", expected, actual)
	}
}

func TestNewCharacterGroup(t *testing.T) {
	actual := CurlyBracketsGroup
	expected := &CharacterGroup{Opening: CurlyBracketsGroup.Opening, Closing: CurlyBracketsGroup.Closing}
	compareCharacterGroup(&t, actual, expected)
}

func TestNewCharacterEntry(t *testing.T) {
	actual := NewCharacterEntry(CurlyBracketsGroup.Opening, CurlyBracketsGroup)
	expected := &CharacterEntry{Char: CurlyBracketsGroup.Opening, Group: CurlyBracketsGroup}
	compareCharacterEntry(&t, actual, expected)
}

func TestNewCharacterGroupStack(t *testing.T) {
	actual := NewCharacterGroupStack()
	expected := &CharacterGroupStack{}
	compareCharacterGroupStack(&t, actual, expected)
}

func TestCharacterEntry_IsClosingCharacter(t *testing.T) {
	entry := NewCharacterEntry(ParenthesisGroup.Closing, ParenthesisGroup)
	actual := entry.IsClosingCharacter()
	compareBoolean(&t, actual, true)

	entry = NewCharacterEntry(ParenthesisGroup.Opening, ParenthesisGroup)
	actual = entry.IsClosingCharacter()
	compareBoolean(&t, actual, false)
}

func TestCharacterGroupStack_Push(t *testing.T) {
	entries := []*CharacterEntry{
		NewCharacterEntry(CurlyBracketsGroup.Opening, CurlyBracketsGroup),
		NewCharacterEntry(AngleBracketsGroup.Opening, AngleBracketsGroup),
		NewCharacterEntry(SquareBracketsGroup.Opening, SquareBracketsGroup),
	}
	stack := NewCharacterGroupStack()
	expected := NewCharacterGroupStack()
	expected.array = make([]*CharacterEntry, 3)

	for i, v := range entries {
		stack.Push(v)
		expected.array[i] = v
	}

	// All pushed entries should be there.
	compareCharacterGroupStack(&t, stack, expected)

	// Now let's see what happens when we add a closing entry
	closingEntry := NewCharacterEntry(0x5D, SquareBracketsGroup)
	stack.Push(closingEntry)
	expected.array = expected.array[:2]
	compareCharacterGroupStack(&t, stack, expected)
}

func TestCharacterGroupStack_Pop(t *testing.T) {
	stack := NewCharacterGroupStack()
	_, err := stack.Pop()

	// Stack has nothing
	if err == nil {
		t.Fatalf("Was expecting error but got none.\n")
	}

	entries := []*CharacterEntry{
		NewCharacterEntry(CurlyBracketsGroup.Opening, CurlyBracketsGroup),
		NewCharacterEntry(AngleBracketsGroup.Opening, AngleBracketsGroup),
		NewCharacterEntry(SquareBracketsGroup.Opening, SquareBracketsGroup),
	}
	expectedStack := NewCharacterGroupStack()
	expectedStack.array = make([]*CharacterEntry, 3)

	for i, v := range entries {
		stack.Push(v)
		expectedStack.array[i] = v
	}

	actual, _ := stack.Pop()
	expected := expectedStack.array[2]
	expectedStack.array = expectedStack.array[:2]

	// Stack removed and returned last element
	compareCharacterEntry(&t, actual, expected)
	compareCharacterGroupStack(&t, stack, expectedStack)
}

func TestCharacterGroupStack_IsBalanced(t *testing.T) {
	// Empty Stack
	emptyStack := NewCharacterGroupStack()
	compareBoolean(&t, emptyStack.IsBalanced(), true)

	stack := NewCharacterGroupStack()
	entries := []*CharacterEntry{
		NewCharacterEntry(CurlyBracketsGroup.Opening, CurlyBracketsGroup),
		NewCharacterEntry(AngleBracketsGroup.Opening, AngleBracketsGroup),
		NewCharacterEntry(SquareBracketsGroup.Opening, SquareBracketsGroup),
		NewCharacterEntry(SquareBracketsGroup.Closing, SquareBracketsGroup),
		NewCharacterEntry(AngleBracketsGroup.Closing, AngleBracketsGroup),
		NewCharacterEntry(CurlyBracketsGroup.Closing, CurlyBracketsGroup),
	}

	for i, v := range entries {
		stack.Push(v)

		// Check midway
		if i == 3 {
			compareBoolean(&t, stack.IsBalanced(), false)
		}
	}

	// Check when complete with all pushes!
	compareBoolean(&t, stack.IsBalanced(), true)
}
