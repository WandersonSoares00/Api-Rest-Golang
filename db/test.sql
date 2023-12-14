
delete from  gravadora          where true;
delete from  album              where true;
delete from  composicao         where true;
delete from  faixa              where true;
delete from  interprete         where true;
delete from  faixa_interprete   where true;
delete from  faixa_compositor   where true;
delete from  compositor         where true;
delete from  faixa_playlist     where true;
delete from  playlist           where true;

insert into gravadora (cod_grav, nome, cidade, pais, end_homep)
values
    (1, 'Deutsche Grammophon', 'Hanôver', 'Alemanha', 'https://www.deutschegrammophon.com/de'),
    (2, 'DGC Records', 'Santa Mônica', 'EUA', 'https://www.discogs.com'),
    (3, 'Sony Classical', 'Nova Iorque', 'EUA', 'https://www.sonyclassical.com');

insert into album (cod_alb, meio, cod_grav, nome, data_grav, pr_compra, pr_venda)
values
        (1, 'CD',    1, 'Mozart: Requiem',              '2020-04-28', 3.00, 40.00),
        (2, 'CD',    2, 'Nevermind',                    '1991-09-24', 4.00, 30.00),
        (2, 'VINIL', 2, 'Nevermind',                    '2009-09-24', 7.00, 65.00),
        (3, 'CD',    3, 'Beethoven: Sinfonia No. 9',    '2021-03-15', 5.00, 50.00),
        (3, 'VINIL', 3, 'Beethoven: Sinfonia No. 9',    '2021-03-15', 8.00, 75.00),
        (5, 'DOWNLOAD', 1, 'Beethoven: Symphony No. 5', '2022-05-10', 2.99, 25.00),
        (6, 'DOWNLOAD', 3, 'Mozart: Piano Sonatas',     '2022-08-15', 4.50, 35.00);


insert into composicao
(cod_composicao, descricao, tipo)
values
    (1, 'Composição musical que tradicionalmente é uma missa fúnebre ou uma peça memorial em honra dos mortos.', 'Réquiem'),
    (2, 'O rock é um gênero amplamente popular e vasto, que reúne uma ampla gama de estilos diferentes.', 'Rock'),
    (3, 'A Nona Sinfonia de Beethoven, conhecida como a Sinfonia Coral, é uma das obras mais famosas do repertório clássico.', 'Sinfonia');


insert into faixa
(nro_faixa, cod_alb, meio, cod_composicao, descricao, tempo_exec, tipo_grav)
values
    ( 1, 1, 'CD', 1, 'Requiem-Kyrie (Chor/Sopran)',        '00:02:58', 'DDD'),
    ( 2, 1, 'CD', 1, 'Dies Irae (Chor)',                   '00:02:06', 'DDD'),
    ( 3, 1, 'CD', 1, 'Tuba Mirum (Sopran-Alt-Tenor-Bass)', '00:02:45', 'DDD'),
    ( 1, 2, 'CD', 2, 'Smells Like Teen Spirit',     	   '00:05:01', 'DDD'),
    ( 2, 2, 'CD', 2, 'In Bloom',                       	   '00:04:14', 'DDD'),
    ( 3, 2, 'CD', 2, 'Come as You Are',                    '00:03:38', 'DDD'),
    ( 4, 2, 'CD', 2, 'Breed',                              '00:03:03', 'DDD'),
    ( 5, 2, 'CD', 2, 'Lithium',                            '00:04:16', 'DDD'),
    ( 6, 2, 'CD', 2, 'Polly',                              '00:02:56', 'DDD'),
    ( 7, 2, 'CD', 2, 'Territorial Pissings',               '00:02:23', 'DDD'),
    ( 8, 2, 'CD', 2, 'Drain You',                          '00:03:43', 'DDD'),
    ( 9, 2, 'CD', 2, 'Lounge Act',                         '00:02:36', 'DDD'),
    (10, 2, 'CD', 2, 'Stay Away',                          '00:03:32', 'DDD'),
    (11, 2, 'CD', 2, 'On a Plain',                         '00:03:16', 'DDD'),
    (12, 2, 'CD', 2, 'Something in the Way',               '00:03:52', 'DDD'),
    (13, 2, 'CD', 2, 'Endless, Nameless',                  '00:06:44', 'DDD'),
    ( 1, 2, 'VINIL', 2, 'Smells Like Teen Spirit',     	   '00:05:01', NULL),
    ( 2, 2, 'VINIL', 2, 'In Bloom',                        '00:04:14', NULL),
    ( 3, 2, 'VINIL', 2, 'Come as You Are',                 '00:03:38', NULL),
    ( 4, 2, 'VINIL', 2, 'Breed',                           '00:03:03', NULL),
    ( 5, 2, 'VINIL', 2, 'Lithium',                         '00:04:16', NULL),
    ( 6, 2, 'VINIL', 2, 'Polly',                           '00:02:56', NULL),
    ( 7, 2, 'VINIL', 2, 'Territorial Pissings',            '00:02:23', NULL),
    ( 8, 2, 'VINIL', 2, 'Drain You',                       '00:03:43', NULL),
    ( 9, 2, 'VINIL', 2, 'Lounge Act',                      '00:02:36', NULL),
    (10, 2, 'VINIL', 2, 'Stay Away',                       '00:03:32', NULL),
    (11, 2, 'VINIL', 2, 'On a Plain',                      '00:03:16', NULL),
    (12, 2, 'VINIL', 2, 'Something in the Way',            '00:03:52', NULL),
    (13, 2, 'VINIL', 2, 'Endless, Nameless',               '00:06:44', NULL),
    (1,  3, 'CD', 3, 'Allegro ma non troppo, un poco maestoso', '00:15:50', 'DDD'),
    (2,  3, 'CD', 3, 'Molto vivace',                        '00:10:50', 'DDD'),
    (3,  3, 'CD', 3, 'Adagio molto e cantabile',            '00:12:30', 'DDD'),
    (4,  3, 'CD', 3, 'Presto',                              '00:05:50', 'DDD'),
    (5,  3, 'CD', 3, 'Finale: Ode to Joy',                  '00:20:30', 'DDD'),
    (1,  3, 'VINIL', 3, 'Allegro ma non troppo, un poco maestoso', '00:15:50', NULL),
    (2,  3, 'VINIL', 3, 'Molto vivace',                     '00:10:50', NULL),
    (3,  3, 'VINIL', 3, 'Adagio molto e cantabile',         '00:12:30', NULL),
    (4,  3, 'VINIL', 3, 'Presto',                           '00:05:50', NULL),
    (5,  3, 'VINIL', 3, 'Finale: Ode to Joy',               '00:20:30', NULL),
    (1,  5, 'DOWNLOAD', 3, 'Allegro con brio',              '00:07:20', NULL),
    (2,  5, 'DOWNLOAD', 3, 'Andante con moto',              '00:09:15', NULL),
    (3,  5, 'DOWNLOAD', 3, 'Scherzo. Allegro',              '00:05:25', NULL),
    (4,  5, 'DOWNLOAD', 3, 'Allegro',                       '00:08:45', NULL),
    (5,  5, 'DOWNLOAD', 3, 'Piano Sonata No. 16 in C Major, K. 545', '00:04:25', NULL),
    (1,  6, 'DOWNLOAD', 1, 'Allegro',                        '00:06:45', NULL),
    (2,  6, 'DOWNLOAD', 1, 'Adagio un poco mosso',          '00:08:30',  NULL),
    (3,  6, 'DOWNLOAD', 1, 'Allegro',                       '00:05:15',  NULL),
    (4,  6, 'DOWNLOAD', 1, 'Allegro',                       '00:11:10',  NULL),
    (5,  6, 'DOWNLOAD', 1, 'Piano Sonata No. 14 in C-sharp Minor, Op. 27, No. 2 Moonlight', '00:05:55', NULL);



insert into interprete
(cod_inter, nome, tipo)
values
    (1, 'Edith Mathis',     'Soprano'),
    (2, 'Julia Hamari',     'Ópera'),
    (3, 'Wiesław Ochman',   'Ópera'),
    (4, 'Karl Ridderbusch', 'Ópera'),
    (5, 'Kurt Cobain',      'Tenor'),
    (6, 'Lang Lang',        'Pianista'),
    (7, 'Murray Perahia',   'Pianista');


insert into faixa_interprete
(nro_faixa, cod_alb, meio, cod_inter)
values
    (1, 1, 'CD', 1),
    (3, 1, 'CD', 1),
    (3, 1, 'CD', 2),
    (3, 1, 'CD', 3),
    (3, 1, 'CD', 4),
    (5, 5, 'DOWNLOAD', 6),
    (1, 6, 'DOWNLOAD', 7),
    (2, 6, 'DOWNLOAD', 7),
    (3, 6, 'DOWNLOAD', 7),
    (4, 6, 'DOWNLOAD', 7),
    (5, 6, 'DOWNLOAD', 6);


insert into compositor
(cod_compositor, cod_pm, nome, dt_nasc, dt_morte)
values
    (1, 4, 'Mozart', '1756-01-27', '1791-12-05'),
    (2, 6, 'Kurt Cobain', '1967-02-20', '1994-04-05'),
    (3, 4, 'Ludwig van Beethoven', '1770-12-16', '1827-03-26');


insert into faixa_compositor
(nro_faixa, cod_alb, meio, cod_compositor)
values
    (1, 1, 'CD', 1),
    (2, 1, 'CD', 1),
    (3, 1, 'CD', 1),
    ( 1, 2, 'CD', 2),
    ( 2, 2, 'CD', 2),
    ( 3, 2, 'CD', 2),
    ( 4, 2, 'CD', 2),
    ( 5, 2, 'CD', 2),
    ( 6, 2, 'CD', 2),
    ( 7, 2, 'CD', 2),
    ( 8, 2, 'CD', 2),
    ( 9, 2, 'CD', 2),
    (10, 2, 'CD', 2),
    (11, 2, 'CD', 2),
    (12, 2, 'CD', 2),
    (13, 2, 'CD', 2),
    ( 1, 2, 'VINIL', 2),
    ( 2, 2, 'VINIL', 2),
    ( 3, 2, 'VINIL', 2),
    ( 4, 2, 'VINIL', 2),
    ( 5, 2, 'VINIL', 2),
    ( 6, 2, 'VINIL', 2),
    ( 7, 2, 'VINIL', 2),
    ( 8, 2, 'VINIL', 2),
    ( 9, 2, 'VINIL', 2),
    (10, 2, 'VINIL', 2),
    (11, 2, 'VINIL', 2),
    (12, 2, 'VINIL', 2),
    (13, 2, 'VINIL', 2),
    (1, 3, 'CD', 3),
    (2, 3, 'CD', 3),
    (3, 3, 'CD', 3),
    (4, 3, 'CD', 3),
    (5, 3, 'CD', 3),
    (1, 3, 'VINIL', 3),
    (2, 3, 'VINIL', 3),
    (3, 3, 'VINIL', 3),
    (4, 3, 'VINIL', 3),
    (5, 3, 'VINIL', 3),
    (1, 5,  'DOWNLOAD', 3),
    (2, 5,  'DOWNLOAD', 3),
    (3, 5,  'DOWNLOAD', 3),
    (4, 5,  'DOWNLOAD', 3),
    (5, 5, 'DOWNLOAD', 1),
    (1, 6, 'DOWNLOAD', 3),
    (2, 6, 'DOWNLOAD', 3),
    (3, 6, 'DOWNLOAD', 3),
    (4, 6, 'DOWNLOAD', 3),
    (5, 6, 'DOWNLOAD', 1);

insert into playlist
(cod_play, nome, tempo_tot)
values
    (1, 'músicas-estudo', '00:05:00'),
    (2, 'Rock Classics', '00:45:00');

insert into faixa_playlist
(nro_faixa, cod_alb, meio, cod_play, dt_ult_repr)
values
    (1, 1, 'CD', 1, null),
    (3, 1, 'CD', 1, null),
    (1, 2, 'CD', 2, null),
    (2, 2, 'CD', 2, null),
    (3, 2, 'CD', 2, null),
    (4, 2, 'CD', 2, null),
    (5, 2, 'CD', 2, null),
    (1, 2, 'VINIL', 2, null),
    (2, 2, 'VINIL', 2, null),
    (3, 2, 'VINIL', 2, null),
    (4, 2, 'VINIL', 2, null),
    (5, 2, 'VINIL', 2, null);



