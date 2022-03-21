package models

import (
	"ushas/database"

	"gorm.io/gorm"
)

// Groups : Combination of task ID and condition ID.
type Groups struct {
	ID int `gorm:"unique;not null;column:id"`

	// GroupID : The ID assigned to the pair of "task IDs" and "condition ID"
	GroupID int `gorm:"not null;column:group_id"`

	// TaskID : The ID of Task
	TaskID int `gorm:"not null;column:task_id"`

	// ConditionID : The ID of Condition
	ConditionID int `gorm:"not null;column:condition_id"`
}

// GroupCounts : How many users are allocated to each group.
type GroupCounts struct {
	// GroupID : The ID assigned to the pair of "task IDs" and "condition ID"
	GroupID int `gorm:"unique;not nulll;column:group_id" json:"groupId"`

	// Count : Shows how many users are assigned to this group.
	Count int `gorm:"not null;column:count" json:"count"`
}

// GetLeastGroupID : Get one of the task that least users are allocated.
func GetLeastGroupID() (int, error) {
	gc := new(GroupCounts)
	db := database.GetDB()
	db.Transaction(func(tx *gorm.DB) error {
		subquery := tx.Table("group_counts").Select("MIN(counts)")
		if err := tx.Where("count = (?)", subquery).First(gc).Error; err != nil {
			return translateGormError(err, "Failed to fetch least count", nil)
		}
		gc.Count++
		if err := tx.Save(gc).Error; err != nil {
			return translateGormError(err, "Failed to update Task count", nil)
		}
		return nil
	})

	return gc.GroupID, nil
}

// GetAllocationByGroupID : Get a record (task IDs and condition ID) from table.
func GetAllocationByGroupID(groupID int) (*[]Groups, error) {
	gs := new([]Groups)
	db := database.GetDB()
	if err := db.Where("group_id = ?", groupID).Find(gs).Error; err != nil {
		return nil, translateGormError(err, "Failed to fetch task and condition group", groupID)
	}
	return gs, nil
}

// ListGroups : Get all records from table.
func ListGroups() ([]Groups, error) {
	groups := []Groups{}
	db := database.GetDB()
	if err := db.Find(&groups).Error; err != nil {
		return groups, translateGormError(err, "Failed to fetch all task and condition groups", nil)
	}
	return groups, nil
}

// UpdateGroup : Update a record with given parameters.
func UpdateGroup(g *Groups) error {
	db := database.GetDB()
	if err := db.Save(g).Error; err != nil {
		return translateGormError(err, "Failed to update task and condition group", g)
	}
	return nil
}

// DeleteGroup : Delete a record with given ID from table.
func DeleteGroup(conditionID int) error {
	db := database.GetDB()
	if err := db.Delete(&Groups{}, conditionID).Error; err != nil {
		return translateGormError(err, "Failed to delete task and condition group", conditionID)
	}
	return nil
}
