-- TODO: answer here
CREATE TABLE presences (
    id INT not null,
    user_id INT not null,
    presence_date DATE not null,
    status VARCHAR(50) not null,
    location VARCHAR(255),
    description VARCHAR(255),
    image_presence VARCHAR(255),
    image_location VARCHAR(255)
)