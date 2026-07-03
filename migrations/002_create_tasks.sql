CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    due_date TIMESTAMP NOT NULL,
    user_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    
    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);
