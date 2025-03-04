package item

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

type Item struct {
	Content  string
	Priority int
	Done     bool
	position int
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (i *Item) PrettyP() string {
	if i.Priority == 1 {
		return "(1)"
	}
	if i.Priority == 3 {
		return "(3)"
	}
	return " "
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "X"
	}
	return ""
}

func (i *Item) PrettyDisplay() string {
	return fmt.Sprintln(i.Label() + "\t" + i.PrettyDone() + "\t" + i.PrettyP() + "\t" + i.Content + "\t")
}

type ByPri []Item

func (s ByPri) Len() int {
	return len(s)
}

func (s ByPri) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByPri) Less(i, j int) bool {
	if s[i].Done != s[j].Done {
		return !s[i].Done
	}

	if s[i].Priority == s[j].Priority {
		return s[i].Priority < s[j].Priority
	}

	return s[i].position < s[j].position
}

func NewItem(content string) *Item {
	if content == "" {
		return nil
	}

	i := Item{Content: content}
	return &i
}

func SaveItems(filename string, items []Item) error {
	data, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := os.ReadFile(filename)
	if errors.Is(err, fs.ErrNotExist) {
		return []Item{}, nil
	}
	if err != nil {
		return nil, err
	}

	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	for i := range items {
		items[i].position = i + 1
	}

	return items, nil
}
