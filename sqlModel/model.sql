
CREATE TABLE books (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    price NUMERIC NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE category (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);


create table bookandcategory (
    id serial primary key,
    category_id UUID not null references category(id),
    book_id UUID not null references books(id)
)

insert into bookandcategory (category_id, book_id) values
('cd99d095-2011-40a8-9b46-35b754616e6f', '71e1a6b6-4a60-4bec-a94e-bf7937efb82d'),
('aaa1c7ed-2fea-427d-a9ce-57bb30f7ee40', '8f748179-51f6-4433-83f9-739bb6eb76c6'),
('cd99d095-2011-40a8-9b46-35b754616e6f', '8f748179-51f6-4433-83f9-739bb6eb76c6'),
('d6932c36-ddc4-4a4d-99bb-6cc19864df06', '1f76818c-b202-4366-a390-243c74c9e6ba'),
('7d15f744-05b4-4aaa-9237-0fd7d72f15c8', '6a04c1a2-2fd2-49b5-b507-c08ca80df44b');