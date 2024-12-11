create table personalities(
    id serial primary key,
    name varchar,
    history varchar
);

INSERT INTO personalities (id, name, history) VALUES
(1, 'Ada Lovelace', 'Pioneira da programação e autora do primeiro algoritmo.'),
(2, 'Alan Turing', 'Fundador da ciência da computação e da criptoanálise moderna.');