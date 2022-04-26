create keyspace IF NOT EXISTS profile with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
CREATE TYPE IF NOT EXISTS profile.user_preferences (
  visible boolean,
  age_min int,
  age_max int,
  gender_preference text,
  max_distance int);
CREATE TYPE IF NOT EXISTS profile.location (
  city text,
  lat float,
  long float);
create table IF NOT EXISTS profile.user_profile_table(
  id uuid,
  email text,
  profile_image_url text,
  first_name text,
  last_name text,
  gender text,
  bio text,
  location FROZEN<location>,
  anthem text,
  school text,
  preferences FROZEN<user_preferences>,
  created_at timestamp,
  updated_at timestamp, PRIMARY KEY(email));
