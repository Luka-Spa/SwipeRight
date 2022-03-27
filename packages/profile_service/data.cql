create keyspace IF NOT EXISTS profile with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
create type IF NOT EXISTS profile.user_profile ( email text,
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
create table IF NOT EXISTS profile.user_profile_table(pk int, value frozen<user_profile>, PRIMARY KEY(pk));
INSERT INTO profile.user_profile_table (pk, value)
VALUES (0, {first_name: 'luka', last_name: 'testing', gender: 'male', bio: 'just testing', location: 'england', anthem: 'Englishman in New York', school: 'Fontys'});
