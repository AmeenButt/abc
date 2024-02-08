-- name: CreateUser :one
INSERT INTO
    users(username, password, email)
VALUES
($1, $2, $3) RETURNING *;

-- name: GetUsers :many
SELECT * FROM users ORDER BY id;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: UpdateUserLastLogin :one
UPDATE users SET last_login = $2 WHERE id = $1 RETURNING *;

-- name: UpdateUsername :one
UPDATE users SET username = $2 WHERE id = $1 RETURNING *;

-- name: UpdateUserEmail :one
UPDATE users SET email = $2 WHERE id = $1 RETURNING *;

-- name: UpdatePassword :one
UPDATE users SET password = $2 WHERE id = $1 RETURNING *;
/*  
    USER PROFILE
*/

-- name: CreateUserProfile :one
INSERT INTO
    user_profile(user_id, phone_number, address)
VALUES
($1, $2, $3) RETURNING *;

-- name: GetUserProfileInfo :one
SELECT * FROM user_profile WHERE user_id = $1;

-- name: GetUserIdFromPhone :one
SELECT user_id FROM user_profile WHERE phone_number = $1;

-- name: VerifyUser :one
UPDATE user_profile SET is_verified = $2 WHERE user_id = $1 RETURNING *;

-- name: UpdateUserPhone :one
UPDATE user_profile SET phone_number = $2 WHERE user_id = $1 RETURNING *;

-- name: UpdateUserAddress :one
UPDATE user_profile SET address = $2 WHERE user_id = $1 RETURNING *;


/*
    JOINS
*/

-- name: GetAllUsersData :many
SELECT * FROM users INNER JOIN user_profile ON users.id = user_profile.user_id;


-- name: GetSingleUserData :one
SELECT * FROM users INNER JOIN user_profile ON users.id = user_profile.user_id WHERE users.id = $1;

-- name: LoginWithUserId :one
SELECT * FROM users INNER JOIN user_profile ON users.id = user_profile.user_id 
WHERE users.id = $1 AND users.password = $2;

-- name: FindUser :one
SELECT * FROM users INNER JOIN user_profile ON users.id = user_profile.user_id 
WHERE users.username = $1 OR users.email = $1 OR user_profile.phone_number = $1;