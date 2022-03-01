package models

import (
	"fmt"
	"ushas/database"

	"gorm.io/gorm"
)

type Groups struct {
	ID int `gorm:"unique;not null;column:id"`

	// GroupID : The ID assigned to the pair of "task IDs" and "condition ID"
	GroupID int `gorm:"not null;column:group_id"`

	// TaskID : The ID of Task
	TaskID int `gorm:"not null;column:task_id"`

	// ConditionID : The ID of Condition
	ConditionID int `gorm:"not null;column:condition_id"`
}

// GroupCounts : Struct for group count
type GroupCounts struct {
	// GroupID : The ID assigned to the pair of "task IDs" and "condition ID"
	GroupID int `gorm:"unique;not nulll;column:group_id" json:"groupId"`

	// Count : Shows how many users are assigned to this group.
	Count int `gorm:"not null;column:count" json:"count"`
}

func GetLeastGroupID() (int, error) {
	gc := new(GroupCounts)
	db := database.GetDB()
	db.Transaction(func(tx *gorm.DB) error {
		subquery := tx.Table("group_counts").Select("MIN(counts)")
		if err := tx.Where("count = (?)", subquery).First(gc).Error; err != nil {
			return RaiseInternalServerError(
				err,
				"Failed to get least count",
				"",
			)
		}
		gc.Count += 1
		if err := tx.Save(gc).Error; err != nil {
			return RaiseInternalServerError(
				err,
				"Failed to update Task count",
				"",
			)
		}
		return nil
	})

	return gc.GroupID, nil
}

func GetAllocationByGroupID(groupID int) (*[]Groups, error) {
	gs := new([]Groups)
	db := database.GetDB()
	if err := db.Where("group_id = ?", groupID).Find(gs).Error; err != nil {
		return nil, RaiseInternalServerError(err, "", "")
	}
	return gs, nil
}

func ListGroups() ([]Groups, error) {
	groups := []Groups{}
	db := database.GetDB()
	if err := db.Find(&groups).Error; err != nil {
		return groups, RaiseInternalServerError(
			err,
			"Failed to fetch all Group resources",
		)
	}
	return groups, nil
}

func UpdateGroup(g *Groups) error {
	db := database.GetDB()
	if err := db.Save(g).Error; err != nil {
		return RaiseInternalServerError(
			err,
			fmt.Sprintf("Failed to update Group resource of ID %d", g.ID),
			g,
		)
	}
	return nil
}

func DeleteGroup(conditionID int) error {
	db := database.GetDB()
	if err := db.Delete(&Groups{}, conditionID).Error; err != nil {
		return RaiseInternalServerError(
			err,
			fmt.Sprintf("Failed to delete Group resource of ID %d", conditionID),
		)
	}
	return nil
}
