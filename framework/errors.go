package framework

import "fmt"

type (
	//
	// InventoryTooSmall is a type to be used as a custom error code when
	// the user's inventory capacity is too low for an operation
	//
	InventoryTooSmall struct {
		required float64
		capacity float64
	}
)

func (e *InventoryTooSmall) Error() string {
	return fmt.Sprintf("inventory too small. required: %v, capacity: %v", e.required, e.capacity)
}
