create keyspace IF NOT EXISTS profile with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
create table IF NOT EXISTS profile.user_profile_table(
   id uuid,
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
   updated_at timestamp, PRIMARY KEY(email));
