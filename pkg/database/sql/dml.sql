INSERT INTO posts (title, content, status, publish_date)
VALUES
    ('First Post', 'This is the content of the first post', 'publish', '2023-01-01 10:00:00'),
    ('Second Post', 'This is the content of the second post', 'draft', NULL);

INSERT INTO tags (label)
VALUES
    ('Go'),
    ('API'),
    ('PostgreSQL');

INSERT INTO post_tags (post_id, tag_id)
VALUES
    ((SELECT id FROM posts WHERE title = 'First Post'), (SELECT id FROM tags WHERE label = 'Go')),
    ((SELECT id FROM posts WHERE title = 'First Post'), (SELECT id FROM tags WHERE label = 'API')),
    ((SELECT id FROM posts WHERE title = 'Second Post'), (SELECT id FROM tags WHERE label = 'PostgreSQL'));