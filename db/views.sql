
/*
Visão materializada que tem como atributos o nome da playlist e a
quantidade de álbuns que a compõem.
*/

CREATE MATERIALIZED VIEW albuns_playlist AS
SELECT p.nome "Playlist", COUNT(DISTINCT P.cod_play) "Qtd albuns" FROM album
JOIN faixa          USING (cod_alb, cod_meio)
JOIN faixa_playlist USING (cod_faixa, cod_alb, cod_meio)
JOIN playlist p     USING (cod_play)
GROUP BY p.cod_play, p.nome
WITH NO DATA;


