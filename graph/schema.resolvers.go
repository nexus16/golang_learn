package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api/graph/generated"
	"api/graph/model"
	"api/internal/links"
	"api/internal/memberSkill"
	"api/internal/members"
	"api/internal/skills"
	"api/internal/users"
	"api/pkg/jwt"
	"context"
	"fmt"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	var link links.Link
	link.Title = input.Title
	link.Address = input.Address
	linkID := link.Save()
	return &model.Link{ID: strconv.FormatInt(linkID, 10), Title: link.Title, Address: link.Address}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	user.Create()
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) CreateSkill(ctx context.Context, input model.NewSkill) (*model.Skill, error) {
	var skill skills.Skill
	spew.Dump(skill)
	skill.Name = input.Name
	skill.Desc = input.Desc
	skillID := skill.Save()
	return &model.Skill{ID: strconv.FormatInt(skillID, 10), Name: skill.Name, Desc: skill.Desc}, nil
}

func (r *mutationResolver) CreateMemberSkill(ctx context.Context, input model.NewMemberSkill) (*model.MemberSkill, error) {
	var mb memberSkill.MemberSkill
	mb.MemberID = input.MemberID
	mb.SkillID = input.SkillID
	memberSkillID := mb.Save()
	return &model.MemberSkill{ID: strconv.FormatInt(memberSkillID, 10), MemberID: mb.MemberID, SkillID: mb.SkillID}, nil
}

func (r *mutationResolver) CreateMember(ctx context.Context, input model.NewMember) (*model.Member, error) {
	var member members.Member
	member.Name = input.Name
	member.Age = input.Age
	memberID := member.Save()
	return &model.Member{ID: strconv.FormatInt(memberID, 10), Name: member.Name, Age: member.Age}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct {
		// 1
		return "", &users.WrongUsernameOrPasswordError{}
	}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var resultLinks []*model.Link
	var dbLinks []links.Link
	dbLinks = links.GetAll()
	for _, link := range dbLinks {
		grahpqlUser := &model.User{
			Name: link.User.Username,
		}
		spew.Dump(grahpqlUser)

		resultLinks = append(resultLinks, &model.Link{ID: link.ID, Title: link.Title, Address: link.Address, User: grahpqlUser})
	}
	return resultLinks, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var resultUsers []*model.User
	var dbUsers []users.User
	dbUsers = users.GetAll()
	for _, user := range dbUsers {
		spew.Dump(model.User{ID: user.ID})
		resultUsers = append(resultUsers, &model.User{ID: user.ID, Name: user.Username})
	}
	return resultUsers, nil
}

func (r *queryResolver) Members(ctx context.Context) ([]*model.Member, error) {
	var resultMembers []*model.Member
	var dbMembers []members.Member
	dbMembers = members.GetAll()
	for _, member := range dbMembers {

		var dbSkills []skills.Skill
		dbSkills = skills.GetListSkillByMemberId(member.ID)
		var grahpqlSkill []*model.Skill
		for _, skill := range dbSkills {
			spew.Dump(skill)
			grahpqlSkill = append(grahpqlSkill, &model.Skill{ID: skill.ID, Name: skill.Name, Desc: skill.Desc})
		}
		// grahpqlSkill := &model.User{
		// 	Name: link.User.Username,
		// }
		resultMembers = append(resultMembers, &model.Member{ID: member.ID, Name: member.Name, Age: member.Age, Skill: grahpqlSkill})
	}
	return resultMembers, nil
}

func (r *queryResolver) Skills(ctx context.Context) ([]*model.Skill, error) {
	var resultSkills []*model.Skill
	var dbSkills []skills.Skill
	dbSkills = skills.GetAll()
	for _, skill := range dbSkills {
		resultSkills = append(resultSkills, &model.Skill{ID: skill.ID, Name: skill.Name, Desc: skill.Desc})
	}
	return resultSkills, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
