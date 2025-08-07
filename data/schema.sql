PRAGMA foreign_keys=1;

CREATE TABLE IF NOT EXISTS libraries (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS groups (
    library INTEGER NOT NULL REFERENCES libraries (id),
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS sagas (
    "group" INTEGER NOT NULL REFERENCES groups (id),
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS books (
    isbn CHAR(13) PRIMARY KEY,

    title TEXT NOT NULL,
    authors TEXT NOT NULL, -- Separated by newlines
    publisher TEXT NOT NULL,

    lang CHAR(2) NOT NULL,
    pages INTEGER,
    price TEXT,

    saga INTEGER REFERENCES sagas (id),
    number INTEGER,

    boughtShop TEXT,
    boughtDate DATE,
    startedDate DATE,
    finishedDate DATE,
    status STRING NOT NULL
);

CREATE TABLE IF NOT EXISTS thumbnails (
    isbn CHAR(13) REFERENCES books (isbn),
    data BLOB NOT NULL
);
