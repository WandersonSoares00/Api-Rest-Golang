
delete from  gravadora          where true;
delete from  album              where true;
delete from  composicao         where true;
delete from  faixa              where true;
delete from  interprete         where true;
delete from  faixa_interprete   where true;
delete from  faixa_compositor   where true;
delete from  compositor         where true;
delete from  periodo_musical    where true;
delete from  faixa_playlist     where true;
delete from  playlist           where true;

insert into gravadora (cod_grav, nome, cidade, pais, end_homep)
values
    (50, 'Deutsche Grammophon', 'Hanôver', 'Alemanha', 'https://www.deutschegrammophon.com/de');
    

insert into album (cod_alb, cod_meio, cod_grav, nome, data_grav, pr_compra, pr_venda, meio)
values
        (1, 100, 50, 'Mozart: Requiem', '2020-04-28', 3.00, 40.00, 'CD');

insert into composicao
(cod_composicao, descricao, tipo)
values
    (90, 'Composição musical que tradicionalmente é uma missa fúnebre ou uma peça memorial em honra dos mortos.', 'Réquiem');

-- cod_alb * 10 + num_faixa
insert into faixa
(cod_faixa, cod_alb, cod_meio, cod_composicao, numero, descricao, tempo_exec, tipo_grav)
values
    (11, 1, 100, 90, 1, 'Requiem-Kyrie (Chor/Sopran)', '00:02:58', 'DDD'),
    (12, 1, 100, 90, 2, 'Dies Irae (Chor)', '00:02:06', 'DDD'),
    (13, 1, 100, 90, 3, 'Tuba Mirum (Sopran-Alt-Tenor-Bass)', '00:02:45', 'DDD');


insert into interprete
(cod_inter, nome, tipo)
values
    (70, 'Edith Mathis', 'Soprano'),
    (71, 'Julia Hamari', 'Ópera'),
    (72, 'Wiesław Ochman', 'Ópera'),
    (73, 'Karl Ridderbusch', 'Ópera');

insert into faixa_interprete
(cod_faixa, cod_alb, cod_meio, cod_inter)
values
    (11, 1, 100, 70),
    (13, 1, 100, 70),
    (13, 1, 100, 71),
    (13, 1, 100, 72),
    (13, 1, 100, 73);

insert into periodo_musical
(cod_pm, periodo, int_tempo)
values
    (300, 'clássico', daterange('1500-01-01', '1599-12-31', '[]'));

insert into compositor
(cod_compositor, cod_pm, nome, dt_nasc, dt_morte)
values
    (200, 300, 'Mozart', '1756-01-27', '1791-12-05');

insert into faixa_compositor
(cod_faixa, cod_alb, cod_meio, cod_compositor)
values
    (11, 1, 100, 200),
    (12, 1, 100, 200),
    (13, 1, 100, 200);

insert into playlist
(cod_play, nome, tempo_tot)
values
    (1, 'músicas-estudo', '00:05:00');

insert into faixa_playlist
(cod_faixa, cod_alb, cod_meio, cod_play, dt_ult_repr)
values
    (11, 1, 100, 1, null),
    (13, 1, 100, 1, null);
