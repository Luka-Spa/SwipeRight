CREATE TABLE USER_PROFILE(
   id int PRIMARY KEY,
   email text,
   profile_image_url text,
   first_name text,
   last_name text,
   gender text,
   bio text,
   location text,
   anthem text,
   school text,
   created_at timestamp,
   updated_at timestamp,
   );

INSERT INTO USER_PROFILE (id, first_name, last_name)
VALUES (0, 'Test', 'Person');
