CREATE TABLE users (
    id INTEGER unsigned NOT NULL AUTO_INCREMENT,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE tasks (
    id INTEGER unsigned NOT NULL AUTO_INCREMENT,
    user_id INTEGER unsigned NOT NULL, 
    title VARCHAR(20) NOT NULL,
    description VARCHAR(20) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) 
        REFERENCES users(id)
        ON DELETE CASCADE
);
