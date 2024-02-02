INSERT INTO users (name, username, email, password) VALUES 
("Thomas Henrique", "thom", "thomas.santos063@gmail.com", "$2a$10$Ch10zvF.xWLYZotTkboWeuftET0R11lILTz2kspvDaaJtDqqu3.tC"),
("Alice Silva", "alice", "alice.silva@example.com", "$2a$10$Ch10zvF.xWLYZotTkboWeuftET0R11lILTz2kspvDaaJtDqqu3.tC"),
("Bob Oliveira", "bob", "bob.oliveira@example.com", "$2a$10$Ch10zvF.xWLYZotTkboWeuftET0R11lILTz2kspvDaaJtDqqu3.tC"),
("Charlie Santos", "charlie", "charlie.santos@example.com", "$2a$10$Ch10zvF.xWLYZotTkboWeuftET0R11lILTz2kspvDaaJtDqqu3.tC"),
("David Costa", "david", "david.costa@example.com", "$2a$10$Ch10zvF.xWLYZotTkboWeuftET0R11lILTz2kspvDaaJtDqqu3.tC"),
("Eva Lima", "eva", "eva.lima@example.com", "$2a$10$Ch10zvF.xWLYZotTkboWeuftET0R11lILTz2kspvDaaJtDqqu3.tC");

INSERT INTO followers (userId, followerId) VALUES 
(1, 2),
(1, 3),
(3, 1),
(4, 6),
(5, 6),
(6, 5); 

INSERT INTO posts (title, content, authorId) VALUES
('First Post by Thomas', 'This is the content of the first post by Thomas Henrique.', 1),
('First Post by Alice', 'This is the content of the first post by Alice Silva.', 2),
('First Post by Bob', 'This is the content of the first post by Bob Oliveira.', 3),
('First Post by Charlie', 'This is the content of the first post by Charlie Santos.', 4),
('First Post by David', 'This is the content of the first post by David Costa.', 5),
('First Post by Eva', 'This is the content of the first post by Eva Lima.', 6);


