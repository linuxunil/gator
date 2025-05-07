
-- +goose up
CREATE TABLE feeds(
		name TEXT NOT NULL,
	url TEXT UNIQUE NOT NULL,
	user_id int NOT NULL,
	CONSTRAINT fk_user
	FOREIGN KEY(user_id)
	references users(id) 
	ON DELETE CASCADE

);
-- +goose down
DROP TABLE feeds;
