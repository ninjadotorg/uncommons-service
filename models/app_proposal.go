package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

var (
	// AppProposals ...
	AppProposals map[int64]*AppProposal
)

// AppProposal ...
type AppProposal struct {
	ID              int64
	AuthorID        int64
	Name            string `orm:"text"`
	Description     string `orm:"text"`
	DescWhyWeLove   string `orm:"text"`
	DescWhyYouBuild string `orm:"text"`
	DescWhatCool    string `orm:"text"`
	HowFar          string `orm:"text"`
	ShippingDate    time.Time
	Amount          int64
	AboutURL        string `orm:"text"`
}

func init() {
	orm.RegisterModel(new(AppProposal))
	orm.RegisterDataBase("default", "mysql", "root:root@/localhost?charset=utf8", 30)
}

// AddOne ...
func AddOne(proposal AppProposal) (proposalID int64) {
	return proposal.ID
}

// GetOne ...
func GetOne(proposalID int64) (proposal *AppProposal, err error) {
	if v, ok := AppProposals[proposalID]; ok {
		return v, nil
	}
	return nil, errors.New("ObjectId Not Exist")
}

// GetAll ...
func GetAll() map[int64]*AppProposal {
	return AppProposals
}

// Update ...
func Update(proposalID int64) (err error) {
	// if v, ok := Objects[ObjectId]; ok {
	// 	v.Score = Score
	// 	return nil
	// }
	return errors.New("ObjectId Not Exist")
}

// Delete ...
func Delete(proposalID int64) {
	delete(AppProposals, proposalID)
}
