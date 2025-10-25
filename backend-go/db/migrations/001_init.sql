-- 001_init.sql
CREATE EXTENSION IF NOT EXISTS postgis;

-- Users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    handicap REAL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Courses
CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    location GEOGRAPHY(Point, 4326),
    total_holes INT DEFAULT 18,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Holes
CREATE TABLE holes (
    id SERIAL PRIMARY KEY,
    course_id INT REFERENCES courses(id) ON DELETE CASCADE,
    hole_number INT NOT NULL,
    par INT NOT NULL,
    yardage INT,
    green_center GEOGRAPHY(Point, 4326),
    fairway_shape GEOGRAPHY(POLYGON, 4326),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Rounds
CREATE TABLE rounds (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    course_id INT REFERENCES courses(id),
    date_played DATE NOT NULL DEFAULT CURRENT_DATE,
    weather JSONB,
    total_score INT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Shots
CREATE TABLE shots (
    id SERIAL PRIMARY KEY,
    round_id INT REFERENCES rounds(id) ON DELETE CASCADE,
    hole_id INT REFERENCES holes(id),
    shot_number INT NOT NULL,
    club TEXT,
    start_point GEOGRAPHY(Point, 4326),
    end_point GEOGRAPHY(Point, 4326),
    distance REAL,
    result TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- AI Recommendations
CREATE TABLE ai_recommendations (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    round_id INT REFERENCES rounds(id),
    hole_id INT REFERENCES holes(id),
    recommendation_type TEXT,
    content JSONB,
    created_at TIMESTAMP DEFAULT NOW()
);
