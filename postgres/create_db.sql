CREATE DATABASE vaalit;
CREATE USER vaalit WITH PASSWORD 'vaalit';
\c vaalit
CREATE TABLE Voting(name TEXT, id TEXT PRIMARY KEY, description TEXT, open INTEGER, ended INTEGER, visible INTEGER, hidden_ID SERIAL);
CREATE TABLE Token(value TEXT PRIMARY KEY, valid INTEGER, hidden_id SERIAL);
CREATE TABLE Candidate(name TEXT, id TEXT PRIMARY KEY, description TEXT, hidden_id SERIAL);
CREATE TABLE Availability(candidateId TEXT REFERENCES Candidate(id) ON DELETE CASCADE ON UPDATE CASCADE, votingId TEXT REFERENCES Voting(id) ON DELETE CASCADE ON UPDATE CASCADE, hidden_id SERIAL, PRIMARY KEY (candidateId, votingId));
CREATE TABLE Vote(id TEXT PRIMARY KEY, votingId TEXT REFERENCES Voting(id) ON DELETE CASCADE ON UPDATE CASCADE, candidateId TEXT REFERENCES Candidate(id) ON DELETE CASCADE ON UPDATE CASCADE, token TEXT REFERENCES Token(value) ON DELETE CASCADE ON UPDATE CASCADE, hidden_id SERIAL);
CREATE TABLE Mastertoken(value TEXT PRIMARY KEY);
GRANT ALL ON DATABASE vaalit TO vaalit;
GRANT ALL ON ALL TABLES IN SCHEMA public TO VAALIT;
GRANT ALL ON ALL SEQUENCES IN SCHEMA public TO VAALIT;
