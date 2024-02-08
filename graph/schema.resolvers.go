package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	db "github.com/ashbeelghouri/project1/db/sqlc"
	"github.com/ashbeelghouri/project1/graph/model"
	models "github.com/ashbeelghouri/project1/model"
	"github.com/ashbeelghouri/project1/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	// log.Printf("create user input: %v", input)
	// errors := make([]string, 0)
	conn := utils.ConnectDB(ctx)
	queries := db.New(conn)

	_, err := queries.GetUserByEmail(ctx, input.Email)

	if err == nil {
		log.Fatalf("User Email Exists: %v", err)
		return nil, errors.New("email address exists")
	}

	_, err = queries.GetUserByUsername(ctx, input.Username)

	if err == nil {
		log.Fatalf("Username exists: %v", err)
		return nil, errors.New("username exists")
	}

	user, err := queries.CreateUser(ctx, db.CreateUserParams{
		Username: input.Username,
		Email:    input.Email,
		Password: utils.Encrypt(input.Password),
	})

	if err != nil {
		log.Fatalf("unable to create user: %v", err)
		return nil, errors.New("unable to create user")
	}

	userProfile, err := queries.CreateUserProfile(ctx, db.CreateUserProfileParams{
		UserID:      user.ID,
		PhoneNumber: input.Profile.Phone,
		Address:     input.Profile.Address,
	})

	if err != nil {
		log.Fatalf("unable to create user profile: %v", err)
		return nil, errors.New("unable to create user profile")
	}

	return &model.User{
		Username: user.Username,
		Email:    user.Email,
		Password: &user.Password,
		UserProfile: &model.UserProfile{
			PhoneNumber: &userProfile.PhoneNumber,
			Address:     &userProfile.Address,
			IsVerified:  &userProfile.IsVerified,
		},
		ID:        int(user.ID),
		LastLogin: &user.LastLogin.Time,
		CreatedAt: &user.CreatedAt.Time,
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	conn := utils.ConnectDB(ctx)
	queries := db.New(conn)
	userid, err := strconv.ParseInt(input.ID, 10, 32)

	if err != nil {
		log.Printf("unable to convert the id to integer!")
		return nil, errors.New("system error")
	}

	if input.Email != nil {
		updateEmailInput := db.UpdateUserEmailParams{
			Email: *input.Email,
			ID:    int32(userid),
		}
		_, err := queries.UpdateUserEmail(ctx, updateEmailInput)

		if err != nil {
			log.Printf("error while updating the email address of the user: %v", err)
			return nil, errors.New("can not update email")
		}
	}

	// log.Printf("user input here is : %v", updateUserInput)

	if input.Username != nil {
		updateUserInput := db.UpdateUsernameParams{
			Username: *input.Username,
			ID:       int32(userid),
		}
		_, err := queries.UpdateUsername(ctx, updateUserInput)

		if err != nil {
			log.Printf("error while updating the username of the user: %v", err)
			return nil, errors.New("can not update username")
		}
	}

	if input.UserProfile != nil {
		if input.UserProfile.Address != nil {
			updateUserAddress := db.UpdateUserAddressParams{
				Address: *input.UserProfile.Address,
				UserID:  int32(userid),
			}
			_, err := queries.UpdateUserAddress(ctx, updateUserAddress)

			if err != nil {
				log.Printf("error while updating the user address: %v", err)
				return nil, errors.New("can not update user's address")
			}
		}

		if input.UserProfile.PhoneNumber != nil {
			updateUserPhoneNumber := db.UpdateUserPhoneParams{
				PhoneNumber: *input.UserProfile.PhoneNumber,
				UserID:      int32(userid),
			}
			_, err := queries.UpdateUserPhone(ctx, updateUserPhoneNumber)

			if err != nil {
				log.Printf("error while updating the user phone number: %v", err)
				return nil, errors.New("can not update user's phone number")
			}
		}

		if input.UserProfile.IsVerified != nil {
			verifyUserInput := db.VerifyUserParams{
				IsVerified: *input.UserProfile.IsVerified,
				UserID:     int32(userid),
			}
			_, err := queries.VerifyUser(ctx, verifyUserInput)

			if err != nil {
				log.Printf("error while updating the user verification status: %v", err)
				return nil, errors.New("unable to verify user")
			}
		}
	}

	userData, err := queries.GetSingleUserData(ctx, int32(userid))

	if err != nil {
		log.Printf("error while getting user data to show for updation: %v", err)
		return nil, errors.New("unable to fetch the user from database")
	}

	return &model.User{
		Username:  userData.Username,
		Email:     userData.Email,
		ID:        int(userData.ID),
		CreatedAt: &userData.CreatedAt.Time,
		LastLogin: &userData.LastLogin.Time,
		UserProfile: &model.UserProfile{
			Address:     &userData.Address,
			PhoneNumber: &userData.PhoneNumber,
			IsVerified:  &userData.IsVerified,
		},
	}, nil
}

// UpdateUserProfile is the resolver for the updateUserProfile field.
func (r *mutationResolver) UpdateUserProfile(ctx context.Context, input model.UpdateUserProfileInput, id string) (*model.User, error) {
	conn := utils.ConnectDB(ctx)
	queries := db.New(conn)
	userid, err := strconv.ParseInt(id, 10, 32)

	if err != nil {
		log.Printf("unable to convert the id to integer!")
		return nil, errors.New("system error")
	}
	if input.Address != nil {
		updateUserAddress := db.UpdateUserAddressParams{
			Address: *input.Address,
			UserID:  int32(userid),
		}
		_, err := queries.UpdateUserAddress(ctx, updateUserAddress)

		if err != nil {
			log.Printf("error while updating the user address: %v", err)
			return nil, errors.New("can not update user's address")
		}
	}

	if input.PhoneNumber != nil {
		updateUserPhoneNumber := db.UpdateUserPhoneParams{
			PhoneNumber: *input.PhoneNumber,
			UserID:      int32(userid),
		}
		_, err := queries.UpdateUserPhone(ctx, updateUserPhoneNumber)

		if err != nil {
			log.Printf("error while updating the user phone number: %v", err)
			return nil, errors.New("can not update user's phone number")
		}
	}

	if input.IsVerified != nil {
		verifyUserInput := db.VerifyUserParams{
			IsVerified: *input.IsVerified,
			UserID:     int32(userid),
		}
		_, err := queries.VerifyUser(ctx, verifyUserInput)

		if err != nil {
			log.Printf("error while updating the user verification status: %v", err)
			return nil, errors.New("unable to verify user")
		}
	}

	userData, err := queries.GetSingleUserData(ctx, int32(userid))

	if err != nil {
		log.Printf("error while getting user data to show for updation: %v", err)
		return nil, errors.New("unable to fetch the user from database")
	}

	return &model.User{
		Username:  userData.Username,
		Email:     userData.Email,
		ID:        int(userData.ID),
		CreatedAt: &userData.CreatedAt.Time,
		LastLogin: &userData.LastLogin.Time,
		UserProfile: &model.UserProfile{
			Address:     &userData.Address,
			PhoneNumber: &userData.PhoneNumber,
			IsVerified:  &userData.IsVerified,
		},
	}, nil
}

// LoginUser is the resolver for the loginUser field.
func (r *mutationResolver) LoginUser(ctx context.Context, input *model.LoginUsernameInput) (*model.LoginOutput, error) {
	conn := utils.ConnectDB(ctx)
	queries := db.New(conn)

	user, err := queries.FindUser(ctx, input.Username)
	if err != nil {
		return nil, errors.New("User Found")
	}

	if utils.CheckPasswd(user.Password, input.Password) {
		// convert user info to jwt token
		updatedUser, err := queries.UpdateUserLastLogin(ctx, db.UpdateUserLastLoginParams{
			LastLogin: pgtype.Timestamp{Time: time.Now(), Valid: true},
			ID:        user.ID,
		})
		if err != nil {
			log.Printf("unable to update the new login time: %v", err)
		}
		userTokenClaim := models.UserToken{
			ID:        int(user.ID),
			Username:  user.Username,
			Email:     user.Email,
			LastLogin: updatedUser.LastLogin.Time,
			UserProfile: models.UserProfile{
				Address:     user.Address,
				PhoneNumber: user.PhoneNumber,
				IsVerified:  user.IsVerified,
			},
		}

		token, err := utils.CreateUserToken(userTokenClaim, "usersecret")
		if err != nil {
			return nil, errors.New("unable to create token")
		}

		return &model.LoginOutput{
			Token: token,
			User: &model.User{
				ID:        int(user.ID),
				Username:  user.Username,
				Email:     user.Email,
				LastLogin: &updatedUser.LastLogin.Time,
				UserProfile: &model.UserProfile{
					Address:     &user.Address,
					PhoneNumber: &user.PhoneNumber,
				},
			},
		}, nil
	}

	return nil, errors.New("username or password is incorrect")
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	conn := utils.ConnectDB(ctx)
	queries := db.New(conn)

	users, err := queries.GetAllUsersData(ctx)

	if err != nil {
		return nil, errors.New("can not get users data from database")
	}
	allUsers := make([]*model.User, len(users))
	for i, v := range users {
		allUsers[i] = &model.User{
			Username:  v.Username,
			Email:     v.Email,
			ID:        int(v.ID),
			LastLogin: &v.LastLogin.Time,
			CreatedAt: &v.CreatedAt.Time,
			UserProfile: &model.UserProfile{
				Address:     &v.Address,
				PhoneNumber: &v.PhoneNumber,
				IsVerified:  &v.IsVerified,
			},
		}
	}
	return allUsers, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	conn := utils.ConnectDB(ctx)
	queries := db.New(conn)
	userid, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		log.Printf("can not convert the user id: %v", userid)
		return nil, errors.New("can not convert the user id")
	}
	user, err := queries.GetSingleUserData(ctx, int32(userid))
	if err != nil {
		return nil, errors.New("can not get user data")
	}
	return &model.User{
		Username: user.Username,
		Email:    user.Email,
		ID:       int(user.ID),
		UserProfile: &model.UserProfile{
			Address:     &user.Address,
			PhoneNumber: &user.PhoneNumber,
			IsVerified:  &user.IsVerified,
		},
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }