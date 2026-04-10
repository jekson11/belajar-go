-- FindAllUserData
SELECT user_id, name, username, email, created_at, password
FROM learngo.td_user
ORDER BY username ASC;