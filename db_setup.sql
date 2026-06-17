-- psql -U postgres to enter psql

-- create database
create database greenlight;

-- enter database
\c greenlight

-- create role
create role greenlight with login password 'pa55word';

-- create extension
create extension if not exists citext;