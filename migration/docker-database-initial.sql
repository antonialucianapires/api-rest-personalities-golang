create table personalidades(
    id serial primary key,
    nome varchar,
    historia varchar
);

INSERT INTO personalidades (id, nome, historia) VALUES
(1, 'Ada Lovelace', 'Pioneira da programação e autora do primeiro algoritmo.'),
(2, 'Alan Turing', 'Fundador da ciência da computação e da criptoanálise moderna.');