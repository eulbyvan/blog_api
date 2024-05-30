-- Create the posts table
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    status VARCHAR(50) CHECK (status IN ('draft', 'publish')) NOT NULL,
    publish_date TIMESTAMP
);

-- Create the tags table
CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    label VARCHAR(100) UNIQUE NOT NULL
);

-- Create the junction table for the many-to-many relationship
CREATE TABLE post_tags (
    post_id INT REFERENCES posts(id) ON DELETE CASCADE,
    tag_id INT REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (post_id, tag_id)
);