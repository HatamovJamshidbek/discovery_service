CREATE TYPE genre AS ENUM ('rock', 'pop', 'jazz', 'classical', 'electronic', 'hip_hop', 'other');

CREATE TABLE composition_metadata (
                                      composition_id uuid,
                                      genre genre,
                                      tags TEXT[],
                                      listen_count INTEGER DEFAULT 0,
                                      like_count INTEGER DEFAULT 0
);

CREATE TABLE user_interactions (
                                   id uuid primary key  default gen_random_uuid(),
                                   user_id uuid ,
                                   composition_id uuid,
                                   interaction_type VARCHAR(20),
                                   created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
