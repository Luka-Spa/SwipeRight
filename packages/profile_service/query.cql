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


create keyspace profile with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
create type profile.user_profile ( email text,
   profile_image_url text,
   first_name text,
   last_name text,
   gender text,
   bio text,
   location text,
   anthem text,
   school text,
   created_at timestamp,
   updated_at timestamp,);
create table profile.user_profile_table(pk int, value frozen<user_profile>, PRIMARY KEY(pk));
INSERT INTO profile.user_profile_table (pk, value)
VALUES (0, {first_name: 'luka', last_name: 'testing', gender: 'male', bio: 'just testing', location: 'england', anthem: 'Englishman in New York', school: 'Fontys'});
