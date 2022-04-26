create keyspace IF NOT EXISTS recommendation with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
CREATE TYPE IF NOT EXISTS recommendation.user_preferences (
  visible boolean,
  age_min int,
  age_max int,
  gender_preference int,
  max_distance int);
create table IF NOT EXISTS recommendation.user_table(
   id uuid,
   profile_image_url text,
   first_name text,
   gender text,
   bio text,
   location text,
   anthem text,
   school text,
   preferences FROZEN<user_preferences>, PRIMARY KEY(id))
